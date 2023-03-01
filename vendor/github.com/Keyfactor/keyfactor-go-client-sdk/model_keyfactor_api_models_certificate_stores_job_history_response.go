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

// checks if the KeyfactorApiModelsCertificateStoresJobHistoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsCertificateStoresJobHistoryResponse{}

// KeyfactorApiModelsCertificateStoresJobHistoryResponse struct for KeyfactorApiModelsCertificateStoresJobHistoryResponse
type KeyfactorApiModelsCertificateStoresJobHistoryResponse struct {
	JobHistoryId *int64 `json:"JobHistoryId,omitempty"`
	AgentMachine *string `json:"AgentMachine,omitempty"`
	JobId *string `json:"JobId,omitempty"`
	Schedule *KeyfactorCommonSchedulingKeyfactorSchedule `json:"Schedule,omitempty"`
	JobType *string `json:"JobType,omitempty"`
	OperationStart *time.Time `json:"OperationStart,omitempty"`
	OperationEnd *time.Time `json:"OperationEnd,omitempty"`
	Message *string `json:"Message,omitempty"`
	Result *int32 `json:"Result,omitempty"`
	Status *int32 `json:"Status,omitempty"`
	StorePath *string `json:"StorePath,omitempty"`
	ClientMachine *string `json:"ClientMachine,omitempty"`
}

// NewKeyfactorApiModelsCertificateStoresJobHistoryResponse instantiates a new KeyfactorApiModelsCertificateStoresJobHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsCertificateStoresJobHistoryResponse() *KeyfactorApiModelsCertificateStoresJobHistoryResponse {
	this := KeyfactorApiModelsCertificateStoresJobHistoryResponse{}
	return &this
}

// NewKeyfactorApiModelsCertificateStoresJobHistoryResponseWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresJobHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsCertificateStoresJobHistoryResponseWithDefaults() *KeyfactorApiModelsCertificateStoresJobHistoryResponse {
	this := KeyfactorApiModelsCertificateStoresJobHistoryResponse{}
	return &this
}

// GetJobHistoryId returns the JobHistoryId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobHistoryId() int64 {
	if o == nil || isNil(o.JobHistoryId) {
		var ret int64
		return ret
	}
	return *o.JobHistoryId
}

// GetJobHistoryIdOk returns a tuple with the JobHistoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobHistoryIdOk() (*int64, bool) {
	if o == nil || isNil(o.JobHistoryId) {
		return nil, false
	}
	return o.JobHistoryId, true
}

// HasJobHistoryId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasJobHistoryId() bool {
	if o != nil && !isNil(o.JobHistoryId) {
		return true
	}

	return false
}

// SetJobHistoryId gets a reference to the given int64 and assigns it to the JobHistoryId field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobHistoryId(v int64) {
	o.JobHistoryId = &v
}

// GetAgentMachine returns the AgentMachine field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetAgentMachine() string {
	if o == nil || isNil(o.AgentMachine) {
		var ret string
		return ret
	}
	return *o.AgentMachine
}

// GetAgentMachineOk returns a tuple with the AgentMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetAgentMachineOk() (*string, bool) {
	if o == nil || isNil(o.AgentMachine) {
		return nil, false
	}
	return o.AgentMachine, true
}

// HasAgentMachine returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasAgentMachine() bool {
	if o != nil && !isNil(o.AgentMachine) {
		return true
	}

	return false
}

// SetAgentMachine gets a reference to the given string and assigns it to the AgentMachine field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetAgentMachine(v string) {
	o.AgentMachine = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobId() string {
	if o == nil || isNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobIdOk() (*string, bool) {
	if o == nil || isNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasJobId() bool {
	if o != nil && !isNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobId(v string) {
	o.JobId = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || isNil(o.Schedule) {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool) {
	if o == nil || isNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given KeyfactorCommonSchedulingKeyfactorSchedule and assigns it to the Schedule field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.Schedule = &v
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobType() string {
	if o == nil || isNil(o.JobType) {
		var ret string
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetJobTypeOk() (*string, bool) {
	if o == nil || isNil(o.JobType) {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasJobType() bool {
	if o != nil && !isNil(o.JobType) {
		return true
	}

	return false
}

// SetJobType gets a reference to the given string and assigns it to the JobType field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetJobType(v string) {
	o.JobType = &v
}

// GetOperationStart returns the OperationStart field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationStart() time.Time {
	if o == nil || isNil(o.OperationStart) {
		var ret time.Time
		return ret
	}
	return *o.OperationStart
}

// GetOperationStartOk returns a tuple with the OperationStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationStartOk() (*time.Time, bool) {
	if o == nil || isNil(o.OperationStart) {
		return nil, false
	}
	return o.OperationStart, true
}

// HasOperationStart returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasOperationStart() bool {
	if o != nil && !isNil(o.OperationStart) {
		return true
	}

	return false
}

// SetOperationStart gets a reference to the given time.Time and assigns it to the OperationStart field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetOperationStart(v time.Time) {
	o.OperationStart = &v
}

// GetOperationEnd returns the OperationEnd field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationEnd() time.Time {
	if o == nil || isNil(o.OperationEnd) {
		var ret time.Time
		return ret
	}
	return *o.OperationEnd
}

// GetOperationEndOk returns a tuple with the OperationEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetOperationEndOk() (*time.Time, bool) {
	if o == nil || isNil(o.OperationEnd) {
		return nil, false
	}
	return o.OperationEnd, true
}

// HasOperationEnd returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasOperationEnd() bool {
	if o != nil && !isNil(o.OperationEnd) {
		return true
	}

	return false
}

// SetOperationEnd gets a reference to the given time.Time and assigns it to the OperationEnd field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetOperationEnd(v time.Time) {
	o.OperationEnd = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetMessage(v string) {
	o.Message = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetResult() int32 {
	if o == nil || isNil(o.Result) {
		var ret int32
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetResultOk() (*int32, bool) {
	if o == nil || isNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given int32 and assigns it to the Result field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetResult(v int32) {
	o.Result = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetStorePath returns the StorePath field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStorePath() string {
	if o == nil || isNil(o.StorePath) {
		var ret string
		return ret
	}
	return *o.StorePath
}

// GetStorePathOk returns a tuple with the StorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetStorePathOk() (*string, bool) {
	if o == nil || isNil(o.StorePath) {
		return nil, false
	}
	return o.StorePath, true
}

// HasStorePath returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasStorePath() bool {
	if o != nil && !isNil(o.StorePath) {
		return true
	}

	return false
}

// SetStorePath gets a reference to the given string and assigns it to the StorePath field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetStorePath(v string) {
	o.StorePath = &v
}

// GetClientMachine returns the ClientMachine field value if set, zero value otherwise.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetClientMachine() string {
	if o == nil || isNil(o.ClientMachine) {
		var ret string
		return ret
	}
	return *o.ClientMachine
}

// GetClientMachineOk returns a tuple with the ClientMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) GetClientMachineOk() (*string, bool) {
	if o == nil || isNil(o.ClientMachine) {
		return nil, false
	}
	return o.ClientMachine, true
}

// HasClientMachine returns a boolean if a field has been set.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) HasClientMachine() bool {
	if o != nil && !isNil(o.ClientMachine) {
		return true
	}

	return false
}

// SetClientMachine gets a reference to the given string and assigns it to the ClientMachine field.
func (o *KeyfactorApiModelsCertificateStoresJobHistoryResponse) SetClientMachine(v string) {
	o.ClientMachine = &v
}

func (o KeyfactorApiModelsCertificateStoresJobHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsCertificateStoresJobHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JobHistoryId) {
		toSerialize["JobHistoryId"] = o.JobHistoryId
	}
	if !isNil(o.AgentMachine) {
		toSerialize["AgentMachine"] = o.AgentMachine
	}
	if !isNil(o.JobId) {
		toSerialize["JobId"] = o.JobId
	}
	if !isNil(o.Schedule) {
		toSerialize["Schedule"] = o.Schedule
	}
	if !isNil(o.JobType) {
		toSerialize["JobType"] = o.JobType
	}
	if !isNil(o.OperationStart) {
		toSerialize["OperationStart"] = o.OperationStart
	}
	if !isNil(o.OperationEnd) {
		toSerialize["OperationEnd"] = o.OperationEnd
	}
	if !isNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !isNil(o.Result) {
		toSerialize["Result"] = o.Result
	}
	if !isNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !isNil(o.StorePath) {
		toSerialize["StorePath"] = o.StorePath
	}
	if !isNil(o.ClientMachine) {
		toSerialize["ClientMachine"] = o.ClientMachine
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsCertificateStoresJobHistoryResponse struct {
	value *KeyfactorApiModelsCertificateStoresJobHistoryResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsCertificateStoresJobHistoryResponse) Get() *KeyfactorApiModelsCertificateStoresJobHistoryResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsCertificateStoresJobHistoryResponse) Set(val *KeyfactorApiModelsCertificateStoresJobHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsCertificateStoresJobHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsCertificateStoresJobHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsCertificateStoresJobHistoryResponse(val *KeyfactorApiModelsCertificateStoresJobHistoryResponse) *NullableKeyfactorApiModelsCertificateStoresJobHistoryResponse {
	return &NullableKeyfactorApiModelsCertificateStoresJobHistoryResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsCertificateStoresJobHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsCertificateStoresJobHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


