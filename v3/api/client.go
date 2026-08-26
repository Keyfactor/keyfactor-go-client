// Copyright 2024 Keyfactor
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/Keyfactor/keyfactor-auth-client-go/auth_providers"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	// DefaultAPIPath is the default API path for Keyfactor Command.
	DefaultAPIPath = auth_providers.DefaultCommandAPIPath

	// EnvCommandHostname is the environment variable for the Keyfactor Command hostname.
	EnvCommandHostname = auth_providers.EnvKeyfactorHostName

	// EnvCommandAPI is the environment variable for the Keyfactor Command API path.
	EnvCommandAPI = auth_providers.EnvKeyfactorAPIPath

	// EnvCommandTimeout is the environment variable for the Keyfactor Command timeout.
	EnvCommandTimeout = auth_providers.EnvKeyfactorClientTimeout

	// EnvCommandUsername is the environment variable for the Keyfactor Command username.
	EnvCommandUsername = auth_providers.EnvKeyfactorUsername

	// EnvCommandPassword is the environment variable for the Keyfactor Command password.
	EnvCommandPassword = auth_providers.EnvKeyfactorPassword

	// EnvCommandDomain is the environment variable for the Keyfactor Command domain.
	EnvCommandDomain = auth_providers.EnvKeyfactorDomain

	// EnvCommandClientId is the environment variable for the Keyfactor Command client ID.
	EnvCommandClientId = auth_providers.EnvKeyfactorClientID

	// EnvCommandClientSecret is the environment variable for the Keyfactor Command client secret.
	EnvCommandClientSecret = auth_providers.EnvKeyfactorClientSecret

	// EnvCommandOAuthTokenUrl is the environment variable for the Keyfactor Command OAuth token URL.
	EnvCommandOAuthTokenUrl = auth_providers.EnvKeyfactorAuthTokenURL
)

type Client struct {
	AuthClient AuthConfig
	LoggerType string

	// httpClient caches the *http.Client returned by AuthClient.GetHttpClient()
	// so that sendRequest reuses a single underlying transport/connection pool
	// across requests instead of asking AuthClient to build a brand new one on
	// every call. Both CommandConfigOauth.GetHttpClient() and
	// CommandAuthConfigBasic.GetHttpClient() (in keyfactor-auth-client-go)
	// construct a fresh http.Transport per invocation, and that transport's
	// IdleConnTimeout is derived from the configured HttpClientTimeout - so
	// without this cache, every request opens its own connection pool whose
	// sockets linger for up to HttpClientTimeout before being reclaimed. This
	// was already true at the old fixed 60s default; plumbing a caller-supplied
	// ClientTimeout (see NewKeyfactorClient) just widens the window, so caching
	// here keeps that fix from amplifying a pre-existing resource leak.
	httpClient   *http.Client
	httpClientMu sync.Mutex
}

// getHttpClient returns the cached *http.Client if one has already been
// resolved for this Client, populating the cache on first use otherwise.
// This guarantees AuthClient.GetHttpClient() is invoked at most once per
// Client instance, so the transport (and its connection pool) is reused
// across requests. It is safe for concurrent use.
//
// Note this does not affect OAuth token refresh: the cached *http.Client's
// transport wraps an oauth2 TokenSource that is consulted (and refreshed as
// needed) on every RoundTrip, independent of how many times the *http.Client
// itself is reused.
func (c *Client) getHttpClient() (*http.Client, error) {
	c.httpClientMu.Lock()
	defer c.httpClientMu.Unlock()

	if c.httpClient != nil {
		return c.httpClient, nil
	}

	httpClient, err := c.AuthClient.GetHttpClient()
	if err != nil {
		return nil, err
	}
	c.httpClient = httpClient
	return httpClient, nil
}

// TerraformLogger wraps the tflog logging to handle Go's log messages with log level mapping.
type TerraformLogger struct {
	ctx context.Context
}

// Write implements the io.Writer interface to redirect Go logs to tflog with log levels.
func (w *TerraformLogger) Write(p []byte) (n int, err error) {
	// Clean up the message (optional, depending on the format you want)
	msg := string(p)
	msg = strings.TrimSpace(msg)

	// Parse the message for log level (optional enhancement)
	w.logWithLevel(msg)

	// Return the length of the written message and no error
	return len(p), nil
}

// logWithLevel processes the message to detect log levels and log to tflog accordingly.
func (w *TerraformLogger) logWithLevel(msg string) {
	msg = strings.TrimSpace(msg)

	switch {
	case strings.Contains(msg, "ERROR"):
		tflog.Error(w.ctx, strings.Trim(strings.Replace(msg, "ERROR", "", 1), "[] "))
	case strings.Contains(msg, "WARN"):
		tflog.Warn(w.ctx, strings.Trim(strings.Replace(msg, "WARN", "", 1), "[] "))
	case strings.Contains(msg, "DEBUG"):
		tflog.Debug(w.ctx, strings.Trim(strings.Replace(msg, "DEBUG", "", 1), "[] "))
	case strings.Contains(msg, "TRACE"):
		tflog.Trace(w.ctx, strings.Trim(strings.Replace(msg, "TRACE", "", 1), "[] "))
	default:
		tflog.Info(w.ctx, strings.Trim(strings.Replace(msg, "INFO", "", 1), "[] "))
	}
}

func initLogger(ctx *context.Context) {
	if ctx != nil {
		log.SetOutput(
			&TerraformLogger{
				ctx: *ctx,
			},
		)
	} else {
		log.SetOutput(io.Discard)
	}
}

// Define an interface that both CommandConfigOauth and CommandAuthConfigBasic implement
type AuthConfig interface {
	Authenticate() error
	GetHttpClient() (*http.Client, error)
	GetServerConfig() *auth_providers.Server
	GetCommandVersion() string
}

// NewKeyfactorClientWithAuth creates a Client with a pre-built AuthConfig, bypassing
// the Authenticate() network call. Used in unit tests with VCR cassettes.
func NewKeyfactorClientWithAuth(auth AuthConfig, ctx *context.Context) *Client {
	initLogger(ctx)
	return &Client{AuthClient: auth}
}

// NewKeyfactorClient creates a new Keyfactor client instance. A configured Client is returned with methods used to
// interact with Keyfactor.
func NewKeyfactorClient(cfg *auth_providers.Server, ctx *context.Context) (*Client, error) {
	initLogger(ctx)
	client := Client{}
	clientAuthType := cfg.GetAuthType()

	baseConfig := auth_providers.CommandAuthConfig{
		CommandHostName:   cfg.Host,
		CommandPort:       cfg.Port,
		CommandAPIPath:    cfg.APIPath,
		CommandCACert:     cfg.CACertPath,
		SkipVerify:        cfg.SkipTLSVerify,
		HttpClientTimeout: cfg.ClientTimeout,
	}

	if clientAuthType == "basic" {
		basicCfg := auth_providers.CommandAuthConfigBasic{
			CommandAuthConfig: baseConfig,
			Username:          cfg.Username,
			Password:          cfg.Password,
			Domain:            cfg.Domain,
		}
		aErr := basicCfg.Authenticate()
		if aErr != nil {
			return nil, aErr
		}
		httpClient, cErr := basicCfg.GetHttpClient()
		if cErr != nil {
			return nil, cErr
		}
		client.AuthClient = &basicCfg
		client.httpClient = httpClient
		return &client, nil
	} else if clientAuthType == "oauth" {
		oauthCfg := auth_providers.CommandConfigOauth{
			CommandAuthConfig: baseConfig,
			ClientID:          cfg.ClientID,
			ClientSecret:      cfg.ClientSecret,
			TokenURL:          cfg.OAuthTokenUrl,
			Audience:          cfg.Audience,
			Scopes:            cfg.Scopes,
			AccessToken:       cfg.AccessToken,
		}
		aErr := oauthCfg.Authenticate()
		if aErr != nil {
			return nil, aErr
		}
		httpClient, cErr := oauthCfg.GetHttpClient()
		if cErr != nil {
			return nil, cErr
		}
		client.AuthClient = &oauthCfg
		client.httpClient = httpClient
		return &client, nil
	} else {
		return nil, fmt.Errorf("unsupported auth type or authentication cfg: '%s'", clientAuthType)
	}
}

func logRequest(req *http.Request) error {
	// Read the request body
	if req == nil {
		log.Printf("[WARN] HTTP Request is nil")
		return nil
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	// Restore the request body so it can be read later
	req.Body = io.NopCloser(bytes.NewBuffer(body))

	// Create a struct to hold request data
	requestData := struct {
		Method  string              `json:"method"`
		URL     string              `json:"url"`
		Headers map[string][]string `json:"headers"`
		Body    string              `json:"body"`
	}{
		Method:  req.Method,
		URL:     req.URL.String(),
		Headers: req.Header,
		Body:    string(body),
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	// Log the JSON data
	log.Printf("[TRACE] HTTP Request: %s", jsonData)
	curlStr, err := requestToCurl(req)
	if err != nil {
		log.Printf("[ERROR] Error converting request to cURL: %s", err)
		return nil
	}
	log.Printf("[TRACE] cURL Request: %s", curlStr)
	log.Printf("[TRACE] cURL Request b64encoded: %s", base64.StdEncoding.EncodeToString([]byte(curlStr)))
	return nil
}

func requestToCurl(req *http.Request) (string, error) {
	var curlCommand strings.Builder

	// Start with the cURL command
	curlCommand.WriteString(fmt.Sprintf("curl -X %s ", req.Method))

	// Add the URL
	curlCommand.WriteString(fmt.Sprintf("%q ", req.URL.String()))

	// Add headers
	for name, values := range req.Header {
		for _, value := range values {
			curlCommand.WriteString(fmt.Sprintf("-H %q ", fmt.Sprintf("%s: %s", name, value)))
		}
	}

	// Add the body if it exists
	if req.Method == http.MethodPost || req.Method == http.MethodPut {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return "", err
		}
		req.Body = io.NopCloser(bytes.NewBuffer(body)) // Restore the request body

		curlCommand.WriteString(fmt.Sprintf("--data %q ", string(body)))
	}

	return curlCommand.String(), nil
}

// sendRequest takes an APIRequest struct as input and generates an API call
// using the configuration data inside. It returns a pointer to an http response
// struct and an error, if applicable.
func (c *Client) sendRequest(request *request) (*http.Response, error) {
	if c == nil {
		return nil, errors.New("invalid Keyfactor client, please check your configuration")
	}
	serverConfig := c.AuthClient.GetServerConfig()
	if serverConfig == nil {
		return nil, errors.New("invalid Keyfactor client, please check your configuration")
	}

	u, err := url.Parse(serverConfig.Host) // Parse raw Hostname into URL structure
	if err != nil {
		return nil, err
	}
	if u.Scheme != "https" {
		log.Printf("[WARN] Forcing https scheme on '%s'", u.Scheme)
		u.Scheme = "https"
	}
	apiPath := serverConfig.APIPath
	if apiPath == "" {
		apiPath = DefaultAPIPath
	}
	endpoint := fmt.Sprintf("%s/", strings.Trim(apiPath, "/")) + request.Endpoint
	log.Printf("[DEBUG] Endpoint: %s", endpoint)
	u.Path = path.Join(u.Path, endpoint) // Attach enroll endpoint
	log.Printf("[DEBUG] URL: %s", u.String())

	// Set request query
	if request.Query != nil {
		log.Printf("[DEBUG] Setting query parameters")
		queryString := u.Query()
		for _, query := range request.Query.Query {
			log.Printf("[TRACE] Setting query parameter %s=%s", query.Elem1, query.Elem2)
			queryString.Set(query.Elem1, query.Elem2)
		}
		log.Printf("[DEBUG] Encoding query string")
		u.RawQuery = queryString.Encode()
		log.Printf("[TRACE] Query string: %s", u.RawQuery)
		log.Printf("[DEBUG] Replacing '+' in query string")
		u.RawQuery = strings.ReplaceAll(u.RawQuery, "+", "%20")
		log.Printf("[DEBUG] Query string: %s", u.RawQuery)
	}

	keyfactorPath := u.String() // Convert absolute path to string

	log.Printf("[INFO] Preparing a %s request to path '%s'", request.Method, keyfactorPath)
	jsonByes, mErr := json.Marshal(request.Payload)
	if mErr != nil {
		return nil, mErr
	}
	log.Printf("[TRACE] Request body: %s", jsonByes)

	req, reqErr := http.NewRequest(request.Method, keyfactorPath, bytes.NewBuffer(jsonByes))
	if reqErr != nil {
		log.Printf("[ERROR] Error creating request: %s", reqErr)
		return nil, reqErr
	}

	log.Printf("[DEBUG] Setting request headers")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Set custom Keyfactor headers
	for _, headers := range request.Headers.Headers {
		//mask the Authorization header
		if strings.ToLower(headers.Elem1) == "authorization" {
			log.Printf("[TRACE] Setting header %s=********", headers.Elem1)
		} else {
			log.Printf("[TRACE] Setting header %s=%s", headers.Elem1, headers.Elem2)
		}
		req.Header.Set(headers.Elem1, headers.Elem2)
	}

	// Log the request
	logRequest(req)
	httpClient, cErr := c.getHttpClient()
	if cErr != nil {
		return nil, cErr
	}
	resp, respErr := httpClient.Do(req)

	// NOTE: this used to silently retry on "context deadline exceeded" (up to
	// MAX_CONTEXT_DEADLINE_RETRIES times) without ever surfacing that a retry
	// happened. That's unsafe for two reasons:
	//   1. Retrying a non-idempotent request (e.g. a POST enrollment) after a
	//      client-side timeout risks creating a second server-side resource
	//      if the original request actually succeeded after the client gave
	//      up on it -- exactly the scenario callers need to detect via the
	//      returned error, not have hidden from them by a "successful" retry.
	//   2. If every retry also failed, `resp` was never reassigned from its
	//      original nil value and there was no `return` for this case, so
	//      control fell through to `resp.StatusCode` below on a nil
	//      *http.Response, panicking the caller (e.g. crashing `terraform
	//      apply` outright).
	// Callers that need retry-with-backoff semantics around a timeout (and
	// that know their request is safe to repeat) should implement that at
	// their own call site, where they have the context to decide; this layer
	// now always returns the transport error untouched.
	switch {
	case respErr != nil:
		log.Printf("[ERROR] Error sending '%s' request to '%s': %s", request.Method, request.Endpoint, respErr)
		return nil, respErr
	case resp == nil:
		log.Printf(
			"[ERROR] No response from Keyfactor Command for '%s' request to '%s'",
			request.Method,
			request.Endpoint,
		)
		return nil, errors.New("no response from Keyfactor Command")
	}
	var stringMessage string
	log.Printf("[DEBUG] Response code from '%s' request to '%s': %d", request.Method, request.Endpoint, resp.StatusCode)
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
		log.Printf("[DEBUG] '%s' request to '%s' succeeded", request.Method, request.Endpoint)
		return resp, nil
	} else if resp.StatusCode == http.StatusNotFound {
		stringMessage = fmt.Sprintf(
			"Error %d - the requested resource was not found. Please check the request and try again.",
			resp.StatusCode,
		)
		log.Printf(
			"[ERROR] '%s' request to '%s' returned status '%d': %s", request.Method, keyfactorPath,
			resp.StatusCode,
			stringMessage,
		)

		var errorMessage map[string]interface{} // Decode JSON body to handle issue
		err = json.NewDecoder(resp.Body).Decode(&errorMessage)

		if err != nil {
			_, derr := httputil.DumpResponse(resp, true)
			if derr != nil {
				log.Printf(
					"[ERROR] Error dumping response body for '%s' request to '%s': %s",
					request.Method,
					request.Endpoint,
					derr,
				)
				return nil, derr
			}
			uerr := errors.New(
				fmt.Sprintf(
					"%d - Unknown error connecting to Keyfactor %s, please check your connection.",
					resp.StatusCode,
					endpoint,
				),
			)
			log.Printf("[ERROR] %s", uerr)
			return nil, uerr
		}

		stringMessage = fmt.Sprintf("%v", errorMessage["Message"])
		return nil, errors.New(stringMessage)
	} else if resp.StatusCode == http.StatusUnauthorized {
		dmp, derr := httputil.DumpResponse(resp, true)
		if derr != nil {
			return nil, derr
		}
		//convert response to string
		if dmp != nil {
			respStr := string(dmp)
			errMsg := fmt.Sprintf("http %d: %s", resp.StatusCode, respStr)
			return nil, errors.New(errMsg)
		}
		errMsg := fmt.Sprintf(
			"%d - %s",
			resp.StatusCode,
			"Unauthorized: Access is denied due to invalid credentials",
		)
		log.Printf("[ERROR] %s", errMsg)
		err = errors.New(errMsg)
		return nil, err
	} else {
		log.Printf(
			"[DEBUG] Attempting to decode error response body for '%s' request to '%s'",
			request.Method,
			request.Endpoint,
		)
		var errorMessage map[string]interface{} // Decode JSON body to handle issue
		err = json.NewDecoder(resp.Body).Decode(&errorMessage)

		if err != nil {
			_, derr := httputil.DumpResponse(resp, true)
			if derr != nil {
				log.Printf(
					"[ERROR] Error dumping response body for '%s' request to '%s': %s",
					request.Method,
					request.Endpoint,
					derr,
				)
				return nil, derr
			}
			uerr := errors.New(
				fmt.Sprintf(
					"%d - Unknown error connecting to Keyfactor %s, please check your connection.",
					resp.StatusCode,
					endpoint,
				),
			)
			log.Printf("[ERROR] %s", uerr)
			return nil, uerr
		}

		log.Printf(
			"[DEBUG] Error response body for '%s' request to '%s': %s",
			request.Method,
			request.Endpoint,
			errorMessage,
		)
		_, hasFailedOps := errorMessage["FailedOperations"]
		if hasFailedOps {
			var fOps []string
			log.Printf("[TRACE] Failed operations: %s", errorMessage["FailedOperations"])
			for _, v := range errorMessage["FailedOperations"].([]interface{}) {
				fOps = append(fOps, fmt.Sprintf("%s", v.(map[string]interface{})["Reason"]))
			}
			fOpsStr := strings.Join(fOps, ", ")
			stringMessage += fmt.Sprintf("%s. %s", errorMessage["Message"], fOpsStr)
			log.Printf("[TRACE] Failed ops string: %s", stringMessage)
		}
		_, hasMsg := errorMessage["Message"]
		if hasMsg {
			log.Printf("[TRACE] Error message: %s", errorMessage["Message"])
			stringMessage += fmt.Sprintf("%s", errorMessage["Message"])
		}
		log.Printf(
			"[ERROR] '%s' request to '%s' returned status '%d': %s",
			request.Method,
			keyfactorPath,
			resp.StatusCode,
			stringMessage,
		)

		return nil, errors.New(stringMessage)
	}
}

func getTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func buildQuery(params map[string]interface{}, filterName string) (apiQuery, error) {
	query := apiQuery{
		Query: []StringTuple{},
	}
	var qStr string
	for k, v := range params {
		switch v.(type) {
		case []string:
			var nestedQuery []string
			for _, val := range v.([]string) {
				if val != "" {
					nestedQuery = append(nestedQuery, fmt.Sprintf("%s -eq \"%s\"", k, val))
				}
			}
			if v != nil && len(v.([]string)) > 0 {
				if qStr != "" {
					qStr += fmt.Sprintf("AND (%s)", strings.Join(nestedQuery, " OR "))
				} else {
					qStr += fmt.Sprintf("(%s)", strings.Join(nestedQuery, " OR "))
				}
			}

		case string:
			if qStr != "" {
				qStr += fmt.Sprintf("AND %s -eq \"%s\"", k, v.(string))
			} else if v != "" {
				qStr += fmt.Sprintf("(%s -eq \"%s\")", k, v.(string))
			}
		case int:
			if qStr != "" {
				qStr += fmt.Sprintf("AND %s -eq %d", k, v.(int))
			} else {
				qStr += fmt.Sprintf("(%s -eq %d)", k, v.(int))
			}
		case bool:
			if qStr != "" {
				qStr += fmt.Sprintf("AND %s -eq %t", k, v.(bool))
			} else {
				qStr += fmt.Sprintf("(%s -eq %t)", k, v.(bool))
			}
		}

	}
	query.Query = append(query.Query, StringTuple{filterName, qStr})
	return query, nil
}
