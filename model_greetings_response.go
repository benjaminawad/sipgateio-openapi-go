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

// GreetingsResponse struct for GreetingsResponse
type GreetingsResponse struct {
	Items []GreetingResponse `json:"items,omitempty"`
}

// NewGreetingsResponse instantiates a new GreetingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreetingsResponse() *GreetingsResponse {
	this := GreetingsResponse{}
	return &this
}

// NewGreetingsResponseWithDefaults instantiates a new GreetingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreetingsResponseWithDefaults() *GreetingsResponse {
	this := GreetingsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GreetingsResponse) GetItems() []GreetingResponse {
	if o == nil || o.Items == nil {
		var ret []GreetingResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreetingsResponse) GetItemsOk() ([]GreetingResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GreetingsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []GreetingResponse and assigns it to the Items field.
func (o *GreetingsResponse) SetItems(v []GreetingResponse) {
	o.Items = v
}

func (o GreetingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableGreetingsResponse struct {
	value *GreetingsResponse
	isSet bool
}

func (v NullableGreetingsResponse) Get() *GreetingsResponse {
	return v.value
}

func (v *NullableGreetingsResponse) Set(val *GreetingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGreetingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGreetingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreetingsResponse(val *GreetingsResponse) *NullableGreetingsResponse {
	return &NullableGreetingsResponse{value: val, isSet: true}
}

func (v NullableGreetingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreetingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

