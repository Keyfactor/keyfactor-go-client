package api

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// GetAgentList returns a list of orchestrators registered in the Keyfactor instance
func (c *Client) GetAgentList() ([]Agent, error) {
	//log.println("[INFO] Getting a list of agents registered in Keyfactor")

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
	//log.println("[INFO] Getting agent by ID or name.")

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
	//log.printf("[INFO] Approving agent %s in Keyfactor.\n", id)

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
	//log.printf("[INFO] Disapproving agent %s in Keyfactor.\n", id)

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
	//log.printf("[INFO] Resetting agent %s in Keyfactor.\n", id)

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
	//log.printf("[INFO] Fetching agent logs for %s.\n", id)

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
