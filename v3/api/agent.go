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
	"regexp"
)

// GetAgentList returns a list of orchestrators registered in the Keyfactor instance
func (c *Client) GetAgentList() ([]Agent, error) {
	// 0
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: "Agents",
		Headers:  headers,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp []Agent
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

func (c *Client) GetAgent(id string) ([]Agent, error) {
	// 0
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

	if isGUID(id) {
		query.Query = append(
			query.Query, StringTuple{
				"pq.queryString", fmt.Sprintf("AgentId -eq \"%s\"", id),
			},
		)
	} else {
		query.Query = append(
			query.Query, StringTuple{
				"pq.queryString", fmt.Sprintf("ClientMachine -eq \"%s\"", id),
			},
		)
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: "Agents",
		Headers:  headers,
		Payload:  nil,
		Query:    &query,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp []Agent
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	if len(jsonResp) == 0 {
		return jsonResp, fmt.Errorf("agent %s not found", id)
	}
	return jsonResp, nil
}

func (c *Client) ApproveAgent(id string) (string, error) {
	// 0
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Agents/Approve",
		Headers:  headers,
		Payload:  &[]string{id},
	}

	resp, rErr := c.sendRequest(keyfactorAPIStruct)
	if rErr != nil {
		return "", rErr
	}

	if resp.StatusCode == 204 {
		return "Approve agent successful.", nil
	}

	var jsonResp string
	err := json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return "", err
	}
	return jsonResp, nil
}

func (c *Client) DisApproveAgent(id string) (string, error) {
	// 0
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "Agents/Disapprove",
		Headers:  headers,
		Payload:  &[]string{id},
	}

	resp, rErr := c.sendRequest(keyfactorAPIStruct)
	if rErr != nil {
		return "", rErr
	}

	if resp.StatusCode == 204 {
		return fmt.Sprintf("Disapproving %s successful.", id), nil
	}

	var jsonResp string
	err := json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return "", err
	}
	return jsonResp, nil
}

func (c *Client) ResetAgent(id string) (string, error) {
	// 0
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: fmt.Sprintf("Agents/%s/Reset", id),
		Headers:  headers,
		Payload:  nil,
	}

	resp, rErr := c.sendRequest(keyfactorAPIStruct)
	if rErr != nil {
		return "", rErr
	}

	if resp.StatusCode == 204 {
		return "Reset agent successful.", nil
	}
	var jsonResp string
	err := json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return "", err
	}
	return jsonResp, nil
}

func (c *Client) FetchAgentLogs(id string) (string, error) {
	// 0
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: fmt.Sprintf("Agents/%s/FetchLogs", id),
		Headers:  headers,
		Payload:  nil,
	}

	resp, rErr := c.sendRequest(keyfactorAPIStruct)
	if rErr != nil {
		return "", rErr
	}

	if resp.StatusCode == 204 {
		return "Fetch logs successful.", nil
	}
	var jsonResp string
	err := json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return "", err
	}
	return jsonResp, nil
}

func isGUID(input string) bool {
	guidPattern := `^[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$`
	match, _ := regexp.MatchString(guidPattern, input)
	return match
}
