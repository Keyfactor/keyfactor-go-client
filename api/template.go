package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Keyfactor/keyfactor-go-client-sdk"
)

// GetTemplate takes arguments for a template ID used to facilitate the retrieval
// of certificate template context. The primary query required to get certificate context is the template ID. A pointer
// to a GetTemplateResponse structure is returned, containing the template context.
func (c *Client) GetTemplate(Id interface{}) (*GetTemplateResponse, error) {
	if Id == 0 {
		return nil, errors.New("template id required to get template")
	}

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	newId := Id.(int32)

	resp, _, err := apiClient.TemplateApi.TemplateGetTemplate(context.Background(), newId).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp GetTemplateResponse
	mapResp, _ := resp.ToMap()
	jsonData, _ := json.Marshal(mapResp)
	json.Unmarshal(jsonData, &newResp)

	return &newResp, err
}

// GetTemplates asks Keyfactor for a complete list of known certificate templates. A list of
// GetTemplateResponse structures is returned, containing the template context.
func (c *Client) GetTemplates() ([]GetTemplateResponse, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	resp, _, err := apiClient.TemplateApi.TemplateGetTemplates(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newResp []GetTemplateResponse
	for i, _ := range resp {
		var newTemp GetTemplateResponse
		mapResp, _ := resp[i].ToMap()
		jsonData, _ := json.Marshal(mapResp)
		json.Unmarshal(jsonData, &newTemp)
		newResp = append(newResp, newTemp)
	}

	return newResp, err
}

// UpdateTemplate takes arguments for a UpdateTemplateArg structure used to facilitate the modification
// of a certificate template. Required parameters for this function are elements of UpdateTemplateArg that can't be set to nil. A pointer
// to a UpdateTemplateResponse structure is returned, containing the template context.
func (c *Client) UpdateTemplate(uta *UpdateTemplateArg) (*UpdateTemplateResponse, error) {

	xKeyfactorRequestedWith := "APIClient"
	xKeyfactorApiVersion := "1"

	configuration := keyfactor_command_client_api.NewConfiguration()
	apiClient := keyfactor_command_client_api.NewAPIClient(configuration)

	var newReq keyfactor_command_client_api.ModelsTemplateUpdateRequest
	jsonData, _ := json.Marshal(newReq)
	json.Unmarshal(jsonData, &newReq)

	resp, _, err := apiClient.TemplateApi.TemplateUpdateTemplate(context.Background()).XKeyfactorRequestedWith(xKeyfactorRequestedWith).Template(newReq).XKeyfactorApiVersion(xKeyfactorApiVersion).Execute()

	if err != nil {
		return nil, err
	}

	var newTemp GetTemplateResponse
	mapResp, _ := resp.ToMap()
	jsonData, _ = json.Marshal(mapResp)
	json.Unmarshal(jsonData, &newTemp)

	newResp := UpdateTemplateResponse{
		GetTemplateResponse: newTemp,
	}

	return &newResp, err
}
