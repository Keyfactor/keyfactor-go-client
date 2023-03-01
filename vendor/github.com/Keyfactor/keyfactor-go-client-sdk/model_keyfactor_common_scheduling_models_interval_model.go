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

// checks if the KeyfactorCommonSchedulingModelsIntervalModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorCommonSchedulingModelsIntervalModel{}

// KeyfactorCommonSchedulingModelsIntervalModel struct for KeyfactorCommonSchedulingModelsIntervalModel
type KeyfactorCommonSchedulingModelsIntervalModel struct {
	Minutes *int32 `json:"Minutes,omitempty"`
}

// NewKeyfactorCommonSchedulingModelsIntervalModel instantiates a new KeyfactorCommonSchedulingModelsIntervalModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorCommonSchedulingModelsIntervalModel() *KeyfactorCommonSchedulingModelsIntervalModel {
	this := KeyfactorCommonSchedulingModelsIntervalModel{}
	return &this
}

// NewKeyfactorCommonSchedulingModelsIntervalModelWithDefaults instantiates a new KeyfactorCommonSchedulingModelsIntervalModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorCommonSchedulingModelsIntervalModelWithDefaults() *KeyfactorCommonSchedulingModelsIntervalModel {
	this := KeyfactorCommonSchedulingModelsIntervalModel{}
	return &this
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *KeyfactorCommonSchedulingModelsIntervalModel) GetMinutes() int32 {
	if o == nil || isNil(o.Minutes) {
		var ret int32
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorCommonSchedulingModelsIntervalModel) GetMinutesOk() (*int32, bool) {
	if o == nil || isNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *KeyfactorCommonSchedulingModelsIntervalModel) HasMinutes() bool {
	if o != nil && !isNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given int32 and assigns it to the Minutes field.
func (o *KeyfactorCommonSchedulingModelsIntervalModel) SetMinutes(v int32) {
	o.Minutes = &v
}

func (o KeyfactorCommonSchedulingModelsIntervalModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorCommonSchedulingModelsIntervalModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Minutes) {
		toSerialize["Minutes"] = o.Minutes
	}
	return toSerialize, nil
}

type NullableKeyfactorCommonSchedulingModelsIntervalModel struct {
	value *KeyfactorCommonSchedulingModelsIntervalModel
	isSet bool
}

func (v NullableKeyfactorCommonSchedulingModelsIntervalModel) Get() *KeyfactorCommonSchedulingModelsIntervalModel {
	return v.value
}

func (v *NullableKeyfactorCommonSchedulingModelsIntervalModel) Set(val *KeyfactorCommonSchedulingModelsIntervalModel) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorCommonSchedulingModelsIntervalModel) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorCommonSchedulingModelsIntervalModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorCommonSchedulingModelsIntervalModel(val *KeyfactorCommonSchedulingModelsIntervalModel) *NullableKeyfactorCommonSchedulingModelsIntervalModel {
	return &NullableKeyfactorCommonSchedulingModelsIntervalModel{value: val, isSet: true}
}

func (v NullableKeyfactorCommonSchedulingModelsIntervalModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorCommonSchedulingModelsIntervalModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


