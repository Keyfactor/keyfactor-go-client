package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	EnvCommandHostname = "KEYFACTOR_HOSTNAME"
	EnvCommandUsername = "KEYFACTOR_USERNAME"
	EnvCommandPassword = "KEYFACTOR_PASSWORD"
	EnvCommandDomain   = "KEYFACTOR_DOMAIN"
	EnvCommandAPI      = "KEYFACTOR_API_PATH"
	EnvCommandTimeout  = "KEYFACTOR_TIMEOUT"
	DefaultAPIPath     = "KeyfactorAPI"
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

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

// NewKeyfactorClient creates a new Keyfactor client instance. A configured Client is returned with methods used to
// interact with Keyfactor.
func NewKeyfactorClient(auth *AuthConfig) (*Client, error) {
	initLogger()
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
	initLogger()
	log.Debug().Msg("Entering loginToKeyfactor")
	if auth.Hostname == "" {
		envHostname := os.Getenv(EnvCommandHostname)
		if envHostname != "" {
			auth.Hostname = envHostname
		} else {
			return nil, fmt.Errorf("%s is required", EnvCommandHostname)
		}
	}
	log.Debug().Str("hostname", auth.Hostname).Send()
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
			log.Error().Msgf("%s is required", EnvCommandUsername)
			return nil, fmt.Errorf("%s is required", EnvCommandUsername)
		}
	}
	log.Debug().Str("username", auth.Username).Send()
	if auth.Password == "" {
		envPassword := os.Getenv(EnvCommandPassword)
		if envPassword != "" {
			auth.Password = envPassword
		} else {
			return nil, fmt.Errorf("%s is required", EnvCommandPassword)
		}
	}
	if auth.Password != "" {
		log.Debug().Str("password", "********").Msg("Password is set")
	}
	log.Trace().Str("password", auth.Password).Send()

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}
	log.Trace().Interface("headers", headers).Msg("Request headers for Keyfactor API call")

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
	log.Debug().
		Int("timeout", timeout).
		Msg("Timeout for Keyfactor API call in seconds")

	log.Debug().
		Str("hostname", auth.Hostname).
		Int("timeout", timeout).
		Str("apiPath", auth.APIPath).
		Msg("Creating Keyfactor Command client")
	c := &Client{
		Hostname:        auth.Hostname,
		httpClient:      &http.Client{Timeout: time.Duration(timeout) * time.Second},
		basicAuthString: buildBasicAuthString(auth),
		ApiPath:         auth.APIPath,
	}

	log.Debug().Msg("Calling sendRequest...")
	_, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		log.Error().Err(err).Msg("Keyfactor Command returned an error.")
		return nil, err
	}

	log.Info().Msg("Successfully connected to Keyfactor Command")
	log.Debug().Msg("Exiting loginToKeyfactor")
	return c, nil
}

// sendRequest takes an APIRequest struct as input and generates an API call
// using the configuration data inside. It returns a pointer to an http response
// struct and an error, if applicable.
func (c *Client) sendRequest(request *request) (*http.Response, error) {
	initLogger()
	if c == nil {
		return nil, errors.New("invalid Keyfactor client, please check your configuration")
	}
	u, err := url.Parse(c.Hostname) // Parse raw Hostname into URL structure
	if err != nil {
		log.Error().Err(err).Msg("Error parsing Keyfactor hostname")
		return nil, err

	}
	if u.Scheme != "https" {
		log.Warn().Str("scheme", u.Scheme).Msg("Keyfactor hostname is not using HTTPS, this is not allowed.")
		u.Scheme = "https"
	}
	endpoint := fmt.Sprintf("%s/", strings.Trim(c.ApiPath, "/")) + request.Endpoint
	log.Debug().Str("endpoint", endpoint).Msg("Endpoint for Keyfactor API call")
	u.Path = path.Join(u.Path, endpoint) // Attach enroll endpoint
	log.Debug().Str("path", u.Path).Msg("Path for Keyfactor API call")

	// Set request query
	if request.Query != nil {
		log.Debug().Msg("Query passed to Keyfactor API call")
		log.Trace().Interface("query", request.Query).Msg("Query for Keyfactor API call")
		queryString := u.Query()
		log.Debug().Interface("query", queryString).Msg("Query string for Keyfactor API call")
		for _, query := range request.Query.Query {
			queryString.Set(query.Elem1, query.Elem2)
		}
		log.Debug().Interface("query", queryString).Msg("Encoding query string")
		u.RawQuery = queryString.Encode()
		log.Trace().Str("rawQuery", u.RawQuery).Msg("Raw query string for Keyfactor API call")
		u.RawQuery = strings.ReplaceAll(u.RawQuery, "+", "%20")
		log.Trace().Str("rawQuery", u.RawQuery).Msg("Raw query string w/ replacements for Keyfactor API call")
	}

	keyfactorPath := u.String() // Convert absolute path to string

	//log.Printf("[INFO] Preparing a %s request to path '%s'", request.Method, keyfactorPath)
	log.Info().
		Str("method", request.Method).
		Str("hostname", c.Hostname).
		Str("endpoint", request.Endpoint).
		Str("apiPath", c.ApiPath).
		Str("path", keyfactorPath).
		Str("query", u.RawQuery).
		Msg("Preparing JSON payload for Keyfactor API call")
	log.Trace().Interface("payload", request.Payload).Msg("Payload for Keyfactor API call")
	jsonByes, mErr := json.Marshal(request.Payload)
	if mErr != nil {
		log.Error().Err(mErr).Msg("Error marshalling JSON payload for Keyfactor API call")
		return nil, mErr
	}
	log.Trace().Bytes("json", jsonByes).Msg("JSON payload for Keyfactor API call")

	log.Debug().Msg("Creating HTTP request for Keyfactor API call")
	req, reqErr := http.NewRequest(request.Method, keyfactorPath, bytes.NewBuffer(jsonByes))
	if reqErr != nil {
		log.Error().Err(reqErr).Msg("Error creating HTTP request for Keyfactor API call")
		return nil, reqErr
	}

	log.Debug().Msg("Setting request headers for Keyfactor API call")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.basicAuthString)

	// Set custom Keyfactor headers
	for _, headers := range request.Headers.Headers {
		req.Header.Set(headers.Elem1, headers.Elem2)
	}
	log.Trace().Interface("headers", req.Header).Msg("Request headers for Keyfactor API call")

	log.Info().Str("method", request.Method).Str("path", keyfactorPath).Msg("Sending HTTP request to Keyfactor API")
	resp, respErr := c.httpClient.Do(req)

	// check if context deadline exceeded
	switch {
	case respErr != nil && (strings.Contains(respErr.Error(), "context deadline exceeded")):
		sleepDuration := time.Duration(1) * time.Second
		log.Warn().
			Err(respErr).
			Str("sleep_duration", sleepDuration.String()).
			Msg("Context deadline exceeded, retrying...")
		for i := 0; i < MAX_CONTEXT_DEADLINE_RETRIES; i++ {
			// sleep for exponential backoff

			if i > 0 {
				sleepDuration *= 2
				if sleepDuration > time.Duration(MAX_WAIT_SECONDS)*time.Second {
					sleepDuration = time.Duration(MAX_WAIT_SECONDS) * time.Second
				}
				//log.Printf("[DEBUG] %s request to %s failed with error %s, retrying in %s seconds...", request.Method, keyfactorPath, respErr.Error(), sleepDuration)
				log.Debug().
					Str("sleep_duration", sleepDuration.String()).Int("retry_count", i).
					Msg("Sleeping for exponential backoff")
				time.Sleep(sleepDuration)
			}

			//log.Printf("[DEBUG] %s request to %s failed with error %s, retrying...", request.Method, keyfactorPath, respErr.Error())
			log.Trace().Msg("Rebuilding HTTP request for Keyfactor API call")
			req, reqErr = http.NewRequest(request.Method, keyfactorPath, bytes.NewBuffer(jsonByes))
			if reqErr != nil {
				log.Error().Err(reqErr).Msg("Error creating HTTP request for Keyfactor API call")
				return nil, reqErr
			}
			log.Debug().
				Str("method", request.Method).
				Str("path", keyfactorPath).
				Str("query", u.RawQuery).
				Msg("Retrying HTTP request to Keyfactor API")
			resp2, respErr2 := c.httpClient.Do(req)
			if respErr2 == nil {
				log.Debug().Msg("Request succeeded after retry")
				resp = resp2
				break
			}
			log.Error().Err(respErr2).Msg("Error sending HTTP request to Keyfactor API")
		}
	case respErr != nil:
		log.Error().Err(respErr).Msg("Error sending HTTP request to Keyfactor API")
		return nil, respErr
	case resp == nil:
		log.Error().Msg("No response from Keyfactor Command")
		return nil, errors.New("no response from Keyfactor Command")
	}
	var stringMessage string
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
		//log.Printf("[DEBUG] %s succeeded with response code %d", request.Method, resp.StatusCode)
		log.Debug().Int("statusCode", resp.StatusCode).Msg("Request succeeded")
		return resp, nil
	} else if resp.StatusCode == http.StatusNotFound {
		stringMessage = fmt.Sprintf("http%d: the requested resource was not found, please check the request and try again.", resp.StatusCode)
		log.Error().Msg(stringMessage)
		//log.Printf("[ERROR] Call to %s returned status %d. %s", keyfactorPath, resp.StatusCode, stringMessage)
		return nil, errors.New(stringMessage)
	} else if resp.StatusCode == http.StatusUnauthorized {
		dmp, derr := httputil.DumpResponse(resp, true)
		//convert dmp from bytes to string
		dmpStr := string(dmp)
		log.Error().Err(derr).Str("response", dmpStr).Msg("Unauthorized request to Keyfactor Command")
		if derr != nil {
			log.Error().Err(derr).Msg("Error dumping response from Keyfactor Command")
			return nil, derr
		}
		err = errors.New(fmt.Sprintf("http%d - Unauthorized: Access is denied due to invalid credentials", resp.StatusCode))
		return nil, err
	} else {
		log.Error().Msg("Request failed")
		var errorMessage map[string]interface{} // Decode JSON body to handle issue
		log.Debug().Msg("Decoding JSON response from Keyfactor Command")
		err = json.NewDecoder(resp.Body).Decode(&errorMessage)

		if err != nil {
			log.Error().Err(err).Msg("Error decoding JSON response from Keyfactor Command")
			_, derr := httputil.DumpResponse(resp, true)
			if derr != nil {
				log.Error().Err(derr).Msg("Error dumping response from Keyfactor Command")
				return nil, derr
			}
			uerr := errors.New(fmt.Sprintf("%d - Unknown error connecting to Keyfactor %s, please check your connection.", resp.StatusCode, endpoint))
			log.Error().
				Err(uerr).
				Str("method", req.Method).
				Str("path", keyfactorPath).
				Msg("Unknown error connecting to Keyfactor")
			return nil, uerr
		}

		//log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
		log.Debug().Int("statusCode", resp.StatusCode).Interface("errorMessage", errorMessage).Msg("Request failed")
		_, hasFailedOps := errorMessage["FailedOperations"]
		if hasFailedOps {
			log.Debug().Msg("Failed operations found in response")
			var fOps []string
			for _, v := range errorMessage["FailedOperations"].([]interface{}) {
				log.Trace().Interface("failedOperation", v).Msg("reason")
				fOps = append(fOps, fmt.Sprintf("%s", v.(map[string]interface{})["Reason"]))
			}
			fOpsStr := strings.Join(fOps, ", ")
			stringMessage += fmt.Sprintf("%s. %s", errorMessage["Message"], fOpsStr)
			log.Trace().Str("failedOperations", stringMessage).Msg("Failed operations")
		}
		_, hasMsg := errorMessage["Message"]
		if hasMsg {
			stringMessage += fmt.Sprintf("%s", errorMessage["Message"])
		}
		log.Error().Msg(stringMessage)
		return nil, errors.New(stringMessage)
	}
}

// buildBasicAuthString constructs a basic authorization string necessary for basic authorization to Keyfactor. It
// returns a base-64 encoded auth string including the 'Basic ' prefix.
func buildBasicAuthString(auth *AuthConfig) string {
	initLogger()
	log.Debug().Msg("Entering buildBasicAuthString")
	var authString string
	if auth.Domain != "" && !strings.Contains(auth.Username, auth.Domain) {
		authString = strings.Join([]string{auth.Domain, "\\", auth.Username, ":", auth.Password}, "")
	} else {
		authString = strings.Join([]string{auth.Username, ":", auth.Password}, "")
	}
	authBytes := []byte(authString)
	authDigest := base64.StdEncoding.EncodeToString(authBytes)
	authField := strings.Join([]string{"Basic ", authDigest}, "")
	//log.Trace().Str("authField", authField).Msg("Basic auth string") //INSECURE: authField contains sensitive information
	if authField == "" || authField == "Basic " {
		log.Error().Msg("Basic auth string is empty")
	}
	log.Debug().Msg("Exiting buildBasicAuthString")
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
