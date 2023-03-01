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

// checks if the KeyfactorApiModelsEventHandlerEventHandlerParameterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsEventHandlerEventHandlerParameterResponse{}

// KeyfactorApiModelsEventHandlerEventHandlerParameterResponse struct for KeyfactorApiModelsEventHandlerEventHandlerParameterResponse
type KeyfactorApiModelsEventHandlerEventHandlerParameterResponse struct {
	Id *int32 `json:"Id,omitempty"`
	Key *string `json:"Key,omitempty"`
	DefaultValue *string `json:"DefaultValue,omitempty"`
	ParameterType *string `json:"ParameterType,omitempty"`
}

// NewKeyfactorApiModelsEventHandlerEventHandlerParameterResponse instantiates a new KeyfactorApiModelsEventHandlerEventHandlerParameterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsEventHandlerEventHandlerParameterResponse() *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse {
	this := KeyfactorApiModelsEventHandlerEventHandlerParameterResponse{}
	return &this
}

// NewKeyfactorApiModelsEventHandlerEventHandlerParameterResponseWithDefaults instantiates a new KeyfactorApiModelsEventHandlerEventHandlerParameterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsEventHandlerEventHandlerParameterResponseWithDefaults() *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse {
	this := KeyfactorApiModelsEventHandlerEventHandlerParameterResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) SetId(v int32) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) SetKey(v string) {
	o.Key = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) GetDefaultValue() string {
	if o == nil || isNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) GetDefaultValueOk() (*string, bool) {
	if o == nil || isNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) HasDefaultValue() bool {
	if o != nil && !isNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetParameterType returns the ParameterType field value if set, zero value otherwise.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) GetParameterType() string {
	if o == nil || isNil(o.ParameterType) {
		var ret string
		return ret
	}
	return *o.ParameterType
}

// GetParameterTypeOk returns a tuple with the ParameterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) GetParameterTypeOk() (*string, bool) {
	if o == nil || isNil(o.ParameterType) {
		return nil, false
	}
	return o.ParameterType, true
}

// HasParameterType returns a boolean if a field has been set.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) HasParameterType() bool {
	if o != nil && !isNil(o.ParameterType) {
		return true
	}

	return false
}

// SetParameterType gets a reference to the given string and assigns it to the ParameterType field.
func (o *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) SetParameterType(v string) {
	o.ParameterType = &v
}

func (o KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !isNil(o.DefaultValue) {
		toSerialize["DefaultValue"] = o.DefaultValue
	}
	if !isNil(o.ParameterType) {
		toSerialize["ParameterType"] = o.ParameterType
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse struct {
	value *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse) Get() *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse) Set(val *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse(val *KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) *NullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse {
	return &NullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsEventHandlerEventHandlerParameterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


