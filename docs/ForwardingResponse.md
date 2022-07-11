# ForwardingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewForwardingResponse

`func NewForwardingResponse() *ForwardingResponse`

NewForwardingResponse instantiates a new ForwardingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardingResponseWithDefaults

`func NewForwardingResponseWithDefaults() *ForwardingResponse`

NewForwardingResponseWithDefaults instantiates a new ForwardingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ForwardingResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ForwardingResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ForwardingResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ForwardingResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDestination

`func (o *ForwardingResponse) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ForwardingResponse) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ForwardingResponse) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ForwardingResponse) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetTimeout

`func (o *ForwardingResponse) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ForwardingResponse) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ForwardingResponse) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ForwardingResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetActive

`func (o *ForwardingResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ForwardingResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ForwardingResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ForwardingResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


