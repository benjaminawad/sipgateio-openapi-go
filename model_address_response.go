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

// AddressResponse struct for AddressResponse
type AddressResponse struct {
	PoBox *string `json:"poBox,omitempty"`
	ExtendedAddress *string `json:"extendedAddress,omitempty"`
	StreetAddress *string `json:"streetAddress,omitempty"`
	Locality *string `json:"locality,omitempty"`
	Region *string `json:"region,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	Country *string `json:"country,omitempty"`
}

// NewAddressResponse instantiates a new AddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressResponse() *AddressResponse {
	this := AddressResponse{}
	return &this
}

// NewAddressResponseWithDefaults instantiates a new AddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressResponseWithDefaults() *AddressResponse {
	this := AddressResponse{}
	return &this
}

// GetPoBox returns the PoBox field value if set, zero value otherwise.
func (o *AddressResponse) GetPoBox() string {
	if o == nil || o.PoBox == nil {
		var ret string
		return ret
	}
	return *o.PoBox
}

// GetPoBoxOk returns a tuple with the PoBox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponse) GetPoBoxOk() (*string, bool) {
	if o == nil || o.PoBox == nil {
		return nil, false
	}
	return o.PoBox, true
}

// HasPoBox returns a boolean if a field has been set.
func (o *AddressResponse) HasPoBox() bool {
	if o != nil && o.PoBox != nil {
		return true
	}

	return false
}

// SetPoBox gets a reference to the given string and assigns it to the PoBox field.
func (o *AddressResponse) SetPoBox(v string) {
	o.PoBox = &v
}

// GetExtendedAddress returns the ExtendedAddress field value if set, zero value otherwise.
func (o *AddressResponse) GetExtendedAddress() string {
	if o == nil || o.ExtendedAddress == nil {
		var ret string
		return ret
	}
	return *o.ExtendedAddress
}

// GetExtendedAddressOk returns a tuple with the ExtendedAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponse) GetExtendedAddressOk() (*string, bool) {
	if o == nil || o.ExtendedAddress == nil {
		return nil, false
	}
	return o.ExtendedAddress, true
}

// HasExtendedAddress returns a boolean if a field has been set.
func (o *AddressResponse) HasExtendedAddress() bool {
	if o != nil && o.ExtendedAddress != nil {
		return true
	}

	return false
}

// SetExtendedAddress gets a reference to the given string and assigns it to the ExtendedAddress field.
func (o *AddressResponse) SetExtendedAddress(v string) {
	o.ExtendedAddress = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *AddressResponse) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponse) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *AddressResponse) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *AddressResponse) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *AddressResponse) GetLocality() string {
	if o == nil || o.Locality == nil {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponse) GetLocalityOk() (*string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *AddressResponse) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *AddressResponse) SetLocality(v string) {
	o.Locality = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AddressResponse) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponse) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AddressResponse) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AddressResponse) SetRegion(v string) {
	o.Region = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AddressResponse) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponse) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressResponse) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AddressResponse) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AddressResponse) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressResponse) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressResponse) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AddressResponse) SetCountry(v string) {
	o.Country = &v
}

func (o AddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PoBox != nil {
		toSerialize["poBox"] = o.PoBox
	}
	if o.ExtendedAddress != nil {
		toSerialize["extendedAddress"] = o.ExtendedAddress
	}
	if o.StreetAddress != nil {
		toSerialize["streetAddress"] = o.StreetAddress
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.PostalCode != nil {
		toSerialize["postalCode"] = o.PostalCode
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	return json.Marshal(toSerialize)
}

type NullableAddressResponse struct {
	value *AddressResponse
	isSet bool
}

func (v NullableAddressResponse) Get() *AddressResponse {
	return v.value
}

func (v *NullableAddressResponse) Set(val *AddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressResponse(val *AddressResponse) *NullableAddressResponse {
	return &NullableAddressResponse{value: val, isSet: true}
}

func (v NullableAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


