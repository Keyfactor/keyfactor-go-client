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

// checks if the ModelsQueryModelsPagedSSHLogonQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsQueryModelsPagedSSHLogonQuery{}

// ModelsQueryModelsPagedSSHLogonQuery struct for ModelsQueryModelsPagedSSHLogonQuery
type ModelsQueryModelsPagedSSHLogonQuery struct {
	// Contents of the query (ex: field1 -eq value1 AND field2 -gt value2)
	QueryString *string `json:"QueryString,omitempty"`
	// The current page within the result set to be returned
	PageReturned *int32 `json:"PageReturned,omitempty"`
	// Maximum number of records to be returned in a single call
	ReturnLimit *int32 `json:"ReturnLimit,omitempty"`
	// Field by which the results should be sorted (view results via Management Portal for sortable columns)
	SortField *string `json:"SortField,omitempty"`
	// Field sort direction [0=ascending, 1=descending]
	SortAscending *int32 `json:"SortAscending,omitempty"`
}

// NewModelsQueryModelsPagedSSHLogonQuery instantiates a new ModelsQueryModelsPagedSSHLogonQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsQueryModelsPagedSSHLogonQuery() *ModelsQueryModelsPagedSSHLogonQuery {
	this := ModelsQueryModelsPagedSSHLogonQuery{}
	return &this
}

// NewModelsQueryModelsPagedSSHLogonQueryWithDefaults instantiates a new ModelsQueryModelsPagedSSHLogonQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsQueryModelsPagedSSHLogonQueryWithDefaults() *ModelsQueryModelsPagedSSHLogonQuery {
	this := ModelsQueryModelsPagedSSHLogonQuery{}
	return &this
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetQueryString() string {
	if o == nil || isNil(o.QueryString) {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetQueryStringOk() (*string, bool) {
	if o == nil || isNil(o.QueryString) {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) HasQueryString() bool {
	if o != nil && !isNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *ModelsQueryModelsPagedSSHLogonQuery) SetQueryString(v string) {
	o.QueryString = &v
}

// GetPageReturned returns the PageReturned field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetPageReturned() int32 {
	if o == nil || isNil(o.PageReturned) {
		var ret int32
		return ret
	}
	return *o.PageReturned
}

// GetPageReturnedOk returns a tuple with the PageReturned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetPageReturnedOk() (*int32, bool) {
	if o == nil || isNil(o.PageReturned) {
		return nil, false
	}
	return o.PageReturned, true
}

// HasPageReturned returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) HasPageReturned() bool {
	if o != nil && !isNil(o.PageReturned) {
		return true
	}

	return false
}

// SetPageReturned gets a reference to the given int32 and assigns it to the PageReturned field.
func (o *ModelsQueryModelsPagedSSHLogonQuery) SetPageReturned(v int32) {
	o.PageReturned = &v
}

// GetReturnLimit returns the ReturnLimit field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetReturnLimit() int32 {
	if o == nil || isNil(o.ReturnLimit) {
		var ret int32
		return ret
	}
	return *o.ReturnLimit
}

// GetReturnLimitOk returns a tuple with the ReturnLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetReturnLimitOk() (*int32, bool) {
	if o == nil || isNil(o.ReturnLimit) {
		return nil, false
	}
	return o.ReturnLimit, true
}

// HasReturnLimit returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) HasReturnLimit() bool {
	if o != nil && !isNil(o.ReturnLimit) {
		return true
	}

	return false
}

// SetReturnLimit gets a reference to the given int32 and assigns it to the ReturnLimit field.
func (o *ModelsQueryModelsPagedSSHLogonQuery) SetReturnLimit(v int32) {
	o.ReturnLimit = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetSortField() string {
	if o == nil || isNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetSortFieldOk() (*string, bool) {
	if o == nil || isNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) HasSortField() bool {
	if o != nil && !isNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *ModelsQueryModelsPagedSSHLogonQuery) SetSortField(v string) {
	o.SortField = &v
}

// GetSortAscending returns the SortAscending field value if set, zero value otherwise.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetSortAscending() int32 {
	if o == nil || isNil(o.SortAscending) {
		var ret int32
		return ret
	}
	return *o.SortAscending
}

// GetSortAscendingOk returns a tuple with the SortAscending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) GetSortAscendingOk() (*int32, bool) {
	if o == nil || isNil(o.SortAscending) {
		return nil, false
	}
	return o.SortAscending, true
}

// HasSortAscending returns a boolean if a field has been set.
func (o *ModelsQueryModelsPagedSSHLogonQuery) HasSortAscending() bool {
	if o != nil && !isNil(o.SortAscending) {
		return true
	}

	return false
}

// SetSortAscending gets a reference to the given int32 and assigns it to the SortAscending field.
func (o *ModelsQueryModelsPagedSSHLogonQuery) SetSortAscending(v int32) {
	o.SortAscending = &v
}

func (o ModelsQueryModelsPagedSSHLogonQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsQueryModelsPagedSSHLogonQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.QueryString) {
		toSerialize["QueryString"] = o.QueryString
	}
	if !isNil(o.PageReturned) {
		toSerialize["PageReturned"] = o.PageReturned
	}
	if !isNil(o.ReturnLimit) {
		toSerialize["ReturnLimit"] = o.ReturnLimit
	}
	if !isNil(o.SortField) {
		toSerialize["SortField"] = o.SortField
	}
	if !isNil(o.SortAscending) {
		toSerialize["SortAscending"] = o.SortAscending
	}
	return toSerialize, nil
}

type NullableModelsQueryModelsPagedSSHLogonQuery struct {
	value *ModelsQueryModelsPagedSSHLogonQuery
	isSet bool
}

func (v NullableModelsQueryModelsPagedSSHLogonQuery) Get() *ModelsQueryModelsPagedSSHLogonQuery {
	return v.value
}

func (v *NullableModelsQueryModelsPagedSSHLogonQuery) Set(val *ModelsQueryModelsPagedSSHLogonQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsQueryModelsPagedSSHLogonQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsQueryModelsPagedSSHLogonQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsQueryModelsPagedSSHLogonQuery(val *ModelsQueryModelsPagedSSHLogonQuery) *NullableModelsQueryModelsPagedSSHLogonQuery {
	return &NullableModelsQueryModelsPagedSSHLogonQuery{value: val, isSet: true}
}

func (v NullableModelsQueryModelsPagedSSHLogonQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsQueryModelsPagedSSHLogonQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


