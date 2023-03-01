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

// checks if the KeyfactorApiModelsWorkflowsAvailableSignalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsWorkflowsAvailableSignalResponse{}

// KeyfactorApiModelsWorkflowsAvailableSignalResponse struct for KeyfactorApiModelsWorkflowsAvailableSignalResponse
type KeyfactorApiModelsWorkflowsAvailableSignalResponse struct {
	// The name of the signal.
	SignalName *string `json:"SignalName,omitempty"`
	// The signal Id.
	StepSignalId *string `json:"StepSignalId,omitempty"`
	// Whether or not the signal has been received.
	SignalReceived *bool `json:"SignalReceived,omitempty"`
}

// NewKeyfactorApiModelsWorkflowsAvailableSignalResponse instantiates a new KeyfactorApiModelsWorkflowsAvailableSignalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsWorkflowsAvailableSignalResponse() *KeyfactorApiModelsWorkflowsAvailableSignalResponse {
	this := KeyfactorApiModelsWorkflowsAvailableSignalResponse{}
	return &this
}

// NewKeyfactorApiModelsWorkflowsAvailableSignalResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsAvailableSignalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsWorkflowsAvailableSignalResponseWithDefaults() *KeyfactorApiModelsWorkflowsAvailableSignalResponse {
	this := KeyfactorApiModelsWorkflowsAvailableSignalResponse{}
	return &this
}

// GetSignalName returns the SignalName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) GetSignalName() string {
	if o == nil || isNil(o.SignalName) {
		var ret string
		return ret
	}
	return *o.SignalName
}

// GetSignalNameOk returns a tuple with the SignalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) GetSignalNameOk() (*string, bool) {
	if o == nil || isNil(o.SignalName) {
		return nil, false
	}
	return o.SignalName, true
}

// HasSignalName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) HasSignalName() bool {
	if o != nil && !isNil(o.SignalName) {
		return true
	}

	return false
}

// SetSignalName gets a reference to the given string and assigns it to the SignalName field.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) SetSignalName(v string) {
	o.SignalName = &v
}

// GetStepSignalId returns the StepSignalId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) GetStepSignalId() string {
	if o == nil || isNil(o.StepSignalId) {
		var ret string
		return ret
	}
	return *o.StepSignalId
}

// GetStepSignalIdOk returns a tuple with the StepSignalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) GetStepSignalIdOk() (*string, bool) {
	if o == nil || isNil(o.StepSignalId) {
		return nil, false
	}
	return o.StepSignalId, true
}

// HasStepSignalId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) HasStepSignalId() bool {
	if o != nil && !isNil(o.StepSignalId) {
		return true
	}

	return false
}

// SetStepSignalId gets a reference to the given string and assigns it to the StepSignalId field.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) SetStepSignalId(v string) {
	o.StepSignalId = &v
}

// GetSignalReceived returns the SignalReceived field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) GetSignalReceived() bool {
	if o == nil || isNil(o.SignalReceived) {
		var ret bool
		return ret
	}
	return *o.SignalReceived
}

// GetSignalReceivedOk returns a tuple with the SignalReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) GetSignalReceivedOk() (*bool, bool) {
	if o == nil || isNil(o.SignalReceived) {
		return nil, false
	}
	return o.SignalReceived, true
}

// HasSignalReceived returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) HasSignalReceived() bool {
	if o != nil && !isNil(o.SignalReceived) {
		return true
	}

	return false
}

// SetSignalReceived gets a reference to the given bool and assigns it to the SignalReceived field.
func (o *KeyfactorApiModelsWorkflowsAvailableSignalResponse) SetSignalReceived(v bool) {
	o.SignalReceived = &v
}

func (o KeyfactorApiModelsWorkflowsAvailableSignalResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsWorkflowsAvailableSignalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SignalName) {
		toSerialize["SignalName"] = o.SignalName
	}
	if !isNil(o.StepSignalId) {
		toSerialize["StepSignalId"] = o.StepSignalId
	}
	if !isNil(o.SignalReceived) {
		toSerialize["SignalReceived"] = o.SignalReceived
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsWorkflowsAvailableSignalResponse struct {
	value *KeyfactorApiModelsWorkflowsAvailableSignalResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsWorkflowsAvailableSignalResponse) Get() *KeyfactorApiModelsWorkflowsAvailableSignalResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsWorkflowsAvailableSignalResponse) Set(val *KeyfactorApiModelsWorkflowsAvailableSignalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsWorkflowsAvailableSignalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsWorkflowsAvailableSignalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsWorkflowsAvailableSignalResponse(val *KeyfactorApiModelsWorkflowsAvailableSignalResponse) *NullableKeyfactorApiModelsWorkflowsAvailableSignalResponse {
	return &NullableKeyfactorApiModelsWorkflowsAvailableSignalResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsWorkflowsAvailableSignalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsWorkflowsAvailableSignalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


