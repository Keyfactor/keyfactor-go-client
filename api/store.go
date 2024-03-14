package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	kfc "github.com/Keyfactor/keyfactor-go-client-sdk/api/keyfactor"
)

// CreateStore takes arguments for CreateStoreFctArgs to facilitate the creation
// of all store types supported by a customer Keyfactor Command instance. Note that various certificate
// store types require different property arguments, and careful attention should be taken to ensure that
// all required elements are included. Required arguments for this method are:
//   - ClientMachine : string
//   - StorePath     : string
//   - Properties    : []StringTuple *Note - Method converts this array of StringTuples to a JSON string if provided
//   - AgentId       : string
//
// TODO?
func (c *Client) CreateStore(ca *CreateStoreFctArgs) (*CreateStoreResponse, error) {
	log.Println("[INFO] Creating new certificate store with Keyfactor")

	// Validate that the required fields are present
	err := validateCreateStoreArgs(ca)
	if err != nil {
		return nil, err
	}

	// API doesn't know what a StringTuple type is. Convert this type to an array of interfaces
	// that the JSON library can serialize. Then, serialize to JSON, and convert to string.
	if ca.PropertiesString == "" {
		propertiesInterface := ca.Properties
		propertiesJson, err := json.Marshal(propertiesInterface)
		if err != nil {
			return nil, err
		}
		ca.PropertiesString = string(propertiesJson)
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "POST",
		Endpoint: "CertificateStores",
		Headers:  headers,
		Payload:  &ca,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &CreateStoreResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// UpdateStore takes arguments for UpdateStoreFctArgs to facilitate the adjustment of a certificate store
// associated with a Keyfactor Command instance. Note that various certificate
// store types require different property arguments, and careful attention should be taken to ensure that
// all required elements are included. Required arguments for this method are:
//   - ClientMachine : string
//   - StorePath     : string
//   - Properties    : []StringTuple *Note - Method converts this slice of StringTuples to a JSON string if provided
//   - AgentId       : string
//
// TODO?
func (c *Client) UpdateStore(ua *UpdateStoreFctArgs) (*UpdateStoreResponse, error) {
	log.Println("[INFO] Creating new certificate store with Keyfactor")

	// Validate that the required fields are present
	err := validateUpdateStoreArgs(ua)
	if err != nil {
		return nil, err
	}

	// API doesn't know what a StringTuple type is. Convert this type to an array of interfaces
	// that the JSON library can serialize. Then, serialize to JSON, and convert to string.
	if ua.PropertiesString == "" {
		propertiesInterface := ua.Properties
		propertiesJson, err := json.Marshal(propertiesInterface)
		if err != nil {
			return nil, err
		}
		ua.PropertiesString = string(propertiesJson)
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	keyfactorAPIStruct := &request{
		Method:   "Put",
		Endpoint: "CertificateStores",
		Headers:  headers,
		Payload:  &ua,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &UpdateStoreResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// DeleteCertificateStore takes arguments for a certificate store ID to facilitate a call to Keyfactor
// that deletes a certificate store. Only the store ID is required.
func (c *Client) DeleteCertificateStore(storeId string) error {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := kfc.NewConfiguration(make(map[string]string))
	apiClient := kfc.NewAPIClient(configuration)

	resp, err := apiClient.CertificateStoreApi.CertificateStoreDeleteCertificateStore(context.Background(), storeId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("[ERROR] Something unexpected happened, DELETE call to /Certificate/Store/{id} returned status %d", resp.StatusCode)
	}

	return nil
}

// ListCertificateStores takes no arguments and returns a slice of CertificateStore objects
// that represent all certificate stores associated with a Keyfactor Command instance.

// TODO?
func (c *Client) ListCertificateStores(params *map[string]interface{}) (*[]GetCertificateStoreResponse, error) {
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
	if params != nil {
		sId, ok := (*params)["Id"]
		if ok {
			var resp, err = c.GetCertificateStoreByID(fmt.Sprintf("%s", sId.([]string)[0]))
			if err != nil {
				return nil, err
			}
			return &[]GetCertificateStoreResponse{*resp}, nil
		}

		query, _ = buildQuery(*params, "certificateStoreQuery.queryString")
	}

	endpoint := "CertificateStores/"
	keyfactorAPIStruct := &request{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  headers,
		Payload:  nil,
		Query:    &query,
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return &[]GetCertificateStoreResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return &[]GetCertificateStoreResponse{}, fmt.Errorf("[ERROR] Something unexpected happened, %s call to %s returned status %d", keyfactorAPIStruct.Method, keyfactorAPIStruct.Endpoint, resp.StatusCode)
	}
	var jsonResp []GetCertificateStoreResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

// GetCertificateStoreByID takes arguments for a certificate store ID to facilitate a call to Keyfactor
// that retrieves a certificate store context. Only the store ID is required. A pointer to a GetStoreByIDResp struct
// is returned that contains information on the certificate store.
// TODO?
func (c *Client) GetCertificateStoreByID(storeId string) (*GetCertificateStoreResponse, error) {
	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "CertificateStores/" + fmt.Sprintf("%s", storeId) // Append GUID to complete endpoint
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

	jsonResp := &GetCertificateStoreResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	jsonResp.Properties = unmarshalPropertiesString(jsonResp.PropertiesString)
	return jsonResp, nil
}

// GetCertificateStoreByID takes arguments for a certificate store ID to facilitate a call to Keyfactor
// that retrieves a certificate store context. Only the store ID is required. A pointer to a GetStoreByIDResp struct
// is returned that contains information on the certificate store.
// TODO?
func (c *Client) GetCertificateStoreByContainerID(containerID interface{}) (*[]GetCertificateStoreResponse, error) {

	query := apiQuery{
		Query: []StringTuple{},
	}
	switch containerID.(type) {
	case int:
		query.Query = append(query.Query, StringTuple{
			"certificateStoreQuery.queryString", fmt.Sprintf(`ContainerId -eq "%d"`, containerID),
		})
	case string:
		ct, ctErr := c.GetStoreContainer(containerID.(string))
		if ctErr != nil {
			return nil, ctErr
		}
		query.Query = append(query.Query, StringTuple{
			"certificateStoreQuery.queryString", fmt.Sprintf(`ContainerId -eq %d`, *ct.Id),
		})
	}

	// Set Keyfactor-specific headers
	headers := &apiHeaders{
		Headers: []StringTuple{
			{"x-keyfactor-api-version", "1"},
			{"x-keyfactor-requested-with", "APIClient"},
		},
	}

	endpoint := "CertificateStores"

	var keyfactorAPIStruct *request
	if query.Query != nil {
		keyfactorAPIStruct = &request{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  headers,
			Query:    &query,
		}
	} else {
		keyfactorAPIStruct = &request{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  headers,
		}
	}

	resp, err := c.sendRequest(keyfactorAPIStruct)
	if err != nil {
		return nil, err
	}

	jsonResp := &[]GetCertificateStoreResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}
	//jsonResp.Properties = unmarshalPropertiesString(jsonResp.PropertiesString)
	return jsonResp, nil
}

// AddCertificateToStores takes argument for a AddCertificateToStore structure and is used to add a configured certificate
// from one or more certificate stores.
func (c *Client) AddCertificateToStores(config *AddCertificateToStore) ([]string, error) {
	log.Printf("[INFO] Adding certificate with ID %d to one or more certificate stores", config.CertificateId)

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := kfc.NewConfiguration(make(map[string]string))
	apiClient := kfc.NewAPIClient(configuration)

	newCollectionId := int32(config.CollectionId)
	var newCertStoresList []kfc.ModelsCertificateStoreEntry
	for _, cert := range *config.CertificateStores {
		newProvider := int32(cert.EntryPassword.Provider)
		var newParams map[string]string
		data, _ := json.Marshal(cert.EntryPassword.Parameters)
		json.Unmarshal(data, &newParams)
		var newEntryPassword = kfc.ModelsKeyfactorAPISecret{
			SecretValue: &cert.EntryPassword.SecretValue,
			Parameters:  &newParams,
			Provider:    &newProvider,
		}
		var newCert = kfc.ModelsCertificateStoreEntry{
			CertificateStoreId: cert.CertificateStoreId,
			Alias:              &cert.Alias,
			JobFields:          nil,
			Overwrite:          &cert.Overwrite,
			EntryPassword:      &newEntryPassword,
			PfxPassword:        nil,
			IncludePrivateKey:  nil,
		}
		newCertStoresList = append(newCertStoresList, newCert)
	}

	jsonInvSched, _ := json.Marshal(config.InventorySchedule)
	var newSchedule kfc.KeyfactorCommonSchedulingKeyfactorSchedule
	json.Unmarshal(jsonInvSched, newSchedule)
	var newReq = kfc.KeyfactorApiModelsCertificateStoresAddCertificateRequest{
		CertificateId:     int32(config.CertificateId),
		CertificateStores: newCertStoresList,
		Schedule:          newSchedule,
		CollectionId:      &newCollectionId,
	}

	resp, _, err := apiClient.CertificateStoreApi.CertificateStoreAddCertificate(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).AddRequest(newReq).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveCertificateFromStores takes argument for a RemoveCertificateFromStore structure, and is used to remove a certificate
// from one or more certificate stores.
func (c *Client) RemoveCertificateFromStores(config *RemoveCertificateFromStore) ([]string, error) {
	log.Println("[INFO] Removing certificate from one or more certificate stores")

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := kfc.NewConfiguration(make(map[string]string))
	apiClient := kfc.NewAPIClient(configuration)

	newCollectionId := int32(config.CollectionId)
	var newCertStoresList []kfc.ModelsCertificateLocationSpecifier
	for _, cert := range *config.CertificateStores {
		var newCert = kfc.ModelsCertificateLocationSpecifier{
			Alias:              &cert.Alias,
			CertificateStoreId: &cert.CertificateStoreId,
			JobFields:          nil,
		}
		newCertStoresList = append(newCertStoresList, newCert)
	}

	jsonInvSched, _ := json.Marshal(config.InventorySchedule)
	var newSchedule kfc.KeyfactorCommonSchedulingKeyfactorSchedule
	json.Unmarshal(jsonInvSched, newSchedule)
	var newReq = kfc.KeyfactorApiModelsCertificateStoresRemoveCertificateRequest{
		CertificateStores: newCertStoresList,
		Schedule:          newSchedule,
		CollectionId:      &newCollectionId,
	}

	resp, _, err := apiClient.CertificateStoreApi.CertificateStoreRemoveCertificate(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).RemovalRequest(newReq).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetCertStoreInventory(storeId string) (*[]CertStoreInventory, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := kfc.NewConfiguration(make(map[string]string))
	apiClient := kfc.NewAPIClient(configuration)

	resp, _, err := apiClient.CertificateStoreApi.CertificateStoreGetCertificateStoreInventory(context.Background(), storeId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp []CertStoreInventory

	if len(resp) == 0 {
		newResp = []CertStoreInventory{}
	} else {
		for _, certInv := range resp {
			var newInvCertList []InventoriedCertificate
			var newParams = make(map[string]interface{})
			for _, param := range certInv.Parameters {
				for key, value := range param {
					newParams[key] = value
				}
			}
			for _, storedCert := range certInv.Certificates {
				var newInvCert = InventoriedCertificate{
					Id:                       int(*storedCert.Id),
					IssuedDN:                 *storedCert.IssuedDN.Get(),
					SerialNumber:             *storedCert.SerialNumber,
					NotBefore:                storedCert.NotBefore.String(),
					NotAfter:                 storedCert.NotAfter.String(),
					SigningAlgorithm:         *storedCert.SigningAlgorithm,
					IssuerDN:                 *storedCert.IssuerDN.Get(),
					Thumbprint:               *storedCert.Thumbprint,
					CertStoreInventoryItemId: int(*storedCert.CertStoreInventoryItemId),
				}
				newInvCertList = append(newInvCertList, newInvCert)
			}
			var newInv = CertStoreInventory{
				CertStoreInventoryItemId: 0,
				Name:                     *certInv.Name,
				Certificates:             newInvCertList,
				Thumbprints:              nil,
				Serials:                  nil,
				Ids:                      nil,
				Properties:               nil,
				Parameters:               newParams,
			}
			newResp = append(newResp, newInv)
		}
	}

	//jsonResp.Properties = unmarshalPropertiesString(jsonResp.PropertiesString)
	return &newResp, nil
}

// unmarshalPropertiesString unmarshalls a JSON string and serializes it into an array of StringTuple.
func unmarshalPropertiesString(properties string) map[string]interface{} {
	if properties != "" {
		// First, unmarshal JSON properties string to []interface{}
		var tempInterface interface{}
		if err := json.Unmarshal([]byte(properties), &tempInterface); err != nil {
			return make(map[string]interface{})
		}
		// Then, iterate through each key:value pair and serialize into map[string]string
		newMap := make(map[string]interface{})
		for key, value := range tempInterface.(map[string]interface{}) {
			newMap[key] = value
		}
		return newMap
	}

	return make(map[string]interface{})
}

func validateCreateStoreArgs(ca *CreateStoreFctArgs) error {
	if ca.ClientMachine == "" {
		return errors.New("client machine is required for creation of new certificate store")
	}
	if ca.StorePath == "" {
		return errors.New("store path is required for creation of new certificate store")
	}
	if ca.AgentId == "" {
		return errors.New("orchestrator agent id is required for creation of new certificate store")
	}

	return nil
}

func validateUpdateStoreArgs(ca *UpdateStoreFctArgs) error {
	if ca.ClientMachine == "" {
		return errors.New("client machine is required for creation of new certificate store")
	}
	if ca.StorePath == "" {
		return errors.New("store path is required for creation of new certificate store")
	}
	if ca.AgentId == "" {
		return errors.New("orchestrator agent id is required for creation of new certificate store")
	}

	return nil
}

// buildPropertiesInterface takes argument for an array of StringTuple and returns an interface of the associated values
// in map[string]interface{} elements.
func buildPropertiesInterface(properties map[string]interface{}) interface{} {
	// Create temporary array of interfaces
	// When updating a property in Keyfactor, API expects {"key": {"value": "key-value"}} - Build this interface
	propertiesInterface := make(map[string]interface{})

	for key, value := range properties {
		inside := make(map[string]interface{}) // Create {"value": "<key-value>"} interface
		inside["value"] = value                // TODO: put plain string if not an interface
		propertiesInterface[key] = inside      // Create {"<key>": {"value": "key-value"}} interface
	}

	return propertiesInterface
}
