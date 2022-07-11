# SetParallelForwardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSetParallelForwardingRequest

`func NewSetParallelForwardingRequest() *SetParallelForwardingRequest`

NewSetParallelForwardingRequest instantiates a new SetParallelForwardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetParallelForwardingRequestWithDefaults

`func NewSetParallelForwardingRequestWithDefaults() *SetParallelForwardingRequest`

NewSetParallelForwardingRequestWithDefaults instantiates a new SetParallelForwardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *SetParallelForwardingRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SetParallelForwardingRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SetParallelForwardingRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SetParallelForwardingRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetAlias

`func (o *SetParallelForwardingRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *SetParallelForwardingRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *SetParallelForwardingRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *SetParallelForwardingRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetActive

`func (o *SetParallelForwardingRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SetParallelForwardingRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SetParallelForwardingRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SetParallelForwardingRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


