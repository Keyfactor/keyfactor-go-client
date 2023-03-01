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

// checks if the KeyfactorApiModelsSslUpdateNetworkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsSslUpdateNetworkRequest{}

// KeyfactorApiModelsSslUpdateNetworkRequest struct for KeyfactorApiModelsSslUpdateNetworkRequest
type KeyfactorApiModelsSslUpdateNetworkRequest struct {
	NetworkId string `json:"NetworkId"`
	Name string `json:"Name"`
	AgentPoolName string `json:"AgentPoolName"`
	Description string `json:"Description"`
	Enabled *bool `json:"Enabled,omitempty"`
	DiscoverSchedule *KeyfactorCommonSchedulingKeyfactorSchedule `json:"DiscoverSchedule,omitempty"`
	MonitorSchedule *KeyfactorCommonSchedulingKeyfactorSchedule `json:"MonitorSchedule,omitempty"`
	SslAlertRecipients []string `json:"SslAlertRecipients,omitempty"`
	AutoMonitor *bool `json:"AutoMonitor,omitempty"`
	GetRobots *bool `json:"GetRobots,omitempty"`
	DiscoverTimeoutMs *float64 `json:"DiscoverTimeoutMs,omitempty"`
	MonitorTimeoutMs *float64 `json:"MonitorTimeoutMs,omitempty"`
	ExpirationAlertDays *float64 `json:"ExpirationAlertDays,omitempty"`
	QuietHours []KeyfactorApiModelsSslQuietHourRequest `json:"QuietHours,omitempty"`
}

// NewKeyfactorApiModelsSslUpdateNetworkRequest instantiates a new KeyfactorApiModelsSslUpdateNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsSslUpdateNetworkRequest(networkId string, name string, agentPoolName string, description string) *KeyfactorApiModelsSslUpdateNetworkRequest {
	this := KeyfactorApiModelsSslUpdateNetworkRequest{}
	this.NetworkId = networkId
	this.Name = name
	this.AgentPoolName = agentPoolName
	this.Description = description
	return &this
}

// NewKeyfactorApiModelsSslUpdateNetworkRequestWithDefaults instantiates a new KeyfactorApiModelsSslUpdateNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsSslUpdateNetworkRequestWithDefaults() *KeyfactorApiModelsSslUpdateNetworkRequest {
	this := KeyfactorApiModelsSslUpdateNetworkRequest{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetName returns the Name field value
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetName(v string) {
	o.Name = v
}

// GetAgentPoolName returns the AgentPoolName field value
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetAgentPoolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentPoolName
}

// GetAgentPoolNameOk returns a tuple with the AgentPoolName field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetAgentPoolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentPoolName, true
}

// SetAgentPoolName sets field value
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetAgentPoolName(v string) {
	o.AgentPoolName = v
}

// GetDescription returns the Description field value
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetDescription(v string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDiscoverSchedule returns the DiscoverSchedule field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.DiscoverSchedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.DiscoverSchedule
}

// GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.DiscoverSchedule) {
		return nil, false
	}
	return o.DiscoverSchedule, true
}

// HasDiscoverSchedule returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasDiscoverSchedule() bool {
	if o != nil && !isNil(o.DiscoverSchedule) {
		return true
	}

	return false
}

// SetDiscoverSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the DiscoverSchedule field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.DiscoverSchedule = &v
}

// GetMonitorSchedule returns the MonitorSchedule field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.MonitorSchedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.MonitorSchedule
}

// GetMonitorScheduleOk returns a tuple with the MonitorSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.MonitorSchedule) {
		return nil, false
	}
	return o.MonitorSchedule, true
}

// HasMonitorSchedule returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasMonitorSchedule() bool {
	if o != nil && !isNil(o.MonitorSchedule) {
		return true
	}

	return false
}

// SetMonitorSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the MonitorSchedule field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.MonitorSchedule = &v
}

// GetSslAlertRecipients returns the SslAlertRecipients field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetSslAlertRecipients() []string {
	if o == nil || isNil(o.SslAlertRecipients) {
		var ret []string
		return ret
	}
	return o.SslAlertRecipients
}

// GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetSslAlertRecipientsOk() ([]string, bool) {
	if o == nil || isNil(o.SslAlertRecipients) {
		return nil, false
	}
	return o.SslAlertRecipients, true
}

// HasSslAlertRecipients returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasSslAlertRecipients() bool {
	if o != nil && !isNil(o.SslAlertRecipients) {
		return true
	}

	return false
}

// SetSslAlertRecipients gets a reference to the given []string and assigns it to the SslAlertRecipients field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetSslAlertRecipients(v []string) {
	o.SslAlertRecipients = v
}

// GetAutoMonitor returns the AutoMonitor field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetAutoMonitor() bool {
	if o == nil || isNil(o.AutoMonitor) {
		var ret bool
		return ret
	}
	return *o.AutoMonitor
}

// GetAutoMonitorOk returns a tuple with the AutoMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetAutoMonitorOk() (*bool, bool) {
	if o == nil || isNil(o.AutoMonitor) {
		return nil, false
	}
	return o.AutoMonitor, true
}

// HasAutoMonitor returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasAutoMonitor() bool {
	if o != nil && !isNil(o.AutoMonitor) {
		return true
	}

	return false
}

// SetAutoMonitor gets a reference to the given bool and assigns it to the AutoMonitor field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetAutoMonitor(v bool) {
	o.AutoMonitor = &v
}

// GetGetRobots returns the GetRobots field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetGetRobots() bool {
	if o == nil || isNil(o.GetRobots) {
		var ret bool
		return ret
	}
	return *o.GetRobots
}

// GetGetRobotsOk returns a tuple with the GetRobots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetGetRobotsOk() (*bool, bool) {
	if o == nil || isNil(o.GetRobots) {
		return nil, false
	}
	return o.GetRobots, true
}

// HasGetRobots returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasGetRobots() bool {
	if o != nil && !isNil(o.GetRobots) {
		return true
	}

	return false
}

// SetGetRobots gets a reference to the given bool and assigns it to the GetRobots field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetGetRobots(v bool) {
	o.GetRobots = &v
}

// GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverTimeoutMs() float64 {
	if o == nil || isNil(o.DiscoverTimeoutMs) {
		var ret float64
		return ret
	}
	return *o.DiscoverTimeoutMs
}

// GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetDiscoverTimeoutMsOk() (*float64, bool) {
	if o == nil || isNil(o.DiscoverTimeoutMs) {
		return nil, false
	}
	return o.DiscoverTimeoutMs, true
}

// HasDiscoverTimeoutMs returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasDiscoverTimeoutMs() bool {
	if o != nil && !isNil(o.DiscoverTimeoutMs) {
		return true
	}

	return false
}

// SetDiscoverTimeoutMs gets a reference to the given float64 and assigns it to the DiscoverTimeoutMs field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetDiscoverTimeoutMs(v float64) {
	o.DiscoverTimeoutMs = &v
}

// GetMonitorTimeoutMs returns the MonitorTimeoutMs field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorTimeoutMs() float64 {
	if o == nil || isNil(o.MonitorTimeoutMs) {
		var ret float64
		return ret
	}
	return *o.MonitorTimeoutMs
}

// GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetMonitorTimeoutMsOk() (*float64, bool) {
	if o == nil || isNil(o.MonitorTimeoutMs) {
		return nil, false
	}
	return o.MonitorTimeoutMs, true
}

// HasMonitorTimeoutMs returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasMonitorTimeoutMs() bool {
	if o != nil && !isNil(o.MonitorTimeoutMs) {
		return true
	}

	return false
}

// SetMonitorTimeoutMs gets a reference to the given float64 and assigns it to the MonitorTimeoutMs field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetMonitorTimeoutMs(v float64) {
	o.MonitorTimeoutMs = &v
}

// GetExpirationAlertDays returns the ExpirationAlertDays field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetExpirationAlertDays() float64 {
	if o == nil || isNil(o.ExpirationAlertDays) {
		var ret float64
		return ret
	}
	return *o.ExpirationAlertDays
}

// GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetExpirationAlertDaysOk() (*float64, bool) {
	if o == nil || isNil(o.ExpirationAlertDays) {
		return nil, false
	}
	return o.ExpirationAlertDays, true
}

// HasExpirationAlertDays returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasExpirationAlertDays() bool {
	if o != nil && !isNil(o.ExpirationAlertDays) {
		return true
	}

	return false
}

// SetExpirationAlertDays gets a reference to the given float64 and assigns it to the ExpirationAlertDays field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetExpirationAlertDays(v float64) {
	o.ExpirationAlertDays = &v
}

// GetQuietHours returns the QuietHours field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetQuietHours() []KeyfactorApiModelsSslQuietHourRequest {
	if o == nil || isNil(o.QuietHours) {
		var ret []KeyfactorApiModelsSslQuietHourRequest
		return ret
	}
	return o.QuietHours
}

// GetQuietHoursOk returns a tuple with the QuietHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) GetQuietHoursOk() ([]KeyfactorApiModelsSslQuietHourRequest, bool) {
	if o == nil || isNil(o.QuietHours) {
		return nil, false
	}
	return o.QuietHours, true
}

// HasQuietHours returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) HasQuietHours() bool {
	if o != nil && !isNil(o.QuietHours) {
		return true
	}

	return false
}

// SetQuietHours gets a reference to the given []KeyfactorApiModelsSslQuietHourRequest and assigns it to the QuietHours field.
func (o *KeyfactorApiModelsSslUpdateNetworkRequest) SetQuietHours(v []KeyfactorApiModelsSslQuietHourRequest) {
	o.QuietHours = v
}

func (o KeyfactorApiModelsSslUpdateNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsSslUpdateNetworkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["NetworkId"] = o.NetworkId
	toSerialize["Name"] = o.Name
	toSerialize["AgentPoolName"] = o.AgentPoolName
	toSerialize["Description"] = o.Description
	if !isNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !isNil(o.DiscoverSchedule) {
		toSerialize["DiscoverSchedule"] = o.DiscoverSchedule
	}
	if !isNil(o.MonitorSchedule) {
		toSerialize["MonitorSchedule"] = o.MonitorSchedule
	}
	if !isNil(o.SslAlertRecipients) {
		toSerialize["SslAlertRecipients"] = o.SslAlertRecipients
	}
	if !isNil(o.AutoMonitor) {
		toSerialize["AutoMonitor"] = o.AutoMonitor
	}
	if !isNil(o.GetRobots) {
		toSerialize["GetRobots"] = o.GetRobots
	}
	if !isNil(o.DiscoverTimeoutMs) {
		toSerialize["DiscoverTimeoutMs"] = o.DiscoverTimeoutMs
	}
	if !isNil(o.MonitorTimeoutMs) {
		toSerialize["MonitorTimeoutMs"] = o.MonitorTimeoutMs
	}
	if !isNil(o.ExpirationAlertDays) {
		toSerialize["ExpirationAlertDays"] = o.ExpirationAlertDays
	}
	if !isNil(o.QuietHours) {
		toSerialize["QuietHours"] = o.QuietHours
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsSslUpdateNetworkRequest struct {
	value *KeyfactorApiModelsSslUpdateNetworkRequest
	isSet bool
}

func (v NullableKeyfactorApiModelsSslUpdateNetworkRequest) Get() *KeyfactorApiModelsSslUpdateNetworkRequest {
	return v.value
}

func (v *NullableKeyfactorApiModelsSslUpdateNetworkRequest) Set(val *KeyfactorApiModelsSslUpdateNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsSslUpdateNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsSslUpdateNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsSslUpdateNetworkRequest(val *KeyfactorApiModelsSslUpdateNetworkRequest) *NullableKeyfactorApiModelsSslUpdateNetworkRequest {
	return &NullableKeyfactorApiModelsSslUpdateNetworkRequest{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsSslUpdateNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsSslUpdateNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


