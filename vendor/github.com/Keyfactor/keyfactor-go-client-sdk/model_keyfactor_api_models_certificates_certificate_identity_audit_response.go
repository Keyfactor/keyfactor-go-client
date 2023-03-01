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

// checks if the KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse{}

// KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse Represents an account with a list of permission granted to it on a given certificate by either a role or collection
type KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse struct {
	// Id of the account represented by the audit response
	Id *int32 `json:"Id,omitempty"`
	// Name of the account represented by the audit response
	AccountName *string `json:"AccountName,omitempty"`
	// The type of account represented by the audit response (User or Group)
	IdentityType *string `json:"IdentityType,omitempty"`
	// The SID of the account represented by the audit reponse
	SID *string `json:"SID,omitempty"`
	// Permissions granted to the account represented by the audit reponse on the specified certifcate
	Permissions []KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission `json:"Permissions,omitempty"`
}

// NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse instantiates a new KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse() *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse {
	this := KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse{}
	return &this
}

// NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseWithDefaults instantiates a new KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsCertificatesCertificateIdentityAuditResponseWithDefaults() *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse {
	this := KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetId(v int32) {
	o.Id = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetAccountName() string {
	if o == nil || isNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetAccountNameOk() (*string, bool) {
	if o == nil || isNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasAccountName() bool {
	if o != nil && !isNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetAccountName(v string) {
	o.AccountName = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetIdentityType() string {
	if o == nil || isNil(o.IdentityType) {
		var ret string
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetIdentityTypeOk() (*string, bool) {
	if o == nil || isNil(o.IdentityType) {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasIdentityType() bool {
	if o != nil && !isNil(o.IdentityType) {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given string and assigns it to the IdentityType field.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetIdentityType(v string) {
	o.IdentityType = &v
}

// GetSID returns the SID field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetSID() string {
	if o == nil || isNil(o.SID) {
		var ret string
		return ret
	}
	return *o.SID
}

// GetSIDOk returns a tuple with the SID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetSIDOk() (*string, bool) {
	if o == nil || isNil(o.SID) {
		return nil, false
	}
	return o.SID, true
}

// HasSID returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasSID() bool {
	if o != nil && !isNil(o.SID) {
		return true
	}

	return false
}

// SetSID gets a reference to the given string and assigns it to the SID field.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetSID(v string) {
	o.SID = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetPermissions() []KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission {
	if o == nil || isNil(o.Permissions) {
		var ret []KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) GetPermissionsOk() ([]KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission, bool) {
	if o == nil || isNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) HasPermissions() bool {
	if o != nil && !isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission and assigns it to the Permissions field.
func (o *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) SetPermissions(v []KeyfactorApiModelsCertificatesCertificateIdentityAuditResponseCertificatePermission) {
	o.Permissions = v
}

func (o KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.AccountName) {
		toSerialize["AccountName"] = o.AccountName
	}
	if !isNil(o.IdentityType) {
		toSerialize["IdentityType"] = o.IdentityType
	}
	if !isNil(o.SID) {
		toSerialize["SID"] = o.SID
	}
	// skip: Permissions is readOnly
	return toSerialize, nil
}

type NullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse struct {
	value *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) Get() *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) Set(val *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse(val *KeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) *NullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse {
	return &NullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsCertificatesCertificateIdentityAuditResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


