# NewCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Caller** | **string** |  | 
**Callee** | **string** |  | 
**CallerId** | Pointer to **string** |  | [optional] 

## Methods

### NewNewCallRequest

`func NewNewCallRequest(caller string, callee string, ) *NewCallRequest`

NewNewCallRequest instantiates a new NewCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCallRequestWithDefaults

`func NewNewCallRequestWithDefaults() *NewCallRequest`

NewNewCallRequestWithDefaults instantiates a new NewCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *NewCallRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *NewCallRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *NewCallRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *NewCallRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetCaller

`func (o *NewCallRequest) GetCaller() string`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *NewCallRequest) GetCallerOk() (*string, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *NewCallRequest) SetCaller(v string)`

SetCaller sets Caller field to given value.


### GetCallee

`func (o *NewCallRequest) GetCallee() string`

GetCallee returns the Callee field if non-nil, zero value otherwise.

### GetCalleeOk

`func (o *NewCallRequest) GetCalleeOk() (*string, bool)`

GetCalleeOk returns a tuple with the Callee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallee

`func (o *NewCallRequest) SetCallee(v string)`

SetCallee sets Callee field to given value.


### GetCallerId

`func (o *NewCallRequest) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *NewCallRequest) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *NewCallRequest) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *NewCallRequest) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


