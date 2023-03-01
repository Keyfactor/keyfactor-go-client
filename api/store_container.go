package api

import (
	"context"
	"encoding/json"
	"fmt"
	keyfactor_command_client_api "github.com/spbsoluble/keyfactor-go-client-sdk"
	"log"
	"strconv"
)

// GetStoreContainers returns a list of store containers
func (c *Client) GetStoreContainers() (*[]CertStoreContainer, error) {
	log.Println("[INFO] Listing certificate store containers.")

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, _, err := apiClient.CertificateStoreContainerApi.CertificateStoreContainerGetAllCertificateStoreContainers(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp []CertStoreContainer
	for i, _ := range resp {
		var newCont CertStoreContainer
		mapResp, _ := resp[i].ToMap()
		jsonData, _ := json.Marshal(mapResp)
		json.Unmarshal(jsonData, &newCont)
		newResp = append(newResp, newCont)
	}

	return &newResp, nil
}

// GetStoreContainer takes an ID and returns a single store container
// TODO?
func (c *Client) GetStoreContainer(id interface{}) (*CertStoreContainer, error) {
	log.Printf("[INFO] Fetching certificate store containers %s.\n", id)
	var endpoint string
	var query apiQuery
	var jsonResp interface{}

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}
	var idInt int
	var cErr error
	switch id.(type) {
	case string:
		idInt, cErr = strconv.Atoi(id.(string))
	case int:
		idInt = id.(int)
	}

	if cErr == nil {
		// Endpoint returns a single store container
		endpoint = fmt.Sprintf("CertificateStoreContainers/%d", idInt)
		jsonResp = &CertStoreContainer{}
	} else {
		// Endpoint returns a list of store containers
		endpoint = "CertificateStoreContainers"
		query = apiQuery{
			Query: []StringTuple{},
		}
		query.Query = append(query.Query, StringTuple{
			"pq.queryString", fmt.Sprintf(`Name -eq "%s"`, id),
		})
		jsonResp = &[]CertStoreContainer{}
	}
	var keyfactorAPIStruct *request
	if query.Query != nil {
		keyfactorAPIStruct = &request{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  headers,
			Query:    &query,
		}
	} else {
		keyfactorAPIStruct = &request{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  headers,
		}
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	switch jsonResp.(type) {
	case *CertStoreContainer:
		return jsonResp.(*CertStoreContainer), nil
	case *[]CertStoreContainer:
		if len(*jsonResp.(*[]CertStoreContainer)) > 0 {
			return &(*jsonResp.(*[]CertStoreContainer))[0], nil
		}
		return nil, fmt.Errorf("no cert store container found with name %s", id)
	}
	return nil, fmt.Errorf("invalid API response from Keyfactor while getting cert store container %s", id)
}
