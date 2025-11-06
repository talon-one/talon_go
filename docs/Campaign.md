# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**UserId** | Pointer to **int64** | The ID of the user associated with this entity. | 
**Name** | Pointer to **string** | A user-facing name for this campaign. | 
**Description** | Pointer to **string** | A detailed description of the campaign. | 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become inactive. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**State** | Pointer to **string** | A disabled or archived campaign is not evaluated for rules or coupons.  | [default to "enabled"]
**ActiveRulesetId** | Pointer to **int64** | [ID of Ruleset](https://docs.talon.one/management-api#operation/getRulesets) this campaign applies on customer session evaluation.  | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign. | 
**Features** | Pointer to **[]string** | The features enabled in this campaign. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | The set of [budget limits](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets) for this campaign.  | 
**CampaignGroups** | Pointer to **[]int64** | The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/managing-campaign-groups) this campaign belongs to.  | [optional] 
**Type** | Pointer to **string** | The campaign type. Possible type values:   - &#x60;cartItem&#x60;: Type of campaign that can apply effects only to cart items.   - &#x60;advanced&#x60;: Type of campaign that can apply effects to customer sessions and cart items.  | [default to "advanced"]
**LinkedStoreIds** | Pointer to **[]int64** | A list of store IDs that you want to link to the campaign.  **Note:** Campaigns with linked store IDs will only be evaluated when there is a [customer session update](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) that references a linked store.  | [optional] 
**Budgets** | Pointer to [**[]CampaignBudget**](CampaignBudget.md) | A list of all the budgets that are defined by this campaign and their usage.  **Note:** Budgets that are not defined do not appear in this list and their usage is not counted until they are defined.  | [optional] 
**CouponRedemptionCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Number of coupons redeemed in the campaign.  | [optional] 
**ReferralRedemptionCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Number of referral codes redeemed in the campaign.  | [optional] 
**DiscountCount** | Pointer to **float32** | This property is **deprecated**. The count should be available under *budgets* property. Total amount of discounts redeemed in the campaign.  | [optional] 
**DiscountEffectCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of times discounts were redeemed in this campaign.  | [optional] 
**CouponCreationCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of coupons created by rules in this campaign.  | [optional] 
**CustomEffectCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of custom effects triggered by rules in this campaign.  | [optional] 
**ReferralCreationCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of referrals created by rules in this campaign.  | [optional] 
**AddFreeItemEffectCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of times the [add free item effect](https://docs.talon.one/docs/dev/integration-api/api-effects#addfreeitem) can be triggered in this campaign.  | [optional] 
**AwardedGiveawaysCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of giveaways awarded by rules in this campaign.  | [optional] 
**CreatedLoyaltyPointsCount** | Pointer to **float32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty points created by rules in this campaign.  | [optional] 
**CreatedLoyaltyPointsEffectCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty point creation effects triggered by rules in this campaign.  | [optional] 
**RedeemedLoyaltyPointsCount** | Pointer to **float32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty points redeemed by rules in this campaign.  | [optional] 
**RedeemedLoyaltyPointsEffectCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty point redemption effects triggered by rules in this campaign.  | [optional] 
**CallApiEffectCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of webhooks triggered by rules in this campaign.  | [optional] 
**ReservecouponEffectCount** | Pointer to **int64** | This property is **deprecated**. The count should be available under *budgets* property. Total number of reserve coupon effects triggered by rules in this campaign.  | [optional] 
**LastActivity** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received by this campaign. | [optional] 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the campaign&#39;s property. Updates to external entities used in this campaign are **not** registered by this property, such as collection or coupon updates.  | [optional] 
**CreatedBy** | Pointer to **string** | Name of the user who created this campaign if available. | [optional] 
**UpdatedBy** | Pointer to **string** | Name of the user who last updated this campaign if available. | [optional] 
**TemplateId** | Pointer to **int64** | The ID of the Campaign Template this Campaign was created from. | [optional] 
**FrontendState** | Pointer to **string** | The campaign state displayed in the Campaign Manager. | 
**StoresImported** | Pointer to **bool** | Indicates whether the linked stores were imported via a CSV file. | 
**ValueMapsIds** | Pointer to **[]int64** | A list of value map IDs for the campaign. | [optional] 
**RevisionFrontendState** | Pointer to **string** | The campaign revision state displayed in the Campaign Manager. | [optional] 
**ActiveRevisionId** | Pointer to **int64** | ID of the revision that was last activated on this campaign.  | [optional] 
**ActiveRevisionVersionId** | Pointer to **int64** | ID of the revision version that is active on the campaign.  | [optional] 
**Version** | Pointer to **int64** | Incrementing number representing how many revisions have been activated on this campaign, starts from 0 for a new campaign.  | [optional] 
**CurrentRevisionId** | Pointer to **int64** | ID of the revision currently being modified for the campaign.  | [optional] 
**CurrentRevisionVersionId** | Pointer to **int64** | ID of the latest version applied on the current revision.  | [optional] 
**StageRevision** | Pointer to **bool** | Flag for determining whether we use current revision when sending requests with staging API key.  | [optional] [default to false]

## Methods

### NewCampaign

`func NewCampaign(id int64, created time.Time, applicationId int64, userId int64, name string, description string, state string, tags []string, features []string, limits []LimitConfig, type_ string, frontendState string, storesImported bool, ) *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Campaign) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Campaign) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Campaign) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Campaign) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *Campaign) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Campaign) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Campaign) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetUserId

`func (o *Campaign) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Campaign) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Campaign) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Campaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Campaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Campaign) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStartTime

`func (o *Campaign) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Campaign) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Campaign) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Campaign) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Campaign) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Campaign) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Campaign) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Campaign) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetAttributes

`func (o *Campaign) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Campaign) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Campaign) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Campaign) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetState

`func (o *Campaign) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Campaign) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Campaign) SetState(v string)`

SetState sets State field to given value.


### GetActiveRulesetId

`func (o *Campaign) GetActiveRulesetId() int64`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *Campaign) GetActiveRulesetIdOk() (*int64, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRulesetId

`func (o *Campaign) SetActiveRulesetId(v int64)`

SetActiveRulesetId sets ActiveRulesetId field to given value.

### HasActiveRulesetId

`func (o *Campaign) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### GetTags

`func (o *Campaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Campaign) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Campaign) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetFeatures

`func (o *Campaign) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Campaign) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Campaign) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetCouponSettings

`func (o *Campaign) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *Campaign) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *Campaign) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *Campaign) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetReferralSettings

`func (o *Campaign) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *Campaign) GetReferralSettingsOk() (*CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSettings

`func (o *Campaign) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings sets ReferralSettings field to given value.

### HasReferralSettings

`func (o *Campaign) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### GetLimits

`func (o *Campaign) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Campaign) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Campaign) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.


### GetCampaignGroups

`func (o *Campaign) GetCampaignGroups() []int64`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *Campaign) GetCampaignGroupsOk() (*[]int64, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignGroups

`func (o *Campaign) SetCampaignGroups(v []int64)`

SetCampaignGroups sets CampaignGroups field to given value.

### HasCampaignGroups

`func (o *Campaign) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### GetType

`func (o *Campaign) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Campaign) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Campaign) SetType(v string)`

SetType sets Type field to given value.


### GetLinkedStoreIds

`func (o *Campaign) GetLinkedStoreIds() []int64`

GetLinkedStoreIds returns the LinkedStoreIds field if non-nil, zero value otherwise.

### GetLinkedStoreIdsOk

`func (o *Campaign) GetLinkedStoreIdsOk() (*[]int64, bool)`

GetLinkedStoreIdsOk returns a tuple with the LinkedStoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedStoreIds

`func (o *Campaign) SetLinkedStoreIds(v []int64)`

SetLinkedStoreIds sets LinkedStoreIds field to given value.

### HasLinkedStoreIds

`func (o *Campaign) HasLinkedStoreIds() bool`

HasLinkedStoreIds returns a boolean if a field has been set.

### GetBudgets

`func (o *Campaign) GetBudgets() []CampaignBudget`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *Campaign) GetBudgetsOk() (*[]CampaignBudget, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *Campaign) SetBudgets(v []CampaignBudget)`

SetBudgets sets Budgets field to given value.

### HasBudgets

`func (o *Campaign) HasBudgets() bool`

HasBudgets returns a boolean if a field has been set.

### GetCouponRedemptionCount

`func (o *Campaign) GetCouponRedemptionCount() int64`

GetCouponRedemptionCount returns the CouponRedemptionCount field if non-nil, zero value otherwise.

### GetCouponRedemptionCountOk

`func (o *Campaign) GetCouponRedemptionCountOk() (*int64, bool)`

GetCouponRedemptionCountOk returns a tuple with the CouponRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRedemptionCount

`func (o *Campaign) SetCouponRedemptionCount(v int64)`

SetCouponRedemptionCount sets CouponRedemptionCount field to given value.

### HasCouponRedemptionCount

`func (o *Campaign) HasCouponRedemptionCount() bool`

HasCouponRedemptionCount returns a boolean if a field has been set.

### GetReferralRedemptionCount

`func (o *Campaign) GetReferralRedemptionCount() int64`

GetReferralRedemptionCount returns the ReferralRedemptionCount field if non-nil, zero value otherwise.

### GetReferralRedemptionCountOk

`func (o *Campaign) GetReferralRedemptionCountOk() (*int64, bool)`

GetReferralRedemptionCountOk returns a tuple with the ReferralRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralRedemptionCount

`func (o *Campaign) SetReferralRedemptionCount(v int64)`

SetReferralRedemptionCount sets ReferralRedemptionCount field to given value.

### HasReferralRedemptionCount

`func (o *Campaign) HasReferralRedemptionCount() bool`

HasReferralRedemptionCount returns a boolean if a field has been set.

### GetDiscountCount

`func (o *Campaign) GetDiscountCount() float32`

GetDiscountCount returns the DiscountCount field if non-nil, zero value otherwise.

### GetDiscountCountOk

`func (o *Campaign) GetDiscountCountOk() (*float32, bool)`

GetDiscountCountOk returns a tuple with the DiscountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCount

`func (o *Campaign) SetDiscountCount(v float32)`

SetDiscountCount sets DiscountCount field to given value.

### HasDiscountCount

`func (o *Campaign) HasDiscountCount() bool`

HasDiscountCount returns a boolean if a field has been set.

### GetDiscountEffectCount

`func (o *Campaign) GetDiscountEffectCount() int64`

GetDiscountEffectCount returns the DiscountEffectCount field if non-nil, zero value otherwise.

### GetDiscountEffectCountOk

`func (o *Campaign) GetDiscountEffectCountOk() (*int64, bool)`

GetDiscountEffectCountOk returns a tuple with the DiscountEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountEffectCount

`func (o *Campaign) SetDiscountEffectCount(v int64)`

SetDiscountEffectCount sets DiscountEffectCount field to given value.

### HasDiscountEffectCount

`func (o *Campaign) HasDiscountEffectCount() bool`

HasDiscountEffectCount returns a boolean if a field has been set.

### GetCouponCreationCount

`func (o *Campaign) GetCouponCreationCount() int64`

GetCouponCreationCount returns the CouponCreationCount field if non-nil, zero value otherwise.

### GetCouponCreationCountOk

`func (o *Campaign) GetCouponCreationCountOk() (*int64, bool)`

GetCouponCreationCountOk returns a tuple with the CouponCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCreationCount

`func (o *Campaign) SetCouponCreationCount(v int64)`

SetCouponCreationCount sets CouponCreationCount field to given value.

### HasCouponCreationCount

`func (o *Campaign) HasCouponCreationCount() bool`

HasCouponCreationCount returns a boolean if a field has been set.

### GetCustomEffectCount

`func (o *Campaign) GetCustomEffectCount() int64`

GetCustomEffectCount returns the CustomEffectCount field if non-nil, zero value otherwise.

### GetCustomEffectCountOk

`func (o *Campaign) GetCustomEffectCountOk() (*int64, bool)`

GetCustomEffectCountOk returns a tuple with the CustomEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEffectCount

`func (o *Campaign) SetCustomEffectCount(v int64)`

SetCustomEffectCount sets CustomEffectCount field to given value.

### HasCustomEffectCount

`func (o *Campaign) HasCustomEffectCount() bool`

HasCustomEffectCount returns a boolean if a field has been set.

### GetReferralCreationCount

`func (o *Campaign) GetReferralCreationCount() int64`

GetReferralCreationCount returns the ReferralCreationCount field if non-nil, zero value otherwise.

### GetReferralCreationCountOk

`func (o *Campaign) GetReferralCreationCountOk() (*int64, bool)`

GetReferralCreationCountOk returns a tuple with the ReferralCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCreationCount

`func (o *Campaign) SetReferralCreationCount(v int64)`

SetReferralCreationCount sets ReferralCreationCount field to given value.

### HasReferralCreationCount

`func (o *Campaign) HasReferralCreationCount() bool`

HasReferralCreationCount returns a boolean if a field has been set.

### GetAddFreeItemEffectCount

`func (o *Campaign) GetAddFreeItemEffectCount() int64`

GetAddFreeItemEffectCount returns the AddFreeItemEffectCount field if non-nil, zero value otherwise.

### GetAddFreeItemEffectCountOk

`func (o *Campaign) GetAddFreeItemEffectCountOk() (*int64, bool)`

GetAddFreeItemEffectCountOk returns a tuple with the AddFreeItemEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFreeItemEffectCount

`func (o *Campaign) SetAddFreeItemEffectCount(v int64)`

SetAddFreeItemEffectCount sets AddFreeItemEffectCount field to given value.

### HasAddFreeItemEffectCount

`func (o *Campaign) HasAddFreeItemEffectCount() bool`

HasAddFreeItemEffectCount returns a boolean if a field has been set.

### GetAwardedGiveawaysCount

`func (o *Campaign) GetAwardedGiveawaysCount() int64`

GetAwardedGiveawaysCount returns the AwardedGiveawaysCount field if non-nil, zero value otherwise.

### GetAwardedGiveawaysCountOk

`func (o *Campaign) GetAwardedGiveawaysCountOk() (*int64, bool)`

GetAwardedGiveawaysCountOk returns a tuple with the AwardedGiveawaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardedGiveawaysCount

`func (o *Campaign) SetAwardedGiveawaysCount(v int64)`

SetAwardedGiveawaysCount sets AwardedGiveawaysCount field to given value.

### HasAwardedGiveawaysCount

`func (o *Campaign) HasAwardedGiveawaysCount() bool`

HasAwardedGiveawaysCount returns a boolean if a field has been set.

### GetCreatedLoyaltyPointsCount

`func (o *Campaign) GetCreatedLoyaltyPointsCount() float32`

GetCreatedLoyaltyPointsCount returns the CreatedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsCountOk

`func (o *Campaign) GetCreatedLoyaltyPointsCountOk() (*float32, bool)`

GetCreatedLoyaltyPointsCountOk returns a tuple with the CreatedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedLoyaltyPointsCount

`func (o *Campaign) SetCreatedLoyaltyPointsCount(v float32)`

SetCreatedLoyaltyPointsCount sets CreatedLoyaltyPointsCount field to given value.

### HasCreatedLoyaltyPointsCount

`func (o *Campaign) HasCreatedLoyaltyPointsCount() bool`

HasCreatedLoyaltyPointsCount returns a boolean if a field has been set.

### GetCreatedLoyaltyPointsEffectCount

`func (o *Campaign) GetCreatedLoyaltyPointsEffectCount() int64`

GetCreatedLoyaltyPointsEffectCount returns the CreatedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsEffectCountOk

`func (o *Campaign) GetCreatedLoyaltyPointsEffectCountOk() (*int64, bool)`

GetCreatedLoyaltyPointsEffectCountOk returns a tuple with the CreatedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedLoyaltyPointsEffectCount

`func (o *Campaign) SetCreatedLoyaltyPointsEffectCount(v int64)`

SetCreatedLoyaltyPointsEffectCount sets CreatedLoyaltyPointsEffectCount field to given value.

### HasCreatedLoyaltyPointsEffectCount

`func (o *Campaign) HasCreatedLoyaltyPointsEffectCount() bool`

HasCreatedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### GetRedeemedLoyaltyPointsCount

`func (o *Campaign) GetRedeemedLoyaltyPointsCount() float32`

GetRedeemedLoyaltyPointsCount returns the RedeemedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsCountOk

`func (o *Campaign) GetRedeemedLoyaltyPointsCountOk() (*float32, bool)`

GetRedeemedLoyaltyPointsCountOk returns a tuple with the RedeemedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedLoyaltyPointsCount

`func (o *Campaign) SetRedeemedLoyaltyPointsCount(v float32)`

SetRedeemedLoyaltyPointsCount sets RedeemedLoyaltyPointsCount field to given value.

### HasRedeemedLoyaltyPointsCount

`func (o *Campaign) HasRedeemedLoyaltyPointsCount() bool`

HasRedeemedLoyaltyPointsCount returns a boolean if a field has been set.

### GetRedeemedLoyaltyPointsEffectCount

`func (o *Campaign) GetRedeemedLoyaltyPointsEffectCount() int64`

GetRedeemedLoyaltyPointsEffectCount returns the RedeemedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsEffectCountOk

`func (o *Campaign) GetRedeemedLoyaltyPointsEffectCountOk() (*int64, bool)`

GetRedeemedLoyaltyPointsEffectCountOk returns a tuple with the RedeemedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedLoyaltyPointsEffectCount

`func (o *Campaign) SetRedeemedLoyaltyPointsEffectCount(v int64)`

SetRedeemedLoyaltyPointsEffectCount sets RedeemedLoyaltyPointsEffectCount field to given value.

### HasRedeemedLoyaltyPointsEffectCount

`func (o *Campaign) HasRedeemedLoyaltyPointsEffectCount() bool`

HasRedeemedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### GetCallApiEffectCount

`func (o *Campaign) GetCallApiEffectCount() int64`

GetCallApiEffectCount returns the CallApiEffectCount field if non-nil, zero value otherwise.

### GetCallApiEffectCountOk

`func (o *Campaign) GetCallApiEffectCountOk() (*int64, bool)`

GetCallApiEffectCountOk returns a tuple with the CallApiEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallApiEffectCount

`func (o *Campaign) SetCallApiEffectCount(v int64)`

SetCallApiEffectCount sets CallApiEffectCount field to given value.

### HasCallApiEffectCount

`func (o *Campaign) HasCallApiEffectCount() bool`

HasCallApiEffectCount returns a boolean if a field has been set.

### GetReservecouponEffectCount

`func (o *Campaign) GetReservecouponEffectCount() int64`

GetReservecouponEffectCount returns the ReservecouponEffectCount field if non-nil, zero value otherwise.

### GetReservecouponEffectCountOk

`func (o *Campaign) GetReservecouponEffectCountOk() (*int64, bool)`

GetReservecouponEffectCountOk returns a tuple with the ReservecouponEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservecouponEffectCount

`func (o *Campaign) SetReservecouponEffectCount(v int64)`

SetReservecouponEffectCount sets ReservecouponEffectCount field to given value.

### HasReservecouponEffectCount

`func (o *Campaign) HasReservecouponEffectCount() bool`

HasReservecouponEffectCount returns a boolean if a field has been set.

### GetLastActivity

`func (o *Campaign) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *Campaign) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *Campaign) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *Campaign) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetUpdated

`func (o *Campaign) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Campaign) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Campaign) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Campaign) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Campaign) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Campaign) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Campaign) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Campaign) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Campaign) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Campaign) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Campaign) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Campaign) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetTemplateId

`func (o *Campaign) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Campaign) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Campaign) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Campaign) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetFrontendState

`func (o *Campaign) GetFrontendState() string`

GetFrontendState returns the FrontendState field if non-nil, zero value otherwise.

### GetFrontendStateOk

`func (o *Campaign) GetFrontendStateOk() (*string, bool)`

GetFrontendStateOk returns a tuple with the FrontendState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendState

`func (o *Campaign) SetFrontendState(v string)`

SetFrontendState sets FrontendState field to given value.


### GetStoresImported

`func (o *Campaign) GetStoresImported() bool`

GetStoresImported returns the StoresImported field if non-nil, zero value otherwise.

### GetStoresImportedOk

`func (o *Campaign) GetStoresImportedOk() (*bool, bool)`

GetStoresImportedOk returns a tuple with the StoresImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresImported

`func (o *Campaign) SetStoresImported(v bool)`

SetStoresImported sets StoresImported field to given value.


### GetValueMapsIds

`func (o *Campaign) GetValueMapsIds() []int64`

GetValueMapsIds returns the ValueMapsIds field if non-nil, zero value otherwise.

### GetValueMapsIdsOk

`func (o *Campaign) GetValueMapsIdsOk() (*[]int64, bool)`

GetValueMapsIdsOk returns a tuple with the ValueMapsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMapsIds

`func (o *Campaign) SetValueMapsIds(v []int64)`

SetValueMapsIds sets ValueMapsIds field to given value.

### HasValueMapsIds

`func (o *Campaign) HasValueMapsIds() bool`

HasValueMapsIds returns a boolean if a field has been set.

### GetRevisionFrontendState

`func (o *Campaign) GetRevisionFrontendState() string`

GetRevisionFrontendState returns the RevisionFrontendState field if non-nil, zero value otherwise.

### GetRevisionFrontendStateOk

`func (o *Campaign) GetRevisionFrontendStateOk() (*string, bool)`

GetRevisionFrontendStateOk returns a tuple with the RevisionFrontendState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionFrontendState

`func (o *Campaign) SetRevisionFrontendState(v string)`

SetRevisionFrontendState sets RevisionFrontendState field to given value.

### HasRevisionFrontendState

`func (o *Campaign) HasRevisionFrontendState() bool`

HasRevisionFrontendState returns a boolean if a field has been set.

### GetActiveRevisionId

`func (o *Campaign) GetActiveRevisionId() int64`

GetActiveRevisionId returns the ActiveRevisionId field if non-nil, zero value otherwise.

### GetActiveRevisionIdOk

`func (o *Campaign) GetActiveRevisionIdOk() (*int64, bool)`

GetActiveRevisionIdOk returns a tuple with the ActiveRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRevisionId

`func (o *Campaign) SetActiveRevisionId(v int64)`

SetActiveRevisionId sets ActiveRevisionId field to given value.

### HasActiveRevisionId

`func (o *Campaign) HasActiveRevisionId() bool`

HasActiveRevisionId returns a boolean if a field has been set.

### GetActiveRevisionVersionId

`func (o *Campaign) GetActiveRevisionVersionId() int64`

GetActiveRevisionVersionId returns the ActiveRevisionVersionId field if non-nil, zero value otherwise.

### GetActiveRevisionVersionIdOk

`func (o *Campaign) GetActiveRevisionVersionIdOk() (*int64, bool)`

GetActiveRevisionVersionIdOk returns a tuple with the ActiveRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRevisionVersionId

`func (o *Campaign) SetActiveRevisionVersionId(v int64)`

SetActiveRevisionVersionId sets ActiveRevisionVersionId field to given value.

### HasActiveRevisionVersionId

`func (o *Campaign) HasActiveRevisionVersionId() bool`

HasActiveRevisionVersionId returns a boolean if a field has been set.

### GetVersion

`func (o *Campaign) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Campaign) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Campaign) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Campaign) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCurrentRevisionId

`func (o *Campaign) GetCurrentRevisionId() int64`

GetCurrentRevisionId returns the CurrentRevisionId field if non-nil, zero value otherwise.

### GetCurrentRevisionIdOk

`func (o *Campaign) GetCurrentRevisionIdOk() (*int64, bool)`

GetCurrentRevisionIdOk returns a tuple with the CurrentRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevisionId

`func (o *Campaign) SetCurrentRevisionId(v int64)`

SetCurrentRevisionId sets CurrentRevisionId field to given value.

### HasCurrentRevisionId

`func (o *Campaign) HasCurrentRevisionId() bool`

HasCurrentRevisionId returns a boolean if a field has been set.

### GetCurrentRevisionVersionId

`func (o *Campaign) GetCurrentRevisionVersionId() int64`

GetCurrentRevisionVersionId returns the CurrentRevisionVersionId field if non-nil, zero value otherwise.

### GetCurrentRevisionVersionIdOk

`func (o *Campaign) GetCurrentRevisionVersionIdOk() (*int64, bool)`

GetCurrentRevisionVersionIdOk returns a tuple with the CurrentRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevisionVersionId

`func (o *Campaign) SetCurrentRevisionVersionId(v int64)`

SetCurrentRevisionVersionId sets CurrentRevisionVersionId field to given value.

### HasCurrentRevisionVersionId

`func (o *Campaign) HasCurrentRevisionVersionId() bool`

HasCurrentRevisionVersionId returns a boolean if a field has been set.

### GetStageRevision

`func (o *Campaign) GetStageRevision() bool`

GetStageRevision returns the StageRevision field if non-nil, zero value otherwise.

### GetStageRevisionOk

`func (o *Campaign) GetStageRevisionOk() (*bool, bool)`

GetStageRevisionOk returns a tuple with the StageRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageRevision

`func (o *Campaign) SetStageRevision(v bool)`

SetStageRevision sets StageRevision field to given value.

### HasStageRevision

`func (o *Campaign) HasStageRevision() bool`

HasStageRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


