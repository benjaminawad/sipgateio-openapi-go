# ContactsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ContactResponse**](ContactResponse.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewContactsResponse

`func NewContactsResponse() *ContactsResponse`

NewContactsResponse instantiates a new ContactsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactsResponseWithDefaults

`func NewContactsResponseWithDefaults() *ContactsResponse`

NewContactsResponseWithDefaults instantiates a new ContactsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ContactsResponse) GetItems() []ContactResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContactsResponse) GetItemsOk() (*[]ContactResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContactsResponse) SetItems(v []ContactResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *ContactsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCount

`func (o *ContactsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ContactsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ContactsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ContactsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


