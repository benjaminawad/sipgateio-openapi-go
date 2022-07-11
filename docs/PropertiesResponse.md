# PropertiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityVerificationRequired** | Pointer to **bool** |  | [optional] [default to false]
**VerificationLetterCodeRequired** | Pointer to **bool** |  | [optional] [default to false]
**ShowSipCredentials** | Pointer to **bool** |  | [optional] [default to false]
**ShowRegisterDndOption** | Pointer to **bool** |  | [optional] [default to false]
**ShowDisplaySingleLineOption** | Pointer to **bool** |  | [optional] [default to false]
**ShowClick2DialButton** | Pointer to **bool** |  | [optional] [default to false]
**ShowVerificationLetter** | Pointer to **bool** |  | [optional] [default to false]
**HasDedicatedAdminArea** | Pointer to **bool** |  | [optional] [default to false]
**ShowSimRelevantOptions** | Pointer to **bool** |  | [optional] [default to false]
**ShowESimOptions** | Pointer to **bool** |  | [optional] [default to false]
**ShowEventTimeToLiveHint** | Pointer to **bool** |  | [optional] [default to false]
**ShowNettoPrices** | Pointer to **bool** |  | [optional] [default to false]
**ShowBasicLogo** | Pointer to **bool** |  | [optional] [default to false]
**ShowSimquadratLogo** | Pointer to **bool** |  | [optional] [default to false]
**IdentityVerificationRequiredForMobile** | Pointer to **bool** |  | [optional] [default to false]
**ChatbotHuman** | Pointer to **bool** |  | [optional] [default to false]
**ChatbotUsingInformalLanguage** | Pointer to **bool** |  | [optional] [default to false]
**MultiSimProduct** | Pointer to **bool** |  | [optional] [default to false]
**ProductWithoutBaseFee** | Pointer to **bool** |  | [optional] [default to false]
**VoipOnlyProduct** | Pointer to **bool** |  | [optional] [default to false]
**MobileOnlyProduct** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPropertiesResponse

`func NewPropertiesResponse() *PropertiesResponse`

NewPropertiesResponse instantiates a new PropertiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertiesResponseWithDefaults

`func NewPropertiesResponseWithDefaults() *PropertiesResponse`

NewPropertiesResponseWithDefaults instantiates a new PropertiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityVerificationRequired

`func (o *PropertiesResponse) GetIdentityVerificationRequired() bool`

GetIdentityVerificationRequired returns the IdentityVerificationRequired field if non-nil, zero value otherwise.

### GetIdentityVerificationRequiredOk

`func (o *PropertiesResponse) GetIdentityVerificationRequiredOk() (*bool, bool)`

GetIdentityVerificationRequiredOk returns a tuple with the IdentityVerificationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityVerificationRequired

`func (o *PropertiesResponse) SetIdentityVerificationRequired(v bool)`

SetIdentityVerificationRequired sets IdentityVerificationRequired field to given value.

### HasIdentityVerificationRequired

`func (o *PropertiesResponse) HasIdentityVerificationRequired() bool`

HasIdentityVerificationRequired returns a boolean if a field has been set.

### GetVerificationLetterCodeRequired

`func (o *PropertiesResponse) GetVerificationLetterCodeRequired() bool`

GetVerificationLetterCodeRequired returns the VerificationLetterCodeRequired field if non-nil, zero value otherwise.

### GetVerificationLetterCodeRequiredOk

`func (o *PropertiesResponse) GetVerificationLetterCodeRequiredOk() (*bool, bool)`

GetVerificationLetterCodeRequiredOk returns a tuple with the VerificationLetterCodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLetterCodeRequired

`func (o *PropertiesResponse) SetVerificationLetterCodeRequired(v bool)`

SetVerificationLetterCodeRequired sets VerificationLetterCodeRequired field to given value.

### HasVerificationLetterCodeRequired

`func (o *PropertiesResponse) HasVerificationLetterCodeRequired() bool`

HasVerificationLetterCodeRequired returns a boolean if a field has been set.

### GetShowSipCredentials

`func (o *PropertiesResponse) GetShowSipCredentials() bool`

GetShowSipCredentials returns the ShowSipCredentials field if non-nil, zero value otherwise.

### GetShowSipCredentialsOk

`func (o *PropertiesResponse) GetShowSipCredentialsOk() (*bool, bool)`

GetShowSipCredentialsOk returns a tuple with the ShowSipCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSipCredentials

`func (o *PropertiesResponse) SetShowSipCredentials(v bool)`

SetShowSipCredentials sets ShowSipCredentials field to given value.

### HasShowSipCredentials

`func (o *PropertiesResponse) HasShowSipCredentials() bool`

HasShowSipCredentials returns a boolean if a field has been set.

### GetShowRegisterDndOption

`func (o *PropertiesResponse) GetShowRegisterDndOption() bool`

GetShowRegisterDndOption returns the ShowRegisterDndOption field if non-nil, zero value otherwise.

### GetShowRegisterDndOptionOk

`func (o *PropertiesResponse) GetShowRegisterDndOptionOk() (*bool, bool)`

GetShowRegisterDndOptionOk returns a tuple with the ShowRegisterDndOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRegisterDndOption

`func (o *PropertiesResponse) SetShowRegisterDndOption(v bool)`

SetShowRegisterDndOption sets ShowRegisterDndOption field to given value.

### HasShowRegisterDndOption

`func (o *PropertiesResponse) HasShowRegisterDndOption() bool`

HasShowRegisterDndOption returns a boolean if a field has been set.

### GetShowDisplaySingleLineOption

`func (o *PropertiesResponse) GetShowDisplaySingleLineOption() bool`

GetShowDisplaySingleLineOption returns the ShowDisplaySingleLineOption field if non-nil, zero value otherwise.

### GetShowDisplaySingleLineOptionOk

`func (o *PropertiesResponse) GetShowDisplaySingleLineOptionOk() (*bool, bool)`

GetShowDisplaySingleLineOptionOk returns a tuple with the ShowDisplaySingleLineOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDisplaySingleLineOption

`func (o *PropertiesResponse) SetShowDisplaySingleLineOption(v bool)`

SetShowDisplaySingleLineOption sets ShowDisplaySingleLineOption field to given value.

### HasShowDisplaySingleLineOption

`func (o *PropertiesResponse) HasShowDisplaySingleLineOption() bool`

HasShowDisplaySingleLineOption returns a boolean if a field has been set.

### GetShowClick2DialButton

`func (o *PropertiesResponse) GetShowClick2DialButton() bool`

GetShowClick2DialButton returns the ShowClick2DialButton field if non-nil, zero value otherwise.

### GetShowClick2DialButtonOk

`func (o *PropertiesResponse) GetShowClick2DialButtonOk() (*bool, bool)`

GetShowClick2DialButtonOk returns a tuple with the ShowClick2DialButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowClick2DialButton

`func (o *PropertiesResponse) SetShowClick2DialButton(v bool)`

SetShowClick2DialButton sets ShowClick2DialButton field to given value.

### HasShowClick2DialButton

`func (o *PropertiesResponse) HasShowClick2DialButton() bool`

HasShowClick2DialButton returns a boolean if a field has been set.

### GetShowVerificationLetter

`func (o *PropertiesResponse) GetShowVerificationLetter() bool`

GetShowVerificationLetter returns the ShowVerificationLetter field if non-nil, zero value otherwise.

### GetShowVerificationLetterOk

`func (o *PropertiesResponse) GetShowVerificationLetterOk() (*bool, bool)`

GetShowVerificationLetterOk returns a tuple with the ShowVerificationLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowVerificationLetter

`func (o *PropertiesResponse) SetShowVerificationLetter(v bool)`

SetShowVerificationLetter sets ShowVerificationLetter field to given value.

### HasShowVerificationLetter

`func (o *PropertiesResponse) HasShowVerificationLetter() bool`

HasShowVerificationLetter returns a boolean if a field has been set.

### GetHasDedicatedAdminArea

`func (o *PropertiesResponse) GetHasDedicatedAdminArea() bool`

GetHasDedicatedAdminArea returns the HasDedicatedAdminArea field if non-nil, zero value otherwise.

### GetHasDedicatedAdminAreaOk

`func (o *PropertiesResponse) GetHasDedicatedAdminAreaOk() (*bool, bool)`

GetHasDedicatedAdminAreaOk returns a tuple with the HasDedicatedAdminArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDedicatedAdminArea

`func (o *PropertiesResponse) SetHasDedicatedAdminArea(v bool)`

SetHasDedicatedAdminArea sets HasDedicatedAdminArea field to given value.

### HasHasDedicatedAdminArea

`func (o *PropertiesResponse) HasHasDedicatedAdminArea() bool`

HasHasDedicatedAdminArea returns a boolean if a field has been set.

### GetShowSimRelevantOptions

`func (o *PropertiesResponse) GetShowSimRelevantOptions() bool`

GetShowSimRelevantOptions returns the ShowSimRelevantOptions field if non-nil, zero value otherwise.

### GetShowSimRelevantOptionsOk

`func (o *PropertiesResponse) GetShowSimRelevantOptionsOk() (*bool, bool)`

GetShowSimRelevantOptionsOk returns a tuple with the ShowSimRelevantOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSimRelevantOptions

`func (o *PropertiesResponse) SetShowSimRelevantOptions(v bool)`

SetShowSimRelevantOptions sets ShowSimRelevantOptions field to given value.

### HasShowSimRelevantOptions

`func (o *PropertiesResponse) HasShowSimRelevantOptions() bool`

HasShowSimRelevantOptions returns a boolean if a field has been set.

### GetShowESimOptions

`func (o *PropertiesResponse) GetShowESimOptions() bool`

GetShowESimOptions returns the ShowESimOptions field if non-nil, zero value otherwise.

### GetShowESimOptionsOk

`func (o *PropertiesResponse) GetShowESimOptionsOk() (*bool, bool)`

GetShowESimOptionsOk returns a tuple with the ShowESimOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowESimOptions

`func (o *PropertiesResponse) SetShowESimOptions(v bool)`

SetShowESimOptions sets ShowESimOptions field to given value.

### HasShowESimOptions

`func (o *PropertiesResponse) HasShowESimOptions() bool`

HasShowESimOptions returns a boolean if a field has been set.

### GetShowEventTimeToLiveHint

`func (o *PropertiesResponse) GetShowEventTimeToLiveHint() bool`

GetShowEventTimeToLiveHint returns the ShowEventTimeToLiveHint field if non-nil, zero value otherwise.

### GetShowEventTimeToLiveHintOk

`func (o *PropertiesResponse) GetShowEventTimeToLiveHintOk() (*bool, bool)`

GetShowEventTimeToLiveHintOk returns a tuple with the ShowEventTimeToLiveHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEventTimeToLiveHint

`func (o *PropertiesResponse) SetShowEventTimeToLiveHint(v bool)`

SetShowEventTimeToLiveHint sets ShowEventTimeToLiveHint field to given value.

### HasShowEventTimeToLiveHint

`func (o *PropertiesResponse) HasShowEventTimeToLiveHint() bool`

HasShowEventTimeToLiveHint returns a boolean if a field has been set.

### GetShowNettoPrices

`func (o *PropertiesResponse) GetShowNettoPrices() bool`

GetShowNettoPrices returns the ShowNettoPrices field if non-nil, zero value otherwise.

### GetShowNettoPricesOk

`func (o *PropertiesResponse) GetShowNettoPricesOk() (*bool, bool)`

GetShowNettoPricesOk returns a tuple with the ShowNettoPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNettoPrices

`func (o *PropertiesResponse) SetShowNettoPrices(v bool)`

SetShowNettoPrices sets ShowNettoPrices field to given value.

### HasShowNettoPrices

`func (o *PropertiesResponse) HasShowNettoPrices() bool`

HasShowNettoPrices returns a boolean if a field has been set.

### GetShowBasicLogo

`func (o *PropertiesResponse) GetShowBasicLogo() bool`

GetShowBasicLogo returns the ShowBasicLogo field if non-nil, zero value otherwise.

### GetShowBasicLogoOk

`func (o *PropertiesResponse) GetShowBasicLogoOk() (*bool, bool)`

GetShowBasicLogoOk returns a tuple with the ShowBasicLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBasicLogo

`func (o *PropertiesResponse) SetShowBasicLogo(v bool)`

SetShowBasicLogo sets ShowBasicLogo field to given value.

### HasShowBasicLogo

`func (o *PropertiesResponse) HasShowBasicLogo() bool`

HasShowBasicLogo returns a boolean if a field has been set.

### GetShowSimquadratLogo

`func (o *PropertiesResponse) GetShowSimquadratLogo() bool`

GetShowSimquadratLogo returns the ShowSimquadratLogo field if non-nil, zero value otherwise.

### GetShowSimquadratLogoOk

`func (o *PropertiesResponse) GetShowSimquadratLogoOk() (*bool, bool)`

GetShowSimquadratLogoOk returns a tuple with the ShowSimquadratLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSimquadratLogo

`func (o *PropertiesResponse) SetShowSimquadratLogo(v bool)`

SetShowSimquadratLogo sets ShowSimquadratLogo field to given value.

### HasShowSimquadratLogo

`func (o *PropertiesResponse) HasShowSimquadratLogo() bool`

HasShowSimquadratLogo returns a boolean if a field has been set.

### GetIdentityVerificationRequiredForMobile

`func (o *PropertiesResponse) GetIdentityVerificationRequiredForMobile() bool`

GetIdentityVerificationRequiredForMobile returns the IdentityVerificationRequiredForMobile field if non-nil, zero value otherwise.

### GetIdentityVerificationRequiredForMobileOk

`func (o *PropertiesResponse) GetIdentityVerificationRequiredForMobileOk() (*bool, bool)`

GetIdentityVerificationRequiredForMobileOk returns a tuple with the IdentityVerificationRequiredForMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityVerificationRequiredForMobile

`func (o *PropertiesResponse) SetIdentityVerificationRequiredForMobile(v bool)`

SetIdentityVerificationRequiredForMobile sets IdentityVerificationRequiredForMobile field to given value.

### HasIdentityVerificationRequiredForMobile

`func (o *PropertiesResponse) HasIdentityVerificationRequiredForMobile() bool`

HasIdentityVerificationRequiredForMobile returns a boolean if a field has been set.

### GetChatbotHuman

`func (o *PropertiesResponse) GetChatbotHuman() bool`

GetChatbotHuman returns the ChatbotHuman field if non-nil, zero value otherwise.

### GetChatbotHumanOk

`func (o *PropertiesResponse) GetChatbotHumanOk() (*bool, bool)`

GetChatbotHumanOk returns a tuple with the ChatbotHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatbotHuman

`func (o *PropertiesResponse) SetChatbotHuman(v bool)`

SetChatbotHuman sets ChatbotHuman field to given value.

### HasChatbotHuman

`func (o *PropertiesResponse) HasChatbotHuman() bool`

HasChatbotHuman returns a boolean if a field has been set.

### GetChatbotUsingInformalLanguage

`func (o *PropertiesResponse) GetChatbotUsingInformalLanguage() bool`

GetChatbotUsingInformalLanguage returns the ChatbotUsingInformalLanguage field if non-nil, zero value otherwise.

### GetChatbotUsingInformalLanguageOk

`func (o *PropertiesResponse) GetChatbotUsingInformalLanguageOk() (*bool, bool)`

GetChatbotUsingInformalLanguageOk returns a tuple with the ChatbotUsingInformalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatbotUsingInformalLanguage

`func (o *PropertiesResponse) SetChatbotUsingInformalLanguage(v bool)`

SetChatbotUsingInformalLanguage sets ChatbotUsingInformalLanguage field to given value.

### HasChatbotUsingInformalLanguage

`func (o *PropertiesResponse) HasChatbotUsingInformalLanguage() bool`

HasChatbotUsingInformalLanguage returns a boolean if a field has been set.

### GetMultiSimProduct

`func (o *PropertiesResponse) GetMultiSimProduct() bool`

GetMultiSimProduct returns the MultiSimProduct field if non-nil, zero value otherwise.

### GetMultiSimProductOk

`func (o *PropertiesResponse) GetMultiSimProductOk() (*bool, bool)`

GetMultiSimProductOk returns a tuple with the MultiSimProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSimProduct

`func (o *PropertiesResponse) SetMultiSimProduct(v bool)`

SetMultiSimProduct sets MultiSimProduct field to given value.

### HasMultiSimProduct

`func (o *PropertiesResponse) HasMultiSimProduct() bool`

HasMultiSimProduct returns a boolean if a field has been set.

### GetProductWithoutBaseFee

`func (o *PropertiesResponse) GetProductWithoutBaseFee() bool`

GetProductWithoutBaseFee returns the ProductWithoutBaseFee field if non-nil, zero value otherwise.

### GetProductWithoutBaseFeeOk

`func (o *PropertiesResponse) GetProductWithoutBaseFeeOk() (*bool, bool)`

GetProductWithoutBaseFeeOk returns a tuple with the ProductWithoutBaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductWithoutBaseFee

`func (o *PropertiesResponse) SetProductWithoutBaseFee(v bool)`

SetProductWithoutBaseFee sets ProductWithoutBaseFee field to given value.

### HasProductWithoutBaseFee

`func (o *PropertiesResponse) HasProductWithoutBaseFee() bool`

HasProductWithoutBaseFee returns a boolean if a field has been set.

### GetVoipOnlyProduct

`func (o *PropertiesResponse) GetVoipOnlyProduct() bool`

GetVoipOnlyProduct returns the VoipOnlyProduct field if non-nil, zero value otherwise.

### GetVoipOnlyProductOk

`func (o *PropertiesResponse) GetVoipOnlyProductOk() (*bool, bool)`

GetVoipOnlyProductOk returns a tuple with the VoipOnlyProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoipOnlyProduct

`func (o *PropertiesResponse) SetVoipOnlyProduct(v bool)`

SetVoipOnlyProduct sets VoipOnlyProduct field to given value.

### HasVoipOnlyProduct

`func (o *PropertiesResponse) HasVoipOnlyProduct() bool`

HasVoipOnlyProduct returns a boolean if a field has been set.

### GetMobileOnlyProduct

`func (o *PropertiesResponse) GetMobileOnlyProduct() bool`

GetMobileOnlyProduct returns the MobileOnlyProduct field if non-nil, zero value otherwise.

### GetMobileOnlyProductOk

`func (o *PropertiesResponse) GetMobileOnlyProductOk() (*bool, bool)`

GetMobileOnlyProductOk returns a tuple with the MobileOnlyProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileOnlyProduct

`func (o *PropertiesResponse) SetMobileOnlyProduct(v bool)`

SetMobileOnlyProduct sets MobileOnlyProduct field to given value.

### HasMobileOnlyProduct

`func (o *PropertiesResponse) HasMobileOnlyProduct() bool`

HasMobileOnlyProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


