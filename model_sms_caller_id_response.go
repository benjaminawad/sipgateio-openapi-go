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

// SmsCallerIdResponse struct for SmsCallerIdResponse
type SmsCallerIdResponse struct {
	Id *int32 `json:"id,omitempty"`
	Phonenumber *string `json:"phonenumber,omitempty"`
	Verified *bool `json:"verified,omitempty"`
	DefaultNumber *bool `json:"defaultNumber,omitempty"`
}

// NewSmsCallerIdResponse instantiates a new SmsCallerIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsCallerIdResponse() *SmsCallerIdResponse {
	this := SmsCallerIdResponse{}
	var verified bool = false
	this.Verified = &verified
	var defaultNumber bool = false
	this.DefaultNumber = &defaultNumber
	return &this
}

// NewSmsCallerIdResponseWithDefaults instantiates a new SmsCallerIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsCallerIdResponseWithDefaults() *SmsCallerIdResponse {
	this := SmsCallerIdResponse{}
	var verified bool = false
	this.Verified = &verified
	var defaultNumber bool = false
	this.DefaultNumber = &defaultNumber
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SmsCallerIdResponse) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsCallerIdResponse) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SmsCallerIdResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SmsCallerIdResponse) SetId(v int32) {
	o.Id = &v
}

// GetPhonenumber returns the Phonenumber field value if set, zero value otherwise.
func (o *SmsCallerIdResponse) GetPhonenumber() string {
	if o == nil || o.Phonenumber == nil {
		var ret string
		return ret
	}
	return *o.Phonenumber
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsCallerIdResponse) GetPhonenumberOk() (*string, bool) {
	if o == nil || o.Phonenumber == nil {
		return nil, false
	}
	return o.Phonenumber, true
}

// HasPhonenumber returns a boolean if a field has been set.
func (o *SmsCallerIdResponse) HasPhonenumber() bool {
	if o != nil && o.Phonenumber != nil {
		return true
	}

	return false
}

// SetPhonenumber gets a reference to the given string and assigns it to the Phonenumber field.
func (o *SmsCallerIdResponse) SetPhonenumber(v string) {
	o.Phonenumber = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *SmsCallerIdResponse) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsCallerIdResponse) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *SmsCallerIdResponse) HasVerified() bool {
	if o != nil && o.Verified != nil {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *SmsCallerIdResponse) SetVerified(v bool) {
	o.Verified = &v
}

// GetDefaultNumber returns the DefaultNumber field value if set, zero value otherwise.
func (o *SmsCallerIdResponse) GetDefaultNumber() bool {
	if o == nil || o.DefaultNumber == nil {
		var ret bool
		return ret
	}
	return *o.DefaultNumber
}

// GetDefaultNumberOk returns a tuple with the DefaultNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsCallerIdResponse) GetDefaultNumberOk() (*bool, bool) {
	if o == nil || o.DefaultNumber == nil {
		return nil, false
	}
	return o.DefaultNumber, true
}

// HasDefaultNumber returns a boolean if a field has been set.
func (o *SmsCallerIdResponse) HasDefaultNumber() bool {
	if o != nil && o.DefaultNumber != nil {
		return true
	}

	return false
}

// SetDefaultNumber gets a reference to the given bool and assigns it to the DefaultNumber field.
func (o *SmsCallerIdResponse) SetDefaultNumber(v bool) {
	o.DefaultNumber = &v
}

func (o SmsCallerIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Phonenumber != nil {
		toSerialize["phonenumber"] = o.Phonenumber
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}
	if o.DefaultNumber != nil {
		toSerialize["defaultNumber"] = o.DefaultNumber
	}
	return json.Marshal(toSerialize)
}

type NullableSmsCallerIdResponse struct {
	value *SmsCallerIdResponse
	isSet bool
}

func (v NullableSmsCallerIdResponse) Get() *SmsCallerIdResponse {
	return v.value
}

func (v *NullableSmsCallerIdResponse) Set(val *SmsCallerIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsCallerIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsCallerIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsCallerIdResponse(val *SmsCallerIdResponse) *NullableSmsCallerIdResponse {
	return &NullableSmsCallerIdResponse{value: val, isSet: true}
}

func (v NullableSmsCallerIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsCallerIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


