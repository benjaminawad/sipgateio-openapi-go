# CreateClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the application the client is intended to use with | 
**Description** | **string** | Short description of the application the client is intended to use with | 
**RedirectUris** | **[]string** | Valid URI pattern a browser can redirect to after a successful login or logout (simple wildcards are allowed e.g. &#39;http://example.com/_*&#39;) | 
**WebOrigins** | **[]string** | Allowed CORS origins (simple wildcards are allowed e.g. &#39;http://_*.example.com&#39;) | 
**PrivacyUrl** | **string** | The privacy policy URL | 
**TermsUrl** | **string** | The terms and conditions URL | 

## Methods

### NewCreateClientRequest

`func NewCreateClientRequest(name string, description string, redirectUris []string, webOrigins []string, privacyUrl string, termsUrl string, ) *CreateClientRequest`

NewCreateClientRequest instantiates a new CreateClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClientRequestWithDefaults

`func NewCreateClientRequestWithDefaults() *CreateClientRequest`

NewCreateClientRequestWithDefaults instantiates a new CreateClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClientRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateClientRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateClientRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateClientRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRedirectUris

`func (o *CreateClientRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *CreateClientRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *CreateClientRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetWebOrigins

`func (o *CreateClientRequest) GetWebOrigins() []string`

GetWebOrigins returns the WebOrigins field if non-nil, zero value otherwise.

### GetWebOriginsOk

`func (o *CreateClientRequest) GetWebOriginsOk() (*[]string, bool)`

GetWebOriginsOk returns a tuple with the WebOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebOrigins

`func (o *CreateClientRequest) SetWebOrigins(v []string)`

SetWebOrigins sets WebOrigins field to given value.


### GetPrivacyUrl

`func (o *CreateClientRequest) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *CreateClientRequest) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *CreateClientRequest) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.


### GetTermsUrl

`func (o *CreateClientRequest) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *CreateClientRequest) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *CreateClientRequest) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


