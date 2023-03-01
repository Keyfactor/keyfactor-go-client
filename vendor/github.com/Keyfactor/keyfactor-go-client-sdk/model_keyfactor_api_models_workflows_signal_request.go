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

// checks if the KeyfactorApiModelsWorkflowsSignalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsWorkflowsSignalRequest{}

// KeyfactorApiModelsWorkflowsSignalRequest struct for KeyfactorApiModelsWorkflowsSignalRequest
type KeyfactorApiModelsWorkflowsSignalRequest struct {
	// The signal key. This is expected to be in a format like \"STEP_NAME.SIGNAL_NAME\"
	SignalKey *string `json:"SignalKey,omitempty"`
	// Arbitrary data to associate with the signal.
	Data map[string]map[string]interface{} `json:"Data,omitempty"`
}

// NewKeyfactorApiModelsWorkflowsSignalRequest instantiates a new KeyfactorApiModelsWorkflowsSignalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsWorkflowsSignalRequest() *KeyfactorApiModelsWorkflowsSignalRequest {
	this := KeyfactorApiModelsWorkflowsSignalRequest{}
	return &this
}

// NewKeyfactorApiModelsWorkflowsSignalRequestWithDefaults instantiates a new KeyfactorApiModelsWorkflowsSignalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsWorkflowsSignalRequestWithDefaults() *KeyfactorApiModelsWorkflowsSignalRequest {
	this := KeyfactorApiModelsWorkflowsSignalRequest{}
	return &this
}

// GetSignalKey returns the SignalKey field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsSignalRequest) GetSignalKey() string {
	if o == nil || isNil(o.SignalKey) {
		var ret string
		return ret
	}
	return *o.SignalKey
}

// GetSignalKeyOk returns a tuple with the SignalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsSignalRequest) GetSignalKeyOk() (*string, bool) {
	if o == nil || isNil(o.SignalKey) {
		return nil, false
	}
	return o.SignalKey, true
}

// HasSignalKey returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsSignalRequest) HasSignalKey() bool {
	if o != nil && !isNil(o.SignalKey) {
		return true
	}

	return false
}

// SetSignalKey gets a reference to the given string and assigns it to the SignalKey field.
func (o *KeyfactorApiModelsWorkflowsSignalRequest) SetSignalKey(v string) {
	o.SignalKey = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsSignalRequest) GetData() map[string]map[string]interface{} {
	if o == nil || isNil(o.Data) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsSignalRequest) GetDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.Data) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsSignalRequest) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]map[string]interface{} and assigns it to the Data field.
func (o *KeyfactorApiModelsWorkflowsSignalRequest) SetData(v map[string]map[string]interface{}) {
	o.Data = v
}

func (o KeyfactorApiModelsWorkflowsSignalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsWorkflowsSignalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SignalKey) {
		toSerialize["SignalKey"] = o.SignalKey
	}
	if !isNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsWorkflowsSignalRequest struct {
	value *KeyfactorApiModelsWorkflowsSignalRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsWorkflowsSignalRequest) Get() *KeyfactorApiModelsWorkflowsSignalRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsWorkflowsSignalRequest) Set(val *KeyfactorApiModelsWorkflowsSignalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsWorkflowsSignalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsWorkflowsSignalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsWorkflowsSignalRequest(val *KeyfactorApiModelsWorkflowsSignalRequest) *NullableKeyfactorApiModelsWorkflowsSignalRequest {
	return &NullableKeyfactorApiModelsWorkflowsSignalRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsWorkflowsSignalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsWorkflowsSignalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


