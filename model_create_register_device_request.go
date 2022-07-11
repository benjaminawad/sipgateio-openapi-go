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

// CreateRegisterDeviceRequest struct for CreateRegisterDeviceRequest
type CreateRegisterDeviceRequest struct {
	Alias *string `json:"alias,omitempty"`
}

// NewCreateRegisterDeviceRequest instantiates a new CreateRegisterDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRegisterDeviceRequest() *CreateRegisterDeviceRequest {
	this := CreateRegisterDeviceRequest{}
	return &this
}

// NewCreateRegisterDeviceRequestWithDefaults instantiates a new CreateRegisterDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRegisterDeviceRequestWithDefaults() *CreateRegisterDeviceRequest {
	this := CreateRegisterDeviceRequest{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *CreateRegisterDeviceRequest) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegisterDeviceRequest) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *CreateRegisterDeviceRequest) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *CreateRegisterDeviceRequest) SetAlias(v string) {
	o.Alias = &v
}

func (o CreateRegisterDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRegisterDeviceRequest struct {
	value *CreateRegisterDeviceRequest
	isSet bool
}

func (v NullableCreateRegisterDeviceRequest) Get() *CreateRegisterDeviceRequest {
	return v.value
}

func (v *NullableCreateRegisterDeviceRequest) Set(val *CreateRegisterDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRegisterDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRegisterDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRegisterDeviceRequest(val *CreateRegisterDeviceRequest) *NullableCreateRegisterDeviceRequest {
	return &NullableCreateRegisterDeviceRequest{value: val, isSet: true}
}

func (v NullableCreateRegisterDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRegisterDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


