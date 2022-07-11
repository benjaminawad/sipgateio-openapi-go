# CallRestrictionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**ForeignSpecial** | Pointer to **bool** |  | [optional] [default to false]
**Charged** | Pointer to **bool** |  | [optional] [default to false]
**Roaming** | Pointer to **bool** |  | [optional] [default to false]
**HighCostForeignNumbers** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCallRestrictionResponse

`func NewCallRestrictionResponse() *CallRestrictionResponse`

NewCallRestrictionResponse instantiates a new CallRestrictionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRestrictionResponseWithDefaults

`func NewCallRestrictionResponseWithDefaults() *CallRestrictionResponse`

NewCallRestrictionResponseWithDefaults instantiates a new CallRestrictionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CallRestrictionResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CallRestrictionResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CallRestrictionResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CallRestrictionResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetForeignSpecial

`func (o *CallRestrictionResponse) GetForeignSpecial() bool`

GetForeignSpecial returns the ForeignSpecial field if non-nil, zero value otherwise.

### GetForeignSpecialOk

`func (o *CallRestrictionResponse) GetForeignSpecialOk() (*bool, bool)`

GetForeignSpecialOk returns a tuple with the ForeignSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignSpecial

`func (o *CallRestrictionResponse) SetForeignSpecial(v bool)`

SetForeignSpecial sets ForeignSpecial field to given value.

### HasForeignSpecial

`func (o *CallRestrictionResponse) HasForeignSpecial() bool`

HasForeignSpecial returns a boolean if a field has been set.

### GetCharged

`func (o *CallRestrictionResponse) GetCharged() bool`

GetCharged returns the Charged field if non-nil, zero value otherwise.

### GetChargedOk

`func (o *CallRestrictionResponse) GetChargedOk() (*bool, bool)`

GetChargedOk returns a tuple with the Charged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharged

`func (o *CallRestrictionResponse) SetCharged(v bool)`

SetCharged sets Charged field to given value.

### HasCharged

`func (o *CallRestrictionResponse) HasCharged() bool`

HasCharged returns a boolean if a field has been set.

### GetRoaming

`func (o *CallRestrictionResponse) GetRoaming() bool`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *CallRestrictionResponse) GetRoamingOk() (*bool, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *CallRestrictionResponse) SetRoaming(v bool)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *CallRestrictionResponse) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.

### GetHighCostForeignNumbers

`func (o *CallRestrictionResponse) GetHighCostForeignNumbers() bool`

GetHighCostForeignNumbers returns the HighCostForeignNumbers field if non-nil, zero value otherwise.

### GetHighCostForeignNumbersOk

`func (o *CallRestrictionResponse) GetHighCostForeignNumbersOk() (*bool, bool)`

GetHighCostForeignNumbersOk returns a tuple with the HighCostForeignNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighCostForeignNumbers

`func (o *CallRestrictionResponse) SetHighCostForeignNumbers(v bool)`

SetHighCostForeignNumbers sets HighCostForeignNumbers field to given value.

### HasHighCostForeignNumbers

`func (o *CallRestrictionResponse) HasHighCostForeignNumbers() bool`

HasHighCostForeignNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


