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

// PhonelineSipgateIoLogsResponse struct for PhonelineSipgateIoLogsResponse
type PhonelineSipgateIoLogsResponse struct {
	Items []PhonelineSipgateIoLogResponse `json:"items,omitempty"`
}

// NewPhonelineSipgateIoLogsResponse instantiates a new PhonelineSipgateIoLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhonelineSipgateIoLogsResponse() *PhonelineSipgateIoLogsResponse {
	this := PhonelineSipgateIoLogsResponse{}
	return &this
}

// NewPhonelineSipgateIoLogsResponseWithDefaults instantiates a new PhonelineSipgateIoLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhonelineSipgateIoLogsResponseWithDefaults() *PhonelineSipgateIoLogsResponse {
	this := PhonelineSipgateIoLogsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PhonelineSipgateIoLogsResponse) GetItems() []PhonelineSipgateIoLogResponse {
	if o == nil || o.Items == nil {
		var ret []PhonelineSipgateIoLogResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineSipgateIoLogsResponse) GetItemsOk() ([]PhonelineSipgateIoLogResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PhonelineSipgateIoLogsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PhonelineSipgateIoLogResponse and assigns it to the Items field.
func (o *PhonelineSipgateIoLogsResponse) SetItems(v []PhonelineSipgateIoLogResponse) {
	o.Items = v
}

func (o PhonelineSipgateIoLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullablePhonelineSipgateIoLogsResponse struct {
	value *PhonelineSipgateIoLogsResponse
	isSet bool
}

func (v NullablePhonelineSipgateIoLogsResponse) Get() *PhonelineSipgateIoLogsResponse {
	return v.value
}

func (v *NullablePhonelineSipgateIoLogsResponse) Set(val *PhonelineSipgateIoLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePhonelineSipgateIoLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePhonelineSipgateIoLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhonelineSipgateIoLogsResponse(val *PhonelineSipgateIoLogsResponse) *NullablePhonelineSipgateIoLogsResponse {
	return &NullablePhonelineSipgateIoLogsResponse{value: val, isSet: true}
}

func (v NullablePhonelineSipgateIoLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhonelineSipgateIoLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


