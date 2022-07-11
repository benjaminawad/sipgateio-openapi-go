# OAuthClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The unique OAuth 2.0 client identifier | 
**ClientSecret** | **string** | The OAuth 2.0 client secret | 
**Name** | **string** | Name of the application the client is intended to use with | 
**Description** | **string** | Short description of the application the client is intended to use with | 
**RedirectUris** | **[]string** | Valid URI pattern a browser can redirect to after a successful login or logout (simple wildcards are allowed e.g. &#39;http://example.com/_*&#39;) | 
**WebOrigins** | **[]string** | Allowed CORS origins (simple wildcards are allowed e.g. &#39;http://_*.example.com&#39;) | 
**PrivacyUrl** | **string** | The privacy policy URL | 
**TermsUrl** | **string** | The terms and conditions URL | 

## Methods

### NewOAuthClient

`func NewOAuthClient(clientId string, clientSecret string, name string, description string, redirectUris []string, webOrigins []string, privacyUrl string, termsUrl string, ) *OAuthClient`

NewOAuthClient instantiates a new OAuthClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientWithDefaults

`func NewOAuthClientWithDefaults() *OAuthClient`

NewOAuthClientWithDefaults instantiates a new OAuthClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuthClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuthClient) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthClient) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthClient) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetName

`func (o *OAuthClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuthClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuthClient) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OAuthClient) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OAuthClient) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OAuthClient) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRedirectUris

`func (o *OAuthClient) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OAuthClient) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OAuthClient) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetWebOrigins

`func (o *OAuthClient) GetWebOrigins() []string`

GetWebOrigins returns the WebOrigins field if non-nil, zero value otherwise.

### GetWebOriginsOk

`func (o *OAuthClient) GetWebOriginsOk() (*[]string, bool)`

GetWebOriginsOk returns a tuple with the WebOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebOrigins

`func (o *OAuthClient) SetWebOrigins(v []string)`

SetWebOrigins sets WebOrigins field to given value.


### GetPrivacyUrl

`func (o *OAuthClient) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *OAuthClient) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *OAuthClient) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.


### GetTermsUrl

`func (o *OAuthClient) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *OAuthClient) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *OAuthClient) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


