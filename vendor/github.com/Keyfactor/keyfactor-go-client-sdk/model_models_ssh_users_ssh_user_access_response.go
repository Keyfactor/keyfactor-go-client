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

// checks if the ModelsSSHUsersSshUserAccessResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHUsersSshUserAccessResponse{}

// ModelsSSHUsersSshUserAccessResponse struct for ModelsSSHUsersSshUserAccessResponse
type ModelsSSHUsersSshUserAccessResponse struct {
	Id *int32 `json:"Id,omitempty"`
	Key *ModelsSSHKeysKeyResponse `json:"Key,omitempty"`
	Username *string `json:"Username,omitempty"`
	Access []ModelsSSHLogonsLogonResponse `json:"Access,omitempty"`
	IsGroup *bool `json:"IsGroup,omitempty"`
}

// NewModelsSSHUsersSshUserAccessResponse instantiates a new ModelsSSHUsersSshUserAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHUsersSshUserAccessResponse() *ModelsSSHUsersSshUserAccessResponse {
	this := ModelsSSHUsersSshUserAccessResponse{}
	return &this
}

// NewModelsSSHUsersSshUserAccessResponseWithDefaults instantiates a new ModelsSSHUsersSshUserAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHUsersSshUserAccessResponseWithDefaults() *ModelsSSHUsersSshUserAccessResponse {
	this := ModelsSSHUsersSshUserAccessResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsSSHUsersSshUserAccessResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsSSHUsersSshUserAccessResponse) SetId(v int32) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ModelsSSHUsersSshUserAccessResponse) GetKey() ModelsSSHKeysKeyResponse {
	if o == nil || isNil(o.Key) {
		var ret ModelsSSHKeysKeyResponse
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) GetKeyOk() (*ModelsSSHKeysKeyResponse, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given ModelsSSHKeysKeyResponse and assigns it to the Key field.
func (o *ModelsSSHUsersSshUserAccessResponse) SetKey(v ModelsSSHKeysKeyResponse) {
	o.Key = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ModelsSSHUsersSshUserAccessResponse) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ModelsSSHUsersSshUserAccessResponse) SetUsername(v string) {
	o.Username = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ModelsSSHUsersSshUserAccessResponse) GetAccess() []ModelsSSHLogonsLogonResponse {
	if o == nil || isNil(o.Access) {
		var ret []ModelsSSHLogonsLogonResponse
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) GetAccessOk() ([]ModelsSSHLogonsLogonResponse, bool) {
	if o == nil || isNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given []ModelsSSHLogonsLogonResponse and assigns it to the Access field.
func (o *ModelsSSHUsersSshUserAccessResponse) SetAccess(v []ModelsSSHLogonsLogonResponse) {
	o.Access = v
}

// GetIsGroup returns the IsGroup field value if set, zero value otherwise.
func (o *ModelsSSHUsersSshUserAccessResponse) GetIsGroup() bool {
	if o == nil || isNil(o.IsGroup) {
		var ret bool
		return ret
	}
	return *o.IsGroup
}

// GetIsGroupOk returns a tuple with the IsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) GetIsGroupOk() (*bool, bool) {
	if o == nil || isNil(o.IsGroup) {
		return nil, false
	}
	return o.IsGroup, true
}

// HasIsGroup returns a boolean if a field has been set.
func (o *ModelsSSHUsersSshUserAccessResponse) HasIsGroup() bool {
	if o != nil && !isNil(o.IsGroup) {
		return true
	}

	return false
}

// SetIsGroup gets a reference to the given bool and assigns it to the IsGroup field.
func (o *ModelsSSHUsersSshUserAccessResponse) SetIsGroup(v bool) {
	o.IsGroup = &v
}

func (o ModelsSSHUsersSshUserAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHUsersSshUserAccessResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !isNil(o.Username) {
		toSerialize["Username"] = o.Username
	}
	if !isNil(o.Access) {
		toSerialize["Access"] = o.Access
	}
	if !isNil(o.IsGroup) {
		toSerialize["IsGroup"] = o.IsGroup
	}
	return toSerialize, nil
}

type NullableModelsSSHUsersSshUserAccessResponse struct {
	value *ModelsSSHUsersSshUserAccessResponse
	isSet bool
}

func (v NullableModelsSSHUsersSshUserAccessResponse) Get() *ModelsSSHUsersSshUserAccessResponse {
	return v.value
}

func (v *NullableModelsSSHUsersSshUserAccessResponse) Set(val *ModelsSSHUsersSshUserAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHUsersSshUserAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHUsersSshUserAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHUsersSshUserAccessResponse(val *ModelsSSHUsersSshUserAccessResponse) *NullableModelsSSHUsersSshUserAccessResponse {
	return &NullableModelsSSHUsersSshUserAccessResponse{value: val, isSet: true}
}

func (v NullableModelsSSHUsersSshUserAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHUsersSshUserAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


