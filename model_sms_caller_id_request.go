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

// SmsCallerIdRequest struct for SmsCallerIdRequest
type SmsCallerIdRequest struct {
	Phonenumber *string `json:"phonenumber,omitempty"`
}

// NewSmsCallerIdRequest instantiates a new SmsCallerIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsCallerIdRequest() *SmsCallerIdRequest {
	this := SmsCallerIdRequest{}
	return &this
}

// NewSmsCallerIdRequestWithDefaults instantiates a new SmsCallerIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsCallerIdRequestWithDefaults() *SmsCallerIdRequest {
	this := SmsCallerIdRequest{}
	return &this
}

// GetPhonenumber returns the Phonenumber field value if set, zero value otherwise.
func (o *SmsCallerIdRequest) GetPhonenumber() string {
	if o == nil || o.Phonenumber == nil {
		var ret string
		return ret
	}
	return *o.Phonenumber
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsCallerIdRequest) GetPhonenumberOk() (*string, bool) {
	if o == nil || o.Phonenumber == nil {
		return nil, false
	}
	return o.Phonenumber, true
}

// HasPhonenumber returns a boolean if a field has been set.
func (o *SmsCallerIdRequest) HasPhonenumber() bool {
	if o != nil && o.Phonenumber != nil {
		return true
	}

	return false
}

// SetPhonenumber gets a reference to the given string and assigns it to the Phonenumber field.
func (o *SmsCallerIdRequest) SetPhonenumber(v string) {
	o.Phonenumber = &v
}

func (o SmsCallerIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Phonenumber != nil {
		toSerialize["phonenumber"] = o.Phonenumber
	}
	return json.Marshal(toSerialize)
}

type NullableSmsCallerIdRequest struct {
	value *SmsCallerIdRequest
	isSet bool
}

func (v NullableSmsCallerIdRequest) Get() *SmsCallerIdRequest {
	return v.value
}

func (v *NullableSmsCallerIdRequest) Set(val *SmsCallerIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsCallerIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsCallerIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsCallerIdRequest(val *SmsCallerIdRequest) *NullableSmsCallerIdRequest {
	return &NullableSmsCallerIdRequest{value: val, isSet: true}
}

func (v NullableSmsCallerIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsCallerIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


