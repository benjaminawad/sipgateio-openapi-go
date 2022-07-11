# SimpleVoicemailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**BelongsToEndpoint** | Pointer to [**EndpointResponse**](EndpointResponse.md) |  | [optional] 

## Methods

### NewSimpleVoicemailResponse

`func NewSimpleVoicemailResponse() *SimpleVoicemailResponse`

NewSimpleVoicemailResponse instantiates a new SimpleVoicemailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleVoicemailResponseWithDefaults

`func NewSimpleVoicemailResponseWithDefaults() *SimpleVoicemailResponse`

NewSimpleVoicemailResponseWithDefaults instantiates a new SimpleVoicemailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleVoicemailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleVoicemailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleVoicemailResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleVoicemailResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *SimpleVoicemailResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *SimpleVoicemailResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *SimpleVoicemailResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *SimpleVoicemailResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetBelongsToEndpoint

`func (o *SimpleVoicemailResponse) GetBelongsToEndpoint() EndpointResponse`

GetBelongsToEndpoint returns the BelongsToEndpoint field if non-nil, zero value otherwise.

### GetBelongsToEndpointOk

`func (o *SimpleVoicemailResponse) GetBelongsToEndpointOk() (*EndpointResponse, bool)`

GetBelongsToEndpointOk returns a tuple with the BelongsToEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelongsToEndpoint

`func (o *SimpleVoicemailResponse) SetBelongsToEndpoint(v EndpointResponse)`

SetBelongsToEndpoint sets BelongsToEndpoint field to given value.

### HasBelongsToEndpoint

`func (o *SimpleVoicemailResponse) HasBelongsToEndpoint() bool`

HasBelongsToEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


