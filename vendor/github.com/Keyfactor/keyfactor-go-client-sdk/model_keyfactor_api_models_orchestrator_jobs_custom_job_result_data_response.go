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

// checks if the KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse{}

// KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse struct for KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse
type KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse struct {
	JobHistoryId *int64 `json:"JobHistoryId,omitempty"`
	Data *string `json:"Data,omitempty"`
}

// NewKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse instantiates a new KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse() *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse {
	this := KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse{}
	return &this
}

// NewKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponseWithDefaults instantiates a new KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponseWithDefaults() *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse {
	this := KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse{}
	return &this
}

// GetJobHistoryId returns the JobHistoryId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) GetJobHistoryId() int64 {
	if o == nil || isNil(o.JobHistoryId) {
		var ret int64
		return ret
	}
	return *o.JobHistoryId
}

// GetJobHistoryIdOk returns a tuple with the JobHistoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) GetJobHistoryIdOk() (*int64, bool) {
	if o == nil || isNil(o.JobHistoryId) {
		return nil, false
	}
	return o.JobHistoryId, true
}

// HasJobHistoryId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) HasJobHistoryId() bool {
	if o != nil && !isNil(o.JobHistoryId) {
		return true
	}

	return false
}

// SetJobHistoryId gets a reference to the given int64 and assigns it to the JobHistoryId field.
func (o *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) SetJobHistoryId(v int64) {
	o.JobHistoryId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) GetData() string {
	if o == nil || isNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) GetDataOk() (*string, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) SetData(v string) {
	o.Data = &v
}

func (o KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JobHistoryId) {
		toSerialize["JobHistoryId"] = o.JobHistoryId
	}
	if !isNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse struct {
	value *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) Get() *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) Set(val *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse(val *KeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) *NullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse {
	return &NullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsCustomJobResultDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


