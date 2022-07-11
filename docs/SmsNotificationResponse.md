# SmsNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** |  | [optional] 
**EndpointAlias** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]SmsEmailTarget**](SmsEmailTarget.md) |  | [optional] 

## Methods

### NewSmsNotificationResponse

`func NewSmsNotificationResponse() *SmsNotificationResponse`

NewSmsNotificationResponse instantiates a new SmsNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsNotificationResponseWithDefaults

`func NewSmsNotificationResponseWithDefaults() *SmsNotificationResponse`

NewSmsNotificationResponseWithDefaults instantiates a new SmsNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *SmsNotificationResponse) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *SmsNotificationResponse) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *SmsNotificationResponse) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *SmsNotificationResponse) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEndpointAlias

`func (o *SmsNotificationResponse) GetEndpointAlias() string`

GetEndpointAlias returns the EndpointAlias field if non-nil, zero value otherwise.

### GetEndpointAliasOk

`func (o *SmsNotificationResponse) GetEndpointAliasOk() (*string, bool)`

GetEndpointAliasOk returns a tuple with the EndpointAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointAlias

`func (o *SmsNotificationResponse) SetEndpointAlias(v string)`

SetEndpointAlias sets EndpointAlias field to given value.

### HasEndpointAlias

`func (o *SmsNotificationResponse) HasEndpointAlias() bool`

HasEndpointAlias returns a boolean if a field has been set.

### GetEmails

`func (o *SmsNotificationResponse) GetEmails() []SmsEmailTarget`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SmsNotificationResponse) GetEmailsOk() (*[]SmsEmailTarget, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SmsNotificationResponse) SetEmails(v []SmsEmailTarget)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *SmsNotificationResponse) HasEmails() bool`

HasEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


