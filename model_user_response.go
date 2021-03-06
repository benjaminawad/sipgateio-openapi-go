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

// UserResponse struct for UserResponse
type UserResponse struct {
	Id *string `json:"id,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Email *string `json:"email,omitempty"`
	DefaultDevice *string `json:"defaultDevice,omitempty"`
	BusyOnBusy *bool `json:"busyOnBusy,omitempty"`
	AddressId *string `json:"addressId,omitempty"`
	DirectDialIds []string `json:"directDialIds,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Admin *bool `json:"admin,omitempty"`
}

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse() *UserResponse {
	this := UserResponse{}
	var busyOnBusy bool = false
	this.BusyOnBusy = &busyOnBusy
	var admin bool = false
	this.Admin = &admin
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	var busyOnBusy bool = false
	this.BusyOnBusy = &busyOnBusy
	var admin bool = false
	this.Admin = &admin
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserResponse) SetId(v string) {
	o.Id = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *UserResponse) GetFirstname() string {
	if o == nil || o.Firstname == nil {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFirstnameOk() (*string, bool) {
	if o == nil || o.Firstname == nil {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *UserResponse) HasFirstname() bool {
	if o != nil && o.Firstname != nil {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *UserResponse) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *UserResponse) GetLastname() string {
	if o == nil || o.Lastname == nil {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLastnameOk() (*string, bool) {
	if o == nil || o.Lastname == nil {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *UserResponse) HasLastname() bool {
	if o != nil && o.Lastname != nil {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *UserResponse) SetLastname(v string) {
	o.Lastname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserResponse) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserResponse) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserResponse) SetEmail(v string) {
	o.Email = &v
}

// GetDefaultDevice returns the DefaultDevice field value if set, zero value otherwise.
func (o *UserResponse) GetDefaultDevice() string {
	if o == nil || o.DefaultDevice == nil {
		var ret string
		return ret
	}
	return *o.DefaultDevice
}

// GetDefaultDeviceOk returns a tuple with the DefaultDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetDefaultDeviceOk() (*string, bool) {
	if o == nil || o.DefaultDevice == nil {
		return nil, false
	}
	return o.DefaultDevice, true
}

// HasDefaultDevice returns a boolean if a field has been set.
func (o *UserResponse) HasDefaultDevice() bool {
	if o != nil && o.DefaultDevice != nil {
		return true
	}

	return false
}

// SetDefaultDevice gets a reference to the given string and assigns it to the DefaultDevice field.
func (o *UserResponse) SetDefaultDevice(v string) {
	o.DefaultDevice = &v
}

// GetBusyOnBusy returns the BusyOnBusy field value if set, zero value otherwise.
func (o *UserResponse) GetBusyOnBusy() bool {
	if o == nil || o.BusyOnBusy == nil {
		var ret bool
		return ret
	}
	return *o.BusyOnBusy
}

// GetBusyOnBusyOk returns a tuple with the BusyOnBusy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetBusyOnBusyOk() (*bool, bool) {
	if o == nil || o.BusyOnBusy == nil {
		return nil, false
	}
	return o.BusyOnBusy, true
}

// HasBusyOnBusy returns a boolean if a field has been set.
func (o *UserResponse) HasBusyOnBusy() bool {
	if o != nil && o.BusyOnBusy != nil {
		return true
	}

	return false
}

// SetBusyOnBusy gets a reference to the given bool and assigns it to the BusyOnBusy field.
func (o *UserResponse) SetBusyOnBusy(v bool) {
	o.BusyOnBusy = &v
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *UserResponse) GetAddressId() string {
	if o == nil || o.AddressId == nil {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetAddressIdOk() (*string, bool) {
	if o == nil || o.AddressId == nil {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *UserResponse) HasAddressId() bool {
	if o != nil && o.AddressId != nil {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *UserResponse) SetAddressId(v string) {
	o.AddressId = &v
}

// GetDirectDialIds returns the DirectDialIds field value if set, zero value otherwise.
func (o *UserResponse) GetDirectDialIds() []string {
	if o == nil || o.DirectDialIds == nil {
		var ret []string
		return ret
	}
	return o.DirectDialIds
}

// GetDirectDialIdsOk returns a tuple with the DirectDialIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetDirectDialIdsOk() ([]string, bool) {
	if o == nil || o.DirectDialIds == nil {
		return nil, false
	}
	return o.DirectDialIds, true
}

// HasDirectDialIds returns a boolean if a field has been set.
func (o *UserResponse) HasDirectDialIds() bool {
	if o != nil && o.DirectDialIds != nil {
		return true
	}

	return false
}

// SetDirectDialIds gets a reference to the given []string and assigns it to the DirectDialIds field.
func (o *UserResponse) SetDirectDialIds(v []string) {
	o.DirectDialIds = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *UserResponse) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UserResponse) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *UserResponse) SetTimezone(v string) {
	o.Timezone = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *UserResponse) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *UserResponse) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *UserResponse) SetAdmin(v bool) {
	o.Admin = &v
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Firstname != nil {
		toSerialize["firstname"] = o.Firstname
	}
	if o.Lastname != nil {
		toSerialize["lastname"] = o.Lastname
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.DefaultDevice != nil {
		toSerialize["defaultDevice"] = o.DefaultDevice
	}
	if o.BusyOnBusy != nil {
		toSerialize["busyOnBusy"] = o.BusyOnBusy
	}
	if o.AddressId != nil {
		toSerialize["addressId"] = o.AddressId
	}
	if o.DirectDialIds != nil {
		toSerialize["directDialIds"] = o.DirectDialIds
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	return json.Marshal(toSerialize)
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


