# VoicemailSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewVoicemailSettingsResponse

`func NewVoicemailSettingsResponse() *VoicemailSettingsResponse`

NewVoicemailSettingsResponse instantiates a new VoicemailSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailSettingsResponseWithDefaults

`func NewVoicemailSettingsResponseWithDefaults() *VoicemailSettingsResponse`

NewVoicemailSettingsResponseWithDefaults instantiates a new VoicemailSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VoicemailSettingsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoicemailSettingsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoicemailSettingsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoicemailSettingsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimeout

`func (o *VoicemailSettingsResponse) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *VoicemailSettingsResponse) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *VoicemailSettingsResponse) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *VoicemailSettingsResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetActive

`func (o *VoicemailSettingsResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *VoicemailSettingsResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *VoicemailSettingsResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *VoicemailSettingsResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


