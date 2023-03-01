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

// checks if the KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse{}

// KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse struct for KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse
type KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse struct {
	// The allowed RSA key sizes.
	RSAValidKeySizes []int32 `json:"RSAValidKeySizes,omitempty"`
	// The allowed ECC curves.
	ECCValidCurves []string `json:"ECCValidCurves,omitempty"`
	// Whether or not keys can be reused.
	AllowKeyReuse *bool `json:"AllowKeyReuse,omitempty"`
	// Whether or not wildcards can be used.
	AllowWildcards *bool `json:"AllowWildcards,omitempty"`
	// Whether or not RFC 2818 compliance should be enforced.
	RFCEnforcement *bool `json:"RFCEnforcement,omitempty"`
}

// NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse instantiates a new KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse() *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse {
	this := KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse{}
	return &this
}

// NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponseWithDefaults instantiates a new KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponseWithDefaults() *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse {
	this := KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse{}
	return &this
}

// GetRSAValidKeySizes returns the RSAValidKeySizes field value if set, zero value otherwise.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetRSAValidKeySizes() []int32 {
	if o == nil || isNil(o.RSAValidKeySizes) {
		var ret []int32
		return ret
	}
	return o.RSAValidKeySizes
}

// GetRSAValidKeySizesOk returns a tuple with the RSAValidKeySizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetRSAValidKeySizesOk() ([]int32, bool) {
	if o == nil || isNil(o.RSAValidKeySizes) {
		return nil, false
	}
	return o.RSAValidKeySizes, true
}

// HasRSAValidKeySizes returns a boolean if a field has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasRSAValidKeySizes() bool {
	if o != nil && !isNil(o.RSAValidKeySizes) {
		return true
	}

	return false
}

// SetRSAValidKeySizes gets a reference to the given []int32 and assigns it to the RSAValidKeySizes field.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetRSAValidKeySizes(v []int32) {
	o.RSAValidKeySizes = v
}

// GetECCValidCurves returns the ECCValidCurves field value if set, zero value otherwise.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetECCValidCurves() []string {
	if o == nil || isNil(o.ECCValidCurves) {
		var ret []string
		return ret
	}
	return o.ECCValidCurves
}

// GetECCValidCurvesOk returns a tuple with the ECCValidCurves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetECCValidCurvesOk() ([]string, bool) {
	if o == nil || isNil(o.ECCValidCurves) {
		return nil, false
	}
	return o.ECCValidCurves, true
}

// HasECCValidCurves returns a boolean if a field has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasECCValidCurves() bool {
	if o != nil && !isNil(o.ECCValidCurves) {
		return true
	}

	return false
}

// SetECCValidCurves gets a reference to the given []string and assigns it to the ECCValidCurves field.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetECCValidCurves(v []string) {
	o.ECCValidCurves = v
}

// GetAllowKeyReuse returns the AllowKeyReuse field value if set, zero value otherwise.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetAllowKeyReuse() bool {
	if o == nil || isNil(o.AllowKeyReuse) {
		var ret bool
		return ret
	}
	return *o.AllowKeyReuse
}

// GetAllowKeyReuseOk returns a tuple with the AllowKeyReuse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetAllowKeyReuseOk() (*bool, bool) {
	if o == nil || isNil(o.AllowKeyReuse) {
		return nil, false
	}
	return o.AllowKeyReuse, true
}

// HasAllowKeyReuse returns a boolean if a field has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasAllowKeyReuse() bool {
	if o != nil && !isNil(o.AllowKeyReuse) {
		return true
	}

	return false
}

// SetAllowKeyReuse gets a reference to the given bool and assigns it to the AllowKeyReuse field.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetAllowKeyReuse(v bool) {
	o.AllowKeyReuse = &v
}

// GetAllowWildcards returns the AllowWildcards field value if set, zero value otherwise.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetAllowWildcards() bool {
	if o == nil || isNil(o.AllowWildcards) {
		var ret bool
		return ret
	}
	return *o.AllowWildcards
}

// GetAllowWildcardsOk returns a tuple with the AllowWildcards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetAllowWildcardsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowWildcards) {
		return nil, false
	}
	return o.AllowWildcards, true
}

// HasAllowWildcards returns a boolean if a field has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasAllowWildcards() bool {
	if o != nil && !isNil(o.AllowWildcards) {
		return true
	}

	return false
}

// SetAllowWildcards gets a reference to the given bool and assigns it to the AllowWildcards field.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetAllowWildcards(v bool) {
	o.AllowWildcards = &v
}

// GetRFCEnforcement returns the RFCEnforcement field value if set, zero value otherwise.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetRFCEnforcement() bool {
	if o == nil || isNil(o.RFCEnforcement) {
		var ret bool
		return ret
	}
	return *o.RFCEnforcement
}

// GetRFCEnforcementOk returns a tuple with the RFCEnforcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) GetRFCEnforcementOk() (*bool, bool) {
	if o == nil || isNil(o.RFCEnforcement) {
		return nil, false
	}
	return o.RFCEnforcement, true
}

// HasRFCEnforcement returns a boolean if a field has been set.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) HasRFCEnforcement() bool {
	if o != nil && !isNil(o.RFCEnforcement) {
		return true
	}

	return false
}

// SetRFCEnforcement gets a reference to the given bool and assigns it to the RFCEnforcement field.
func (o *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) SetRFCEnforcement(v bool) {
	o.RFCEnforcement = &v
}

func (o KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RSAValidKeySizes) {
		toSerialize["RSAValidKeySizes"] = o.RSAValidKeySizes
	}
	if !isNil(o.ECCValidCurves) {
		toSerialize["ECCValidCurves"] = o.ECCValidCurves
	}
	if !isNil(o.AllowKeyReuse) {
		toSerialize["AllowKeyReuse"] = o.AllowKeyReuse
	}
	if !isNil(o.AllowWildcards) {
		toSerialize["AllowWildcards"] = o.AllowWildcards
	}
	if !isNil(o.RFCEnforcement) {
		toSerialize["RFCEnforcement"] = o.RFCEnforcement
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse struct {
	value *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) Get() *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) Set(val *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse(val *KeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) *NullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse {
	return &NullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsTemplatesTemplateEnrollmentPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


