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

// checks if the KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse{}

// KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse struct for KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse
type KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse struct {
	WorkflowType *string `json:"WorkflowType,omitempty"`
	KeyType *string `json:"KeyType,omitempty"`
	ContextParameters []string `json:"ContextParameters,omitempty"`
	BuiltInSteps []KeyfactorApiModelsWorkflowsAvailableStepResponse `json:"BuiltInSteps,omitempty"`
}

// NewKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse instantiates a new KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse() *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse {
	this := KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse{}
	return &this
}

// NewKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponseWithDefaults() *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse {
	this := KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse{}
	return &this
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetWorkflowType() string {
	if o == nil || isNil(o.WorkflowType) {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || isNil(o.WorkflowType) {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) HasWorkflowType() bool {
	if o != nil && !isNil(o.WorkflowType) {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetKeyType() string {
	if o == nil || isNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetKeyTypeOk() (*string, bool) {
	if o == nil || isNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) HasKeyType() bool {
	if o != nil && !isNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetKeyType(v string) {
	o.KeyType = &v
}

// GetContextParameters returns the ContextParameters field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetContextParameters() []string {
	if o == nil || isNil(o.ContextParameters) {
		var ret []string
		return ret
	}
	return o.ContextParameters
}

// GetContextParametersOk returns a tuple with the ContextParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetContextParametersOk() ([]string, bool) {
	if o == nil || isNil(o.ContextParameters) {
		return nil, false
	}
	return o.ContextParameters, true
}

// HasContextParameters returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) HasContextParameters() bool {
	if o != nil && !isNil(o.ContextParameters) {
		return true
	}

	return false
}

// SetContextParameters gets a reference to the given []string and assigns it to the ContextParameters field.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetContextParameters(v []string) {
	o.ContextParameters = v
}

// GetBuiltInSteps returns the BuiltInSteps field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetBuiltInSteps() []KeyfactorApiModelsWorkflowsAvailableStepResponse {
	if o == nil || isNil(o.BuiltInSteps) {
		var ret []KeyfactorApiModelsWorkflowsAvailableStepResponse
		return ret
	}
	return o.BuiltInSteps
}

// GetBuiltInStepsOk returns a tuple with the BuiltInSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) GetBuiltInStepsOk() ([]KeyfactorApiModelsWorkflowsAvailableStepResponse, bool) {
	if o == nil || isNil(o.BuiltInSteps) {
		return nil, false
	}
	return o.BuiltInSteps, true
}

// HasBuiltInSteps returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) HasBuiltInSteps() bool {
	if o != nil && !isNil(o.BuiltInSteps) {
		return true
	}

	return false
}

// SetBuiltInSteps gets a reference to the given []KeyfactorApiModelsWorkflowsAvailableStepResponse and assigns it to the BuiltInSteps field.
func (o *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) SetBuiltInSteps(v []KeyfactorApiModelsWorkflowsAvailableStepResponse) {
	o.BuiltInSteps = v
}

func (o KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WorkflowType) {
		toSerialize["WorkflowType"] = o.WorkflowType
	}
	if !isNil(o.KeyType) {
		toSerialize["KeyType"] = o.KeyType
	}
	if !isNil(o.ContextParameters) {
		toSerialize["ContextParameters"] = o.ContextParameters
	}
	if !isNil(o.BuiltInSteps) {
		toSerialize["BuiltInSteps"] = o.BuiltInSteps
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse struct {
	value *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) Get() *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) Set(val *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse(val *KeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) *NullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse {
	return &NullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsWorkflowsWorkflowTypeQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


