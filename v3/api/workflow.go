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
)

func (c *Client) ListPendingCertificates(q map[string]string) ([]WorkflowCertificate, error) {
	return c.ListWorkflowCert("Pending")
}

func (c *Client) ListDeniedCertificates(q map[string]string) ([]WorkflowCertificate, error) {
	return c.ListWorkflowCert("Denied")
}

func (c *Client) ListExternalValidationPendingCertificates(q map[string]string) ([]WorkflowCertificate, error) {
	return c.ListWorkflowCert("ExternalValidation")
}

func (c *Client) ListWorkflowCert(endpoint string) ([]WorkflowCertificate, error) {
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
	query.Query = append(
		query.Query, StringTuple{
			"pagedQuery.returnLimit", "1000",
		},
	)

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: fmt.Sprintf("Workflow/Certificates/%s", endpoint),
		Headers:  headers,
		Query:    &query,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp []WorkflowCertificate
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, err

}
