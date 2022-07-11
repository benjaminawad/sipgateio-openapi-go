# CreateContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Family** | Pointer to **string** |  | [optional] 
**Given** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]EmailRequest**](EmailRequest.md) |  | [optional] 
**Numbers** | Pointer to [**[]NumberRequest**](NumberRequest.md) |  | [optional] 
**Addresses** | Pointer to [**[]AddressRequest**](AddressRequest.md) |  | [optional] 
**Organization** | Pointer to **[][]string** |  | [optional] 
**Scope** | **string** |  | 

## Methods

### NewCreateContactRequest

`func NewCreateContactRequest(scope string, ) *CreateContactRequest`

NewCreateContactRequest instantiates a new CreateContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContactRequestWithDefaults

`func NewCreateContactRequestWithDefaults() *CreateContactRequest`

NewCreateContactRequestWithDefaults instantiates a new CreateContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateContactRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContactRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContactRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateContactRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFamily

`func (o *CreateContactRequest) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *CreateContactRequest) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *CreateContactRequest) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *CreateContactRequest) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetGiven

`func (o *CreateContactRequest) GetGiven() string`

GetGiven returns the Given field if non-nil, zero value otherwise.

### GetGivenOk

`func (o *CreateContactRequest) GetGivenOk() (*string, bool)`

GetGivenOk returns a tuple with the Given field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiven

`func (o *CreateContactRequest) SetGiven(v string)`

SetGiven sets Given field to given value.

### HasGiven

`func (o *CreateContactRequest) HasGiven() bool`

HasGiven returns a boolean if a field has been set.

### GetPicture

`func (o *CreateContactRequest) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *CreateContactRequest) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *CreateContactRequest) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *CreateContactRequest) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetEmails

`func (o *CreateContactRequest) GetEmails() []EmailRequest`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CreateContactRequest) GetEmailsOk() (*[]EmailRequest, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CreateContactRequest) SetEmails(v []EmailRequest)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *CreateContactRequest) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetNumbers

`func (o *CreateContactRequest) GetNumbers() []NumberRequest`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *CreateContactRequest) GetNumbersOk() (*[]NumberRequest, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *CreateContactRequest) SetNumbers(v []NumberRequest)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *CreateContactRequest) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetAddresses

`func (o *CreateContactRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CreateContactRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CreateContactRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *CreateContactRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetOrganization

`func (o *CreateContactRequest) GetOrganization() [][]string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CreateContactRequest) GetOrganizationOk() (*[][]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CreateContactRequest) SetOrganization(v [][]string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CreateContactRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetScope

`func (o *CreateContactRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateContactRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateContactRequest) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


