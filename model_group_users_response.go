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

// GroupUsersResponse struct for GroupUsersResponse
type GroupUsersResponse struct {
	Items []GroupUserResponse `json:"items,omitempty"`
}

// NewGroupUsersResponse instantiates a new GroupUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUsersResponse() *GroupUsersResponse {
	this := GroupUsersResponse{}
	return &this
}

// NewGroupUsersResponseWithDefaults instantiates a new GroupUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUsersResponseWithDefaults() *GroupUsersResponse {
	this := GroupUsersResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GroupUsersResponse) GetItems() []GroupUserResponse {
	if o == nil || o.Items == nil {
		var ret []GroupUserResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUsersResponse) GetItemsOk() ([]GroupUserResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GroupUsersResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []GroupUserResponse and assigns it to the Items field.
func (o *GroupUsersResponse) SetItems(v []GroupUserResponse) {
	o.Items = v
}

func (o GroupUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableGroupUsersResponse struct {
	value *GroupUsersResponse
	isSet bool
}

func (v NullableGroupUsersResponse) Get() *GroupUsersResponse {
	return v.value
}

func (v *NullableGroupUsersResponse) Set(val *GroupUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUsersResponse(val *GroupUsersResponse) *NullableGroupUsersResponse {
	return &NullableGroupUsersResponse{value: val, isSet: true}
}

func (v NullableGroupUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


