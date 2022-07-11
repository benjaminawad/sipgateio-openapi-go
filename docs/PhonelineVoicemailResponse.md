# PhonelineVoicemailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [default to false]
**TranscriptionsUrl** | Pointer to **string** |  | [optional] 
**Transcription** | Pointer to **bool** |  | [optional] [default to false]
**CanTranscribe** | Pointer to **bool** |  | [optional] [default to false]
**GreetingsUrl** | Pointer to **string** |  | [optional] 
**ActiveGreetingAlias** | Pointer to **string** |  | [optional] 
**ActiveGreetingId** | Pointer to **string** |  | [optional] 
**AccessNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewPhonelineVoicemailResponse

`func NewPhonelineVoicemailResponse() *PhonelineVoicemailResponse`

NewPhonelineVoicemailResponse instantiates a new PhonelineVoicemailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhonelineVoicemailResponseWithDefaults

`func NewPhonelineVoicemailResponseWithDefaults() *PhonelineVoicemailResponse`

NewPhonelineVoicemailResponseWithDefaults instantiates a new PhonelineVoicemailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhonelineVoicemailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhonelineVoicemailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhonelineVoicemailResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhonelineVoicemailResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *PhonelineVoicemailResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PhonelineVoicemailResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PhonelineVoicemailResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PhonelineVoicemailResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetTimeout

`func (o *PhonelineVoicemailResponse) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PhonelineVoicemailResponse) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PhonelineVoicemailResponse) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PhonelineVoicemailResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetActive

`func (o *PhonelineVoicemailResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PhonelineVoicemailResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PhonelineVoicemailResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PhonelineVoicemailResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTranscriptionsUrl

`func (o *PhonelineVoicemailResponse) GetTranscriptionsUrl() string`

GetTranscriptionsUrl returns the TranscriptionsUrl field if non-nil, zero value otherwise.

### GetTranscriptionsUrlOk

`func (o *PhonelineVoicemailResponse) GetTranscriptionsUrlOk() (*string, bool)`

GetTranscriptionsUrlOk returns a tuple with the TranscriptionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionsUrl

`func (o *PhonelineVoicemailResponse) SetTranscriptionsUrl(v string)`

SetTranscriptionsUrl sets TranscriptionsUrl field to given value.

### HasTranscriptionsUrl

`func (o *PhonelineVoicemailResponse) HasTranscriptionsUrl() bool`

HasTranscriptionsUrl returns a boolean if a field has been set.

### GetTranscription

`func (o *PhonelineVoicemailResponse) GetTranscription() bool`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *PhonelineVoicemailResponse) GetTranscriptionOk() (*bool, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *PhonelineVoicemailResponse) SetTranscription(v bool)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *PhonelineVoicemailResponse) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetCanTranscribe

`func (o *PhonelineVoicemailResponse) GetCanTranscribe() bool`

GetCanTranscribe returns the CanTranscribe field if non-nil, zero value otherwise.

### GetCanTranscribeOk

`func (o *PhonelineVoicemailResponse) GetCanTranscribeOk() (*bool, bool)`

GetCanTranscribeOk returns a tuple with the CanTranscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTranscribe

`func (o *PhonelineVoicemailResponse) SetCanTranscribe(v bool)`

SetCanTranscribe sets CanTranscribe field to given value.

### HasCanTranscribe

`func (o *PhonelineVoicemailResponse) HasCanTranscribe() bool`

HasCanTranscribe returns a boolean if a field has been set.

### GetGreetingsUrl

`func (o *PhonelineVoicemailResponse) GetGreetingsUrl() string`

GetGreetingsUrl returns the GreetingsUrl field if non-nil, zero value otherwise.

### GetGreetingsUrlOk

`func (o *PhonelineVoicemailResponse) GetGreetingsUrlOk() (*string, bool)`

GetGreetingsUrlOk returns a tuple with the GreetingsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingsUrl

`func (o *PhonelineVoicemailResponse) SetGreetingsUrl(v string)`

SetGreetingsUrl sets GreetingsUrl field to given value.

### HasGreetingsUrl

`func (o *PhonelineVoicemailResponse) HasGreetingsUrl() bool`

HasGreetingsUrl returns a boolean if a field has been set.

### GetActiveGreetingAlias

`func (o *PhonelineVoicemailResponse) GetActiveGreetingAlias() string`

GetActiveGreetingAlias returns the ActiveGreetingAlias field if non-nil, zero value otherwise.

### GetActiveGreetingAliasOk

`func (o *PhonelineVoicemailResponse) GetActiveGreetingAliasOk() (*string, bool)`

GetActiveGreetingAliasOk returns a tuple with the ActiveGreetingAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGreetingAlias

`func (o *PhonelineVoicemailResponse) SetActiveGreetingAlias(v string)`

SetActiveGreetingAlias sets ActiveGreetingAlias field to given value.

### HasActiveGreetingAlias

`func (o *PhonelineVoicemailResponse) HasActiveGreetingAlias() bool`

HasActiveGreetingAlias returns a boolean if a field has been set.

### GetActiveGreetingId

`func (o *PhonelineVoicemailResponse) GetActiveGreetingId() string`

GetActiveGreetingId returns the ActiveGreetingId field if non-nil, zero value otherwise.

### GetActiveGreetingIdOk

`func (o *PhonelineVoicemailResponse) GetActiveGreetingIdOk() (*string, bool)`

GetActiveGreetingIdOk returns a tuple with the ActiveGreetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGreetingId

`func (o *PhonelineVoicemailResponse) SetActiveGreetingId(v string)`

SetActiveGreetingId sets ActiveGreetingId field to given value.

### HasActiveGreetingId

`func (o *PhonelineVoicemailResponse) HasActiveGreetingId() bool`

HasActiveGreetingId returns a boolean if a field has been set.

### GetAccessNumber

`func (o *PhonelineVoicemailResponse) GetAccessNumber() string`

GetAccessNumber returns the AccessNumber field if non-nil, zero value otherwise.

### GetAccessNumberOk

`func (o *PhonelineVoicemailResponse) GetAccessNumberOk() (*string, bool)`

GetAccessNumberOk returns a tuple with the AccessNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNumber

`func (o *PhonelineVoicemailResponse) SetAccessNumber(v string)`

SetAccessNumber sets AccessNumber field to given value.

### HasAccessNumber

`func (o *PhonelineVoicemailResponse) HasAccessNumber() bool`

HasAccessNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


