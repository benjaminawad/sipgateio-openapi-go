# SendFaxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FaxlineId** | **string** |  | 
**Recipient** | **string** |  | 
**Filename** | **string** |  | 
**Base64Content** | **string** |  | 

## Methods

### NewSendFaxRequest

`func NewSendFaxRequest(faxlineId string, recipient string, filename string, base64Content string, ) *SendFaxRequest`

NewSendFaxRequest instantiates a new SendFaxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendFaxRequestWithDefaults

`func NewSendFaxRequestWithDefaults() *SendFaxRequest`

NewSendFaxRequestWithDefaults instantiates a new SendFaxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaxlineId

`func (o *SendFaxRequest) GetFaxlineId() string`

GetFaxlineId returns the FaxlineId field if non-nil, zero value otherwise.

### GetFaxlineIdOk

`func (o *SendFaxRequest) GetFaxlineIdOk() (*string, bool)`

GetFaxlineIdOk returns a tuple with the FaxlineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxlineId

`func (o *SendFaxRequest) SetFaxlineId(v string)`

SetFaxlineId sets FaxlineId field to given value.


### GetRecipient

`func (o *SendFaxRequest) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *SendFaxRequest) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *SendFaxRequest) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetFilename

`func (o *SendFaxRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *SendFaxRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *SendFaxRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetBase64Content

`func (o *SendFaxRequest) GetBase64Content() string`

GetBase64Content returns the Base64Content field if non-nil, zero value otherwise.

### GetBase64ContentOk

`func (o *SendFaxRequest) GetBase64ContentOk() (*string, bool)`

GetBase64ContentOk returns a tuple with the Base64Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Content

`func (o *SendFaxRequest) SetBase64Content(v string)`

SetBase64Content sets Base64Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


