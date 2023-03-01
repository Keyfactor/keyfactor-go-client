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

// checks if the ModelsCertStoreLocationsCertificateStoreLocationsDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertStoreLocationsCertificateStoreLocationsDetail{}

// ModelsCertStoreLocationsCertificateStoreLocationsDetail struct for ModelsCertStoreLocationsCertificateStoreLocationsDetail
type ModelsCertStoreLocationsCertificateStoreLocationsDetail struct {
	StoreId *string `json:"StoreId,omitempty"`
	StoreTypeId *int32 `json:"StoreTypeId,omitempty"`
	ClientMachine *string `json:"ClientMachine,omitempty"`
	StorePath *string `json:"StorePath,omitempty"`
	AgentPool *string `json:"AgentPool,omitempty"`
	Alias *string `json:"Alias,omitempty"`
	IPAddress *string `json:"IPAddress,omitempty"`
	Port *int32 `json:"Port,omitempty"`
	NetworkName *string `json:"NetworkName,omitempty"`
}

// NewModelsCertStoreLocationsCertificateStoreLocationsDetail instantiates a new ModelsCertStoreLocationsCertificateStoreLocationsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertStoreLocationsCertificateStoreLocationsDetail() *ModelsCertStoreLocationsCertificateStoreLocationsDetail {
	this := ModelsCertStoreLocationsCertificateStoreLocationsDetail{}
	return &this
}

// NewModelsCertStoreLocationsCertificateStoreLocationsDetailWithDefaults instantiates a new ModelsCertStoreLocationsCertificateStoreLocationsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertStoreLocationsCertificateStoreLocationsDetailWithDefaults() *ModelsCertStoreLocationsCertificateStoreLocationsDetail {
	this := ModelsCertStoreLocationsCertificateStoreLocationsDetail{}
	return &this
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetStoreId() string {
	if o == nil || isNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetStoreIdOk() (*string, bool) {
	if o == nil || isNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) HasStoreId() bool {
	if o != nil && !isNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) SetStoreId(v string) {
	o.StoreId = &v
}

// GetStoreTypeId returns the StoreTypeId field value if set, zero value otherwise.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetStoreTypeId() int32 {
	if o == nil || isNil(o.StoreTypeId) {
		var ret int32
		return ret
	}
	return *o.StoreTypeId
}

// GetStoreTypeIdOk returns a tuple with the StoreTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetStoreTypeIdOk() (*int32, bool) {
	if o == nil || isNil(o.StoreTypeId) {
		return nil, false
	}
	return o.StoreTypeId, true
}

// HasStoreTypeId returns a boolean if a field has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) HasStoreTypeId() bool {
	if o != nil && !isNil(o.StoreTypeId) {
		return true
	}

	return false
}

// SetStoreTypeId gets a reference to the given int32 and assigns it to the StoreTypeId field.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) SetStoreTypeId(v int32) {
	o.StoreTypeId = &v
}

// GetClientMachine returns the ClientMachine field value if set, zero value otherwise.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetClientMachine() string {
	if o == nil || isNil(o.ClientMachine) {
		var ret string
		return ret
	}
	return *o.ClientMachine
}

// GetClientMachineOk returns a tuple with the ClientMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetClientMachineOk() (*string, bool) {
	if o == nil || isNil(o.ClientMachine) {
		return nil, false
	}
	return o.ClientMachine, true
}

// HasClientMachine returns a boolean if a field has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) HasClientMachine() bool {
	if o != nil && !isNil(o.ClientMachine) {
		return true
	}

	return false
}

// SetClientMachine gets a reference to the given string and assigns it to the ClientMachine field.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) SetClientMachine(v string) {
	o.ClientMachine = &v
}

// GetStorePath returns the StorePath field value if set, zero value otherwise.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetStorePath() string {
	if o == nil || isNil(o.StorePath) {
		var ret string
		return ret
	}
	return *o.StorePath
}

// GetStorePathOk returns a tuple with the StorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetStorePathOk() (*string, bool) {
	if o == nil || isNil(o.StorePath) {
		return nil, false
	}
	return o.StorePath, true
}

// HasStorePath returns a boolean if a field has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) HasStorePath() bool {
	if o != nil && !isNil(o.StorePath) {
		return true
	}

	return false
}

// SetStorePath gets a reference to the given string and assigns it to the StorePath field.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) SetStorePath(v string) {
	o.StorePath = &v
}

// GetAgentPool returns the AgentPool field value if set, zero value otherwise.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetAgentPool() string {
	if o == nil || isNil(o.AgentPool) {
		var ret string
		return ret
	}
	return *o.AgentPool
}

// GetAgentPoolOk returns a tuple with the AgentPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetAgentPoolOk() (*string, bool) {
	if o == nil || isNil(o.AgentPool) {
		return nil, false
	}
	return o.AgentPool, true
}

// HasAgentPool returns a boolean if a field has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) HasAgentPool() bool {
	if o != nil && !isNil(o.AgentPool) {
		return true
	}

	return false
}

// SetAgentPool gets a reference to the given string and assigns it to the AgentPool field.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) SetAgentPool(v string) {
	o.AgentPool = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetAlias() string {
	if o == nil || isNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetAliasOk() (*string, bool) {
	if o == nil || isNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) HasAlias() bool {
	if o != nil && !isNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) SetAlias(v string) {
	o.Alias = &v
}

// GetIPAddress returns the IPAddress field value if set, zero value otherwise.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetIPAddress() string {
	if o == nil || isNil(o.IPAddress) {
		var ret string
		return ret
	}
	return *o.IPAddress
}

// GetIPAddressOk returns a tuple with the IPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetIPAddressOk() (*string, bool) {
	if o == nil || isNil(o.IPAddress) {
		return nil, false
	}
	return o.IPAddress, true
}

// HasIPAddress returns a boolean if a field has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) HasIPAddress() bool {
	if o != nil && !isNil(o.IPAddress) {
		return true
	}

	return false
}

// SetIPAddress gets a reference to the given string and assigns it to the IPAddress field.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) SetIPAddress(v string) {
	o.IPAddress = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) SetPort(v int32) {
	o.Port = &v
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetNetworkName() string {
	if o == nil || isNil(o.NetworkName) {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) GetNetworkNameOk() (*string, bool) {
	if o == nil || isNil(o.NetworkName) {
		return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) HasNetworkName() bool {
	if o != nil && !isNil(o.NetworkName) {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *ModelsCertStoreLocationsCertificateStoreLocationsDetail) SetNetworkName(v string) {
	o.NetworkName = &v
}

func (o ModelsCertStoreLocationsCertificateStoreLocationsDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertStoreLocationsCertificateStoreLocationsDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StoreId) {
		toSerialize["StoreId"] = o.StoreId
	}
	if !isNil(o.StoreTypeId) {
		toSerialize["StoreTypeId"] = o.StoreTypeId
	}
	if !isNil(o.ClientMachine) {
		toSerialize["ClientMachine"] = o.ClientMachine
	}
	if !isNil(o.StorePath) {
		toSerialize["StorePath"] = o.StorePath
	}
	if !isNil(o.AgentPool) {
		toSerialize["AgentPool"] = o.AgentPool
	}
	if !isNil(o.Alias) {
		toSerialize["Alias"] = o.Alias
	}
	if !isNil(o.IPAddress) {
		toSerialize["IPAddress"] = o.IPAddress
	}
	if !isNil(o.Port) {
		toSerialize["Port"] = o.Port
	}
	if !isNil(o.NetworkName) {
		toSerialize["NetworkName"] = o.NetworkName
	}
	return toSerialize, nil
}

type NullableModelsCertStoreLocationsCertificateStoreLocationsDetail struct {
	value *ModelsCertStoreLocationsCertificateStoreLocationsDetail
	isSet bool
}

func (v NullableModelsCertStoreLocationsCertificateStoreLocationsDetail) Get() *ModelsCertStoreLocationsCertificateStoreLocationsDetail {
	return v.value
}

func (v *NullableModelsCertStoreLocationsCertificateStoreLocationsDetail) Set(val *ModelsCertStoreLocationsCertificateStoreLocationsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertStoreLocationsCertificateStoreLocationsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertStoreLocationsCertificateStoreLocationsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertStoreLocationsCertificateStoreLocationsDetail(val *ModelsCertStoreLocationsCertificateStoreLocationsDetail) *NullableModelsCertStoreLocationsCertificateStoreLocationsDetail {
	return &NullableModelsCertStoreLocationsCertificateStoreLocationsDetail{value: val, isSet: true}
}

func (v NullableModelsCertStoreLocationsCertificateStoreLocationsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertStoreLocationsCertificateStoreLocationsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


