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

// AddFaxSmsNotificationRequest struct for AddFaxSmsNotificationRequest
type AddFaxSmsNotificationRequest struct {
	FaxlineId string `json:"faxlineId"`
	Direction string `json:"direction"`
	Number string `json:"number"`
}

// NewAddFaxSmsNotificationRequest instantiates a new AddFaxSmsNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFaxSmsNotificationRequest(faxlineId string, direction string, number string) *AddFaxSmsNotificationRequest {
	this := AddFaxSmsNotificationRequest{}
	this.FaxlineId = faxlineId
	this.Direction = direction
	this.Number = number
	return &this
}

// NewAddFaxSmsNotificationRequestWithDefaults instantiates a new AddFaxSmsNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFaxSmsNotificationRequestWithDefaults() *AddFaxSmsNotificationRequest {
	this := AddFaxSmsNotificationRequest{}
	return &this
}

// GetFaxlineId returns the FaxlineId field value
func (o *AddFaxSmsNotificationRequest) GetFaxlineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FaxlineId
}

// GetFaxlineIdOk returns a tuple with the FaxlineId field value
// and a boolean to check if the value has been set.
func (o *AddFaxSmsNotificationRequest) GetFaxlineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FaxlineId, true
}

// SetFaxlineId sets field value
func (o *AddFaxSmsNotificationRequest) SetFaxlineId(v string) {
	o.FaxlineId = v
}

// GetDirection returns the Direction field value
func (o *AddFaxSmsNotificationRequest) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *AddFaxSmsNotificationRequest) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *AddFaxSmsNotificationRequest) SetDirection(v string) {
	o.Direction = v
}

// GetNumber returns the Number field value
func (o *AddFaxSmsNotificationRequest) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *AddFaxSmsNotificationRequest) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *AddFaxSmsNotificationRequest) SetNumber(v string) {
	o.Number = v
}

func (o AddFaxSmsNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["faxlineId"] = o.FaxlineId
	}
	if true {
		toSerialize["direction"] = o.Direction
	}
	if true {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableAddFaxSmsNotificationRequest struct {
	value *AddFaxSmsNotificationRequest
	isSet bool
}

func (v NullableAddFaxSmsNotificationRequest) Get() *AddFaxSmsNotificationRequest {
	return v.value
}

func (v *NullableAddFaxSmsNotificationRequest) Set(val *AddFaxSmsNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFaxSmsNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFaxSmsNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFaxSmsNotificationRequest(val *AddFaxSmsNotificationRequest) *NullableAddFaxSmsNotificationRequest {
	return &NullableAddFaxSmsNotificationRequest{value: val, isSet: true}
}

func (v NullableAddFaxSmsNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFaxSmsNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


