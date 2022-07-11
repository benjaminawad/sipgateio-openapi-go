# UpdateContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### NewUpdateContactRequest

`func NewUpdateContactRequest(scope string, ) *UpdateContactRequest`

NewUpdateContactRequest instantiates a new UpdateContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContactRequestWithDefaults

`func NewUpdateContactRequestWithDefaults() *UpdateContactRequest`

NewUpdateContactRequestWithDefaults instantiates a new UpdateContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateContactRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateContactRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateContactRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateContactRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateContactRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateContactRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateContactRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateContactRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFamily

`func (o *UpdateContactRequest) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *UpdateContactRequest) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *UpdateContactRequest) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *UpdateContactRequest) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetGiven

`func (o *UpdateContactRequest) GetGiven() string`

GetGiven returns the Given field if non-nil, zero value otherwise.

### GetGivenOk

`func (o *UpdateContactRequest) GetGivenOk() (*string, bool)`

GetGivenOk returns a tuple with the Given field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiven

`func (o *UpdateContactRequest) SetGiven(v string)`

SetGiven sets Given field to given value.

### HasGiven

`func (o *UpdateContactRequest) HasGiven() bool`

HasGiven returns a boolean if a field has been set.

### GetPicture

`func (o *UpdateContactRequest) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *UpdateContactRequest) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *UpdateContactRequest) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *UpdateContactRequest) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetEmails

`func (o *UpdateContactRequest) GetEmails() []EmailRequest`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UpdateContactRequest) GetEmailsOk() (*[]EmailRequest, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UpdateContactRequest) SetEmails(v []EmailRequest)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UpdateContactRequest) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetNumbers

`func (o *UpdateContactRequest) GetNumbers() []NumberRequest`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *UpdateContactRequest) GetNumbersOk() (*[]NumberRequest, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *UpdateContactRequest) SetNumbers(v []NumberRequest)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *UpdateContactRequest) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetAddresses

`func (o *UpdateContactRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *UpdateContactRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *UpdateContactRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *UpdateContactRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetOrganization

`func (o *UpdateContactRequest) GetOrganization() [][]string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *UpdateContactRequest) GetOrganizationOk() (*[][]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *UpdateContactRequest) SetOrganization(v [][]string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *UpdateContactRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetScope

`func (o *UpdateContactRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateContactRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateContactRequest) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


