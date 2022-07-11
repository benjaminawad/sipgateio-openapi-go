# RestrictionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restriction** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewRestrictionItem

`func NewRestrictionItem() *RestrictionItem`

NewRestrictionItem instantiates a new RestrictionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictionItemWithDefaults

`func NewRestrictionItemWithDefaults() *RestrictionItem`

NewRestrictionItemWithDefaults instantiates a new RestrictionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestriction

`func (o *RestrictionItem) GetRestriction() string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *RestrictionItem) GetRestrictionOk() (*string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *RestrictionItem) SetRestriction(v string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *RestrictionItem) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetValue

`func (o *RestrictionItem) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RestrictionItem) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RestrictionItem) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *RestrictionItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTarget

`func (o *RestrictionItem) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RestrictionItem) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RestrictionItem) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RestrictionItem) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


