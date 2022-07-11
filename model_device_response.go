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

// DeviceResponse struct for DeviceResponse
type DeviceResponse struct {
	Id *string `json:"id,omitempty"`
	Alias *string `json:"alias,omitempty"`
	Type *string `json:"type,omitempty"`
	Online *bool `json:"online,omitempty"`
	Dnd *bool `json:"dnd,omitempty"`
	ActivePhonelines []ActiveRouting `json:"activePhonelines,omitempty"`
	ActiveGroups []ActiveRouting `json:"activeGroups,omitempty"`
}

// NewDeviceResponse instantiates a new DeviceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceResponse() *DeviceResponse {
	this := DeviceResponse{}
	var online bool = false
	this.Online = &online
	var dnd bool = false
	this.Dnd = &dnd
	return &this
}

// NewDeviceResponseWithDefaults instantiates a new DeviceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceResponseWithDefaults() *DeviceResponse {
	this := DeviceResponse{}
	var online bool = false
	this.Online = &online
	var dnd bool = false
	this.Dnd = &dnd
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceResponse) SetId(v string) {
	o.Id = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *DeviceResponse) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *DeviceResponse) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *DeviceResponse) SetAlias(v string) {
	o.Alias = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeviceResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeviceResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeviceResponse) SetType(v string) {
	o.Type = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *DeviceResponse) GetOnline() bool {
	if o == nil || o.Online == nil {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetOnlineOk() (*bool, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *DeviceResponse) HasOnline() bool {
	if o != nil && o.Online != nil {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *DeviceResponse) SetOnline(v bool) {
	o.Online = &v
}

// GetDnd returns the Dnd field value if set, zero value otherwise.
func (o *DeviceResponse) GetDnd() bool {
	if o == nil || o.Dnd == nil {
		var ret bool
		return ret
	}
	return *o.Dnd
}

// GetDndOk returns a tuple with the Dnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetDndOk() (*bool, bool) {
	if o == nil || o.Dnd == nil {
		return nil, false
	}
	return o.Dnd, true
}

// HasDnd returns a boolean if a field has been set.
func (o *DeviceResponse) HasDnd() bool {
	if o != nil && o.Dnd != nil {
		return true
	}

	return false
}

// SetDnd gets a reference to the given bool and assigns it to the Dnd field.
func (o *DeviceResponse) SetDnd(v bool) {
	o.Dnd = &v
}

// GetActivePhonelines returns the ActivePhonelines field value if set, zero value otherwise.
func (o *DeviceResponse) GetActivePhonelines() []ActiveRouting {
	if o == nil || o.ActivePhonelines == nil {
		var ret []ActiveRouting
		return ret
	}
	return o.ActivePhonelines
}

// GetActivePhonelinesOk returns a tuple with the ActivePhonelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetActivePhonelinesOk() ([]ActiveRouting, bool) {
	if o == nil || o.ActivePhonelines == nil {
		return nil, false
	}
	return o.ActivePhonelines, true
}

// HasActivePhonelines returns a boolean if a field has been set.
func (o *DeviceResponse) HasActivePhonelines() bool {
	if o != nil && o.ActivePhonelines != nil {
		return true
	}

	return false
}

// SetActivePhonelines gets a reference to the given []ActiveRouting and assigns it to the ActivePhonelines field.
func (o *DeviceResponse) SetActivePhonelines(v []ActiveRouting) {
	o.ActivePhonelines = v
}

// GetActiveGroups returns the ActiveGroups field value if set, zero value otherwise.
func (o *DeviceResponse) GetActiveGroups() []ActiveRouting {
	if o == nil || o.ActiveGroups == nil {
		var ret []ActiveRouting
		return ret
	}
	return o.ActiveGroups
}

// GetActiveGroupsOk returns a tuple with the ActiveGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceResponse) GetActiveGroupsOk() ([]ActiveRouting, bool) {
	if o == nil || o.ActiveGroups == nil {
		return nil, false
	}
	return o.ActiveGroups, true
}

// HasActiveGroups returns a boolean if a field has been set.
func (o *DeviceResponse) HasActiveGroups() bool {
	if o != nil && o.ActiveGroups != nil {
		return true
	}

	return false
}

// SetActiveGroups gets a reference to the given []ActiveRouting and assigns it to the ActiveGroups field.
func (o *DeviceResponse) SetActiveGroups(v []ActiveRouting) {
	o.ActiveGroups = v
}

func (o DeviceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Online != nil {
		toSerialize["online"] = o.Online
	}
	if o.Dnd != nil {
		toSerialize["dnd"] = o.Dnd
	}
	if o.ActivePhonelines != nil {
		toSerialize["activePhonelines"] = o.ActivePhonelines
	}
	if o.ActiveGroups != nil {
		toSerialize["activeGroups"] = o.ActiveGroups
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceResponse struct {
	value *DeviceResponse
	isSet bool
}

func (v NullableDeviceResponse) Get() *DeviceResponse {
	return v.value
}

func (v *NullableDeviceResponse) Set(val *DeviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceResponse(val *DeviceResponse) *NullableDeviceResponse {
	return &NullableDeviceResponse{value: val, isSet: true}
}

func (v NullableDeviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


