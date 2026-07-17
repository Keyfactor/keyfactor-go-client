package api

import (
	"encoding/json"
	"fmt"
	"log"
)

// ListPAMProviders returns all PAM providers according to the provided filter and output parameters
func (c *Client) ListPAMProviders(query *GetPAMProviderQuery) (*[]ProviderResponseLegacy, error) {
	log.Println("[INFO] Listing all PAM providers")

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "PamProviders"
	if query != nil {
		queryParams := query.toQueryString()
		if queryParams != "" {
			endpoint = fmt.Sprintf("%s?%s", endpoint, queryParams)
		}
	}

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

	var jsonResp []ProviderResponseLegacy
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// GetPamProviderByName returns a specific PAM provider by name
func (c *Client) GetPamProviderByName(name string) (*ProviderResponseLegacy, error) {
	log.Printf("[INFO] Getting PAM provider with name: %s", name)

	query := &GetPAMProviderQuery{
		QueryString: fmt.Sprintf("Name eq '%s'", name),
	}
	providers, err := c.ListPAMProviders(query)
	if err != nil {
		return nil, err
	}

	if providers == nil || len(*providers) == 0 {
		return nil, fmt.Errorf("PAM provider with name '%s' not found", name)
	}
	return &(*providers)[0], nil
}

// GetPAMProvider returns a specific PAM provider by ID
func (c *Client) GetPAMProvider(id int) (*ProviderResponseLegacy, error) {
	log.Printf("[INFO] Getting PAM provider with ID: %d", id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := fmt.Sprintf("PamProviders/%d", id)
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

	var jsonResp ProviderResponseLegacy
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// CreatePAMProvider creates a new PAM provider with the associated properties
func (c *Client) CreatePAMProvider(provider *ProviderCreateRequest) (*ProviderResponseLegacy, error) {
	log.Printf("[INFO] Creating new PAM provider: %s", provider.Name)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "PamProviders",
		Headers:  headers,
		Payload:  provider,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp ProviderResponseLegacy
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// UpdatePAMProvider updates an existing PAM provider
func (c *Client) UpdatePAMProvider(provider *ProviderUpdateRequestLegacy) (*ProviderResponseLegacy, error) {
	log.Printf("[INFO] Updating PAM provider with ID: %d", provider.Id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: "PamProviders",
		Headers:  headers,
		Payload:  provider,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp ProviderResponseLegacy
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// DeletePAMProvider deletes a PAM provider by ID
func (c *Client) DeletePAMProvider(id int) error {
	log.Printf("[INFO] Deleting PAM provider with ID: %d", id)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := fmt.Sprintf("PamProviders/%d", id)
	keyfactorAPIStruct := &request{
		Method:   "DELETE",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  nil,
	}

	_, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return err
	}

	return nil
}

// ListLocalPAMEntries returns local PAM entries for the given PAM provider according to the provided filter
func (c *Client) ListLocalPAMEntries(providerId int, query *GetPAMProviderQuery) (*[]LocalPAMEntryResponse, error) {
	log.Printf("[INFO] Listing local PAM entries for provider ID: %d", providerId)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := fmt.Sprintf("PamProviders/Local/%d/Entries", providerId)
	if query != nil {
		queryParams := query.toQueryString()
		if queryParams != "" {
			endpoint = fmt.Sprintf("%s?%s", endpoint, queryParams)
		}
	}

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

	var jsonResp []LocalPAMEntryResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// CreateLocalPAMEntry creates a new local PAM entry for the given PAM provider
func (c *Client) CreateLocalPAMEntry(providerId int, entry *LocalPAMEntryCreateRequest) (
	*LocalPAMEntryResponse,
	error,
) {
	log.Printf("[INFO] Creating local PAM entry for provider ID: %d", providerId)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	endpoint := fmt.Sprintf("PamProviders/Local/%d/Entries", providerId)
	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  entry,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp LocalPAMEntryResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// UpdateLocalPAMEntry updates an existing local PAM entry for the given PAM provider
func (c *Client) UpdateLocalPAMEntry(
	providerId int,
	secretName string,
	entry *LocalPAMEntryUpdateRequest,
) (*LocalPAMEntryResponse, error) {
	log.Printf("[INFO] Updating local PAM entry '%s' for provider ID: %d", secretName, providerId)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
			{"Content-Type", "application/json"},
		},
	}

	endpoint := fmt.Sprintf("PamProviders/Local/%d/Entries/%s", providerId, secretName)
	keyfactorAPIStruct := &request{
		Method:   "PUT",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  entry,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	var jsonResp LocalPAMEntryResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// DeleteLocalPAMEntry deletes a local PAM entry for the given PAM provider
func (c *Client) DeleteLocalPAMEntry(providerId int, secretName string) error {
	log.Printf("[INFO] Deleting local PAM entry '%s' for provider ID: %d", secretName, providerId)

	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := fmt.Sprintf("PamProviders/Local/%d/Entries/%s", providerId, secretName)
	keyfactorAPIStruct := &request{
		Method:   "DELETE",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  nil,
	}

	_, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return err
	}

	return nil
}

// GetPAMProviderQuery represents query parameters for PAM provider listing
type GetPAMProviderQuery struct {
	QueryString   string
	PageReturned  int
	ReturnLimit   int
	SortField     string
	SortAscending int
}

// toQueryString converts query parameters to URL query string
func (q *GetPAMProviderQuery) toQueryString() string {
	if q == nil {
		return ""
	}

	params := ""
	if q.QueryString != "" {
		params += fmt.Sprintf("QueryString=%s&", q.QueryString)
	}
	if q.PageReturned > 0 {
		params += fmt.Sprintf("PageReturned=%d&", q.PageReturned)
	}
	if q.ReturnLimit > 0 {
		params += fmt.Sprintf("ReturnLimit=%d&", q.ReturnLimit)
	}
	if q.SortField != "" {
		params += fmt.Sprintf("SortField=%s&", q.SortField)
		// Only add SortAscending if SortField is provided
		params += fmt.Sprintf("SortAscending=%d&", q.SortAscending)
	}

	// Remove trailing '&'
	if len(params) > 0 && params[len(params)-1] == '&' {
		params = params[:len(params)-1]
	}

	return params
}
