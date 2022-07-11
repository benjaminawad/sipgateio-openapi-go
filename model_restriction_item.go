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

// RestrictionItem struct for RestrictionItem
type RestrictionItem struct {
	Restriction *string `json:"restriction,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
	Target *string `json:"target,omitempty"`
}

// NewRestrictionItem instantiates a new RestrictionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestrictionItem() *RestrictionItem {
	this := RestrictionItem{}
	return &this
}

// NewRestrictionItemWithDefaults instantiates a new RestrictionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictionItemWithDefaults() *RestrictionItem {
	this := RestrictionItem{}
	return &this
}

// GetRestriction returns the Restriction field value if set, zero value otherwise.
func (o *RestrictionItem) GetRestriction() string {
	if o == nil || o.Restriction == nil {
		var ret string
		return ret
	}
	return *o.Restriction
}

// GetRestrictionOk returns a tuple with the Restriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionItem) GetRestrictionOk() (*string, bool) {
	if o == nil || o.Restriction == nil {
		return nil, false
	}
	return o.Restriction, true
}

// HasRestriction returns a boolean if a field has been set.
func (o *RestrictionItem) HasRestriction() bool {
	if o != nil && o.Restriction != nil {
		return true
	}

	return false
}

// SetRestriction gets a reference to the given string and assigns it to the Restriction field.
func (o *RestrictionItem) SetRestriction(v string) {
	o.Restriction = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RestrictionItem) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionItem) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RestrictionItem) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *RestrictionItem) SetValue(v map[string]interface{}) {
	o.Value = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *RestrictionItem) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionItem) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *RestrictionItem) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *RestrictionItem) SetTarget(v string) {
	o.Target = &v
}

func (o RestrictionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Restriction != nil {
		toSerialize["restriction"] = o.Restriction
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableRestrictionItem struct {
	value *RestrictionItem
	isSet bool
}

func (v NullableRestrictionItem) Get() *RestrictionItem {
	return v.value
}

func (v *NullableRestrictionItem) Set(val *RestrictionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictionItem(val *RestrictionItem) *NullableRestrictionItem {
	return &NullableRestrictionItem{value: val, isSet: true}
}

func (v NullableRestrictionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


