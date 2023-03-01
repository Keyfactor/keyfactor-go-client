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

// checks if the KeyfactorApiModelsWorkflowsInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsWorkflowsInstanceResponse{}

// KeyfactorApiModelsWorkflowsInstanceResponse struct for KeyfactorApiModelsWorkflowsInstanceResponse
type KeyfactorApiModelsWorkflowsInstanceResponse struct {
	Id *string `json:"Id,omitempty"`
	Status *int32 `json:"Status,omitempty"`
	CurrentStepId *string `json:"CurrentStepId,omitempty"`
	StatusMessage *string `json:"StatusMessage,omitempty"`
	Signals []KeyfactorApiModelsWorkflowsAvailableSignalResponse `json:"Signals,omitempty"`
	Definition *KeyfactorApiModelsWorkflowsInstanceDefinitionResponse `json:"Definition,omitempty"`
	CurrentStepDisplayName *string `json:"CurrentStepDisplayName,omitempty"`
	CurrentStepUniqueName *string `json:"CurrentStepUniqueName,omitempty"`
	Title *string `json:"Title,omitempty"`
	LastModified *time.Time `json:"LastModified,omitempty"`
	StartDate *time.Time `json:"StartDate,omitempty"`
	InitialData map[string]map[string]interface{} `json:"InitialData,omitempty"`
	CurrentStateData map[string]map[string]interface{} `json:"CurrentStateData,omitempty"`
	ReferenceId *int64 `json:"ReferenceId,omitempty"`
}

// NewKeyfactorApiModelsWorkflowsInstanceResponse instantiates a new KeyfactorApiModelsWorkflowsInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsWorkflowsInstanceResponse() *KeyfactorApiModelsWorkflowsInstanceResponse {
	this := KeyfactorApiModelsWorkflowsInstanceResponse{}
	return &this
}

// NewKeyfactorApiModelsWorkflowsInstanceResponseWithDefaults instantiates a new KeyfactorApiModelsWorkflowsInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsWorkflowsInstanceResponseWithDefaults() *KeyfactorApiModelsWorkflowsInstanceResponse {
	this := KeyfactorApiModelsWorkflowsInstanceResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetCurrentStepId returns the CurrentStepId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepId() string {
	if o == nil || isNil(o.CurrentStepId) {
		var ret string
		return ret
	}
	return *o.CurrentStepId
}

// GetCurrentStepIdOk returns a tuple with the CurrentStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepIdOk() (*string, bool) {
	if o == nil || isNil(o.CurrentStepId) {
		return nil, false
	}
	return o.CurrentStepId, true
}

// HasCurrentStepId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasCurrentStepId() bool {
	if o != nil && !isNil(o.CurrentStepId) {
		return true
	}

	return false
}

// SetCurrentStepId gets a reference to the given string and assigns it to the CurrentStepId field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetCurrentStepId(v string) {
	o.CurrentStepId = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStatusMessage() string {
	if o == nil || isNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStatusMessageOk() (*string, bool) {
	if o == nil || isNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasStatusMessage() bool {
	if o != nil && !isNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetSignals returns the Signals field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetSignals() []KeyfactorApiModelsWorkflowsAvailableSignalResponse {
	if o == nil || isNil(o.Signals) {
		var ret []KeyfactorApiModelsWorkflowsAvailableSignalResponse
		return ret
	}
	return o.Signals
}

// GetSignalsOk returns a tuple with the Signals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetSignalsOk() ([]KeyfactorApiModelsWorkflowsAvailableSignalResponse, bool) {
	if o == nil || isNil(o.Signals) {
		return nil, false
	}
	return o.Signals, true
}

// HasSignals returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasSignals() bool {
	if o != nil && !isNil(o.Signals) {
		return true
	}

	return false
}

// SetSignals gets a reference to the given []KeyfactorApiModelsWorkflowsAvailableSignalResponse and assigns it to the Signals field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetSignals(v []KeyfactorApiModelsWorkflowsAvailableSignalResponse) {
	o.Signals = v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetDefinition() KeyfactorApiModelsWorkflowsInstanceDefinitionResponse {
	if o == nil || isNil(o.Definition) {
		var ret KeyfactorApiModelsWorkflowsInstanceDefinitionResponse
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetDefinitionOk() (*KeyfactorApiModelsWorkflowsInstanceDefinitionResponse, bool) {
	if o == nil || isNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasDefinition() bool {
	if o != nil && !isNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given KeyfactorApiModelsWorkflowsInstanceDefinitionResponse and assigns it to the Definition field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetDefinition(v KeyfactorApiModelsWorkflowsInstanceDefinitionResponse) {
	o.Definition = &v
}

// GetCurrentStepDisplayName returns the CurrentStepDisplayName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepDisplayName() string {
	if o == nil || isNil(o.CurrentStepDisplayName) {
		var ret string
		return ret
	}
	return *o.CurrentStepDisplayName
}

// GetCurrentStepDisplayNameOk returns a tuple with the CurrentStepDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.CurrentStepDisplayName) {
		return nil, false
	}
	return o.CurrentStepDisplayName, true
}

// HasCurrentStepDisplayName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasCurrentStepDisplayName() bool {
	if o != nil && !isNil(o.CurrentStepDisplayName) {
		return true
	}

	return false
}

// SetCurrentStepDisplayName gets a reference to the given string and assigns it to the CurrentStepDisplayName field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetCurrentStepDisplayName(v string) {
	o.CurrentStepDisplayName = &v
}

// GetCurrentStepUniqueName returns the CurrentStepUniqueName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepUniqueName() string {
	if o == nil || isNil(o.CurrentStepUniqueName) {
		var ret string
		return ret
	}
	return *o.CurrentStepUniqueName
}

// GetCurrentStepUniqueNameOk returns a tuple with the CurrentStepUniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStepUniqueNameOk() (*string, bool) {
	if o == nil || isNil(o.CurrentStepUniqueName) {
		return nil, false
	}
	return o.CurrentStepUniqueName, true
}

// HasCurrentStepUniqueName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasCurrentStepUniqueName() bool {
	if o != nil && !isNil(o.CurrentStepUniqueName) {
		return true
	}

	return false
}

// SetCurrentStepUniqueName gets a reference to the given string and assigns it to the CurrentStepUniqueName field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetCurrentStepUniqueName(v string) {
	o.CurrentStepUniqueName = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetTitle(v string) {
	o.Title = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetLastModified() time.Time {
	if o == nil || isNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasLastModified() bool {
	if o != nil && !isNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStartDate() time.Time {
	if o == nil || isNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetInitialData returns the InitialData field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetInitialData() map[string]map[string]interface{} {
	if o == nil || isNil(o.InitialData) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.InitialData
}

// GetInitialDataOk returns a tuple with the InitialData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetInitialDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.InitialData) {
		return map[string]map[string]interface{}{}, false
	}
	return o.InitialData, true
}

// HasInitialData returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasInitialData() bool {
	if o != nil && !isNil(o.InitialData) {
		return true
	}

	return false
}

// SetInitialData gets a reference to the given map[string]map[string]interface{} and assigns it to the InitialData field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetInitialData(v map[string]map[string]interface{}) {
	o.InitialData = v
}

// GetCurrentStateData returns the CurrentStateData field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStateData() map[string]map[string]interface{} {
	if o == nil || isNil(o.CurrentStateData) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.CurrentStateData
}

// GetCurrentStateDataOk returns a tuple with the CurrentStateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetCurrentStateDataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || isNil(o.CurrentStateData) {
		return map[string]map[string]interface{}{}, false
	}
	return o.CurrentStateData, true
}

// HasCurrentStateData returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasCurrentStateData() bool {
	if o != nil && !isNil(o.CurrentStateData) {
		return true
	}

	return false
}

// SetCurrentStateData gets a reference to the given map[string]map[string]interface{} and assigns it to the CurrentStateData field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetCurrentStateData(v map[string]map[string]interface{}) {
	o.CurrentStateData = v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetReferenceId() int64 {
	if o == nil || isNil(o.ReferenceId) {
		var ret int64
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) GetReferenceIdOk() (*int64, bool) {
	if o == nil || isNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given int64 and assigns it to the ReferenceId field.
func (o *KeyfactorApiModelsWorkflowsInstanceResponse) SetReferenceId(v int64) {
	o.ReferenceId = &v
}

func (o KeyfactorApiModelsWorkflowsInstanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsWorkflowsInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !isNil(o.CurrentStepId) {
		toSerialize["CurrentStepId"] = o.CurrentStepId
	}
	if !isNil(o.StatusMessage) {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if !isNil(o.Signals) {
		toSerialize["Signals"] = o.Signals
	}
	if !isNil(o.Definition) {
		toSerialize["Definition"] = o.Definition
	}
	if !isNil(o.CurrentStepDisplayName) {
		toSerialize["CurrentStepDisplayName"] = o.CurrentStepDisplayName
	}
	if !isNil(o.CurrentStepUniqueName) {
		toSerialize["CurrentStepUniqueName"] = o.CurrentStepUniqueName
	}
	if !isNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !isNil(o.LastModified) {
		toSerialize["LastModified"] = o.LastModified
	}
	if !isNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !isNil(o.InitialData) {
		toSerialize["InitialData"] = o.InitialData
	}
	if !isNil(o.CurrentStateData) {
		toSerialize["CurrentStateData"] = o.CurrentStateData
	}
	if !isNil(o.ReferenceId) {
		toSerialize["ReferenceId"] = o.ReferenceId
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsWorkflowsInstanceResponse struct {
	value *KeyfactorApiModelsWorkflowsInstanceResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsWorkflowsInstanceResponse) Get() *KeyfactorApiModelsWorkflowsInstanceResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsWorkflowsInstanceResponse) Set(val *KeyfactorApiModelsWorkflowsInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsWorkflowsInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsWorkflowsInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsWorkflowsInstanceResponse(val *KeyfactorApiModelsWorkflowsInstanceResponse) *NullableKeyfactorApiModelsWorkflowsInstanceResponse {
	return &NullableKeyfactorApiModelsWorkflowsInstanceResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsWorkflowsInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsWorkflowsInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


