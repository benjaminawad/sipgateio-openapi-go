# HistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]HistoryEntryResponse**](HistoryEntryResponse.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewHistoryResponse

`func NewHistoryResponse() *HistoryResponse`

NewHistoryResponse instantiates a new HistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryResponseWithDefaults

`func NewHistoryResponseWithDefaults() *HistoryResponse`

NewHistoryResponseWithDefaults instantiates a new HistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *HistoryResponse) GetItems() []HistoryEntryResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HistoryResponse) GetItemsOk() (*[]HistoryEntryResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HistoryResponse) SetItems(v []HistoryEntryResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *HistoryResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCount

`func (o *HistoryResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *HistoryResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *HistoryResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *HistoryResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


