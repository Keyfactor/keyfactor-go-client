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

// checks if the KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse{}

// KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse struct for KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
type KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse struct {
	Id *int32 `json:"Id,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty"`
	Subject *string `json:"Subject,omitempty"`
	Message *string `json:"Message,omitempty"`
	Recipients []string `json:"Recipients,omitempty"`
	Template *KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse `json:"Template,omitempty"`
	RegisteredEventHandler *KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse `json:"RegisteredEventHandler,omitempty"`
	EventHandlerParameters []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse `json:"EventHandlerParameters,omitempty"`
}

// NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse instantiates a new KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse() *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse {
	this := KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse{}
	return &this
}

// NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponseWithDefaults instantiates a new KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponseWithDefaults() *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse {
	this := KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetId(v int32) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetMessage(v string) {
	o.Message = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetRecipients() []string {
	if o == nil || isNil(o.Recipients) {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetRecipientsOk() ([]string, bool) {
	if o == nil || isNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasRecipients() bool {
	if o != nil && !isNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetRecipients(v []string) {
	o.Recipients = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetTemplate() KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse {
	if o == nil || isNil(o.Template) {
		var ret KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetTemplateOk() (*KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse, bool) {
	if o == nil || isNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasTemplate() bool {
	if o != nil && !isNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse and assigns it to the Template field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetTemplate(v KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse) {
	o.Template = &v
}

// GetRegisteredEventHandler returns the RegisteredEventHandler field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse {
	if o == nil || isNil(o.RegisteredEventHandler) {
		var ret KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse
		return ret
	}
	return *o.RegisteredEventHandler
}

// GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse, bool) {
	if o == nil || isNil(o.RegisteredEventHandler) {
		return nil, false
	}
	return o.RegisteredEventHandler, true
}

// HasRegisteredEventHandler returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasRegisteredEventHandler() bool {
	if o != nil && !isNil(o.RegisteredEventHandler) {
		return true
	}

	return false
}

// SetRegisteredEventHandler gets a reference to the given KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse and assigns it to the RegisteredEventHandler field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse) {
	o.RegisteredEventHandler = &v
}

// GetEventHandlerParameters returns the EventHandlerParameters field value if set, zero value otherwise.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse {
	if o == nil || isNil(o.EventHandlerParameters) {
		var ret []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse
		return ret
	}
	return o.EventHandlerParameters
}

// GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetEventHandlerParametersOk() ([]KeyfactorApiModelsEventHandlerEventHandlerParameterResponse, bool) {
	if o == nil || isNil(o.EventHandlerParameters) {
		return nil, false
	}
	return o.EventHandlerParameters, true
}

// HasEventHandlerParameters returns a boolean if a field has been set.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasEventHandlerParameters() bool {
	if o != nil && !isNil(o.EventHandlerParameters) {
		return true
	}

	return false
}

// SetEventHandlerParameters gets a reference to the given []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse and assigns it to the EventHandlerParameters field.
func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse) {
	o.EventHandlerParameters = v
}

func (o KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !isNil(o.Subject) {
		toSerialize["Subject"] = o.Subject
	}
	if !isNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !isNil(o.Recipients) {
		toSerialize["Recipients"] = o.Recipients
	}
	if !isNil(o.Template) {
		toSerialize["Template"] = o.Template
	}
	if !isNil(o.RegisteredEventHandler) {
		toSerialize["RegisteredEventHandler"] = o.RegisteredEventHandler
	}
	if !isNil(o.EventHandlerParameters) {
		toSerialize["EventHandlerParameters"] = o.EventHandlerParameters
	}
	return toSerialize, nil
}

type NullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse struct {
	value *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse
	isSet bool
}

func (v NullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) Get() *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse {
	return v.value
}

func (v *NullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) Set(val *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse(val *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) *NullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse {
	return &NullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse{value: val, isSet: true}
}

func (v NullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


