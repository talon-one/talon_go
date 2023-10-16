# AccountAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to **int32** | Total number of applications in the account. | 
**LiveApplications** | Pointer to **int32** | Total number of live applications in the account. | 
**SandboxApplications** | Pointer to **int32** | Total number of sandbox applications in the account. | 
**Campaigns** | Pointer to **int32** | Total number of campaigns in the account. | 
**ActiveCampaigns** | Pointer to **int32** | Total number of active campaigns in the account. | 
**LiveActiveCampaigns** | Pointer to **int32** | Total number of active campaigns in live applications in the account. | 
**Coupons** | Pointer to **int32** | Total number of coupons in the account. | 
**ActiveCoupons** | Pointer to **int32** | Total number of active coupons in the account. | 
**ExpiredCoupons** | Pointer to **int32** | Total number of expired coupons in the account. | 
**ReferralCodes** | Pointer to **int32** | Total number of referral codes in the account. | 
**ActiveReferralCodes** | Pointer to **int32** | Total number of active referral codes in the account. | 
**ExpiredReferralCodes** | Pointer to **int32** | Total number of expired referral codes in the account. | 
**ActiveRules** | Pointer to **int32** | Total number of active rules in the account. | 
**Users** | Pointer to **int32** | Total number of users in the account. | 
**Roles** | Pointer to **int32** | Total number of roles in the account. | 
**CustomAttributes** | Pointer to **int32** | Total number of custom attributes in the account. | 
**Webhooks** | Pointer to **int32** | Total number of webhooks in the account. | 
**LoyaltyPrograms** | Pointer to **int32** | Total number of all loyalty programs in the account. | 
**LiveLoyaltyPrograms** | Pointer to **int32** | Total number of live loyalty programs in the account. | 
**LastUpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The point in time when the analytics numbers were updated last. | 

## Methods

### GetApplications

`func (o *AccountAnalytics) GetApplications() int32`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *AccountAnalytics) GetApplicationsOk() (int32, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplications

`func (o *AccountAnalytics) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplications

`func (o *AccountAnalytics) SetApplications(v int32)`

SetApplications gets a reference to the given int32 and assigns it to the Applications field.

### GetLiveApplications

`func (o *AccountAnalytics) GetLiveApplications() int32`

GetLiveApplications returns the LiveApplications field if non-nil, zero value otherwise.

### GetLiveApplicationsOk

`func (o *AccountAnalytics) GetLiveApplicationsOk() (int32, bool)`

GetLiveApplicationsOk returns a tuple with the LiveApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLiveApplications

`func (o *AccountAnalytics) HasLiveApplications() bool`

HasLiveApplications returns a boolean if a field has been set.

### SetLiveApplications

`func (o *AccountAnalytics) SetLiveApplications(v int32)`

SetLiveApplications gets a reference to the given int32 and assigns it to the LiveApplications field.

### GetSandboxApplications

`func (o *AccountAnalytics) GetSandboxApplications() int32`

GetSandboxApplications returns the SandboxApplications field if non-nil, zero value otherwise.

### GetSandboxApplicationsOk

`func (o *AccountAnalytics) GetSandboxApplicationsOk() (int32, bool)`

GetSandboxApplicationsOk returns a tuple with the SandboxApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandboxApplications

`func (o *AccountAnalytics) HasSandboxApplications() bool`

HasSandboxApplications returns a boolean if a field has been set.

### SetSandboxApplications

`func (o *AccountAnalytics) SetSandboxApplications(v int32)`

SetSandboxApplications gets a reference to the given int32 and assigns it to the SandboxApplications field.

### GetCampaigns

`func (o *AccountAnalytics) GetCampaigns() int32`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *AccountAnalytics) GetCampaignsOk() (int32, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaigns

`func (o *AccountAnalytics) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### SetCampaigns

`func (o *AccountAnalytics) SetCampaigns(v int32)`

SetCampaigns gets a reference to the given int32 and assigns it to the Campaigns field.

### GetActiveCampaigns

`func (o *AccountAnalytics) GetActiveCampaigns() int32`

GetActiveCampaigns returns the ActiveCampaigns field if non-nil, zero value otherwise.

### GetActiveCampaignsOk

`func (o *AccountAnalytics) GetActiveCampaignsOk() (int32, bool)`

GetActiveCampaignsOk returns a tuple with the ActiveCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveCampaigns

`func (o *AccountAnalytics) HasActiveCampaigns() bool`

HasActiveCampaigns returns a boolean if a field has been set.

### SetActiveCampaigns

`func (o *AccountAnalytics) SetActiveCampaigns(v int32)`

SetActiveCampaigns gets a reference to the given int32 and assigns it to the ActiveCampaigns field.

### GetLiveActiveCampaigns

`func (o *AccountAnalytics) GetLiveActiveCampaigns() int32`

GetLiveActiveCampaigns returns the LiveActiveCampaigns field if non-nil, zero value otherwise.

### GetLiveActiveCampaignsOk

`func (o *AccountAnalytics) GetLiveActiveCampaignsOk() (int32, bool)`

GetLiveActiveCampaignsOk returns a tuple with the LiveActiveCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLiveActiveCampaigns

`func (o *AccountAnalytics) HasLiveActiveCampaigns() bool`

HasLiveActiveCampaigns returns a boolean if a field has been set.

### SetLiveActiveCampaigns

`func (o *AccountAnalytics) SetLiveActiveCampaigns(v int32)`

SetLiveActiveCampaigns gets a reference to the given int32 and assigns it to the LiveActiveCampaigns field.

### GetCoupons

`func (o *AccountAnalytics) GetCoupons() int32`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *AccountAnalytics) GetCouponsOk() (int32, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoupons

`func (o *AccountAnalytics) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### SetCoupons

`func (o *AccountAnalytics) SetCoupons(v int32)`

SetCoupons gets a reference to the given int32 and assigns it to the Coupons field.

### GetActiveCoupons

`func (o *AccountAnalytics) GetActiveCoupons() int32`

GetActiveCoupons returns the ActiveCoupons field if non-nil, zero value otherwise.

### GetActiveCouponsOk

`func (o *AccountAnalytics) GetActiveCouponsOk() (int32, bool)`

GetActiveCouponsOk returns a tuple with the ActiveCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveCoupons

`func (o *AccountAnalytics) HasActiveCoupons() bool`

HasActiveCoupons returns a boolean if a field has been set.

### SetActiveCoupons

`func (o *AccountAnalytics) SetActiveCoupons(v int32)`

SetActiveCoupons gets a reference to the given int32 and assigns it to the ActiveCoupons field.

### GetExpiredCoupons

`func (o *AccountAnalytics) GetExpiredCoupons() int32`

GetExpiredCoupons returns the ExpiredCoupons field if non-nil, zero value otherwise.

### GetExpiredCouponsOk

`func (o *AccountAnalytics) GetExpiredCouponsOk() (int32, bool)`

GetExpiredCouponsOk returns a tuple with the ExpiredCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiredCoupons

`func (o *AccountAnalytics) HasExpiredCoupons() bool`

HasExpiredCoupons returns a boolean if a field has been set.

### SetExpiredCoupons

`func (o *AccountAnalytics) SetExpiredCoupons(v int32)`

SetExpiredCoupons gets a reference to the given int32 and assigns it to the ExpiredCoupons field.

### GetReferralCodes

`func (o *AccountAnalytics) GetReferralCodes() int32`

GetReferralCodes returns the ReferralCodes field if non-nil, zero value otherwise.

### GetReferralCodesOk

`func (o *AccountAnalytics) GetReferralCodesOk() (int32, bool)`

GetReferralCodesOk returns a tuple with the ReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralCodes

`func (o *AccountAnalytics) HasReferralCodes() bool`

HasReferralCodes returns a boolean if a field has been set.

### SetReferralCodes

`func (o *AccountAnalytics) SetReferralCodes(v int32)`

SetReferralCodes gets a reference to the given int32 and assigns it to the ReferralCodes field.

### GetActiveReferralCodes

`func (o *AccountAnalytics) GetActiveReferralCodes() int32`

GetActiveReferralCodes returns the ActiveReferralCodes field if non-nil, zero value otherwise.

### GetActiveReferralCodesOk

`func (o *AccountAnalytics) GetActiveReferralCodesOk() (int32, bool)`

GetActiveReferralCodesOk returns a tuple with the ActiveReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveReferralCodes

`func (o *AccountAnalytics) HasActiveReferralCodes() bool`

HasActiveReferralCodes returns a boolean if a field has been set.

### SetActiveReferralCodes

`func (o *AccountAnalytics) SetActiveReferralCodes(v int32)`

SetActiveReferralCodes gets a reference to the given int32 and assigns it to the ActiveReferralCodes field.

### GetExpiredReferralCodes

`func (o *AccountAnalytics) GetExpiredReferralCodes() int32`

GetExpiredReferralCodes returns the ExpiredReferralCodes field if non-nil, zero value otherwise.

### GetExpiredReferralCodesOk

`func (o *AccountAnalytics) GetExpiredReferralCodesOk() (int32, bool)`

GetExpiredReferralCodesOk returns a tuple with the ExpiredReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiredReferralCodes

`func (o *AccountAnalytics) HasExpiredReferralCodes() bool`

HasExpiredReferralCodes returns a boolean if a field has been set.

### SetExpiredReferralCodes

`func (o *AccountAnalytics) SetExpiredReferralCodes(v int32)`

SetExpiredReferralCodes gets a reference to the given int32 and assigns it to the ExpiredReferralCodes field.

### GetActiveRules

`func (o *AccountAnalytics) GetActiveRules() int32`

GetActiveRules returns the ActiveRules field if non-nil, zero value otherwise.

### GetActiveRulesOk

`func (o *AccountAnalytics) GetActiveRulesOk() (int32, bool)`

GetActiveRulesOk returns a tuple with the ActiveRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRules

`func (o *AccountAnalytics) HasActiveRules() bool`

HasActiveRules returns a boolean if a field has been set.

### SetActiveRules

`func (o *AccountAnalytics) SetActiveRules(v int32)`

SetActiveRules gets a reference to the given int32 and assigns it to the ActiveRules field.

### GetUsers

`func (o *AccountAnalytics) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccountAnalytics) GetUsersOk() (int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsers

`func (o *AccountAnalytics) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsers

`func (o *AccountAnalytics) SetUsers(v int32)`

SetUsers gets a reference to the given int32 and assigns it to the Users field.

### GetRoles

`func (o *AccountAnalytics) GetRoles() int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccountAnalytics) GetRolesOk() (int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *AccountAnalytics) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *AccountAnalytics) SetRoles(v int32)`

SetRoles gets a reference to the given int32 and assigns it to the Roles field.

### GetCustomAttributes

`func (o *AccountAnalytics) GetCustomAttributes() int32`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *AccountAnalytics) GetCustomAttributesOk() (int32, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomAttributes

`func (o *AccountAnalytics) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### SetCustomAttributes

`func (o *AccountAnalytics) SetCustomAttributes(v int32)`

SetCustomAttributes gets a reference to the given int32 and assigns it to the CustomAttributes field.

### GetWebhooks

`func (o *AccountAnalytics) GetWebhooks() int32`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AccountAnalytics) GetWebhooksOk() (int32, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhooks

`func (o *AccountAnalytics) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### SetWebhooks

`func (o *AccountAnalytics) SetWebhooks(v int32)`

SetWebhooks gets a reference to the given int32 and assigns it to the Webhooks field.

### GetLoyaltyPrograms

`func (o *AccountAnalytics) GetLoyaltyPrograms() int32`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *AccountAnalytics) GetLoyaltyProgramsOk() (int32, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyPrograms

`func (o *AccountAnalytics) HasLoyaltyPrograms() bool`

HasLoyaltyPrograms returns a boolean if a field has been set.

### SetLoyaltyPrograms

`func (o *AccountAnalytics) SetLoyaltyPrograms(v int32)`

SetLoyaltyPrograms gets a reference to the given int32 and assigns it to the LoyaltyPrograms field.

### GetLiveLoyaltyPrograms

`func (o *AccountAnalytics) GetLiveLoyaltyPrograms() int32`

GetLiveLoyaltyPrograms returns the LiveLoyaltyPrograms field if non-nil, zero value otherwise.

### GetLiveLoyaltyProgramsOk

`func (o *AccountAnalytics) GetLiveLoyaltyProgramsOk() (int32, bool)`

GetLiveLoyaltyProgramsOk returns a tuple with the LiveLoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLiveLoyaltyPrograms

`func (o *AccountAnalytics) HasLiveLoyaltyPrograms() bool`

HasLiveLoyaltyPrograms returns a boolean if a field has been set.

### SetLiveLoyaltyPrograms

`func (o *AccountAnalytics) SetLiveLoyaltyPrograms(v int32)`

SetLiveLoyaltyPrograms gets a reference to the given int32 and assigns it to the LiveLoyaltyPrograms field.

### GetLastUpdatedAt

`func (o *AccountAnalytics) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *AccountAnalytics) GetLastUpdatedAtOk() (time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdatedAt

`func (o *AccountAnalytics) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### SetLastUpdatedAt

`func (o *AccountAnalytics) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


