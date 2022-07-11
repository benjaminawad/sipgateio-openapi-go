# GroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**DepartmentId** | Pointer to **int64** |  | [optional] 
**Voicemails** | Pointer to [**[]VoicemailSettingsResponse**](VoicemailSettingsResponse.md) |  | [optional] 

## Methods

### NewGroupResponse

`func NewGroupResponse() *GroupResponse`

NewGroupResponse instantiates a new GroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupResponseWithDefaults

`func NewGroupResponseWithDefaults() *GroupResponse`

NewGroupResponseWithDefaults instantiates a new GroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *GroupResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GroupResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GroupResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *GroupResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDepartmentId

`func (o *GroupResponse) GetDepartmentId() int64`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *GroupResponse) GetDepartmentIdOk() (*int64, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *GroupResponse) SetDepartmentId(v int64)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *GroupResponse) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### GetVoicemails

`func (o *GroupResponse) GetVoicemails() []VoicemailSettingsResponse`

GetVoicemails returns the Voicemails field if non-nil, zero value otherwise.

### GetVoicemailsOk

`func (o *GroupResponse) GetVoicemailsOk() (*[]VoicemailSettingsResponse, bool)`

GetVoicemailsOk returns a tuple with the Voicemails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemails

`func (o *GroupResponse) SetVoicemails(v []VoicemailSettingsResponse)`

SetVoicemails sets Voicemails field to given value.

### HasVoicemails

`func (o *GroupResponse) HasVoicemails() bool`

HasVoicemails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


