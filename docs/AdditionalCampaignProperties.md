# AdditionalCampaignProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### GetBudgets

`func (o *AdditionalCampaignProperties) GetBudgets() []CampaignBudget`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *AdditionalCampaignProperties) GetBudgetsOk() ([]CampaignBudget, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBudgets

`func (o *AdditionalCampaignProperties) HasBudgets() bool`

HasBudgets returns a boolean if a field has been set.

### SetBudgets

`func (o *AdditionalCampaignProperties) SetBudgets(v []CampaignBudget)`

SetBudgets gets a reference to the given []CampaignBudget and assigns it to the Budgets field.

### GetCouponRedemptionCount

`func (o *AdditionalCampaignProperties) GetCouponRedemptionCount() int32`

GetCouponRedemptionCount returns the CouponRedemptionCount field if non-nil, zero value otherwise.

### GetCouponRedemptionCountOk

`func (o *AdditionalCampaignProperties) GetCouponRedemptionCountOk() (int32, bool)`

GetCouponRedemptionCountOk returns a tuple with the CouponRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRedemptionCount

`func (o *AdditionalCampaignProperties) HasCouponRedemptionCount() bool`

HasCouponRedemptionCount returns a boolean if a field has been set.

### SetCouponRedemptionCount

`func (o *AdditionalCampaignProperties) SetCouponRedemptionCount(v int32)`

SetCouponRedemptionCount gets a reference to the given int32 and assigns it to the CouponRedemptionCount field.

### GetReferralRedemptionCount

`func (o *AdditionalCampaignProperties) GetReferralRedemptionCount() int32`

GetReferralRedemptionCount returns the ReferralRedemptionCount field if non-nil, zero value otherwise.

### GetReferralRedemptionCountOk

`func (o *AdditionalCampaignProperties) GetReferralRedemptionCountOk() (int32, bool)`

GetReferralRedemptionCountOk returns a tuple with the ReferralRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralRedemptionCount

`func (o *AdditionalCampaignProperties) HasReferralRedemptionCount() bool`

HasReferralRedemptionCount returns a boolean if a field has been set.

### SetReferralRedemptionCount

`func (o *AdditionalCampaignProperties) SetReferralRedemptionCount(v int32)`

SetReferralRedemptionCount gets a reference to the given int32 and assigns it to the ReferralRedemptionCount field.

### GetDiscountCount

`func (o *AdditionalCampaignProperties) GetDiscountCount() float32`

GetDiscountCount returns the DiscountCount field if non-nil, zero value otherwise.

### GetDiscountCountOk

`func (o *AdditionalCampaignProperties) GetDiscountCountOk() (float32, bool)`

GetDiscountCountOk returns a tuple with the DiscountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountCount

`func (o *AdditionalCampaignProperties) HasDiscountCount() bool`

HasDiscountCount returns a boolean if a field has been set.

### SetDiscountCount

`func (o *AdditionalCampaignProperties) SetDiscountCount(v float32)`

SetDiscountCount gets a reference to the given float32 and assigns it to the DiscountCount field.

### GetDiscountEffectCount

`func (o *AdditionalCampaignProperties) GetDiscountEffectCount() int32`

GetDiscountEffectCount returns the DiscountEffectCount field if non-nil, zero value otherwise.

### GetDiscountEffectCountOk

`func (o *AdditionalCampaignProperties) GetDiscountEffectCountOk() (int32, bool)`

GetDiscountEffectCountOk returns a tuple with the DiscountEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountEffectCount

`func (o *AdditionalCampaignProperties) HasDiscountEffectCount() bool`

HasDiscountEffectCount returns a boolean if a field has been set.

### SetDiscountEffectCount

`func (o *AdditionalCampaignProperties) SetDiscountEffectCount(v int32)`

SetDiscountEffectCount gets a reference to the given int32 and assigns it to the DiscountEffectCount field.

### GetCouponCreationCount

`func (o *AdditionalCampaignProperties) GetCouponCreationCount() int32`

GetCouponCreationCount returns the CouponCreationCount field if non-nil, zero value otherwise.

### GetCouponCreationCountOk

`func (o *AdditionalCampaignProperties) GetCouponCreationCountOk() (int32, bool)`

GetCouponCreationCountOk returns a tuple with the CouponCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponCreationCount

`func (o *AdditionalCampaignProperties) HasCouponCreationCount() bool`

HasCouponCreationCount returns a boolean if a field has been set.

### SetCouponCreationCount

`func (o *AdditionalCampaignProperties) SetCouponCreationCount(v int32)`

SetCouponCreationCount gets a reference to the given int32 and assigns it to the CouponCreationCount field.

### GetCustomEffectCount

`func (o *AdditionalCampaignProperties) GetCustomEffectCount() int32`

GetCustomEffectCount returns the CustomEffectCount field if non-nil, zero value otherwise.

### GetCustomEffectCountOk

`func (o *AdditionalCampaignProperties) GetCustomEffectCountOk() (int32, bool)`

GetCustomEffectCountOk returns a tuple with the CustomEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomEffectCount

`func (o *AdditionalCampaignProperties) HasCustomEffectCount() bool`

HasCustomEffectCount returns a boolean if a field has been set.

### SetCustomEffectCount

`func (o *AdditionalCampaignProperties) SetCustomEffectCount(v int32)`

SetCustomEffectCount gets a reference to the given int32 and assigns it to the CustomEffectCount field.

### GetReferralCreationCount

`func (o *AdditionalCampaignProperties) GetReferralCreationCount() int32`

GetReferralCreationCount returns the ReferralCreationCount field if non-nil, zero value otherwise.

### GetReferralCreationCountOk

`func (o *AdditionalCampaignProperties) GetReferralCreationCountOk() (int32, bool)`

GetReferralCreationCountOk returns a tuple with the ReferralCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralCreationCount

`func (o *AdditionalCampaignProperties) HasReferralCreationCount() bool`

HasReferralCreationCount returns a boolean if a field has been set.

### SetReferralCreationCount

`func (o *AdditionalCampaignProperties) SetReferralCreationCount(v int32)`

SetReferralCreationCount gets a reference to the given int32 and assigns it to the ReferralCreationCount field.

### GetAddFreeItemEffectCount

`func (o *AdditionalCampaignProperties) GetAddFreeItemEffectCount() int32`

GetAddFreeItemEffectCount returns the AddFreeItemEffectCount field if non-nil, zero value otherwise.

### GetAddFreeItemEffectCountOk

`func (o *AdditionalCampaignProperties) GetAddFreeItemEffectCountOk() (int32, bool)`

GetAddFreeItemEffectCountOk returns a tuple with the AddFreeItemEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddFreeItemEffectCount

`func (o *AdditionalCampaignProperties) HasAddFreeItemEffectCount() bool`

HasAddFreeItemEffectCount returns a boolean if a field has been set.

### SetAddFreeItemEffectCount

`func (o *AdditionalCampaignProperties) SetAddFreeItemEffectCount(v int32)`

SetAddFreeItemEffectCount gets a reference to the given int32 and assigns it to the AddFreeItemEffectCount field.

### GetAwardedGiveawaysCount

`func (o *AdditionalCampaignProperties) GetAwardedGiveawaysCount() int32`

GetAwardedGiveawaysCount returns the AwardedGiveawaysCount field if non-nil, zero value otherwise.

### GetAwardedGiveawaysCountOk

`func (o *AdditionalCampaignProperties) GetAwardedGiveawaysCountOk() (int32, bool)`

GetAwardedGiveawaysCountOk returns a tuple with the AwardedGiveawaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAwardedGiveawaysCount

`func (o *AdditionalCampaignProperties) HasAwardedGiveawaysCount() bool`

HasAwardedGiveawaysCount returns a boolean if a field has been set.

### SetAwardedGiveawaysCount

`func (o *AdditionalCampaignProperties) SetAwardedGiveawaysCount(v int32)`

SetAwardedGiveawaysCount gets a reference to the given int32 and assigns it to the AwardedGiveawaysCount field.

### GetCreatedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsCount() float32`

GetCreatedLoyaltyPointsCount returns the CreatedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsCountOk

`func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsCountOk() (float32, bool)`

GetCreatedLoyaltyPointsCountOk returns a tuple with the CreatedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) HasCreatedLoyaltyPointsCount() bool`

HasCreatedLoyaltyPointsCount returns a boolean if a field has been set.

### SetCreatedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) SetCreatedLoyaltyPointsCount(v float32)`

SetCreatedLoyaltyPointsCount gets a reference to the given float32 and assigns it to the CreatedLoyaltyPointsCount field.

### GetCreatedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsEffectCount() int32`

GetCreatedLoyaltyPointsEffectCount returns the CreatedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsEffectCountOk

`func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsEffectCountOk() (int32, bool)`

GetCreatedLoyaltyPointsEffectCountOk returns a tuple with the CreatedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) HasCreatedLoyaltyPointsEffectCount() bool`

HasCreatedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### SetCreatedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) SetCreatedLoyaltyPointsEffectCount(v int32)`

SetCreatedLoyaltyPointsEffectCount gets a reference to the given int32 and assigns it to the CreatedLoyaltyPointsEffectCount field.

### GetRedeemedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsCount() float32`

GetRedeemedLoyaltyPointsCount returns the RedeemedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsCountOk

`func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsCountOk() (float32, bool)`

GetRedeemedLoyaltyPointsCountOk returns a tuple with the RedeemedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRedeemedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) HasRedeemedLoyaltyPointsCount() bool`

HasRedeemedLoyaltyPointsCount returns a boolean if a field has been set.

### SetRedeemedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) SetRedeemedLoyaltyPointsCount(v float32)`

SetRedeemedLoyaltyPointsCount gets a reference to the given float32 and assigns it to the RedeemedLoyaltyPointsCount field.

### GetRedeemedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsEffectCount() int32`

GetRedeemedLoyaltyPointsEffectCount returns the RedeemedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsEffectCountOk

`func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsEffectCountOk() (int32, bool)`

GetRedeemedLoyaltyPointsEffectCountOk returns a tuple with the RedeemedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRedeemedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) HasRedeemedLoyaltyPointsEffectCount() bool`

HasRedeemedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### SetRedeemedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) SetRedeemedLoyaltyPointsEffectCount(v int32)`

SetRedeemedLoyaltyPointsEffectCount gets a reference to the given int32 and assigns it to the RedeemedLoyaltyPointsEffectCount field.

### GetCallApiEffectCount

`func (o *AdditionalCampaignProperties) GetCallApiEffectCount() int32`

GetCallApiEffectCount returns the CallApiEffectCount field if non-nil, zero value otherwise.

### GetCallApiEffectCountOk

`func (o *AdditionalCampaignProperties) GetCallApiEffectCountOk() (int32, bool)`

GetCallApiEffectCountOk returns a tuple with the CallApiEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallApiEffectCount

`func (o *AdditionalCampaignProperties) HasCallApiEffectCount() bool`

HasCallApiEffectCount returns a boolean if a field has been set.

### SetCallApiEffectCount

`func (o *AdditionalCampaignProperties) SetCallApiEffectCount(v int32)`

SetCallApiEffectCount gets a reference to the given int32 and assigns it to the CallApiEffectCount field.

### GetReservecouponEffectCount

`func (o *AdditionalCampaignProperties) GetReservecouponEffectCount() int32`

GetReservecouponEffectCount returns the ReservecouponEffectCount field if non-nil, zero value otherwise.

### GetReservecouponEffectCountOk

`func (o *AdditionalCampaignProperties) GetReservecouponEffectCountOk() (int32, bool)`

GetReservecouponEffectCountOk returns a tuple with the ReservecouponEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservecouponEffectCount

`func (o *AdditionalCampaignProperties) HasReservecouponEffectCount() bool`

HasReservecouponEffectCount returns a boolean if a field has been set.

### SetReservecouponEffectCount

`func (o *AdditionalCampaignProperties) SetReservecouponEffectCount(v int32)`

SetReservecouponEffectCount gets a reference to the given int32 and assigns it to the ReservecouponEffectCount field.

### GetLastActivity

`func (o *AdditionalCampaignProperties) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *AdditionalCampaignProperties) GetLastActivityOk() (time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActivity

`func (o *AdditionalCampaignProperties) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivity

`func (o *AdditionalCampaignProperties) SetLastActivity(v time.Time)`

SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.

### GetUpdated

`func (o *AdditionalCampaignProperties) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AdditionalCampaignProperties) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *AdditionalCampaignProperties) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *AdditionalCampaignProperties) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.

### GetCreatedBy

`func (o *AdditionalCampaignProperties) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AdditionalCampaignProperties) GetCreatedByOk() (string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *AdditionalCampaignProperties) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *AdditionalCampaignProperties) SetCreatedBy(v string)`

SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.

### GetUpdatedBy

`func (o *AdditionalCampaignProperties) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AdditionalCampaignProperties) GetUpdatedByOk() (string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedBy

`func (o *AdditionalCampaignProperties) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedBy

`func (o *AdditionalCampaignProperties) SetUpdatedBy(v string)`

SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.

### GetTemplateId

`func (o *AdditionalCampaignProperties) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *AdditionalCampaignProperties) GetTemplateIdOk() (int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateId

`func (o *AdditionalCampaignProperties) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateId

`func (o *AdditionalCampaignProperties) SetTemplateId(v int32)`

SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.

### GetFrontendState

`func (o *AdditionalCampaignProperties) GetFrontendState() string`

GetFrontendState returns the FrontendState field if non-nil, zero value otherwise.

### GetFrontendStateOk

`func (o *AdditionalCampaignProperties) GetFrontendStateOk() (string, bool)`

GetFrontendStateOk returns a tuple with the FrontendState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrontendState

`func (o *AdditionalCampaignProperties) HasFrontendState() bool`

HasFrontendState returns a boolean if a field has been set.

### SetFrontendState

`func (o *AdditionalCampaignProperties) SetFrontendState(v string)`

SetFrontendState gets a reference to the given string and assigns it to the FrontendState field.

### GetStoresImported

`func (o *AdditionalCampaignProperties) GetStoresImported() bool`

GetStoresImported returns the StoresImported field if non-nil, zero value otherwise.

### GetStoresImportedOk

`func (o *AdditionalCampaignProperties) GetStoresImportedOk() (bool, bool)`

GetStoresImportedOk returns a tuple with the StoresImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStoresImported

`func (o *AdditionalCampaignProperties) HasStoresImported() bool`

HasStoresImported returns a boolean if a field has been set.

### SetStoresImported

`func (o *AdditionalCampaignProperties) SetStoresImported(v bool)`

SetStoresImported gets a reference to the given bool and assigns it to the StoresImported field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


