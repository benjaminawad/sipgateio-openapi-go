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

// NumberResponseWithEndpoint struct for NumberResponseWithEndpoint
type NumberResponseWithEndpoint struct {
	Id *string `json:"id,omitempty"`
	Number *string `json:"number,omitempty"`
	Localized *string `json:"localized,omitempty"`
	Type *string `json:"type,omitempty"`
	EndpointId *string `json:"endpointId,omitempty"`
	EndpointAlias *string `json:"endpointAlias,omitempty"`
	EndpointUrl *string `json:"endpointUrl,omitempty"`
}

// NewNumberResponseWithEndpoint instantiates a new NumberResponseWithEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberResponseWithEndpoint() *NumberResponseWithEndpoint {
	this := NumberResponseWithEndpoint{}
	return &this
}

// NewNumberResponseWithEndpointWithDefaults instantiates a new NumberResponseWithEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberResponseWithEndpointWithDefaults() *NumberResponseWithEndpoint {
	this := NumberResponseWithEndpoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NumberResponseWithEndpoint) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberResponseWithEndpoint) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NumberResponseWithEndpoint) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NumberResponseWithEndpoint) SetId(v string) {
	o.Id = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *NumberResponseWithEndpoint) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberResponseWithEndpoint) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *NumberResponseWithEndpoint) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *NumberResponseWithEndpoint) SetNumber(v string) {
	o.Number = &v
}

// GetLocalized returns the Localized field value if set, zero value otherwise.
func (o *NumberResponseWithEndpoint) GetLocalized() string {
	if o == nil || o.Localized == nil {
		var ret string
		return ret
	}
	return *o.Localized
}

// GetLocalizedOk returns a tuple with the Localized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberResponseWithEndpoint) GetLocalizedOk() (*string, bool) {
	if o == nil || o.Localized == nil {
		return nil, false
	}
	return o.Localized, true
}

// HasLocalized returns a boolean if a field has been set.
func (o *NumberResponseWithEndpoint) HasLocalized() bool {
	if o != nil && o.Localized != nil {
		return true
	}

	return false
}

// SetLocalized gets a reference to the given string and assigns it to the Localized field.
func (o *NumberResponseWithEndpoint) SetLocalized(v string) {
	o.Localized = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NumberResponseWithEndpoint) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberResponseWithEndpoint) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NumberResponseWithEndpoint) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NumberResponseWithEndpoint) SetType(v string) {
	o.Type = &v
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *NumberResponseWithEndpoint) GetEndpointId() string {
	if o == nil || o.EndpointId == nil {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberResponseWithEndpoint) GetEndpointIdOk() (*string, bool) {
	if o == nil || o.EndpointId == nil {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *NumberResponseWithEndpoint) HasEndpointId() bool {
	if o != nil && o.EndpointId != nil {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *NumberResponseWithEndpoint) SetEndpointId(v string) {
	o.EndpointId = &v
}

// GetEndpointAlias returns the EndpointAlias field value if set, zero value otherwise.
func (o *NumberResponseWithEndpoint) GetEndpointAlias() string {
	if o == nil || o.EndpointAlias == nil {
		var ret string
		return ret
	}
	return *o.EndpointAlias
}

// GetEndpointAliasOk returns a tuple with the EndpointAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberResponseWithEndpoint) GetEndpointAliasOk() (*string, bool) {
	if o == nil || o.EndpointAlias == nil {
		return nil, false
	}
	return o.EndpointAlias, true
}

// HasEndpointAlias returns a boolean if a field has been set.
func (o *NumberResponseWithEndpoint) HasEndpointAlias() bool {
	if o != nil && o.EndpointAlias != nil {
		return true
	}

	return false
}

// SetEndpointAlias gets a reference to the given string and assigns it to the EndpointAlias field.
func (o *NumberResponseWithEndpoint) SetEndpointAlias(v string) {
	o.EndpointAlias = &v
}

// GetEndpointUrl returns the EndpointUrl field value if set, zero value otherwise.
func (o *NumberResponseWithEndpoint) GetEndpointUrl() string {
	if o == nil || o.EndpointUrl == nil {
		var ret string
		return ret
	}
	return *o.EndpointUrl
}

// GetEndpointUrlOk returns a tuple with the EndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberResponseWithEndpoint) GetEndpointUrlOk() (*string, bool) {
	if o == nil || o.EndpointUrl == nil {
		return nil, false
	}
	return o.EndpointUrl, true
}

// HasEndpointUrl returns a boolean if a field has been set.
func (o *NumberResponseWithEndpoint) HasEndpointUrl() bool {
	if o != nil && o.EndpointUrl != nil {
		return true
	}

	return false
}

// SetEndpointUrl gets a reference to the given string and assigns it to the EndpointUrl field.
func (o *NumberResponseWithEndpoint) SetEndpointUrl(v string) {
	o.EndpointUrl = &v
}

func (o NumberResponseWithEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Localized != nil {
		toSerialize["localized"] = o.Localized
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.EndpointId != nil {
		toSerialize["endpointId"] = o.EndpointId
	}
	if o.EndpointAlias != nil {
		toSerialize["endpointAlias"] = o.EndpointAlias
	}
	if o.EndpointUrl != nil {
		toSerialize["endpointUrl"] = o.EndpointUrl
	}
	return json.Marshal(toSerialize)
}

type NullableNumberResponseWithEndpoint struct {
	value *NumberResponseWithEndpoint
	isSet bool
}

func (v NullableNumberResponseWithEndpoint) Get() *NumberResponseWithEndpoint {
	return v.value
}

func (v *NullableNumberResponseWithEndpoint) Set(val *NumberResponseWithEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberResponseWithEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberResponseWithEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberResponseWithEndpoint(val *NumberResponseWithEndpoint) *NullableNumberResponseWithEndpoint {
	return &NullableNumberResponseWithEndpoint{value: val, isSet: true}
}

func (v NullableNumberResponseWithEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberResponseWithEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


