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

// SetHistoryEntryArchiveRequest struct for SetHistoryEntryArchiveRequest
type SetHistoryEntryArchiveRequest struct {
	Value *bool `json:"value,omitempty"`
}

// NewSetHistoryEntryArchiveRequest instantiates a new SetHistoryEntryArchiveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetHistoryEntryArchiveRequest() *SetHistoryEntryArchiveRequest {
	this := SetHistoryEntryArchiveRequest{}
	var value bool = false
	this.Value = &value
	return &this
}

// NewSetHistoryEntryArchiveRequestWithDefaults instantiates a new SetHistoryEntryArchiveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetHistoryEntryArchiveRequestWithDefaults() *SetHistoryEntryArchiveRequest {
	this := SetHistoryEntryArchiveRequest{}
	var value bool = false
	this.Value = &value
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SetHistoryEntryArchiveRequest) GetValue() bool {
	if o == nil || o.Value == nil {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetHistoryEntryArchiveRequest) GetValueOk() (*bool, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SetHistoryEntryArchiveRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *SetHistoryEntryArchiveRequest) SetValue(v bool) {
	o.Value = &v
}

func (o SetHistoryEntryArchiveRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSetHistoryEntryArchiveRequest struct {
	value *SetHistoryEntryArchiveRequest
	isSet bool
}

func (v NullableSetHistoryEntryArchiveRequest) Get() *SetHistoryEntryArchiveRequest {
	return v.value
}

func (v *NullableSetHistoryEntryArchiveRequest) Set(val *SetHistoryEntryArchiveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetHistoryEntryArchiveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetHistoryEntryArchiveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetHistoryEntryArchiveRequest(val *SetHistoryEntryArchiveRequest) *NullableSetHistoryEntryArchiveRequest {
	return &NullableSetHistoryEntryArchiveRequest{value: val, isSet: true}
}

func (v NullableSetHistoryEntryArchiveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetHistoryEntryArchiveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


