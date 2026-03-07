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
	"fmt"
	"log"
)

// ListApplications returns all applications from the /Applications endpoint.
func (c *Client) ListApplications() ([]ApplicationListItem, error) {
	log.Println("[INFO] Listing applications.")

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	req := &request{
		Method:   "GET",
		Endpoint: "Applications",
		Headers:  headers,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var result []ApplicationListItem
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetApplication returns the full details of an application by its integer ID.
func (c *Client) GetApplication(id int) (*ApplicationResponse, error) {
	log.Printf("[INFO] Fetching application with ID %d.", id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	req := &request{
		Method:   "GET",
		Endpoint: fmt.Sprintf("Applications/%d", id),
		Headers:  headers,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var result ApplicationResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApplicationByName returns the application matching the given name by
// listing all applications and then fetching the matching one by ID.
// Returns an error if no application with that name exists.
func (c *Client) GetApplicationByName(name string) (*ApplicationResponse, error) {
	log.Printf("[INFO] Fetching application with name %q.", name)

	apps, err := c.ListApplications()
	if err != nil {
		return nil, err
	}

	for _, app := range apps {
		if app.Name == name {
			return c.GetApplication(app.Id)
		}
	}
	return nil, fmt.Errorf("application %q not found", name)
}

// CreateApplication creates a new application and returns the created resource.
func (c *Client) CreateApplication(createReq *ApplicationCreateRequest) (*ApplicationResponse, error) {
	log.Println("[INFO] Creating application.")

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	req := &request{
		Method:   "POST",
		Endpoint: "Applications",
		Headers:  headers,
		Payload:  createReq,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var result ApplicationResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateApplication performs a full replacement (PUT) of an existing application.
// The API uses PUT /Applications (base URL, ID in body), not PUT /Applications/{id}.
// The Id field in updateReq is set automatically from the id argument.
func (c *Client) UpdateApplication(id int, updateReq *ApplicationUpdateRequest) (*ApplicationResponse, error) {
	log.Printf("[INFO] Updating application with ID %d.", id)

	updateReq.Id = id

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	req := &request{
		Method:   "PUT",
		Endpoint: "Applications",
		Headers:  headers,
		Payload:  updateReq,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var result ApplicationResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteApplication deletes an application by its integer ID.
// The server returns 204 No Content on success.
func (c *Client) DeleteApplication(id int) error {
	log.Printf("[INFO] Deleting application with ID %d.", id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	req := &request{
		Method:   "DELETE",
		Endpoint: fmt.Sprintf("Applications/%d", id),
		Headers:  headers,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to delete application: HTTP %d", resp.StatusCode)
	}

	return nil
}
