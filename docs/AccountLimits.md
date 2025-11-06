# AccountLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LiveApplications** | Pointer to **int64** | Total number of allowed live applications in the account. | 
**SandboxApplications** | Pointer to **int64** | Total number of allowed sandbox applications in the account. | 
**ActiveCampaigns** | Pointer to **int64** | Total number of allowed active campaigns in live applications in the account. | 
**Coupons** | Pointer to **int64** | Total number of allowed coupons in the account. | 
**ReferralCodes** | Pointer to **int64** | Total number of allowed referral codes in the account. | 
**ActiveRules** | Pointer to **int64** | Total number of allowed active rulesets in the account. | 
**LiveLoyaltyPrograms** | Pointer to **int64** | Total number of allowed live loyalty programs in the account. | 
**SandboxLoyaltyPrograms** | Pointer to **int64** | Total number of allowed sandbox loyalty programs in the account. | 
**Webhooks** | Pointer to **int64** | Total number of allowed webhooks in the account. | 
**Users** | Pointer to **int64** | Total number of allowed users in the account. | 
**ApiVolume** | Pointer to **int64** | Allowed volume of API requests to the account. | 
**PromotionTypes** | Pointer to **[]string** | Array of promotion types that are employed in the account. | 
**SecondaryDeploymentPrice** | Pointer to **int64** | The price for a secondary deployment according to contractual agreements. | 
**CurrencyCode** | Pointer to **string** | The currency of the contract. | 

## Methods

### NewAccountLimits

`func NewAccountLimits(liveApplications int64, sandboxApplications int64, activeCampaigns int64, coupons int64, referralCodes int64, activeRules int64, liveLoyaltyPrograms int64, sandboxLoyaltyPrograms int64, webhooks int64, users int64, apiVolume int64, promotionTypes []string, secondaryDeploymentPrice int64, currencyCode string, ) *AccountLimits`

NewAccountLimits instantiates a new AccountLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLimitsWithDefaults

`func NewAccountLimitsWithDefaults() *AccountLimits`

NewAccountLimitsWithDefaults instantiates a new AccountLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiveApplications

`func (o *AccountLimits) GetLiveApplications() int64`

GetLiveApplications returns the LiveApplications field if non-nil, zero value otherwise.

### GetLiveApplicationsOk

`func (o *AccountLimits) GetLiveApplicationsOk() (*int64, bool)`

GetLiveApplicationsOk returns a tuple with the LiveApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveApplications

`func (o *AccountLimits) SetLiveApplications(v int64)`

SetLiveApplications sets LiveApplications field to given value.


### GetSandboxApplications

`func (o *AccountLimits) GetSandboxApplications() int64`

GetSandboxApplications returns the SandboxApplications field if non-nil, zero value otherwise.

### GetSandboxApplicationsOk

`func (o *AccountLimits) GetSandboxApplicationsOk() (*int64, bool)`

GetSandboxApplicationsOk returns a tuple with the SandboxApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxApplications

`func (o *AccountLimits) SetSandboxApplications(v int64)`

SetSandboxApplications sets SandboxApplications field to given value.


### GetActiveCampaigns

`func (o *AccountLimits) GetActiveCampaigns() int64`

GetActiveCampaigns returns the ActiveCampaigns field if non-nil, zero value otherwise.

### GetActiveCampaignsOk

`func (o *AccountLimits) GetActiveCampaignsOk() (*int64, bool)`

GetActiveCampaignsOk returns a tuple with the ActiveCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCampaigns

`func (o *AccountLimits) SetActiveCampaigns(v int64)`

SetActiveCampaigns sets ActiveCampaigns field to given value.


### GetCoupons

`func (o *AccountLimits) GetCoupons() int64`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *AccountLimits) GetCouponsOk() (*int64, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *AccountLimits) SetCoupons(v int64)`

SetCoupons sets Coupons field to given value.


### GetReferralCodes

`func (o *AccountLimits) GetReferralCodes() int64`

GetReferralCodes returns the ReferralCodes field if non-nil, zero value otherwise.

### GetReferralCodesOk

`func (o *AccountLimits) GetReferralCodesOk() (*int64, bool)`

GetReferralCodesOk returns a tuple with the ReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCodes

`func (o *AccountLimits) SetReferralCodes(v int64)`

SetReferralCodes sets ReferralCodes field to given value.


### GetActiveRules

`func (o *AccountLimits) GetActiveRules() int64`

GetActiveRules returns the ActiveRules field if non-nil, zero value otherwise.

### GetActiveRulesOk

`func (o *AccountLimits) GetActiveRulesOk() (*int64, bool)`

GetActiveRulesOk returns a tuple with the ActiveRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRules

`func (o *AccountLimits) SetActiveRules(v int64)`

SetActiveRules sets ActiveRules field to given value.


### GetLiveLoyaltyPrograms

`func (o *AccountLimits) GetLiveLoyaltyPrograms() int64`

GetLiveLoyaltyPrograms returns the LiveLoyaltyPrograms field if non-nil, zero value otherwise.

### GetLiveLoyaltyProgramsOk

`func (o *AccountLimits) GetLiveLoyaltyProgramsOk() (*int64, bool)`

GetLiveLoyaltyProgramsOk returns a tuple with the LiveLoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveLoyaltyPrograms

`func (o *AccountLimits) SetLiveLoyaltyPrograms(v int64)`

SetLiveLoyaltyPrograms sets LiveLoyaltyPrograms field to given value.


### GetSandboxLoyaltyPrograms

`func (o *AccountLimits) GetSandboxLoyaltyPrograms() int64`

GetSandboxLoyaltyPrograms returns the SandboxLoyaltyPrograms field if non-nil, zero value otherwise.

### GetSandboxLoyaltyProgramsOk

`func (o *AccountLimits) GetSandboxLoyaltyProgramsOk() (*int64, bool)`

GetSandboxLoyaltyProgramsOk returns a tuple with the SandboxLoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxLoyaltyPrograms

`func (o *AccountLimits) SetSandboxLoyaltyPrograms(v int64)`

SetSandboxLoyaltyPrograms sets SandboxLoyaltyPrograms field to given value.


### GetWebhooks

`func (o *AccountLimits) GetWebhooks() int64`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AccountLimits) GetWebhooksOk() (*int64, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *AccountLimits) SetWebhooks(v int64)`

SetWebhooks sets Webhooks field to given value.


### GetUsers

`func (o *AccountLimits) GetUsers() int64`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccountLimits) GetUsersOk() (*int64, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AccountLimits) SetUsers(v int64)`

SetUsers sets Users field to given value.


### GetApiVolume

`func (o *AccountLimits) GetApiVolume() int64`

GetApiVolume returns the ApiVolume field if non-nil, zero value otherwise.

### GetApiVolumeOk

`func (o *AccountLimits) GetApiVolumeOk() (*int64, bool)`

GetApiVolumeOk returns a tuple with the ApiVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVolume

`func (o *AccountLimits) SetApiVolume(v int64)`

SetApiVolume sets ApiVolume field to given value.


### GetPromotionTypes

`func (o *AccountLimits) GetPromotionTypes() []string`

GetPromotionTypes returns the PromotionTypes field if non-nil, zero value otherwise.

### GetPromotionTypesOk

`func (o *AccountLimits) GetPromotionTypesOk() (*[]string, bool)`

GetPromotionTypesOk returns a tuple with the PromotionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionTypes

`func (o *AccountLimits) SetPromotionTypes(v []string)`

SetPromotionTypes sets PromotionTypes field to given value.


### GetSecondaryDeploymentPrice

`func (o *AccountLimits) GetSecondaryDeploymentPrice() int64`

GetSecondaryDeploymentPrice returns the SecondaryDeploymentPrice field if non-nil, zero value otherwise.

### GetSecondaryDeploymentPriceOk

`func (o *AccountLimits) GetSecondaryDeploymentPriceOk() (*int64, bool)`

GetSecondaryDeploymentPriceOk returns a tuple with the SecondaryDeploymentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDeploymentPrice

`func (o *AccountLimits) SetSecondaryDeploymentPrice(v int64)`

SetSecondaryDeploymentPrice sets SecondaryDeploymentPrice field to given value.


### GetCurrencyCode

`func (o *AccountLimits) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountLimits) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountLimits) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


