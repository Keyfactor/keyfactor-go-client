package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Keyfactor/keyfactor-go-client-sdk/api/keyfactor"
	"log"
)

//type StringInt int32
//
//// UnmarshalJSON create a custom unmarshal for the StringInt
///// this helps us check the type of our value before unmarshalling it
//
//func (st *StringInt) UnmarshalJSON(b []byte) error {
//	//convert the bytes into an interface
//	//this will help us check the type of our value
//	//if it is a string that can be converted into an int we convert it
//	///otherwise we return an error
//	var item interface{}
//	if err := json.Unmarshal(b, &item); err != nil {
//		return err
//	}
//	switch v := item.(type) {
//	case int32:
//		*st = StringInt(v)
//	case float64:
//		*st = StringInt(int(v))
//	case string:
//		///here convert the string into
//		///an integer
//		i, err := strconv.Atoi(v)
//		if err != nil {
//			///the string might not be of integer type
//			///so return an error
//			return err
//
//		}
//		*st = StringInt(i)
//
//	}
//	return nil
//}

// GetCertificateStoreType takes arguments for a certificate store type ID or name and if found will return the certificate store type
func (c *Client) GetCertificateStoreType(id interface{}) (*CertificateStoreType, error) {
	switch id.(type) {
	case int:
		return c.GetCertificateStoreTypeById(id.(int))
	case string:
		return c.GetCertificateStoreTypeByName(id.(string))
	}

	return nil, errors.New("invalid type for id, must pass either string or integer")
}

// GetCertificateStoreTypeByName takes arguments for a certificate store type ID to facilitate a call to Keyfactor
// that retrieves certificate store context associated with a store type ID
func (c *Client) GetCertificateStoreTypeByName(name string) (*CertificateStoreType, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration()
	apiClient := keyfactor.NewAPIClient(configuration)

	resp, _, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeGetCertificateStoreType1(context.Background(), name).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp []CertificateStoreType
	for i, _ := range resp {
		var newCertType CertificateStoreType
		mapResp, _ := resp[i].ToMap()
		jsonData, _ := json.Marshal(mapResp)
		json.Unmarshal(jsonData, &newCertType)
		newResp = append(newResp, newCertType)
	}
	for _, v := range newResp {
		// TODO: Assumes that there really should only be one type with a given shortname but this is not guaranteed
		return &v, nil
	}
	return nil, errors.New("no certificate store type found with the given name")
}

// GetCertificateStoreTypeById takes arguments for a certificate store type ID to facilitate a call to Keyfactor
// that retrieves certificate store context associated with a store type ID
func (c *Client) GetCertificateStoreTypeById(id int) (*CertificateStoreType, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration()
	apiClient := keyfactor.NewAPIClient(configuration)

	resp, _, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeGetCertificateStoreType0(context.Background(), int32(id)).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp CertificateStoreType
	mapResp, _ := resp.ToMap()
	jsonData, _ := json.Marshal(mapResp)
	json.Unmarshal(jsonData, &newResp)

	return &newResp, nil
}

// ListCertificateStoreTypes takes no arguments and returns a list of certificate store types from Keyfactor.
func (c *Client) ListCertificateStoreTypes() (*[]CertificateStoreType, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration()
	apiClient := keyfactor.NewAPIClient(configuration)

	resp, _, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeGetTypes(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp []CertificateStoreType
	for i, _ := range resp {
		var newCertType CertificateStoreType
		mapResp, _ := resp[i].ToMap()
		jsonData, _ := json.Marshal(mapResp)
		json.Unmarshal(jsonData, &newCertType)
		newResp = append(newResp, newCertType)
	}
	return &newResp, nil
}

// CreateStoreType takes arguments for CreateStoreFctArgs to facilitate the creation
// of all store types supported by a customer Keyfactor Command instance. Note that various certificate
// store types require different property arguments, and careful attention should be taken to ensure that
// all required elements are included. Required arguments for this method are:
//   - ClientMachine : string
//   - StorePath     : string
//   - Properties    : []StringTuple *Note - Method converts this array of StringTuples to a JSON string if provided
//   - AgentId       : string
func (c *Client) CreateStoreType(ca *CertificateStoreType) (*CertificateStoreType, error) {
	log.Println("[INFO] Creating new certificate store type with Keyfactor")

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration()
	apiClient := keyfactor.NewAPIClient(configuration)

	var newReq keyfactor.KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest
	jsonData, _ := json.Marshal(ca)
	err := json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return nil, err
	}

	resp, _, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeCreateCertificateStoreType(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CertStoreType(newReq).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()
	if err != nil {
		return nil, err
	}

	var newResp CertificateStoreType
	mapResp, mErr := resp.ToMap()
	if mErr != nil {
		return nil, mErr
	}
	jsonData, jErr := json.Marshal(mapResp)
	if jErr != nil {
		return nil, jErr
	}
	json.Unmarshal(jsonData, &newResp)

	return &newResp, nil
}

func (c *Client) UpdateStoreType(ca *CertificateStoreType) (*CertificateStoreType, error) {
	log.Println("[INFO] Creating new certificate store type with Keyfactor")

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration()
	apiClient := keyfactor.NewAPIClient(configuration)

	var newReq keyfactor.KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest
	jsonData, _ := json.Marshal(ca)
	err := json.Unmarshal(jsonData, &newReq)
	if err != nil {
		return nil, err
	}

	resp, _, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeUpdateCertificateStoreType(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).CertStoreType(newReq).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp CertificateStoreType
	mapResp, _ := resp.ToMap()
	jsonData, _ = json.Marshal(mapResp)
	json.Unmarshal(jsonData, &newResp)

	return &newResp, nil
}
func (c *Client) DeleteCertificateStoreType(id int) (*DeleteStoreType, error) {
	log.Printf("[INFO] Attempting to delete certificate store type %d", id)

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor.NewConfiguration()
	apiClient := keyfactor.NewAPIClient(configuration)

	resp, err := apiClient.CertificateStoreTypeApi.CertificateStoreTypeDeleteCertificateStoreType(context.Background(), int32(id)).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 204 {
		return nil, fmt.Errorf("error deleting certificate store type %d. %s", id, resp.Body)
	}
	return &DeleteStoreType{ID: id}, nil
}
