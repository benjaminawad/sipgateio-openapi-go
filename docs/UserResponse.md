# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DefaultDevice** | Pointer to **string** |  | [optional] 
**BusyOnBusy** | Pointer to **bool** |  | [optional] [default to false]
**AddressId** | Pointer to **string** |  | [optional] 
**DirectDialIds** | Pointer to **[]string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewUserResponse

`func NewUserResponse() *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstname

`func (o *UserResponse) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserResponse) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserResponse) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserResponse) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserResponse) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserResponse) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserResponse) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserResponse) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetEmail

`func (o *UserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDefaultDevice

`func (o *UserResponse) GetDefaultDevice() string`

GetDefaultDevice returns the DefaultDevice field if non-nil, zero value otherwise.

### GetDefaultDeviceOk

`func (o *UserResponse) GetDefaultDeviceOk() (*string, bool)`

GetDefaultDeviceOk returns a tuple with the DefaultDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDevice

`func (o *UserResponse) SetDefaultDevice(v string)`

SetDefaultDevice sets DefaultDevice field to given value.

### HasDefaultDevice

`func (o *UserResponse) HasDefaultDevice() bool`

HasDefaultDevice returns a boolean if a field has been set.

### GetBusyOnBusy

`func (o *UserResponse) GetBusyOnBusy() bool`

GetBusyOnBusy returns the BusyOnBusy field if non-nil, zero value otherwise.

### GetBusyOnBusyOk

`func (o *UserResponse) GetBusyOnBusyOk() (*bool, bool)`

GetBusyOnBusyOk returns a tuple with the BusyOnBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyOnBusy

`func (o *UserResponse) SetBusyOnBusy(v bool)`

SetBusyOnBusy sets BusyOnBusy field to given value.

### HasBusyOnBusy

`func (o *UserResponse) HasBusyOnBusy() bool`

HasBusyOnBusy returns a boolean if a field has been set.

### GetAddressId

`func (o *UserResponse) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *UserResponse) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *UserResponse) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *UserResponse) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetDirectDialIds

`func (o *UserResponse) GetDirectDialIds() []string`

GetDirectDialIds returns the DirectDialIds field if non-nil, zero value otherwise.

### GetDirectDialIdsOk

`func (o *UserResponse) GetDirectDialIdsOk() (*[]string, bool)`

GetDirectDialIdsOk returns a tuple with the DirectDialIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDialIds

`func (o *UserResponse) SetDirectDialIds(v []string)`

SetDirectDialIds sets DirectDialIds field to given value.

### HasDirectDialIds

`func (o *UserResponse) HasDirectDialIds() bool`

HasDirectDialIds returns a boolean if a field has been set.

### GetTimezone

`func (o *UserResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAdmin

`func (o *UserResponse) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserResponse) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserResponse) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserResponse) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


