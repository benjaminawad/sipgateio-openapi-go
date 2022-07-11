# MobileDeviceResponse

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
**Credentials** | Pointer to [**MobileCredentialsResponse**](MobileCredentialsResponse.md) |  | [optional] 
**SimState** | Pointer to **string** |  | [optional] 
**Esim** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewMobileDeviceResponse

`func NewMobileDeviceResponse() *MobileDeviceResponse`

NewMobileDeviceResponse instantiates a new MobileDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceResponseWithDefaults

`func NewMobileDeviceResponseWithDefaults() *MobileDeviceResponse`

NewMobileDeviceResponseWithDefaults instantiates a new MobileDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobileDeviceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobileDeviceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobileDeviceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MobileDeviceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *MobileDeviceResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *MobileDeviceResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *MobileDeviceResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *MobileDeviceResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetType

`func (o *MobileDeviceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileDeviceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileDeviceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MobileDeviceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOnline

`func (o *MobileDeviceResponse) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *MobileDeviceResponse) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *MobileDeviceResponse) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *MobileDeviceResponse) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetDnd

`func (o *MobileDeviceResponse) GetDnd() bool`

GetDnd returns the Dnd field if non-nil, zero value otherwise.

### GetDndOk

`func (o *MobileDeviceResponse) GetDndOk() (*bool, bool)`

GetDndOk returns a tuple with the Dnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnd

`func (o *MobileDeviceResponse) SetDnd(v bool)`

SetDnd sets Dnd field to given value.

### HasDnd

`func (o *MobileDeviceResponse) HasDnd() bool`

HasDnd returns a boolean if a field has been set.

### GetActivePhonelines

`func (o *MobileDeviceResponse) GetActivePhonelines() []ActiveRouting`

GetActivePhonelines returns the ActivePhonelines field if non-nil, zero value otherwise.

### GetActivePhonelinesOk

`func (o *MobileDeviceResponse) GetActivePhonelinesOk() (*[]ActiveRouting, bool)`

GetActivePhonelinesOk returns a tuple with the ActivePhonelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePhonelines

`func (o *MobileDeviceResponse) SetActivePhonelines(v []ActiveRouting)`

SetActivePhonelines sets ActivePhonelines field to given value.

### HasActivePhonelines

`func (o *MobileDeviceResponse) HasActivePhonelines() bool`

HasActivePhonelines returns a boolean if a field has been set.

### GetActiveGroups

`func (o *MobileDeviceResponse) GetActiveGroups() []ActiveRouting`

GetActiveGroups returns the ActiveGroups field if non-nil, zero value otherwise.

### GetActiveGroupsOk

`func (o *MobileDeviceResponse) GetActiveGroupsOk() (*[]ActiveRouting, bool)`

GetActiveGroupsOk returns a tuple with the ActiveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGroups

`func (o *MobileDeviceResponse) SetActiveGroups(v []ActiveRouting)`

SetActiveGroups sets ActiveGroups field to given value.

### HasActiveGroups

`func (o *MobileDeviceResponse) HasActiveGroups() bool`

HasActiveGroups returns a boolean if a field has been set.

### GetCredentials

`func (o *MobileDeviceResponse) GetCredentials() MobileCredentialsResponse`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *MobileDeviceResponse) GetCredentialsOk() (*MobileCredentialsResponse, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *MobileDeviceResponse) SetCredentials(v MobileCredentialsResponse)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *MobileDeviceResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetSimState

`func (o *MobileDeviceResponse) GetSimState() string`

GetSimState returns the SimState field if non-nil, zero value otherwise.

### GetSimStateOk

`func (o *MobileDeviceResponse) GetSimStateOk() (*string, bool)`

GetSimStateOk returns a tuple with the SimState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimState

`func (o *MobileDeviceResponse) SetSimState(v string)`

SetSimState sets SimState field to given value.

### HasSimState

`func (o *MobileDeviceResponse) HasSimState() bool`

HasSimState returns a boolean if a field has been set.

### GetEsim

`func (o *MobileDeviceResponse) GetEsim() bool`

GetEsim returns the Esim field if non-nil, zero value otherwise.

### GetEsimOk

`func (o *MobileDeviceResponse) GetEsimOk() (*bool, bool)`

GetEsimOk returns a tuple with the Esim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsim

`func (o *MobileDeviceResponse) SetEsim(v bool)`

SetEsim sets Esim field to given value.

### HasEsim

`func (o *MobileDeviceResponse) HasEsim() bool`

HasEsim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


