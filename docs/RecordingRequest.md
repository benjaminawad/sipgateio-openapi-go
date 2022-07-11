# RecordingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Announcement** | **bool** | Announce recording start/stop to all participants | 
**Value** | **bool** | Start/stop the recording | 

## Methods

### NewRecordingRequest

`func NewRecordingRequest(announcement bool, value bool, ) *RecordingRequest`

NewRecordingRequest instantiates a new RecordingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingRequestWithDefaults

`func NewRecordingRequestWithDefaults() *RecordingRequest`

NewRecordingRequestWithDefaults instantiates a new RecordingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnouncement

`func (o *RecordingRequest) GetAnnouncement() bool`

GetAnnouncement returns the Announcement field if non-nil, zero value otherwise.

### GetAnnouncementOk

`func (o *RecordingRequest) GetAnnouncementOk() (*bool, bool)`

GetAnnouncementOk returns a tuple with the Announcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncement

`func (o *RecordingRequest) SetAnnouncement(v bool)`

SetAnnouncement sets Announcement field to given value.


### GetValue

`func (o *RecordingRequest) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RecordingRequest) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RecordingRequest) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


