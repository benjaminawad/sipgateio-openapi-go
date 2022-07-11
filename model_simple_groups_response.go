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

// SimpleGroupsResponse struct for SimpleGroupsResponse
type SimpleGroupsResponse struct {
	Items []SimpleGroupResponse `json:"items,omitempty"`
}

// NewSimpleGroupsResponse instantiates a new SimpleGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleGroupsResponse() *SimpleGroupsResponse {
	this := SimpleGroupsResponse{}
	return &this
}

// NewSimpleGroupsResponseWithDefaults instantiates a new SimpleGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleGroupsResponseWithDefaults() *SimpleGroupsResponse {
	this := SimpleGroupsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SimpleGroupsResponse) GetItems() []SimpleGroupResponse {
	if o == nil || o.Items == nil {
		var ret []SimpleGroupResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleGroupsResponse) GetItemsOk() ([]SimpleGroupResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SimpleGroupsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SimpleGroupResponse and assigns it to the Items field.
func (o *SimpleGroupsResponse) SetItems(v []SimpleGroupResponse) {
	o.Items = v
}

func (o SimpleGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleGroupsResponse struct {
	value *SimpleGroupsResponse
	isSet bool
}

func (v NullableSimpleGroupsResponse) Get() *SimpleGroupsResponse {
	return v.value
}

func (v *NullableSimpleGroupsResponse) Set(val *SimpleGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleGroupsResponse(val *SimpleGroupsResponse) *NullableSimpleGroupsResponse {
	return &NullableSimpleGroupsResponse{value: val, isSet: true}
}

func (v NullableSimpleGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


