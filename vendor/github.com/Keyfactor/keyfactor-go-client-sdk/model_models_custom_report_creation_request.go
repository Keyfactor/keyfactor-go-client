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

// checks if the ModelsCustomReportCreationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsCustomReportCreationRequest{}

// ModelsCustomReportCreationRequest The CustomReport can be used to create and update a custom report.
type ModelsCustomReportCreationRequest struct {
	CustomURL string `json:"CustomURL"`
	DisplayName string `json:"DisplayName"`
	Description string `json:"Description"`
	InNavigator *bool `json:"InNavigator,omitempty"`
	Favorite *bool `json:"Favorite,omitempty"`
}

// NewModelsCustomReportCreationRequest instantiates a new ModelsCustomReportCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsCustomReportCreationRequest(customURL string, displayName string, description string) *ModelsCustomReportCreationRequest {
	this := ModelsCustomReportCreationRequest{}
	this.CustomURL = customURL
	this.DisplayName = displayName
	this.Description = description
	return &this
}

// NewModelsCustomReportCreationRequestWithDefaults instantiates a new ModelsCustomReportCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsCustomReportCreationRequestWithDefaults() *ModelsCustomReportCreationRequest {
	this := ModelsCustomReportCreationRequest{}
	return &this
}

// GetCustomURL returns the CustomURL field value
func (o *ModelsCustomReportCreationRequest) GetCustomURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomURL
}

// GetCustomURLOk returns a tuple with the CustomURL field value
// and a boolean to check if the value has been set.
func (o *ModelsCustomReportCreationRequest) GetCustomURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomURL, true
}

// SetCustomURL sets field value
func (o *ModelsCustomReportCreationRequest) SetCustomURL(v string) {
	o.CustomURL = v
}

// GetDisplayName returns the DisplayName field value
func (o *ModelsCustomReportCreationRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ModelsCustomReportCreationRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ModelsCustomReportCreationRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value
func (o *ModelsCustomReportCreationRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelsCustomReportCreationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModelsCustomReportCreationRequest) SetDescription(v string) {
	o.Description = v
}

// GetInNavigator returns the InNavigator field value if set, zero value otherwise.
func (o *ModelsCustomReportCreationRequest) GetInNavigator() bool {
	if o == nil || isNil(o.InNavigator) {
		var ret bool
		return ret
	}
	return *o.InNavigator
}

// GetInNavigatorOk returns a tuple with the InNavigator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCustomReportCreationRequest) GetInNavigatorOk() (*bool, bool) {
	if o == nil || isNil(o.InNavigator) {
		return nil, false
	}
	return o.InNavigator, true
}

// HasInNavigator returns a boolean if a field has been set.
func (o *ModelsCustomReportCreationRequest) HasInNavigator() bool {
	if o != nil && !isNil(o.InNavigator) {
		return true
	}

	return false
}

// SetInNavigator gets a reference to the given bool and assigns it to the InNavigator field.
func (o *ModelsCustomReportCreationRequest) SetInNavigator(v bool) {
	o.InNavigator = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *ModelsCustomReportCreationRequest) GetFavorite() bool {
	if o == nil || isNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsCustomReportCreationRequest) GetFavoriteOk() (*bool, bool) {
	if o == nil || isNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *ModelsCustomReportCreationRequest) HasFavorite() bool {
	if o != nil && !isNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *ModelsCustomReportCreationRequest) SetFavorite(v bool) {
	o.Favorite = &v
}

func (o ModelsCustomReportCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsCustomReportCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CustomURL"] = o.CustomURL
	toSerialize["DisplayName"] = o.DisplayName
	toSerialize["Description"] = o.Description
	if !isNil(o.InNavigator) {
		toSerialize["InNavigator"] = o.InNavigator
	}
	if !isNil(o.Favorite) {
		toSerialize["Favorite"] = o.Favorite
	}
	return toSerialize, nil
}

type NullableModelsCustomReportCreationRequest struct {
	value *ModelsCustomReportCreationRequest
	isSet bool
}

func (v NullableModelsCustomReportCreationRequest) Get() *ModelsCustomReportCreationRequest {
	return v.value
}

func (v *NullableModelsCustomReportCreationRequest) Set(val *ModelsCustomReportCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsCustomReportCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsCustomReportCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsCustomReportCreationRequest(val *ModelsCustomReportCreationRequest) *NullableModelsCustomReportCreationRequest {
	return &NullableModelsCustomReportCreationRequest{value: val, isSet: true}
}

func (v NullableModelsCustomReportCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsCustomReportCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


