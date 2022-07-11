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

// SetBlockAnonymousSettingsRequest struct for SetBlockAnonymousSettingsRequest
type SetBlockAnonymousSettingsRequest struct {
	Enabled *bool `json:"enabled,omitempty"`
	Target *string `json:"target,omitempty"`
}

// NewSetBlockAnonymousSettingsRequest instantiates a new SetBlockAnonymousSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetBlockAnonymousSettingsRequest() *SetBlockAnonymousSettingsRequest {
	this := SetBlockAnonymousSettingsRequest{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSetBlockAnonymousSettingsRequestWithDefaults instantiates a new SetBlockAnonymousSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetBlockAnonymousSettingsRequestWithDefaults() *SetBlockAnonymousSettingsRequest {
	this := SetBlockAnonymousSettingsRequest{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SetBlockAnonymousSettingsRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetBlockAnonymousSettingsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SetBlockAnonymousSettingsRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SetBlockAnonymousSettingsRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SetBlockAnonymousSettingsRequest) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetBlockAnonymousSettingsRequest) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SetBlockAnonymousSettingsRequest) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *SetBlockAnonymousSettingsRequest) SetTarget(v string) {
	o.Target = &v
}

func (o SetBlockAnonymousSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableSetBlockAnonymousSettingsRequest struct {
	value *SetBlockAnonymousSettingsRequest
	isSet bool
}

func (v NullableSetBlockAnonymousSettingsRequest) Get() *SetBlockAnonymousSettingsRequest {
	return v.value
}

func (v *NullableSetBlockAnonymousSettingsRequest) Set(val *SetBlockAnonymousSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetBlockAnonymousSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetBlockAnonymousSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetBlockAnonymousSettingsRequest(val *SetBlockAnonymousSettingsRequest) *NullableSetBlockAnonymousSettingsRequest {
	return &NullableSetBlockAnonymousSettingsRequest{value: val, isSet: true}
}

func (v NullableSetBlockAnonymousSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetBlockAnonymousSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


