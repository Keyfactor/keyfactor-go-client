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
	"strings"
	"time"
)

var (
	EnvCommandHostname = "KEYFACTOR_HOSTNAME"
	EnvCommandUsername = "KEYFACTOR_USERNAME"
	EnvCommandPassword = "KEYFACTOR_PASSWORD"
	EnvCommandDomain   = "KEYFACTOR_DOMAIN"
)

type Client struct {
	hostname        string
	httpClient      *http.Client
	basicAuthString string
}

// AuthConfig is a struct holding all necessary client configuration data
// for communicating with the Keyfactor API. This includes the hostname,
// username, password, and domain.
type AuthConfig struct {
	Hostname string
	Username string
	Password string
	Domain   string
}

// NewKeyfactorClient creates a new Keyfactor client instance. A configured Client is returned with methods used to
// interact with Keyfactor.
func NewKeyfactorClient(auth *AuthConfig) (*Client, error) {
	c, err := loginToKeyfactor(auth)
	if err != nil {
		return nil, err
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

	c := &Client{
		hostname:        auth.Hostname,
		httpClient:      &http.Client{Timeout: 10 * time.Second},
		basicAuthString: buildBasicAuthString(auth),
	}

	_, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] Successfully logged into Keyfactor at host %s", c.hostname)

	return c, nil
}

// sendRequest takes an APIRequest struct as input and generates an API call
// using the configuration data inside. It returns a pointer to an http response
// struct and an error, if applicable.
func (c *Client) sendRequest(request *request) (*http.Response, error) {
	if c == nil {
		return nil, errors.New("invalid Keyfactor client, please check your configuration")
	}
	u, err := url.Parse(c.hostname) // Parse raw hostname into URL structure
	if err != nil {
		return nil, err
	}
	if u.Scheme != "https" {
		u.Scheme = "https"
	}
	endpoint := "KeyfactorAPI/" + request.Endpoint
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

	log.Printf("[INFO] Preparing a %s request to path '%s'", request.Method, keyfactorPath)
	jsonByes, mErr := json.Marshal(request.Payload)
	if mErr != nil {
		return nil, mErr
	}
	//log.Printf("[TRACE] Request body: %s", jsonByes)

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
	if respErr != nil {
		return nil, respErr
	}
	var stringMessage string
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
		log.Printf("[DEBUG] %s succeeded with response code %d", request.Method, resp.StatusCode)
		return resp, nil
	} else if resp.StatusCode == http.StatusNotFound {
		stringMessage = fmt.Sprintf("Error %d - the requested resource was not found. Please check the request and try again.", resp.StatusCode)
		log.Printf("[ERROR] Call to %s returned status %d. %s", keyfactorPath, resp.StatusCode, stringMessage)
		return nil, errors.New(stringMessage)
	} else if resp.StatusCode == http.StatusUnauthorized {
		_, derr := httputil.DumpResponse(resp, true)
		if derr != nil {
			return nil, derr
		}
		err = errors.New("401 - Unauthorized: Access is denied due to invalid credentials")
		return nil, err
	} else {
		var errorMessage map[string]interface{} // Decode JSON body to handle issue
		err = json.NewDecoder(resp.Body).Decode(&errorMessage)

		if err != nil {
			_, derr := httputil.DumpResponse(resp, true)
			if derr != nil {
				return nil, derr
			}
			uerr := errors.New(fmt.Sprintf("%d - Unknown error connecting to Keyfactor %s, please check your connection.", resp.StatusCode, endpoint))
			return nil, uerr
		}

		log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
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
	//log.Println("[TRACE] Building Authorization field")
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
