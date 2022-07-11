# ExternalDeviceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] [default to false]
**Dnd** | Pointer to **bool** |  | [optional] [default to false]
**ActivePhonelines** | Pointer to [**[]ActiveRouting**](ActiveRouting.md) |  | [optional] 
**ActiveGroups** | Pointer to [**[]ActiveRouting**](ActiveRouting.md) |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**IncomingCallDisplayState** | Pointer to **string** |  | [optional] 

## Methods

### NewExternalDeviceResponse

`func NewExternalDeviceResponse() *ExternalDeviceResponse`

NewExternalDeviceResponse instantiates a new ExternalDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalDeviceResponseWithDefaults

`func NewExternalDeviceResponseWithDefaults() *ExternalDeviceResponse`

NewExternalDeviceResponseWithDefaults instantiates a new ExternalDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalDeviceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalDeviceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalDeviceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalDeviceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *ExternalDeviceResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ExternalDeviceResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ExternalDeviceResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ExternalDeviceResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetType

`func (o *ExternalDeviceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalDeviceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalDeviceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalDeviceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOnline

`func (o *ExternalDeviceResponse) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ExternalDeviceResponse) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ExternalDeviceResponse) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ExternalDeviceResponse) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetDnd

`func (o *ExternalDeviceResponse) GetDnd() bool`

GetDnd returns the Dnd field if non-nil, zero value otherwise.

### GetDndOk

`func (o *ExternalDeviceResponse) GetDndOk() (*bool, bool)`

GetDndOk returns a tuple with the Dnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnd

`func (o *ExternalDeviceResponse) SetDnd(v bool)`

SetDnd sets Dnd field to given value.

### HasDnd

`func (o *ExternalDeviceResponse) HasDnd() bool`

HasDnd returns a boolean if a field has been set.

### GetActivePhonelines

`func (o *ExternalDeviceResponse) GetActivePhonelines() []ActiveRouting`

GetActivePhonelines returns the ActivePhonelines field if non-nil, zero value otherwise.

### GetActivePhonelinesOk

`func (o *ExternalDeviceResponse) GetActivePhonelinesOk() (*[]ActiveRouting, bool)`

GetActivePhonelinesOk returns a tuple with the ActivePhonelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePhonelines

`func (o *ExternalDeviceResponse) SetActivePhonelines(v []ActiveRouting)`

SetActivePhonelines sets ActivePhonelines field to given value.

### HasActivePhonelines

`func (o *ExternalDeviceResponse) HasActivePhonelines() bool`

HasActivePhonelines returns a boolean if a field has been set.

### GetActiveGroups

`func (o *ExternalDeviceResponse) GetActiveGroups() []ActiveRouting`

GetActiveGroups returns the ActiveGroups field if non-nil, zero value otherwise.

### GetActiveGroupsOk

`func (o *ExternalDeviceResponse) GetActiveGroupsOk() (*[]ActiveRouting, bool)`

GetActiveGroupsOk returns a tuple with the ActiveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGroups

`func (o *ExternalDeviceResponse) SetActiveGroups(v []ActiveRouting)`

SetActiveGroups sets ActiveGroups field to given value.

### HasActiveGroups

`func (o *ExternalDeviceResponse) HasActiveGroups() bool`

HasActiveGroups returns a boolean if a field has been set.

### GetNumber

`func (o *ExternalDeviceResponse) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ExternalDeviceResponse) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ExternalDeviceResponse) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ExternalDeviceResponse) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetIncomingCallDisplayState

`func (o *ExternalDeviceResponse) GetIncomingCallDisplayState() string`

GetIncomingCallDisplayState returns the IncomingCallDisplayState field if non-nil, zero value otherwise.

### GetIncomingCallDisplayStateOk

`func (o *ExternalDeviceResponse) GetIncomingCallDisplayStateOk() (*string, bool)`

GetIncomingCallDisplayStateOk returns a tuple with the IncomingCallDisplayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingCallDisplayState

`func (o *ExternalDeviceResponse) SetIncomingCallDisplayState(v string)`

SetIncomingCallDisplayState sets IncomingCallDisplayState field to given value.

### HasIncomingCallDisplayState

`func (o *ExternalDeviceResponse) HasIncomingCallDisplayState() bool`

HasIncomingCallDisplayState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


