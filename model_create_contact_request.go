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

// CreateContactRequest struct for CreateContactRequest
type CreateContactRequest struct {
	Name *string `json:"name,omitempty"`
	Family *string `json:"family,omitempty"`
	Given *string `json:"given,omitempty"`
	Picture *string `json:"picture,omitempty"`
	Emails []EmailRequest `json:"emails,omitempty"`
	Numbers []NumberRequest `json:"numbers,omitempty"`
	Addresses []AddressRequest `json:"addresses,omitempty"`
	Organization [][]string `json:"organization,omitempty"`
	Scope string `json:"scope"`
}

// NewCreateContactRequest instantiates a new CreateContactRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContactRequest(scope string) *CreateContactRequest {
	this := CreateContactRequest{}
	this.Scope = scope
	return &this
}

// NewCreateContactRequestWithDefaults instantiates a new CreateContactRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContactRequestWithDefaults() *CreateContactRequest {
	this := CreateContactRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateContactRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContactRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateContactRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateContactRequest) SetName(v string) {
	o.Name = &v
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *CreateContactRequest) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContactRequest) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *CreateContactRequest) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *CreateContactRequest) SetFamily(v string) {
	o.Family = &v
}

// GetGiven returns the Given field value if set, zero value otherwise.
func (o *CreateContactRequest) GetGiven() string {
	if o == nil || o.Given == nil {
		var ret string
		return ret
	}
	return *o.Given
}

// GetGivenOk returns a tuple with the Given field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContactRequest) GetGivenOk() (*string, bool) {
	if o == nil || o.Given == nil {
		return nil, false
	}
	return o.Given, true
}

// HasGiven returns a boolean if a field has been set.
func (o *CreateContactRequest) HasGiven() bool {
	if o != nil && o.Given != nil {
		return true
	}

	return false
}

// SetGiven gets a reference to the given string and assigns it to the Given field.
func (o *CreateContactRequest) SetGiven(v string) {
	o.Given = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *CreateContactRequest) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContactRequest) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *CreateContactRequest) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *CreateContactRequest) SetPicture(v string) {
	o.Picture = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *CreateContactRequest) GetEmails() []EmailRequest {
	if o == nil || o.Emails == nil {
		var ret []EmailRequest
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContactRequest) GetEmailsOk() ([]EmailRequest, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *CreateContactRequest) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []EmailRequest and assigns it to the Emails field.
func (o *CreateContactRequest) SetEmails(v []EmailRequest) {
	o.Emails = v
}

// GetNumbers returns the Numbers field value if set, zero value otherwise.
func (o *CreateContactRequest) GetNumbers() []NumberRequest {
	if o == nil || o.Numbers == nil {
		var ret []NumberRequest
		return ret
	}
	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContactRequest) GetNumbersOk() ([]NumberRequest, bool) {
	if o == nil || o.Numbers == nil {
		return nil, false
	}
	return o.Numbers, true
}

// HasNumbers returns a boolean if a field has been set.
func (o *CreateContactRequest) HasNumbers() bool {
	if o != nil && o.Numbers != nil {
		return true
	}

	return false
}

// SetNumbers gets a reference to the given []NumberRequest and assigns it to the Numbers field.
func (o *CreateContactRequest) SetNumbers(v []NumberRequest) {
	o.Numbers = v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *CreateContactRequest) GetAddresses() []AddressRequest {
	if o == nil || o.Addresses == nil {
		var ret []AddressRequest
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContactRequest) GetAddressesOk() ([]AddressRequest, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *CreateContactRequest) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressRequest and assigns it to the Addresses field.
func (o *CreateContactRequest) SetAddresses(v []AddressRequest) {
	o.Addresses = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CreateContactRequest) GetOrganization() [][]string {
	if o == nil || o.Organization == nil {
		var ret [][]string
		return ret
	}
	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContactRequest) GetOrganizationOk() ([][]string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CreateContactRequest) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given [][]string and assigns it to the Organization field.
func (o *CreateContactRequest) SetOrganization(v [][]string) {
	o.Organization = v
}

// GetScope returns the Scope field value
func (o *CreateContactRequest) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CreateContactRequest) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CreateContactRequest) SetScope(v string) {
	o.Scope = v
}

func (o CreateContactRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Given != nil {
		toSerialize["given"] = o.Given
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
	if true {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableCreateContactRequest struct {
	value *CreateContactRequest
	isSet bool
}

func (v NullableCreateContactRequest) Get() *CreateContactRequest {
	return v.value
}

func (v *NullableCreateContactRequest) Set(val *CreateContactRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContactRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContactRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContactRequest(val *CreateContactRequest) *NullableCreateContactRequest {
	return &NullableCreateContactRequest{value: val, isSet: true}
}

func (v NullableCreateContactRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContactRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


