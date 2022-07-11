# SetSipgateIoPhonelineSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingActive** | Pointer to **bool** |  | [optional] [default to false]
**OutgoingActive** | Pointer to **bool** |  | [optional] [default to false]
**IncomingUrl** | **string** |  | 
**OutgoingUrl** | **string** |  | 
**Log** | Pointer to **bool** |  | [optional] [default to false]
**Certificate** | Pointer to **string** |  | [optional] 
**AllowSelfSigned** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSetSipgateIoPhonelineSettingsRequest

`func NewSetSipgateIoPhonelineSettingsRequest(incomingUrl string, outgoingUrl string, ) *SetSipgateIoPhonelineSettingsRequest`

NewSetSipgateIoPhonelineSettingsRequest instantiates a new SetSipgateIoPhonelineSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetSipgateIoPhonelineSettingsRequestWithDefaults

`func NewSetSipgateIoPhonelineSettingsRequestWithDefaults() *SetSipgateIoPhonelineSettingsRequest`

NewSetSipgateIoPhonelineSettingsRequestWithDefaults instantiates a new SetSipgateIoPhonelineSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingActive

`func (o *SetSipgateIoPhonelineSettingsRequest) GetIncomingActive() bool`

GetIncomingActive returns the IncomingActive field if non-nil, zero value otherwise.

### GetIncomingActiveOk

`func (o *SetSipgateIoPhonelineSettingsRequest) GetIncomingActiveOk() (*bool, bool)`

GetIncomingActiveOk returns a tuple with the IncomingActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingActive

`func (o *SetSipgateIoPhonelineSettingsRequest) SetIncomingActive(v bool)`

SetIncomingActive sets IncomingActive field to given value.

### HasIncomingActive

`func (o *SetSipgateIoPhonelineSettingsRequest) HasIncomingActive() bool`

HasIncomingActive returns a boolean if a field has been set.

### GetOutgoingActive

`func (o *SetSipgateIoPhonelineSettingsRequest) GetOutgoingActive() bool`

GetOutgoingActive returns the OutgoingActive field if non-nil, zero value otherwise.

### GetOutgoingActiveOk

`func (o *SetSipgateIoPhonelineSettingsRequest) GetOutgoingActiveOk() (*bool, bool)`

GetOutgoingActiveOk returns a tuple with the OutgoingActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingActive

`func (o *SetSipgateIoPhonelineSettingsRequest) SetOutgoingActive(v bool)`

SetOutgoingActive sets OutgoingActive field to given value.

### HasOutgoingActive

`func (o *SetSipgateIoPhonelineSettingsRequest) HasOutgoingActive() bool`

HasOutgoingActive returns a boolean if a field has been set.

### GetIncomingUrl

`func (o *SetSipgateIoPhonelineSettingsRequest) GetIncomingUrl() string`

GetIncomingUrl returns the IncomingUrl field if non-nil, zero value otherwise.

### GetIncomingUrlOk

`func (o *SetSipgateIoPhonelineSettingsRequest) GetIncomingUrlOk() (*string, bool)`

GetIncomingUrlOk returns a tuple with the IncomingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingUrl

`func (o *SetSipgateIoPhonelineSettingsRequest) SetIncomingUrl(v string)`

SetIncomingUrl sets IncomingUrl field to given value.


### GetOutgoingUrl

`func (o *SetSipgateIoPhonelineSettingsRequest) GetOutgoingUrl() string`

GetOutgoingUrl returns the OutgoingUrl field if non-nil, zero value otherwise.

### GetOutgoingUrlOk

`func (o *SetSipgateIoPhonelineSettingsRequest) GetOutgoingUrlOk() (*string, bool)`

GetOutgoingUrlOk returns a tuple with the OutgoingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingUrl

`func (o *SetSipgateIoPhonelineSettingsRequest) SetOutgoingUrl(v string)`

SetOutgoingUrl sets OutgoingUrl field to given value.


### GetLog

`func (o *SetSipgateIoPhonelineSettingsRequest) GetLog() bool`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *SetSipgateIoPhonelineSettingsRequest) GetLogOk() (*bool, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *SetSipgateIoPhonelineSettingsRequest) SetLog(v bool)`

SetLog sets Log field to given value.

### HasLog

`func (o *SetSipgateIoPhonelineSettingsRequest) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetCertificate

`func (o *SetSipgateIoPhonelineSettingsRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SetSipgateIoPhonelineSettingsRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SetSipgateIoPhonelineSettingsRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SetSipgateIoPhonelineSettingsRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetAllowSelfSigned

`func (o *SetSipgateIoPhonelineSettingsRequest) GetAllowSelfSigned() bool`

GetAllowSelfSigned returns the AllowSelfSigned field if non-nil, zero value otherwise.

### GetAllowSelfSignedOk

`func (o *SetSipgateIoPhonelineSettingsRequest) GetAllowSelfSignedOk() (*bool, bool)`

GetAllowSelfSignedOk returns a tuple with the AllowSelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSelfSigned

`func (o *SetSipgateIoPhonelineSettingsRequest) SetAllowSelfSigned(v bool)`

SetAllowSelfSigned sets AllowSelfSigned field to given value.

### HasAllowSelfSigned

`func (o *SetSipgateIoPhonelineSettingsRequest) HasAllowSelfSigned() bool`

HasAllowSelfSigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


