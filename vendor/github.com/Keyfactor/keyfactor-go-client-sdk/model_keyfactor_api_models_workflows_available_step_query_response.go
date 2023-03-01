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

// checks if the KeyfactorApiModelsWorkflowsAvailableStepQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsWorkflowsAvailableStepQueryResponse{}

// KeyfactorApiModelsWorkflowsAvailableStepQueryResponse struct for KeyfactorApiModelsWorkflowsAvailableStepQueryResponse
type KeyfactorApiModelsWorkflowsAvailableStepQueryResponse struct {
	// The display name of the step.
	DisplayName *string `json:"DisplayName,omitempty"`
	// The extension name of the step.
	ExtensionName *string `json:"ExtensionName,omitempty"`
	// The workflow types which this step can be a part of.
	SupportedWorkflowTypes []string `json:"SupportedWorkflowTypes,omitempty"`
	ConfigurationParametersDefinition *map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse `json:"ConfigurationParametersDefinition,omitempty"`
	SignalsDefinition *map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse `json:"SignalsDefinition,omitempty"`
}

// NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponse instantiates a new KeyfactorApiModelsWorkflowsAvailableStepQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponse() *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse {
	this := KeyfactorApiModelsWorkflowsAvailableStepQueryResponse{}
	return &this
}

// NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsAvailableStepQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsWorkflowsAvailableStepQueryResponseWithDefaults() *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse {
	this := KeyfactorApiModelsWorkflowsAvailableStepQueryResponse{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExtensionName returns the ExtensionName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetExtensionName() string {
	if o == nil || isNil(o.ExtensionName) {
		var ret string
		return ret
	}
	return *o.ExtensionName
}

// GetExtensionNameOk returns a tuple with the ExtensionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetExtensionNameOk() (*string, bool) {
	if o == nil || isNil(o.ExtensionName) {
		return nil, false
	}
	return o.ExtensionName, true
}

// HasExtensionName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasExtensionName() bool {
	if o != nil && !isNil(o.ExtensionName) {
		return true
	}

	return false
}

// SetExtensionName gets a reference to the given string and assigns it to the ExtensionName field.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetExtensionName(v string) {
	o.ExtensionName = &v
}

// GetSupportedWorkflowTypes returns the SupportedWorkflowTypes field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSupportedWorkflowTypes() []string {
	if o == nil || isNil(o.SupportedWorkflowTypes) {
		var ret []string
		return ret
	}
	return o.SupportedWorkflowTypes
}

// GetSupportedWorkflowTypesOk returns a tuple with the SupportedWorkflowTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSupportedWorkflowTypesOk() ([]string, bool) {
	if o == nil || isNil(o.SupportedWorkflowTypes) {
		return nil, false
	}
	return o.SupportedWorkflowTypes, true
}

// HasSupportedWorkflowTypes returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasSupportedWorkflowTypes() bool {
	if o != nil && !isNil(o.SupportedWorkflowTypes) {
		return true
	}

	return false
}

// SetSupportedWorkflowTypes gets a reference to the given []string and assigns it to the SupportedWorkflowTypes field.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetSupportedWorkflowTypes(v []string) {
	o.SupportedWorkflowTypes = v
}

// GetConfigurationParametersDefinition returns the ConfigurationParametersDefinition field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetConfigurationParametersDefinition() map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse {
	if o == nil || isNil(o.ConfigurationParametersDefinition) {
		var ret map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse
		return ret
	}
	return *o.ConfigurationParametersDefinition
}

// GetConfigurationParametersDefinitionOk returns a tuple with the ConfigurationParametersDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetConfigurationParametersDefinitionOk() (*map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse, bool) {
	if o == nil || isNil(o.ConfigurationParametersDefinition) {
		return nil, false
	}
	return o.ConfigurationParametersDefinition, true
}

// HasConfigurationParametersDefinition returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasConfigurationParametersDefinition() bool {
	if o != nil && !isNil(o.ConfigurationParametersDefinition) {
		return true
	}

	return false
}

// SetConfigurationParametersDefinition gets a reference to the given map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse and assigns it to the ConfigurationParametersDefinition field.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetConfigurationParametersDefinition(v map[string]KeyfactorApiModelsWorkflowsParameterDefinitionResponse) {
	o.ConfigurationParametersDefinition = &v
}

// GetSignalsDefinition returns the SignalsDefinition field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSignalsDefinition() map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse {
	if o == nil || isNil(o.SignalsDefinition) {
		var ret map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse
		return ret
	}
	return *o.SignalsDefinition
}

// GetSignalsDefinitionOk returns a tuple with the SignalsDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) GetSignalsDefinitionOk() (*map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse, bool) {
	if o == nil || isNil(o.SignalsDefinition) {
		return nil, false
	}
	return o.SignalsDefinition, true
}

// HasSignalsDefinition returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) HasSignalsDefinition() bool {
	if o != nil && !isNil(o.SignalsDefinition) {
		return true
	}

	return false
}

// SetSignalsDefinition gets a reference to the given map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse and assigns it to the SignalsDefinition field.
func (o *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) SetSignalsDefinition(v map[string]KeyfactorApiModelsWorkflowsSignalDefinitionResponse) {
	o.SignalsDefinition = &v
}

func (o KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !isNil(o.ExtensionName) {
		toSerialize["ExtensionName"] = o.ExtensionName
	}
	if !isNil(o.SupportedWorkflowTypes) {
		toSerialize["SupportedWorkflowTypes"] = o.SupportedWorkflowTypes
	}
	if !isNil(o.ConfigurationParametersDefinition) {
		toSerialize["ConfigurationParametersDefinition"] = o.ConfigurationParametersDefinition
	}
	if !isNil(o.SignalsDefinition) {
		toSerialize["SignalsDefinition"] = o.SignalsDefinition
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse struct {
	value *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) Get() *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) Set(val *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse(val *KeyfactorApiModelsWorkflowsAvailableStepQueryResponse) *NullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse {
	return &NullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsWorkflowsAvailableStepQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


