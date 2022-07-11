/*
sipgate API

This is the sipgate REST API documentation. We build our applications on this API and we invite you to use it too.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PhonelineResponse struct for PhonelineResponse
type PhonelineResponse struct {
	Id *string `json:"id,omitempty"`
	Alias *string `json:"alias,omitempty"`
	Voicemails []VoicemailSettingsResponse `json:"voicemails,omitempty"`
}

// NewPhonelineResponse instantiates a new PhonelineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhonelineResponse() *PhonelineResponse {
	this := PhonelineResponse{}
	return &this
}

// NewPhonelineResponseWithDefaults instantiates a new PhonelineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhonelineResponseWithDefaults() *PhonelineResponse {
	this := PhonelineResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PhonelineResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PhonelineResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PhonelineResponse) SetId(v string) {
	o.Id = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *PhonelineResponse) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineResponse) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *PhonelineResponse) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *PhonelineResponse) SetAlias(v string) {
	o.Alias = &v
}

// GetVoicemails returns the Voicemails field value if set, zero value otherwise.
func (o *PhonelineResponse) GetVoicemails() []VoicemailSettingsResponse {
	if o == nil || o.Voicemails == nil {
		var ret []VoicemailSettingsResponse
		return ret
	}
	return o.Voicemails
}

// GetVoicemailsOk returns a tuple with the Voicemails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineResponse) GetVoicemailsOk() ([]VoicemailSettingsResponse, bool) {
	if o == nil || o.Voicemails == nil {
		return nil, false
	}
	return o.Voicemails, true
}

// HasVoicemails returns a boolean if a field has been set.
func (o *PhonelineResponse) HasVoicemails() bool {
	if o != nil && o.Voicemails != nil {
		return true
	}

	return false
}

// SetVoicemails gets a reference to the given []VoicemailSettingsResponse and assigns it to the Voicemails field.
func (o *PhonelineResponse) SetVoicemails(v []VoicemailSettingsResponse) {
	o.Voicemails = v
}

func (o PhonelineResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Voicemails != nil {
		toSerialize["voicemails"] = o.Voicemails
	}
	return json.Marshal(toSerialize)
}

type NullablePhonelineResponse struct {
	value *PhonelineResponse
	isSet bool
}

func (v NullablePhonelineResponse) Get() *PhonelineResponse {
	return v.value
}

func (v *NullablePhonelineResponse) Set(val *PhonelineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePhonelineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePhonelineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhonelineResponse(val *PhonelineResponse) *NullablePhonelineResponse {
	return &NullablePhonelineResponse{value: val, isSet: true}
}

func (v NullablePhonelineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhonelineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


