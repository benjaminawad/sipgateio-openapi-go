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

// SimpleGroupResponse struct for SimpleGroupResponse
type SimpleGroupResponse struct {
	Id *string `json:"id,omitempty"`
	Alias *string `json:"alias,omitempty"`
	GroupNumbersUrl *string `json:"groupNumbersUrl,omitempty"`
	GroupUsersUrl *string `json:"groupUsersUrl,omitempty"`
	DepartmentId *int64 `json:"departmentId,omitempty"`
}

// NewSimpleGroupResponse instantiates a new SimpleGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleGroupResponse() *SimpleGroupResponse {
	this := SimpleGroupResponse{}
	return &this
}

// NewSimpleGroupResponseWithDefaults instantiates a new SimpleGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleGroupResponseWithDefaults() *SimpleGroupResponse {
	this := SimpleGroupResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimpleGroupResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleGroupResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimpleGroupResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SimpleGroupResponse) SetId(v string) {
	o.Id = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *SimpleGroupResponse) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleGroupResponse) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *SimpleGroupResponse) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *SimpleGroupResponse) SetAlias(v string) {
	o.Alias = &v
}

// GetGroupNumbersUrl returns the GroupNumbersUrl field value if set, zero value otherwise.
func (o *SimpleGroupResponse) GetGroupNumbersUrl() string {
	if o == nil || o.GroupNumbersUrl == nil {
		var ret string
		return ret
	}
	return *o.GroupNumbersUrl
}

// GetGroupNumbersUrlOk returns a tuple with the GroupNumbersUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleGroupResponse) GetGroupNumbersUrlOk() (*string, bool) {
	if o == nil || o.GroupNumbersUrl == nil {
		return nil, false
	}
	return o.GroupNumbersUrl, true
}

// HasGroupNumbersUrl returns a boolean if a field has been set.
func (o *SimpleGroupResponse) HasGroupNumbersUrl() bool {
	if o != nil && o.GroupNumbersUrl != nil {
		return true
	}

	return false
}

// SetGroupNumbersUrl gets a reference to the given string and assigns it to the GroupNumbersUrl field.
func (o *SimpleGroupResponse) SetGroupNumbersUrl(v string) {
	o.GroupNumbersUrl = &v
}

// GetGroupUsersUrl returns the GroupUsersUrl field value if set, zero value otherwise.
func (o *SimpleGroupResponse) GetGroupUsersUrl() string {
	if o == nil || o.GroupUsersUrl == nil {
		var ret string
		return ret
	}
	return *o.GroupUsersUrl
}

// GetGroupUsersUrlOk returns a tuple with the GroupUsersUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleGroupResponse) GetGroupUsersUrlOk() (*string, bool) {
	if o == nil || o.GroupUsersUrl == nil {
		return nil, false
	}
	return o.GroupUsersUrl, true
}

// HasGroupUsersUrl returns a boolean if a field has been set.
func (o *SimpleGroupResponse) HasGroupUsersUrl() bool {
	if o != nil && o.GroupUsersUrl != nil {
		return true
	}

	return false
}

// SetGroupUsersUrl gets a reference to the given string and assigns it to the GroupUsersUrl field.
func (o *SimpleGroupResponse) SetGroupUsersUrl(v string) {
	o.GroupUsersUrl = &v
}

// GetDepartmentId returns the DepartmentId field value if set, zero value otherwise.
func (o *SimpleGroupResponse) GetDepartmentId() int64 {
	if o == nil || o.DepartmentId == nil {
		var ret int64
		return ret
	}
	return *o.DepartmentId
}

// GetDepartmentIdOk returns a tuple with the DepartmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleGroupResponse) GetDepartmentIdOk() (*int64, bool) {
	if o == nil || o.DepartmentId == nil {
		return nil, false
	}
	return o.DepartmentId, true
}

// HasDepartmentId returns a boolean if a field has been set.
func (o *SimpleGroupResponse) HasDepartmentId() bool {
	if o != nil && o.DepartmentId != nil {
		return true
	}

	return false
}

// SetDepartmentId gets a reference to the given int64 and assigns it to the DepartmentId field.
func (o *SimpleGroupResponse) SetDepartmentId(v int64) {
	o.DepartmentId = &v
}

func (o SimpleGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.GroupNumbersUrl != nil {
		toSerialize["groupNumbersUrl"] = o.GroupNumbersUrl
	}
	if o.GroupUsersUrl != nil {
		toSerialize["groupUsersUrl"] = o.GroupUsersUrl
	}
	if o.DepartmentId != nil {
		toSerialize["departmentId"] = o.DepartmentId
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleGroupResponse struct {
	value *SimpleGroupResponse
	isSet bool
}

func (v NullableSimpleGroupResponse) Get() *SimpleGroupResponse {
	return v.value
}

func (v *NullableSimpleGroupResponse) Set(val *SimpleGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleGroupResponse(val *SimpleGroupResponse) *NullableSimpleGroupResponse {
	return &NullableSimpleGroupResponse{value: val, isSet: true}
}

func (v NullableSimpleGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

