# IncomingEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number | 
**IsBlock** | Pointer to **bool** | Should phone numbers be blocked beginning with given number | [optional] [default to false]

## Methods

### NewIncomingEntry

`func NewIncomingEntry(phoneNumber string, ) *IncomingEntry`

NewIncomingEntry instantiates a new IncomingEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingEntryWithDefaults

`func NewIncomingEntryWithDefaults() *IncomingEntry`

NewIncomingEntryWithDefaults instantiates a new IncomingEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *IncomingEntry) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IncomingEntry) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IncomingEntry) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetIsBlock

`func (o *IncomingEntry) GetIsBlock() bool`

GetIsBlock returns the IsBlock field if non-nil, zero value otherwise.

### GetIsBlockOk

`func (o *IncomingEntry) GetIsBlockOk() (*bool, bool)`

GetIsBlockOk returns a tuple with the IsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlock

`func (o *IncomingEntry) SetIsBlock(v bool)`

SetIsBlock sets IsBlock field to given value.

### HasIsBlock

`func (o *IncomingEntry) HasIsBlock() bool`

HasIsBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


