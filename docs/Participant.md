# Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipantId** | **string** | The unique participant identifier | 
**PhoneNumber** | **string** | Participants phone number (can be empty for various reasons, e.g. anonymous participant) | 
**Muted** | **bool** | Participant is muted | 
**Hold** | **bool** | Participant is on hold | 
**Owner** | **bool** | Participant is the call owner | 

## Methods

### NewParticipant

`func NewParticipant(participantId string, phoneNumber string, muted bool, hold bool, owner bool, ) *Participant`

NewParticipant instantiates a new Participant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantWithDefaults

`func NewParticipantWithDefaults() *Participant`

NewParticipantWithDefaults instantiates a new Participant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipantId

`func (o *Participant) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *Participant) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *Participant) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.


### GetPhoneNumber

`func (o *Participant) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Participant) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Participant) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetMuted

`func (o *Participant) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *Participant) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *Participant) SetMuted(v bool)`

SetMuted sets Muted field to given value.


### GetHold

`func (o *Participant) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *Participant) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *Participant) SetHold(v bool)`

SetHold sets Hold field to given value.


### GetOwner

`func (o *Participant) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Participant) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Participant) SetOwner(v bool)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


