# DetailedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | The status code | [optional] 
**Title** | Pointer to **string** | The title of the error | [optional] 
**Detail** | Pointer to **string** | The detailed description of the error | [optional] 
**Source** | Pointer to [**ErrorSource**](ErrorSource.md) |  | [optional] 

## Methods

### NewDetailedError

`func NewDetailedError() *DetailedError`

NewDetailedError instantiates a new DetailedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedErrorWithDefaults

`func NewDetailedErrorWithDefaults() *DetailedError`

NewDetailedErrorWithDefaults instantiates a new DetailedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DetailedError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DetailedError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DetailedError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DetailedError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *DetailedError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DetailedError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DetailedError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DetailedError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *DetailedError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DetailedError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DetailedError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DetailedError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetSource

`func (o *DetailedError) GetSource() ErrorSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DetailedError) GetSourceOk() (*ErrorSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DetailedError) SetSource(v ErrorSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *DetailedError) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


