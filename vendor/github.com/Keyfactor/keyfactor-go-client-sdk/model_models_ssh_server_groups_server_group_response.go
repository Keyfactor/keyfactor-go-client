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

// checks if the ModelsSSHServerGroupsServerGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHServerGroupsServerGroupResponse{}

// ModelsSSHServerGroupsServerGroupResponse struct for ModelsSSHServerGroupsServerGroupResponse
type ModelsSSHServerGroupsServerGroupResponse struct {
	Id *string `json:"Id,omitempty"`
	Owner *ModelsSSHUsersSshUserResponse `json:"Owner,omitempty"`
	GroupName *string `json:"GroupName,omitempty"`
	SyncSchedule *KeyfactorCommonSchedulingKeyfactorSchedule `json:"SyncSchedule,omitempty"`
	UnderManagement *bool `json:"UnderManagement,omitempty"`
	ServerCount *int32 `json:"ServerCount,omitempty"`
}

// NewModelsSSHServerGroupsServerGroupResponse instantiates a new ModelsSSHServerGroupsServerGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHServerGroupsServerGroupResponse() *ModelsSSHServerGroupsServerGroupResponse {
	this := ModelsSSHServerGroupsServerGroupResponse{}
	return &this
}

// NewModelsSSHServerGroupsServerGroupResponseWithDefaults instantiates a new ModelsSSHServerGroupsServerGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHServerGroupsServerGroupResponseWithDefaults() *ModelsSSHServerGroupsServerGroupResponse {
	this := ModelsSSHServerGroupsServerGroupResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsSSHServerGroupsServerGroupResponse) SetId(v string) {
	o.Id = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetOwner() ModelsSSHUsersSshUserResponse {
	if o == nil || isNil(o.Owner) {
		var ret ModelsSSHUsersSshUserResponse
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetOwnerOk() (*ModelsSSHUsersSshUserResponse, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ModelsSSHUsersSshUserResponse and assigns it to the Owner field.
func (o *ModelsSSHServerGroupsServerGroupResponse) SetOwner(v ModelsSSHUsersSshUserResponse) {
	o.Owner = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetGroupName() string {
	if o == nil || isNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetGroupNameOk() (*string, bool) {
	if o == nil || isNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) HasGroupName() bool {
	if o != nil && !isNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ModelsSSHServerGroupsServerGroupResponse) SetGroupName(v string) {
	o.GroupName = &v
}

// GetSyncSchedule returns the SyncSchedule field value if set, zero value otherwise.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetSyncSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.SyncSchedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.SyncSchedule
}

// GetSyncScheduleOk returns a tuple with the SyncSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetSyncScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.SyncSchedule) {
		return nil, false
	}
	return o.SyncSchedule, true
}

// HasSyncSchedule returns a boolean if a field has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) HasSyncSchedule() bool {
	if o != nil && !isNil(o.SyncSchedule) {
		return true
	}

	return false
}

// SetSyncSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the SyncSchedule field.
func (o *ModelsSSHServerGroupsServerGroupResponse) SetSyncSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.SyncSchedule = &v
}

// GetUnderManagement returns the UnderManagement field value if set, zero value otherwise.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetUnderManagement() bool {
	if o == nil || isNil(o.UnderManagement) {
		var ret bool
		return ret
	}
	return *o.UnderManagement
}

// GetUnderManagementOk returns a tuple with the UnderManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetUnderManagementOk() (*bool, bool) {
	if o == nil || isNil(o.UnderManagement) {
		return nil, false
	}
	return o.UnderManagement, true
}

// HasUnderManagement returns a boolean if a field has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) HasUnderManagement() bool {
	if o != nil && !isNil(o.UnderManagement) {
		return true
	}

	return false
}

// SetUnderManagement gets a reference to the given bool and assigns it to the UnderManagement field.
func (o *ModelsSSHServerGroupsServerGroupResponse) SetUnderManagement(v bool) {
	o.UnderManagement = &v
}

// GetServerCount returns the ServerCount field value if set, zero value otherwise.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetServerCount() int32 {
	if o == nil || isNil(o.ServerCount) {
		var ret int32
		return ret
	}
	return *o.ServerCount
}

// GetServerCountOk returns a tuple with the ServerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) GetServerCountOk() (*int32, bool) {
	if o == nil || isNil(o.ServerCount) {
		return nil, false
	}
	return o.ServerCount, true
}

// HasServerCount returns a boolean if a field has been set.
func (o *ModelsSSHServerGroupsServerGroupResponse) HasServerCount() bool {
	if o != nil && !isNil(o.ServerCount) {
		return true
	}

	return false
}

// SetServerCount gets a reference to the given int32 and assigns it to the ServerCount field.
func (o *ModelsSSHServerGroupsServerGroupResponse) SetServerCount(v int32) {
	o.ServerCount = &v
}

func (o ModelsSSHServerGroupsServerGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHServerGroupsServerGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.Owner) {
		toSerialize["Owner"] = o.Owner
	}
	if !isNil(o.GroupName) {
		toSerialize["GroupName"] = o.GroupName
	}
	if !isNil(o.SyncSchedule) {
		toSerialize["SyncSchedule"] = o.SyncSchedule
	}
	if !isNil(o.UnderManagement) {
		toSerialize["UnderManagement"] = o.UnderManagement
	}
	if !isNil(o.ServerCount) {
		toSerialize["ServerCount"] = o.ServerCount
	}
	return toSerialize, nil
}

type NullableModelsSSHServerGroupsServerGroupResponse struct {
	value *ModelsSSHServerGroupsServerGroupResponse
	isSet bool
}

func (v NullableModelsSSHServerGroupsServerGroupResponse) Get() *ModelsSSHServerGroupsServerGroupResponse {
	return v.value
}

func (v *NullableModelsSSHServerGroupsServerGroupResponse) Set(val *ModelsSSHServerGroupsServerGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHServerGroupsServerGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHServerGroupsServerGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHServerGroupsServerGroupResponse(val *ModelsSSHServerGroupsServerGroupResponse) *NullableModelsSSHServerGroupsServerGroupResponse {
	return &NullableModelsSSHServerGroupsServerGroupResponse{value: val, isSet: true}
}

func (v NullableModelsSSHServerGroupsServerGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHServerGroupsServerGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


