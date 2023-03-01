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

// checks if the ModelsSecuritySecurityRolesSecurityRoleCreationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsSecuritySecurityRolesSecurityRoleCreationRequest{}

// ModelsSecuritySecurityRolesSecurityRoleCreationRequest struct for ModelsSecuritySecurityRolesSecurityRoleCreationRequest
type ModelsSecuritySecurityRolesSecurityRoleCreationRequest struct {
	// The name of the security role to create
	Name string `json:"Name"`
	// The description to be used on the created security role
	Description string `json:"Description"`
	// Whether or not the security role should be enabled
	Enabled *bool `json:"Enabled,omitempty"`
	// Whether or not the security role should be private
	Private *bool `json:"Private,omitempty"`
	// The permissions to include in the role. These must be supplied in the format \"Area:Permission\"
	Permissions []string `json:"Permissions,omitempty"`
	// The Keyfactor identities to assign to the created role
	Identities []ModelsSecurityIdentitiesSecurityIdentityIdentifier `json:"Identities,omitempty"`
}

// NewModelsSecuritySecurityRolesSecurityRoleCreationRequest instantiates a new ModelsSecuritySecurityRolesSecurityRoleCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsSecuritySecurityRolesSecurityRoleCreationRequest(name string, description string) *ModelsSecuritySecurityRolesSecurityRoleCreationRequest {
	this := ModelsSecuritySecurityRolesSecurityRoleCreationRequest{}
	this.Name = name
	this.Description = description
	return &this
}

// NewModelsSecuritySecurityRolesSecurityRoleCreationRequestWithDefaults instantiates a new ModelsSecuritySecurityRolesSecurityRoleCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsSecuritySecurityRolesSecurityRoleCreationRequestWithDefaults() *ModelsSecuritySecurityRolesSecurityRoleCreationRequest {
	this := ModelsSecuritySecurityRolesSecurityRoleCreationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetDescription(v string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetPrivate() bool {
	if o == nil || isNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetPrivateOk() (*bool, bool) {
	if o == nil || isNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) HasPrivate() bool {
	if o != nil && !isNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetPrivate(v bool) {
	o.Private = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetPermissions() []string {
	if o == nil || isNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil || isNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) HasPermissions() bool {
	if o != nil && !isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetPermissions(v []string) {
	o.Permissions = v
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetIdentities() []ModelsSecurityIdentitiesSecurityIdentityIdentifier {
	if o == nil || isNil(o.Identities) {
		var ret []ModelsSecurityIdentitiesSecurityIdentityIdentifier
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetIdentitiesOk() ([]ModelsSecurityIdentitiesSecurityIdentityIdentifier, bool) {
	if o == nil || isNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) HasIdentities() bool {
	if o != nil && !isNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []ModelsSecurityIdentitiesSecurityIdentityIdentifier and assigns it to the Identities field.
func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetIdentities(v []ModelsSecurityIdentitiesSecurityIdentityIdentifier) {
	o.Identities = v
}

func (o ModelsSecuritySecurityRolesSecurityRoleCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsSecuritySecurityRolesSecurityRoleCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	toSerialize["Description"] = o.Description
	if !isNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !isNil(o.Private) {
		toSerialize["Private"] = o.Private
	}
	if !isNil(o.Permissions) {
		toSerialize["Permissions"] = o.Permissions
	}
	if !isNil(o.Identities) {
		toSerialize["Identities"] = o.Identities
	}
	return toSerialize, nil
}

type NullableModelsSecuritySecurityRolesSecurityRoleCreationRequest struct {
	value *ModelsSecuritySecurityRolesSecurityRoleCreationRequest
	isSet bool
}

func (v NullableModelsSecuritySecurityRolesSecurityRoleCreationRequest) Get() *ModelsSecuritySecurityRolesSecurityRoleCreationRequest {
	return v.value
}

func (v *NullableModelsSecuritySecurityRolesSecurityRoleCreationRequest) Set(val *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsSecuritySecurityRolesSecurityRoleCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsSecuritySecurityRolesSecurityRoleCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsSecuritySecurityRolesSecurityRoleCreationRequest(val *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) *NullableModelsSecuritySecurityRolesSecurityRoleCreationRequest {
	return &NullableModelsSecuritySecurityRolesSecurityRoleCreationRequest{value: val, isSet: true}
}

func (v NullableModelsSecuritySecurityRolesSecurityRoleCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsSecuritySecurityRolesSecurityRoleCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


