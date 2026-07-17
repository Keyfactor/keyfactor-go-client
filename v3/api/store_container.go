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
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// getStoreContainersMaxPages is the maximum number of pages GetStoreContainers
// will fetch before aborting with an error. It is an unexported package-level
// var (not a const) so that tests can lower it without iterating thousands of
// times.
var getStoreContainersMaxPages = 10000

// GetStoreContainers returns a list of store containers, paginating
// automatically so that instances with more than the server's default page
// size (100) return all containers.
func (c *Client) GetStoreContainers() (*[]CertStoreContainer, error) {
	log.Println("[INFO] Listing certificate store containers.")

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	const pageSize = 100
	var all []CertStoreContainer
	var page int
	for page = 1; page <= getStoreContainersMaxPages; page++ {
		keyfactorAPIStruct := &request{
			Method:   "GET",
			Endpoint: "CertificateStoreContainers",
			Headers:  headers,
			Query: &apiQuery{
				Query: []StringTuple{
					{"PageReturned", strconv.Itoa(page)},
					{"ReturnLimit", strconv.Itoa(pageSize)},
				},
			},
			Payload: nil,
		}

		resp, err := c.sendRequest(keyfactorAPIStruct)
		if err != nil {
			log.Printf("[ERROR] GetStoreContainers: request for page %d failed: %s", page, err)
			return nil, err
		}

		var pageResults []CertStoreContainer
		decodeErr := json.NewDecoder(resp.Body).Decode(&pageResults)
		resp.Body.Close()
		if decodeErr != nil {
			log.Printf("[ERROR] GetStoreContainers: failed to decode page %d: %s", page, decodeErr)
			return nil, decodeErr
		}

		all = append(all, pageResults...)
		if len(pageResults) == 0 || len(pageResults) < pageSize {
			break
		}
	}

	if page > getStoreContainersMaxPages {
		return nil, fmt.Errorf("GetStoreContainers: exceeded max pages (%d); server may be ignoring pagination", getStoreContainersMaxPages)
	}

	log.Printf("[INFO] Listed %d certificate store containers across %d page(s).", len(all), page)
	return &all, nil
}

// GetStoreContainer takes an ID and returns a single store container
func (c *Client) GetStoreContainer(id interface{}) (*CertStoreContainer, error) {
	log.Printf("[INFO] Fetching certificat store containers %s.\n", id)

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

// CreateStoreContainer creates a new certificate store container (pre-v25 legacy API).
func (c *Client) CreateStoreContainer(req *CertStoreContainer) (*CertStoreContainer, error) {
	log.Println("[INFO] Creating certificate store container.")

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "CertificateStoreContainers",
		Headers:  headers,
		Payload:  bytes.NewReader(payload),
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &CertStoreContainer{}
	if err = json.NewDecoder(resp.Body).Decode(jsonResp); err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// UpdateStoreContainer updates an existing certificate store container (pre-v25 legacy API).
func (c *Client) UpdateStoreContainer(id int, req *CertStoreContainer) (*CertStoreContainer, error) {
	log.Printf("[INFO] Updating certificate store container %d.\n", id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	req.Id = &id
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: fmt.Sprintf("CertificateStoreContainers/%d", id),
		Headers:  headers,
		Payload:  bytes.NewReader(payload),
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &CertStoreContainer{}
	if err = json.NewDecoder(resp.Body).Decode(jsonResp); err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// DeleteStoreContainer deletes a certificate store container by ID (pre-v25 legacy API).
func (c *Client) DeleteStoreContainer(id int) error {
	log.Printf("[INFO] Deleting certificate store container %d.\n", id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "DELETE",
		Endpoint: fmt.Sprintf("CertificateStoreContainers/%d", id),
		Headers:  headers,
	}

	_, err := c.sendRequest(keyfactorAPIStruct)
	return err
}
