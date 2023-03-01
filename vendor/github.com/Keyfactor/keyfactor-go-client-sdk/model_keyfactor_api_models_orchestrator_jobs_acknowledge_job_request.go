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

// checks if the KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest{}

// KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest struct for KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest
type KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest struct {
	// List of orchestrator job audit ids to be acknowledged
	JobAuditIds []int64 `json:"JobAuditIds,omitempty"`
	// Query identifying orchestrator jobs to be acknowledged
	Query *string `json:"Query,omitempty"`
}

// NewKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest instantiates a new KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest() *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest {
	this := KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest{}
	return &this
}

// NewKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequestWithDefaults instantiates a new KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequestWithDefaults() *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest {
	this := KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest{}
	return &this
}

// GetJobAuditIds returns the JobAuditIds field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) GetJobAuditIds() []int64 {
	if o == nil || isNil(o.JobAuditIds) {
		var ret []int64
		return ret
	}
	return o.JobAuditIds
}

// GetJobAuditIdsOk returns a tuple with the JobAuditIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) GetJobAuditIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.JobAuditIds) {
		return nil, false
	}
	return o.JobAuditIds, true
}

// HasJobAuditIds returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) HasJobAuditIds() bool {
	if o != nil && !isNil(o.JobAuditIds) {
		return true
	}

	return false
}

// SetJobAuditIds gets a reference to the given []int64 and assigns it to the JobAuditIds field.
func (o *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) SetJobAuditIds(v []int64) {
	o.JobAuditIds = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) SetQuery(v string) {
	o.Query = &v
}

func (o KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JobAuditIds) {
		toSerialize["JobAuditIds"] = o.JobAuditIds
	}
	if !isNil(o.Query) {
		toSerialize["Query"] = o.Query
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest struct {
	value *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) Get() *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) Set(val *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest(val *KeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) *NullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest {
	return &NullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsOrchestratorJobsAcknowledgeJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


