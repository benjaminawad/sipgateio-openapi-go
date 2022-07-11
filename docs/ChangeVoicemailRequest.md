# ChangeVoicemailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** |  | [optional] 
**Active** | **bool** |  | [default to false]
**Transcription** | **bool** |  | [default to false]

## Methods

### NewChangeVoicemailRequest

`func NewChangeVoicemailRequest(active bool, transcription bool, ) *ChangeVoicemailRequest`

NewChangeVoicemailRequest instantiates a new ChangeVoicemailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeVoicemailRequestWithDefaults

`func NewChangeVoicemailRequestWithDefaults() *ChangeVoicemailRequest`

NewChangeVoicemailRequestWithDefaults instantiates a new ChangeVoicemailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *ChangeVoicemailRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ChangeVoicemailRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ChangeVoicemailRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ChangeVoicemailRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetActive

`func (o *ChangeVoicemailRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ChangeVoicemailRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ChangeVoicemailRequest) SetActive(v bool)`

SetActive sets Active field to given value.


### GetTranscription

`func (o *ChangeVoicemailRequest) GetTranscription() bool`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *ChangeVoicemailRequest) GetTranscriptionOk() (*bool, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *ChangeVoicemailRequest) SetTranscription(v bool)`

SetTranscription sets Transcription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


