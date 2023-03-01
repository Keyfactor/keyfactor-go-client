/*
Keyfactor-v1

This reference serves to document REST-based methods to manage and integrate with Keyfactor. In addition, an embedded interface allows for the execution of calls against the current Keyfactor API instance.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keyfactor_command_client_api

import (
	"encoding/json"
	"time"
)

// checks if the KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest{}

// KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest struct for KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest
type KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest struct {
	AlertId *int32 `json:"AlertId,omitempty"`
	EvaluationDate *time.Time `json:"EvaluationDate,omitempty"`
	SendAlerts *bool `json:"SendAlerts,omitempty"`
}

// NewKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest() *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest {
	this := KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest{}
	return &this
}

// NewKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequestWithDefaults instantiates a new KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequestWithDefaults() *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest {
	this := KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest{}
	return &this
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetAlertId() int32 {
	if o == nil || isNil(o.AlertId) {
		var ret int32
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetAlertIdOk() (*int32, bool) {
	if o == nil || isNil(o.AlertId) {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) HasAlertId() bool {
	if o != nil && !isNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given int32 and assigns it to the AlertId field.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) SetAlertId(v int32) {
	o.AlertId = &v
}

// GetEvaluationDate returns the EvaluationDate field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetEvaluationDate() time.Time {
	if o == nil || isNil(o.EvaluationDate) {
		var ret time.Time
		return ret
	}
	return *o.EvaluationDate
}

// GetEvaluationDateOk returns a tuple with the EvaluationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetEvaluationDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.EvaluationDate) {
		return nil, false
	}
	return o.EvaluationDate, true
}

// HasEvaluationDate returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) HasEvaluationDate() bool {
	if o != nil && !isNil(o.EvaluationDate) {
		return true
	}

	return false
}

// SetEvaluationDate gets a reference to the given time.Time and assigns it to the EvaluationDate field.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) SetEvaluationDate(v time.Time) {
	o.EvaluationDate = &v
}

// GetSendAlerts returns the SendAlerts field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetSendAlerts() bool {
	if o == nil || isNil(o.SendAlerts) {
		var ret bool
		return ret
	}
	return *o.SendAlerts
}

// GetSendAlertsOk returns a tuple with the SendAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) GetSendAlertsOk() (*bool, bool) {
	if o == nil || isNil(o.SendAlerts) {
		return nil, false
	}
	return o.SendAlerts, true
}

// HasSendAlerts returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) HasSendAlerts() bool {
	if o != nil && !isNil(o.SendAlerts) {
		return true
	}

	return false
}

// SetSendAlerts gets a reference to the given bool and assigns it to the SendAlerts field.
func (o *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) SetSendAlerts(v bool) {
	o.SendAlerts = &v
}

func (o KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AlertId) {
		toSerialize["AlertId"] = o.AlertId
	}
	if !isNil(o.EvaluationDate) {
		toSerialize["EvaluationDate"] = o.EvaluationDate
	}
	if !isNil(o.SendAlerts) {
		toSerialize["SendAlerts"] = o.SendAlerts
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest struct {
	value *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) Get() *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) Set(val *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest(val *KeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) *NullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest {
	return &NullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsMonitoringRevocationMonitoringAlertTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


