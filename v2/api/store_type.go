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
	"encoding/json"
	"errors"
	"fmt"
)

// GetCertificateStoreType takes arguments for a certificate store type ID or name and if found will return the certificate store type
func (c *Client) GetCertificateStoreType(id interface{}) (*CertificateStoreType, error) {
	switch id.(type) {
	case int:
		return c.GetCertificateStoreTypeById(id.(int))
	case string:
		return c.GetCertificateStoreTypeByName(id.(string))
	}

	return nil, errors.New("invalid type for id, must pass either string or integer")
}

// GetCertificateStoreTypeByName takes arguments for a certificate store type ID to facilitate a call to Keyfactor
// that retrieves certificate store context associated with a store type ID
func (c *Client) GetCertificateStoreTypeByName(name string) (*CertificateStoreType, error) {
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := fmt.Sprintf("CertificateStoreTypes/Name/%s", name)
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

	var jsonResp []CertificateStoreType
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	for _, v := range jsonResp {
		// TODO: Assumes that there really should only be one type with a given shortname but this is not guaranteed
		return &v, nil
	}
	return nil, errors.New("no certificate store type found with the given name")
}

// GetCertificateStoreTypeById takes arguments for a certificate store type ID to facilitate a call to Keyfactor
// that retrieves certificate store context associated with a store type ID
func (c *Client) GetCertificateStoreTypeById(id int) (*CertificateStoreType, error) {
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := fmt.Sprintf("CertificateStoreTypes/%d", id)
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

	jsonResp := CertificateStoreType{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// ListCertificateStoreTypes takes no arguments and returns a list of certificate store types from Keyfactor.
func (c *Client) ListCertificateStoreTypes() (*[]CertificateStoreType, error) {
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "CertificateStoreTypes"
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

	var jsonResp []CertificateStoreType
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// CreateStoreType takes arguments for CreateStoreFctArgs to facilitate the creation
// of all store types supported by a customer Keyfactor Command instance. Note that various certificate
// store types require different property arguments, and careful attention should be taken to ensure that
// all required elements are included. Required arguments for this method are:
//   - ClientMachine : string
//   - StorePath     : string
//   - Properties    : []StringTuple *Note - Method converts this array of StringTuples to a JSON string if provided
//   - AgentId       : string
func (c *Client) CreateStoreType(ca *CertificateStoreType) (*CertificateStoreType, error) {
	//log.println("[INFO] Creating new certificate store type with Keyfactor")

	// Validate that the required fields are present
	//err := validateCreateStoreTypeArgs(ca)
	//if err != nil {
	//	return nil, err
	//}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "CertificateStoreTypes",
		Headers:  headers,
		Payload:  &ca,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &CertificateStoreType{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

func (c *Client) UpdateStoreType(ca *CertificateStoreType) (*CertificateStoreType, error) {
	//log.println("[INFO] Creating new certificate store type with Keyfactor")

	// Validate that the required fields are present
	//err := validateCreateStoreTypeArgs(ca)
	//if err != nil {
	//	return nil, err
	//}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: "CertificateStoreTypes",
		Headers:  headers,
		Payload:  &ca,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &CertificateStoreType{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}
func (c *Client) DeleteCertificateStoreType(id int) (*DeleteStoreType, error) {
	//log.printf("[INFO] Attempting to delete certificate store type %d", id)

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "DELETE",
		Endpoint: fmt.Sprintf("CertificateStoreTypes/%d", id),
		Headers:  headers,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 204 {
		return nil, fmt.Errorf("error deleting certificate store type %d. %s", id, resp.Body)
	}
	return &DeleteStoreType{ID: id}, nil
}
