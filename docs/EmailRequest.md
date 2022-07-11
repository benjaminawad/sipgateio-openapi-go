# EmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEmailRequest

`func NewEmailRequest() *EmailRequest`

NewEmailRequest instantiates a new EmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRequestWithDefaults

`func NewEmailRequestWithDefaults() *EmailRequest`

NewEmailRequestWithDefaults instantiates a new EmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetType

`func (o *EmailRequest) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailRequest) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailRequest) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *EmailRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


