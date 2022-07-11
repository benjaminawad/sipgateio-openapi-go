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

// SetAliasRequest struct for SetAliasRequest
type SetAliasRequest struct {
	Alias *string `json:"alias,omitempty"`
}

// NewSetAliasRequest instantiates a new SetAliasRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetAliasRequest() *SetAliasRequest {
	this := SetAliasRequest{}
	return &this
}

// NewSetAliasRequestWithDefaults instantiates a new SetAliasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetAliasRequestWithDefaults() *SetAliasRequest {
	this := SetAliasRequest{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *SetAliasRequest) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetAliasRequest) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *SetAliasRequest) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *SetAliasRequest) SetAlias(v string) {
	o.Alias = &v
}

func (o SetAliasRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	return json.Marshal(toSerialize)
}

type NullableSetAliasRequest struct {
	value *SetAliasRequest
	isSet bool
}

func (v NullableSetAliasRequest) Get() *SetAliasRequest {
	return v.value
}

func (v *NullableSetAliasRequest) Set(val *SetAliasRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetAliasRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetAliasRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetAliasRequest(val *SetAliasRequest) *NullableSetAliasRequest {
	return &NullableSetAliasRequest{value: val, isSet: true}
}

func (v NullableSetAliasRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetAliasRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


