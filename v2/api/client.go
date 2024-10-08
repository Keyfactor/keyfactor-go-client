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
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var (
	EnvCommandHostname     = "KEYFACTOR_HOSTNAME"
	EnvCommandUsername     = "KEYFACTOR_USERNAME"
	EnvCommandPassword     = "KEYFACTOR_PASSWORD"
	EnvCommandDomain       = "KEYFACTOR_DOMAIN"
	EnvCommandAPI          = "KEYFACTOR_API_PATH"
	EnvCommandTimeout      = "KEYFACTOR_TIMEOUT"
	DefaultAPIPath         = "KeyfactorAPI"
	EnvCommandClientId     = "KEYFACTOR_CLIENT_ID"
	EnvCommandClientSecret = "KEYFACTOR_CLIENT_SECRET"
)

type Client struct {
	Hostname        string
	httpClient      *http.Client
	basicAuthString string
	ApiPath         string
}

// AuthConfig is a struct holding all necessary client configuration data
// for communicating with the Keyfactor API. This includes the Hostname,
// username, password, and domain.
type AuthConfig struct {
	Hostname string
	Username string
	Password string
	Domain   string
	APIPath  string
	Timeout  int
}

// NewKeyfactorClient creates a new Keyfactor client instance. A configured Client is returned with methods used to
// interact with Keyfactor.
func NewKeyfactorClient(auth *AuthConfig) (*Client, error) {
	// set log to stdout
	log.SetOutput(os.Stdout)
	////log.printf("[DEBUG] Logging into Keyfactor at host %s", auth.Hostname)
	c, err := loginToKeyfactor(auth)
	if err != nil {
		return nil, err
	}

	if auth.Timeout > 0 {
		c.httpClient = &http.Client{Timeout: time.Duration(auth.Timeout) * time.Second}
	} else {
		c.httpClient = &http.Client{Timeout: MAX_WAIT_SECONDS * time.Second}
	}

	return c, nil
}

// loginToKeyfactor is a helper function that creates a new Keyfactor client instance. A configured Client is returned
// with methods used to interact with Keyfactor.
func loginToKeyfactor(auth *AuthConfig) (*Client, error) {
	if auth.Hostname == "" {
		envHostname := os.Getenv(EnvCommandHostname)
		if envHostname != "" {
			auth.Hostname = envHostname
		} else {
			return nil, fmt.Errorf("%s is required", EnvCommandHostname)
		}
	}
	if auth.APIPath == "" {
		envAPIPath := os.Getenv(EnvCommandAPI)
		if envAPIPath != "" {
			auth.APIPath = envAPIPath
		} else {
			auth.APIPath = DefaultAPIPath
		}
	}
	if auth.Username == "" {
		envUsername := os.Getenv(EnvCommandUsername)
		if envUsername != "" {
			envDomain := os.Getenv(EnvCommandDomain)
			if envDomain != "" && auth.Domain == "" && !strings.Contains(envUsername, envDomain) {
				auth.Domain = envDomain
			}
			if auth.Domain != "" && !strings.Contains(envUsername, envDomain) {
				auth.Username = auth.Domain + "\\" + envUsername
			} else {
				auth.Username = envUsername
			}
		} else {
			return nil, fmt.Errorf("%s is required", EnvCommandUsername)
		}
	}
	if auth.Password == "" {
		envPassword := os.Getenv(EnvCommandPassword)
		if envPassword != "" {
			auth.Password = envPassword
		} else {
			return nil, fmt.Errorf("%s is required", EnvCommandPassword)
		}
	}

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: "Status/Endpoints",
		Headers:  headers,
	}

	timeoutStr := os.Getenv(EnvCommandTimeout)
	timeout := MAX_WAIT_SECONDS
	if timeoutStr != "" {
		//convert to int and check if greater than 0
		timeoutInt, err := strconv.Atoi(timeoutStr)
		if err == nil && timeoutInt > 0 {
			timeout = timeoutInt
		}
	} else if auth.Timeout > 0 {
		timeout = auth.Timeout
	}

	c := &Client{
		Hostname:        auth.Hostname,
		httpClient:      &http.Client{Timeout: time.Duration(timeout) * time.Second},
		basicAuthString: buildBasicAuthString(auth),
		ApiPath:         auth.APIPath,
	}

	_, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	////log.printf("[INFO] Successfully logged into Keyfactor at host %s", c.Hostname)

	return c, nil
}

// sendRequest takes an APIRequest struct as input and generates an API call
// using the configuration data inside. It returns a pointer to an http response
// struct and an error, if applicable.
func (c *Client) sendRequest(request *request) (*http.Response, error) {
	if c == nil {
		return nil, errors.New("invalid Keyfactor client, please check your configuration")
	}
	u, err := url.Parse(c.Hostname) // Parse raw Hostname into URL structure
	if err != nil {
		return nil, err
	}
	if u.Scheme != "https" {
		u.Scheme = "https"
	}
	endpoint := fmt.Sprintf("%s/", strings.Trim(c.ApiPath, "/")) + request.Endpoint
	u.Path = path.Join(u.Path, endpoint) // Attach enroll endpoint

	// Set request query
	if request.Query != nil {
		queryString := u.Query()
		for _, query := range request.Query.Query {
			queryString.Set(query.Elem1, query.Elem2)
		}
		u.RawQuery = queryString.Encode()
		u.RawQuery = strings.ReplaceAll(u.RawQuery, "+", "%20")
	}

	keyfactorPath := u.String() // Convert absolute path to string

	////log.printf("[INFO] Preparing a %s request to path '%s'", request.Method, keyfactorPath)
	jsonByes, mErr := json.Marshal(request.Payload)
	if mErr != nil {
		return nil, mErr
	}
	////log.printf("[TRACE] Request body: %s", jsonByes)

	req, reqErr := http.NewRequest(request.Method, keyfactorPath, bytes.NewBuffer(jsonByes))
	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.basicAuthString)

	// Set custom Keyfactor headers
	for _, headers := range request.Headers.Headers {
		req.Header.Set(headers.Elem1, headers.Elem2)
	}

	resp, respErr := c.httpClient.Do(req)

	// check if context deadline exceeded
	switch {
	case respErr != nil && (strings.Contains(respErr.Error(), "context deadline exceeded")):
		sleepDuration := time.Duration(1) * time.Second
		for i := 0; i < MAX_CONTEXT_DEADLINE_RETRIES; i++ {
			// sleep for exponential backoff
			if i > 0 {
				sleepDuration *= 2
				if sleepDuration > time.Duration(MAX_WAIT_SECONDS)*time.Second {
					sleepDuration = time.Duration(MAX_WAIT_SECONDS) * time.Second
				}
				////log.printf(
				//	"[DEBUG] %s request to %s failed with error %s, retrying in %s seconds...",
				//	request.Method,
				//	keyfactorPath,
				//	respErr.Error(),
				//	sleepDuration,
				//)
				time.Sleep(sleepDuration)
			}

			////log.printf(
			//	"[DEBUG] %s request to %s failed with error %s, retrying...",
			//	request.Method,
			//	keyfactorPath,
			//	respErr.Error(),
			//)
			req, reqErr = http.NewRequest(request.Method, keyfactorPath, bytes.NewBuffer(jsonByes))
			if reqErr != nil {
				return nil, reqErr
			}
			resp2, respErr2 := c.httpClient.Do(req)
			if respErr2 == nil {
				resp = resp2
				break
			}
		}
	case respErr != nil:
		return nil, respErr
	case resp == nil:
		return nil, errors.New("no response from Keyfactor Command")
	}
	var stringMessage string
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
		////log.printf("[DEBUG] %s succeeded with response code %d", request.Method, resp.StatusCode)
		return resp, nil
	} else if resp.StatusCode == http.StatusNotFound {
		stringMessage = fmt.Sprintf(
			"Error %d - the requested resource was not found. Please check the request and try again.",
			resp.StatusCode,
		)
		////log.printf("[ERROR] Call to %s returned status %d. %s", keyfactorPath, resp.StatusCode, stringMessage)
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
			"http %d: %s",
			resp.StatusCode,
			"Unauthorized: Access is denied due to invalid credentials",
		)
		err = errors.New(errMsg)
		return nil, err
	} else {
		var errorMessage map[string]interface{} // Decode JSON body to handle issue
		err = json.NewDecoder(resp.Body).Decode(&errorMessage)

		if err != nil {
			_, derr := httputil.DumpResponse(resp, true)
			if derr != nil {
				return nil, derr
			}
			uerr := errors.New(
				fmt.Sprintf(
					"%d - Unknown error connecting to Keyfactor %s, please check your connection.",
					resp.StatusCode,
					endpoint,
				),
			)
			return nil, uerr
		}

		////log.printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
		_, hasFailedOps := errorMessage["FailedOperations"]
		if hasFailedOps {
			var fOps []string
			for _, v := range errorMessage["FailedOperations"].([]interface{}) {
				fOps = append(fOps, fmt.Sprintf("%s", v.(map[string]interface{})["Reason"]))
			}
			fOpsStr := strings.Join(fOps, ", ")
			stringMessage += fmt.Sprintf("%s. %s", errorMessage["Message"], fOpsStr)
		}
		_, hasMsg := errorMessage["Message"]
		if hasMsg {
			stringMessage += fmt.Sprintf("%s", errorMessage["Message"])
		}
		return nil, errors.New(stringMessage)
	}
}

// buildBasicAuthString constructs a basic authorization string necessary for basic authorization to Keyfactor. It
// returns a base-64 encoded auth string including the 'Basic ' prefix.
func buildBasicAuthString(auth *AuthConfig) string {
	var authString string
	////log.println("[TRACE] Building Authorization field")
	if auth.Domain != "" && !strings.Contains(auth.Username, auth.Domain) {
		authString = strings.Join([]string{auth.Domain, "\\", auth.Username, ":", auth.Password}, "")
	} else {
		authString = strings.Join([]string{auth.Username, ":", auth.Password}, "")
	}
	authBytes := []byte(authString)
	authDigest := base64.StdEncoding.EncodeToString(authBytes)
	authField := strings.Join([]string{"Basic ", authDigest}, "")
	return authField
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
