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

// AddFaxReportEmailNotificationRequest struct for AddFaxReportEmailNotificationRequest
type AddFaxReportEmailNotificationRequest struct {
	FaxlineId string `json:"faxlineId"`
	Email string `json:"email"`
}

// NewAddFaxReportEmailNotificationRequest instantiates a new AddFaxReportEmailNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFaxReportEmailNotificationRequest(faxlineId string, email string) *AddFaxReportEmailNotificationRequest {
	this := AddFaxReportEmailNotificationRequest{}
	this.FaxlineId = faxlineId
	this.Email = email
	return &this
}

// NewAddFaxReportEmailNotificationRequestWithDefaults instantiates a new AddFaxReportEmailNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFaxReportEmailNotificationRequestWithDefaults() *AddFaxReportEmailNotificationRequest {
	this := AddFaxReportEmailNotificationRequest{}
	return &this
}

// GetFaxlineId returns the FaxlineId field value
func (o *AddFaxReportEmailNotificationRequest) GetFaxlineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FaxlineId
}

// GetFaxlineIdOk returns a tuple with the FaxlineId field value
// and a boolean to check if the value has been set.
func (o *AddFaxReportEmailNotificationRequest) GetFaxlineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FaxlineId, true
}

// SetFaxlineId sets field value
func (o *AddFaxReportEmailNotificationRequest) SetFaxlineId(v string) {
	o.FaxlineId = v
}

// GetEmail returns the Email field value
func (o *AddFaxReportEmailNotificationRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AddFaxReportEmailNotificationRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AddFaxReportEmailNotificationRequest) SetEmail(v string) {
	o.Email = v
}

func (o AddFaxReportEmailNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["faxlineId"] = o.FaxlineId
	}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableAddFaxReportEmailNotificationRequest struct {
	value *AddFaxReportEmailNotificationRequest
	isSet bool
}

func (v NullableAddFaxReportEmailNotificationRequest) Get() *AddFaxReportEmailNotificationRequest {
	return v.value
}

func (v *NullableAddFaxReportEmailNotificationRequest) Set(val *AddFaxReportEmailNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFaxReportEmailNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFaxReportEmailNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFaxReportEmailNotificationRequest(val *AddFaxReportEmailNotificationRequest) *NullableAddFaxReportEmailNotificationRequest {
	return &NullableAddFaxReportEmailNotificationRequest{value: val, isSet: true}
}

func (v NullableAddFaxReportEmailNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFaxReportEmailNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


