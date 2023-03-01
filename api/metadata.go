package api

import (
	"context"
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

	jsonData, _ := json.Marshal(um.Metadata)
	var newReq keyfactor_command_client_api.ModelsMetadataUpdateRequest
	json.Unmarshal(jsonData, &newReq)

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, err := apiClient.CertificateApi.CertificateUpdateMetadata(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).MetadataUpdate(newReq).CollectionId(int32(um.CollectionId)).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("[ERROR] Something unexpected happened, PUT call to /Certificates/Metadata returned status %d", resp.StatusCode)
	}
	return nil
}

func (c *Client) GetAllMetadataFields() ([]MetadataField, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, _, err := apiClient.MetadataFieldApi.MetadataFieldGetAllMetadataFields(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp []MetadataField
	for i := range resp {
		mapResp, _ := resp[i].ToMap()
		jsonData, _ := json.Marshal(mapResp)
		var newMetField MetadataField
		json.Unmarshal(jsonData, &newMetField)
		newResp = append(newResp, newMetField)
	}

	return newResp, nil

}

//func mapTupleArrayToString(i []StringTuple) map[string]string {
//	temp := make(map[string]string, len(i)) // Create string-index-able interface array from tuple struct
//	for _, field := range i {
//		temp[field.Elem1] = field.Elem2
//	}
//	return temp
//}
