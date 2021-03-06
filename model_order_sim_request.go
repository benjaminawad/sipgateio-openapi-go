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

// OrderSimRequest struct for OrderSimRequest
type OrderSimRequest struct {
	AddressId string `json:"addressId"`
}

// NewOrderSimRequest instantiates a new OrderSimRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSimRequest(addressId string) *OrderSimRequest {
	this := OrderSimRequest{}
	this.AddressId = addressId
	return &this
}

// NewOrderSimRequestWithDefaults instantiates a new OrderSimRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSimRequestWithDefaults() *OrderSimRequest {
	this := OrderSimRequest{}
	return &this
}

// GetAddressId returns the AddressId field value
func (o *OrderSimRequest) GetAddressId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value
// and a boolean to check if the value has been set.
func (o *OrderSimRequest) GetAddressIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressId, true
}

// SetAddressId sets field value
func (o *OrderSimRequest) SetAddressId(v string) {
	o.AddressId = v
}

func (o OrderSimRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addressId"] = o.AddressId
	}
	return json.Marshal(toSerialize)
}

type NullableOrderSimRequest struct {
	value *OrderSimRequest
	isSet bool
}

func (v NullableOrderSimRequest) Get() *OrderSimRequest {
	return v.value
}

func (v *NullableOrderSimRequest) Set(val *OrderSimRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSimRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSimRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSimRequest(val *OrderSimRequest) *NullableOrderSimRequest {
	return &NullableOrderSimRequest{value: val, isSet: true}
}

func (v NullableOrderSimRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSimRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


