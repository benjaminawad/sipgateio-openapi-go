# PhonelineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Voicemails** | Pointer to [**[]VoicemailSettingsResponse**](VoicemailSettingsResponse.md) |  | [optional] 

## Methods

### NewPhonelineResponse

`func NewPhonelineResponse() *PhonelineResponse`

NewPhonelineResponse instantiates a new PhonelineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhonelineResponseWithDefaults

`func NewPhonelineResponseWithDefaults() *PhonelineResponse`

NewPhonelineResponseWithDefaults instantiates a new PhonelineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhonelineResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhonelineResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhonelineResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhonelineResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *PhonelineResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PhonelineResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PhonelineResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PhonelineResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetVoicemails

`func (o *PhonelineResponse) GetVoicemails() []VoicemailSettingsResponse`

GetVoicemails returns the Voicemails field if non-nil, zero value otherwise.

### GetVoicemailsOk

`func (o *PhonelineResponse) GetVoicemailsOk() (*[]VoicemailSettingsResponse, bool)`

GetVoicemailsOk returns a tuple with the Voicemails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemails

`func (o *PhonelineResponse) SetVoicemails(v []VoicemailSettingsResponse)`

SetVoicemails sets Voicemails field to given value.

### HasVoicemails

`func (o *PhonelineResponse) HasVoicemails() bool`

HasVoicemails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


