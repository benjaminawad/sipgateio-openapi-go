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

// ChangeVoicemailGreetingRequest struct for ChangeVoicemailGreetingRequest
type ChangeVoicemailGreetingRequest struct {
	Active *bool `json:"active,omitempty"`
}

// NewChangeVoicemailGreetingRequest instantiates a new ChangeVoicemailGreetingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeVoicemailGreetingRequest() *ChangeVoicemailGreetingRequest {
	this := ChangeVoicemailGreetingRequest{}
	var active bool = false
	this.Active = &active
	return &this
}

// NewChangeVoicemailGreetingRequestWithDefaults instantiates a new ChangeVoicemailGreetingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeVoicemailGreetingRequestWithDefaults() *ChangeVoicemailGreetingRequest {
	this := ChangeVoicemailGreetingRequest{}
	var active bool = false
	this.Active = &active
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ChangeVoicemailGreetingRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeVoicemailGreetingRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ChangeVoicemailGreetingRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ChangeVoicemailGreetingRequest) SetActive(v bool) {
	o.Active = &v
}

func (o ChangeVoicemailGreetingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableChangeVoicemailGreetingRequest struct {
	value *ChangeVoicemailGreetingRequest
	isSet bool
}

func (v NullableChangeVoicemailGreetingRequest) Get() *ChangeVoicemailGreetingRequest {
	return v.value
}

func (v *NullableChangeVoicemailGreetingRequest) Set(val *ChangeVoicemailGreetingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeVoicemailGreetingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeVoicemailGreetingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeVoicemailGreetingRequest(val *ChangeVoicemailGreetingRequest) *NullableChangeVoicemailGreetingRequest {
	return &NullableChangeVoicemailGreetingRequest{value: val, isSet: true}
}

func (v NullableChangeVoicemailGreetingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeVoicemailGreetingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

