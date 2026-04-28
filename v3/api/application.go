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
	"strconv"
	"strings"
)

// isLegacyContainerAPI returns true when the server is pre-v25 and uses
// CertificateStoreContainers instead of the v25+ Applications endpoint.
// Both endpoints accept and return the same JSON schedule format.
func (c *Client) isLegacyContainerAPI() bool {
	v := c.AuthClient.GetCommandVersion()
	major := commandVersionMajor(v)
	return major > 0 && major < 25
}

// commandVersionMajor parses the major version number from a product version
// string such as "24.4.0.0" or "25.1.0.0". Returns 0 if unparseable.
func commandVersionMajor(version string) int {
	if version == "" {
		return 0
	}
	parts := strings.SplitN(version, ".", 2)
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}
	return major
}

// appEndpoint returns the base API endpoint for the connected Command version.
// Pre-v25 uses CertificateStoreContainers; v25+ uses Applications.
// Both endpoints share the same JSON request/response format.
func (c *Client) appEndpoint() string {
	if c.isLegacyContainerAPI() {
		return "CertificateStoreContainers"
	}
	return "Applications"
}

// ListApplications returns all applications/containers, paginating automatically.
func (c *Client) ListApplications() ([]ApplicationListItem, error) {
	log.Println("[INFO] Listing applications.")

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	const pageSize = 100
	var all []ApplicationListItem
	for page := 1; ; page++ {
		req := &request{
			Method:   "GET",
			Endpoint: c.appEndpoint(),
			Headers:  headers,
			Query: &apiQuery{
				Query: []StringTuple{
					{"PageReturned", strconv.Itoa(page)},
					{"ReturnLimit", strconv.Itoa(pageSize)},
				},
			},
		}
		resp, err := c.sendRequest(req)
		if err != nil {
			return nil, err
		}
		var pageResults []ApplicationListItem
		if err = json.NewDecoder(resp.Body).Decode(&pageResults); err != nil {
			return nil, err
		}
		all = append(all, pageResults...)
		if len(pageResults) < pageSize {
			break
		}
	}
	return all, nil
}

// GetApplication returns the full details of an application/container by integer ID.
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
		Endpoint: fmt.Sprintf("%s/%d", c.appEndpoint(), id),
		Headers:  headers,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var result ApplicationResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
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

// CreateApplication creates a new application/container and returns the created resource.
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
		Endpoint: c.appEndpoint(),
		Headers:  headers,
		Payload:  createReq,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var result ApplicationResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateApplication performs a full replacement of an existing application/container.
// For v25+ the API uses PUT /Applications with the ID in the body.
// For pre-v25 the API uses PUT /CertificateStoreContainers/{id} with the ID in the path.
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

	// Both pre-v25 and v25+ use PUT /{endpoint} with the ID in the request body.
	endpoint := c.appEndpoint()

	req := &request{
		Method:   "PUT",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  updateReq,
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var result ApplicationResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteApplication deletes an application/container by its integer ID.
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
		Endpoint: fmt.Sprintf("%s/%d", c.appEndpoint(), id),
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
