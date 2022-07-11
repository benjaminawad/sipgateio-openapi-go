# VoicemailNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** |  | [optional] 
**EndpointAlias** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]EmailTarget**](EmailTarget.md) |  | [optional] 
**Sms** | Pointer to [**[]SmsTarget**](SmsTarget.md) |  | [optional] 

## Methods

### NewVoicemailNotificationResponse

`func NewVoicemailNotificationResponse() *VoicemailNotificationResponse`

NewVoicemailNotificationResponse instantiates a new VoicemailNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailNotificationResponseWithDefaults

`func NewVoicemailNotificationResponseWithDefaults() *VoicemailNotificationResponse`

NewVoicemailNotificationResponseWithDefaults instantiates a new VoicemailNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *VoicemailNotificationResponse) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *VoicemailNotificationResponse) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *VoicemailNotificationResponse) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *VoicemailNotificationResponse) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEndpointAlias

`func (o *VoicemailNotificationResponse) GetEndpointAlias() string`

GetEndpointAlias returns the EndpointAlias field if non-nil, zero value otherwise.

### GetEndpointAliasOk

`func (o *VoicemailNotificationResponse) GetEndpointAliasOk() (*string, bool)`

GetEndpointAliasOk returns a tuple with the EndpointAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointAlias

`func (o *VoicemailNotificationResponse) SetEndpointAlias(v string)`

SetEndpointAlias sets EndpointAlias field to given value.

### HasEndpointAlias

`func (o *VoicemailNotificationResponse) HasEndpointAlias() bool`

HasEndpointAlias returns a boolean if a field has been set.

### GetEmails

`func (o *VoicemailNotificationResponse) GetEmails() []EmailTarget`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *VoicemailNotificationResponse) GetEmailsOk() (*[]EmailTarget, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *VoicemailNotificationResponse) SetEmails(v []EmailTarget)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *VoicemailNotificationResponse) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetSms

`func (o *VoicemailNotificationResponse) GetSms() []SmsTarget`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *VoicemailNotificationResponse) GetSmsOk() (*[]SmsTarget, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *VoicemailNotificationResponse) SetSms(v []SmsTarget)`

SetSms sets Sms field to given value.

### HasSms

`func (o *VoicemailNotificationResponse) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


