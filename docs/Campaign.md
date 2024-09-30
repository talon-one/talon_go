# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**UserId** | Pointer to **int32** | The ID of the user associated with this entity. | 
**Name** | Pointer to **string** | A user-facing name for this campaign. | 
**Description** | Pointer to **string** | A detailed description of the campaign. | 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become inactive. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**State** | Pointer to **string** | A disabled or archived campaign is not evaluated for rules or coupons.  | [default to STATE_ENABLED]
**ActiveRulesetId** | Pointer to **int32** | [ID of Ruleset](https://docs.talon.one/management-api#operation/getRulesets) this campaign applies on customer session evaluation.  | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign. | 
**Features** | Pointer to **[]string** | The features enabled in this campaign. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**ReferralSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | The set of [budget limits](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets) for this campaign.  | 
**CampaignGroups** | Pointer to **[]int32** | The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/managing-campaign-groups) this campaign belongs to.  | [optional] 
**Type** | Pointer to **string** | The campaign type. Possible type values:   - &#x60;cartItem&#x60;: Type of campaign that can apply effects only to cart items.   - &#x60;advanced&#x60;: Type of campaign that can apply effects to customer sessions and cart items.  | [default to TYPE_ADVANCED]
**LinkedStoreIds** | Pointer to **[]int32** | A list of store IDs that you want to link to the campaign.  **Note:** Campaigns with linked store IDs will only be evaluated when there is a [customer session update](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) that references a linked store.  | [optional] 
**Budgets** | Pointer to [**[]CampaignBudget**](CampaignBudget.md) | A list of all the budgets that are defined by this campaign and their usage.  **Note:** Budgets that are not defined do not appear in this list and their usage is not counted until they are defined.  | 
**CouponRedemptionCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Number of coupons redeemed in the campaign.  | [optional] 
**ReferralRedemptionCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Number of referral codes redeemed in the campaign.  | [optional] 
**DiscountCount** | Pointer to **float32** | This property is **deprecated**. The count should be available under *budgets* property. Total amount of discounts redeemed in the campaign.  | [optional] 
**DiscountEffectCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of times discounts were redeemed in this campaign.  | [optional] 
**CouponCreationCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of coupons created by rules in this campaign.  | [optional] 
**CustomEffectCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of custom effects triggered by rules in this campaign.  | [optional] 
**ReferralCreationCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of referrals created by rules in this campaign.  | [optional] 
**AddFreeItemEffectCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of times the [add free item effect](https://docs.talon.one/docs/dev/integration-api/api-effects#addfreeitem) can be triggered in this campaign.  | [optional] 
**AwardedGiveawaysCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of giveaways awarded by rules in this campaign.  | [optional] 
**CreatedLoyaltyPointsCount** | Pointer to **float32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty points created by rules in this campaign.  | [optional] 
**CreatedLoyaltyPointsEffectCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty point creation effects triggered by rules in this campaign.  | [optional] 
**RedeemedLoyaltyPointsCount** | Pointer to **float32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty points redeemed by rules in this campaign.  | [optional] 
**RedeemedLoyaltyPointsEffectCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty point redemption effects triggered by rules in this campaign.  | [optional] 
**CallApiEffectCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of webhooks triggered by rules in this campaign.  | [optional] 
**ReservecouponEffectCount** | Pointer to **int32** | This property is **deprecated**. The count should be available under *budgets* property. Total number of reserve coupon effects triggered by rules in this campaign.  | [optional] 
**LastActivity** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received by this campaign. | [optional] 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the campaign&#39;s property. Updates to external entities used in this campaign are **not** registered by this property, such as collection or coupon updates.  | [optional] 
**CreatedBy** | Pointer to **string** | Name of the user who created this campaign if available. | [optional] 
**UpdatedBy** | Pointer to **string** | Name of the user who last updated this campaign if available. | [optional] 
**TemplateId** | Pointer to **int32** | The ID of the Campaign Template this Campaign was created from. | [optional] 
**FrontendState** | Pointer to **string** | A campaign state described exactly as in the Campaign Manager. | 
**StoresImported** | Pointer to **bool** | Indicates whether the linked stores were imported via a CSV file. | 
**ActiveRevisionId** | Pointer to **int32** | ID of the revision that was last activated on this campaign.  | [optional] 
**ActiveRevisionVersionId** | Pointer to **int32** | ID of the revision version that is active on the campaign.  | [optional] 
**Version** | Pointer to **int32** | Incrementing number representing how many revisions have been activated on this campaign, starts from 0 for a new campaign.  | [optional] 
**CurrentRevisionId** | Pointer to **int32** | ID of the revision currently being modified for the campaign.  | [optional] 
**CurrentRevisionVersionId** | Pointer to **int32** | ID of the latest version applied on the current revision.  | [optional] 
**StageRevision** | Pointer to **bool** | Flag for determining whether we use current revision when sending requests with staging API key.  | [optional] [default to false]

## Methods

### GetId

`func (o *Campaign) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Campaign) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Campaign) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Campaign) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Campaign) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Campaign) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Campaign) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *Campaign) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Campaign) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *Campaign) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *Campaign) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetUserId

`func (o *Campaign) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Campaign) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *Campaign) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *Campaign) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.

### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Campaign) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Campaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Campaign) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Campaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Campaign) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetStartTime

`func (o *Campaign) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Campaign) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *Campaign) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *Campaign) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *Campaign) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Campaign) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *Campaign) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *Campaign) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetAttributes

`func (o *Campaign) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Campaign) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Campaign) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Campaign) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetState

`func (o *Campaign) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Campaign) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Campaign) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Campaign) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetActiveRulesetId

`func (o *Campaign) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *Campaign) GetActiveRulesetIdOk() (int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRulesetId

`func (o *Campaign) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### SetActiveRulesetId

`func (o *Campaign) SetActiveRulesetId(v int32)`

SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.

### GetTags

`func (o *Campaign) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Campaign) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Campaign) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Campaign) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetFeatures

`func (o *Campaign) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Campaign) GetFeaturesOk() ([]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *Campaign) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *Campaign) SetFeatures(v []string)`

SetFeatures gets a reference to the given []string and assigns it to the Features field.

### GetCouponSettings

`func (o *Campaign) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *Campaign) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *Campaign) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *Campaign) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetReferralSettings

`func (o *Campaign) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *Campaign) GetReferralSettingsOk() (CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralSettings

`func (o *Campaign) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### SetReferralSettings

`func (o *Campaign) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.

### GetLimits

`func (o *Campaign) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Campaign) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *Campaign) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *Campaign) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetCampaignGroups

`func (o *Campaign) GetCampaignGroups() []int32`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *Campaign) GetCampaignGroupsOk() ([]int32, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignGroups

`func (o *Campaign) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### SetCampaignGroups

`func (o *Campaign) SetCampaignGroups(v []int32)`

SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.

### GetType

`func (o *Campaign) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Campaign) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Campaign) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Campaign) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetLinkedStoreIds

`func (o *Campaign) GetLinkedStoreIds() []int32`

GetLinkedStoreIds returns the LinkedStoreIds field if non-nil, zero value otherwise.

### GetLinkedStoreIdsOk

`func (o *Campaign) GetLinkedStoreIdsOk() ([]int32, bool)`

GetLinkedStoreIdsOk returns a tuple with the LinkedStoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkedStoreIds

`func (o *Campaign) HasLinkedStoreIds() bool`

HasLinkedStoreIds returns a boolean if a field has been set.

### SetLinkedStoreIds

`func (o *Campaign) SetLinkedStoreIds(v []int32)`

SetLinkedStoreIds gets a reference to the given []int32 and assigns it to the LinkedStoreIds field.

### GetBudgets

`func (o *Campaign) GetBudgets() []CampaignBudget`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *Campaign) GetBudgetsOk() ([]CampaignBudget, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBudgets

`func (o *Campaign) HasBudgets() bool`

HasBudgets returns a boolean if a field has been set.

### SetBudgets

`func (o *Campaign) SetBudgets(v []CampaignBudget)`

SetBudgets gets a reference to the given []CampaignBudget and assigns it to the Budgets field.

### GetCouponRedemptionCount

`func (o *Campaign) GetCouponRedemptionCount() int32`

GetCouponRedemptionCount returns the CouponRedemptionCount field if non-nil, zero value otherwise.

### GetCouponRedemptionCountOk

`func (o *Campaign) GetCouponRedemptionCountOk() (int32, bool)`

GetCouponRedemptionCountOk returns a tuple with the CouponRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRedemptionCount

`func (o *Campaign) HasCouponRedemptionCount() bool`

HasCouponRedemptionCount returns a boolean if a field has been set.

### SetCouponRedemptionCount

`func (o *Campaign) SetCouponRedemptionCount(v int32)`

SetCouponRedemptionCount gets a reference to the given int32 and assigns it to the CouponRedemptionCount field.

### GetReferralRedemptionCount

`func (o *Campaign) GetReferralRedemptionCount() int32`

GetReferralRedemptionCount returns the ReferralRedemptionCount field if non-nil, zero value otherwise.

### GetReferralRedemptionCountOk

`func (o *Campaign) GetReferralRedemptionCountOk() (int32, bool)`

GetReferralRedemptionCountOk returns a tuple with the ReferralRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralRedemptionCount

`func (o *Campaign) HasReferralRedemptionCount() bool`

HasReferralRedemptionCount returns a boolean if a field has been set.

### SetReferralRedemptionCount

`func (o *Campaign) SetReferralRedemptionCount(v int32)`

SetReferralRedemptionCount gets a reference to the given int32 and assigns it to the ReferralRedemptionCount field.

### GetDiscountCount

`func (o *Campaign) GetDiscountCount() float32`

GetDiscountCount returns the DiscountCount field if non-nil, zero value otherwise.

### GetDiscountCountOk

`func (o *Campaign) GetDiscountCountOk() (float32, bool)`

GetDiscountCountOk returns a tuple with the DiscountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountCount

`func (o *Campaign) HasDiscountCount() bool`

HasDiscountCount returns a boolean if a field has been set.

### SetDiscountCount

`func (o *Campaign) SetDiscountCount(v float32)`

SetDiscountCount gets a reference to the given float32 and assigns it to the DiscountCount field.

### GetDiscountEffectCount

`func (o *Campaign) GetDiscountEffectCount() int32`

GetDiscountEffectCount returns the DiscountEffectCount field if non-nil, zero value otherwise.

### GetDiscountEffectCountOk

`func (o *Campaign) GetDiscountEffectCountOk() (int32, bool)`

GetDiscountEffectCountOk returns a tuple with the DiscountEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountEffectCount

`func (o *Campaign) HasDiscountEffectCount() bool`

HasDiscountEffectCount returns a boolean if a field has been set.

### SetDiscountEffectCount

`func (o *Campaign) SetDiscountEffectCount(v int32)`

SetDiscountEffectCount gets a reference to the given int32 and assigns it to the DiscountEffectCount field.

### GetCouponCreationCount

`func (o *Campaign) GetCouponCreationCount() int32`

GetCouponCreationCount returns the CouponCreationCount field if non-nil, zero value otherwise.

### GetCouponCreationCountOk

`func (o *Campaign) GetCouponCreationCountOk() (int32, bool)`

GetCouponCreationCountOk returns a tuple with the CouponCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponCreationCount

`func (o *Campaign) HasCouponCreationCount() bool`

HasCouponCreationCount returns a boolean if a field has been set.

### SetCouponCreationCount

`func (o *Campaign) SetCouponCreationCount(v int32)`

SetCouponCreationCount gets a reference to the given int32 and assigns it to the CouponCreationCount field.

### GetCustomEffectCount

`func (o *Campaign) GetCustomEffectCount() int32`

GetCustomEffectCount returns the CustomEffectCount field if non-nil, zero value otherwise.

### GetCustomEffectCountOk

`func (o *Campaign) GetCustomEffectCountOk() (int32, bool)`

GetCustomEffectCountOk returns a tuple with the CustomEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomEffectCount

`func (o *Campaign) HasCustomEffectCount() bool`

HasCustomEffectCount returns a boolean if a field has been set.

### SetCustomEffectCount

`func (o *Campaign) SetCustomEffectCount(v int32)`

SetCustomEffectCount gets a reference to the given int32 and assigns it to the CustomEffectCount field.

### GetReferralCreationCount

`func (o *Campaign) GetReferralCreationCount() int32`

GetReferralCreationCount returns the ReferralCreationCount field if non-nil, zero value otherwise.

### GetReferralCreationCountOk

`func (o *Campaign) GetReferralCreationCountOk() (int32, bool)`

GetReferralCreationCountOk returns a tuple with the ReferralCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralCreationCount

`func (o *Campaign) HasReferralCreationCount() bool`

HasReferralCreationCount returns a boolean if a field has been set.

### SetReferralCreationCount

`func (o *Campaign) SetReferralCreationCount(v int32)`

SetReferralCreationCount gets a reference to the given int32 and assigns it to the ReferralCreationCount field.

### GetAddFreeItemEffectCount

`func (o *Campaign) GetAddFreeItemEffectCount() int32`

GetAddFreeItemEffectCount returns the AddFreeItemEffectCount field if non-nil, zero value otherwise.

### GetAddFreeItemEffectCountOk

`func (o *Campaign) GetAddFreeItemEffectCountOk() (int32, bool)`

GetAddFreeItemEffectCountOk returns a tuple with the AddFreeItemEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddFreeItemEffectCount

`func (o *Campaign) HasAddFreeItemEffectCount() bool`

HasAddFreeItemEffectCount returns a boolean if a field has been set.

### SetAddFreeItemEffectCount

`func (o *Campaign) SetAddFreeItemEffectCount(v int32)`

SetAddFreeItemEffectCount gets a reference to the given int32 and assigns it to the AddFreeItemEffectCount field.

### GetAwardedGiveawaysCount

`func (o *Campaign) GetAwardedGiveawaysCount() int32`

GetAwardedGiveawaysCount returns the AwardedGiveawaysCount field if non-nil, zero value otherwise.

### GetAwardedGiveawaysCountOk

`func (o *Campaign) GetAwardedGiveawaysCountOk() (int32, bool)`

GetAwardedGiveawaysCountOk returns a tuple with the AwardedGiveawaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAwardedGiveawaysCount

`func (o *Campaign) HasAwardedGiveawaysCount() bool`

HasAwardedGiveawaysCount returns a boolean if a field has been set.

### SetAwardedGiveawaysCount

`func (o *Campaign) SetAwardedGiveawaysCount(v int32)`

SetAwardedGiveawaysCount gets a reference to the given int32 and assigns it to the AwardedGiveawaysCount field.

### GetCreatedLoyaltyPointsCount

`func (o *Campaign) GetCreatedLoyaltyPointsCount() float32`

GetCreatedLoyaltyPointsCount returns the CreatedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsCountOk

`func (o *Campaign) GetCreatedLoyaltyPointsCountOk() (float32, bool)`

GetCreatedLoyaltyPointsCountOk returns a tuple with the CreatedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedLoyaltyPointsCount

`func (o *Campaign) HasCreatedLoyaltyPointsCount() bool`

HasCreatedLoyaltyPointsCount returns a boolean if a field has been set.

### SetCreatedLoyaltyPointsCount

`func (o *Campaign) SetCreatedLoyaltyPointsCount(v float32)`

SetCreatedLoyaltyPointsCount gets a reference to the given float32 and assigns it to the CreatedLoyaltyPointsCount field.

### GetCreatedLoyaltyPointsEffectCount

`func (o *Campaign) GetCreatedLoyaltyPointsEffectCount() int32`

GetCreatedLoyaltyPointsEffectCount returns the CreatedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsEffectCountOk

`func (o *Campaign) GetCreatedLoyaltyPointsEffectCountOk() (int32, bool)`

GetCreatedLoyaltyPointsEffectCountOk returns a tuple with the CreatedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedLoyaltyPointsEffectCount

`func (o *Campaign) HasCreatedLoyaltyPointsEffectCount() bool`

HasCreatedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### SetCreatedLoyaltyPointsEffectCount

`func (o *Campaign) SetCreatedLoyaltyPointsEffectCount(v int32)`

SetCreatedLoyaltyPointsEffectCount gets a reference to the given int32 and assigns it to the CreatedLoyaltyPointsEffectCount field.

### GetRedeemedLoyaltyPointsCount

`func (o *Campaign) GetRedeemedLoyaltyPointsCount() float32`

GetRedeemedLoyaltyPointsCount returns the RedeemedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsCountOk

`func (o *Campaign) GetRedeemedLoyaltyPointsCountOk() (float32, bool)`

GetRedeemedLoyaltyPointsCountOk returns a tuple with the RedeemedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRedeemedLoyaltyPointsCount

`func (o *Campaign) HasRedeemedLoyaltyPointsCount() bool`

HasRedeemedLoyaltyPointsCount returns a boolean if a field has been set.

### SetRedeemedLoyaltyPointsCount

`func (o *Campaign) SetRedeemedLoyaltyPointsCount(v float32)`

SetRedeemedLoyaltyPointsCount gets a reference to the given float32 and assigns it to the RedeemedLoyaltyPointsCount field.

### GetRedeemedLoyaltyPointsEffectCount

`func (o *Campaign) GetRedeemedLoyaltyPointsEffectCount() int32`

GetRedeemedLoyaltyPointsEffectCount returns the RedeemedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsEffectCountOk

`func (o *Campaign) GetRedeemedLoyaltyPointsEffectCountOk() (int32, bool)`

GetRedeemedLoyaltyPointsEffectCountOk returns a tuple with the RedeemedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRedeemedLoyaltyPointsEffectCount

`func (o *Campaign) HasRedeemedLoyaltyPointsEffectCount() bool`

HasRedeemedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### SetRedeemedLoyaltyPointsEffectCount

`func (o *Campaign) SetRedeemedLoyaltyPointsEffectCount(v int32)`

SetRedeemedLoyaltyPointsEffectCount gets a reference to the given int32 and assigns it to the RedeemedLoyaltyPointsEffectCount field.

### GetCallApiEffectCount

`func (o *Campaign) GetCallApiEffectCount() int32`

GetCallApiEffectCount returns the CallApiEffectCount field if non-nil, zero value otherwise.

### GetCallApiEffectCountOk

`func (o *Campaign) GetCallApiEffectCountOk() (int32, bool)`

GetCallApiEffectCountOk returns a tuple with the CallApiEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallApiEffectCount

`func (o *Campaign) HasCallApiEffectCount() bool`

HasCallApiEffectCount returns a boolean if a field has been set.

### SetCallApiEffectCount

`func (o *Campaign) SetCallApiEffectCount(v int32)`

SetCallApiEffectCount gets a reference to the given int32 and assigns it to the CallApiEffectCount field.

### GetReservecouponEffectCount

`func (o *Campaign) GetReservecouponEffectCount() int32`

GetReservecouponEffectCount returns the ReservecouponEffectCount field if non-nil, zero value otherwise.

### GetReservecouponEffectCountOk

`func (o *Campaign) GetReservecouponEffectCountOk() (int32, bool)`

GetReservecouponEffectCountOk returns a tuple with the ReservecouponEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservecouponEffectCount

`func (o *Campaign) HasReservecouponEffectCount() bool`

HasReservecouponEffectCount returns a boolean if a field has been set.

### SetReservecouponEffectCount

`func (o *Campaign) SetReservecouponEffectCount(v int32)`

SetReservecouponEffectCount gets a reference to the given int32 and assigns it to the ReservecouponEffectCount field.

### GetLastActivity

`func (o *Campaign) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *Campaign) GetLastActivityOk() (time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActivity

`func (o *Campaign) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivity

`func (o *Campaign) SetLastActivity(v time.Time)`

SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.

### GetUpdated

`func (o *Campaign) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Campaign) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *Campaign) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *Campaign) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.

### GetCreatedBy

`func (o *Campaign) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Campaign) GetCreatedByOk() (string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *Campaign) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *Campaign) SetCreatedBy(v string)`

SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.

### GetUpdatedBy

`func (o *Campaign) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Campaign) GetUpdatedByOk() (string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedBy

`func (o *Campaign) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedBy

`func (o *Campaign) SetUpdatedBy(v string)`

SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.

### GetTemplateId

`func (o *Campaign) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Campaign) GetTemplateIdOk() (int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateId

`func (o *Campaign) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateId

`func (o *Campaign) SetTemplateId(v int32)`

SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.

### GetFrontendState

`func (o *Campaign) GetFrontendState() string`

GetFrontendState returns the FrontendState field if non-nil, zero value otherwise.

### GetFrontendStateOk

`func (o *Campaign) GetFrontendStateOk() (string, bool)`

GetFrontendStateOk returns a tuple with the FrontendState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrontendState

`func (o *Campaign) HasFrontendState() bool`

HasFrontendState returns a boolean if a field has been set.

### SetFrontendState

`func (o *Campaign) SetFrontendState(v string)`

SetFrontendState gets a reference to the given string and assigns it to the FrontendState field.

### GetStoresImported

`func (o *Campaign) GetStoresImported() bool`

GetStoresImported returns the StoresImported field if non-nil, zero value otherwise.

### GetStoresImportedOk

`func (o *Campaign) GetStoresImportedOk() (bool, bool)`

GetStoresImportedOk returns a tuple with the StoresImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStoresImported

`func (o *Campaign) HasStoresImported() bool`

HasStoresImported returns a boolean if a field has been set.

### SetStoresImported

`func (o *Campaign) SetStoresImported(v bool)`

SetStoresImported gets a reference to the given bool and assigns it to the StoresImported field.

### GetActiveRevisionId

`func (o *Campaign) GetActiveRevisionId() int32`

GetActiveRevisionId returns the ActiveRevisionId field if non-nil, zero value otherwise.

### GetActiveRevisionIdOk

`func (o *Campaign) GetActiveRevisionIdOk() (int32, bool)`

GetActiveRevisionIdOk returns a tuple with the ActiveRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRevisionId

`func (o *Campaign) HasActiveRevisionId() bool`

HasActiveRevisionId returns a boolean if a field has been set.

### SetActiveRevisionId

`func (o *Campaign) SetActiveRevisionId(v int32)`

SetActiveRevisionId gets a reference to the given int32 and assigns it to the ActiveRevisionId field.

### GetActiveRevisionVersionId

`func (o *Campaign) GetActiveRevisionVersionId() int32`

GetActiveRevisionVersionId returns the ActiveRevisionVersionId field if non-nil, zero value otherwise.

### GetActiveRevisionVersionIdOk

`func (o *Campaign) GetActiveRevisionVersionIdOk() (int32, bool)`

GetActiveRevisionVersionIdOk returns a tuple with the ActiveRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRevisionVersionId

`func (o *Campaign) HasActiveRevisionVersionId() bool`

HasActiveRevisionVersionId returns a boolean if a field has been set.

### SetActiveRevisionVersionId

`func (o *Campaign) SetActiveRevisionVersionId(v int32)`

SetActiveRevisionVersionId gets a reference to the given int32 and assigns it to the ActiveRevisionVersionId field.

### GetVersion

`func (o *Campaign) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Campaign) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *Campaign) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *Campaign) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetCurrentRevisionId

`func (o *Campaign) GetCurrentRevisionId() int32`

GetCurrentRevisionId returns the CurrentRevisionId field if non-nil, zero value otherwise.

### GetCurrentRevisionIdOk

`func (o *Campaign) GetCurrentRevisionIdOk() (int32, bool)`

GetCurrentRevisionIdOk returns a tuple with the CurrentRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentRevisionId

`func (o *Campaign) HasCurrentRevisionId() bool`

HasCurrentRevisionId returns a boolean if a field has been set.

### SetCurrentRevisionId

`func (o *Campaign) SetCurrentRevisionId(v int32)`

SetCurrentRevisionId gets a reference to the given int32 and assigns it to the CurrentRevisionId field.

### GetCurrentRevisionVersionId

`func (o *Campaign) GetCurrentRevisionVersionId() int32`

GetCurrentRevisionVersionId returns the CurrentRevisionVersionId field if non-nil, zero value otherwise.

### GetCurrentRevisionVersionIdOk

`func (o *Campaign) GetCurrentRevisionVersionIdOk() (int32, bool)`

GetCurrentRevisionVersionIdOk returns a tuple with the CurrentRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentRevisionVersionId

`func (o *Campaign) HasCurrentRevisionVersionId() bool`

HasCurrentRevisionVersionId returns a boolean if a field has been set.

### SetCurrentRevisionVersionId

`func (o *Campaign) SetCurrentRevisionVersionId(v int32)`

SetCurrentRevisionVersionId gets a reference to the given int32 and assigns it to the CurrentRevisionVersionId field.

### GetStageRevision

`func (o *Campaign) GetStageRevision() bool`

GetStageRevision returns the StageRevision field if non-nil, zero value otherwise.

### GetStageRevisionOk

`func (o *Campaign) GetStageRevisionOk() (bool, bool)`

GetStageRevisionOk returns a tuple with the StageRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStageRevision

`func (o *Campaign) HasStageRevision() bool`

HasStageRevision returns a boolean if a field has been set.

### SetStageRevision

`func (o *Campaign) SetStageRevision(v bool)`

SetStageRevision gets a reference to the given bool and assigns it to the StageRevision field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


