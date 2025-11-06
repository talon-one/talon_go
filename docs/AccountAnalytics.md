# AccountAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to **int64** | Total number of applications in the account. | 
**LiveApplications** | Pointer to **int64** | Total number of live applications in the account. | 
**SandboxApplications** | Pointer to **int64** | Total number of sandbox applications in the account. | 
**Campaigns** | Pointer to **int64** | Total number of campaigns in the account. | 
**ActiveCampaigns** | Pointer to **int64** | Total number of active campaigns in the account. | 
**LiveActiveCampaigns** | Pointer to **int64** | Total number of active campaigns in live applications in the account. | 
**Coupons** | Pointer to **int64** | Total number of coupons in the account. | 
**ActiveCoupons** | Pointer to **int64** | Total number of active coupons in the account. | 
**ExpiredCoupons** | Pointer to **int64** | Total number of expired coupons in the account. | 
**ReferralCodes** | Pointer to **int64** | Total number of referral codes in the account. | 
**ActiveReferralCodes** | Pointer to **int64** | Total number of active referral codes in the account. | 
**ExpiredReferralCodes** | Pointer to **int64** | Total number of expired referral codes in the account. | 
**ActiveRules** | Pointer to **int64** | Total number of active rules in the account. | 
**Users** | Pointer to **int64** | Total number of users in the account. | 
**Roles** | Pointer to **int64** | Total number of roles in the account. | 
**CustomAttributes** | Pointer to **int64** | Total number of custom attributes in the account. | 
**Webhooks** | Pointer to **int64** | Total number of webhooks in the account. | 
**LoyaltyPrograms** | Pointer to **int64** | Total number of all loyalty programs in the account. | 
**LiveLoyaltyPrograms** | Pointer to **int64** | Total number of live loyalty programs in the account. | 
**LastUpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The point in time when the analytics numbers were updated last. | 

## Methods

### NewAccountAnalytics

`func NewAccountAnalytics(applications int64, liveApplications int64, sandboxApplications int64, campaigns int64, activeCampaigns int64, liveActiveCampaigns int64, coupons int64, activeCoupons int64, expiredCoupons int64, referralCodes int64, activeReferralCodes int64, expiredReferralCodes int64, activeRules int64, users int64, roles int64, customAttributes int64, webhooks int64, loyaltyPrograms int64, liveLoyaltyPrograms int64, lastUpdatedAt time.Time, ) *AccountAnalytics`

NewAccountAnalytics instantiates a new AccountAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAnalyticsWithDefaults

`func NewAccountAnalyticsWithDefaults() *AccountAnalytics`

NewAccountAnalyticsWithDefaults instantiates a new AccountAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *AccountAnalytics) GetApplications() int64`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *AccountAnalytics) GetApplicationsOk() (*int64, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *AccountAnalytics) SetApplications(v int64)`

SetApplications sets Applications field to given value.


### GetLiveApplications

`func (o *AccountAnalytics) GetLiveApplications() int64`

GetLiveApplications returns the LiveApplications field if non-nil, zero value otherwise.

### GetLiveApplicationsOk

`func (o *AccountAnalytics) GetLiveApplicationsOk() (*int64, bool)`

GetLiveApplicationsOk returns a tuple with the LiveApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveApplications

`func (o *AccountAnalytics) SetLiveApplications(v int64)`

SetLiveApplications sets LiveApplications field to given value.


### GetSandboxApplications

`func (o *AccountAnalytics) GetSandboxApplications() int64`

GetSandboxApplications returns the SandboxApplications field if non-nil, zero value otherwise.

### GetSandboxApplicationsOk

`func (o *AccountAnalytics) GetSandboxApplicationsOk() (*int64, bool)`

GetSandboxApplicationsOk returns a tuple with the SandboxApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxApplications

`func (o *AccountAnalytics) SetSandboxApplications(v int64)`

SetSandboxApplications sets SandboxApplications field to given value.


### GetCampaigns

`func (o *AccountAnalytics) GetCampaigns() int64`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *AccountAnalytics) GetCampaignsOk() (*int64, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *AccountAnalytics) SetCampaigns(v int64)`

SetCampaigns sets Campaigns field to given value.


### GetActiveCampaigns

`func (o *AccountAnalytics) GetActiveCampaigns() int64`

GetActiveCampaigns returns the ActiveCampaigns field if non-nil, zero value otherwise.

### GetActiveCampaignsOk

`func (o *AccountAnalytics) GetActiveCampaignsOk() (*int64, bool)`

GetActiveCampaignsOk returns a tuple with the ActiveCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCampaigns

`func (o *AccountAnalytics) SetActiveCampaigns(v int64)`

SetActiveCampaigns sets ActiveCampaigns field to given value.


### GetLiveActiveCampaigns

`func (o *AccountAnalytics) GetLiveActiveCampaigns() int64`

GetLiveActiveCampaigns returns the LiveActiveCampaigns field if non-nil, zero value otherwise.

### GetLiveActiveCampaignsOk

`func (o *AccountAnalytics) GetLiveActiveCampaignsOk() (*int64, bool)`

GetLiveActiveCampaignsOk returns a tuple with the LiveActiveCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveActiveCampaigns

`func (o *AccountAnalytics) SetLiveActiveCampaigns(v int64)`

SetLiveActiveCampaigns sets LiveActiveCampaigns field to given value.


### GetCoupons

`func (o *AccountAnalytics) GetCoupons() int64`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *AccountAnalytics) GetCouponsOk() (*int64, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *AccountAnalytics) SetCoupons(v int64)`

SetCoupons sets Coupons field to given value.


### GetActiveCoupons

`func (o *AccountAnalytics) GetActiveCoupons() int64`

GetActiveCoupons returns the ActiveCoupons field if non-nil, zero value otherwise.

### GetActiveCouponsOk

`func (o *AccountAnalytics) GetActiveCouponsOk() (*int64, bool)`

GetActiveCouponsOk returns a tuple with the ActiveCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCoupons

`func (o *AccountAnalytics) SetActiveCoupons(v int64)`

SetActiveCoupons sets ActiveCoupons field to given value.


### GetExpiredCoupons

`func (o *AccountAnalytics) GetExpiredCoupons() int64`

GetExpiredCoupons returns the ExpiredCoupons field if non-nil, zero value otherwise.

### GetExpiredCouponsOk

`func (o *AccountAnalytics) GetExpiredCouponsOk() (*int64, bool)`

GetExpiredCouponsOk returns a tuple with the ExpiredCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredCoupons

`func (o *AccountAnalytics) SetExpiredCoupons(v int64)`

SetExpiredCoupons sets ExpiredCoupons field to given value.


### GetReferralCodes

`func (o *AccountAnalytics) GetReferralCodes() int64`

GetReferralCodes returns the ReferralCodes field if non-nil, zero value otherwise.

### GetReferralCodesOk

`func (o *AccountAnalytics) GetReferralCodesOk() (*int64, bool)`

GetReferralCodesOk returns a tuple with the ReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCodes

`func (o *AccountAnalytics) SetReferralCodes(v int64)`

SetReferralCodes sets ReferralCodes field to given value.


### GetActiveReferralCodes

`func (o *AccountAnalytics) GetActiveReferralCodes() int64`

GetActiveReferralCodes returns the ActiveReferralCodes field if non-nil, zero value otherwise.

### GetActiveReferralCodesOk

`func (o *AccountAnalytics) GetActiveReferralCodesOk() (*int64, bool)`

GetActiveReferralCodesOk returns a tuple with the ActiveReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveReferralCodes

`func (o *AccountAnalytics) SetActiveReferralCodes(v int64)`

SetActiveReferralCodes sets ActiveReferralCodes field to given value.


### GetExpiredReferralCodes

`func (o *AccountAnalytics) GetExpiredReferralCodes() int64`

GetExpiredReferralCodes returns the ExpiredReferralCodes field if non-nil, zero value otherwise.

### GetExpiredReferralCodesOk

`func (o *AccountAnalytics) GetExpiredReferralCodesOk() (*int64, bool)`

GetExpiredReferralCodesOk returns a tuple with the ExpiredReferralCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredReferralCodes

`func (o *AccountAnalytics) SetExpiredReferralCodes(v int64)`

SetExpiredReferralCodes sets ExpiredReferralCodes field to given value.


### GetActiveRules

`func (o *AccountAnalytics) GetActiveRules() int64`

GetActiveRules returns the ActiveRules field if non-nil, zero value otherwise.

### GetActiveRulesOk

`func (o *AccountAnalytics) GetActiveRulesOk() (*int64, bool)`

GetActiveRulesOk returns a tuple with the ActiveRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRules

`func (o *AccountAnalytics) SetActiveRules(v int64)`

SetActiveRules sets ActiveRules field to given value.


### GetUsers

`func (o *AccountAnalytics) GetUsers() int64`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccountAnalytics) GetUsersOk() (*int64, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AccountAnalytics) SetUsers(v int64)`

SetUsers sets Users field to given value.


### GetRoles

`func (o *AccountAnalytics) GetRoles() int64`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccountAnalytics) GetRolesOk() (*int64, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccountAnalytics) SetRoles(v int64)`

SetRoles sets Roles field to given value.


### GetCustomAttributes

`func (o *AccountAnalytics) GetCustomAttributes() int64`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *AccountAnalytics) GetCustomAttributesOk() (*int64, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *AccountAnalytics) SetCustomAttributes(v int64)`

SetCustomAttributes sets CustomAttributes field to given value.


### GetWebhooks

`func (o *AccountAnalytics) GetWebhooks() int64`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AccountAnalytics) GetWebhooksOk() (*int64, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *AccountAnalytics) SetWebhooks(v int64)`

SetWebhooks sets Webhooks field to given value.


### GetLoyaltyPrograms

`func (o *AccountAnalytics) GetLoyaltyPrograms() int64`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *AccountAnalytics) GetLoyaltyProgramsOk() (*int64, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyPrograms

`func (o *AccountAnalytics) SetLoyaltyPrograms(v int64)`

SetLoyaltyPrograms sets LoyaltyPrograms field to given value.


### GetLiveLoyaltyPrograms

`func (o *AccountAnalytics) GetLiveLoyaltyPrograms() int64`

GetLiveLoyaltyPrograms returns the LiveLoyaltyPrograms field if non-nil, zero value otherwise.

### GetLiveLoyaltyProgramsOk

`func (o *AccountAnalytics) GetLiveLoyaltyProgramsOk() (*int64, bool)`

GetLiveLoyaltyProgramsOk returns a tuple with the LiveLoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveLoyaltyPrograms

`func (o *AccountAnalytics) SetLiveLoyaltyPrograms(v int64)`

SetLiveLoyaltyPrograms sets LiveLoyaltyPrograms field to given value.


### GetLastUpdatedAt

`func (o *AccountAnalytics) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *AccountAnalytics) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *AccountAnalytics) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


