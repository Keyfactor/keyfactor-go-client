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

// checks if the ModelsCertificateDownloadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCertificateDownloadRequest{}

// ModelsCertificateDownloadRequest struct for ModelsCertificateDownloadRequest
type ModelsCertificateDownloadRequest struct {
	CertID *int32 `json:"CertID,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty"`
	IssuerDN NullableString `json:"IssuerDN,omitempty"`
	Thumbprint *string `json:"Thumbprint,omitempty"`
	IncludeChain *bool `json:"IncludeChain,omitempty"`
}

// NewModelsCertificateDownloadRequest instantiates a new ModelsCertificateDownloadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCertificateDownloadRequest() *ModelsCertificateDownloadRequest {
	this := ModelsCertificateDownloadRequest{}
	return &this
}

// NewModelsCertificateDownloadRequestWithDefaults instantiates a new ModelsCertificateDownloadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCertificateDownloadRequestWithDefaults() *ModelsCertificateDownloadRequest {
	this := ModelsCertificateDownloadRequest{}
	return &this
}

// GetCertID returns the CertID field value if set, zero value otherwise.
func (o *ModelsCertificateDownloadRequest) GetCertID() int32 {
	if o == nil || isNil(o.CertID) {
		var ret int32
		return ret
	}
	return *o.CertID
}

// GetCertIDOk returns a tuple with the CertID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateDownloadRequest) GetCertIDOk() (*int32, bool) {
	if o == nil || isNil(o.CertID) {
		return nil, false
	}
	return o.CertID, true
}

// HasCertID returns a boolean if a field has been set.
func (o *ModelsCertificateDownloadRequest) HasCertID() bool {
	if o != nil && !isNil(o.CertID) {
		return true
	}

	return false
}

// SetCertID gets a reference to the given int32 and assigns it to the CertID field.
func (o *ModelsCertificateDownloadRequest) SetCertID(v int32) {
	o.CertID = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *ModelsCertificateDownloadRequest) GetSerialNumber() string {
	if o == nil || isNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateDownloadRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil || isNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *ModelsCertificateDownloadRequest) HasSerialNumber() bool {
	if o != nil && !isNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *ModelsCertificateDownloadRequest) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelsCertificateDownloadRequest) GetIssuerDN() string {
	if o == nil || isNil(o.IssuerDN.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerDN.Get()
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelsCertificateDownloadRequest) GetIssuerDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerDN.Get(), o.IssuerDN.IsSet()
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *ModelsCertificateDownloadRequest) HasIssuerDN() bool {
	if o != nil && o.IssuerDN.IsSet() {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given NullableString and assigns it to the IssuerDN field.
func (o *ModelsCertificateDownloadRequest) SetIssuerDN(v string) {
	o.IssuerDN.Set(&v)
}
// SetIssuerDNNil sets the value for IssuerDN to be an explicit nil
func (o *ModelsCertificateDownloadRequest) SetIssuerDNNil() {
	o.IssuerDN.Set(nil)
}

// UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
func (o *ModelsCertificateDownloadRequest) UnsetIssuerDN() {
	o.IssuerDN.Unset()
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *ModelsCertificateDownloadRequest) GetThumbprint() string {
	if o == nil || isNil(o.Thumbprint) {
		var ret string
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateDownloadRequest) GetThumbprintOk() (*string, bool) {
	if o == nil || isNil(o.Thumbprint) {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *ModelsCertificateDownloadRequest) HasThumbprint() bool {
	if o != nil && !isNil(o.Thumbprint) {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given string and assigns it to the Thumbprint field.
func (o *ModelsCertificateDownloadRequest) SetThumbprint(v string) {
	o.Thumbprint = &v
}

// GetIncludeChain returns the IncludeChain field value if set, zero value otherwise.
func (o *ModelsCertificateDownloadRequest) GetIncludeChain() bool {
	if o == nil || isNil(o.IncludeChain) {
		var ret bool
		return ret
	}
	return *o.IncludeChain
}

// GetIncludeChainOk returns a tuple with the IncludeChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCertificateDownloadRequest) GetIncludeChainOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeChain) {
		return nil, false
	}
	return o.IncludeChain, true
}

// HasIncludeChain returns a boolean if a field has been set.
func (o *ModelsCertificateDownloadRequest) HasIncludeChain() bool {
	if o != nil && !isNil(o.IncludeChain) {
		return true
	}

	return false
}

// SetIncludeChain gets a reference to the given bool and assigns it to the IncludeChain field.
func (o *ModelsCertificateDownloadRequest) SetIncludeChain(v bool) {
	o.IncludeChain = &v
}

func (o ModelsCertificateDownloadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCertificateDownloadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CertID) {
		toSerialize["CertID"] = o.CertID
	}
	if !isNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.IssuerDN.IsSet() {
		toSerialize["IssuerDN"] = o.IssuerDN.Get()
	}
	if !isNil(o.Thumbprint) {
		toSerialize["Thumbprint"] = o.Thumbprint
	}
	if !isNil(o.IncludeChain) {
		toSerialize["IncludeChain"] = o.IncludeChain
	}
	return toSerialize, nil
}

type NullableModelsCertificateDownloadRequest struct {
	value *ModelsCertificateDownloadRequest
	isSet bool
}

func (v NullableModelsCertificateDownloadRequest) Get() *ModelsCertificateDownloadRequest {
	return v.value
}

func (v *NullableModelsCertificateDownloadRequest) Set(val *ModelsCertificateDownloadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCertificateDownloadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCertificateDownloadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCertificateDownloadRequest(val *ModelsCertificateDownloadRequest) *NullableModelsCertificateDownloadRequest {
	return &NullableModelsCertificateDownloadRequest{value: val, isSet: true}
}

func (v NullableModelsCertificateDownloadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCertificateDownloadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


