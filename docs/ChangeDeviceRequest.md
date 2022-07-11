# ChangeDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnd** | Pointer to **bool** |  | [optional] [default to false]
**EmergencyAddressId** | Pointer to **int32** |  | [optional] 

## Methods

### NewChangeDeviceRequest

`func NewChangeDeviceRequest() *ChangeDeviceRequest`

NewChangeDeviceRequest instantiates a new ChangeDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeDeviceRequestWithDefaults

`func NewChangeDeviceRequestWithDefaults() *ChangeDeviceRequest`

NewChangeDeviceRequestWithDefaults instantiates a new ChangeDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnd

`func (o *ChangeDeviceRequest) GetDnd() bool`

GetDnd returns the Dnd field if non-nil, zero value otherwise.

### GetDndOk

`func (o *ChangeDeviceRequest) GetDndOk() (*bool, bool)`

GetDndOk returns a tuple with the Dnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnd

`func (o *ChangeDeviceRequest) SetDnd(v bool)`

SetDnd sets Dnd field to given value.

### HasDnd

`func (o *ChangeDeviceRequest) HasDnd() bool`

HasDnd returns a boolean if a field has been set.

### GetEmergencyAddressId

`func (o *ChangeDeviceRequest) GetEmergencyAddressId() int32`

GetEmergencyAddressId returns the EmergencyAddressId field if non-nil, zero value otherwise.

### GetEmergencyAddressIdOk

`func (o *ChangeDeviceRequest) GetEmergencyAddressIdOk() (*int32, bool)`

GetEmergencyAddressIdOk returns a tuple with the EmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressId

`func (o *ChangeDeviceRequest) SetEmergencyAddressId(v int32)`

SetEmergencyAddressId sets EmergencyAddressId field to given value.

### HasEmergencyAddressId

`func (o *ChangeDeviceRequest) HasEmergencyAddressId() bool`

HasEmergencyAddressId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


