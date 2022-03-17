package keyfactor

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// GetTemplate takes arguments for a template ID used to facilitate the retrieval
// of certificate template context. The primary query required to get certificate context is the template ID. A pointer
//to a GetTemplateResponse structure is returned, containing the template context.
func (c *Client) GetTemplate(Id int) (*GetTemplateResponse, error) {
	if Id == 0 {
		return nil, errors.New("template id required to get template")
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "KeyfactorAPI/Templates/" + fmt.Sprintf("%d", Id) // Append ID to complete endpoint

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  headers,
		Query:    nil,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		log.Printf("[DEBUG] GET succeeded with response code %d", resp.StatusCode)

		jsonResp := &GetTemplateResponse{}
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		if err != nil {
			return nil, err
		}
		return jsonResp, err
	}

	var errorMessage interface{} // Decode JSON body to handle issue
	err = json.NewDecoder(resp.Body).Decode(&errorMessage)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
	stringMessage := fmt.Sprintf("%v", errorMessage)

	return nil, errors.New(stringMessage)
}

// UpdateTemplate takes arguments for a template ID used to facilitate the retrieval
// of certificate template context. The primary query required to get certificate context is the template ID. A pointer
//to a GetTemplateResponse structure is returned, containing the template context.
func (c *Client) UpdateTemplate(Id int) (*GetTemplateResponse, error) {
	if Id == 0 {
		return nil, errors.New("template id required to get template")
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "KeyfactorAPI/Templates/" + fmt.Sprintf("%d", Id) // Append ID to complete endpoint

	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: endpoint,
		Headers:  headers,
		Query:    nil,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		log.Printf("[DEBUG] %s succeeded with response code %d", keyfactorAPIStruct.Method, resp.StatusCode)

		jsonResp := &GetTemplateResponse{}
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		if err != nil {
			return nil, err
		}
		return jsonResp, err
	}

	var errorMessage interface{} // Decode JSON body to handle issue
	err = json.NewDecoder(resp.Body).Decode(&errorMessage)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Request failed with code %d and message %v", resp.StatusCode, errorMessage)
	stringMessage := fmt.Sprintf("%v", errorMessage)

	return nil, errors.New(stringMessage)
}
