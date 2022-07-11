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

// CreateMobileDeviceRequest struct for CreateMobileDeviceRequest
type CreateMobileDeviceRequest struct {
	Alias *string `json:"alias,omitempty"`
}

// NewCreateMobileDeviceRequest instantiates a new CreateMobileDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMobileDeviceRequest() *CreateMobileDeviceRequest {
	this := CreateMobileDeviceRequest{}
	return &this
}

// NewCreateMobileDeviceRequestWithDefaults instantiates a new CreateMobileDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMobileDeviceRequestWithDefaults() *CreateMobileDeviceRequest {
	this := CreateMobileDeviceRequest{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *CreateMobileDeviceRequest) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMobileDeviceRequest) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *CreateMobileDeviceRequest) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *CreateMobileDeviceRequest) SetAlias(v string) {
	o.Alias = &v
}

func (o CreateMobileDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMobileDeviceRequest struct {
	value *CreateMobileDeviceRequest
	isSet bool
}

func (v NullableCreateMobileDeviceRequest) Get() *CreateMobileDeviceRequest {
	return v.value
}

func (v *NullableCreateMobileDeviceRequest) Set(val *CreateMobileDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMobileDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMobileDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMobileDeviceRequest(val *CreateMobileDeviceRequest) *NullableCreateMobileDeviceRequest {
	return &NullableCreateMobileDeviceRequest{value: val, isSet: true}
}

func (v NullableCreateMobileDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMobileDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


