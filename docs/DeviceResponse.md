# DeviceResponse

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

## Methods

### NewDeviceResponse

`func NewDeviceResponse() *DeviceResponse`

NewDeviceResponse instantiates a new DeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceResponseWithDefaults

`func NewDeviceResponseWithDefaults() *DeviceResponse`

NewDeviceResponseWithDefaults instantiates a new DeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *DeviceResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DeviceResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DeviceResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DeviceResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetType

`func (o *DeviceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOnline

`func (o *DeviceResponse) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *DeviceResponse) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *DeviceResponse) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *DeviceResponse) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetDnd

`func (o *DeviceResponse) GetDnd() bool`

GetDnd returns the Dnd field if non-nil, zero value otherwise.

### GetDndOk

`func (o *DeviceResponse) GetDndOk() (*bool, bool)`

GetDndOk returns a tuple with the Dnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnd

`func (o *DeviceResponse) SetDnd(v bool)`

SetDnd sets Dnd field to given value.

### HasDnd

`func (o *DeviceResponse) HasDnd() bool`

HasDnd returns a boolean if a field has been set.

### GetActivePhonelines

`func (o *DeviceResponse) GetActivePhonelines() []ActiveRouting`

GetActivePhonelines returns the ActivePhonelines field if non-nil, zero value otherwise.

### GetActivePhonelinesOk

`func (o *DeviceResponse) GetActivePhonelinesOk() (*[]ActiveRouting, bool)`

GetActivePhonelinesOk returns a tuple with the ActivePhonelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePhonelines

`func (o *DeviceResponse) SetActivePhonelines(v []ActiveRouting)`

SetActivePhonelines sets ActivePhonelines field to given value.

### HasActivePhonelines

`func (o *DeviceResponse) HasActivePhonelines() bool`

HasActivePhonelines returns a boolean if a field has been set.

### GetActiveGroups

`func (o *DeviceResponse) GetActiveGroups() []ActiveRouting`

GetActiveGroups returns the ActiveGroups field if non-nil, zero value otherwise.

### GetActiveGroupsOk

`func (o *DeviceResponse) GetActiveGroupsOk() (*[]ActiveRouting, bool)`

GetActiveGroupsOk returns a tuple with the ActiveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGroups

`func (o *DeviceResponse) SetActiveGroups(v []ActiveRouting)`

SetActiveGroups sets ActiveGroups field to given value.

### HasActiveGroups

`func (o *DeviceResponse) HasActiveGroups() bool`

HasActiveGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


