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

// checks if the ModelsSSHKeysKeyUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSSHKeysKeyUpdateRequest{}

// ModelsSSHKeysKeyUpdateRequest struct for ModelsSSHKeysKeyUpdateRequest
type ModelsSSHKeysKeyUpdateRequest struct {
	Id int32 `json:"Id"`
	Email string `json:"Email"`
	Comment *string `json:"Comment,omitempty"`
}

// NewModelsSSHKeysKeyUpdateRequest instantiates a new ModelsSSHKeysKeyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSSHKeysKeyUpdateRequest(id int32, email string) *ModelsSSHKeysKeyUpdateRequest {
	this := ModelsSSHKeysKeyUpdateRequest{}
	this.Id = id
	this.Email = email
	return &this
}

// NewModelsSSHKeysKeyUpdateRequestWithDefaults instantiates a new ModelsSSHKeysKeyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSSHKeysKeyUpdateRequestWithDefaults() *ModelsSSHKeysKeyUpdateRequest {
	this := ModelsSSHKeysKeyUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ModelsSSHKeysKeyUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHKeysKeyUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelsSSHKeysKeyUpdateRequest) SetId(v int32) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *ModelsSSHKeysKeyUpdateRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ModelsSSHKeysKeyUpdateRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ModelsSSHKeysKeyUpdateRequest) SetEmail(v string) {
	o.Email = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ModelsSSHKeysKeyUpdateRequest) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSSHKeysKeyUpdateRequest) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ModelsSSHKeysKeyUpdateRequest) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ModelsSSHKeysKeyUpdateRequest) SetComment(v string) {
	o.Comment = &v
}

func (o ModelsSSHKeysKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSSHKeysKeyUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["Email"] = o.Email
	if !isNil(o.Comment) {
		toSerialize["Comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableModelsSSHKeysKeyUpdateRequest struct {
	value *ModelsSSHKeysKeyUpdateRequest
	isSet bool
}

func (v NullableModelsSSHKeysKeyUpdateRequest) Get() *ModelsSSHKeysKeyUpdateRequest {
	return v.value
}

func (v *NullableModelsSSHKeysKeyUpdateRequest) Set(val *ModelsSSHKeysKeyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSSHKeysKeyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSSHKeysKeyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSSHKeysKeyUpdateRequest(val *ModelsSSHKeysKeyUpdateRequest) *NullableModelsSSHKeysKeyUpdateRequest {
	return &NullableModelsSSHKeysKeyUpdateRequest{value: val, isSet: true}
}

func (v NullableModelsSSHKeysKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSSHKeysKeyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


