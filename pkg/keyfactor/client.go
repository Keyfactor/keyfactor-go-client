package keyfactor

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type APIClient struct {
	Hostname string
	Username string
	Password string
	Domain   string
}

type StringTuple struct {
	Elem1 string
	Elem2 string
}

type APIHeaders struct {
	Headers []StringTuple
}

type APIRequest struct {
	KFClient *APIClient
	Method   string
	Endpoint string
	Headers  *APIHeaders
	Payload  interface{}
}

func SendRequest(request *APIRequest) (*http.Response, error) {
	kfApiClientData := request.KFClient // Extract Keyfactor client data

	u, err := url.Parse(kfApiClientData.Hostname) // Parse raw hostname into URL structure
	if err != nil {
		return nil, err
	}
	if u.Scheme != "https" {
		u.Scheme = "https"
	}
	u.Path = path.Join(u.Path, request.Endpoint) // Attach enroll endpoint
	keyfactorPath := u.String()

	log.Printf("[INFO] Preparing a %s request to path '%s'", request.Method, keyfactorPath)
	jsonByes, err := json.Marshal(request.Payload)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Request body: %s", jsonByes)

	req, err := http.NewRequest(request.Method, keyfactorPath, bytes.NewBuffer(jsonByes))
	if err != nil {
		return nil, err
	}

	authField := buildAuthField(request.KFClient)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", authField)

	// Set custom Keyfactor headers
	for _, headers := range request.Headers.Headers {
		req.Header.Set(headers.Elem1, headers.Elem2)
	}

	log.Println("[TRACE] Executing request with client.Do")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func buildAuthField(creds *APIClient) string {
	var authString string
	log.Println("[TRACE] Building Authorization field")
	if creds.Domain == "" {
		authString = strings.Join([]string{creds.Username, ":", creds.Password}, "")
	} else {
		authString = strings.Join([]string{creds.Domain, "\\", creds.Username, ":", creds.Password}, "")
	}
	authBytes := []byte(authString)
	authDigest := base64.StdEncoding.EncodeToString(authBytes)
	authField := strings.Join([]string{"Basic ", authDigest}, "")
	return authField
}

func getTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}
