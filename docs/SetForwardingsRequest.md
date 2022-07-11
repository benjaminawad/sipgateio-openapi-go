# SetForwardingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forwardings** | Pointer to [**[]Forwarding**](Forwarding.md) |  | [optional] 

## Methods

### NewSetForwardingsRequest

`func NewSetForwardingsRequest() *SetForwardingsRequest`

NewSetForwardingsRequest instantiates a new SetForwardingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetForwardingsRequestWithDefaults

`func NewSetForwardingsRequestWithDefaults() *SetForwardingsRequest`

NewSetForwardingsRequestWithDefaults instantiates a new SetForwardingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardings

`func (o *SetForwardingsRequest) GetForwardings() []Forwarding`

GetForwardings returns the Forwardings field if non-nil, zero value otherwise.

### GetForwardingsOk

`func (o *SetForwardingsRequest) GetForwardingsOk() (*[]Forwarding, bool)`

GetForwardingsOk returns a tuple with the Forwardings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardings

`func (o *SetForwardingsRequest) SetForwardings(v []Forwarding)`

SetForwardings sets Forwardings field to given value.

### HasForwardings

`func (o *SetForwardingsRequest) HasForwardings() bool`

HasForwardings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


