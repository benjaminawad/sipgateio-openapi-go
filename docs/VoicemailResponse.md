# VoicemailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**BelongsToEndpoint** | Pointer to [**EndpointResponse**](EndpointResponse.md) |  | [optional] 
**Transcription** | Pointer to **bool** |  | [optional] [default to false]
**CanTranscribe** | Pointer to **bool** |  | [optional] [default to false]
**ActiveGreetingAlias** | Pointer to **string** |  | [optional] 
**ActiveGreetingId** | Pointer to **string** |  | [optional] 
**AccessNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewVoicemailResponse

`func NewVoicemailResponse() *VoicemailResponse`

NewVoicemailResponse instantiates a new VoicemailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailResponseWithDefaults

`func NewVoicemailResponseWithDefaults() *VoicemailResponse`

NewVoicemailResponseWithDefaults instantiates a new VoicemailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VoicemailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoicemailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoicemailResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoicemailResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *VoicemailResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *VoicemailResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *VoicemailResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *VoicemailResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetBelongsToEndpoint

`func (o *VoicemailResponse) GetBelongsToEndpoint() EndpointResponse`

GetBelongsToEndpoint returns the BelongsToEndpoint field if non-nil, zero value otherwise.

### GetBelongsToEndpointOk

`func (o *VoicemailResponse) GetBelongsToEndpointOk() (*EndpointResponse, bool)`

GetBelongsToEndpointOk returns a tuple with the BelongsToEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelongsToEndpoint

`func (o *VoicemailResponse) SetBelongsToEndpoint(v EndpointResponse)`

SetBelongsToEndpoint sets BelongsToEndpoint field to given value.

### HasBelongsToEndpoint

`func (o *VoicemailResponse) HasBelongsToEndpoint() bool`

HasBelongsToEndpoint returns a boolean if a field has been set.

### GetTranscription

`func (o *VoicemailResponse) GetTranscription() bool`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *VoicemailResponse) GetTranscriptionOk() (*bool, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *VoicemailResponse) SetTranscription(v bool)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *VoicemailResponse) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetCanTranscribe

`func (o *VoicemailResponse) GetCanTranscribe() bool`

GetCanTranscribe returns the CanTranscribe field if non-nil, zero value otherwise.

### GetCanTranscribeOk

`func (o *VoicemailResponse) GetCanTranscribeOk() (*bool, bool)`

GetCanTranscribeOk returns a tuple with the CanTranscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTranscribe

`func (o *VoicemailResponse) SetCanTranscribe(v bool)`

SetCanTranscribe sets CanTranscribe field to given value.

### HasCanTranscribe

`func (o *VoicemailResponse) HasCanTranscribe() bool`

HasCanTranscribe returns a boolean if a field has been set.

### GetActiveGreetingAlias

`func (o *VoicemailResponse) GetActiveGreetingAlias() string`

GetActiveGreetingAlias returns the ActiveGreetingAlias field if non-nil, zero value otherwise.

### GetActiveGreetingAliasOk

`func (o *VoicemailResponse) GetActiveGreetingAliasOk() (*string, bool)`

GetActiveGreetingAliasOk returns a tuple with the ActiveGreetingAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGreetingAlias

`func (o *VoicemailResponse) SetActiveGreetingAlias(v string)`

SetActiveGreetingAlias sets ActiveGreetingAlias field to given value.

### HasActiveGreetingAlias

`func (o *VoicemailResponse) HasActiveGreetingAlias() bool`

HasActiveGreetingAlias returns a boolean if a field has been set.

### GetActiveGreetingId

`func (o *VoicemailResponse) GetActiveGreetingId() string`

GetActiveGreetingId returns the ActiveGreetingId field if non-nil, zero value otherwise.

### GetActiveGreetingIdOk

`func (o *VoicemailResponse) GetActiveGreetingIdOk() (*string, bool)`

GetActiveGreetingIdOk returns a tuple with the ActiveGreetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGreetingId

`func (o *VoicemailResponse) SetActiveGreetingId(v string)`

SetActiveGreetingId sets ActiveGreetingId field to given value.

### HasActiveGreetingId

`func (o *VoicemailResponse) HasActiveGreetingId() bool`

HasActiveGreetingId returns a boolean if a field has been set.

### GetAccessNumber

`func (o *VoicemailResponse) GetAccessNumber() string`

GetAccessNumber returns the AccessNumber field if non-nil, zero value otherwise.

### GetAccessNumberOk

`func (o *VoicemailResponse) GetAccessNumberOk() (*string, bool)`

GetAccessNumberOk returns a tuple with the AccessNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNumber

`func (o *VoicemailResponse) SetAccessNumber(v string)`

SetAccessNumber sets AccessNumber field to given value.

### HasAccessNumber

`func (o *VoicemailResponse) HasAccessNumber() bool`

HasAccessNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


