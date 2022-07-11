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

// Forwarding struct for Forwarding
type Forwarding struct {
	Destination *string `json:"destination,omitempty"`
	Timeout *int32 `json:"timeout,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewForwarding instantiates a new Forwarding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForwarding() *Forwarding {
	this := Forwarding{}
	var active bool = false
	this.Active = &active
	return &this
}

// NewForwardingWithDefaults instantiates a new Forwarding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForwardingWithDefaults() *Forwarding {
	this := Forwarding{}
	var active bool = false
	this.Active = &active
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *Forwarding) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forwarding) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *Forwarding) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *Forwarding) SetDestination(v string) {
	o.Destination = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *Forwarding) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forwarding) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *Forwarding) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *Forwarding) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Forwarding) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forwarding) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Forwarding) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Forwarding) SetActive(v bool) {
	o.Active = &v
}

func (o Forwarding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableForwarding struct {
	value *Forwarding
	isSet bool
}

func (v NullableForwarding) Get() *Forwarding {
	return v.value
}

func (v *NullableForwarding) Set(val *Forwarding) {
	v.value = val
	v.isSet = true
}

func (v NullableForwarding) IsSet() bool {
	return v.isSet
}

func (v *NullableForwarding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForwarding(val *Forwarding) *NullableForwarding {
	return &NullableForwarding{value: val, isSet: true}
}

func (v NullableForwarding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForwarding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


