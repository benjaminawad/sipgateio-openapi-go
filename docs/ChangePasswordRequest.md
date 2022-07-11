# ChangePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**PasswordRepeat** | Pointer to **string** |  | [optional] 

## Methods

### NewChangePasswordRequest

`func NewChangePasswordRequest() *ChangePasswordRequest`

NewChangePasswordRequest instantiates a new ChangePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePasswordRequestWithDefaults

`func NewChangePasswordRequestWithDefaults() *ChangePasswordRequest`

NewChangePasswordRequestWithDefaults instantiates a new ChangePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ChangePasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ChangePasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ChangePasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ChangePasswordRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordRepeat

`func (o *ChangePasswordRequest) GetPasswordRepeat() string`

GetPasswordRepeat returns the PasswordRepeat field if non-nil, zero value otherwise.

### GetPasswordRepeatOk

`func (o *ChangePasswordRequest) GetPasswordRepeatOk() (*string, bool)`

GetPasswordRepeatOk returns a tuple with the PasswordRepeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRepeat

`func (o *ChangePasswordRequest) SetPasswordRepeat(v string)`

SetPasswordRepeat sets PasswordRepeat field to given value.

### HasPasswordRepeat

`func (o *ChangePasswordRequest) HasPasswordRepeat() bool`

HasPasswordRepeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


