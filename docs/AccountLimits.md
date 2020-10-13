# AccountLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LiveApplications** | Pointer to **int32** | Total number of allowed live applications in the account | 
**SandboxApplications** | Pointer to **int32** | Total number of allowed sandbox applications in the account | 
**ActiveCampaigns** | Pointer to **int32** | Total number of allowed active campaigns in live applications in the account | 
**Coupons** | Pointer to **int32** | Total number of allowed coupons in the account | 
**ReferralCodes** | Pointer to **int32** | Total number of allowed referral codes in the account | 
**ActiveRules** | Pointer to **int32** | Total number of allowed active rulesets in the account | 
**LiveLoyaltyPrograms** | Pointer to **int32** | Total number of allowed live loyalty programs in the account | 
**SandboxLoyaltyPrograms** | Pointer to **int32** | Total number of allowed sandbox loyalty programs in the account | 
**Webhooks** | Pointer to **int32** | Total number of allowed webhooks in the account | 
**Users** | Pointer to **int32** | Total number of allowed users in the account | 
**ApiVolume** | Pointer to **int32** | Allowed volume of API requests to the account | 
**PromotionTypes** | Pointer to **[]string** | Array of promotion types that are employed in the account | 

## Methods

### GetLiveApplications

`func (o *AccountLimits) GetLiveApplications() int32`

GetLiveApplications returns the LiveApplications field if non-nil, zero value otherwise.

### GetLiveApplicationsOk

`func (o *AccountLimits) GetLiveApplicationsOk() (int32, bool)`

GetLiveApplicationsOk returns a tuple with the LiveApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLiveApplications

`func (o *AccountLimits) HasLiveApplications() bool`

HasLiveApplications returns a boolean if a field has been set.

### SetLiveApplications

`func (o *AccountLimits) SetLiveApplications(v int32)`

SetLiveApplications gets a reference to the given int32 and assigns it to the LiveApplications field.

### GetSandboxApplications

`func (o *AccountLimits) GetSandboxApplications() int32`

GetSandboxApplications returns the SandboxApplications field if non-nil, zero value otherwise.

### GetSandboxApplicationsOk

`func (o *AccountLimits) GetSandboxApplicationsOk() (int32, bool)`

GetSandboxApplicationsOk returns a tuple with the SandboxApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandboxApplications

`func (o *AccountLimits) HasSandboxApplications() bool`

HasSandboxApplications returns a boolean if a field has been set.

### SetSandboxApplications

`func (o *AccountLimits) SetSandboxApplications(v int32)`

SetSandboxApplications gets a reference to the given int32 and assigns it to the SandboxApplications field.

### GetActiveCampaigns

`func (o *AccountLimits) GetActiveCampaigns() int32`

GetActiveCampaigns returns the ActiveCampaigns field if non-nil, zero value otherwise.

### GetActiveCampaignsOk

`func (o *AccountLimits) GetActiveCampaignsOk() (int32, bool)`

GetActiveCampaignsOk returns a tuple with the ActiveCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveCampaigns

`func (o *AccountLimits) HasActiveCampaigns() bool`

HasActiveCampaigns returns a boolean if a field has been set.

### SetActiveCampaigns

`func (o *AccountLimits) SetActiveCampaigns(v int32)`

SetActiveCampaigns gets a reference to the given int32 and assigns it to the ActiveCampaigns field.

### GetCoupons

`func (o *AccountLimits) GetCoupons() int32`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *AccountLimits) GetCouponsOk() (int32, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoupons

`func (o *AccountLimits) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### SetCoupons

`func (o *AccountLimits) SetCoupons(v int32)`

SetCoupons gets a reference to the given int32 and assigns it to the Coupons field.

### GetReferralCodes

`func (o *AccountLimits) GetReferralCodes() int32`

GetReferralCodes returns the ReferralCodes field if non-nil, zero value otherwise.

### GetReferralCodesOk

`func (o *AccountLimits) GetReferralCodesOk() (int32, bool)`

GetReferralCodesOk returns a tuple with the ReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralCodes

`func (o *AccountLimits) HasReferralCodes() bool`

HasReferralCodes returns a boolean if a field has been set.

### SetReferralCodes

`func (o *AccountLimits) SetReferralCodes(v int32)`

SetReferralCodes gets a reference to the given int32 and assigns it to the ReferralCodes field.

### GetActiveRules

`func (o *AccountLimits) GetActiveRules() int32`

GetActiveRules returns the ActiveRules field if non-nil, zero value otherwise.

### GetActiveRulesOk

`func (o *AccountLimits) GetActiveRulesOk() (int32, bool)`

GetActiveRulesOk returns a tuple with the ActiveRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRules

`func (o *AccountLimits) HasActiveRules() bool`

HasActiveRules returns a boolean if a field has been set.

### SetActiveRules

`func (o *AccountLimits) SetActiveRules(v int32)`

SetActiveRules gets a reference to the given int32 and assigns it to the ActiveRules field.

### GetLiveLoyaltyPrograms

`func (o *AccountLimits) GetLiveLoyaltyPrograms() int32`

GetLiveLoyaltyPrograms returns the LiveLoyaltyPrograms field if non-nil, zero value otherwise.

### GetLiveLoyaltyProgramsOk

`func (o *AccountLimits) GetLiveLoyaltyProgramsOk() (int32, bool)`

GetLiveLoyaltyProgramsOk returns a tuple with the LiveLoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLiveLoyaltyPrograms

`func (o *AccountLimits) HasLiveLoyaltyPrograms() bool`

HasLiveLoyaltyPrograms returns a boolean if a field has been set.

### SetLiveLoyaltyPrograms

`func (o *AccountLimits) SetLiveLoyaltyPrograms(v int32)`

SetLiveLoyaltyPrograms gets a reference to the given int32 and assigns it to the LiveLoyaltyPrograms field.

### GetSandboxLoyaltyPrograms

`func (o *AccountLimits) GetSandboxLoyaltyPrograms() int32`

GetSandboxLoyaltyPrograms returns the SandboxLoyaltyPrograms field if non-nil, zero value otherwise.

### GetSandboxLoyaltyProgramsOk

`func (o *AccountLimits) GetSandboxLoyaltyProgramsOk() (int32, bool)`

GetSandboxLoyaltyProgramsOk returns a tuple with the SandboxLoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandboxLoyaltyPrograms

`func (o *AccountLimits) HasSandboxLoyaltyPrograms() bool`

HasSandboxLoyaltyPrograms returns a boolean if a field has been set.

### SetSandboxLoyaltyPrograms

`func (o *AccountLimits) SetSandboxLoyaltyPrograms(v int32)`

SetSandboxLoyaltyPrograms gets a reference to the given int32 and assigns it to the SandboxLoyaltyPrograms field.

### GetWebhooks

`func (o *AccountLimits) GetWebhooks() int32`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AccountLimits) GetWebhooksOk() (int32, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhooks

`func (o *AccountLimits) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### SetWebhooks

`func (o *AccountLimits) SetWebhooks(v int32)`

SetWebhooks gets a reference to the given int32 and assigns it to the Webhooks field.

### GetUsers

`func (o *AccountLimits) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccountLimits) GetUsersOk() (int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsers

`func (o *AccountLimits) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsers

`func (o *AccountLimits) SetUsers(v int32)`

SetUsers gets a reference to the given int32 and assigns it to the Users field.

### GetApiVolume

`func (o *AccountLimits) GetApiVolume() int32`

GetApiVolume returns the ApiVolume field if non-nil, zero value otherwise.

### GetApiVolumeOk

`func (o *AccountLimits) GetApiVolumeOk() (int32, bool)`

GetApiVolumeOk returns a tuple with the ApiVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApiVolume

`func (o *AccountLimits) HasApiVolume() bool`

HasApiVolume returns a boolean if a field has been set.

### SetApiVolume

`func (o *AccountLimits) SetApiVolume(v int32)`

SetApiVolume gets a reference to the given int32 and assigns it to the ApiVolume field.

### GetPromotionTypes

`func (o *AccountLimits) GetPromotionTypes() []string`

GetPromotionTypes returns the PromotionTypes field if non-nil, zero value otherwise.

### GetPromotionTypesOk

`func (o *AccountLimits) GetPromotionTypesOk() ([]string, bool)`

GetPromotionTypesOk returns a tuple with the PromotionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPromotionTypes

`func (o *AccountLimits) HasPromotionTypes() bool`

HasPromotionTypes returns a boolean if a field has been set.

### SetPromotionTypes

`func (o *AccountLimits) SetPromotionTypes(v []string)`

SetPromotionTypes gets a reference to the given []string and assigns it to the PromotionTypes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


