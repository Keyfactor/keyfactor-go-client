package api

import (
	"encoding/json"
	"fmt"
	"log"
)

// GetStoreContainers returns a list of store containers
func (c *Client) GetStoreContainers() (*[]CertStoreContainer, error) {
	log.Println("[INFO] Listing certificate store containers.")

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: "CertificateStoreContainers",
		Headers:  headers,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &[]CertStoreContainer{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// GetStoreContainer takes an ID and returns a single store container
func (c *Client) GetStoreContainer(id interface{}) (*CertStoreContainer, error) {
	log.Printf("[INFO] Fetching certificat store containers %s.\n", id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	var endpoint string
	var query apiQuery
	switch id.(type) {
	case int:
		endpoint = fmt.Sprintf("CertificateStoreContainers/%s", id)
	case string:
		endpoint = "CertificateStoreContainers"
		query = apiQuery{
			Query: []StringTuple{},
		}
		query.Query = append(query.Query, StringTuple{
			"pq.queryString", fmt.Sprintf(`name -eq "%s"`, id),
		})
	}
	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  headers,
		Query:    &query,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &CertStoreContainer{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}
