# EventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEventResponse

`func NewEventResponse() *EventResponse`

NewEventResponse instantiates a new EventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseWithDefaults

`func NewEventResponseWithDefaults() *EventResponse`

NewEventResponseWithDefaults instantiates a new EventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *EventResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *EventResponse) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *EventResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventResponse) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventName

`func (o *EventResponse) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventResponse) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventResponse) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *EventResponse) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetPayload

`func (o *EventResponse) GetPayload() map[string]string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EventResponse) GetPayloadOk() (*map[string]string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *EventResponse) SetPayload(v map[string]string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *EventResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


