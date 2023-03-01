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

// checks if the ModelsCertificateImportRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertificateImportRequestModel{}

// ModelsCertificateImportRequestModel Class representing a request to import a certificate into Keyfactor Command
type ModelsCertificateImportRequestModel struct {
	// Base64-encoded certificate contents
	Certificate string `json:"Certificate"`
	// Optional password associated if required for a PFX
	Password *string `json:"Password,omitempty"`
	// Colleciton of metadata to be associated with the imported certificate
	Metadata *map[string]string `json:"Metadata,omitempty"`
	// List of the Keyfactor certificate store identifiers (GUID) with which the imported certificate should be associated
	StoreIds []string `json:"StoreIds,omitempty"`
	// List of the certificate store types with with the imported certificate should be associated
	StoreTypes []ModelsEnrollmentManagementStoreType `json:"StoreTypes,omitempty"`
	// Schedule on which the certificate should be imported
	Schedule *time.Time `json:"Schedule,omitempty"`
	// Whether or not we should validate and import the certificate's metadata.
	ImportMetadata *bool `json:"ImportMetadata,omitempty"`
}

// NewModelsCertificateImportRequestModel instantiates a new ModelsCertificateImportRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertificateImportRequestModel(certificate string) *ModelsCertificateImportRequestModel {
	this := ModelsCertificateImportRequestModel{}
	this.Certificate = certificate
	return &this
}

// NewModelsCertificateImportRequestModelWithDefaults instantiates a new ModelsCertificateImportRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertificateImportRequestModelWithDefaults() *ModelsCertificateImportRequestModel {
	this := ModelsCertificateImportRequestModel{}
	return &this
}

// GetCertificate returns the Certificate field value
func (o *ModelsCertificateImportRequestModel) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportRequestModel) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *ModelsCertificateImportRequestModel) SetCertificate(v string) {
	o.Certificate = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ModelsCertificateImportRequestModel) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportRequestModel) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ModelsCertificateImportRequestModel) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ModelsCertificateImportRequestModel) SetPassword(v string) {
	o.Password = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ModelsCertificateImportRequestModel) GetMetadata() map[string]string {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportRequestModel) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ModelsCertificateImportRequestModel) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ModelsCertificateImportRequestModel) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetStoreIds returns the StoreIds field value if set, zero value otherwise.
func (o *ModelsCertificateImportRequestModel) GetStoreIds() []string {
	if o == nil || isNil(o.StoreIds) {
		var ret []string
		return ret
	}
	return o.StoreIds
}

// GetStoreIdsOk returns a tuple with the StoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportRequestModel) GetStoreIdsOk() ([]string, bool) {
	if o == nil || isNil(o.StoreIds) {
		return nil, false
	}
	return o.StoreIds, true
}

// HasStoreIds returns a boolean if a field has been set.
func (o *ModelsCertificateImportRequestModel) HasStoreIds() bool {
	if o != nil && !isNil(o.StoreIds) {
		return true
	}

	return false
}

// SetStoreIds gets a reference to the given []string and assigns it to the StoreIds field.
func (o *ModelsCertificateImportRequestModel) SetStoreIds(v []string) {
	o.StoreIds = v
}

// GetStoreTypes returns the StoreTypes field value if set, zero value otherwise.
func (o *ModelsCertificateImportRequestModel) GetStoreTypes() []ModelsEnrollmentManagementStoreType {
	if o == nil || isNil(o.StoreTypes) {
		var ret []ModelsEnrollmentManagementStoreType
		return ret
	}
	return o.StoreTypes
}

// GetStoreTypesOk returns a tuple with the StoreTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportRequestModel) GetStoreTypesOk() ([]ModelsEnrollmentManagementStoreType, bool) {
	if o == nil || isNil(o.StoreTypes) {
		return nil, false
	}
	return o.StoreTypes, true
}

// HasStoreTypes returns a boolean if a field has been set.
func (o *ModelsCertificateImportRequestModel) HasStoreTypes() bool {
	if o != nil && !isNil(o.StoreTypes) {
		return true
	}

	return false
}

// SetStoreTypes gets a reference to the given []ModelsEnrollmentManagementStoreType and assigns it to the StoreTypes field.
func (o *ModelsCertificateImportRequestModel) SetStoreTypes(v []ModelsEnrollmentManagementStoreType) {
	o.StoreTypes = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ModelsCertificateImportRequestModel) GetSchedule() time.Time {
	if o == nil || isNil(o.Schedule) {
		var ret time.Time
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportRequestModel) GetScheduleOk() (*time.Time, bool) {
	if o == nil || isNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ModelsCertificateImportRequestModel) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given time.Time and assigns it to the Schedule field.
func (o *ModelsCertificateImportRequestModel) SetSchedule(v time.Time) {
	o.Schedule = &v
}

// GetImportMetadata returns the ImportMetadata field value if set, zero value otherwise.
func (o *ModelsCertificateImportRequestModel) GetImportMetadata() bool {
	if o == nil || isNil(o.ImportMetadata) {
		var ret bool
		return ret
	}
	return *o.ImportMetadata
}

// GetImportMetadataOk returns a tuple with the ImportMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateImportRequestModel) GetImportMetadataOk() (*bool, bool) {
	if o == nil || isNil(o.ImportMetadata) {
		return nil, false
	}
	return o.ImportMetadata, true
}

// HasImportMetadata returns a boolean if a field has been set.
func (o *ModelsCertificateImportRequestModel) HasImportMetadata() bool {
	if o != nil && !isNil(o.ImportMetadata) {
		return true
	}

	return false
}

// SetImportMetadata gets a reference to the given bool and assigns it to the ImportMetadata field.
func (o *ModelsCertificateImportRequestModel) SetImportMetadata(v bool) {
	o.ImportMetadata = &v
}

func (o ModelsCertificateImportRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertificateImportRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Certificate"] = o.Certificate
	if !isNil(o.Password) {
		toSerialize["Password"] = o.Password
	}
	if !isNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	if !isNil(o.StoreIds) {
		toSerialize["StoreIds"] = o.StoreIds
	}
	if !isNil(o.StoreTypes) {
		toSerialize["StoreTypes"] = o.StoreTypes
	}
	if !isNil(o.Schedule) {
		toSerialize["Schedule"] = o.Schedule
	}
	if !isNil(o.ImportMetadata) {
		toSerialize["ImportMetadata"] = o.ImportMetadata
	}
	return toSerialize, nil
}

type NullableModelsCertificateImportRequestModel struct {
	value *ModelsCertificateImportRequestModel
	isSet bool
}

func (v NullableModelsCertificateImportRequestModel) Get() *ModelsCertificateImportRequestModel {
	return v.value
}

func (v *NullableModelsCertificateImportRequestModel) Set(val *ModelsCertificateImportRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertificateImportRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertificateImportRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertificateImportRequestModel(val *ModelsCertificateImportRequestModel) *NullableModelsCertificateImportRequestModel {
	return &NullableModelsCertificateImportRequestModel{value: val, isSet: true}
}

func (v NullableModelsCertificateImportRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertificateImportRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


