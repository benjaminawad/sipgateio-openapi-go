# FaxNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** |  | [optional] 
**EndpointAlias** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]FaxEmailTarget**](FaxEmailTarget.md) |  | [optional] 
**Sms** | Pointer to [**[]FaxSmsTarget**](FaxSmsTarget.md) |  | [optional] 
**Reports** | Pointer to [**[]ReportTarget**](ReportTarget.md) |  | [optional] 

## Methods

### NewFaxNotificationResponse

`func NewFaxNotificationResponse() *FaxNotificationResponse`

NewFaxNotificationResponse instantiates a new FaxNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxNotificationResponseWithDefaults

`func NewFaxNotificationResponseWithDefaults() *FaxNotificationResponse`

NewFaxNotificationResponseWithDefaults instantiates a new FaxNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *FaxNotificationResponse) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *FaxNotificationResponse) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *FaxNotificationResponse) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *FaxNotificationResponse) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEndpointAlias

`func (o *FaxNotificationResponse) GetEndpointAlias() string`

GetEndpointAlias returns the EndpointAlias field if non-nil, zero value otherwise.

### GetEndpointAliasOk

`func (o *FaxNotificationResponse) GetEndpointAliasOk() (*string, bool)`

GetEndpointAliasOk returns a tuple with the EndpointAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointAlias

`func (o *FaxNotificationResponse) SetEndpointAlias(v string)`

SetEndpointAlias sets EndpointAlias field to given value.

### HasEndpointAlias

`func (o *FaxNotificationResponse) HasEndpointAlias() bool`

HasEndpointAlias returns a boolean if a field has been set.

### GetEmails

`func (o *FaxNotificationResponse) GetEmails() []FaxEmailTarget`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *FaxNotificationResponse) GetEmailsOk() (*[]FaxEmailTarget, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *FaxNotificationResponse) SetEmails(v []FaxEmailTarget)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *FaxNotificationResponse) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetSms

`func (o *FaxNotificationResponse) GetSms() []FaxSmsTarget`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *FaxNotificationResponse) GetSmsOk() (*[]FaxSmsTarget, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *FaxNotificationResponse) SetSms(v []FaxSmsTarget)`

SetSms sets Sms field to given value.

### HasSms

`func (o *FaxNotificationResponse) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetReports

`func (o *FaxNotificationResponse) GetReports() []ReportTarget`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *FaxNotificationResponse) GetReportsOk() (*[]ReportTarget, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *FaxNotificationResponse) SetReports(v []ReportTarget)`

SetReports sets Reports field to given value.

### HasReports

`func (o *FaxNotificationResponse) HasReports() bool`

HasReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


