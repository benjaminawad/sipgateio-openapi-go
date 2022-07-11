# ApiUrlsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountUrl** | Pointer to **string** |  | [optional] 
**AccountVerifiedUrl** | Pointer to **string** |  | [optional] 
**AddressUrl** | Pointer to **string** |  | [optional] 
**AddressesUrl** | Pointer to **string** |  | [optional] 
**AddressNumbersUrl** | Pointer to **string** |  | [optional] 
**AuthorizationOauthClientsUrl** | Pointer to **string** |  | [optional] 
**AuthorizationOauthClientUrl** | Pointer to **string** |  | [optional] 
**AuthorizationTokenUrl** | Pointer to **string** |  | [optional] 
**BalanceUrl** | Pointer to **string** |  | [optional] 
**GetCallRestrictionsUrl** | Pointer to **string** |  | [optional] 
**CallRestrictionsUrl** | Pointer to **string** |  | [optional] 
**ContactsUrl** | Pointer to **string** |  | [optional] 
**CrmBridgeUrl** | Pointer to **string** |  | [optional] 
**InternalContactsUrl** | Pointer to **string** |  | [optional] 
**DeviceCallerIdUrl** | Pointer to **string** |  | [optional] 
**DeviceUrl** | Pointer to **string** |  | [optional] 
**DevicesTariffAnnouncementUrl** | Pointer to **string** |  | [optional] 
**DevicesSingleRowDisplayUrl** | Pointer to **string** |  | [optional] 
**DeviceSimOrdersUrl** | Pointer to **string** |  | [optional] 
**FaxlinesUrl** | Pointer to **string** |  | [optional] 
**FaxlineUrl** | Pointer to **string** |  | [optional] 
**FaxlineNumbersUrl** | Pointer to **string** |  | [optional] 
**GroupsUrl** | Pointer to **string** |  | [optional] 
**GroupDetailUrl** | Pointer to **string** |  | [optional] 
**GroupUserUrl** | Pointer to **string** |  | [optional] 
**GroupDevicesUrl** | Pointer to **string** |  | [optional] 
**GroupNumbersUrl** | Pointer to **string** |  | [optional] 
**HistoryEntryUrl** | Pointer to **string** |  | [optional] 
**HistoryUrl** | Pointer to **string** |  | [optional] 
**NotificationsUrl** | Pointer to **string** |  | [optional] 
**NumberUrl** | Pointer to **string** |  | [optional] 
**NumbersUrl** | Pointer to **string** |  | [optional] 
**WebuserNumbersUrl** | Pointer to **string** |  | [optional] 
**AddDirectDialUrl** | Pointer to **string** |  | [optional] 
**ChangeDirectDialUrl** | Pointer to **string** |  | [optional] 
**DeleteDirectDialUrl** | Pointer to **string** |  | [optional] 
**PhonelineUrl** | Pointer to **string** |  | [optional] 
**PhonelineBlockAnonymousUrl** | Pointer to **string** |  | [optional] 
**PhonelineDevicesUrl** | Pointer to **string** |  | [optional] 
**PhonelineForwardingsUrl** | Pointer to **string** |  | [optional] 
**PhonelineNumbersUrl** | Pointer to **string** |  | [optional] 
**PhonelinesParallelForwardingUrl** | Pointer to **string** |  | [optional] 
**PhonelinesParallelForwardingsUrl** | Pointer to **string** |  | [optional] 
**PhonelineSipgateIoUrl** | Pointer to **string** |  | [optional] 
**PhonelineSipgateIoLogUrl** | Pointer to **string** |  | [optional] 
**PhonelineVoicemailGreetingUrl** | Pointer to **string** |  | [optional] 
**PhonelineVoicemailGreetingsUrl** | Pointer to **string** |  | [optional] 
**PhonelineVoicemailUrl** | Pointer to **string** |  | [optional] 
**PhonelineVoicemailsUrl** | Pointer to **string** |  | [optional] 
**PhonelinesUrl** | Pointer to **string** |  | [optional] 
**PhonelineDetailUrl** | Pointer to **string** |  | [optional] 
**PortingUrl** | Pointer to **string** |  | [optional] 
**PortingsUrl** | Pointer to **string** |  | [optional] 
**RestrictionsUrl** | Pointer to **string** |  | [optional] 
**SessionsCallUrl** | Pointer to **string** |  | [optional] 
**SettingsSipgateioUrl** | Pointer to **string** |  | [optional] 
**SmsUrl** | Pointer to **string** |  | [optional] 
**SmsCallerIdsUrl** | Pointer to **string** |  | [optional] 
**TranslationsUrl** | Pointer to **string** |  | [optional] 
**UserInfoUrl** | Pointer to **string** |  | [optional] 
**UserUrl** | Pointer to **string** |  | [optional] 
**UserBusyOnBusyUrl** | Pointer to **string** |  | [optional] 
**UserDefaultDeviceUrl** | Pointer to **string** |  | [optional] 
**UserDevicesUrl** | Pointer to **string** |  | [optional] 
**UserDeviceSimUrl** | Pointer to **string** |  | [optional] 
**UserExternalDevicesUrl** | Pointer to **string** |  | [optional] 
**UserMobileDevicesUrl** | Pointer to **string** |  | [optional] 
**UserRegisterDevicesUrl** | Pointer to **string** |  | [optional] 
**UserRoleUrl** | Pointer to **string** |  | [optional] 
**UsersUrl** | Pointer to **string** |  | [optional] 
**VoicemailsUrl** | Pointer to **string** |  | [optional] 
**VoicemailDetailUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewApiUrlsResponse

`func NewApiUrlsResponse() *ApiUrlsResponse`

NewApiUrlsResponse instantiates a new ApiUrlsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUrlsResponseWithDefaults

`func NewApiUrlsResponseWithDefaults() *ApiUrlsResponse`

NewApiUrlsResponseWithDefaults instantiates a new ApiUrlsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountUrl

`func (o *ApiUrlsResponse) GetAccountUrl() string`

GetAccountUrl returns the AccountUrl field if non-nil, zero value otherwise.

### GetAccountUrlOk

`func (o *ApiUrlsResponse) GetAccountUrlOk() (*string, bool)`

GetAccountUrlOk returns a tuple with the AccountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUrl

`func (o *ApiUrlsResponse) SetAccountUrl(v string)`

SetAccountUrl sets AccountUrl field to given value.

### HasAccountUrl

`func (o *ApiUrlsResponse) HasAccountUrl() bool`

HasAccountUrl returns a boolean if a field has been set.

### GetAccountVerifiedUrl

`func (o *ApiUrlsResponse) GetAccountVerifiedUrl() string`

GetAccountVerifiedUrl returns the AccountVerifiedUrl field if non-nil, zero value otherwise.

### GetAccountVerifiedUrlOk

`func (o *ApiUrlsResponse) GetAccountVerifiedUrlOk() (*string, bool)`

GetAccountVerifiedUrlOk returns a tuple with the AccountVerifiedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountVerifiedUrl

`func (o *ApiUrlsResponse) SetAccountVerifiedUrl(v string)`

SetAccountVerifiedUrl sets AccountVerifiedUrl field to given value.

### HasAccountVerifiedUrl

`func (o *ApiUrlsResponse) HasAccountVerifiedUrl() bool`

HasAccountVerifiedUrl returns a boolean if a field has been set.

### GetAddressUrl

`func (o *ApiUrlsResponse) GetAddressUrl() string`

GetAddressUrl returns the AddressUrl field if non-nil, zero value otherwise.

### GetAddressUrlOk

`func (o *ApiUrlsResponse) GetAddressUrlOk() (*string, bool)`

GetAddressUrlOk returns a tuple with the AddressUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressUrl

`func (o *ApiUrlsResponse) SetAddressUrl(v string)`

SetAddressUrl sets AddressUrl field to given value.

### HasAddressUrl

`func (o *ApiUrlsResponse) HasAddressUrl() bool`

HasAddressUrl returns a boolean if a field has been set.

### GetAddressesUrl

`func (o *ApiUrlsResponse) GetAddressesUrl() string`

GetAddressesUrl returns the AddressesUrl field if non-nil, zero value otherwise.

### GetAddressesUrlOk

`func (o *ApiUrlsResponse) GetAddressesUrlOk() (*string, bool)`

GetAddressesUrlOk returns a tuple with the AddressesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressesUrl

`func (o *ApiUrlsResponse) SetAddressesUrl(v string)`

SetAddressesUrl sets AddressesUrl field to given value.

### HasAddressesUrl

`func (o *ApiUrlsResponse) HasAddressesUrl() bool`

HasAddressesUrl returns a boolean if a field has been set.

### GetAddressNumbersUrl

`func (o *ApiUrlsResponse) GetAddressNumbersUrl() string`

GetAddressNumbersUrl returns the AddressNumbersUrl field if non-nil, zero value otherwise.

### GetAddressNumbersUrlOk

`func (o *ApiUrlsResponse) GetAddressNumbersUrlOk() (*string, bool)`

GetAddressNumbersUrlOk returns a tuple with the AddressNumbersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressNumbersUrl

`func (o *ApiUrlsResponse) SetAddressNumbersUrl(v string)`

SetAddressNumbersUrl sets AddressNumbersUrl field to given value.

### HasAddressNumbersUrl

`func (o *ApiUrlsResponse) HasAddressNumbersUrl() bool`

HasAddressNumbersUrl returns a boolean if a field has been set.

### GetAuthorizationOauthClientsUrl

`func (o *ApiUrlsResponse) GetAuthorizationOauthClientsUrl() string`

GetAuthorizationOauthClientsUrl returns the AuthorizationOauthClientsUrl field if non-nil, zero value otherwise.

### GetAuthorizationOauthClientsUrlOk

`func (o *ApiUrlsResponse) GetAuthorizationOauthClientsUrlOk() (*string, bool)`

GetAuthorizationOauthClientsUrlOk returns a tuple with the AuthorizationOauthClientsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationOauthClientsUrl

`func (o *ApiUrlsResponse) SetAuthorizationOauthClientsUrl(v string)`

SetAuthorizationOauthClientsUrl sets AuthorizationOauthClientsUrl field to given value.

### HasAuthorizationOauthClientsUrl

`func (o *ApiUrlsResponse) HasAuthorizationOauthClientsUrl() bool`

HasAuthorizationOauthClientsUrl returns a boolean if a field has been set.

### GetAuthorizationOauthClientUrl

`func (o *ApiUrlsResponse) GetAuthorizationOauthClientUrl() string`

GetAuthorizationOauthClientUrl returns the AuthorizationOauthClientUrl field if non-nil, zero value otherwise.

### GetAuthorizationOauthClientUrlOk

`func (o *ApiUrlsResponse) GetAuthorizationOauthClientUrlOk() (*string, bool)`

GetAuthorizationOauthClientUrlOk returns a tuple with the AuthorizationOauthClientUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationOauthClientUrl

`func (o *ApiUrlsResponse) SetAuthorizationOauthClientUrl(v string)`

SetAuthorizationOauthClientUrl sets AuthorizationOauthClientUrl field to given value.

### HasAuthorizationOauthClientUrl

`func (o *ApiUrlsResponse) HasAuthorizationOauthClientUrl() bool`

HasAuthorizationOauthClientUrl returns a boolean if a field has been set.

### GetAuthorizationTokenUrl

`func (o *ApiUrlsResponse) GetAuthorizationTokenUrl() string`

GetAuthorizationTokenUrl returns the AuthorizationTokenUrl field if non-nil, zero value otherwise.

### GetAuthorizationTokenUrlOk

`func (o *ApiUrlsResponse) GetAuthorizationTokenUrlOk() (*string, bool)`

GetAuthorizationTokenUrlOk returns a tuple with the AuthorizationTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationTokenUrl

`func (o *ApiUrlsResponse) SetAuthorizationTokenUrl(v string)`

SetAuthorizationTokenUrl sets AuthorizationTokenUrl field to given value.

### HasAuthorizationTokenUrl

`func (o *ApiUrlsResponse) HasAuthorizationTokenUrl() bool`

HasAuthorizationTokenUrl returns a boolean if a field has been set.

### GetBalanceUrl

`func (o *ApiUrlsResponse) GetBalanceUrl() string`

GetBalanceUrl returns the BalanceUrl field if non-nil, zero value otherwise.

### GetBalanceUrlOk

`func (o *ApiUrlsResponse) GetBalanceUrlOk() (*string, bool)`

GetBalanceUrlOk returns a tuple with the BalanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceUrl

`func (o *ApiUrlsResponse) SetBalanceUrl(v string)`

SetBalanceUrl sets BalanceUrl field to given value.

### HasBalanceUrl

`func (o *ApiUrlsResponse) HasBalanceUrl() bool`

HasBalanceUrl returns a boolean if a field has been set.

### GetGetCallRestrictionsUrl

`func (o *ApiUrlsResponse) GetGetCallRestrictionsUrl() string`

GetGetCallRestrictionsUrl returns the GetCallRestrictionsUrl field if non-nil, zero value otherwise.

### GetGetCallRestrictionsUrlOk

`func (o *ApiUrlsResponse) GetGetCallRestrictionsUrlOk() (*string, bool)`

GetGetCallRestrictionsUrlOk returns a tuple with the GetCallRestrictionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetCallRestrictionsUrl

`func (o *ApiUrlsResponse) SetGetCallRestrictionsUrl(v string)`

SetGetCallRestrictionsUrl sets GetCallRestrictionsUrl field to given value.

### HasGetCallRestrictionsUrl

`func (o *ApiUrlsResponse) HasGetCallRestrictionsUrl() bool`

HasGetCallRestrictionsUrl returns a boolean if a field has been set.

### GetCallRestrictionsUrl

`func (o *ApiUrlsResponse) GetCallRestrictionsUrl() string`

GetCallRestrictionsUrl returns the CallRestrictionsUrl field if non-nil, zero value otherwise.

### GetCallRestrictionsUrlOk

`func (o *ApiUrlsResponse) GetCallRestrictionsUrlOk() (*string, bool)`

GetCallRestrictionsUrlOk returns a tuple with the CallRestrictionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRestrictionsUrl

`func (o *ApiUrlsResponse) SetCallRestrictionsUrl(v string)`

SetCallRestrictionsUrl sets CallRestrictionsUrl field to given value.

### HasCallRestrictionsUrl

`func (o *ApiUrlsResponse) HasCallRestrictionsUrl() bool`

HasCallRestrictionsUrl returns a boolean if a field has been set.

### GetContactsUrl

`func (o *ApiUrlsResponse) GetContactsUrl() string`

GetContactsUrl returns the ContactsUrl field if non-nil, zero value otherwise.

### GetContactsUrlOk

`func (o *ApiUrlsResponse) GetContactsUrlOk() (*string, bool)`

GetContactsUrlOk returns a tuple with the ContactsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactsUrl

`func (o *ApiUrlsResponse) SetContactsUrl(v string)`

SetContactsUrl sets ContactsUrl field to given value.

### HasContactsUrl

`func (o *ApiUrlsResponse) HasContactsUrl() bool`

HasContactsUrl returns a boolean if a field has been set.

### GetCrmBridgeUrl

`func (o *ApiUrlsResponse) GetCrmBridgeUrl() string`

GetCrmBridgeUrl returns the CrmBridgeUrl field if non-nil, zero value otherwise.

### GetCrmBridgeUrlOk

`func (o *ApiUrlsResponse) GetCrmBridgeUrlOk() (*string, bool)`

GetCrmBridgeUrlOk returns a tuple with the CrmBridgeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmBridgeUrl

`func (o *ApiUrlsResponse) SetCrmBridgeUrl(v string)`

SetCrmBridgeUrl sets CrmBridgeUrl field to given value.

### HasCrmBridgeUrl

`func (o *ApiUrlsResponse) HasCrmBridgeUrl() bool`

HasCrmBridgeUrl returns a boolean if a field has been set.

### GetInternalContactsUrl

`func (o *ApiUrlsResponse) GetInternalContactsUrl() string`

GetInternalContactsUrl returns the InternalContactsUrl field if non-nil, zero value otherwise.

### GetInternalContactsUrlOk

`func (o *ApiUrlsResponse) GetInternalContactsUrlOk() (*string, bool)`

GetInternalContactsUrlOk returns a tuple with the InternalContactsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalContactsUrl

`func (o *ApiUrlsResponse) SetInternalContactsUrl(v string)`

SetInternalContactsUrl sets InternalContactsUrl field to given value.

### HasInternalContactsUrl

`func (o *ApiUrlsResponse) HasInternalContactsUrl() bool`

HasInternalContactsUrl returns a boolean if a field has been set.

### GetDeviceCallerIdUrl

`func (o *ApiUrlsResponse) GetDeviceCallerIdUrl() string`

GetDeviceCallerIdUrl returns the DeviceCallerIdUrl field if non-nil, zero value otherwise.

### GetDeviceCallerIdUrlOk

`func (o *ApiUrlsResponse) GetDeviceCallerIdUrlOk() (*string, bool)`

GetDeviceCallerIdUrlOk returns a tuple with the DeviceCallerIdUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCallerIdUrl

`func (o *ApiUrlsResponse) SetDeviceCallerIdUrl(v string)`

SetDeviceCallerIdUrl sets DeviceCallerIdUrl field to given value.

### HasDeviceCallerIdUrl

`func (o *ApiUrlsResponse) HasDeviceCallerIdUrl() bool`

HasDeviceCallerIdUrl returns a boolean if a field has been set.

### GetDeviceUrl

`func (o *ApiUrlsResponse) GetDeviceUrl() string`

GetDeviceUrl returns the DeviceUrl field if non-nil, zero value otherwise.

### GetDeviceUrlOk

`func (o *ApiUrlsResponse) GetDeviceUrlOk() (*string, bool)`

GetDeviceUrlOk returns a tuple with the DeviceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUrl

`func (o *ApiUrlsResponse) SetDeviceUrl(v string)`

SetDeviceUrl sets DeviceUrl field to given value.

### HasDeviceUrl

`func (o *ApiUrlsResponse) HasDeviceUrl() bool`

HasDeviceUrl returns a boolean if a field has been set.

### GetDevicesTariffAnnouncementUrl

`func (o *ApiUrlsResponse) GetDevicesTariffAnnouncementUrl() string`

GetDevicesTariffAnnouncementUrl returns the DevicesTariffAnnouncementUrl field if non-nil, zero value otherwise.

### GetDevicesTariffAnnouncementUrlOk

`func (o *ApiUrlsResponse) GetDevicesTariffAnnouncementUrlOk() (*string, bool)`

GetDevicesTariffAnnouncementUrlOk returns a tuple with the DevicesTariffAnnouncementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesTariffAnnouncementUrl

`func (o *ApiUrlsResponse) SetDevicesTariffAnnouncementUrl(v string)`

SetDevicesTariffAnnouncementUrl sets DevicesTariffAnnouncementUrl field to given value.

### HasDevicesTariffAnnouncementUrl

`func (o *ApiUrlsResponse) HasDevicesTariffAnnouncementUrl() bool`

HasDevicesTariffAnnouncementUrl returns a boolean if a field has been set.

### GetDevicesSingleRowDisplayUrl

`func (o *ApiUrlsResponse) GetDevicesSingleRowDisplayUrl() string`

GetDevicesSingleRowDisplayUrl returns the DevicesSingleRowDisplayUrl field if non-nil, zero value otherwise.

### GetDevicesSingleRowDisplayUrlOk

`func (o *ApiUrlsResponse) GetDevicesSingleRowDisplayUrlOk() (*string, bool)`

GetDevicesSingleRowDisplayUrlOk returns a tuple with the DevicesSingleRowDisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesSingleRowDisplayUrl

`func (o *ApiUrlsResponse) SetDevicesSingleRowDisplayUrl(v string)`

SetDevicesSingleRowDisplayUrl sets DevicesSingleRowDisplayUrl field to given value.

### HasDevicesSingleRowDisplayUrl

`func (o *ApiUrlsResponse) HasDevicesSingleRowDisplayUrl() bool`

HasDevicesSingleRowDisplayUrl returns a boolean if a field has been set.

### GetDeviceSimOrdersUrl

`func (o *ApiUrlsResponse) GetDeviceSimOrdersUrl() string`

GetDeviceSimOrdersUrl returns the DeviceSimOrdersUrl field if non-nil, zero value otherwise.

### GetDeviceSimOrdersUrlOk

`func (o *ApiUrlsResponse) GetDeviceSimOrdersUrlOk() (*string, bool)`

GetDeviceSimOrdersUrlOk returns a tuple with the DeviceSimOrdersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSimOrdersUrl

`func (o *ApiUrlsResponse) SetDeviceSimOrdersUrl(v string)`

SetDeviceSimOrdersUrl sets DeviceSimOrdersUrl field to given value.

### HasDeviceSimOrdersUrl

`func (o *ApiUrlsResponse) HasDeviceSimOrdersUrl() bool`

HasDeviceSimOrdersUrl returns a boolean if a field has been set.

### GetFaxlinesUrl

`func (o *ApiUrlsResponse) GetFaxlinesUrl() string`

GetFaxlinesUrl returns the FaxlinesUrl field if non-nil, zero value otherwise.

### GetFaxlinesUrlOk

`func (o *ApiUrlsResponse) GetFaxlinesUrlOk() (*string, bool)`

GetFaxlinesUrlOk returns a tuple with the FaxlinesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxlinesUrl

`func (o *ApiUrlsResponse) SetFaxlinesUrl(v string)`

SetFaxlinesUrl sets FaxlinesUrl field to given value.

### HasFaxlinesUrl

`func (o *ApiUrlsResponse) HasFaxlinesUrl() bool`

HasFaxlinesUrl returns a boolean if a field has been set.

### GetFaxlineUrl

`func (o *ApiUrlsResponse) GetFaxlineUrl() string`

GetFaxlineUrl returns the FaxlineUrl field if non-nil, zero value otherwise.

### GetFaxlineUrlOk

`func (o *ApiUrlsResponse) GetFaxlineUrlOk() (*string, bool)`

GetFaxlineUrlOk returns a tuple with the FaxlineUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxlineUrl

`func (o *ApiUrlsResponse) SetFaxlineUrl(v string)`

SetFaxlineUrl sets FaxlineUrl field to given value.

### HasFaxlineUrl

`func (o *ApiUrlsResponse) HasFaxlineUrl() bool`

HasFaxlineUrl returns a boolean if a field has been set.

### GetFaxlineNumbersUrl

`func (o *ApiUrlsResponse) GetFaxlineNumbersUrl() string`

GetFaxlineNumbersUrl returns the FaxlineNumbersUrl field if non-nil, zero value otherwise.

### GetFaxlineNumbersUrlOk

`func (o *ApiUrlsResponse) GetFaxlineNumbersUrlOk() (*string, bool)`

GetFaxlineNumbersUrlOk returns a tuple with the FaxlineNumbersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxlineNumbersUrl

`func (o *ApiUrlsResponse) SetFaxlineNumbersUrl(v string)`

SetFaxlineNumbersUrl sets FaxlineNumbersUrl field to given value.

### HasFaxlineNumbersUrl

`func (o *ApiUrlsResponse) HasFaxlineNumbersUrl() bool`

HasFaxlineNumbersUrl returns a boolean if a field has been set.

### GetGroupsUrl

`func (o *ApiUrlsResponse) GetGroupsUrl() string`

GetGroupsUrl returns the GroupsUrl field if non-nil, zero value otherwise.

### GetGroupsUrlOk

`func (o *ApiUrlsResponse) GetGroupsUrlOk() (*string, bool)`

GetGroupsUrlOk returns a tuple with the GroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsUrl

`func (o *ApiUrlsResponse) SetGroupsUrl(v string)`

SetGroupsUrl sets GroupsUrl field to given value.

### HasGroupsUrl

`func (o *ApiUrlsResponse) HasGroupsUrl() bool`

HasGroupsUrl returns a boolean if a field has been set.

### GetGroupDetailUrl

`func (o *ApiUrlsResponse) GetGroupDetailUrl() string`

GetGroupDetailUrl returns the GroupDetailUrl field if non-nil, zero value otherwise.

### GetGroupDetailUrlOk

`func (o *ApiUrlsResponse) GetGroupDetailUrlOk() (*string, bool)`

GetGroupDetailUrlOk returns a tuple with the GroupDetailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDetailUrl

`func (o *ApiUrlsResponse) SetGroupDetailUrl(v string)`

SetGroupDetailUrl sets GroupDetailUrl field to given value.

### HasGroupDetailUrl

`func (o *ApiUrlsResponse) HasGroupDetailUrl() bool`

HasGroupDetailUrl returns a boolean if a field has been set.

### GetGroupUserUrl

`func (o *ApiUrlsResponse) GetGroupUserUrl() string`

GetGroupUserUrl returns the GroupUserUrl field if non-nil, zero value otherwise.

### GetGroupUserUrlOk

`func (o *ApiUrlsResponse) GetGroupUserUrlOk() (*string, bool)`

GetGroupUserUrlOk returns a tuple with the GroupUserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUserUrl

`func (o *ApiUrlsResponse) SetGroupUserUrl(v string)`

SetGroupUserUrl sets GroupUserUrl field to given value.

### HasGroupUserUrl

`func (o *ApiUrlsResponse) HasGroupUserUrl() bool`

HasGroupUserUrl returns a boolean if a field has been set.

### GetGroupDevicesUrl

`func (o *ApiUrlsResponse) GetGroupDevicesUrl() string`

GetGroupDevicesUrl returns the GroupDevicesUrl field if non-nil, zero value otherwise.

### GetGroupDevicesUrlOk

`func (o *ApiUrlsResponse) GetGroupDevicesUrlOk() (*string, bool)`

GetGroupDevicesUrlOk returns a tuple with the GroupDevicesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDevicesUrl

`func (o *ApiUrlsResponse) SetGroupDevicesUrl(v string)`

SetGroupDevicesUrl sets GroupDevicesUrl field to given value.

### HasGroupDevicesUrl

`func (o *ApiUrlsResponse) HasGroupDevicesUrl() bool`

HasGroupDevicesUrl returns a boolean if a field has been set.

### GetGroupNumbersUrl

`func (o *ApiUrlsResponse) GetGroupNumbersUrl() string`

GetGroupNumbersUrl returns the GroupNumbersUrl field if non-nil, zero value otherwise.

### GetGroupNumbersUrlOk

`func (o *ApiUrlsResponse) GetGroupNumbersUrlOk() (*string, bool)`

GetGroupNumbersUrlOk returns a tuple with the GroupNumbersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNumbersUrl

`func (o *ApiUrlsResponse) SetGroupNumbersUrl(v string)`

SetGroupNumbersUrl sets GroupNumbersUrl field to given value.

### HasGroupNumbersUrl

`func (o *ApiUrlsResponse) HasGroupNumbersUrl() bool`

HasGroupNumbersUrl returns a boolean if a field has been set.

### GetHistoryEntryUrl

`func (o *ApiUrlsResponse) GetHistoryEntryUrl() string`

GetHistoryEntryUrl returns the HistoryEntryUrl field if non-nil, zero value otherwise.

### GetHistoryEntryUrlOk

`func (o *ApiUrlsResponse) GetHistoryEntryUrlOk() (*string, bool)`

GetHistoryEntryUrlOk returns a tuple with the HistoryEntryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryEntryUrl

`func (o *ApiUrlsResponse) SetHistoryEntryUrl(v string)`

SetHistoryEntryUrl sets HistoryEntryUrl field to given value.

### HasHistoryEntryUrl

`func (o *ApiUrlsResponse) HasHistoryEntryUrl() bool`

HasHistoryEntryUrl returns a boolean if a field has been set.

### GetHistoryUrl

`func (o *ApiUrlsResponse) GetHistoryUrl() string`

GetHistoryUrl returns the HistoryUrl field if non-nil, zero value otherwise.

### GetHistoryUrlOk

`func (o *ApiUrlsResponse) GetHistoryUrlOk() (*string, bool)`

GetHistoryUrlOk returns a tuple with the HistoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryUrl

`func (o *ApiUrlsResponse) SetHistoryUrl(v string)`

SetHistoryUrl sets HistoryUrl field to given value.

### HasHistoryUrl

`func (o *ApiUrlsResponse) HasHistoryUrl() bool`

HasHistoryUrl returns a boolean if a field has been set.

### GetNotificationsUrl

`func (o *ApiUrlsResponse) GetNotificationsUrl() string`

GetNotificationsUrl returns the NotificationsUrl field if non-nil, zero value otherwise.

### GetNotificationsUrlOk

`func (o *ApiUrlsResponse) GetNotificationsUrlOk() (*string, bool)`

GetNotificationsUrlOk returns a tuple with the NotificationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsUrl

`func (o *ApiUrlsResponse) SetNotificationsUrl(v string)`

SetNotificationsUrl sets NotificationsUrl field to given value.

### HasNotificationsUrl

`func (o *ApiUrlsResponse) HasNotificationsUrl() bool`

HasNotificationsUrl returns a boolean if a field has been set.

### GetNumberUrl

`func (o *ApiUrlsResponse) GetNumberUrl() string`

GetNumberUrl returns the NumberUrl field if non-nil, zero value otherwise.

### GetNumberUrlOk

`func (o *ApiUrlsResponse) GetNumberUrlOk() (*string, bool)`

GetNumberUrlOk returns a tuple with the NumberUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberUrl

`func (o *ApiUrlsResponse) SetNumberUrl(v string)`

SetNumberUrl sets NumberUrl field to given value.

### HasNumberUrl

`func (o *ApiUrlsResponse) HasNumberUrl() bool`

HasNumberUrl returns a boolean if a field has been set.

### GetNumbersUrl

`func (o *ApiUrlsResponse) GetNumbersUrl() string`

GetNumbersUrl returns the NumbersUrl field if non-nil, zero value otherwise.

### GetNumbersUrlOk

`func (o *ApiUrlsResponse) GetNumbersUrlOk() (*string, bool)`

GetNumbersUrlOk returns a tuple with the NumbersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbersUrl

`func (o *ApiUrlsResponse) SetNumbersUrl(v string)`

SetNumbersUrl sets NumbersUrl field to given value.

### HasNumbersUrl

`func (o *ApiUrlsResponse) HasNumbersUrl() bool`

HasNumbersUrl returns a boolean if a field has been set.

### GetWebuserNumbersUrl

`func (o *ApiUrlsResponse) GetWebuserNumbersUrl() string`

GetWebuserNumbersUrl returns the WebuserNumbersUrl field if non-nil, zero value otherwise.

### GetWebuserNumbersUrlOk

`func (o *ApiUrlsResponse) GetWebuserNumbersUrlOk() (*string, bool)`

GetWebuserNumbersUrlOk returns a tuple with the WebuserNumbersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuserNumbersUrl

`func (o *ApiUrlsResponse) SetWebuserNumbersUrl(v string)`

SetWebuserNumbersUrl sets WebuserNumbersUrl field to given value.

### HasWebuserNumbersUrl

`func (o *ApiUrlsResponse) HasWebuserNumbersUrl() bool`

HasWebuserNumbersUrl returns a boolean if a field has been set.

### GetAddDirectDialUrl

`func (o *ApiUrlsResponse) GetAddDirectDialUrl() string`

GetAddDirectDialUrl returns the AddDirectDialUrl field if non-nil, zero value otherwise.

### GetAddDirectDialUrlOk

`func (o *ApiUrlsResponse) GetAddDirectDialUrlOk() (*string, bool)`

GetAddDirectDialUrlOk returns a tuple with the AddDirectDialUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDirectDialUrl

`func (o *ApiUrlsResponse) SetAddDirectDialUrl(v string)`

SetAddDirectDialUrl sets AddDirectDialUrl field to given value.

### HasAddDirectDialUrl

`func (o *ApiUrlsResponse) HasAddDirectDialUrl() bool`

HasAddDirectDialUrl returns a boolean if a field has been set.

### GetChangeDirectDialUrl

`func (o *ApiUrlsResponse) GetChangeDirectDialUrl() string`

GetChangeDirectDialUrl returns the ChangeDirectDialUrl field if non-nil, zero value otherwise.

### GetChangeDirectDialUrlOk

`func (o *ApiUrlsResponse) GetChangeDirectDialUrlOk() (*string, bool)`

GetChangeDirectDialUrlOk returns a tuple with the ChangeDirectDialUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDirectDialUrl

`func (o *ApiUrlsResponse) SetChangeDirectDialUrl(v string)`

SetChangeDirectDialUrl sets ChangeDirectDialUrl field to given value.

### HasChangeDirectDialUrl

`func (o *ApiUrlsResponse) HasChangeDirectDialUrl() bool`

HasChangeDirectDialUrl returns a boolean if a field has been set.

### GetDeleteDirectDialUrl

`func (o *ApiUrlsResponse) GetDeleteDirectDialUrl() string`

GetDeleteDirectDialUrl returns the DeleteDirectDialUrl field if non-nil, zero value otherwise.

### GetDeleteDirectDialUrlOk

`func (o *ApiUrlsResponse) GetDeleteDirectDialUrlOk() (*string, bool)`

GetDeleteDirectDialUrlOk returns a tuple with the DeleteDirectDialUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteDirectDialUrl

`func (o *ApiUrlsResponse) SetDeleteDirectDialUrl(v string)`

SetDeleteDirectDialUrl sets DeleteDirectDialUrl field to given value.

### HasDeleteDirectDialUrl

`func (o *ApiUrlsResponse) HasDeleteDirectDialUrl() bool`

HasDeleteDirectDialUrl returns a boolean if a field has been set.

### GetPhonelineUrl

`func (o *ApiUrlsResponse) GetPhonelineUrl() string`

GetPhonelineUrl returns the PhonelineUrl field if non-nil, zero value otherwise.

### GetPhonelineUrlOk

`func (o *ApiUrlsResponse) GetPhonelineUrlOk() (*string, bool)`

GetPhonelineUrlOk returns a tuple with the PhonelineUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineUrl

`func (o *ApiUrlsResponse) SetPhonelineUrl(v string)`

SetPhonelineUrl sets PhonelineUrl field to given value.

### HasPhonelineUrl

`func (o *ApiUrlsResponse) HasPhonelineUrl() bool`

HasPhonelineUrl returns a boolean if a field has been set.

### GetPhonelineBlockAnonymousUrl

`func (o *ApiUrlsResponse) GetPhonelineBlockAnonymousUrl() string`

GetPhonelineBlockAnonymousUrl returns the PhonelineBlockAnonymousUrl field if non-nil, zero value otherwise.

### GetPhonelineBlockAnonymousUrlOk

`func (o *ApiUrlsResponse) GetPhonelineBlockAnonymousUrlOk() (*string, bool)`

GetPhonelineBlockAnonymousUrlOk returns a tuple with the PhonelineBlockAnonymousUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineBlockAnonymousUrl

`func (o *ApiUrlsResponse) SetPhonelineBlockAnonymousUrl(v string)`

SetPhonelineBlockAnonymousUrl sets PhonelineBlockAnonymousUrl field to given value.

### HasPhonelineBlockAnonymousUrl

`func (o *ApiUrlsResponse) HasPhonelineBlockAnonymousUrl() bool`

HasPhonelineBlockAnonymousUrl returns a boolean if a field has been set.

### GetPhonelineDevicesUrl

`func (o *ApiUrlsResponse) GetPhonelineDevicesUrl() string`

GetPhonelineDevicesUrl returns the PhonelineDevicesUrl field if non-nil, zero value otherwise.

### GetPhonelineDevicesUrlOk

`func (o *ApiUrlsResponse) GetPhonelineDevicesUrlOk() (*string, bool)`

GetPhonelineDevicesUrlOk returns a tuple with the PhonelineDevicesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineDevicesUrl

`func (o *ApiUrlsResponse) SetPhonelineDevicesUrl(v string)`

SetPhonelineDevicesUrl sets PhonelineDevicesUrl field to given value.

### HasPhonelineDevicesUrl

`func (o *ApiUrlsResponse) HasPhonelineDevicesUrl() bool`

HasPhonelineDevicesUrl returns a boolean if a field has been set.

### GetPhonelineForwardingsUrl

`func (o *ApiUrlsResponse) GetPhonelineForwardingsUrl() string`

GetPhonelineForwardingsUrl returns the PhonelineForwardingsUrl field if non-nil, zero value otherwise.

### GetPhonelineForwardingsUrlOk

`func (o *ApiUrlsResponse) GetPhonelineForwardingsUrlOk() (*string, bool)`

GetPhonelineForwardingsUrlOk returns a tuple with the PhonelineForwardingsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineForwardingsUrl

`func (o *ApiUrlsResponse) SetPhonelineForwardingsUrl(v string)`

SetPhonelineForwardingsUrl sets PhonelineForwardingsUrl field to given value.

### HasPhonelineForwardingsUrl

`func (o *ApiUrlsResponse) HasPhonelineForwardingsUrl() bool`

HasPhonelineForwardingsUrl returns a boolean if a field has been set.

### GetPhonelineNumbersUrl

`func (o *ApiUrlsResponse) GetPhonelineNumbersUrl() string`

GetPhonelineNumbersUrl returns the PhonelineNumbersUrl field if non-nil, zero value otherwise.

### GetPhonelineNumbersUrlOk

`func (o *ApiUrlsResponse) GetPhonelineNumbersUrlOk() (*string, bool)`

GetPhonelineNumbersUrlOk returns a tuple with the PhonelineNumbersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineNumbersUrl

`func (o *ApiUrlsResponse) SetPhonelineNumbersUrl(v string)`

SetPhonelineNumbersUrl sets PhonelineNumbersUrl field to given value.

### HasPhonelineNumbersUrl

`func (o *ApiUrlsResponse) HasPhonelineNumbersUrl() bool`

HasPhonelineNumbersUrl returns a boolean if a field has been set.

### GetPhonelinesParallelForwardingUrl

`func (o *ApiUrlsResponse) GetPhonelinesParallelForwardingUrl() string`

GetPhonelinesParallelForwardingUrl returns the PhonelinesParallelForwardingUrl field if non-nil, zero value otherwise.

### GetPhonelinesParallelForwardingUrlOk

`func (o *ApiUrlsResponse) GetPhonelinesParallelForwardingUrlOk() (*string, bool)`

GetPhonelinesParallelForwardingUrlOk returns a tuple with the PhonelinesParallelForwardingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelinesParallelForwardingUrl

`func (o *ApiUrlsResponse) SetPhonelinesParallelForwardingUrl(v string)`

SetPhonelinesParallelForwardingUrl sets PhonelinesParallelForwardingUrl field to given value.

### HasPhonelinesParallelForwardingUrl

`func (o *ApiUrlsResponse) HasPhonelinesParallelForwardingUrl() bool`

HasPhonelinesParallelForwardingUrl returns a boolean if a field has been set.

### GetPhonelinesParallelForwardingsUrl

`func (o *ApiUrlsResponse) GetPhonelinesParallelForwardingsUrl() string`

GetPhonelinesParallelForwardingsUrl returns the PhonelinesParallelForwardingsUrl field if non-nil, zero value otherwise.

### GetPhonelinesParallelForwardingsUrlOk

`func (o *ApiUrlsResponse) GetPhonelinesParallelForwardingsUrlOk() (*string, bool)`

GetPhonelinesParallelForwardingsUrlOk returns a tuple with the PhonelinesParallelForwardingsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelinesParallelForwardingsUrl

`func (o *ApiUrlsResponse) SetPhonelinesParallelForwardingsUrl(v string)`

SetPhonelinesParallelForwardingsUrl sets PhonelinesParallelForwardingsUrl field to given value.

### HasPhonelinesParallelForwardingsUrl

`func (o *ApiUrlsResponse) HasPhonelinesParallelForwardingsUrl() bool`

HasPhonelinesParallelForwardingsUrl returns a boolean if a field has been set.

### GetPhonelineSipgateIoUrl

`func (o *ApiUrlsResponse) GetPhonelineSipgateIoUrl() string`

GetPhonelineSipgateIoUrl returns the PhonelineSipgateIoUrl field if non-nil, zero value otherwise.

### GetPhonelineSipgateIoUrlOk

`func (o *ApiUrlsResponse) GetPhonelineSipgateIoUrlOk() (*string, bool)`

GetPhonelineSipgateIoUrlOk returns a tuple with the PhonelineSipgateIoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineSipgateIoUrl

`func (o *ApiUrlsResponse) SetPhonelineSipgateIoUrl(v string)`

SetPhonelineSipgateIoUrl sets PhonelineSipgateIoUrl field to given value.

### HasPhonelineSipgateIoUrl

`func (o *ApiUrlsResponse) HasPhonelineSipgateIoUrl() bool`

HasPhonelineSipgateIoUrl returns a boolean if a field has been set.

### GetPhonelineSipgateIoLogUrl

`func (o *ApiUrlsResponse) GetPhonelineSipgateIoLogUrl() string`

GetPhonelineSipgateIoLogUrl returns the PhonelineSipgateIoLogUrl field if non-nil, zero value otherwise.

### GetPhonelineSipgateIoLogUrlOk

`func (o *ApiUrlsResponse) GetPhonelineSipgateIoLogUrlOk() (*string, bool)`

GetPhonelineSipgateIoLogUrlOk returns a tuple with the PhonelineSipgateIoLogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineSipgateIoLogUrl

`func (o *ApiUrlsResponse) SetPhonelineSipgateIoLogUrl(v string)`

SetPhonelineSipgateIoLogUrl sets PhonelineSipgateIoLogUrl field to given value.

### HasPhonelineSipgateIoLogUrl

`func (o *ApiUrlsResponse) HasPhonelineSipgateIoLogUrl() bool`

HasPhonelineSipgateIoLogUrl returns a boolean if a field has been set.

### GetPhonelineVoicemailGreetingUrl

`func (o *ApiUrlsResponse) GetPhonelineVoicemailGreetingUrl() string`

GetPhonelineVoicemailGreetingUrl returns the PhonelineVoicemailGreetingUrl field if non-nil, zero value otherwise.

### GetPhonelineVoicemailGreetingUrlOk

`func (o *ApiUrlsResponse) GetPhonelineVoicemailGreetingUrlOk() (*string, bool)`

GetPhonelineVoicemailGreetingUrlOk returns a tuple with the PhonelineVoicemailGreetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineVoicemailGreetingUrl

`func (o *ApiUrlsResponse) SetPhonelineVoicemailGreetingUrl(v string)`

SetPhonelineVoicemailGreetingUrl sets PhonelineVoicemailGreetingUrl field to given value.

### HasPhonelineVoicemailGreetingUrl

`func (o *ApiUrlsResponse) HasPhonelineVoicemailGreetingUrl() bool`

HasPhonelineVoicemailGreetingUrl returns a boolean if a field has been set.

### GetPhonelineVoicemailGreetingsUrl

`func (o *ApiUrlsResponse) GetPhonelineVoicemailGreetingsUrl() string`

GetPhonelineVoicemailGreetingsUrl returns the PhonelineVoicemailGreetingsUrl field if non-nil, zero value otherwise.

### GetPhonelineVoicemailGreetingsUrlOk

`func (o *ApiUrlsResponse) GetPhonelineVoicemailGreetingsUrlOk() (*string, bool)`

GetPhonelineVoicemailGreetingsUrlOk returns a tuple with the PhonelineVoicemailGreetingsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineVoicemailGreetingsUrl

`func (o *ApiUrlsResponse) SetPhonelineVoicemailGreetingsUrl(v string)`

SetPhonelineVoicemailGreetingsUrl sets PhonelineVoicemailGreetingsUrl field to given value.

### HasPhonelineVoicemailGreetingsUrl

`func (o *ApiUrlsResponse) HasPhonelineVoicemailGreetingsUrl() bool`

HasPhonelineVoicemailGreetingsUrl returns a boolean if a field has been set.

### GetPhonelineVoicemailUrl

`func (o *ApiUrlsResponse) GetPhonelineVoicemailUrl() string`

GetPhonelineVoicemailUrl returns the PhonelineVoicemailUrl field if non-nil, zero value otherwise.

### GetPhonelineVoicemailUrlOk

`func (o *ApiUrlsResponse) GetPhonelineVoicemailUrlOk() (*string, bool)`

GetPhonelineVoicemailUrlOk returns a tuple with the PhonelineVoicemailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineVoicemailUrl

`func (o *ApiUrlsResponse) SetPhonelineVoicemailUrl(v string)`

SetPhonelineVoicemailUrl sets PhonelineVoicemailUrl field to given value.

### HasPhonelineVoicemailUrl

`func (o *ApiUrlsResponse) HasPhonelineVoicemailUrl() bool`

HasPhonelineVoicemailUrl returns a boolean if a field has been set.

### GetPhonelineVoicemailsUrl

`func (o *ApiUrlsResponse) GetPhonelineVoicemailsUrl() string`

GetPhonelineVoicemailsUrl returns the PhonelineVoicemailsUrl field if non-nil, zero value otherwise.

### GetPhonelineVoicemailsUrlOk

`func (o *ApiUrlsResponse) GetPhonelineVoicemailsUrlOk() (*string, bool)`

GetPhonelineVoicemailsUrlOk returns a tuple with the PhonelineVoicemailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineVoicemailsUrl

`func (o *ApiUrlsResponse) SetPhonelineVoicemailsUrl(v string)`

SetPhonelineVoicemailsUrl sets PhonelineVoicemailsUrl field to given value.

### HasPhonelineVoicemailsUrl

`func (o *ApiUrlsResponse) HasPhonelineVoicemailsUrl() bool`

HasPhonelineVoicemailsUrl returns a boolean if a field has been set.

### GetPhonelinesUrl

`func (o *ApiUrlsResponse) GetPhonelinesUrl() string`

GetPhonelinesUrl returns the PhonelinesUrl field if non-nil, zero value otherwise.

### GetPhonelinesUrlOk

`func (o *ApiUrlsResponse) GetPhonelinesUrlOk() (*string, bool)`

GetPhonelinesUrlOk returns a tuple with the PhonelinesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelinesUrl

`func (o *ApiUrlsResponse) SetPhonelinesUrl(v string)`

SetPhonelinesUrl sets PhonelinesUrl field to given value.

### HasPhonelinesUrl

`func (o *ApiUrlsResponse) HasPhonelinesUrl() bool`

HasPhonelinesUrl returns a boolean if a field has been set.

### GetPhonelineDetailUrl

`func (o *ApiUrlsResponse) GetPhonelineDetailUrl() string`

GetPhonelineDetailUrl returns the PhonelineDetailUrl field if non-nil, zero value otherwise.

### GetPhonelineDetailUrlOk

`func (o *ApiUrlsResponse) GetPhonelineDetailUrlOk() (*string, bool)`

GetPhonelineDetailUrlOk returns a tuple with the PhonelineDetailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonelineDetailUrl

`func (o *ApiUrlsResponse) SetPhonelineDetailUrl(v string)`

SetPhonelineDetailUrl sets PhonelineDetailUrl field to given value.

### HasPhonelineDetailUrl

`func (o *ApiUrlsResponse) HasPhonelineDetailUrl() bool`

HasPhonelineDetailUrl returns a boolean if a field has been set.

### GetPortingUrl

`func (o *ApiUrlsResponse) GetPortingUrl() string`

GetPortingUrl returns the PortingUrl field if non-nil, zero value otherwise.

### GetPortingUrlOk

`func (o *ApiUrlsResponse) GetPortingUrlOk() (*string, bool)`

GetPortingUrlOk returns a tuple with the PortingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingUrl

`func (o *ApiUrlsResponse) SetPortingUrl(v string)`

SetPortingUrl sets PortingUrl field to given value.

### HasPortingUrl

`func (o *ApiUrlsResponse) HasPortingUrl() bool`

HasPortingUrl returns a boolean if a field has been set.

### GetPortingsUrl

`func (o *ApiUrlsResponse) GetPortingsUrl() string`

GetPortingsUrl returns the PortingsUrl field if non-nil, zero value otherwise.

### GetPortingsUrlOk

`func (o *ApiUrlsResponse) GetPortingsUrlOk() (*string, bool)`

GetPortingsUrlOk returns a tuple with the PortingsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingsUrl

`func (o *ApiUrlsResponse) SetPortingsUrl(v string)`

SetPortingsUrl sets PortingsUrl field to given value.

### HasPortingsUrl

`func (o *ApiUrlsResponse) HasPortingsUrl() bool`

HasPortingsUrl returns a boolean if a field has been set.

### GetRestrictionsUrl

`func (o *ApiUrlsResponse) GetRestrictionsUrl() string`

GetRestrictionsUrl returns the RestrictionsUrl field if non-nil, zero value otherwise.

### GetRestrictionsUrlOk

`func (o *ApiUrlsResponse) GetRestrictionsUrlOk() (*string, bool)`

GetRestrictionsUrlOk returns a tuple with the RestrictionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionsUrl

`func (o *ApiUrlsResponse) SetRestrictionsUrl(v string)`

SetRestrictionsUrl sets RestrictionsUrl field to given value.

### HasRestrictionsUrl

`func (o *ApiUrlsResponse) HasRestrictionsUrl() bool`

HasRestrictionsUrl returns a boolean if a field has been set.

### GetSessionsCallUrl

`func (o *ApiUrlsResponse) GetSessionsCallUrl() string`

GetSessionsCallUrl returns the SessionsCallUrl field if non-nil, zero value otherwise.

### GetSessionsCallUrlOk

`func (o *ApiUrlsResponse) GetSessionsCallUrlOk() (*string, bool)`

GetSessionsCallUrlOk returns a tuple with the SessionsCallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsCallUrl

`func (o *ApiUrlsResponse) SetSessionsCallUrl(v string)`

SetSessionsCallUrl sets SessionsCallUrl field to given value.

### HasSessionsCallUrl

`func (o *ApiUrlsResponse) HasSessionsCallUrl() bool`

HasSessionsCallUrl returns a boolean if a field has been set.

### GetSettingsSipgateioUrl

`func (o *ApiUrlsResponse) GetSettingsSipgateioUrl() string`

GetSettingsSipgateioUrl returns the SettingsSipgateioUrl field if non-nil, zero value otherwise.

### GetSettingsSipgateioUrlOk

`func (o *ApiUrlsResponse) GetSettingsSipgateioUrlOk() (*string, bool)`

GetSettingsSipgateioUrlOk returns a tuple with the SettingsSipgateioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsSipgateioUrl

`func (o *ApiUrlsResponse) SetSettingsSipgateioUrl(v string)`

SetSettingsSipgateioUrl sets SettingsSipgateioUrl field to given value.

### HasSettingsSipgateioUrl

`func (o *ApiUrlsResponse) HasSettingsSipgateioUrl() bool`

HasSettingsSipgateioUrl returns a boolean if a field has been set.

### GetSmsUrl

`func (o *ApiUrlsResponse) GetSmsUrl() string`

GetSmsUrl returns the SmsUrl field if non-nil, zero value otherwise.

### GetSmsUrlOk

`func (o *ApiUrlsResponse) GetSmsUrlOk() (*string, bool)`

GetSmsUrlOk returns a tuple with the SmsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsUrl

`func (o *ApiUrlsResponse) SetSmsUrl(v string)`

SetSmsUrl sets SmsUrl field to given value.

### HasSmsUrl

`func (o *ApiUrlsResponse) HasSmsUrl() bool`

HasSmsUrl returns a boolean if a field has been set.

### GetSmsCallerIdsUrl

`func (o *ApiUrlsResponse) GetSmsCallerIdsUrl() string`

GetSmsCallerIdsUrl returns the SmsCallerIdsUrl field if non-nil, zero value otherwise.

### GetSmsCallerIdsUrlOk

`func (o *ApiUrlsResponse) GetSmsCallerIdsUrlOk() (*string, bool)`

GetSmsCallerIdsUrlOk returns a tuple with the SmsCallerIdsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCallerIdsUrl

`func (o *ApiUrlsResponse) SetSmsCallerIdsUrl(v string)`

SetSmsCallerIdsUrl sets SmsCallerIdsUrl field to given value.

### HasSmsCallerIdsUrl

`func (o *ApiUrlsResponse) HasSmsCallerIdsUrl() bool`

HasSmsCallerIdsUrl returns a boolean if a field has been set.

### GetTranslationsUrl

`func (o *ApiUrlsResponse) GetTranslationsUrl() string`

GetTranslationsUrl returns the TranslationsUrl field if non-nil, zero value otherwise.

### GetTranslationsUrlOk

`func (o *ApiUrlsResponse) GetTranslationsUrlOk() (*string, bool)`

GetTranslationsUrlOk returns a tuple with the TranslationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationsUrl

`func (o *ApiUrlsResponse) SetTranslationsUrl(v string)`

SetTranslationsUrl sets TranslationsUrl field to given value.

### HasTranslationsUrl

`func (o *ApiUrlsResponse) HasTranslationsUrl() bool`

HasTranslationsUrl returns a boolean if a field has been set.

### GetUserInfoUrl

`func (o *ApiUrlsResponse) GetUserInfoUrl() string`

GetUserInfoUrl returns the UserInfoUrl field if non-nil, zero value otherwise.

### GetUserInfoUrlOk

`func (o *ApiUrlsResponse) GetUserInfoUrlOk() (*string, bool)`

GetUserInfoUrlOk returns a tuple with the UserInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoUrl

`func (o *ApiUrlsResponse) SetUserInfoUrl(v string)`

SetUserInfoUrl sets UserInfoUrl field to given value.

### HasUserInfoUrl

`func (o *ApiUrlsResponse) HasUserInfoUrl() bool`

HasUserInfoUrl returns a boolean if a field has been set.

### GetUserUrl

`func (o *ApiUrlsResponse) GetUserUrl() string`

GetUserUrl returns the UserUrl field if non-nil, zero value otherwise.

### GetUserUrlOk

`func (o *ApiUrlsResponse) GetUserUrlOk() (*string, bool)`

GetUserUrlOk returns a tuple with the UserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUrl

`func (o *ApiUrlsResponse) SetUserUrl(v string)`

SetUserUrl sets UserUrl field to given value.

### HasUserUrl

`func (o *ApiUrlsResponse) HasUserUrl() bool`

HasUserUrl returns a boolean if a field has been set.

### GetUserBusyOnBusyUrl

`func (o *ApiUrlsResponse) GetUserBusyOnBusyUrl() string`

GetUserBusyOnBusyUrl returns the UserBusyOnBusyUrl field if non-nil, zero value otherwise.

### GetUserBusyOnBusyUrlOk

`func (o *ApiUrlsResponse) GetUserBusyOnBusyUrlOk() (*string, bool)`

GetUserBusyOnBusyUrlOk returns a tuple with the UserBusyOnBusyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBusyOnBusyUrl

`func (o *ApiUrlsResponse) SetUserBusyOnBusyUrl(v string)`

SetUserBusyOnBusyUrl sets UserBusyOnBusyUrl field to given value.

### HasUserBusyOnBusyUrl

`func (o *ApiUrlsResponse) HasUserBusyOnBusyUrl() bool`

HasUserBusyOnBusyUrl returns a boolean if a field has been set.

### GetUserDefaultDeviceUrl

`func (o *ApiUrlsResponse) GetUserDefaultDeviceUrl() string`

GetUserDefaultDeviceUrl returns the UserDefaultDeviceUrl field if non-nil, zero value otherwise.

### GetUserDefaultDeviceUrlOk

`func (o *ApiUrlsResponse) GetUserDefaultDeviceUrlOk() (*string, bool)`

GetUserDefaultDeviceUrlOk returns a tuple with the UserDefaultDeviceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefaultDeviceUrl

`func (o *ApiUrlsResponse) SetUserDefaultDeviceUrl(v string)`

SetUserDefaultDeviceUrl sets UserDefaultDeviceUrl field to given value.

### HasUserDefaultDeviceUrl

`func (o *ApiUrlsResponse) HasUserDefaultDeviceUrl() bool`

HasUserDefaultDeviceUrl returns a boolean if a field has been set.

### GetUserDevicesUrl

`func (o *ApiUrlsResponse) GetUserDevicesUrl() string`

GetUserDevicesUrl returns the UserDevicesUrl field if non-nil, zero value otherwise.

### GetUserDevicesUrlOk

`func (o *ApiUrlsResponse) GetUserDevicesUrlOk() (*string, bool)`

GetUserDevicesUrlOk returns a tuple with the UserDevicesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDevicesUrl

`func (o *ApiUrlsResponse) SetUserDevicesUrl(v string)`

SetUserDevicesUrl sets UserDevicesUrl field to given value.

### HasUserDevicesUrl

`func (o *ApiUrlsResponse) HasUserDevicesUrl() bool`

HasUserDevicesUrl returns a boolean if a field has been set.

### GetUserDeviceSimUrl

`func (o *ApiUrlsResponse) GetUserDeviceSimUrl() string`

GetUserDeviceSimUrl returns the UserDeviceSimUrl field if non-nil, zero value otherwise.

### GetUserDeviceSimUrlOk

`func (o *ApiUrlsResponse) GetUserDeviceSimUrlOk() (*string, bool)`

GetUserDeviceSimUrlOk returns a tuple with the UserDeviceSimUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeviceSimUrl

`func (o *ApiUrlsResponse) SetUserDeviceSimUrl(v string)`

SetUserDeviceSimUrl sets UserDeviceSimUrl field to given value.

### HasUserDeviceSimUrl

`func (o *ApiUrlsResponse) HasUserDeviceSimUrl() bool`

HasUserDeviceSimUrl returns a boolean if a field has been set.

### GetUserExternalDevicesUrl

`func (o *ApiUrlsResponse) GetUserExternalDevicesUrl() string`

GetUserExternalDevicesUrl returns the UserExternalDevicesUrl field if non-nil, zero value otherwise.

### GetUserExternalDevicesUrlOk

`func (o *ApiUrlsResponse) GetUserExternalDevicesUrlOk() (*string, bool)`

GetUserExternalDevicesUrlOk returns a tuple with the UserExternalDevicesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExternalDevicesUrl

`func (o *ApiUrlsResponse) SetUserExternalDevicesUrl(v string)`

SetUserExternalDevicesUrl sets UserExternalDevicesUrl field to given value.

### HasUserExternalDevicesUrl

`func (o *ApiUrlsResponse) HasUserExternalDevicesUrl() bool`

HasUserExternalDevicesUrl returns a boolean if a field has been set.

### GetUserMobileDevicesUrl

`func (o *ApiUrlsResponse) GetUserMobileDevicesUrl() string`

GetUserMobileDevicesUrl returns the UserMobileDevicesUrl field if non-nil, zero value otherwise.

### GetUserMobileDevicesUrlOk

`func (o *ApiUrlsResponse) GetUserMobileDevicesUrlOk() (*string, bool)`

GetUserMobileDevicesUrlOk returns a tuple with the UserMobileDevicesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMobileDevicesUrl

`func (o *ApiUrlsResponse) SetUserMobileDevicesUrl(v string)`

SetUserMobileDevicesUrl sets UserMobileDevicesUrl field to given value.

### HasUserMobileDevicesUrl

`func (o *ApiUrlsResponse) HasUserMobileDevicesUrl() bool`

HasUserMobileDevicesUrl returns a boolean if a field has been set.

### GetUserRegisterDevicesUrl

`func (o *ApiUrlsResponse) GetUserRegisterDevicesUrl() string`

GetUserRegisterDevicesUrl returns the UserRegisterDevicesUrl field if non-nil, zero value otherwise.

### GetUserRegisterDevicesUrlOk

`func (o *ApiUrlsResponse) GetUserRegisterDevicesUrlOk() (*string, bool)`

GetUserRegisterDevicesUrlOk returns a tuple with the UserRegisterDevicesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRegisterDevicesUrl

`func (o *ApiUrlsResponse) SetUserRegisterDevicesUrl(v string)`

SetUserRegisterDevicesUrl sets UserRegisterDevicesUrl field to given value.

### HasUserRegisterDevicesUrl

`func (o *ApiUrlsResponse) HasUserRegisterDevicesUrl() bool`

HasUserRegisterDevicesUrl returns a boolean if a field has been set.

### GetUserRoleUrl

`func (o *ApiUrlsResponse) GetUserRoleUrl() string`

GetUserRoleUrl returns the UserRoleUrl field if non-nil, zero value otherwise.

### GetUserRoleUrlOk

`func (o *ApiUrlsResponse) GetUserRoleUrlOk() (*string, bool)`

GetUserRoleUrlOk returns a tuple with the UserRoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoleUrl

`func (o *ApiUrlsResponse) SetUserRoleUrl(v string)`

SetUserRoleUrl sets UserRoleUrl field to given value.

### HasUserRoleUrl

`func (o *ApiUrlsResponse) HasUserRoleUrl() bool`

HasUserRoleUrl returns a boolean if a field has been set.

### GetUsersUrl

`func (o *ApiUrlsResponse) GetUsersUrl() string`

GetUsersUrl returns the UsersUrl field if non-nil, zero value otherwise.

### GetUsersUrlOk

`func (o *ApiUrlsResponse) GetUsersUrlOk() (*string, bool)`

GetUsersUrlOk returns a tuple with the UsersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersUrl

`func (o *ApiUrlsResponse) SetUsersUrl(v string)`

SetUsersUrl sets UsersUrl field to given value.

### HasUsersUrl

`func (o *ApiUrlsResponse) HasUsersUrl() bool`

HasUsersUrl returns a boolean if a field has been set.

### GetVoicemailsUrl

`func (o *ApiUrlsResponse) GetVoicemailsUrl() string`

GetVoicemailsUrl returns the VoicemailsUrl field if non-nil, zero value otherwise.

### GetVoicemailsUrlOk

`func (o *ApiUrlsResponse) GetVoicemailsUrlOk() (*string, bool)`

GetVoicemailsUrlOk returns a tuple with the VoicemailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemailsUrl

`func (o *ApiUrlsResponse) SetVoicemailsUrl(v string)`

SetVoicemailsUrl sets VoicemailsUrl field to given value.

### HasVoicemailsUrl

`func (o *ApiUrlsResponse) HasVoicemailsUrl() bool`

HasVoicemailsUrl returns a boolean if a field has been set.

### GetVoicemailDetailUrl

`func (o *ApiUrlsResponse) GetVoicemailDetailUrl() string`

GetVoicemailDetailUrl returns the VoicemailDetailUrl field if non-nil, zero value otherwise.

### GetVoicemailDetailUrlOk

`func (o *ApiUrlsResponse) GetVoicemailDetailUrlOk() (*string, bool)`

GetVoicemailDetailUrlOk returns a tuple with the VoicemailDetailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemailDetailUrl

`func (o *ApiUrlsResponse) SetVoicemailDetailUrl(v string)`

SetVoicemailDetailUrl sets VoicemailDetailUrl field to given value.

### HasVoicemailDetailUrl

`func (o *ApiUrlsResponse) HasVoicemailDetailUrl() bool`

HasVoicemailDetailUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


