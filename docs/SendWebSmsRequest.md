# SendWebSmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmsId** | **string** |  | 
**Recipient** | **string** |  | 
**Message** | **string** |  | 
**SendAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewSendWebSmsRequest

`func NewSendWebSmsRequest(smsId string, recipient string, message string, ) *SendWebSmsRequest`

NewSendWebSmsRequest instantiates a new SendWebSmsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendWebSmsRequestWithDefaults

`func NewSendWebSmsRequestWithDefaults() *SendWebSmsRequest`

NewSendWebSmsRequestWithDefaults instantiates a new SendWebSmsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmsId

`func (o *SendWebSmsRequest) GetSmsId() string`

GetSmsId returns the SmsId field if non-nil, zero value otherwise.

### GetSmsIdOk

`func (o *SendWebSmsRequest) GetSmsIdOk() (*string, bool)`

GetSmsIdOk returns a tuple with the SmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsId

`func (o *SendWebSmsRequest) SetSmsId(v string)`

SetSmsId sets SmsId field to given value.


### GetRecipient

`func (o *SendWebSmsRequest) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *SendWebSmsRequest) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *SendWebSmsRequest) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetMessage

`func (o *SendWebSmsRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendWebSmsRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendWebSmsRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSendAt

`func (o *SendWebSmsRequest) GetSendAt() int64`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *SendWebSmsRequest) GetSendAtOk() (*int64, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *SendWebSmsRequest) SetSendAt(v int64)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *SendWebSmsRequest) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


