package api

import (
	"encoding/json"
	"fmt"
	"log"
)

// GetAgentList returns a list of orchestrators registered in the Keyfactor instance
func (c *Client) GetAgentList() ([]Agent, error) {
	log.Println("[INFO] Getting a list of agents registered in Keyfactor")

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
	log.Println("[INFO] Getting agent by ID or name.")

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
	query.Query = append(query.Query, StringTuple{
		"pq.queryString", fmt.Sprintf(`ClientMachine -eq "%s"`, id),
	})

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
	log.Printf("[INFO] Approving agent %s in Keyfactor.\n", id)

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
	log.Printf("[INFO] Disapproving agent %s in Keyfactor.\n", id)

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
	log.Printf("[INFO] Resetting agent %s in Keyfactor.\n", id)

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
	log.Printf("[INFO] Fetching agent logs for %s.\n", id)

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
