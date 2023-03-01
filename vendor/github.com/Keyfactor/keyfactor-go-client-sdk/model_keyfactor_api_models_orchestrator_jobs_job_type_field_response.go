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

// checks if the KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse{}

// KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse struct for KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse
type KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse struct {
	Name *string `json:"Name,omitempty"`
	Type *int32 `json:"Type,omitempty"`
	DefaultValue *string `json:"DefaultValue,omitempty"`
	Required *bool `json:"Required,omitempty"`
}

// NewKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse instantiates a new KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse() *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse {
	this := KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse{}
	return &this
}

// NewKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponseWithDefaults instantiates a new KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponseWithDefaults() *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse {
	this := KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) GetType() int32 {
	if o == nil || isNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) GetTypeOk() (*int32, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) SetType(v int32) {
	o.Type = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) GetDefaultValue() string {
	if o == nil || isNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) GetDefaultValueOk() (*string, bool) {
	if o == nil || isNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) HasDefaultValue() bool {
	if o != nil && !isNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) GetRequired() bool {
	if o == nil || isNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) GetRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) HasRequired() bool {
	if o != nil && !isNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) SetRequired(v bool) {
	o.Required = &v
}

func (o KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !isNil(o.DefaultValue) {
		toSerialize["DefaultValue"] = o.DefaultValue
	}
	if !isNil(o.Required) {
		toSerialize["Required"] = o.Required
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse struct {
	value *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) Get() *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) Set(val *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse(val *KeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) *NullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse {
	return &NullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsJobTypeFieldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


