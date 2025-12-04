// Copyright 2025 Keyfactor
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
	"encoding/json"
	"fmt"
	"log"
)

// ListPAMProviderTypes returns all PAM provider types in the Keyfactor instance
func (c *Client) ListPAMProviderTypes() (*[]ProviderTypeResponse, error) {
	log.Println("[INFO] Listing all PAM provider types")

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: "PamProviders/Types",
		Headers:  headers,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp []ProviderTypeResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

func (c *Client) GetPAMProviderTypeByName(name string) (*ProviderTypeResponse, error) {
	// list all provider types
	types, err := c.ListPAMProviderTypes()
	if err != nil {
		return nil, err
	}

	// find the provider type with the matching name
	for _, t := range *types {
		if t.Name != nil && *t.Name == name {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("PAM provider type with name '%s' not found", name)
}

// GetPAMProviderType returns a specific PAM provider type by ID
func (c *Client) GetPAMProviderType(id string) (*ProviderTypeResponse, error) {
	log.Printf("[INFO] Getting PAM provider type with ID: %s", id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := fmt.Sprintf("PamProviders/Types/%s", id)
	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp ProviderTypeResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// CreatePAMProviderType creates a new PAM provider type with the associated properties
func (c *Client) CreatePAMProviderType(providerType *ProviderTypeCreateRequest) (*ProviderTypeResponse, error) {
	log.Printf("[INFO] Creating new PAM provider type: %s", providerType.Name)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "PamProviders/Types",
		Headers:  headers,
		Payload:  providerType,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp ProviderTypeResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// DeletePAMProviderType deletes a PAM provider type by ID, as long as it's not currently in use
func (c *Client) DeletePAMProviderType(id string) error {
	log.Printf("[INFO] Deleting PAM provider type with ID: %s", id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := fmt.Sprintf("PamProviders/Types/%s", id)
	keyfactorAPIStruct := &request{
		Method:   "DELETE",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  nil,
	}

	_, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return err
	}

	return nil
}
