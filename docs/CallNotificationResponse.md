# CallNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** |  | [optional] 
**EndpointAlias** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]CallEmailTarget**](CallEmailTarget.md) |  | [optional] 
**Sms** | Pointer to [**[]CallSmsTarget**](CallSmsTarget.md) |  | [optional] 

## Methods

### NewCallNotificationResponse

`func NewCallNotificationResponse() *CallNotificationResponse`

NewCallNotificationResponse instantiates a new CallNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallNotificationResponseWithDefaults

`func NewCallNotificationResponseWithDefaults() *CallNotificationResponse`

NewCallNotificationResponseWithDefaults instantiates a new CallNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *CallNotificationResponse) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *CallNotificationResponse) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *CallNotificationResponse) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *CallNotificationResponse) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEndpointAlias

`func (o *CallNotificationResponse) GetEndpointAlias() string`

GetEndpointAlias returns the EndpointAlias field if non-nil, zero value otherwise.

### GetEndpointAliasOk

`func (o *CallNotificationResponse) GetEndpointAliasOk() (*string, bool)`

GetEndpointAliasOk returns a tuple with the EndpointAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointAlias

`func (o *CallNotificationResponse) SetEndpointAlias(v string)`

SetEndpointAlias sets EndpointAlias field to given value.

### HasEndpointAlias

`func (o *CallNotificationResponse) HasEndpointAlias() bool`

HasEndpointAlias returns a boolean if a field has been set.

### GetEmails

`func (o *CallNotificationResponse) GetEmails() []CallEmailTarget`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CallNotificationResponse) GetEmailsOk() (*[]CallEmailTarget, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CallNotificationResponse) SetEmails(v []CallEmailTarget)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *CallNotificationResponse) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetSms

`func (o *CallNotificationResponse) GetSms() []CallSmsTarget`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *CallNotificationResponse) GetSmsOk() (*[]CallSmsTarget, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *CallNotificationResponse) SetSms(v []CallSmsTarget)`

SetSms sets Sms field to given value.

### HasSms

`func (o *CallNotificationResponse) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


