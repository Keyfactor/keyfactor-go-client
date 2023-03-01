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

// checks if the KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest{}

// KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest struct for KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest
type KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest struct {
	Id int32 `json:"Id"`
	UseHandler bool `json:"UseHandler"`
}

// NewKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest instantiates a new KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest(id int32, useHandler bool) *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest {
	this := KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest{}
	this.Id = id
	this.UseHandler = useHandler
	return &this
}

// NewKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequestWithDefaults instantiates a new KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequestWithDefaults() *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest {
	this := KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest{}
	return &this
}

// GetId returns the Id field value
func (o *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) SetId(v int32) {
	o.Id = v
}

// GetUseHandler returns the UseHandler field value
func (o *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) GetUseHandler() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseHandler
}

// GetUseHandlerOk returns a tuple with the UseHandler field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) GetUseHandlerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseHandler, true
}

// SetUseHandler sets field value
func (o *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) SetUseHandler(v bool) {
	o.UseHandler = v
}

func (o KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["UseHandler"] = o.UseHandler
	return toSerialize, nil
}

type NullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest struct {
	value *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) Get() *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) Set(val *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest(val *KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) *NullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest {
	return &NullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


