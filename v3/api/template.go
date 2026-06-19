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
)

// getTemplatesMaxPages is the maximum number of pages GetTemplates will fetch
// before aborting with an error. It is an unexported package-level var (not a
// const) so that tests can lower it without iterating thousands of times.
var getTemplatesMaxPages = 10000

// GetTemplate takes arguments for a template ID used to facilitate the retrieval
// of certificate template context. The primary query required to get certificate context is the template ID. A pointer
// to a GetTemplateResponse structure is returned, containing the template context.
func (c *Client) GetTemplate(Id interface{}) (*GetTemplateResponse, error) {
	if Id == 0 {
		return nil, errors.New("template id required to get template")
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "Templates/" + fmt.Sprintf("%d", Id) // Append ID to complete endpoint

	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  headers,
		Query:    nil,
		Payload:  nil,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &GetTemplateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, err
}

// GetTemplates asks Keyfactor for a complete list of known certificate templates,
// paginating automatically so that instances with more than the server's default
// page size (50) return all templates. A list of GetTemplateResponse structures
// is returned, containing the template context.
func (c *Client) GetTemplates() ([]GetTemplateResponse, error) {
	log.Println("[INFO] Listing certificate templates.")

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	const pageSize = 100
	var all []GetTemplateResponse
	var page int
	for page = 1; page <= getTemplatesMaxPages; page++ {
		keyfactorAPIStruct := &request{
			Method:   "GET",
			Endpoint: "Templates/",
			Headers:  headers,
			Query: &apiQuery{
				Query: []StringTuple{
					{"PageReturned", strconv.Itoa(page)},
					{"ReturnLimit", strconv.Itoa(pageSize)},
				},
			},
			Payload: nil,
		}

		resp, err := c.sendRequest(keyfactorAPIStruct)
		if err != nil {
			log.Printf("[ERROR] GetTemplates: request for page %d failed: %s", page, err)
			return nil, err
		}

		var pageResults []GetTemplateResponse
		decodeErr := json.NewDecoder(resp.Body).Decode(&pageResults)
		resp.Body.Close()
		if decodeErr != nil {
			log.Printf("[ERROR] GetTemplates: failed to decode page %d: %s", page, decodeErr)
			return nil, decodeErr
		}

		all = append(all, pageResults...)
		if len(pageResults) == 0 || len(pageResults) < pageSize {
			break
		}
	}

	if page > getTemplatesMaxPages {
		return nil, fmt.Errorf("GetTemplates: exceeded max pages (%d); server may be ignoring pagination", getTemplatesMaxPages)
	}

	log.Printf("[INFO] Listed %d certificate templates across %d page(s).", len(all), page)
	return all, nil
}

// UpdateTemplate takes arguments for a UpdateTemplateArg structure used to facilitate the modification
// of a certificate template. Required parameters for this function are elements of UpdateTemplateArg that can't be set to nil. A pointer
// to a UpdateTemplateResponse structure is returned, containing the template context.
func (c *Client) UpdateTemplate(uta *UpdateTemplateArg) (*UpdateTemplateResponse, error) {

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: "Templates/",
		Headers:  headers,
		Query:    nil,
		Payload:  uta,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &UpdateTemplateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, err
}
