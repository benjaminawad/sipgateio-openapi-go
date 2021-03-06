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

// LocalprefixResponse struct for LocalprefixResponse
type LocalprefixResponse struct {
	Value *string `json:"value,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewLocalprefixResponse instantiates a new LocalprefixResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalprefixResponse() *LocalprefixResponse {
	this := LocalprefixResponse{}
	var active bool = false
	this.Active = &active
	return &this
}

// NewLocalprefixResponseWithDefaults instantiates a new LocalprefixResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalprefixResponseWithDefaults() *LocalprefixResponse {
	this := LocalprefixResponse{}
	var active bool = false
	this.Active = &active
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LocalprefixResponse) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalprefixResponse) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LocalprefixResponse) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *LocalprefixResponse) SetValue(v string) {
	o.Value = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *LocalprefixResponse) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalprefixResponse) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *LocalprefixResponse) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *LocalprefixResponse) SetActive(v bool) {
	o.Active = &v
}

func (o LocalprefixResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableLocalprefixResponse struct {
	value *LocalprefixResponse
	isSet bool
}

func (v NullableLocalprefixResponse) Get() *LocalprefixResponse {
	return v.value
}

func (v *NullableLocalprefixResponse) Set(val *LocalprefixResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalprefixResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalprefixResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalprefixResponse(val *LocalprefixResponse) *NullableLocalprefixResponse {
	return &NullableLocalprefixResponse{value: val, isSet: true}
}

func (v NullableLocalprefixResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalprefixResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


