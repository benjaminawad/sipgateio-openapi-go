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

// PhonelineBlockAnonymousResponse struct for PhonelineBlockAnonymousResponse
type PhonelineBlockAnonymousResponse struct {
	Enabled *bool `json:"enabled,omitempty"`
	Target *string `json:"target,omitempty"`
}

// NewPhonelineBlockAnonymousResponse instantiates a new PhonelineBlockAnonymousResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhonelineBlockAnonymousResponse() *PhonelineBlockAnonymousResponse {
	this := PhonelineBlockAnonymousResponse{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewPhonelineBlockAnonymousResponseWithDefaults instantiates a new PhonelineBlockAnonymousResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhonelineBlockAnonymousResponseWithDefaults() *PhonelineBlockAnonymousResponse {
	this := PhonelineBlockAnonymousResponse{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PhonelineBlockAnonymousResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineBlockAnonymousResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PhonelineBlockAnonymousResponse) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PhonelineBlockAnonymousResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *PhonelineBlockAnonymousResponse) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineBlockAnonymousResponse) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *PhonelineBlockAnonymousResponse) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *PhonelineBlockAnonymousResponse) SetTarget(v string) {
	o.Target = &v
}

func (o PhonelineBlockAnonymousResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullablePhonelineBlockAnonymousResponse struct {
	value *PhonelineBlockAnonymousResponse
	isSet bool
}

func (v NullablePhonelineBlockAnonymousResponse) Get() *PhonelineBlockAnonymousResponse {
	return v.value
}

func (v *NullablePhonelineBlockAnonymousResponse) Set(val *PhonelineBlockAnonymousResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePhonelineBlockAnonymousResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePhonelineBlockAnonymousResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhonelineBlockAnonymousResponse(val *PhonelineBlockAnonymousResponse) *NullablePhonelineBlockAnonymousResponse {
	return &NullablePhonelineBlockAnonymousResponse{value: val, isSet: true}
}

func (v NullablePhonelineBlockAnonymousResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhonelineBlockAnonymousResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


