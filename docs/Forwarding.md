# Forwarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewForwarding

`func NewForwarding() *Forwarding`

NewForwarding instantiates a new Forwarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardingWithDefaults

`func NewForwardingWithDefaults() *Forwarding`

NewForwardingWithDefaults instantiates a new Forwarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *Forwarding) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Forwarding) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Forwarding) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Forwarding) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetTimeout

`func (o *Forwarding) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Forwarding) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Forwarding) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Forwarding) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetActive

`func (o *Forwarding) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Forwarding) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Forwarding) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Forwarding) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


