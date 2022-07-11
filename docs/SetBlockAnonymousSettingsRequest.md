# SetBlockAnonymousSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewSetBlockAnonymousSettingsRequest

`func NewSetBlockAnonymousSettingsRequest() *SetBlockAnonymousSettingsRequest`

NewSetBlockAnonymousSettingsRequest instantiates a new SetBlockAnonymousSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetBlockAnonymousSettingsRequestWithDefaults

`func NewSetBlockAnonymousSettingsRequestWithDefaults() *SetBlockAnonymousSettingsRequest`

NewSetBlockAnonymousSettingsRequestWithDefaults instantiates a new SetBlockAnonymousSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SetBlockAnonymousSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SetBlockAnonymousSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SetBlockAnonymousSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SetBlockAnonymousSettingsRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTarget

`func (o *SetBlockAnonymousSettingsRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SetBlockAnonymousSettingsRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SetBlockAnonymousSettingsRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SetBlockAnonymousSettingsRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


