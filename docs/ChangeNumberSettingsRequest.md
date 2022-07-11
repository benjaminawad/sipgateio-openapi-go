# ChangeNumberSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** |  | [optional] 
**ReleaseForMnp** | Pointer to **bool** |  | [optional] [default to false]
**QuickDial** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewChangeNumberSettingsRequest

`func NewChangeNumberSettingsRequest() *ChangeNumberSettingsRequest`

NewChangeNumberSettingsRequest instantiates a new ChangeNumberSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeNumberSettingsRequestWithDefaults

`func NewChangeNumberSettingsRequestWithDefaults() *ChangeNumberSettingsRequest`

NewChangeNumberSettingsRequestWithDefaults instantiates a new ChangeNumberSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *ChangeNumberSettingsRequest) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ChangeNumberSettingsRequest) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ChangeNumberSettingsRequest) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ChangeNumberSettingsRequest) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetReleaseForMnp

`func (o *ChangeNumberSettingsRequest) GetReleaseForMnp() bool`

GetReleaseForMnp returns the ReleaseForMnp field if non-nil, zero value otherwise.

### GetReleaseForMnpOk

`func (o *ChangeNumberSettingsRequest) GetReleaseForMnpOk() (*bool, bool)`

GetReleaseForMnpOk returns a tuple with the ReleaseForMnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseForMnp

`func (o *ChangeNumberSettingsRequest) SetReleaseForMnp(v bool)`

SetReleaseForMnp sets ReleaseForMnp field to given value.

### HasReleaseForMnp

`func (o *ChangeNumberSettingsRequest) HasReleaseForMnp() bool`

HasReleaseForMnp returns a boolean if a field has been set.

### GetQuickDial

`func (o *ChangeNumberSettingsRequest) GetQuickDial() bool`

GetQuickDial returns the QuickDial field if non-nil, zero value otherwise.

### GetQuickDialOk

`func (o *ChangeNumberSettingsRequest) GetQuickDialOk() (*bool, bool)`

GetQuickDialOk returns a tuple with the QuickDial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickDial

`func (o *ChangeNumberSettingsRequest) SetQuickDial(v bool)`

SetQuickDial sets QuickDial field to given value.

### HasQuickDial

`func (o *ChangeNumberSettingsRequest) HasQuickDial() bool`

HasQuickDial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


