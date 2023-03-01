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

// checks if the ModelsSecurityIdentitiesSecurityIdentityIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSecurityIdentitiesSecurityIdentityIdentifier{}

// ModelsSecurityIdentitiesSecurityIdentityIdentifier Model for looking up a security identity
type ModelsSecurityIdentitiesSecurityIdentityIdentifier struct {
	// The username of the security identity.
	AccountName *string `json:"AccountName,omitempty"`
	// The SID of the security identity.
	SID *string `json:"SID,omitempty"`
}

// NewModelsSecurityIdentitiesSecurityIdentityIdentifier instantiates a new ModelsSecurityIdentitiesSecurityIdentityIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSecurityIdentitiesSecurityIdentityIdentifier() *ModelsSecurityIdentitiesSecurityIdentityIdentifier {
	this := ModelsSecurityIdentitiesSecurityIdentityIdentifier{}
	return &this
}

// NewModelsSecurityIdentitiesSecurityIdentityIdentifierWithDefaults instantiates a new ModelsSecurityIdentitiesSecurityIdentityIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSecurityIdentitiesSecurityIdentityIdentifierWithDefaults() *ModelsSecurityIdentitiesSecurityIdentityIdentifier {
	this := ModelsSecurityIdentitiesSecurityIdentityIdentifier{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ModelsSecurityIdentitiesSecurityIdentityIdentifier) GetAccountName() string {
	if o == nil || isNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecurityIdentitiesSecurityIdentityIdentifier) GetAccountNameOk() (*string, bool) {
	if o == nil || isNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ModelsSecurityIdentitiesSecurityIdentityIdentifier) HasAccountName() bool {
	if o != nil && !isNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ModelsSecurityIdentitiesSecurityIdentityIdentifier) SetAccountName(v string) {
	o.AccountName = &v
}

// GetSID returns the SID field value if set, zero value otherwise.
func (o *ModelsSecurityIdentitiesSecurityIdentityIdentifier) GetSID() string {
	if o == nil || isNil(o.SID) {
		var ret string
		return ret
	}
	return *o.SID
}

// GetSIDOk returns a tuple with the SID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecurityIdentitiesSecurityIdentityIdentifier) GetSIDOk() (*string, bool) {
	if o == nil || isNil(o.SID) {
		return nil, false
	}
	return o.SID, true
}

// HasSID returns a boolean if a field has been set.
func (o *ModelsSecurityIdentitiesSecurityIdentityIdentifier) HasSID() bool {
	if o != nil && !isNil(o.SID) {
		return true
	}

	return false
}

// SetSID gets a reference to the given string and assigns it to the SID field.
func (o *ModelsSecurityIdentitiesSecurityIdentityIdentifier) SetSID(v string) {
	o.SID = &v
}

func (o ModelsSecurityIdentitiesSecurityIdentityIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSecurityIdentitiesSecurityIdentityIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountName) {
		toSerialize["AccountName"] = o.AccountName
	}
	if !isNil(o.SID) {
		toSerialize["SID"] = o.SID
	}
	return toSerialize, nil
}

type NullableModelsSecurityIdentitiesSecurityIdentityIdentifier struct {
	value *ModelsSecurityIdentitiesSecurityIdentityIdentifier
	isSet bool
}

func (v NullableModelsSecurityIdentitiesSecurityIdentityIdentifier) Get() *ModelsSecurityIdentitiesSecurityIdentityIdentifier {
	return v.value
}

func (v *NullableModelsSecurityIdentitiesSecurityIdentityIdentifier) Set(val *ModelsSecurityIdentitiesSecurityIdentityIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSecurityIdentitiesSecurityIdentityIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSecurityIdentitiesSecurityIdentityIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSecurityIdentitiesSecurityIdentityIdentifier(val *ModelsSecurityIdentitiesSecurityIdentityIdentifier) *NullableModelsSecurityIdentitiesSecurityIdentityIdentifier {
	return &NullableModelsSecurityIdentitiesSecurityIdentityIdentifier{value: val, isSet: true}
}

func (v NullableModelsSecurityIdentitiesSecurityIdentityIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSecurityIdentitiesSecurityIdentityIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


