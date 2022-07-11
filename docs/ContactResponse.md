# ContactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]EMailResponse**](EMailResponse.md) |  | [optional] 
**Numbers** | Pointer to [**[]NumberResponse**](NumberResponse.md) |  | [optional] 
**Addresses** | Pointer to [**[]AddressResponse**](AddressResponse.md) |  | [optional] 
**Organization** | Pointer to **[][]string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 

## Methods

### NewContactResponse

`func NewContactResponse() *ContactResponse`

NewContactResponse instantiates a new ContactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactResponseWithDefaults

`func NewContactResponseWithDefaults() *ContactResponse`

NewContactResponseWithDefaults instantiates a new ContactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ContactResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPicture

`func (o *ContactResponse) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *ContactResponse) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *ContactResponse) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *ContactResponse) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetEmails

`func (o *ContactResponse) GetEmails() []EMailResponse`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ContactResponse) GetEmailsOk() (*[]EMailResponse, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ContactResponse) SetEmails(v []EMailResponse)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *ContactResponse) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetNumbers

`func (o *ContactResponse) GetNumbers() []NumberResponse`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *ContactResponse) GetNumbersOk() (*[]NumberResponse, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *ContactResponse) SetNumbers(v []NumberResponse)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *ContactResponse) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetAddresses

`func (o *ContactResponse) GetAddresses() []AddressResponse`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ContactResponse) GetAddressesOk() (*[]AddressResponse, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ContactResponse) SetAddresses(v []AddressResponse)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ContactResponse) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetOrganization

`func (o *ContactResponse) GetOrganization() [][]string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ContactResponse) GetOrganizationOk() (*[][]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ContactResponse) SetOrganization(v [][]string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ContactResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetScope

`func (o *ContactResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ContactResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ContactResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ContactResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


