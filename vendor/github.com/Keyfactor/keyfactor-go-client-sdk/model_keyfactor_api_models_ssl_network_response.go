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

// checks if the KeyfactorApiModelsSslNetworkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsSslNetworkResponse{}

// KeyfactorApiModelsSslNetworkResponse struct for KeyfactorApiModelsSslNetworkResponse
type KeyfactorApiModelsSslNetworkResponse struct {
	NetworkId *string `json:"NetworkId,omitempty"`
	Name *string `json:"Name,omitempty"`
	AgentPoolName *string `json:"AgentPoolName,omitempty"`
	AgentPoolId *string `json:"AgentPoolId,omitempty"`
	Description *string `json:"Description,omitempty"`
	Enabled *bool `json:"Enabled,omitempty"`
	DiscoverSchedule *KeyfactorCommonSchedulingKeyfactorSchedule `json:"DiscoverSchedule,omitempty"`
	MonitorSchedule *KeyfactorCommonSchedulingKeyfactorSchedule `json:"MonitorSchedule,omitempty"`
	DiscoverPercentComplete *float64 `json:"DiscoverPercentComplete,omitempty"`
	MonitorPercentComplete *float64 `json:"MonitorPercentComplete,omitempty"`
	DiscoverStatus *int32 `json:"DiscoverStatus,omitempty"`
	MonitorStatus *int32 `json:"MonitorStatus,omitempty"`
	DiscoverLastScanned *time.Time `json:"DiscoverLastScanned,omitempty"`
	MonitorLastScanned *time.Time `json:"MonitorLastScanned,omitempty"`
	SslAlertRecipients []string `json:"SslAlertRecipients,omitempty"`
	AutoMonitor *bool `json:"AutoMonitor,omitempty"`
	GetRobots *bool `json:"GetRobots,omitempty"`
	DiscoverTimeoutMs *float64 `json:"DiscoverTimeoutMs,omitempty"`
	MonitorTimeoutMs *float64 `json:"MonitorTimeoutMs,omitempty"`
	ExpirationAlertDays *float64 `json:"ExpirationAlertDays,omitempty"`
	DiscoverJobParts *int32 `json:"DiscoverJobParts,omitempty"`
	MonitorJobParts *int32 `json:"MonitorJobParts,omitempty"`
	QuietHours []KeyfactorApiModelsSslQuietHourResponse `json:"QuietHours,omitempty"`
}

// NewKeyfactorApiModelsSslNetworkResponse instantiates a new KeyfactorApiModelsSslNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsSslNetworkResponse() *KeyfactorApiModelsSslNetworkResponse {
	this := KeyfactorApiModelsSslNetworkResponse{}
	return &this
}

// NewKeyfactorApiModelsSslNetworkResponseWithDefaults instantiates a new KeyfactorApiModelsSslNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsSslNetworkResponseWithDefaults() *KeyfactorApiModelsSslNetworkResponse {
	this := KeyfactorApiModelsSslNetworkResponse{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetName(v string) {
	o.Name = &v
}

// GetAgentPoolName returns the AgentPoolName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetAgentPoolName() string {
	if o == nil || isNil(o.AgentPoolName) {
		var ret string
		return ret
	}
	return *o.AgentPoolName
}

// GetAgentPoolNameOk returns a tuple with the AgentPoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetAgentPoolNameOk() (*string, bool) {
	if o == nil || isNil(o.AgentPoolName) {
		return nil, false
	}
	return o.AgentPoolName, true
}

// HasAgentPoolName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasAgentPoolName() bool {
	if o != nil && !isNil(o.AgentPoolName) {
		return true
	}

	return false
}

// SetAgentPoolName gets a reference to the given string and assigns it to the AgentPoolName field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetAgentPoolName(v string) {
	o.AgentPoolName = &v
}

// GetAgentPoolId returns the AgentPoolId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetAgentPoolId() string {
	if o == nil || isNil(o.AgentPoolId) {
		var ret string
		return ret
	}
	return *o.AgentPoolId
}

// GetAgentPoolIdOk returns a tuple with the AgentPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetAgentPoolIdOk() (*string, bool) {
	if o == nil || isNil(o.AgentPoolId) {
		return nil, false
	}
	return o.AgentPoolId, true
}

// HasAgentPoolId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasAgentPoolId() bool {
	if o != nil && !isNil(o.AgentPoolId) {
		return true
	}

	return false
}

// SetAgentPoolId gets a reference to the given string and assigns it to the AgentPoolId field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetAgentPoolId(v string) {
	o.AgentPoolId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDiscoverSchedule returns the DiscoverSchedule field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.DiscoverSchedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.DiscoverSchedule
}

// GetDiscoverScheduleOk returns a tuple with the DiscoverSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.DiscoverSchedule) {
		return nil, false
	}
	return o.DiscoverSchedule, true
}

// HasDiscoverSchedule returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverSchedule() bool {
	if o != nil && !isNil(o.DiscoverSchedule) {
		return true
	}

	return false
}

// SetDiscoverSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the DiscoverSchedule field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.DiscoverSchedule = &v
}

// GetMonitorSchedule returns the MonitorSchedule field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.MonitorSchedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.MonitorSchedule
}

// GetMonitorScheduleOk returns a tuple with the MonitorSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.MonitorSchedule) {
		return nil, false
	}
	return o.MonitorSchedule, true
}

// HasMonitorSchedule returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorSchedule() bool {
	if o != nil && !isNil(o.MonitorSchedule) {
		return true
	}

	return false
}

// SetMonitorSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the MonitorSchedule field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.MonitorSchedule = &v
}

// GetDiscoverPercentComplete returns the DiscoverPercentComplete field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverPercentComplete() float64 {
	if o == nil || isNil(o.DiscoverPercentComplete) {
		var ret float64
		return ret
	}
	return *o.DiscoverPercentComplete
}

// GetDiscoverPercentCompleteOk returns a tuple with the DiscoverPercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverPercentCompleteOk() (*float64, bool) {
	if o == nil || isNil(o.DiscoverPercentComplete) {
		return nil, false
	}
	return o.DiscoverPercentComplete, true
}

// HasDiscoverPercentComplete returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverPercentComplete() bool {
	if o != nil && !isNil(o.DiscoverPercentComplete) {
		return true
	}

	return false
}

// SetDiscoverPercentComplete gets a reference to the given float64 and assigns it to the DiscoverPercentComplete field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverPercentComplete(v float64) {
	o.DiscoverPercentComplete = &v
}

// GetMonitorPercentComplete returns the MonitorPercentComplete field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorPercentComplete() float64 {
	if o == nil || isNil(o.MonitorPercentComplete) {
		var ret float64
		return ret
	}
	return *o.MonitorPercentComplete
}

// GetMonitorPercentCompleteOk returns a tuple with the MonitorPercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorPercentCompleteOk() (*float64, bool) {
	if o == nil || isNil(o.MonitorPercentComplete) {
		return nil, false
	}
	return o.MonitorPercentComplete, true
}

// HasMonitorPercentComplete returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorPercentComplete() bool {
	if o != nil && !isNil(o.MonitorPercentComplete) {
		return true
	}

	return false
}

// SetMonitorPercentComplete gets a reference to the given float64 and assigns it to the MonitorPercentComplete field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorPercentComplete(v float64) {
	o.MonitorPercentComplete = &v
}

// GetDiscoverStatus returns the DiscoverStatus field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverStatus() int32 {
	if o == nil || isNil(o.DiscoverStatus) {
		var ret int32
		return ret
	}
	return *o.DiscoverStatus
}

// GetDiscoverStatusOk returns a tuple with the DiscoverStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverStatusOk() (*int32, bool) {
	if o == nil || isNil(o.DiscoverStatus) {
		return nil, false
	}
	return o.DiscoverStatus, true
}

// HasDiscoverStatus returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverStatus() bool {
	if o != nil && !isNil(o.DiscoverStatus) {
		return true
	}

	return false
}

// SetDiscoverStatus gets a reference to the given int32 and assigns it to the DiscoverStatus field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverStatus(v int32) {
	o.DiscoverStatus = &v
}

// GetMonitorStatus returns the MonitorStatus field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorStatus() int32 {
	if o == nil || isNil(o.MonitorStatus) {
		var ret int32
		return ret
	}
	return *o.MonitorStatus
}

// GetMonitorStatusOk returns a tuple with the MonitorStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorStatusOk() (*int32, bool) {
	if o == nil || isNil(o.MonitorStatus) {
		return nil, false
	}
	return o.MonitorStatus, true
}

// HasMonitorStatus returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorStatus() bool {
	if o != nil && !isNil(o.MonitorStatus) {
		return true
	}

	return false
}

// SetMonitorStatus gets a reference to the given int32 and assigns it to the MonitorStatus field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorStatus(v int32) {
	o.MonitorStatus = &v
}

// GetDiscoverLastScanned returns the DiscoverLastScanned field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverLastScanned() time.Time {
	if o == nil || isNil(o.DiscoverLastScanned) {
		var ret time.Time
		return ret
	}
	return *o.DiscoverLastScanned
}

// GetDiscoverLastScannedOk returns a tuple with the DiscoverLastScanned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverLastScannedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DiscoverLastScanned) {
		return nil, false
	}
	return o.DiscoverLastScanned, true
}

// HasDiscoverLastScanned returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverLastScanned() bool {
	if o != nil && !isNil(o.DiscoverLastScanned) {
		return true
	}

	return false
}

// SetDiscoverLastScanned gets a reference to the given time.Time and assigns it to the DiscoverLastScanned field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverLastScanned(v time.Time) {
	o.DiscoverLastScanned = &v
}

// GetMonitorLastScanned returns the MonitorLastScanned field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorLastScanned() time.Time {
	if o == nil || isNil(o.MonitorLastScanned) {
		var ret time.Time
		return ret
	}
	return *o.MonitorLastScanned
}

// GetMonitorLastScannedOk returns a tuple with the MonitorLastScanned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorLastScannedOk() (*time.Time, bool) {
	if o == nil || isNil(o.MonitorLastScanned) {
		return nil, false
	}
	return o.MonitorLastScanned, true
}

// HasMonitorLastScanned returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorLastScanned() bool {
	if o != nil && !isNil(o.MonitorLastScanned) {
		return true
	}

	return false
}

// SetMonitorLastScanned gets a reference to the given time.Time and assigns it to the MonitorLastScanned field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorLastScanned(v time.Time) {
	o.MonitorLastScanned = &v
}

// GetSslAlertRecipients returns the SslAlertRecipients field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetSslAlertRecipients() []string {
	if o == nil || isNil(o.SslAlertRecipients) {
		var ret []string
		return ret
	}
	return o.SslAlertRecipients
}

// GetSslAlertRecipientsOk returns a tuple with the SslAlertRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetSslAlertRecipientsOk() ([]string, bool) {
	if o == nil || isNil(o.SslAlertRecipients) {
		return nil, false
	}
	return o.SslAlertRecipients, true
}

// HasSslAlertRecipients returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasSslAlertRecipients() bool {
	if o != nil && !isNil(o.SslAlertRecipients) {
		return true
	}

	return false
}

// SetSslAlertRecipients gets a reference to the given []string and assigns it to the SslAlertRecipients field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetSslAlertRecipients(v []string) {
	o.SslAlertRecipients = v
}

// GetAutoMonitor returns the AutoMonitor field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetAutoMonitor() bool {
	if o == nil || isNil(o.AutoMonitor) {
		var ret bool
		return ret
	}
	return *o.AutoMonitor
}

// GetAutoMonitorOk returns a tuple with the AutoMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetAutoMonitorOk() (*bool, bool) {
	if o == nil || isNil(o.AutoMonitor) {
		return nil, false
	}
	return o.AutoMonitor, true
}

// HasAutoMonitor returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasAutoMonitor() bool {
	if o != nil && !isNil(o.AutoMonitor) {
		return true
	}

	return false
}

// SetAutoMonitor gets a reference to the given bool and assigns it to the AutoMonitor field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetAutoMonitor(v bool) {
	o.AutoMonitor = &v
}

// GetGetRobots returns the GetRobots field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetGetRobots() bool {
	if o == nil || isNil(o.GetRobots) {
		var ret bool
		return ret
	}
	return *o.GetRobots
}

// GetGetRobotsOk returns a tuple with the GetRobots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetGetRobotsOk() (*bool, bool) {
	if o == nil || isNil(o.GetRobots) {
		return nil, false
	}
	return o.GetRobots, true
}

// HasGetRobots returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasGetRobots() bool {
	if o != nil && !isNil(o.GetRobots) {
		return true
	}

	return false
}

// SetGetRobots gets a reference to the given bool and assigns it to the GetRobots field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetGetRobots(v bool) {
	o.GetRobots = &v
}

// GetDiscoverTimeoutMs returns the DiscoverTimeoutMs field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverTimeoutMs() float64 {
	if o == nil || isNil(o.DiscoverTimeoutMs) {
		var ret float64
		return ret
	}
	return *o.DiscoverTimeoutMs
}

// GetDiscoverTimeoutMsOk returns a tuple with the DiscoverTimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverTimeoutMsOk() (*float64, bool) {
	if o == nil || isNil(o.DiscoverTimeoutMs) {
		return nil, false
	}
	return o.DiscoverTimeoutMs, true
}

// HasDiscoverTimeoutMs returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverTimeoutMs() bool {
	if o != nil && !isNil(o.DiscoverTimeoutMs) {
		return true
	}

	return false
}

// SetDiscoverTimeoutMs gets a reference to the given float64 and assigns it to the DiscoverTimeoutMs field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverTimeoutMs(v float64) {
	o.DiscoverTimeoutMs = &v
}

// GetMonitorTimeoutMs returns the MonitorTimeoutMs field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorTimeoutMs() float64 {
	if o == nil || isNil(o.MonitorTimeoutMs) {
		var ret float64
		return ret
	}
	return *o.MonitorTimeoutMs
}

// GetMonitorTimeoutMsOk returns a tuple with the MonitorTimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorTimeoutMsOk() (*float64, bool) {
	if o == nil || isNil(o.MonitorTimeoutMs) {
		return nil, false
	}
	return o.MonitorTimeoutMs, true
}

// HasMonitorTimeoutMs returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorTimeoutMs() bool {
	if o != nil && !isNil(o.MonitorTimeoutMs) {
		return true
	}

	return false
}

// SetMonitorTimeoutMs gets a reference to the given float64 and assigns it to the MonitorTimeoutMs field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorTimeoutMs(v float64) {
	o.MonitorTimeoutMs = &v
}

// GetExpirationAlertDays returns the ExpirationAlertDays field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetExpirationAlertDays() float64 {
	if o == nil || isNil(o.ExpirationAlertDays) {
		var ret float64
		return ret
	}
	return *o.ExpirationAlertDays
}

// GetExpirationAlertDaysOk returns a tuple with the ExpirationAlertDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetExpirationAlertDaysOk() (*float64, bool) {
	if o == nil || isNil(o.ExpirationAlertDays) {
		return nil, false
	}
	return o.ExpirationAlertDays, true
}

// HasExpirationAlertDays returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasExpirationAlertDays() bool {
	if o != nil && !isNil(o.ExpirationAlertDays) {
		return true
	}

	return false
}

// SetExpirationAlertDays gets a reference to the given float64 and assigns it to the ExpirationAlertDays field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetExpirationAlertDays(v float64) {
	o.ExpirationAlertDays = &v
}

// GetDiscoverJobParts returns the DiscoverJobParts field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverJobParts() int32 {
	if o == nil || isNil(o.DiscoverJobParts) {
		var ret int32
		return ret
	}
	return *o.DiscoverJobParts
}

// GetDiscoverJobPartsOk returns a tuple with the DiscoverJobParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetDiscoverJobPartsOk() (*int32, bool) {
	if o == nil || isNil(o.DiscoverJobParts) {
		return nil, false
	}
	return o.DiscoverJobParts, true
}

// HasDiscoverJobParts returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasDiscoverJobParts() bool {
	if o != nil && !isNil(o.DiscoverJobParts) {
		return true
	}

	return false
}

// SetDiscoverJobParts gets a reference to the given int32 and assigns it to the DiscoverJobParts field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetDiscoverJobParts(v int32) {
	o.DiscoverJobParts = &v
}

// GetMonitorJobParts returns the MonitorJobParts field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorJobParts() int32 {
	if o == nil || isNil(o.MonitorJobParts) {
		var ret int32
		return ret
	}
	return *o.MonitorJobParts
}

// GetMonitorJobPartsOk returns a tuple with the MonitorJobParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetMonitorJobPartsOk() (*int32, bool) {
	if o == nil || isNil(o.MonitorJobParts) {
		return nil, false
	}
	return o.MonitorJobParts, true
}

// HasMonitorJobParts returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasMonitorJobParts() bool {
	if o != nil && !isNil(o.MonitorJobParts) {
		return true
	}

	return false
}

// SetMonitorJobParts gets a reference to the given int32 and assigns it to the MonitorJobParts field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetMonitorJobParts(v int32) {
	o.MonitorJobParts = &v
}

// GetQuietHours returns the QuietHours field value if set, zero value otherwise.
func (o *KeyfactorApiModelsSslNetworkResponse) GetQuietHours() []KeyfactorApiModelsSslQuietHourResponse {
	if o == nil || isNil(o.QuietHours) {
		var ret []KeyfactorApiModelsSslQuietHourResponse
		return ret
	}
	return o.QuietHours
}

// GetQuietHoursOk returns a tuple with the QuietHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) GetQuietHoursOk() ([]KeyfactorApiModelsSslQuietHourResponse, bool) {
	if o == nil || isNil(o.QuietHours) {
		return nil, false
	}
	return o.QuietHours, true
}

// HasQuietHours returns a boolean if a field has been set.
func (o *KeyfactorApiModelsSslNetworkResponse) HasQuietHours() bool {
	if o != nil && !isNil(o.QuietHours) {
		return true
	}

	return false
}

// SetQuietHours gets a reference to the given []KeyfactorApiModelsSslQuietHourResponse and assigns it to the QuietHours field.
func (o *KeyfactorApiModelsSslNetworkResponse) SetQuietHours(v []KeyfactorApiModelsSslQuietHourResponse) {
	o.QuietHours = v
}

func (o KeyfactorApiModelsSslNetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsSslNetworkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["NetworkId"] = o.NetworkId
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.AgentPoolName) {
		toSerialize["AgentPoolName"] = o.AgentPoolName
	}
	if !isNil(o.AgentPoolId) {
		toSerialize["AgentPoolId"] = o.AgentPoolId
	}
	if !isNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !isNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !isNil(o.DiscoverSchedule) {
		toSerialize["DiscoverSchedule"] = o.DiscoverSchedule
	}
	if !isNil(o.MonitorSchedule) {
		toSerialize["MonitorSchedule"] = o.MonitorSchedule
	}
	if !isNil(o.DiscoverPercentComplete) {
		toSerialize["DiscoverPercentComplete"] = o.DiscoverPercentComplete
	}
	if !isNil(o.MonitorPercentComplete) {
		toSerialize["MonitorPercentComplete"] = o.MonitorPercentComplete
	}
	if !isNil(o.DiscoverStatus) {
		toSerialize["DiscoverStatus"] = o.DiscoverStatus
	}
	if !isNil(o.MonitorStatus) {
		toSerialize["MonitorStatus"] = o.MonitorStatus
	}
	if !isNil(o.DiscoverLastScanned) {
		toSerialize["DiscoverLastScanned"] = o.DiscoverLastScanned
	}
	if !isNil(o.MonitorLastScanned) {
		toSerialize["MonitorLastScanned"] = o.MonitorLastScanned
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
	if !isNil(o.DiscoverJobParts) {
		toSerialize["DiscoverJobParts"] = o.DiscoverJobParts
	}
	if !isNil(o.MonitorJobParts) {
		toSerialize["MonitorJobParts"] = o.MonitorJobParts
	}
	if !isNil(o.QuietHours) {
		toSerialize["QuietHours"] = o.QuietHours
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsSslNetworkResponse struct {
	value *KeyfactorApiModelsSslNetworkResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsSslNetworkResponse) Get() *KeyfactorApiModelsSslNetworkResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsSslNetworkResponse) Set(val *KeyfactorApiModelsSslNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsSslNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsSslNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsSslNetworkResponse(val *KeyfactorApiModelsSslNetworkResponse) *NullableKeyfactorApiModelsSslNetworkResponse {
	return &NullableKeyfactorApiModelsSslNetworkResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsSslNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsSslNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


