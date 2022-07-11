# NotificationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fax** | Pointer to [**[]FaxNotificationResponse**](FaxNotificationResponse.md) |  | [optional] 
**Voicemail** | Pointer to [**[]VoicemailNotificationResponse**](VoicemailNotificationResponse.md) |  | [optional] 
**Call** | Pointer to [**[]CallNotificationResponse**](CallNotificationResponse.md) |  | [optional] 
**Sms** | Pointer to [**[]SmsNotificationResponse**](SmsNotificationResponse.md) |  | [optional] 

## Methods

### NewNotificationsResponse

`func NewNotificationsResponse() *NotificationsResponse`

NewNotificationsResponse instantiates a new NotificationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsResponseWithDefaults

`func NewNotificationsResponseWithDefaults() *NotificationsResponse`

NewNotificationsResponseWithDefaults instantiates a new NotificationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFax

`func (o *NotificationsResponse) GetFax() []FaxNotificationResponse`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *NotificationsResponse) GetFaxOk() (*[]FaxNotificationResponse, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *NotificationsResponse) SetFax(v []FaxNotificationResponse)`

SetFax sets Fax field to given value.

### HasFax

`func (o *NotificationsResponse) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetVoicemail

`func (o *NotificationsResponse) GetVoicemail() []VoicemailNotificationResponse`

GetVoicemail returns the Voicemail field if non-nil, zero value otherwise.

### GetVoicemailOk

`func (o *NotificationsResponse) GetVoicemailOk() (*[]VoicemailNotificationResponse, bool)`

GetVoicemailOk returns a tuple with the Voicemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemail

`func (o *NotificationsResponse) SetVoicemail(v []VoicemailNotificationResponse)`

SetVoicemail sets Voicemail field to given value.

### HasVoicemail

`func (o *NotificationsResponse) HasVoicemail() bool`

HasVoicemail returns a boolean if a field has been set.

### GetCall

`func (o *NotificationsResponse) GetCall() []CallNotificationResponse`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *NotificationsResponse) GetCallOk() (*[]CallNotificationResponse, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *NotificationsResponse) SetCall(v []CallNotificationResponse)`

SetCall sets Call field to given value.

### HasCall

`func (o *NotificationsResponse) HasCall() bool`

HasCall returns a boolean if a field has been set.

### GetSms

`func (o *NotificationsResponse) GetSms() []SmsNotificationResponse`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *NotificationsResponse) GetSmsOk() (*[]SmsNotificationResponse, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *NotificationsResponse) SetSms(v []SmsNotificationResponse)`

SetSms sets Sms field to given value.

### HasSms

`func (o *NotificationsResponse) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


