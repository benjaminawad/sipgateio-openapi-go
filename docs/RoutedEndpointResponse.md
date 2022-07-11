# RoutedEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to [**EndpointResponse**](EndpointResponse.md) |  | [optional] 

## Methods

### NewRoutedEndpointResponse

`func NewRoutedEndpointResponse() *RoutedEndpointResponse`

NewRoutedEndpointResponse instantiates a new RoutedEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutedEndpointResponseWithDefaults

`func NewRoutedEndpointResponseWithDefaults() *RoutedEndpointResponse`

NewRoutedEndpointResponseWithDefaults instantiates a new RoutedEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutedEndpointResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutedEndpointResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutedEndpointResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoutedEndpointResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEndpoint

`func (o *RoutedEndpointResponse) GetEndpoint() EndpointResponse`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *RoutedEndpointResponse) GetEndpointOk() (*EndpointResponse, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *RoutedEndpointResponse) SetEndpoint(v EndpointResponse)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *RoutedEndpointResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


