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
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/Keyfactor/keyfactor-go-client-sdk/api/keyfactor"
)

// GetStoreContainers returns a list of store containers
func (c *Client) GetStoreContainers() (*[]CertStoreContainer, error) {
	log.Println("[INFO] Listing certificate store containers.")

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration(make(map[string]string))
	apiClient := keyfactor.NewAPIClient(configuration)

	resp, _, err := apiClient.CertificateStoreContainerApi.CertificateStoreContainerGetAllCertificateStoreContainers(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp []CertStoreContainer
	for i := range resp {
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
		query.Query = append(
			query.Query, StringTuple{
				"pq.queryString", fmt.Sprintf(`Name -eq "%s"`, id),
			},
		)
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
