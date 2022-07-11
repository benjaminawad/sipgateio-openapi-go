# SipgateIoUrlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingUrl** | Pointer to **string** |  | [optional] 
**OutgoingUrl** | Pointer to **string** |  | [optional] 
**Log** | Pointer to **bool** |  | [optional] [default to false]
**Whitelist** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSipgateIoUrlResponse

`func NewSipgateIoUrlResponse() *SipgateIoUrlResponse`

NewSipgateIoUrlResponse instantiates a new SipgateIoUrlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSipgateIoUrlResponseWithDefaults

`func NewSipgateIoUrlResponseWithDefaults() *SipgateIoUrlResponse`

NewSipgateIoUrlResponseWithDefaults instantiates a new SipgateIoUrlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingUrl

`func (o *SipgateIoUrlResponse) GetIncomingUrl() string`

GetIncomingUrl returns the IncomingUrl field if non-nil, zero value otherwise.

### GetIncomingUrlOk

`func (o *SipgateIoUrlResponse) GetIncomingUrlOk() (*string, bool)`

GetIncomingUrlOk returns a tuple with the IncomingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingUrl

`func (o *SipgateIoUrlResponse) SetIncomingUrl(v string)`

SetIncomingUrl sets IncomingUrl field to given value.

### HasIncomingUrl

`func (o *SipgateIoUrlResponse) HasIncomingUrl() bool`

HasIncomingUrl returns a boolean if a field has been set.

### GetOutgoingUrl

`func (o *SipgateIoUrlResponse) GetOutgoingUrl() string`

GetOutgoingUrl returns the OutgoingUrl field if non-nil, zero value otherwise.

### GetOutgoingUrlOk

`func (o *SipgateIoUrlResponse) GetOutgoingUrlOk() (*string, bool)`

GetOutgoingUrlOk returns a tuple with the OutgoingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingUrl

`func (o *SipgateIoUrlResponse) SetOutgoingUrl(v string)`

SetOutgoingUrl sets OutgoingUrl field to given value.

### HasOutgoingUrl

`func (o *SipgateIoUrlResponse) HasOutgoingUrl() bool`

HasOutgoingUrl returns a boolean if a field has been set.

### GetLog

`func (o *SipgateIoUrlResponse) GetLog() bool`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *SipgateIoUrlResponse) GetLogOk() (*bool, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *SipgateIoUrlResponse) SetLog(v bool)`

SetLog sets Log field to given value.

### HasLog

`func (o *SipgateIoUrlResponse) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetWhitelist

`func (o *SipgateIoUrlResponse) GetWhitelist() []string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *SipgateIoUrlResponse) GetWhitelistOk() (*[]string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *SipgateIoUrlResponse) SetWhitelist(v []string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *SipgateIoUrlResponse) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


