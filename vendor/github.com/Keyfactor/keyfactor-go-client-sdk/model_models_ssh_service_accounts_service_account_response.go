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

// checks if the ModelsSSHServiceAccountsServiceAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHServiceAccountsServiceAccountResponse{}

// ModelsSSHServiceAccountsServiceAccountResponse struct for ModelsSSHServiceAccountsServiceAccountResponse
type ModelsSSHServiceAccountsServiceAccountResponse struct {
	Id *int32 `json:"Id,omitempty"`
	ClientHostname *string `json:"ClientHostname,omitempty"`
	ServerGroup *ModelsSSHServerGroupsServerGroupResponse `json:"ServerGroup,omitempty"`
	User *ModelsSSHUsersSshUserResponse `json:"User,omitempty"`
}

// NewModelsSSHServiceAccountsServiceAccountResponse instantiates a new ModelsSSHServiceAccountsServiceAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHServiceAccountsServiceAccountResponse() *ModelsSSHServiceAccountsServiceAccountResponse {
	this := ModelsSSHServiceAccountsServiceAccountResponse{}
	return &this
}

// NewModelsSSHServiceAccountsServiceAccountResponseWithDefaults instantiates a new ModelsSSHServiceAccountsServiceAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHServiceAccountsServiceAccountResponseWithDefaults() *ModelsSSHServiceAccountsServiceAccountResponse {
	this := ModelsSSHServiceAccountsServiceAccountResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) SetId(v int32) {
	o.Id = &v
}

// GetClientHostname returns the ClientHostname field value if set, zero value otherwise.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetClientHostname() string {
	if o == nil || isNil(o.ClientHostname) {
		var ret string
		return ret
	}
	return *o.ClientHostname
}

// GetClientHostnameOk returns a tuple with the ClientHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetClientHostnameOk() (*string, bool) {
	if o == nil || isNil(o.ClientHostname) {
		return nil, false
	}
	return o.ClientHostname, true
}

// HasClientHostname returns a boolean if a field has been set.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) HasClientHostname() bool {
	if o != nil && !isNil(o.ClientHostname) {
		return true
	}

	return false
}

// SetClientHostname gets a reference to the given string and assigns it to the ClientHostname field.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) SetClientHostname(v string) {
	o.ClientHostname = &v
}

// GetServerGroup returns the ServerGroup field value if set, zero value otherwise.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetServerGroup() ModelsSSHServerGroupsServerGroupResponse {
	if o == nil || isNil(o.ServerGroup) {
		var ret ModelsSSHServerGroupsServerGroupResponse
		return ret
	}
	return *o.ServerGroup
}

// GetServerGroupOk returns a tuple with the ServerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetServerGroupOk() (*ModelsSSHServerGroupsServerGroupResponse, bool) {
	if o == nil || isNil(o.ServerGroup) {
		return nil, false
	}
	return o.ServerGroup, true
}

// HasServerGroup returns a boolean if a field has been set.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) HasServerGroup() bool {
	if o != nil && !isNil(o.ServerGroup) {
		return true
	}

	return false
}

// SetServerGroup gets a reference to the given ModelsSSHServerGroupsServerGroupResponse and assigns it to the ServerGroup field.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) SetServerGroup(v ModelsSSHServerGroupsServerGroupResponse) {
	o.ServerGroup = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetUser() ModelsSSHUsersSshUserResponse {
	if o == nil || isNil(o.User) {
		var ret ModelsSSHUsersSshUserResponse
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) GetUserOk() (*ModelsSSHUsersSshUserResponse, bool) {
	if o == nil || isNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given ModelsSSHUsersSshUserResponse and assigns it to the User field.
func (o *ModelsSSHServiceAccountsServiceAccountResponse) SetUser(v ModelsSSHUsersSshUserResponse) {
	o.User = &v
}

func (o ModelsSSHServiceAccountsServiceAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHServiceAccountsServiceAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.ClientHostname) {
		toSerialize["ClientHostname"] = o.ClientHostname
	}
	if !isNil(o.ServerGroup) {
		toSerialize["ServerGroup"] = o.ServerGroup
	}
	if !isNil(o.User) {
		toSerialize["User"] = o.User
	}
	return toSerialize, nil
}

type NullableModelsSSHServiceAccountsServiceAccountResponse struct {
	value *ModelsSSHServiceAccountsServiceAccountResponse
	isSet bool
}

func (v NullableModelsSSHServiceAccountsServiceAccountResponse) Get() *ModelsSSHServiceAccountsServiceAccountResponse {
	return v.value
}

func (v *NullableModelsSSHServiceAccountsServiceAccountResponse) Set(val *ModelsSSHServiceAccountsServiceAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHServiceAccountsServiceAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHServiceAccountsServiceAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHServiceAccountsServiceAccountResponse(val *ModelsSSHServiceAccountsServiceAccountResponse) *NullableModelsSSHServiceAccountsServiceAccountResponse {
	return &NullableModelsSSHServiceAccountsServiceAccountResponse{value: val, isSet: true}
}

func (v NullableModelsSSHServiceAccountsServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHServiceAccountsServiceAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


