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
	"io"
	"log"
	"net/http"
	"strconv"
)

// CreateStore takes arguments for CreateStoreFctArgs to facilitate the creation
// of all store types supported by a customer Keyfactor Command instance. Note that various certificate
// store types require different property arguments, and careful attention should be taken to ensure that
// all required elements are included. Required arguments for this method are:
//   - ClientMachine : string
//   - StorePath     : string
//   - Properties    : []StringTuple *Note - Method converts this array of StringTuples to a JSON string if provided
//   - AgentId       : string
func (c *Client) CreateStore(ca *CreateStoreFctArgs) (*CreateStoreResponse, error) {
	log.Println("[INFO] Creating new certificate store with Keyfactor")

	// Validate that the required fields are present
	err := validateCreateStoreArgs(ca)
	if err != nil {
		return nil, err
	}

	// API doesn't know what a StringTuple type is. Convert this type to an array of interfaces
	// that the JSON library can serialize. Then, serialize to JSON, and convert to string.
	if ca.PropertiesString == "" {
		propertiesInterface := ca.Properties
		propertiesJson, err := json.Marshal(propertiesInterface)
		if err != nil {
			return nil, err
		}
		ca.PropertiesString = string(propertiesJson)
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "CertificateStores",
		Headers:  headers,
		Payload:  &ca,
	}

	resp, respErr := c.sendRequest(keyfactorAPIStruct)
	if respErr != nil {
		return nil, respErr
	}

	jsonResp := &CreateStoreResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// UpdateStore takes arguments for UpdateStoreFctArgs to facilitate the adjustment of a certificate store
// associated with a Keyfactor Command instance. Note that various certificate
// store types require different property arguments, and careful attention should be taken to ensure that
// all required elements are included. Required arguments for this method are:
//   - ClientMachine : string
//   - StorePath     : string
//   - Properties    : []StringTuple *Note - Method converts this slice of StringTuples to a JSON string if provided
//   - AgentId       : string
func (c *Client) UpdateStore(ua *UpdateStoreFctArgs) (*UpdateStoreResponse, error) {
	log.Println("[INFO] Creating new certificate store with Keyfactor")

	// Validate that the required fields are present
	err := validateUpdateStoreArgs(ua)
	if err != nil {
		return nil, err
	}

	// API doesn't know what a StringTuple type is. Convert this type to an array of interfaces
	// that the JSON library can serialize. Then, serialize to JSON, and convert to string.
	if ua.PropertiesString == "" {
		propertiesInterface := ua.Properties
		propertiesJson, err := json.Marshal(propertiesInterface)
		if err != nil {
			return nil, err
		}
		ua.PropertiesString = string(propertiesJson)
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: "CertificateStores",
		Headers:  headers,
		Payload:  &ua,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &UpdateStoreResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// DeleteCertificateStore takes arguments for a certificate store ID to facilitate a call to Keyfactor
// that deletes a certificate store. Only the store ID is required.
func (c *Client) DeleteCertificateStore(storeId string) error {
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "CertificateStores/" + fmt.Sprintf("%s", storeId) // Append GUID to complete endpoint
	keyfactorAPIStruct := &request{
		Method:   "DELETE",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf(
			"[ERROR] Something unexpected happened, %s call to %s returned status %d",
			keyfactorAPIStruct.Method,
			keyfactorAPIStruct.Endpoint,
			resp.StatusCode,
		)
	}

	return nil
}

// ListCertificateStores takes no arguments and returns a slice of CertificateStore objects
// that represent all certificate stores associated with a Keyfactor Command instance.

func (c *Client) ListCertificateStores(params *map[string]interface{}) (*[]GetCertificateStoreResponse, error) {
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	query := apiQuery{
		Query: []StringTuple{},
	}
	if params != nil {
		sId, ok := (*params)["Id"]
		if ok && len(sId.([]string)) > 0 {
			var resp, err = c.GetCertificateStoreByID(fmt.Sprintf("%s", sId.([]string)[0]))
			if err != nil {
				return nil, err
			}
			return &[]GetCertificateStoreResponse{*resp}, nil
		}

		query, _ = buildQuery(*params, "certificateStoreQuery.queryString")
	}

	currentPage := 1
	var output []GetCertificateStoreResponse
	var errs []error
	// Loop through pages until no more results are returned
	for {
		query.Query = append(query.Query, StringTuple{"certificateStoreQuery.pageReturned", strconv.Itoa(currentPage)})
		endpoint := "CertificateStores/"
		keyfactorAPIStruct := &request{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  headers,
			Payload:  nil,
			Query:    &query,
		}

		resp, err := c.sendRequest(keyfactorAPIStruct)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf(
				"[ERROR] Something unexpected happened, %s call to %s returned status %d",
				keyfactorAPIStruct.Method,
				keyfactorAPIStruct.Endpoint,
				resp.StatusCode,
			)
		}
		var jsonResp []GetCertificateStoreResponse
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		if err != nil {
			errs = append(errs, err)
		}
		if len(jsonResp) == 0 {
			break
		}
		output = append(output, jsonResp...)
		currentPage++
	}

	// combine all errors into one
	if len(errs) > 0 {
		errStr := ""
		for _, e := range errs {
			errStr += e.Error() + "\n"
		}
		return nil, errors.New(errStr)
	}
	return &output, nil
}

// GetCertificateStoreByID takes arguments for a certificate store ID to facilitate a call to Keyfactor
// that retrieves a certificate store context. Only the store ID is required. A pointer to a GetStoreByIDResp struct
// is returned that contains information on the certificate store.
func (c *Client) GetCertificateStoreByID(storeId string) (*GetCertificateStoreResponse, error) {
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "CertificateStores/" + fmt.Sprintf("%s", storeId) // Append GUID to complete endpoint
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
	defer resp.Body.Close()

	bodyBytes, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	jsonResp := &GetCertificateStoreResponse{}
	if jErr := json.Unmarshal(bodyBytes, &jsonResp); jErr != nil {
		rawJson := make(map[string]interface{})
		if mErr := json.Unmarshal(bodyBytes, &rawJson); mErr != nil {
			return nil, fmt.Errorf("error decoding response: %v", mErr)
		}
		return nil, fmt.Errorf("error decoding response: %v, raw response: %v", jErr, rawJson)
	}
	jsonResp.Properties = unmarshalPropertiesString(jsonResp.PropertiesString)
	return jsonResp, nil
}

// GetCertificateStoreByContainerID takes arguments for a certificate store ID to facilitate a call to Keyfactor
// that retrieves a certificate store context. Only the store ID is required. A pointer to a GetStoreByIDResp struct
// is returned that contains information on the certificate store.
func (c *Client) GetCertificateStoreByContainerID(containerID interface{}) (*[]GetCertificateStoreResponse, error) {

	query := apiQuery{
		Query: []StringTuple{},
	}
	switch containerID.(type) {
	case int:
		query.Query = append(
			query.Query, StringTuple{
				"certificateStoreQuery.queryString", fmt.Sprintf(`ContainerId -eq "%d"`, containerID),
			},
		)
	case string:
		ct, ctErr := c.GetStoreContainer(containerID.(string))
		if ctErr != nil {
			return nil, ctErr
		}
		query.Query = append(
			query.Query, StringTuple{
				"certificateStoreQuery.queryString", fmt.Sprintf(`ContainerId -eq %d`, *ct.Id),
			},
		)
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "CertificateStores"

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

	jsonResp := &[]GetCertificateStoreResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	//jsonResp.Properties = unmarshalPropertiesString(jsonResp.PropertiesString)
	return jsonResp, nil
}

func (c *Client) GetCertificateStoreByClientAndStorePath(
	clientMachine string,
	storePath, containerID interface{},
) (*[]GetCertificateStoreResponse, error) {

	query := apiQuery{
		Query: []StringTuple{},
	}

	fullQueryString := ""
	switch containerID.(type) {
	case int, int64:
		contIdInt := int(containerID.(int64))
		if contIdInt > 0 {
			//query.Query = append(query.Query, StringTuple{
			//	"certificateStoreQuery.queryString", fmt.Sprintf(`ContainerId -eq "%d"`, containerID),
			//})
			//append to fullQueryString
			fullQueryString = fmt.Sprintf(`ContainerId -eq "%d"`, contIdInt)
		}
	case string:
		//ct, ctErr := c.GetStoreContainer(containerID.(string))
		//if ctErr != nil {
		//	return nil, ctErr
		//}
		//query.Query = append(query.Query, StringTuple{
		//	"certificateStoreQuery.queryString", fmt.Sprintf(`ContainerId -eq %d`, *ct.Id),
		//})
		//append to fullQueryString
		fullQueryString = fmt.Sprintf(`ContainerId -eq "%s"`, containerID)
	}

	if storePath != nil {
		if fullQueryString != "" {
			fullQueryString = fmt.Sprintf(`%s AND StorePath -eq "%s"`, fullQueryString, storePath)
		} else {
			fullQueryString = fmt.Sprintf(`StorePath -eq "%s"`, storePath)
		}
	}

	if clientMachine != "" {
		if fullQueryString != "" {
			fullQueryString = fmt.Sprintf(`%s AND ClientMachine -eq "%s"`, fullQueryString, clientMachine)
		} else {
			fullQueryString = fmt.Sprintf(`ClientMachine -eq "%s"`, clientMachine)
		}
	}

	if fullQueryString != "" {
		query.Query = append(
			query.Query, StringTuple{
				"certificateStoreQuery.queryString", fullQueryString,
			},
		)
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "CertificateStores"

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

	jsonResp := &[]GetCertificateStoreResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	//jsonResp.Properties = unmarshalPropertiesString(jsonResp.PropertiesString)
	return jsonResp, nil
}

// AddCertificateToStores takes argument for a AddCertificateToStore structure and is used to remove a configured certificate
// from one or more certificate stores.
func (c *Client) AddCertificateToStores(config *AddCertificateToStore) ([]string, error) {
	log.Printf("[INFO] Adding certificate with ID %d to one or more certificate stores", config.CertificateId)

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "CertificateStores/Certificates/Add",
		Headers:  headers,
		Payload:  &config,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp []string
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// RemoveCertificateFromStores takes argument for a RemoveCertificateFromStore structure, and is used to remove a certificate
// from one or more certificate stores.
func (c *Client) RemoveCertificateFromStores(config *RemoveCertificateFromStore) ([]string, error) {
	log.Println("[INFO] Removing certificate from one or more certificate stores")

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "CertificateStores/Certificates/Remove",
		Headers:  headers,
		Payload:  &config,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp []string
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

func (c *Client) GetCertStoreInventory(storeId string) (*[]CertStoreInventory, error) {
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	var fullInventory []CertStoreInventory
	hasMore := true
	page := 1
	for hasMore {
		endpoint := fmt.Sprintf("CertificateStores/%s/Inventory", storeId)
		keyfactorAPIStruct := &request{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  headers,
			Payload:  nil,
			Query: &apiQuery{
				Query: []StringTuple{
					{"query.pageReturned", strconv.Itoa(page)},
				},
			},
		}

		resp, err := c.sendRequest(keyfactorAPIStruct)
		if err != nil {
			return nil, err
		}
		var inv []interface{}
		jsonResp := inv
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		//err = json.Unmarshal(resp.Body, &jsonResp)
		if err != nil {
			return nil, err
		}
		var invResp []CertStoreInventory
		if len(jsonResp) == 0 {
			invResp = []CertStoreInventory{}
			hasMore = false
			continue
		} else {
			// Convert jsonResp to json string
			jsonString, convErr := json.Marshal(jsonResp)
			if convErr != nil {
				return nil, convErr
			}
			// Unmarshal json string into CertStoreInventory struct
			convErr = json.Unmarshal(jsonString, &invResp)
			if convErr != nil {
				return nil, convErr
			}
		}

		// iterate invResp and populate lookups Thumbprints, Serials and Ids
		for i, storeInv := range invResp {
			for _, cert := range storeInv.Certificates {
				invResp[i].Thumbprints = append(invResp[i].Thumbprints, cert.Thumbprint)
				invResp[i].Serials = append(invResp[i].Serials, cert.SerialNumber)
				invResp[i].Ids = append(invResp[i].Ids, cert.Id)
			}
		}
		fullInventory = append(fullInventory, invResp...)
		page++
	}

	//jsonResp.Properties = unmarshalPropertiesString(jsonResp.PropertiesString)
	return &fullInventory, nil
}

// unmarshalPropertiesString unmarshalls a JSON string and serializes it into an array of StringTuple.
func unmarshalPropertiesString(properties string) map[string]interface{} {
	if properties != "" {
		// First, unmarshal JSON properties string to []interface{}
		var tempInterface interface{}
		if err := json.Unmarshal([]byte(properties), &tempInterface); err != nil {
			return make(map[string]interface{})
		}
		// Then, iterate through each key:value pair and serialize into map[string]string
		//newMap := make(map[string]string)
		//for key, value := range tempInterface.(map[string]interface{}) {
		//	newMap[key] = value.(string)
		//}
		//return newMap
		return tempInterface.(map[string]interface{})
	}

	return make(map[string]interface{})
}

// ScheduleImmediateInventory triggers an immediate inventory job on the given certificate store.
// If the store already has any inventory schedule configured this is a no-op.
// When no schedule is present (common for newly-created stores that have never had inventory run)
// it sends a PUT with InventorySchedule.Immediate=true so the orchestrator runs inventory on its
// next check-in and updates Command's inventory record.
//
// Password and server-credential Properties (ServerUsername/Password/UseSsl) are not included in
// the PUT body because they are write-only and are not returned by the GET endpoint. With the
// Password field now tagged json:"Password,omitempty", nil is omitted from JSON entirely so
// Command does not receive a null and will preserve whatever password is already configured.
func (c *Client) ScheduleImmediateInventory(storeId string) error {
	storeResp, err := c.GetCertificateStoreByID(storeId)
	if err != nil {
		return fmt.Errorf("ScheduleImmediateInventory: could not read store %s: %w", storeId, err)
	}

	// No-op if any schedule is already configured.
	sched := storeResp.InventorySchedule
	if sched.Immediate != nil || sched.Interval != nil || sched.Daily != nil || sched.ExactlyOnce != nil {
		return nil
	}

	immediate := true
	_, err = c.UpdateStore(&UpdateStoreFctArgs{
		Id:            storeResp.Id,
		ClientMachine: storeResp.ClientMachine,
		StorePath:     storeResp.StorePath,
		CertStoreType: storeResp.CertStoreType,
		AgentId:       storeResp.AgentId,
		// Password intentionally nil (omitted from JSON via omitempty) — preserves existing password.
		// PropertiesString intentionally empty (omitted from JSON via omitempty) — preserves existing properties.
		InventorySchedule: &InventorySchedule{
			Immediate: &immediate,
		},
	})
	if err != nil {
		return fmt.Errorf("ScheduleImmediateInventory: UpdateStore failed for store %s: %w", storeId, err)
	}
	return nil
}

func validateCreateStoreArgs(ca *CreateStoreFctArgs) error {
	if ca.ClientMachine == "" {
		return errors.New("client machine is required for creation of new certificate store")
	}
	if ca.StorePath == "" {
		return errors.New("store path is required for creation of new certificate store")
	}
	if ca.AgentId == "" {
		return errors.New("orchestrator agent id is required for creation of new certificate store")
	}

	return nil
}

func validateUpdateStoreArgs(ca *UpdateStoreFctArgs) error {
	if ca.ClientMachine == "" {
		return errors.New("client machine is required for creation of new certificate store")
	}
	if ca.StorePath == "" {
		return errors.New("store path is required for creation of new certificate store")
	}
	if ca.AgentId == "" {
		return errors.New("orchestrator agent id is required for creation of new certificate store")
	}

	return nil
}

// buildPropertiesInterface takes argument for an array of StringTuple and returns an interface of the associated values
// in map[string]interface{} elements.
func buildPropertiesInterface(properties map[string]string) interface{} {
	// Create temporary array of interfaces
	// When updating a property in Keyfactor, API expects {"key": {"value": "key-value"}} - Build this interface
	propertiesInterface := make(map[string]interface{})

	for key, value := range properties {
		inside := make(map[string]interface{}) // Create {"value": "<key-value>"} interface
		inside["value"] = value
		propertiesInterface[key] = inside // Create {"<key>": {"value": "key-value"}} interface
	}

	return propertiesInterface
}

func mapToEscapedJSONString(m map[string]interface{}) (string, error) {
	// Convert the map to a byte slice of JSON
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	// Escape any special characters in the JSON string
	escapedString := string(jsonBytes)

	return escapedString, nil
}
