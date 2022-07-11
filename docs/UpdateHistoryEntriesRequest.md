# UpdateHistoryEntriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Archived** | Pointer to **bool** |  | [optional] [default to false]
**Read** | Pointer to **bool** |  | [optional] [default to false]
**Starred** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewUpdateHistoryEntriesRequest

`func NewUpdateHistoryEntriesRequest(id string, ) *UpdateHistoryEntriesRequest`

NewUpdateHistoryEntriesRequest instantiates a new UpdateHistoryEntriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHistoryEntriesRequestWithDefaults

`func NewUpdateHistoryEntriesRequestWithDefaults() *UpdateHistoryEntriesRequest`

NewUpdateHistoryEntriesRequestWithDefaults instantiates a new UpdateHistoryEntriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateHistoryEntriesRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateHistoryEntriesRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateHistoryEntriesRequest) SetId(v string)`

SetId sets Id field to given value.


### GetArchived

`func (o *UpdateHistoryEntriesRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateHistoryEntriesRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateHistoryEntriesRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateHistoryEntriesRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetRead

`func (o *UpdateHistoryEntriesRequest) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *UpdateHistoryEntriesRequest) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *UpdateHistoryEntriesRequest) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *UpdateHistoryEntriesRequest) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetStarred

`func (o *UpdateHistoryEntriesRequest) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *UpdateHistoryEntriesRequest) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *UpdateHistoryEntriesRequest) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *UpdateHistoryEntriesRequest) HasStarred() bool`

HasStarred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


