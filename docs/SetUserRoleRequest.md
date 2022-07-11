# SetUserRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSetUserRoleRequest

`func NewSetUserRoleRequest() *SetUserRoleRequest`

NewSetUserRoleRequest instantiates a new SetUserRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetUserRoleRequestWithDefaults

`func NewSetUserRoleRequestWithDefaults() *SetUserRoleRequest`

NewSetUserRoleRequestWithDefaults instantiates a new SetUserRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *SetUserRoleRequest) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *SetUserRoleRequest) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *SetUserRoleRequest) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *SetUserRoleRequest) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


