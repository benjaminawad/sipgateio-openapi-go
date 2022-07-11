# EMailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEMailResponse

`func NewEMailResponse() *EMailResponse`

NewEMailResponse instantiates a new EMailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEMailResponseWithDefaults

`func NewEMailResponseWithDefaults() *EMailResponse`

NewEMailResponseWithDefaults instantiates a new EMailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EMailResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EMailResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EMailResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EMailResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetType

`func (o *EMailResponse) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EMailResponse) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EMailResponse) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *EMailResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


