/*
Keyfactor-v1

This reference serves to document REST-based methods to manage and integrate with Keyfactor. In addition, an embedded interface allows for the execution of calls against the current Keyfactor API instance.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keyfactor_command_client_api

import (
	"encoding/json"
)

// checks if the ModelsSSLEndpointStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSLEndpointStatusRequest{}

// ModelsSSLEndpointStatusRequest struct for ModelsSSLEndpointStatusRequest
type ModelsSSLEndpointStatusRequest struct {
	Id string `json:"Id"`
	Status bool `json:"Status"`
}

// NewModelsSSLEndpointStatusRequest instantiates a new ModelsSSLEndpointStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSLEndpointStatusRequest(id string, status bool) *ModelsSSLEndpointStatusRequest {
	this := ModelsSSLEndpointStatusRequest{}
	this.Id = id
	this.Status = status
	return &this
}

// NewModelsSSLEndpointStatusRequestWithDefaults instantiates a new ModelsSSLEndpointStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSLEndpointStatusRequestWithDefaults() *ModelsSSLEndpointStatusRequest {
	this := ModelsSSLEndpointStatusRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ModelsSSLEndpointStatusRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpointStatusRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelsSSLEndpointStatusRequest) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *ModelsSSLEndpointStatusRequest) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModelsSSLEndpointStatusRequest) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModelsSSLEndpointStatusRequest) SetStatus(v bool) {
	o.Status = v
}

func (o ModelsSSLEndpointStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSLEndpointStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["Status"] = o.Status
	return toSerialize, nil
}

type NullableModelsSSLEndpointStatusRequest struct {
	value *ModelsSSLEndpointStatusRequest
	isSet bool
}

func (v NullableModelsSSLEndpointStatusRequest) Get() *ModelsSSLEndpointStatusRequest {
	return v.value
}

func (v *NullableModelsSSLEndpointStatusRequest) Set(val *ModelsSSLEndpointStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSLEndpointStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSLEndpointStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSLEndpointStatusRequest(val *ModelsSSLEndpointStatusRequest) *NullableModelsSSLEndpointStatusRequest {
	return &NullableModelsSSLEndpointStatusRequest{value: val, isSet: true}
}

func (v NullableModelsSSLEndpointStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSLEndpointStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


