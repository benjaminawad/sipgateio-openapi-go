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

// SmsTarget struct for SmsTarget
type SmsTarget struct {
	Id *string `json:"id,omitempty"`
	Number *string `json:"number,omitempty"`
}

// NewSmsTarget instantiates a new SmsTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsTarget() *SmsTarget {
	this := SmsTarget{}
	return &this
}

// NewSmsTargetWithDefaults instantiates a new SmsTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsTargetWithDefaults() *SmsTarget {
	this := SmsTarget{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SmsTarget) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTarget) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SmsTarget) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SmsTarget) SetId(v string) {
	o.Id = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SmsTarget) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTarget) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SmsTarget) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *SmsTarget) SetNumber(v string) {
	o.Number = &v
}

func (o SmsTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableSmsTarget struct {
	value *SmsTarget
	isSet bool
}

func (v NullableSmsTarget) Get() *SmsTarget {
	return v.value
}

func (v *NullableSmsTarget) Set(val *SmsTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsTarget(val *SmsTarget) *NullableSmsTarget {
	return &NullableSmsTarget{value: val, isSet: true}
}

func (v NullableSmsTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

