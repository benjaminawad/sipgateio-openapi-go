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

// AddVoicemailEmailVoicemailNotificationRequest struct for AddVoicemailEmailVoicemailNotificationRequest
type AddVoicemailEmailVoicemailNotificationRequest struct {
	VoicemailId string `json:"voicemailId"`
	Email string `json:"email"`
}

// NewAddVoicemailEmailVoicemailNotificationRequest instantiates a new AddVoicemailEmailVoicemailNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVoicemailEmailVoicemailNotificationRequest(voicemailId string, email string) *AddVoicemailEmailVoicemailNotificationRequest {
	this := AddVoicemailEmailVoicemailNotificationRequest{}
	this.VoicemailId = voicemailId
	this.Email = email
	return &this
}

// NewAddVoicemailEmailVoicemailNotificationRequestWithDefaults instantiates a new AddVoicemailEmailVoicemailNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVoicemailEmailVoicemailNotificationRequestWithDefaults() *AddVoicemailEmailVoicemailNotificationRequest {
	this := AddVoicemailEmailVoicemailNotificationRequest{}
	return &this
}

// GetVoicemailId returns the VoicemailId field value
func (o *AddVoicemailEmailVoicemailNotificationRequest) GetVoicemailId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VoicemailId
}

// GetVoicemailIdOk returns a tuple with the VoicemailId field value
// and a boolean to check if the value has been set.
func (o *AddVoicemailEmailVoicemailNotificationRequest) GetVoicemailIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VoicemailId, true
}

// SetVoicemailId sets field value
func (o *AddVoicemailEmailVoicemailNotificationRequest) SetVoicemailId(v string) {
	o.VoicemailId = v
}

// GetEmail returns the Email field value
func (o *AddVoicemailEmailVoicemailNotificationRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AddVoicemailEmailVoicemailNotificationRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AddVoicemailEmailVoicemailNotificationRequest) SetEmail(v string) {
	o.Email = v
}

func (o AddVoicemailEmailVoicemailNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["voicemailId"] = o.VoicemailId
	}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableAddVoicemailEmailVoicemailNotificationRequest struct {
	value *AddVoicemailEmailVoicemailNotificationRequest
	isSet bool
}

func (v NullableAddVoicemailEmailVoicemailNotificationRequest) Get() *AddVoicemailEmailVoicemailNotificationRequest {
	return v.value
}

func (v *NullableAddVoicemailEmailVoicemailNotificationRequest) Set(val *AddVoicemailEmailVoicemailNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVoicemailEmailVoicemailNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVoicemailEmailVoicemailNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVoicemailEmailVoicemailNotificationRequest(val *AddVoicemailEmailVoicemailNotificationRequest) *NullableAddVoicemailEmailVoicemailNotificationRequest {
	return &NullableAddVoicemailEmailVoicemailNotificationRequest{value: val, isSet: true}
}

func (v NullableAddVoicemailEmailVoicemailNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVoicemailEmailVoicemailNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

