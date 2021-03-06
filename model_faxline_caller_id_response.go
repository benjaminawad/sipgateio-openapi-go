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

// FaxlineCallerIdResponse struct for FaxlineCallerIdResponse
type FaxlineCallerIdResponse struct {
	Value *string `json:"value,omitempty"`
}

// NewFaxlineCallerIdResponse instantiates a new FaxlineCallerIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaxlineCallerIdResponse() *FaxlineCallerIdResponse {
	this := FaxlineCallerIdResponse{}
	return &this
}

// NewFaxlineCallerIdResponseWithDefaults instantiates a new FaxlineCallerIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaxlineCallerIdResponseWithDefaults() *FaxlineCallerIdResponse {
	this := FaxlineCallerIdResponse{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FaxlineCallerIdResponse) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaxlineCallerIdResponse) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FaxlineCallerIdResponse) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FaxlineCallerIdResponse) SetValue(v string) {
	o.Value = &v
}

func (o FaxlineCallerIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFaxlineCallerIdResponse struct {
	value *FaxlineCallerIdResponse
	isSet bool
}

func (v NullableFaxlineCallerIdResponse) Get() *FaxlineCallerIdResponse {
	return v.value
}

func (v *NullableFaxlineCallerIdResponse) Set(val *FaxlineCallerIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFaxlineCallerIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFaxlineCallerIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaxlineCallerIdResponse(val *FaxlineCallerIdResponse) *NullableFaxlineCallerIdResponse {
	return &NullableFaxlineCallerIdResponse{value: val, isSet: true}
}

func (v NullableFaxlineCallerIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaxlineCallerIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


