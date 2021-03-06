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

// SetBusyOnBusySettingsRequest struct for SetBusyOnBusySettingsRequest
type SetBusyOnBusySettingsRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSetBusyOnBusySettingsRequest instantiates a new SetBusyOnBusySettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetBusyOnBusySettingsRequest() *SetBusyOnBusySettingsRequest {
	this := SetBusyOnBusySettingsRequest{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSetBusyOnBusySettingsRequestWithDefaults instantiates a new SetBusyOnBusySettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetBusyOnBusySettingsRequestWithDefaults() *SetBusyOnBusySettingsRequest {
	this := SetBusyOnBusySettingsRequest{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SetBusyOnBusySettingsRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetBusyOnBusySettingsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SetBusyOnBusySettingsRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SetBusyOnBusySettingsRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SetBusyOnBusySettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSetBusyOnBusySettingsRequest struct {
	value *SetBusyOnBusySettingsRequest
	isSet bool
}

func (v NullableSetBusyOnBusySettingsRequest) Get() *SetBusyOnBusySettingsRequest {
	return v.value
}

func (v *NullableSetBusyOnBusySettingsRequest) Set(val *SetBusyOnBusySettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetBusyOnBusySettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetBusyOnBusySettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetBusyOnBusySettingsRequest(val *SetBusyOnBusySettingsRequest) *NullableSetBusyOnBusySettingsRequest {
	return &NullableSetBusyOnBusySettingsRequest{value: val, isSet: true}
}

func (v NullableSetBusyOnBusySettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetBusyOnBusySettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


