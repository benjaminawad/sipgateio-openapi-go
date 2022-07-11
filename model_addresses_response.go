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

// AddressesResponse struct for AddressesResponse
type AddressesResponse struct {
	Items []AddressResponse `json:"items,omitempty"`
}

// NewAddressesResponse instantiates a new AddressesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressesResponse() *AddressesResponse {
	this := AddressesResponse{}
	return &this
}

// NewAddressesResponseWithDefaults instantiates a new AddressesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressesResponseWithDefaults() *AddressesResponse {
	this := AddressesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AddressesResponse) GetItems() []AddressResponse {
	if o == nil || o.Items == nil {
		var ret []AddressResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressesResponse) GetItemsOk() ([]AddressResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AddressesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AddressResponse and assigns it to the Items field.
func (o *AddressesResponse) SetItems(v []AddressResponse) {
	o.Items = v
}

func (o AddressesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableAddressesResponse struct {
	value *AddressesResponse
	isSet bool
}

func (v NullableAddressesResponse) Get() *AddressesResponse {
	return v.value
}

func (v *NullableAddressesResponse) Set(val *AddressesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressesResponse(val *AddressesResponse) *NullableAddressesResponse {
	return &NullableAddressesResponse{value: val, isSet: true}
}

func (v NullableAddressesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


