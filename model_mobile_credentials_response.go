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

// MobileCredentialsResponse struct for MobileCredentialsResponse
type MobileCredentialsResponse struct {
	SimId *string `json:"simId,omitempty"`
	Puk1 *string `json:"puk1,omitempty"`
	Puk2 *string `json:"puk2,omitempty"`
}

// NewMobileCredentialsResponse instantiates a new MobileCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileCredentialsResponse() *MobileCredentialsResponse {
	this := MobileCredentialsResponse{}
	return &this
}

// NewMobileCredentialsResponseWithDefaults instantiates a new MobileCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileCredentialsResponseWithDefaults() *MobileCredentialsResponse {
	this := MobileCredentialsResponse{}
	return &this
}

// GetSimId returns the SimId field value if set, zero value otherwise.
func (o *MobileCredentialsResponse) GetSimId() string {
	if o == nil || o.SimId == nil {
		var ret string
		return ret
	}
	return *o.SimId
}

// GetSimIdOk returns a tuple with the SimId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileCredentialsResponse) GetSimIdOk() (*string, bool) {
	if o == nil || o.SimId == nil {
		return nil, false
	}
	return o.SimId, true
}

// HasSimId returns a boolean if a field has been set.
func (o *MobileCredentialsResponse) HasSimId() bool {
	if o != nil && o.SimId != nil {
		return true
	}

	return false
}

// SetSimId gets a reference to the given string and assigns it to the SimId field.
func (o *MobileCredentialsResponse) SetSimId(v string) {
	o.SimId = &v
}

// GetPuk1 returns the Puk1 field value if set, zero value otherwise.
func (o *MobileCredentialsResponse) GetPuk1() string {
	if o == nil || o.Puk1 == nil {
		var ret string
		return ret
	}
	return *o.Puk1
}

// GetPuk1Ok returns a tuple with the Puk1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileCredentialsResponse) GetPuk1Ok() (*string, bool) {
	if o == nil || o.Puk1 == nil {
		return nil, false
	}
	return o.Puk1, true
}

// HasPuk1 returns a boolean if a field has been set.
func (o *MobileCredentialsResponse) HasPuk1() bool {
	if o != nil && o.Puk1 != nil {
		return true
	}

	return false
}

// SetPuk1 gets a reference to the given string and assigns it to the Puk1 field.
func (o *MobileCredentialsResponse) SetPuk1(v string) {
	o.Puk1 = &v
}

// GetPuk2 returns the Puk2 field value if set, zero value otherwise.
func (o *MobileCredentialsResponse) GetPuk2() string {
	if o == nil || o.Puk2 == nil {
		var ret string
		return ret
	}
	return *o.Puk2
}

// GetPuk2Ok returns a tuple with the Puk2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileCredentialsResponse) GetPuk2Ok() (*string, bool) {
	if o == nil || o.Puk2 == nil {
		return nil, false
	}
	return o.Puk2, true
}

// HasPuk2 returns a boolean if a field has been set.
func (o *MobileCredentialsResponse) HasPuk2() bool {
	if o != nil && o.Puk2 != nil {
		return true
	}

	return false
}

// SetPuk2 gets a reference to the given string and assigns it to the Puk2 field.
func (o *MobileCredentialsResponse) SetPuk2(v string) {
	o.Puk2 = &v
}

func (o MobileCredentialsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SimId != nil {
		toSerialize["simId"] = o.SimId
	}
	if o.Puk1 != nil {
		toSerialize["puk1"] = o.Puk1
	}
	if o.Puk2 != nil {
		toSerialize["puk2"] = o.Puk2
	}
	return json.Marshal(toSerialize)
}

type NullableMobileCredentialsResponse struct {
	value *MobileCredentialsResponse
	isSet bool
}

func (v NullableMobileCredentialsResponse) Get() *MobileCredentialsResponse {
	return v.value
}

func (v *NullableMobileCredentialsResponse) Set(val *MobileCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileCredentialsResponse(val *MobileCredentialsResponse) *NullableMobileCredentialsResponse {
	return &NullableMobileCredentialsResponse{value: val, isSet: true}
}

func (v NullableMobileCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


