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

// checks if the ModelsAgentsAgentPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsAgentsAgentPool{}

// ModelsAgentsAgentPool Class representing an SSL agent pool
type ModelsAgentsAgentPool struct {
	// GUID identifier of the agent pool
	AgentPoolId *string `json:"AgentPoolId,omitempty"`
	// Name of the agent pool
	Name string `json:"Name"`
	// Number of agents that can perform discovery jobs
	DiscoverAgentsCount *int32 `json:"DiscoverAgentsCount,omitempty"`
	// Number of agents that can perform monitoring jobs
	MonitorAgentsCount *int32 `json:"MonitorAgentsCount,omitempty"`
	// List of the agents assigned to the pool
	Agents []ModelsAgentsAgentPoolAgent `json:"Agents,omitempty"`
}

// NewModelsAgentsAgentPool instantiates a new ModelsAgentsAgentPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsAgentsAgentPool(name string) *ModelsAgentsAgentPool {
	this := ModelsAgentsAgentPool{}
	this.Name = name
	return &this
}

// NewModelsAgentsAgentPoolWithDefaults instantiates a new ModelsAgentsAgentPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsAgentsAgentPoolWithDefaults() *ModelsAgentsAgentPool {
	this := ModelsAgentsAgentPool{}
	return &this
}

// GetAgentPoolId returns the AgentPoolId field value if set, zero value otherwise.
func (o *ModelsAgentsAgentPool) GetAgentPoolId() string {
	if o == nil || isNil(o.AgentPoolId) {
		var ret string
		return ret
	}
	return *o.AgentPoolId
}

// GetAgentPoolIdOk returns a tuple with the AgentPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAgentsAgentPool) GetAgentPoolIdOk() (*string, bool) {
	if o == nil || isNil(o.AgentPoolId) {
		return nil, false
	}
	return o.AgentPoolId, true
}

// HasAgentPoolId returns a boolean if a field has been set.
func (o *ModelsAgentsAgentPool) HasAgentPoolId() bool {
	if o != nil && !isNil(o.AgentPoolId) {
		return true
	}

	return false
}

// SetAgentPoolId gets a reference to the given string and assigns it to the AgentPoolId field.
func (o *ModelsAgentsAgentPool) SetAgentPoolId(v string) {
	o.AgentPoolId = &v
}

// GetName returns the Name field value
func (o *ModelsAgentsAgentPool) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelsAgentsAgentPool) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelsAgentsAgentPool) SetName(v string) {
	o.Name = v
}

// GetDiscoverAgentsCount returns the DiscoverAgentsCount field value if set, zero value otherwise.
func (o *ModelsAgentsAgentPool) GetDiscoverAgentsCount() int32 {
	if o == nil || isNil(o.DiscoverAgentsCount) {
		var ret int32
		return ret
	}
	return *o.DiscoverAgentsCount
}

// GetDiscoverAgentsCountOk returns a tuple with the DiscoverAgentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAgentsAgentPool) GetDiscoverAgentsCountOk() (*int32, bool) {
	if o == nil || isNil(o.DiscoverAgentsCount) {
		return nil, false
	}
	return o.DiscoverAgentsCount, true
}

// HasDiscoverAgentsCount returns a boolean if a field has been set.
func (o *ModelsAgentsAgentPool) HasDiscoverAgentsCount() bool {
	if o != nil && !isNil(o.DiscoverAgentsCount) {
		return true
	}

	return false
}

// SetDiscoverAgentsCount gets a reference to the given int32 and assigns it to the DiscoverAgentsCount field.
func (o *ModelsAgentsAgentPool) SetDiscoverAgentsCount(v int32) {
	o.DiscoverAgentsCount = &v
}

// GetMonitorAgentsCount returns the MonitorAgentsCount field value if set, zero value otherwise.
func (o *ModelsAgentsAgentPool) GetMonitorAgentsCount() int32 {
	if o == nil || isNil(o.MonitorAgentsCount) {
		var ret int32
		return ret
	}
	return *o.MonitorAgentsCount
}

// GetMonitorAgentsCountOk returns a tuple with the MonitorAgentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAgentsAgentPool) GetMonitorAgentsCountOk() (*int32, bool) {
	if o == nil || isNil(o.MonitorAgentsCount) {
		return nil, false
	}
	return o.MonitorAgentsCount, true
}

// HasMonitorAgentsCount returns a boolean if a field has been set.
func (o *ModelsAgentsAgentPool) HasMonitorAgentsCount() bool {
	if o != nil && !isNil(o.MonitorAgentsCount) {
		return true
	}

	return false
}

// SetMonitorAgentsCount gets a reference to the given int32 and assigns it to the MonitorAgentsCount field.
func (o *ModelsAgentsAgentPool) SetMonitorAgentsCount(v int32) {
	o.MonitorAgentsCount = &v
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *ModelsAgentsAgentPool) GetAgents() []ModelsAgentsAgentPoolAgent {
	if o == nil || isNil(o.Agents) {
		var ret []ModelsAgentsAgentPoolAgent
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAgentsAgentPool) GetAgentsOk() ([]ModelsAgentsAgentPoolAgent, bool) {
	if o == nil || isNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *ModelsAgentsAgentPool) HasAgents() bool {
	if o != nil && !isNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []ModelsAgentsAgentPoolAgent and assigns it to the Agents field.
func (o *ModelsAgentsAgentPool) SetAgents(v []ModelsAgentsAgentPoolAgent) {
	o.Agents = v
}

func (o ModelsAgentsAgentPool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsAgentsAgentPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AgentPoolId) {
		toSerialize["AgentPoolId"] = o.AgentPoolId
	}
	toSerialize["Name"] = o.Name
	if !isNil(o.DiscoverAgentsCount) {
		toSerialize["DiscoverAgentsCount"] = o.DiscoverAgentsCount
	}
	if !isNil(o.MonitorAgentsCount) {
		toSerialize["MonitorAgentsCount"] = o.MonitorAgentsCount
	}
	if !isNil(o.Agents) {
		toSerialize["Agents"] = o.Agents
	}
	return toSerialize, nil
}

type NullableModelsAgentsAgentPool struct {
	value *ModelsAgentsAgentPool
	isSet bool
}

func (v NullableModelsAgentsAgentPool) Get() *ModelsAgentsAgentPool {
	return v.value
}

func (v *NullableModelsAgentsAgentPool) Set(val *ModelsAgentsAgentPool) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsAgentsAgentPool) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsAgentsAgentPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsAgentsAgentPool(val *ModelsAgentsAgentPool) *NullableModelsAgentsAgentPool {
	return &NullableModelsAgentsAgentPool{value: val, isSet: true}
}

func (v NullableModelsAgentsAgentPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsAgentsAgentPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


