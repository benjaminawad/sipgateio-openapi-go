# HistoryEntryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**SourceAlias** | Pointer to **string** |  | [optional] 
**TargetAlias** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Incoming** | Pointer to **bool** |  | [optional] [default to false]
**Status** | Pointer to **string** |  | [optional] 
**ConnectionIds** | Pointer to **[]string** |  | [optional] 
**Read** | Pointer to **bool** |  | [optional] [default to false]
**Archived** | Pointer to **bool** |  | [optional] [default to false]
**Note** | Pointer to **string** |  | [optional] 
**Endpoints** | Pointer to [**[]RoutedEndpointResponse**](RoutedEndpointResponse.md) |  | [optional] 
**Starred** | Pointer to **bool** |  | [optional] [default to false]
**Labels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHistoryEntryResponse

`func NewHistoryEntryResponse() *HistoryEntryResponse`

NewHistoryEntryResponse instantiates a new HistoryEntryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryEntryResponseWithDefaults

`func NewHistoryEntryResponseWithDefaults() *HistoryEntryResponse`

NewHistoryEntryResponseWithDefaults instantiates a new HistoryEntryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryEntryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryEntryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryEntryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HistoryEntryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *HistoryEntryResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HistoryEntryResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HistoryEntryResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *HistoryEntryResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *HistoryEntryResponse) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *HistoryEntryResponse) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *HistoryEntryResponse) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *HistoryEntryResponse) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetSourceAlias

`func (o *HistoryEntryResponse) GetSourceAlias() string`

GetSourceAlias returns the SourceAlias field if non-nil, zero value otherwise.

### GetSourceAliasOk

`func (o *HistoryEntryResponse) GetSourceAliasOk() (*string, bool)`

GetSourceAliasOk returns a tuple with the SourceAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAlias

`func (o *HistoryEntryResponse) SetSourceAlias(v string)`

SetSourceAlias sets SourceAlias field to given value.

### HasSourceAlias

`func (o *HistoryEntryResponse) HasSourceAlias() bool`

HasSourceAlias returns a boolean if a field has been set.

### GetTargetAlias

`func (o *HistoryEntryResponse) GetTargetAlias() string`

GetTargetAlias returns the TargetAlias field if non-nil, zero value otherwise.

### GetTargetAliasOk

`func (o *HistoryEntryResponse) GetTargetAliasOk() (*string, bool)`

GetTargetAliasOk returns a tuple with the TargetAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAlias

`func (o *HistoryEntryResponse) SetTargetAlias(v string)`

SetTargetAlias sets TargetAlias field to given value.

### HasTargetAlias

`func (o *HistoryEntryResponse) HasTargetAlias() bool`

HasTargetAlias returns a boolean if a field has been set.

### GetType

`func (o *HistoryEntryResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoryEntryResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoryEntryResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoryEntryResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreated

`func (o *HistoryEntryResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *HistoryEntryResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *HistoryEntryResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *HistoryEntryResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *HistoryEntryResponse) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *HistoryEntryResponse) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *HistoryEntryResponse) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *HistoryEntryResponse) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetDirection

`func (o *HistoryEntryResponse) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *HistoryEntryResponse) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *HistoryEntryResponse) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *HistoryEntryResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetIncoming

`func (o *HistoryEntryResponse) GetIncoming() bool`

GetIncoming returns the Incoming field if non-nil, zero value otherwise.

### GetIncomingOk

`func (o *HistoryEntryResponse) GetIncomingOk() (*bool, bool)`

GetIncomingOk returns a tuple with the Incoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoming

`func (o *HistoryEntryResponse) SetIncoming(v bool)`

SetIncoming sets Incoming field to given value.

### HasIncoming

`func (o *HistoryEntryResponse) HasIncoming() bool`

HasIncoming returns a boolean if a field has been set.

### GetStatus

`func (o *HistoryEntryResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HistoryEntryResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HistoryEntryResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HistoryEntryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnectionIds

`func (o *HistoryEntryResponse) GetConnectionIds() []string`

GetConnectionIds returns the ConnectionIds field if non-nil, zero value otherwise.

### GetConnectionIdsOk

`func (o *HistoryEntryResponse) GetConnectionIdsOk() (*[]string, bool)`

GetConnectionIdsOk returns a tuple with the ConnectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionIds

`func (o *HistoryEntryResponse) SetConnectionIds(v []string)`

SetConnectionIds sets ConnectionIds field to given value.

### HasConnectionIds

`func (o *HistoryEntryResponse) HasConnectionIds() bool`

HasConnectionIds returns a boolean if a field has been set.

### GetRead

`func (o *HistoryEntryResponse) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *HistoryEntryResponse) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *HistoryEntryResponse) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *HistoryEntryResponse) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetArchived

`func (o *HistoryEntryResponse) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *HistoryEntryResponse) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *HistoryEntryResponse) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *HistoryEntryResponse) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetNote

`func (o *HistoryEntryResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *HistoryEntryResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *HistoryEntryResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *HistoryEntryResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetEndpoints

`func (o *HistoryEntryResponse) GetEndpoints() []RoutedEndpointResponse`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *HistoryEntryResponse) GetEndpointsOk() (*[]RoutedEndpointResponse, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *HistoryEntryResponse) SetEndpoints(v []RoutedEndpointResponse)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *HistoryEntryResponse) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetStarred

`func (o *HistoryEntryResponse) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *HistoryEntryResponse) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *HistoryEntryResponse) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *HistoryEntryResponse) HasStarred() bool`

HasStarred returns a boolean if a field has been set.

### GetLabels

`func (o *HistoryEntryResponse) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HistoryEntryResponse) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HistoryEntryResponse) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HistoryEntryResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


