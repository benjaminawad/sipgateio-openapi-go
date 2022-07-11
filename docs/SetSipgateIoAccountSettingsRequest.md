# SetSipgateIoAccountSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingUrl** | **string** |  | 
**OutgoingUrl** | **string** |  | 
**Log** | Pointer to **bool** |  | [optional] [default to false]
**Whitelist** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSetSipgateIoAccountSettingsRequest

`func NewSetSipgateIoAccountSettingsRequest(incomingUrl string, outgoingUrl string, ) *SetSipgateIoAccountSettingsRequest`

NewSetSipgateIoAccountSettingsRequest instantiates a new SetSipgateIoAccountSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetSipgateIoAccountSettingsRequestWithDefaults

`func NewSetSipgateIoAccountSettingsRequestWithDefaults() *SetSipgateIoAccountSettingsRequest`

NewSetSipgateIoAccountSettingsRequestWithDefaults instantiates a new SetSipgateIoAccountSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingUrl

`func (o *SetSipgateIoAccountSettingsRequest) GetIncomingUrl() string`

GetIncomingUrl returns the IncomingUrl field if non-nil, zero value otherwise.

### GetIncomingUrlOk

`func (o *SetSipgateIoAccountSettingsRequest) GetIncomingUrlOk() (*string, bool)`

GetIncomingUrlOk returns a tuple with the IncomingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingUrl

`func (o *SetSipgateIoAccountSettingsRequest) SetIncomingUrl(v string)`

SetIncomingUrl sets IncomingUrl field to given value.


### GetOutgoingUrl

`func (o *SetSipgateIoAccountSettingsRequest) GetOutgoingUrl() string`

GetOutgoingUrl returns the OutgoingUrl field if non-nil, zero value otherwise.

### GetOutgoingUrlOk

`func (o *SetSipgateIoAccountSettingsRequest) GetOutgoingUrlOk() (*string, bool)`

GetOutgoingUrlOk returns a tuple with the OutgoingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingUrl

`func (o *SetSipgateIoAccountSettingsRequest) SetOutgoingUrl(v string)`

SetOutgoingUrl sets OutgoingUrl field to given value.


### GetLog

`func (o *SetSipgateIoAccountSettingsRequest) GetLog() bool`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *SetSipgateIoAccountSettingsRequest) GetLogOk() (*bool, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *SetSipgateIoAccountSettingsRequest) SetLog(v bool)`

SetLog sets Log field to given value.

### HasLog

`func (o *SetSipgateIoAccountSettingsRequest) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetWhitelist

`func (o *SetSipgateIoAccountSettingsRequest) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *SetSipgateIoAccountSettingsRequest) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *SetSipgateIoAccountSettingsRequest) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *SetSipgateIoAccountSettingsRequest) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


