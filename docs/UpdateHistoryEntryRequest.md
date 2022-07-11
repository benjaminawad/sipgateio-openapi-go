# UpdateHistoryEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to **string** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] [default to false]
**Read** | Pointer to **bool** |  | [optional] [default to false]
**Starred** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewUpdateHistoryEntryRequest

`func NewUpdateHistoryEntryRequest() *UpdateHistoryEntryRequest`

NewUpdateHistoryEntryRequest instantiates a new UpdateHistoryEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHistoryEntryRequestWithDefaults

`func NewUpdateHistoryEntryRequestWithDefaults() *UpdateHistoryEntryRequest`

NewUpdateHistoryEntryRequestWithDefaults instantiates a new UpdateHistoryEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *UpdateHistoryEntryRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UpdateHistoryEntryRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UpdateHistoryEntryRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UpdateHistoryEntryRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetArchived

`func (o *UpdateHistoryEntryRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateHistoryEntryRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateHistoryEntryRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateHistoryEntryRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetRead

`func (o *UpdateHistoryEntryRequest) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *UpdateHistoryEntryRequest) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *UpdateHistoryEntryRequest) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *UpdateHistoryEntryRequest) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetStarred

`func (o *UpdateHistoryEntryRequest) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *UpdateHistoryEntryRequest) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *UpdateHistoryEntryRequest) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *UpdateHistoryEntryRequest) HasStarred() bool`

HasStarred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


