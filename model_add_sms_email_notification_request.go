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

// AddSmsEmailNotificationRequest struct for AddSmsEmailNotificationRequest
type AddSmsEmailNotificationRequest struct {
	EndpointId string `json:"endpointId"`
	Email string `json:"email"`
}

// NewAddSmsEmailNotificationRequest instantiates a new AddSmsEmailNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSmsEmailNotificationRequest(endpointId string, email string) *AddSmsEmailNotificationRequest {
	this := AddSmsEmailNotificationRequest{}
	this.EndpointId = endpointId
	this.Email = email
	return &this
}

// NewAddSmsEmailNotificationRequestWithDefaults instantiates a new AddSmsEmailNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSmsEmailNotificationRequestWithDefaults() *AddSmsEmailNotificationRequest {
	this := AddSmsEmailNotificationRequest{}
	return &this
}

// GetEndpointId returns the EndpointId field value
func (o *AddSmsEmailNotificationRequest) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *AddSmsEmailNotificationRequest) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *AddSmsEmailNotificationRequest) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetEmail returns the Email field value
func (o *AddSmsEmailNotificationRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AddSmsEmailNotificationRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AddSmsEmailNotificationRequest) SetEmail(v string) {
	o.Email = v
}

func (o AddSmsEmailNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["endpointId"] = o.EndpointId
	}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableAddSmsEmailNotificationRequest struct {
	value *AddSmsEmailNotificationRequest
	isSet bool
}

func (v NullableAddSmsEmailNotificationRequest) Get() *AddSmsEmailNotificationRequest {
	return v.value
}

func (v *NullableAddSmsEmailNotificationRequest) Set(val *AddSmsEmailNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSmsEmailNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSmsEmailNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSmsEmailNotificationRequest(val *AddSmsEmailNotificationRequest) *NullableAddSmsEmailNotificationRequest {
	return &NullableAddSmsEmailNotificationRequest{value: val, isSet: true}
}

func (v NullableAddSmsEmailNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSmsEmailNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


