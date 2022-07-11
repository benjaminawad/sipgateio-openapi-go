# PortingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PortDate** | Pointer to **string** |  | [optional] 
**Numbers** | Pointer to [**[]PortingNumberResponse**](PortingNumberResponse.md) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Revocable** | Pointer to **bool** |  | [optional] [default to false]
**Pending** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPortingResponse

`func NewPortingResponse() *PortingResponse`

NewPortingResponse instantiates a new PortingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingResponseWithDefaults

`func NewPortingResponseWithDefaults() *PortingResponse`

NewPortingResponseWithDefaults instantiates a new PortingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortDate

`func (o *PortingResponse) GetPortDate() string`

GetPortDate returns the PortDate field if non-nil, zero value otherwise.

### GetPortDateOk

`func (o *PortingResponse) GetPortDateOk() (*string, bool)`

GetPortDateOk returns a tuple with the PortDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDate

`func (o *PortingResponse) SetPortDate(v string)`

SetPortDate sets PortDate field to given value.

### HasPortDate

`func (o *PortingResponse) HasPortDate() bool`

HasPortDate returns a boolean if a field has been set.

### GetNumbers

`func (o *PortingResponse) GetNumbers() []PortingNumberResponse`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *PortingResponse) GetNumbersOk() (*[]PortingNumberResponse, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *PortingResponse) SetNumbers(v []PortingNumberResponse)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *PortingResponse) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetStatus

`func (o *PortingResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortingResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortingResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortingResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRevocable

`func (o *PortingResponse) GetRevocable() bool`

GetRevocable returns the Revocable field if non-nil, zero value otherwise.

### GetRevocableOk

`func (o *PortingResponse) GetRevocableOk() (*bool, bool)`

GetRevocableOk returns a tuple with the Revocable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocable

`func (o *PortingResponse) SetRevocable(v bool)`

SetRevocable sets Revocable field to given value.

### HasRevocable

`func (o *PortingResponse) HasRevocable() bool`

HasRevocable returns a boolean if a field has been set.

### GetPending

`func (o *PortingResponse) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *PortingResponse) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *PortingResponse) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *PortingResponse) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


