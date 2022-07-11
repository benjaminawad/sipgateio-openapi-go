# RegisterDeviceResponse

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
**Credentials** | Pointer to [**CredentialsResponse**](CredentialsResponse.md) |  | [optional] 
**Registered** | Pointer to [**[]RegisteredDevice**](RegisteredDevice.md) |  | [optional] 
**EmergencyAddressId** | Pointer to **string** |  | [optional] 
**AddressUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewRegisterDeviceResponse

`func NewRegisterDeviceResponse() *RegisterDeviceResponse`

NewRegisterDeviceResponse instantiates a new RegisterDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDeviceResponseWithDefaults

`func NewRegisterDeviceResponseWithDefaults() *RegisterDeviceResponse`

NewRegisterDeviceResponseWithDefaults instantiates a new RegisterDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegisterDeviceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterDeviceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterDeviceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegisterDeviceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *RegisterDeviceResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *RegisterDeviceResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *RegisterDeviceResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *RegisterDeviceResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetType

`func (o *RegisterDeviceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegisterDeviceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegisterDeviceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegisterDeviceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOnline

`func (o *RegisterDeviceResponse) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *RegisterDeviceResponse) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *RegisterDeviceResponse) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *RegisterDeviceResponse) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetDnd

`func (o *RegisterDeviceResponse) GetDnd() bool`

GetDnd returns the Dnd field if non-nil, zero value otherwise.

### GetDndOk

`func (o *RegisterDeviceResponse) GetDndOk() (*bool, bool)`

GetDndOk returns a tuple with the Dnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnd

`func (o *RegisterDeviceResponse) SetDnd(v bool)`

SetDnd sets Dnd field to given value.

### HasDnd

`func (o *RegisterDeviceResponse) HasDnd() bool`

HasDnd returns a boolean if a field has been set.

### GetActivePhonelines

`func (o *RegisterDeviceResponse) GetActivePhonelines() []ActiveRouting`

GetActivePhonelines returns the ActivePhonelines field if non-nil, zero value otherwise.

### GetActivePhonelinesOk

`func (o *RegisterDeviceResponse) GetActivePhonelinesOk() (*[]ActiveRouting, bool)`

GetActivePhonelinesOk returns a tuple with the ActivePhonelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePhonelines

`func (o *RegisterDeviceResponse) SetActivePhonelines(v []ActiveRouting)`

SetActivePhonelines sets ActivePhonelines field to given value.

### HasActivePhonelines

`func (o *RegisterDeviceResponse) HasActivePhonelines() bool`

HasActivePhonelines returns a boolean if a field has been set.

### GetActiveGroups

`func (o *RegisterDeviceResponse) GetActiveGroups() []ActiveRouting`

GetActiveGroups returns the ActiveGroups field if non-nil, zero value otherwise.

### GetActiveGroupsOk

`func (o *RegisterDeviceResponse) GetActiveGroupsOk() (*[]ActiveRouting, bool)`

GetActiveGroupsOk returns a tuple with the ActiveGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGroups

`func (o *RegisterDeviceResponse) SetActiveGroups(v []ActiveRouting)`

SetActiveGroups sets ActiveGroups field to given value.

### HasActiveGroups

`func (o *RegisterDeviceResponse) HasActiveGroups() bool`

HasActiveGroups returns a boolean if a field has been set.

### GetCredentials

`func (o *RegisterDeviceResponse) GetCredentials() CredentialsResponse`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RegisterDeviceResponse) GetCredentialsOk() (*CredentialsResponse, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RegisterDeviceResponse) SetCredentials(v CredentialsResponse)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RegisterDeviceResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetRegistered

`func (o *RegisterDeviceResponse) GetRegistered() []RegisteredDevice`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *RegisterDeviceResponse) GetRegisteredOk() (*[]RegisteredDevice, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *RegisterDeviceResponse) SetRegistered(v []RegisteredDevice)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *RegisterDeviceResponse) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.

### GetEmergencyAddressId

`func (o *RegisterDeviceResponse) GetEmergencyAddressId() string`

GetEmergencyAddressId returns the EmergencyAddressId field if non-nil, zero value otherwise.

### GetEmergencyAddressIdOk

`func (o *RegisterDeviceResponse) GetEmergencyAddressIdOk() (*string, bool)`

GetEmergencyAddressIdOk returns a tuple with the EmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressId

`func (o *RegisterDeviceResponse) SetEmergencyAddressId(v string)`

SetEmergencyAddressId sets EmergencyAddressId field to given value.

### HasEmergencyAddressId

`func (o *RegisterDeviceResponse) HasEmergencyAddressId() bool`

HasEmergencyAddressId returns a boolean if a field has been set.

### GetAddressUrl

`func (o *RegisterDeviceResponse) GetAddressUrl() string`

GetAddressUrl returns the AddressUrl field if non-nil, zero value otherwise.

### GetAddressUrlOk

`func (o *RegisterDeviceResponse) GetAddressUrlOk() (*string, bool)`

GetAddressUrlOk returns a tuple with the AddressUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressUrl

`func (o *RegisterDeviceResponse) SetAddressUrl(v string)`

SetAddressUrl sets AddressUrl field to given value.

### HasAddressUrl

`func (o *RegisterDeviceResponse) HasAddressUrl() bool`

HasAddressUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


