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

// checks if the ModelsPendingCSRResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsPendingCSRResponse{}

// ModelsPendingCSRResponse struct for ModelsPendingCSRResponse
type ModelsPendingCSRResponse struct {
	Id *int32 `json:"Id,omitempty"`
	CSR *string `json:"CSR,omitempty"`
	RequestTime *time.Time `json:"RequestTime,omitempty"`
	Subject []string `json:"Subject,omitempty"`
}

// NewModelsPendingCSRResponse instantiates a new ModelsPendingCSRResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsPendingCSRResponse() *ModelsPendingCSRResponse {
	this := ModelsPendingCSRResponse{}
	return &this
}

// NewModelsPendingCSRResponseWithDefaults instantiates a new ModelsPendingCSRResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsPendingCSRResponseWithDefaults() *ModelsPendingCSRResponse {
	this := ModelsPendingCSRResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsPendingCSRResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPendingCSRResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsPendingCSRResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelsPendingCSRResponse) SetId(v int32) {
	o.Id = &v
}

// GetCSR returns the CSR field value if set, zero value otherwise.
func (o *ModelsPendingCSRResponse) GetCSR() string {
	if o == nil || isNil(o.CSR) {
		var ret string
		return ret
	}
	return *o.CSR
}

// GetCSROk returns a tuple with the CSR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPendingCSRResponse) GetCSROk() (*string, bool) {
	if o == nil || isNil(o.CSR) {
		return nil, false
	}
	return o.CSR, true
}

// HasCSR returns a boolean if a field has been set.
func (o *ModelsPendingCSRResponse) HasCSR() bool {
	if o != nil && !isNil(o.CSR) {
		return true
	}

	return false
}

// SetCSR gets a reference to the given string and assigns it to the CSR field.
func (o *ModelsPendingCSRResponse) SetCSR(v string) {
	o.CSR = &v
}

// GetRequestTime returns the RequestTime field value if set, zero value otherwise.
func (o *ModelsPendingCSRResponse) GetRequestTime() time.Time {
	if o == nil || isNil(o.RequestTime) {
		var ret time.Time
		return ret
	}
	return *o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPendingCSRResponse) GetRequestTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.RequestTime) {
		return nil, false
	}
	return o.RequestTime, true
}

// HasRequestTime returns a boolean if a field has been set.
func (o *ModelsPendingCSRResponse) HasRequestTime() bool {
	if o != nil && !isNil(o.RequestTime) {
		return true
	}

	return false
}

// SetRequestTime gets a reference to the given time.Time and assigns it to the RequestTime field.
func (o *ModelsPendingCSRResponse) SetRequestTime(v time.Time) {
	o.RequestTime = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ModelsPendingCSRResponse) GetSubject() []string {
	if o == nil || isNil(o.Subject) {
		var ret []string
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPendingCSRResponse) GetSubjectOk() ([]string, bool) {
	if o == nil || isNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ModelsPendingCSRResponse) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given []string and assigns it to the Subject field.
func (o *ModelsPendingCSRResponse) SetSubject(v []string) {
	o.Subject = v
}

func (o ModelsPendingCSRResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsPendingCSRResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.CSR) {
		toSerialize["CSR"] = o.CSR
	}
	if !isNil(o.RequestTime) {
		toSerialize["RequestTime"] = o.RequestTime
	}
	if !isNil(o.Subject) {
		toSerialize["Subject"] = o.Subject
	}
	return toSerialize, nil
}

type NullableModelsPendingCSRResponse struct {
	value *ModelsPendingCSRResponse
	isSet bool
}

func (v NullableModelsPendingCSRResponse) Get() *ModelsPendingCSRResponse {
	return v.value
}

func (v *NullableModelsPendingCSRResponse) Set(val *ModelsPendingCSRResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsPendingCSRResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsPendingCSRResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsPendingCSRResponse(val *ModelsPendingCSRResponse) *NullableModelsPendingCSRResponse {
	return &NullableModelsPendingCSRResponse{value: val, isSet: true}
}

func (v NullableModelsPendingCSRResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsPendingCSRResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


