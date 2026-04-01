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
	"log"
	"strconv"
	"strings"
)

// CreateEnrollmentPattern creates a new enrollment pattern with the provided properties
func (c *Client) CreateEnrollmentPattern(
	req *EnrollmentPatternCreateRequest,
	forceTemplateDefault ...bool,
) (*EnrollmentPatternResponse, error) {
	log.Println("[INFO] Creating enrollment pattern with Keyfactor")

	// Validate required fields
	var missingFields []string
	if req.Name == "" {
		missingFields = append(missingFields, "Name")
	}
	if req.Template == 0 {
		missingFields = append(missingFields, "Template")
	}

	if len(missingFields) > 0 {
		return nil, errors.New("Required field(s) missing: " + strings.Join(missingFields, ", "))
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	// Build URL with query parameters
	endpoint := "EnrollmentPatterns"
	if len(forceTemplateDefault) > 0 && forceTemplateDefault[0] {
		endpoint += "?forceTemplateDefault=true"
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  req,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &EnrollmentPatternResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return jsonResp, nil
}

// GetEnrollmentPatterns returns all enrollment patterns according to the provided filter and output parameters
func (c *Client) GetEnrollmentPatterns(params ...*EnrollmentPatternsQueryParams) ([]EnrollmentPatternResponse, error) {
	log.Println("[INFO] Fetching enrollment patterns from Keyfactor")

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	// Build URL with query parameters
	endpoint := "EnrollmentPatterns"
	var queryParams []string

	if len(params) > 0 && params[0] != nil {
		param := params[0]
		if param.QueryString != "" {
			queryParams = append(queryParams, "QueryString="+param.QueryString)
		}
		if param.PageReturned > 0 {
			queryParams = append(queryParams, "PageReturned="+strconv.Itoa(param.PageReturned))
		}
		if param.ReturnLimit > 0 {
			queryParams = append(queryParams, "ReturnLimit="+strconv.Itoa(param.ReturnLimit))
		}
		if param.SortField != "" {
			queryParams = append(queryParams, "SortField="+param.SortField)
		}
		if param.SortAscending != nil {
			queryParams = append(queryParams, "SortAscending="+strconv.Itoa(*param.SortAscending))
		}
	}

	if len(queryParams) > 0 {
		endpoint += "?" + strings.Join(queryParams, "&")
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  headers,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp []EnrollmentPatternResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return jsonResp, nil
}

// GetEnrollmentPattern returns the enrollment pattern associated with the provided ID
func (c *Client) GetEnrollmentPattern(id int) (*EnrollmentPatternResponse, error) {
	log.Printf("[INFO] Fetching enrollment pattern with ID %d from Keyfactor", id)

	if id <= 0 {
		return nil, errors.New("ID must be a positive integer")
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: fmt.Sprintf("EnrollmentPatterns/%d", id),
		Headers:  headers,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp EnrollmentPatternResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil
}

// UpdateEnrollmentPattern updates an enrollment pattern according to the provided properties and Keyfactor identifier
func (c *Client) UpdateEnrollmentPattern(
	id int,
	req *EnrollmentPatternRequest,
	forceTemplateDefault ...bool,
) (*EnrollmentPatternResponse, error) {
	log.Printf("[INFO] Updating enrollment pattern with ID %d in Keyfactor", id)

	if id <= 0 {
		return nil, errors.New("ID must be a positive integer")
	}

	// Validate required fields
	var missingFields []string
	if req.Name == "" {
		missingFields = append(missingFields, "Name")
	}

	if len(missingFields) > 0 {
		return nil, errors.New("Required field(s) missing: " + strings.Join(missingFields, ", "))
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	// Build URL with query parameters
	endpoint := fmt.Sprintf("EnrollmentPatterns/%d", id)
	if len(forceTemplateDefault) > 0 && forceTemplateDefault[0] {
		endpoint += "?forceTemplateDefault=true"
	}

	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  req,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp EnrollmentPatternResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil
}

// DeleteEnrollmentPattern deletes an enrollment pattern by ID
// Note: This method assumes DELETE is supported based on REST conventions,
// though it may not be explicitly defined in the provided schema
func (c *Client) DeleteEnrollmentPattern(id int) error {
	log.Printf("[INFO] Deleting enrollment pattern with ID %d from Keyfactor", id)

	if id <= 0 {
		return errors.New("ID must be a positive integer")
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "DELETE",
		Endpoint: fmt.Sprintf("EnrollmentPatterns/%d", id),
		Headers:  headers,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return err
	}

	// Check if the response indicates success (2xx status codes)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to delete enrollment pattern: HTTP %d", resp.StatusCode)
	}

	return nil
}
