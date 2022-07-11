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

// AddCallSmsNotificationRequest struct for AddCallSmsNotificationRequest
type AddCallSmsNotificationRequest struct {
	EndpointId string `json:"endpointId"`
	Direction string `json:"direction"`
	Cause string `json:"cause"`
	Number string `json:"number"`
}

// NewAddCallSmsNotificationRequest instantiates a new AddCallSmsNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCallSmsNotificationRequest(endpointId string, direction string, cause string, number string) *AddCallSmsNotificationRequest {
	this := AddCallSmsNotificationRequest{}
	this.EndpointId = endpointId
	this.Direction = direction
	this.Cause = cause
	this.Number = number
	return &this
}

// NewAddCallSmsNotificationRequestWithDefaults instantiates a new AddCallSmsNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCallSmsNotificationRequestWithDefaults() *AddCallSmsNotificationRequest {
	this := AddCallSmsNotificationRequest{}
	return &this
}

// GetEndpointId returns the EndpointId field value
func (o *AddCallSmsNotificationRequest) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *AddCallSmsNotificationRequest) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *AddCallSmsNotificationRequest) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetDirection returns the Direction field value
func (o *AddCallSmsNotificationRequest) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *AddCallSmsNotificationRequest) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *AddCallSmsNotificationRequest) SetDirection(v string) {
	o.Direction = v
}

// GetCause returns the Cause field value
func (o *AddCallSmsNotificationRequest) GetCause() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *AddCallSmsNotificationRequest) GetCauseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *AddCallSmsNotificationRequest) SetCause(v string) {
	o.Cause = v
}

// GetNumber returns the Number field value
func (o *AddCallSmsNotificationRequest) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *AddCallSmsNotificationRequest) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *AddCallSmsNotificationRequest) SetNumber(v string) {
	o.Number = v
}

func (o AddCallSmsNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["endpointId"] = o.EndpointId
	}
	if true {
		toSerialize["direction"] = o.Direction
	}
	if true {
		toSerialize["cause"] = o.Cause
	}
	if true {
		toSerialize["number"] = o.Number
	}
	return json.Marshal(toSerialize)
}

type NullableAddCallSmsNotificationRequest struct {
	value *AddCallSmsNotificationRequest
	isSet bool
}

func (v NullableAddCallSmsNotificationRequest) Get() *AddCallSmsNotificationRequest {
	return v.value
}

func (v *NullableAddCallSmsNotificationRequest) Set(val *AddCallSmsNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCallSmsNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCallSmsNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCallSmsNotificationRequest(val *AddCallSmsNotificationRequest) *NullableAddCallSmsNotificationRequest {
	return &NullableAddCallSmsNotificationRequest{value: val, isSet: true}
}

func (v NullableAddCallSmsNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCallSmsNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

