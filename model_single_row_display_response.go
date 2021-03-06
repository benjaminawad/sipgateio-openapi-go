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

// SingleRowDisplayResponse struct for SingleRowDisplayResponse
type SingleRowDisplayResponse struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSingleRowDisplayResponse instantiates a new SingleRowDisplayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleRowDisplayResponse() *SingleRowDisplayResponse {
	this := SingleRowDisplayResponse{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSingleRowDisplayResponseWithDefaults instantiates a new SingleRowDisplayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleRowDisplayResponseWithDefaults() *SingleRowDisplayResponse {
	this := SingleRowDisplayResponse{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SingleRowDisplayResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleRowDisplayResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SingleRowDisplayResponse) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SingleRowDisplayResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SingleRowDisplayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSingleRowDisplayResponse struct {
	value *SingleRowDisplayResponse
	isSet bool
}

func (v NullableSingleRowDisplayResponse) Get() *SingleRowDisplayResponse {
	return v.value
}

func (v *NullableSingleRowDisplayResponse) Set(val *SingleRowDisplayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleRowDisplayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleRowDisplayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleRowDisplayResponse(val *SingleRowDisplayResponse) *NullableSingleRowDisplayResponse {
	return &NullableSingleRowDisplayResponse{value: val, isSet: true}
}

func (v NullableSingleRowDisplayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleRowDisplayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


