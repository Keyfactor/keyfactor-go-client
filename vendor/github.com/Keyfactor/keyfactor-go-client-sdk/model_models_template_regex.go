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

// checks if the ModelsTemplateRegex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsTemplateRegex{}

// ModelsTemplateRegex struct for ModelsTemplateRegex
type ModelsTemplateRegex struct {
	TemplateId *int32 `json:"TemplateId,omitempty"`
	SubjectPart *string `json:"SubjectPart,omitempty"`
	Regex *string `json:"Regex,omitempty"`
	Error *string `json:"Error,omitempty"`
}

// NewModelsTemplateRegex instantiates a new ModelsTemplateRegex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsTemplateRegex() *ModelsTemplateRegex {
	this := ModelsTemplateRegex{}
	return &this
}

// NewModelsTemplateRegexWithDefaults instantiates a new ModelsTemplateRegex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsTemplateRegexWithDefaults() *ModelsTemplateRegex {
	this := ModelsTemplateRegex{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ModelsTemplateRegex) GetTemplateId() int32 {
	if o == nil || isNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplateRegex) GetTemplateIdOk() (*int32, bool) {
	if o == nil || isNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ModelsTemplateRegex) HasTemplateId() bool {
	if o != nil && !isNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *ModelsTemplateRegex) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetSubjectPart returns the SubjectPart field value if set, zero value otherwise.
func (o *ModelsTemplateRegex) GetSubjectPart() string {
	if o == nil || isNil(o.SubjectPart) {
		var ret string
		return ret
	}
	return *o.SubjectPart
}

// GetSubjectPartOk returns a tuple with the SubjectPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplateRegex) GetSubjectPartOk() (*string, bool) {
	if o == nil || isNil(o.SubjectPart) {
		return nil, false
	}
	return o.SubjectPart, true
}

// HasSubjectPart returns a boolean if a field has been set.
func (o *ModelsTemplateRegex) HasSubjectPart() bool {
	if o != nil && !isNil(o.SubjectPart) {
		return true
	}

	return false
}

// SetSubjectPart gets a reference to the given string and assigns it to the SubjectPart field.
func (o *ModelsTemplateRegex) SetSubjectPart(v string) {
	o.SubjectPart = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *ModelsTemplateRegex) GetRegex() string {
	if o == nil || isNil(o.Regex) {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplateRegex) GetRegexOk() (*string, bool) {
	if o == nil || isNil(o.Regex) {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *ModelsTemplateRegex) HasRegex() bool {
	if o != nil && !isNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *ModelsTemplateRegex) SetRegex(v string) {
	o.Regex = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ModelsTemplateRegex) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsTemplateRegex) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ModelsTemplateRegex) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ModelsTemplateRegex) SetError(v string) {
	o.Error = &v
}

func (o ModelsTemplateRegex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsTemplateRegex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TemplateId) {
		toSerialize["TemplateId"] = o.TemplateId
	}
	if !isNil(o.SubjectPart) {
		toSerialize["SubjectPart"] = o.SubjectPart
	}
	if !isNil(o.Regex) {
		toSerialize["Regex"] = o.Regex
	}
	if !isNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	return toSerialize, nil
}

type NullableModelsTemplateRegex struct {
	value *ModelsTemplateRegex
	isSet bool
}

func (v NullableModelsTemplateRegex) Get() *ModelsTemplateRegex {
	return v.value
}

func (v *NullableModelsTemplateRegex) Set(val *ModelsTemplateRegex) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsTemplateRegex) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsTemplateRegex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsTemplateRegex(val *ModelsTemplateRegex) *NullableModelsTemplateRegex {
	return &NullableModelsTemplateRegex{value: val, isSet: true}
}

func (v NullableModelsTemplateRegex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsTemplateRegex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


