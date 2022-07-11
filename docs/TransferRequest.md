# TransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attended** | **bool** | The transfer is attended or blind | 
**PhoneNumber** | **string** | Transfer target phone number | 
**CallerId** | Pointer to **string** | Caller phone nuber | [optional] 

## Methods

### NewTransferRequest

`func NewTransferRequest(attended bool, phoneNumber string, ) *TransferRequest`

NewTransferRequest instantiates a new TransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestWithDefaults

`func NewTransferRequestWithDefaults() *TransferRequest`

NewTransferRequestWithDefaults instantiates a new TransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttended

`func (o *TransferRequest) GetAttended() bool`

GetAttended returns the Attended field if non-nil, zero value otherwise.

### GetAttendedOk

`func (o *TransferRequest) GetAttendedOk() (*bool, bool)`

GetAttendedOk returns a tuple with the Attended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttended

`func (o *TransferRequest) SetAttended(v bool)`

SetAttended sets Attended field to given value.


### GetPhoneNumber

`func (o *TransferRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TransferRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TransferRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetCallerId

`func (o *TransferRequest) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *TransferRequest) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *TransferRequest) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *TransferRequest) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


