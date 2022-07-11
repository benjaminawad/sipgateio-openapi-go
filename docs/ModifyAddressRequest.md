# ModifyAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countrycode** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**Postcode** | **string** |  | 
**City** | **string** |  | 
**Street** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 

## Methods

### NewModifyAddressRequest

`func NewModifyAddressRequest(countrycode string, postcode string, city string, ) *ModifyAddressRequest`

NewModifyAddressRequest instantiates a new ModifyAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyAddressRequestWithDefaults

`func NewModifyAddressRequestWithDefaults() *ModifyAddressRequest`

NewModifyAddressRequestWithDefaults instantiates a new ModifyAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountrycode

`func (o *ModifyAddressRequest) GetCountrycode() string`

GetCountrycode returns the Countrycode field if non-nil, zero value otherwise.

### GetCountrycodeOk

`func (o *ModifyAddressRequest) GetCountrycodeOk() (*string, bool)`

GetCountrycodeOk returns a tuple with the Countrycode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrycode

`func (o *ModifyAddressRequest) SetCountrycode(v string)`

SetCountrycode sets Countrycode field to given value.


### GetState

`func (o *ModifyAddressRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModifyAddressRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModifyAddressRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ModifyAddressRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostcode

`func (o *ModifyAddressRequest) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *ModifyAddressRequest) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *ModifyAddressRequest) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetCity

`func (o *ModifyAddressRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ModifyAddressRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ModifyAddressRequest) SetCity(v string)`

SetCity sets City field to given value.


### GetStreet

`func (o *ModifyAddressRequest) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *ModifyAddressRequest) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *ModifyAddressRequest) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *ModifyAddressRequest) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetNumber

`func (o *ModifyAddressRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ModifyAddressRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ModifyAddressRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ModifyAddressRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAddress1

`func (o *ModifyAddressRequest) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ModifyAddressRequest) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ModifyAddressRequest) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ModifyAddressRequest) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *ModifyAddressRequest) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ModifyAddressRequest) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ModifyAddressRequest) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ModifyAddressRequest) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


