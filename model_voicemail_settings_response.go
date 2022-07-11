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

// VoicemailSettingsResponse struct for VoicemailSettingsResponse
type VoicemailSettingsResponse struct {
	Id *string `json:"id,omitempty"`
	Timeout *int32 `json:"timeout,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewVoicemailSettingsResponse instantiates a new VoicemailSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoicemailSettingsResponse() *VoicemailSettingsResponse {
	this := VoicemailSettingsResponse{}
	var active bool = false
	this.Active = &active
	return &this
}

// NewVoicemailSettingsResponseWithDefaults instantiates a new VoicemailSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoicemailSettingsResponseWithDefaults() *VoicemailSettingsResponse {
	this := VoicemailSettingsResponse{}
	var active bool = false
	this.Active = &active
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VoicemailSettingsResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoicemailSettingsResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VoicemailSettingsResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VoicemailSettingsResponse) SetId(v string) {
	o.Id = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *VoicemailSettingsResponse) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoicemailSettingsResponse) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *VoicemailSettingsResponse) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *VoicemailSettingsResponse) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *VoicemailSettingsResponse) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoicemailSettingsResponse) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *VoicemailSettingsResponse) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *VoicemailSettingsResponse) SetActive(v bool) {
	o.Active = &v
}

func (o VoicemailSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableVoicemailSettingsResponse struct {
	value *VoicemailSettingsResponse
	isSet bool
}

func (v NullableVoicemailSettingsResponse) Get() *VoicemailSettingsResponse {
	return v.value
}

func (v *NullableVoicemailSettingsResponse) Set(val *VoicemailSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVoicemailSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVoicemailSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoicemailSettingsResponse(val *VoicemailSettingsResponse) *NullableVoicemailSettingsResponse {
	return &NullableVoicemailSettingsResponse{value: val, isSet: true}
}

func (v NullableVoicemailSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoicemailSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


