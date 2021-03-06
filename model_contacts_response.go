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

// ContactsResponse struct for ContactsResponse
type ContactsResponse struct {
	Items []ContactResponse `json:"items,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewContactsResponse instantiates a new ContactsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactsResponse() *ContactsResponse {
	this := ContactsResponse{}
	return &this
}

// NewContactsResponseWithDefaults instantiates a new ContactsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactsResponseWithDefaults() *ContactsResponse {
	this := ContactsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ContactsResponse) GetItems() []ContactResponse {
	if o == nil || o.Items == nil {
		var ret []ContactResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactsResponse) GetItemsOk() ([]ContactResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ContactsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ContactResponse and assigns it to the Items field.
func (o *ContactsResponse) SetItems(v []ContactResponse) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ContactsResponse) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactsResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ContactsResponse) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ContactsResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ContactsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableContactsResponse struct {
	value *ContactsResponse
	isSet bool
}

func (v NullableContactsResponse) Get() *ContactsResponse {
	return v.value
}

func (v *NullableContactsResponse) Set(val *ContactsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactsResponse(val *ContactsResponse) *NullableContactsResponse {
	return &NullableContactsResponse{value: val, isSet: true}
}

func (v NullableContactsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


