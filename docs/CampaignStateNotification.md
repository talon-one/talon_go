# CampaignStateNotification

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
**EvaluationGroupId** | Pointer to **int32** | The ID of the campaign evaluation group the campaign belongs to. | [optional] 
**Type** | Pointer to **string** | The campaign type. Possible type values:   - &#x60;cartItem&#x60;: Type of campaign that can apply effects only to cart items.   - &#x60;advanced&#x60;: Type of campaign that can apply effects to customer sessions and cart items.  | [default to TYPE_ADVANCED]
**LinkedStoreIds** | Pointer to **[]int32** | A list of store IDs that you want to link to the campaign.  **Note:** Campaigns with linked store IDs will only be evaluated when there is a [customer session update](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) that references a linked store.  | [optional] 
**CouponRedemptionCount** | Pointer to **int32** | Number of coupons redeemed in the campaign. | [optional] 
**ReferralRedemptionCount** | Pointer to **int32** | Number of referral codes redeemed in the campaign. | [optional] 
**DiscountCount** | Pointer to **float32** | Total amount of discounts redeemed in the campaign. | [optional] 
**DiscountEffectCount** | Pointer to **int32** | Total number of times discounts were redeemed in this campaign. | [optional] 
**CouponCreationCount** | Pointer to **int32** | Total number of coupons created by rules in this campaign. | [optional] 
**CustomEffectCount** | Pointer to **int32** | Total number of custom effects triggered by rules in this campaign. | [optional] 
**ReferralCreationCount** | Pointer to **int32** | Total number of referrals created by rules in this campaign. | [optional] 
**AddFreeItemEffectCount** | Pointer to **int32** | Total number of times the [add free item effect](https://docs.talon.one/docs/dev/integration-api/api-effects#addfreeitem) can be triggered in this campaign. | [optional] 
**AwardedGiveawaysCount** | Pointer to **int32** | Total number of giveaways awarded by rules in this campaign. | [optional] 
**CreatedLoyaltyPointsCount** | Pointer to **float32** | Total number of loyalty points created by rules in this campaign. | [optional] 
**CreatedLoyaltyPointsEffectCount** | Pointer to **int32** | Total number of loyalty point creation effects triggered by rules in this campaign. | [optional] 
**RedeemedLoyaltyPointsCount** | Pointer to **float32** | Total number of loyalty points redeemed by rules in this campaign. | [optional] 
**RedeemedLoyaltyPointsEffectCount** | Pointer to **int32** | Total number of loyalty point redemption effects triggered by rules in this campaign. | [optional] 
**CallApiEffectCount** | Pointer to **int32** | Total number of webhooks triggered by rules in this campaign. | [optional] 
**ReservecouponEffectCount** | Pointer to **int32** | Total number of reserve coupon effects triggered by rules in this campaign. | [optional] 
**LastActivity** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received by this campaign. | [optional] 
**WeeklyUsageCount** | Pointer to **int32** | The total number of effects created in this campaign in the past 7 days. | [optional] 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the campaign&#39;s property. Updates to external entities used in this campaign are **not** registered by this property, such as collection or coupon updates.  | [optional] 
**CreatedBy** | Pointer to **string** | Name of the user who created this campaign if available. | [optional] 
**UpdatedBy** | Pointer to **string** | Name of the user who last updated this campaign if available. | [optional] 
**TemplateId** | Pointer to **int32** | The ID of the Campaign Template this Campaign was created from. | [optional] 
**FrontendState** | Pointer to **string** | A campaign state described exactly as in the Campaign Manager. | 

## Methods

### GetId

`func (o *CampaignStateNotification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignStateNotification) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CampaignStateNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CampaignStateNotification) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CampaignStateNotification) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignStateNotification) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CampaignStateNotification) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CampaignStateNotification) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *CampaignStateNotification) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignStateNotification) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CampaignStateNotification) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CampaignStateNotification) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetUserId

`func (o *CampaignStateNotification) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CampaignStateNotification) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *CampaignStateNotification) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *CampaignStateNotification) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.

### GetName

`func (o *CampaignStateNotification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignStateNotification) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CampaignStateNotification) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CampaignStateNotification) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *CampaignStateNotification) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignStateNotification) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CampaignStateNotification) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CampaignStateNotification) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetStartTime

`func (o *CampaignStateNotification) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CampaignStateNotification) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *CampaignStateNotification) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *CampaignStateNotification) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *CampaignStateNotification) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CampaignStateNotification) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *CampaignStateNotification) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *CampaignStateNotification) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetAttributes

`func (o *CampaignStateNotification) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignStateNotification) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CampaignStateNotification) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CampaignStateNotification) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetState

`func (o *CampaignStateNotification) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CampaignStateNotification) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *CampaignStateNotification) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *CampaignStateNotification) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetActiveRulesetId

`func (o *CampaignStateNotification) GetActiveRulesetId() int32`

GetActiveRulesetId returns the ActiveRulesetId field if non-nil, zero value otherwise.

### GetActiveRulesetIdOk

`func (o *CampaignStateNotification) GetActiveRulesetIdOk() (int32, bool)`

GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveRulesetId

`func (o *CampaignStateNotification) HasActiveRulesetId() bool`

HasActiveRulesetId returns a boolean if a field has been set.

### SetActiveRulesetId

`func (o *CampaignStateNotification) SetActiveRulesetId(v int32)`

SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.

### GetTags

`func (o *CampaignStateNotification) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CampaignStateNotification) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *CampaignStateNotification) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *CampaignStateNotification) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetFeatures

`func (o *CampaignStateNotification) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CampaignStateNotification) GetFeaturesOk() ([]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatures

`func (o *CampaignStateNotification) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeatures

`func (o *CampaignStateNotification) SetFeatures(v []string)`

SetFeatures gets a reference to the given []string and assigns it to the Features field.

### GetCouponSettings

`func (o *CampaignStateNotification) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *CampaignStateNotification) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *CampaignStateNotification) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *CampaignStateNotification) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetReferralSettings

`func (o *CampaignStateNotification) GetReferralSettings() CodeGeneratorSettings`

GetReferralSettings returns the ReferralSettings field if non-nil, zero value otherwise.

### GetReferralSettingsOk

`func (o *CampaignStateNotification) GetReferralSettingsOk() (CodeGeneratorSettings, bool)`

GetReferralSettingsOk returns a tuple with the ReferralSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralSettings

`func (o *CampaignStateNotification) HasReferralSettings() bool`

HasReferralSettings returns a boolean if a field has been set.

### SetReferralSettings

`func (o *CampaignStateNotification) SetReferralSettings(v CodeGeneratorSettings)`

SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.

### GetLimits

`func (o *CampaignStateNotification) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CampaignStateNotification) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *CampaignStateNotification) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *CampaignStateNotification) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetCampaignGroups

`func (o *CampaignStateNotification) GetCampaignGroups() []int32`

GetCampaignGroups returns the CampaignGroups field if non-nil, zero value otherwise.

### GetCampaignGroupsOk

`func (o *CampaignStateNotification) GetCampaignGroupsOk() ([]int32, bool)`

GetCampaignGroupsOk returns a tuple with the CampaignGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignGroups

`func (o *CampaignStateNotification) HasCampaignGroups() bool`

HasCampaignGroups returns a boolean if a field has been set.

### SetCampaignGroups

`func (o *CampaignStateNotification) SetCampaignGroups(v []int32)`

SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.

### GetEvaluationGroupId

`func (o *CampaignStateNotification) GetEvaluationGroupId() int32`

GetEvaluationGroupId returns the EvaluationGroupId field if non-nil, zero value otherwise.

### GetEvaluationGroupIdOk

`func (o *CampaignStateNotification) GetEvaluationGroupIdOk() (int32, bool)`

GetEvaluationGroupIdOk returns a tuple with the EvaluationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationGroupId

`func (o *CampaignStateNotification) HasEvaluationGroupId() bool`

HasEvaluationGroupId returns a boolean if a field has been set.

### SetEvaluationGroupId

`func (o *CampaignStateNotification) SetEvaluationGroupId(v int32)`

SetEvaluationGroupId gets a reference to the given int32 and assigns it to the EvaluationGroupId field.

### GetType

`func (o *CampaignStateNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignStateNotification) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *CampaignStateNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *CampaignStateNotification) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetLinkedStoreIds

`func (o *CampaignStateNotification) GetLinkedStoreIds() []int32`

GetLinkedStoreIds returns the LinkedStoreIds field if non-nil, zero value otherwise.

### GetLinkedStoreIdsOk

`func (o *CampaignStateNotification) GetLinkedStoreIdsOk() ([]int32, bool)`

GetLinkedStoreIdsOk returns a tuple with the LinkedStoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkedStoreIds

`func (o *CampaignStateNotification) HasLinkedStoreIds() bool`

HasLinkedStoreIds returns a boolean if a field has been set.

### SetLinkedStoreIds

`func (o *CampaignStateNotification) SetLinkedStoreIds(v []int32)`

SetLinkedStoreIds gets a reference to the given []int32 and assigns it to the LinkedStoreIds field.

### GetCouponRedemptionCount

`func (o *CampaignStateNotification) GetCouponRedemptionCount() int32`

GetCouponRedemptionCount returns the CouponRedemptionCount field if non-nil, zero value otherwise.

### GetCouponRedemptionCountOk

`func (o *CampaignStateNotification) GetCouponRedemptionCountOk() (int32, bool)`

GetCouponRedemptionCountOk returns a tuple with the CouponRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRedemptionCount

`func (o *CampaignStateNotification) HasCouponRedemptionCount() bool`

HasCouponRedemptionCount returns a boolean if a field has been set.

### SetCouponRedemptionCount

`func (o *CampaignStateNotification) SetCouponRedemptionCount(v int32)`

SetCouponRedemptionCount gets a reference to the given int32 and assigns it to the CouponRedemptionCount field.

### GetReferralRedemptionCount

`func (o *CampaignStateNotification) GetReferralRedemptionCount() int32`

GetReferralRedemptionCount returns the ReferralRedemptionCount field if non-nil, zero value otherwise.

### GetReferralRedemptionCountOk

`func (o *CampaignStateNotification) GetReferralRedemptionCountOk() (int32, bool)`

GetReferralRedemptionCountOk returns a tuple with the ReferralRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralRedemptionCount

`func (o *CampaignStateNotification) HasReferralRedemptionCount() bool`

HasReferralRedemptionCount returns a boolean if a field has been set.

### SetReferralRedemptionCount

`func (o *CampaignStateNotification) SetReferralRedemptionCount(v int32)`

SetReferralRedemptionCount gets a reference to the given int32 and assigns it to the ReferralRedemptionCount field.

### GetDiscountCount

`func (o *CampaignStateNotification) GetDiscountCount() float32`

GetDiscountCount returns the DiscountCount field if non-nil, zero value otherwise.

### GetDiscountCountOk

`func (o *CampaignStateNotification) GetDiscountCountOk() (float32, bool)`

GetDiscountCountOk returns a tuple with the DiscountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountCount

`func (o *CampaignStateNotification) HasDiscountCount() bool`

HasDiscountCount returns a boolean if a field has been set.

### SetDiscountCount

`func (o *CampaignStateNotification) SetDiscountCount(v float32)`

SetDiscountCount gets a reference to the given float32 and assigns it to the DiscountCount field.

### GetDiscountEffectCount

`func (o *CampaignStateNotification) GetDiscountEffectCount() int32`

GetDiscountEffectCount returns the DiscountEffectCount field if non-nil, zero value otherwise.

### GetDiscountEffectCountOk

`func (o *CampaignStateNotification) GetDiscountEffectCountOk() (int32, bool)`

GetDiscountEffectCountOk returns a tuple with the DiscountEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountEffectCount

`func (o *CampaignStateNotification) HasDiscountEffectCount() bool`

HasDiscountEffectCount returns a boolean if a field has been set.

### SetDiscountEffectCount

`func (o *CampaignStateNotification) SetDiscountEffectCount(v int32)`

SetDiscountEffectCount gets a reference to the given int32 and assigns it to the DiscountEffectCount field.

### GetCouponCreationCount

`func (o *CampaignStateNotification) GetCouponCreationCount() int32`

GetCouponCreationCount returns the CouponCreationCount field if non-nil, zero value otherwise.

### GetCouponCreationCountOk

`func (o *CampaignStateNotification) GetCouponCreationCountOk() (int32, bool)`

GetCouponCreationCountOk returns a tuple with the CouponCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponCreationCount

`func (o *CampaignStateNotification) HasCouponCreationCount() bool`

HasCouponCreationCount returns a boolean if a field has been set.

### SetCouponCreationCount

`func (o *CampaignStateNotification) SetCouponCreationCount(v int32)`

SetCouponCreationCount gets a reference to the given int32 and assigns it to the CouponCreationCount field.

### GetCustomEffectCount

`func (o *CampaignStateNotification) GetCustomEffectCount() int32`

GetCustomEffectCount returns the CustomEffectCount field if non-nil, zero value otherwise.

### GetCustomEffectCountOk

`func (o *CampaignStateNotification) GetCustomEffectCountOk() (int32, bool)`

GetCustomEffectCountOk returns a tuple with the CustomEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomEffectCount

`func (o *CampaignStateNotification) HasCustomEffectCount() bool`

HasCustomEffectCount returns a boolean if a field has been set.

### SetCustomEffectCount

`func (o *CampaignStateNotification) SetCustomEffectCount(v int32)`

SetCustomEffectCount gets a reference to the given int32 and assigns it to the CustomEffectCount field.

### GetReferralCreationCount

`func (o *CampaignStateNotification) GetReferralCreationCount() int32`

GetReferralCreationCount returns the ReferralCreationCount field if non-nil, zero value otherwise.

### GetReferralCreationCountOk

`func (o *CampaignStateNotification) GetReferralCreationCountOk() (int32, bool)`

GetReferralCreationCountOk returns a tuple with the ReferralCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralCreationCount

`func (o *CampaignStateNotification) HasReferralCreationCount() bool`

HasReferralCreationCount returns a boolean if a field has been set.

### SetReferralCreationCount

`func (o *CampaignStateNotification) SetReferralCreationCount(v int32)`

SetReferralCreationCount gets a reference to the given int32 and assigns it to the ReferralCreationCount field.

### GetAddFreeItemEffectCount

`func (o *CampaignStateNotification) GetAddFreeItemEffectCount() int32`

GetAddFreeItemEffectCount returns the AddFreeItemEffectCount field if non-nil, zero value otherwise.

### GetAddFreeItemEffectCountOk

`func (o *CampaignStateNotification) GetAddFreeItemEffectCountOk() (int32, bool)`

GetAddFreeItemEffectCountOk returns a tuple with the AddFreeItemEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddFreeItemEffectCount

`func (o *CampaignStateNotification) HasAddFreeItemEffectCount() bool`

HasAddFreeItemEffectCount returns a boolean if a field has been set.

### SetAddFreeItemEffectCount

`func (o *CampaignStateNotification) SetAddFreeItemEffectCount(v int32)`

SetAddFreeItemEffectCount gets a reference to the given int32 and assigns it to the AddFreeItemEffectCount field.

### GetAwardedGiveawaysCount

`func (o *CampaignStateNotification) GetAwardedGiveawaysCount() int32`

GetAwardedGiveawaysCount returns the AwardedGiveawaysCount field if non-nil, zero value otherwise.

### GetAwardedGiveawaysCountOk

`func (o *CampaignStateNotification) GetAwardedGiveawaysCountOk() (int32, bool)`

GetAwardedGiveawaysCountOk returns a tuple with the AwardedGiveawaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAwardedGiveawaysCount

`func (o *CampaignStateNotification) HasAwardedGiveawaysCount() bool`

HasAwardedGiveawaysCount returns a boolean if a field has been set.

### SetAwardedGiveawaysCount

`func (o *CampaignStateNotification) SetAwardedGiveawaysCount(v int32)`

SetAwardedGiveawaysCount gets a reference to the given int32 and assigns it to the AwardedGiveawaysCount field.

### GetCreatedLoyaltyPointsCount

`func (o *CampaignStateNotification) GetCreatedLoyaltyPointsCount() float32`

GetCreatedLoyaltyPointsCount returns the CreatedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsCountOk

`func (o *CampaignStateNotification) GetCreatedLoyaltyPointsCountOk() (float32, bool)`

GetCreatedLoyaltyPointsCountOk returns a tuple with the CreatedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedLoyaltyPointsCount

`func (o *CampaignStateNotification) HasCreatedLoyaltyPointsCount() bool`

HasCreatedLoyaltyPointsCount returns a boolean if a field has been set.

### SetCreatedLoyaltyPointsCount

`func (o *CampaignStateNotification) SetCreatedLoyaltyPointsCount(v float32)`

SetCreatedLoyaltyPointsCount gets a reference to the given float32 and assigns it to the CreatedLoyaltyPointsCount field.

### GetCreatedLoyaltyPointsEffectCount

`func (o *CampaignStateNotification) GetCreatedLoyaltyPointsEffectCount() int32`

GetCreatedLoyaltyPointsEffectCount returns the CreatedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsEffectCountOk

`func (o *CampaignStateNotification) GetCreatedLoyaltyPointsEffectCountOk() (int32, bool)`

GetCreatedLoyaltyPointsEffectCountOk returns a tuple with the CreatedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedLoyaltyPointsEffectCount

`func (o *CampaignStateNotification) HasCreatedLoyaltyPointsEffectCount() bool`

HasCreatedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### SetCreatedLoyaltyPointsEffectCount

`func (o *CampaignStateNotification) SetCreatedLoyaltyPointsEffectCount(v int32)`

SetCreatedLoyaltyPointsEffectCount gets a reference to the given int32 and assigns it to the CreatedLoyaltyPointsEffectCount field.

### GetRedeemedLoyaltyPointsCount

`func (o *CampaignStateNotification) GetRedeemedLoyaltyPointsCount() float32`

GetRedeemedLoyaltyPointsCount returns the RedeemedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsCountOk

`func (o *CampaignStateNotification) GetRedeemedLoyaltyPointsCountOk() (float32, bool)`

GetRedeemedLoyaltyPointsCountOk returns a tuple with the RedeemedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRedeemedLoyaltyPointsCount

`func (o *CampaignStateNotification) HasRedeemedLoyaltyPointsCount() bool`

HasRedeemedLoyaltyPointsCount returns a boolean if a field has been set.

### SetRedeemedLoyaltyPointsCount

`func (o *CampaignStateNotification) SetRedeemedLoyaltyPointsCount(v float32)`

SetRedeemedLoyaltyPointsCount gets a reference to the given float32 and assigns it to the RedeemedLoyaltyPointsCount field.

### GetRedeemedLoyaltyPointsEffectCount

`func (o *CampaignStateNotification) GetRedeemedLoyaltyPointsEffectCount() int32`

GetRedeemedLoyaltyPointsEffectCount returns the RedeemedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsEffectCountOk

`func (o *CampaignStateNotification) GetRedeemedLoyaltyPointsEffectCountOk() (int32, bool)`

GetRedeemedLoyaltyPointsEffectCountOk returns a tuple with the RedeemedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRedeemedLoyaltyPointsEffectCount

`func (o *CampaignStateNotification) HasRedeemedLoyaltyPointsEffectCount() bool`

HasRedeemedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### SetRedeemedLoyaltyPointsEffectCount

`func (o *CampaignStateNotification) SetRedeemedLoyaltyPointsEffectCount(v int32)`

SetRedeemedLoyaltyPointsEffectCount gets a reference to the given int32 and assigns it to the RedeemedLoyaltyPointsEffectCount field.

### GetCallApiEffectCount

`func (o *CampaignStateNotification) GetCallApiEffectCount() int32`

GetCallApiEffectCount returns the CallApiEffectCount field if non-nil, zero value otherwise.

### GetCallApiEffectCountOk

`func (o *CampaignStateNotification) GetCallApiEffectCountOk() (int32, bool)`

GetCallApiEffectCountOk returns a tuple with the CallApiEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallApiEffectCount

`func (o *CampaignStateNotification) HasCallApiEffectCount() bool`

HasCallApiEffectCount returns a boolean if a field has been set.

### SetCallApiEffectCount

`func (o *CampaignStateNotification) SetCallApiEffectCount(v int32)`

SetCallApiEffectCount gets a reference to the given int32 and assigns it to the CallApiEffectCount field.

### GetReservecouponEffectCount

`func (o *CampaignStateNotification) GetReservecouponEffectCount() int32`

GetReservecouponEffectCount returns the ReservecouponEffectCount field if non-nil, zero value otherwise.

### GetReservecouponEffectCountOk

`func (o *CampaignStateNotification) GetReservecouponEffectCountOk() (int32, bool)`

GetReservecouponEffectCountOk returns a tuple with the ReservecouponEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservecouponEffectCount

`func (o *CampaignStateNotification) HasReservecouponEffectCount() bool`

HasReservecouponEffectCount returns a boolean if a field has been set.

### SetReservecouponEffectCount

`func (o *CampaignStateNotification) SetReservecouponEffectCount(v int32)`

SetReservecouponEffectCount gets a reference to the given int32 and assigns it to the ReservecouponEffectCount field.

### GetLastActivity

`func (o *CampaignStateNotification) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CampaignStateNotification) GetLastActivityOk() (time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActivity

`func (o *CampaignStateNotification) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivity

`func (o *CampaignStateNotification) SetLastActivity(v time.Time)`

SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.

### GetWeeklyUsageCount

`func (o *CampaignStateNotification) GetWeeklyUsageCount() int32`

GetWeeklyUsageCount returns the WeeklyUsageCount field if non-nil, zero value otherwise.

### GetWeeklyUsageCountOk

`func (o *CampaignStateNotification) GetWeeklyUsageCountOk() (int32, bool)`

GetWeeklyUsageCountOk returns a tuple with the WeeklyUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWeeklyUsageCount

`func (o *CampaignStateNotification) HasWeeklyUsageCount() bool`

HasWeeklyUsageCount returns a boolean if a field has been set.

### SetWeeklyUsageCount

`func (o *CampaignStateNotification) SetWeeklyUsageCount(v int32)`

SetWeeklyUsageCount gets a reference to the given int32 and assigns it to the WeeklyUsageCount field.

### GetUpdated

`func (o *CampaignStateNotification) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CampaignStateNotification) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *CampaignStateNotification) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *CampaignStateNotification) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.

### GetCreatedBy

`func (o *CampaignStateNotification) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CampaignStateNotification) GetCreatedByOk() (string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *CampaignStateNotification) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *CampaignStateNotification) SetCreatedBy(v string)`

SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.

### GetUpdatedBy

`func (o *CampaignStateNotification) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CampaignStateNotification) GetUpdatedByOk() (string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedBy

`func (o *CampaignStateNotification) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedBy

`func (o *CampaignStateNotification) SetUpdatedBy(v string)`

SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.

### GetTemplateId

`func (o *CampaignStateNotification) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CampaignStateNotification) GetTemplateIdOk() (int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateId

`func (o *CampaignStateNotification) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateId

`func (o *CampaignStateNotification) SetTemplateId(v int32)`

SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.

### GetFrontendState

`func (o *CampaignStateNotification) GetFrontendState() string`

GetFrontendState returns the FrontendState field if non-nil, zero value otherwise.

### GetFrontendStateOk

`func (o *CampaignStateNotification) GetFrontendStateOk() (string, bool)`

GetFrontendStateOk returns a tuple with the FrontendState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrontendState

`func (o *CampaignStateNotification) HasFrontendState() bool`

HasFrontendState returns a boolean if a field has been set.

### SetFrontendState

`func (o *CampaignStateNotification) SetFrontendState(v string)`

SetFrontendState gets a reference to the given string and assigns it to the FrontendState field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


