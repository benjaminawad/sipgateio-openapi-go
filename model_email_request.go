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

// EmailRequest struct for EmailRequest
type EmailRequest struct {
	Email *string `json:"email,omitempty"`
	Type []string `json:"type,omitempty"`
}

// NewEmailRequest instantiates a new EmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailRequest() *EmailRequest {
	this := EmailRequest{}
	return &this
}

// NewEmailRequestWithDefaults instantiates a new EmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailRequestWithDefaults() *EmailRequest {
	this := EmailRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EmailRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EmailRequest) SetEmail(v string) {
	o.Email = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EmailRequest) GetType() []string {
	if o == nil || o.Type == nil {
		var ret []string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetTypeOk() ([]string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EmailRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given []string and assigns it to the Type field.
func (o *EmailRequest) SetType(v []string) {
	o.Type = v
}

func (o EmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEmailRequest struct {
	value *EmailRequest
	isSet bool
}

func (v NullableEmailRequest) Get() *EmailRequest {
	return v.value
}

func (v *NullableEmailRequest) Set(val *EmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailRequest(val *EmailRequest) *NullableEmailRequest {
	return &NullableEmailRequest{value: val, isSet: true}
}

func (v NullableEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


