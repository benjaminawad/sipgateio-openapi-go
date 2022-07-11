# AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** |  | [optional] 
**MainProductType** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] [default to false]
**HistoryLifeTime** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountResponse

`func NewAccountResponse() *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *AccountResponse) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AccountResponse) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AccountResponse) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AccountResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetMainProductType

`func (o *AccountResponse) GetMainProductType() string`

GetMainProductType returns the MainProductType field if non-nil, zero value otherwise.

### GetMainProductTypeOk

`func (o *AccountResponse) GetMainProductTypeOk() (*string, bool)`

GetMainProductTypeOk returns a tuple with the MainProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainProductType

`func (o *AccountResponse) SetMainProductType(v string)`

SetMainProductType sets MainProductType field to given value.

### HasMainProductType

`func (o *AccountResponse) HasMainProductType() bool`

HasMainProductType returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AccountResponse) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AccountResponse) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AccountResponse) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AccountResponse) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetVerified

`func (o *AccountResponse) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AccountResponse) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AccountResponse) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *AccountResponse) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetHistoryLifeTime

`func (o *AccountResponse) GetHistoryLifeTime() string`

GetHistoryLifeTime returns the HistoryLifeTime field if non-nil, zero value otherwise.

### GetHistoryLifeTimeOk

`func (o *AccountResponse) GetHistoryLifeTimeOk() (*string, bool)`

GetHistoryLifeTimeOk returns a tuple with the HistoryLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryLifeTime

`func (o *AccountResponse) SetHistoryLifeTime(v string)`

SetHistoryLifeTime sets HistoryLifeTime field to given value.

### HasHistoryLifeTime

`func (o *AccountResponse) HasHistoryLifeTime() bool`

HasHistoryLifeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


