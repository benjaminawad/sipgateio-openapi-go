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

// RegisterDeviceResponse struct for RegisterDeviceResponse
type RegisterDeviceResponse struct {
	Id *string `json:"id,omitempty"`
	Alias *string `json:"alias,omitempty"`
	Type *string `json:"type,omitempty"`
	Online *bool `json:"online,omitempty"`
	Dnd *bool `json:"dnd,omitempty"`
	ActivePhonelines []ActiveRouting `json:"activePhonelines,omitempty"`
	ActiveGroups []ActiveRouting `json:"activeGroups,omitempty"`
	Credentials *CredentialsResponse `json:"credentials,omitempty"`
	Registered []RegisteredDevice `json:"registered,omitempty"`
	EmergencyAddressId *string `json:"emergencyAddressId,omitempty"`
	AddressUrl *string `json:"addressUrl,omitempty"`
}

// NewRegisterDeviceResponse instantiates a new RegisterDeviceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterDeviceResponse() *RegisterDeviceResponse {
	this := RegisterDeviceResponse{}
	var online bool = false
	this.Online = &online
	var dnd bool = false
	this.Dnd = &dnd
	return &this
}

// NewRegisterDeviceResponseWithDefaults instantiates a new RegisterDeviceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterDeviceResponseWithDefaults() *RegisterDeviceResponse {
	this := RegisterDeviceResponse{}
	var online bool = false
	this.Online = &online
	var dnd bool = false
	this.Dnd = &dnd
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegisterDeviceResponse) SetId(v string) {
	o.Id = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *RegisterDeviceResponse) SetAlias(v string) {
	o.Alias = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RegisterDeviceResponse) SetType(v string) {
	o.Type = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetOnline() bool {
	if o == nil || o.Online == nil {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetOnlineOk() (*bool, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasOnline() bool {
	if o != nil && o.Online != nil {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *RegisterDeviceResponse) SetOnline(v bool) {
	o.Online = &v
}

// GetDnd returns the Dnd field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetDnd() bool {
	if o == nil || o.Dnd == nil {
		var ret bool
		return ret
	}
	return *o.Dnd
}

// GetDndOk returns a tuple with the Dnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetDndOk() (*bool, bool) {
	if o == nil || o.Dnd == nil {
		return nil, false
	}
	return o.Dnd, true
}

// HasDnd returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasDnd() bool {
	if o != nil && o.Dnd != nil {
		return true
	}

	return false
}

// SetDnd gets a reference to the given bool and assigns it to the Dnd field.
func (o *RegisterDeviceResponse) SetDnd(v bool) {
	o.Dnd = &v
}

// GetActivePhonelines returns the ActivePhonelines field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetActivePhonelines() []ActiveRouting {
	if o == nil || o.ActivePhonelines == nil {
		var ret []ActiveRouting
		return ret
	}
	return o.ActivePhonelines
}

// GetActivePhonelinesOk returns a tuple with the ActivePhonelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetActivePhonelinesOk() ([]ActiveRouting, bool) {
	if o == nil || o.ActivePhonelines == nil {
		return nil, false
	}
	return o.ActivePhonelines, true
}

// HasActivePhonelines returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasActivePhonelines() bool {
	if o != nil && o.ActivePhonelines != nil {
		return true
	}

	return false
}

// SetActivePhonelines gets a reference to the given []ActiveRouting and assigns it to the ActivePhonelines field.
func (o *RegisterDeviceResponse) SetActivePhonelines(v []ActiveRouting) {
	o.ActivePhonelines = v
}

// GetActiveGroups returns the ActiveGroups field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetActiveGroups() []ActiveRouting {
	if o == nil || o.ActiveGroups == nil {
		var ret []ActiveRouting
		return ret
	}
	return o.ActiveGroups
}

// GetActiveGroupsOk returns a tuple with the ActiveGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetActiveGroupsOk() ([]ActiveRouting, bool) {
	if o == nil || o.ActiveGroups == nil {
		return nil, false
	}
	return o.ActiveGroups, true
}

// HasActiveGroups returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasActiveGroups() bool {
	if o != nil && o.ActiveGroups != nil {
		return true
	}

	return false
}

// SetActiveGroups gets a reference to the given []ActiveRouting and assigns it to the ActiveGroups field.
func (o *RegisterDeviceResponse) SetActiveGroups(v []ActiveRouting) {
	o.ActiveGroups = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetCredentials() CredentialsResponse {
	if o == nil || o.Credentials == nil {
		var ret CredentialsResponse
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetCredentialsOk() (*CredentialsResponse, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given CredentialsResponse and assigns it to the Credentials field.
func (o *RegisterDeviceResponse) SetCredentials(v CredentialsResponse) {
	o.Credentials = &v
}

// GetRegistered returns the Registered field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetRegistered() []RegisteredDevice {
	if o == nil || o.Registered == nil {
		var ret []RegisteredDevice
		return ret
	}
	return o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetRegisteredOk() ([]RegisteredDevice, bool) {
	if o == nil || o.Registered == nil {
		return nil, false
	}
	return o.Registered, true
}

// HasRegistered returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasRegistered() bool {
	if o != nil && o.Registered != nil {
		return true
	}

	return false
}

// SetRegistered gets a reference to the given []RegisteredDevice and assigns it to the Registered field.
func (o *RegisterDeviceResponse) SetRegistered(v []RegisteredDevice) {
	o.Registered = v
}

// GetEmergencyAddressId returns the EmergencyAddressId field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetEmergencyAddressId() string {
	if o == nil || o.EmergencyAddressId == nil {
		var ret string
		return ret
	}
	return *o.EmergencyAddressId
}

// GetEmergencyAddressIdOk returns a tuple with the EmergencyAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetEmergencyAddressIdOk() (*string, bool) {
	if o == nil || o.EmergencyAddressId == nil {
		return nil, false
	}
	return o.EmergencyAddressId, true
}

// HasEmergencyAddressId returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasEmergencyAddressId() bool {
	if o != nil && o.EmergencyAddressId != nil {
		return true
	}

	return false
}

// SetEmergencyAddressId gets a reference to the given string and assigns it to the EmergencyAddressId field.
func (o *RegisterDeviceResponse) SetEmergencyAddressId(v string) {
	o.EmergencyAddressId = &v
}

// GetAddressUrl returns the AddressUrl field value if set, zero value otherwise.
func (o *RegisterDeviceResponse) GetAddressUrl() string {
	if o == nil || o.AddressUrl == nil {
		var ret string
		return ret
	}
	return *o.AddressUrl
}

// GetAddressUrlOk returns a tuple with the AddressUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeviceResponse) GetAddressUrlOk() (*string, bool) {
	if o == nil || o.AddressUrl == nil {
		return nil, false
	}
	return o.AddressUrl, true
}

// HasAddressUrl returns a boolean if a field has been set.
func (o *RegisterDeviceResponse) HasAddressUrl() bool {
	if o != nil && o.AddressUrl != nil {
		return true
	}

	return false
}

// SetAddressUrl gets a reference to the given string and assigns it to the AddressUrl field.
func (o *RegisterDeviceResponse) SetAddressUrl(v string) {
	o.AddressUrl = &v
}

func (o RegisterDeviceResponse) MarshalJSON() ([]byte, error) {
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
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Registered != nil {
		toSerialize["registered"] = o.Registered
	}
	if o.EmergencyAddressId != nil {
		toSerialize["emergencyAddressId"] = o.EmergencyAddressId
	}
	if o.AddressUrl != nil {
		toSerialize["addressUrl"] = o.AddressUrl
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterDeviceResponse struct {
	value *RegisterDeviceResponse
	isSet bool
}

func (v NullableRegisterDeviceResponse) Get() *RegisterDeviceResponse {
	return v.value
}

func (v *NullableRegisterDeviceResponse) Set(val *RegisterDeviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterDeviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterDeviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterDeviceResponse(val *RegisterDeviceResponse) *NullableRegisterDeviceResponse {
	return &NullableRegisterDeviceResponse{value: val, isSet: true}
}

func (v NullableRegisterDeviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterDeviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


