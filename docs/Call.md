# Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallId** | **string** | The unique call identifier | 
**Muted** | **bool** | The call is muted | 
**Recording** | **bool** | The call is being recorded | 
**Hold** | **bool** | The call is on hold | 
**Participants** | [**[]Participant**](Participant.md) |  | 

## Methods

### NewCall

`func NewCall(callId string, muted bool, recording bool, hold bool, participants []Participant, ) *Call`

NewCall instantiates a new Call object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallWithDefaults

`func NewCallWithDefaults() *Call`

NewCallWithDefaults instantiates a new Call object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallId

`func (o *Call) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *Call) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *Call) SetCallId(v string)`

SetCallId sets CallId field to given value.


### GetMuted

`func (o *Call) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *Call) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *Call) SetMuted(v bool)`

SetMuted sets Muted field to given value.


### GetRecording

`func (o *Call) GetRecording() bool`

GetRecording returns the Recording field if non-nil, zero value otherwise.

### GetRecordingOk

`func (o *Call) GetRecordingOk() (*bool, bool)`

GetRecordingOk returns a tuple with the Recording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecording

`func (o *Call) SetRecording(v bool)`

SetRecording sets Recording field to given value.


### GetHold

`func (o *Call) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *Call) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *Call) SetHold(v bool)`

SetHold sets Hold field to given value.


### GetParticipants

`func (o *Call) GetParticipants() []Participant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *Call) GetParticipantsOk() (*[]Participant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *Call) SetParticipants(v []Participant)`

SetParticipants sets Participants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


