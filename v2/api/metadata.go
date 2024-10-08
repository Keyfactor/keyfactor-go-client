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
	"net/http"
)

// UpdateMetadata takes arguments for UpdateMetadataArgs to facilitate the
// updating of metadata fields in Keyfactor. It returns nil upon successful revocation,
// and an error if not. Required fields to update certificate metadata are:
//   - CertID              : int
//   - CertificateMetadata : []CertificateMetadata OR Metadata : map[string]string
//
// UpdateMetadata sets the metadata associated with a certificate EXACTLY. IE; if
// CertificateMetadata or Metadata are blank, any metadata associated with a certificate
// will be erased.
func (c *Client) UpdateMetadata(um *UpdateMetadataArgs) error {
	// Metadata in Keyfactor varies between deployments
	// Instead of hard coding possibilities, take array of string tuple types and create
	// string-indexed array of interfaces for JSON compilation
	metadata := make(map[string]interface{})
	if um.Metadata == nil {
		if len(um.CertificateMetadata) > 0 {
			metadata = mapTupleArrayToInterface(um.CertificateMetadata)
		} else {
			return fmt.Errorf("metadata is required to update metadata. configure CertificateMetadata or Metadata")
		}
	} else {
		metadata = um.Metadata
	}

	query := apiQuery{
		Query: []StringTuple{},
	}

	if um.CollectionId > 0 {
		query.Query = append(query.Query, StringTuple{"collectionId", fmt.Sprintf("%d", um.CollectionId)})
	}

	fields, err := c.GetAllMetadataFields()
	if err != nil {
		return err
	}

	// Build new interface with every metadata field
	allFields := make(map[string]interface{})
	for _, field := range fields {
		value, ok := metadata[field.Name]
		if ok {
			allFields[field.Name] = value
		} else {
			allFields[field.Name] = ""
		}
	}
	// Then, set um.Metadata back to this new field
	um.Metadata = allFields

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: "Certificates/Metadata",
		Headers:  headers,
		Payload:  um,
		Query:    &query,
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

func (c *Client) GetAllMetadataFields() ([]MetadataField, error) {

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: "MetadataFields",
		Headers:  headers,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp []MetadataField
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, err

}
