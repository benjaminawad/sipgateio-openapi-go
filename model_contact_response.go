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

// ContactResponse struct for ContactResponse
type ContactResponse struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Picture *string `json:"picture,omitempty"`
	Emails []EMailResponse `json:"emails,omitempty"`
	Numbers []NumberResponse `json:"numbers,omitempty"`
	Addresses []AddressResponse `json:"addresses,omitempty"`
	Organization [][]string `json:"organization,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

// NewContactResponse instantiates a new ContactResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactResponse() *ContactResponse {
	this := ContactResponse{}
	return &this
}

// NewContactResponseWithDefaults instantiates a new ContactResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactResponseWithDefaults() *ContactResponse {
	this := ContactResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContactResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContactResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContactResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContactResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContactResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContactResponse) SetName(v string) {
	o.Name = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *ContactResponse) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *ContactResponse) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *ContactResponse) SetPicture(v string) {
	o.Picture = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *ContactResponse) GetEmails() []EMailResponse {
	if o == nil || o.Emails == nil {
		var ret []EMailResponse
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetEmailsOk() ([]EMailResponse, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *ContactResponse) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []EMailResponse and assigns it to the Emails field.
func (o *ContactResponse) SetEmails(v []EMailResponse) {
	o.Emails = v
}

// GetNumbers returns the Numbers field value if set, zero value otherwise.
func (o *ContactResponse) GetNumbers() []NumberResponse {
	if o == nil || o.Numbers == nil {
		var ret []NumberResponse
		return ret
	}
	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetNumbersOk() ([]NumberResponse, bool) {
	if o == nil || o.Numbers == nil {
		return nil, false
	}
	return o.Numbers, true
}

// HasNumbers returns a boolean if a field has been set.
func (o *ContactResponse) HasNumbers() bool {
	if o != nil && o.Numbers != nil {
		return true
	}

	return false
}

// SetNumbers gets a reference to the given []NumberResponse and assigns it to the Numbers field.
func (o *ContactResponse) SetNumbers(v []NumberResponse) {
	o.Numbers = v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ContactResponse) GetAddresses() []AddressResponse {
	if o == nil || o.Addresses == nil {
		var ret []AddressResponse
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetAddressesOk() ([]AddressResponse, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ContactResponse) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressResponse and assigns it to the Addresses field.
func (o *ContactResponse) SetAddresses(v []AddressResponse) {
	o.Addresses = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ContactResponse) GetOrganization() [][]string {
	if o == nil || o.Organization == nil {
		var ret [][]string
		return ret
	}
	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetOrganizationOk() ([][]string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ContactResponse) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given [][]string and assigns it to the Organization field.
func (o *ContactResponse) SetOrganization(v [][]string) {
	o.Organization = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ContactResponse) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ContactResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ContactResponse) SetScope(v string) {
	o.Scope = &v
}

func (o ContactResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.Numbers != nil {
		toSerialize["numbers"] = o.Numbers
	}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableContactResponse struct {
	value *ContactResponse
	isSet bool
}

func (v NullableContactResponse) Get() *ContactResponse {
	return v.value
}

func (v *NullableContactResponse) Set(val *ContactResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactResponse(val *ContactResponse) *NullableContactResponse {
	return &NullableContactResponse{value: val, isSet: true}
}

func (v NullableContactResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


