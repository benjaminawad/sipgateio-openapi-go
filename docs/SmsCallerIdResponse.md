# SmsCallerIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Phonenumber** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] [default to false]
**DefaultNumber** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSmsCallerIdResponse

`func NewSmsCallerIdResponse() *SmsCallerIdResponse`

NewSmsCallerIdResponse instantiates a new SmsCallerIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsCallerIdResponseWithDefaults

`func NewSmsCallerIdResponseWithDefaults() *SmsCallerIdResponse`

NewSmsCallerIdResponseWithDefaults instantiates a new SmsCallerIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsCallerIdResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsCallerIdResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsCallerIdResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SmsCallerIdResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhonenumber

`func (o *SmsCallerIdResponse) GetPhonenumber() string`

GetPhonenumber returns the Phonenumber field if non-nil, zero value otherwise.

### GetPhonenumberOk

`func (o *SmsCallerIdResponse) GetPhonenumberOk() (*string, bool)`

GetPhonenumberOk returns a tuple with the Phonenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonenumber

`func (o *SmsCallerIdResponse) SetPhonenumber(v string)`

SetPhonenumber sets Phonenumber field to given value.

### HasPhonenumber

`func (o *SmsCallerIdResponse) HasPhonenumber() bool`

HasPhonenumber returns a boolean if a field has been set.

### GetVerified

`func (o *SmsCallerIdResponse) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *SmsCallerIdResponse) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *SmsCallerIdResponse) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *SmsCallerIdResponse) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetDefaultNumber

`func (o *SmsCallerIdResponse) GetDefaultNumber() bool`

GetDefaultNumber returns the DefaultNumber field if non-nil, zero value otherwise.

### GetDefaultNumberOk

`func (o *SmsCallerIdResponse) GetDefaultNumberOk() (*bool, bool)`

GetDefaultNumberOk returns a tuple with the DefaultNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNumber

`func (o *SmsCallerIdResponse) SetDefaultNumber(v bool)`

SetDefaultNumber sets DefaultNumber field to given value.

### HasDefaultNumber

`func (o *SmsCallerIdResponse) HasDefaultNumber() bool`

HasDefaultNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


