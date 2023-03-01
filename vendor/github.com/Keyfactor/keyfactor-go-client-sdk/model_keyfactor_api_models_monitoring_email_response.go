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

// checks if the KeyfactorApiModelsMonitoringEmailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsMonitoringEmailResponse{}

// KeyfactorApiModelsMonitoringEmailResponse struct for KeyfactorApiModelsMonitoringEmailResponse
type KeyfactorApiModelsMonitoringEmailResponse struct {
	EnableReminder *bool `json:"EnableReminder,omitempty"`
	WarningDays *int32 `json:"WarningDays,omitempty"`
	Recipients []string `json:"Recipients,omitempty"`
}

// NewKeyfactorApiModelsMonitoringEmailResponse instantiates a new KeyfactorApiModelsMonitoringEmailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsMonitoringEmailResponse() *KeyfactorApiModelsMonitoringEmailResponse {
	this := KeyfactorApiModelsMonitoringEmailResponse{}
	return &this
}

// NewKeyfactorApiModelsMonitoringEmailResponseWithDefaults instantiates a new KeyfactorApiModelsMonitoringEmailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsMonitoringEmailResponseWithDefaults() *KeyfactorApiModelsMonitoringEmailResponse {
	this := KeyfactorApiModelsMonitoringEmailResponse{}
	return &this
}

// GetEnableReminder returns the EnableReminder field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringEmailResponse) GetEnableReminder() bool {
	if o == nil || isNil(o.EnableReminder) {
		var ret bool
		return ret
	}
	return *o.EnableReminder
}

// GetEnableReminderOk returns a tuple with the EnableReminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringEmailResponse) GetEnableReminderOk() (*bool, bool) {
	if o == nil || isNil(o.EnableReminder) {
		return nil, false
	}
	return o.EnableReminder, true
}

// HasEnableReminder returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringEmailResponse) HasEnableReminder() bool {
	if o != nil && !isNil(o.EnableReminder) {
		return true
	}

	return false
}

// SetEnableReminder gets a reference to the given bool and assigns it to the EnableReminder field.
func (o *KeyfactorApiModelsMonitoringEmailResponse) SetEnableReminder(v bool) {
	o.EnableReminder = &v
}

// GetWarningDays returns the WarningDays field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringEmailResponse) GetWarningDays() int32 {
	if o == nil || isNil(o.WarningDays) {
		var ret int32
		return ret
	}
	return *o.WarningDays
}

// GetWarningDaysOk returns a tuple with the WarningDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringEmailResponse) GetWarningDaysOk() (*int32, bool) {
	if o == nil || isNil(o.WarningDays) {
		return nil, false
	}
	return o.WarningDays, true
}

// HasWarningDays returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringEmailResponse) HasWarningDays() bool {
	if o != nil && !isNil(o.WarningDays) {
		return true
	}

	return false
}

// SetWarningDays gets a reference to the given int32 and assigns it to the WarningDays field.
func (o *KeyfactorApiModelsMonitoringEmailResponse) SetWarningDays(v int32) {
	o.WarningDays = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *KeyfactorApiModelsMonitoringEmailResponse) GetRecipients() []string {
	if o == nil || isNil(o.Recipients) {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsMonitoringEmailResponse) GetRecipientsOk() ([]string, bool) {
	if o == nil || isNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *KeyfactorApiModelsMonitoringEmailResponse) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *KeyfactorApiModelsMonitoringEmailResponse) SetRecipients(v []string) {
	o.Recipients = v
}

func (o KeyfactorApiModelsMonitoringEmailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsMonitoringEmailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EnableReminder) {
		toSerialize["EnableReminder"] = o.EnableReminder
	}
	if !isNil(o.WarningDays) {
		toSerialize["WarningDays"] = o.WarningDays
	}
	if !isNil(o.Recipients) {
		toSerialize["Recipients"] = o.Recipients
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsMonitoringEmailResponse struct {
	value *KeyfactorApiModelsMonitoringEmailResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsMonitoringEmailResponse) Get() *KeyfactorApiModelsMonitoringEmailResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsMonitoringEmailResponse) Set(val *KeyfactorApiModelsMonitoringEmailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsMonitoringEmailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsMonitoringEmailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsMonitoringEmailResponse(val *KeyfactorApiModelsMonitoringEmailResponse) *NullableKeyfactorApiModelsMonitoringEmailResponse {
	return &NullableKeyfactorApiModelsMonitoringEmailResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsMonitoringEmailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsMonitoringEmailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


