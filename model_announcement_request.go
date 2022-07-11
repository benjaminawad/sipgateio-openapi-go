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

// AnnouncementRequest struct for AnnouncementRequest
type AnnouncementRequest struct {
	// The URL pointing to the announcement WAV file.
	Url string `json:"url"`
}

// NewAnnouncementRequest instantiates a new AnnouncementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnouncementRequest(url string) *AnnouncementRequest {
	this := AnnouncementRequest{}
	this.Url = url
	return &this
}

// NewAnnouncementRequestWithDefaults instantiates a new AnnouncementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnouncementRequestWithDefaults() *AnnouncementRequest {
	this := AnnouncementRequest{}
	return &this
}

// GetUrl returns the Url field value
func (o *AnnouncementRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AnnouncementRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AnnouncementRequest) SetUrl(v string) {
	o.Url = v
}

func (o AnnouncementRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableAnnouncementRequest struct {
	value *AnnouncementRequest
	isSet bool
}

func (v NullableAnnouncementRequest) Get() *AnnouncementRequest {
	return v.value
}

func (v *NullableAnnouncementRequest) Set(val *AnnouncementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnouncementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnouncementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnouncementRequest(val *AnnouncementRequest) *NullableAnnouncementRequest {
	return &NullableAnnouncementRequest{value: val, isSet: true}
}

func (v NullableAnnouncementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnouncementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


