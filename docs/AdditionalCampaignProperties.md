# AdditionalCampaignProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewAdditionalCampaignProperties

`func NewAdditionalCampaignProperties(frontendState string, storesImported bool, ) *AdditionalCampaignProperties`

NewAdditionalCampaignProperties instantiates a new AdditionalCampaignProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalCampaignPropertiesWithDefaults

`func NewAdditionalCampaignPropertiesWithDefaults() *AdditionalCampaignProperties`

NewAdditionalCampaignPropertiesWithDefaults instantiates a new AdditionalCampaignProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgets

`func (o *AdditionalCampaignProperties) GetBudgets() []CampaignBudget`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *AdditionalCampaignProperties) GetBudgetsOk() (*[]CampaignBudget, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *AdditionalCampaignProperties) SetBudgets(v []CampaignBudget)`

SetBudgets sets Budgets field to given value.

### HasBudgets

`func (o *AdditionalCampaignProperties) HasBudgets() bool`

HasBudgets returns a boolean if a field has been set.

### GetCouponRedemptionCount

`func (o *AdditionalCampaignProperties) GetCouponRedemptionCount() int64`

GetCouponRedemptionCount returns the CouponRedemptionCount field if non-nil, zero value otherwise.

### GetCouponRedemptionCountOk

`func (o *AdditionalCampaignProperties) GetCouponRedemptionCountOk() (*int64, bool)`

GetCouponRedemptionCountOk returns a tuple with the CouponRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRedemptionCount

`func (o *AdditionalCampaignProperties) SetCouponRedemptionCount(v int64)`

SetCouponRedemptionCount sets CouponRedemptionCount field to given value.

### HasCouponRedemptionCount

`func (o *AdditionalCampaignProperties) HasCouponRedemptionCount() bool`

HasCouponRedemptionCount returns a boolean if a field has been set.

### GetReferralRedemptionCount

`func (o *AdditionalCampaignProperties) GetReferralRedemptionCount() int64`

GetReferralRedemptionCount returns the ReferralRedemptionCount field if non-nil, zero value otherwise.

### GetReferralRedemptionCountOk

`func (o *AdditionalCampaignProperties) GetReferralRedemptionCountOk() (*int64, bool)`

GetReferralRedemptionCountOk returns a tuple with the ReferralRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralRedemptionCount

`func (o *AdditionalCampaignProperties) SetReferralRedemptionCount(v int64)`

SetReferralRedemptionCount sets ReferralRedemptionCount field to given value.

### HasReferralRedemptionCount

`func (o *AdditionalCampaignProperties) HasReferralRedemptionCount() bool`

HasReferralRedemptionCount returns a boolean if a field has been set.

### GetDiscountCount

`func (o *AdditionalCampaignProperties) GetDiscountCount() float32`

GetDiscountCount returns the DiscountCount field if non-nil, zero value otherwise.

### GetDiscountCountOk

`func (o *AdditionalCampaignProperties) GetDiscountCountOk() (*float32, bool)`

GetDiscountCountOk returns a tuple with the DiscountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCount

`func (o *AdditionalCampaignProperties) SetDiscountCount(v float32)`

SetDiscountCount sets DiscountCount field to given value.

### HasDiscountCount

`func (o *AdditionalCampaignProperties) HasDiscountCount() bool`

HasDiscountCount returns a boolean if a field has been set.

### GetDiscountEffectCount

`func (o *AdditionalCampaignProperties) GetDiscountEffectCount() int64`

GetDiscountEffectCount returns the DiscountEffectCount field if non-nil, zero value otherwise.

### GetDiscountEffectCountOk

`func (o *AdditionalCampaignProperties) GetDiscountEffectCountOk() (*int64, bool)`

GetDiscountEffectCountOk returns a tuple with the DiscountEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountEffectCount

`func (o *AdditionalCampaignProperties) SetDiscountEffectCount(v int64)`

SetDiscountEffectCount sets DiscountEffectCount field to given value.

### HasDiscountEffectCount

`func (o *AdditionalCampaignProperties) HasDiscountEffectCount() bool`

HasDiscountEffectCount returns a boolean if a field has been set.

### GetCouponCreationCount

`func (o *AdditionalCampaignProperties) GetCouponCreationCount() int64`

GetCouponCreationCount returns the CouponCreationCount field if non-nil, zero value otherwise.

### GetCouponCreationCountOk

`func (o *AdditionalCampaignProperties) GetCouponCreationCountOk() (*int64, bool)`

GetCouponCreationCountOk returns a tuple with the CouponCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCreationCount

`func (o *AdditionalCampaignProperties) SetCouponCreationCount(v int64)`

SetCouponCreationCount sets CouponCreationCount field to given value.

### HasCouponCreationCount

`func (o *AdditionalCampaignProperties) HasCouponCreationCount() bool`

HasCouponCreationCount returns a boolean if a field has been set.

### GetCustomEffectCount

`func (o *AdditionalCampaignProperties) GetCustomEffectCount() int64`

GetCustomEffectCount returns the CustomEffectCount field if non-nil, zero value otherwise.

### GetCustomEffectCountOk

`func (o *AdditionalCampaignProperties) GetCustomEffectCountOk() (*int64, bool)`

GetCustomEffectCountOk returns a tuple with the CustomEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEffectCount

`func (o *AdditionalCampaignProperties) SetCustomEffectCount(v int64)`

SetCustomEffectCount sets CustomEffectCount field to given value.

### HasCustomEffectCount

`func (o *AdditionalCampaignProperties) HasCustomEffectCount() bool`

HasCustomEffectCount returns a boolean if a field has been set.

### GetReferralCreationCount

`func (o *AdditionalCampaignProperties) GetReferralCreationCount() int64`

GetReferralCreationCount returns the ReferralCreationCount field if non-nil, zero value otherwise.

### GetReferralCreationCountOk

`func (o *AdditionalCampaignProperties) GetReferralCreationCountOk() (*int64, bool)`

GetReferralCreationCountOk returns a tuple with the ReferralCreationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCreationCount

`func (o *AdditionalCampaignProperties) SetReferralCreationCount(v int64)`

SetReferralCreationCount sets ReferralCreationCount field to given value.

### HasReferralCreationCount

`func (o *AdditionalCampaignProperties) HasReferralCreationCount() bool`

HasReferralCreationCount returns a boolean if a field has been set.

### GetAddFreeItemEffectCount

`func (o *AdditionalCampaignProperties) GetAddFreeItemEffectCount() int64`

GetAddFreeItemEffectCount returns the AddFreeItemEffectCount field if non-nil, zero value otherwise.

### GetAddFreeItemEffectCountOk

`func (o *AdditionalCampaignProperties) GetAddFreeItemEffectCountOk() (*int64, bool)`

GetAddFreeItemEffectCountOk returns a tuple with the AddFreeItemEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFreeItemEffectCount

`func (o *AdditionalCampaignProperties) SetAddFreeItemEffectCount(v int64)`

SetAddFreeItemEffectCount sets AddFreeItemEffectCount field to given value.

### HasAddFreeItemEffectCount

`func (o *AdditionalCampaignProperties) HasAddFreeItemEffectCount() bool`

HasAddFreeItemEffectCount returns a boolean if a field has been set.

### GetAwardedGiveawaysCount

`func (o *AdditionalCampaignProperties) GetAwardedGiveawaysCount() int64`

GetAwardedGiveawaysCount returns the AwardedGiveawaysCount field if non-nil, zero value otherwise.

### GetAwardedGiveawaysCountOk

`func (o *AdditionalCampaignProperties) GetAwardedGiveawaysCountOk() (*int64, bool)`

GetAwardedGiveawaysCountOk returns a tuple with the AwardedGiveawaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardedGiveawaysCount

`func (o *AdditionalCampaignProperties) SetAwardedGiveawaysCount(v int64)`

SetAwardedGiveawaysCount sets AwardedGiveawaysCount field to given value.

### HasAwardedGiveawaysCount

`func (o *AdditionalCampaignProperties) HasAwardedGiveawaysCount() bool`

HasAwardedGiveawaysCount returns a boolean if a field has been set.

### GetCreatedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsCount() float32`

GetCreatedLoyaltyPointsCount returns the CreatedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsCountOk

`func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsCountOk() (*float32, bool)`

GetCreatedLoyaltyPointsCountOk returns a tuple with the CreatedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) SetCreatedLoyaltyPointsCount(v float32)`

SetCreatedLoyaltyPointsCount sets CreatedLoyaltyPointsCount field to given value.

### HasCreatedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) HasCreatedLoyaltyPointsCount() bool`

HasCreatedLoyaltyPointsCount returns a boolean if a field has been set.

### GetCreatedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsEffectCount() int64`

GetCreatedLoyaltyPointsEffectCount returns the CreatedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetCreatedLoyaltyPointsEffectCountOk

`func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsEffectCountOk() (*int64, bool)`

GetCreatedLoyaltyPointsEffectCountOk returns a tuple with the CreatedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) SetCreatedLoyaltyPointsEffectCount(v int64)`

SetCreatedLoyaltyPointsEffectCount sets CreatedLoyaltyPointsEffectCount field to given value.

### HasCreatedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) HasCreatedLoyaltyPointsEffectCount() bool`

HasCreatedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### GetRedeemedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsCount() float32`

GetRedeemedLoyaltyPointsCount returns the RedeemedLoyaltyPointsCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsCountOk

`func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsCountOk() (*float32, bool)`

GetRedeemedLoyaltyPointsCountOk returns a tuple with the RedeemedLoyaltyPointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) SetRedeemedLoyaltyPointsCount(v float32)`

SetRedeemedLoyaltyPointsCount sets RedeemedLoyaltyPointsCount field to given value.

### HasRedeemedLoyaltyPointsCount

`func (o *AdditionalCampaignProperties) HasRedeemedLoyaltyPointsCount() bool`

HasRedeemedLoyaltyPointsCount returns a boolean if a field has been set.

### GetRedeemedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsEffectCount() int64`

GetRedeemedLoyaltyPointsEffectCount returns the RedeemedLoyaltyPointsEffectCount field if non-nil, zero value otherwise.

### GetRedeemedLoyaltyPointsEffectCountOk

`func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsEffectCountOk() (*int64, bool)`

GetRedeemedLoyaltyPointsEffectCountOk returns a tuple with the RedeemedLoyaltyPointsEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) SetRedeemedLoyaltyPointsEffectCount(v int64)`

SetRedeemedLoyaltyPointsEffectCount sets RedeemedLoyaltyPointsEffectCount field to given value.

### HasRedeemedLoyaltyPointsEffectCount

`func (o *AdditionalCampaignProperties) HasRedeemedLoyaltyPointsEffectCount() bool`

HasRedeemedLoyaltyPointsEffectCount returns a boolean if a field has been set.

### GetCallApiEffectCount

`func (o *AdditionalCampaignProperties) GetCallApiEffectCount() int64`

GetCallApiEffectCount returns the CallApiEffectCount field if non-nil, zero value otherwise.

### GetCallApiEffectCountOk

`func (o *AdditionalCampaignProperties) GetCallApiEffectCountOk() (*int64, bool)`

GetCallApiEffectCountOk returns a tuple with the CallApiEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallApiEffectCount

`func (o *AdditionalCampaignProperties) SetCallApiEffectCount(v int64)`

SetCallApiEffectCount sets CallApiEffectCount field to given value.

### HasCallApiEffectCount

`func (o *AdditionalCampaignProperties) HasCallApiEffectCount() bool`

HasCallApiEffectCount returns a boolean if a field has been set.

### GetReservecouponEffectCount

`func (o *AdditionalCampaignProperties) GetReservecouponEffectCount() int64`

GetReservecouponEffectCount returns the ReservecouponEffectCount field if non-nil, zero value otherwise.

### GetReservecouponEffectCountOk

`func (o *AdditionalCampaignProperties) GetReservecouponEffectCountOk() (*int64, bool)`

GetReservecouponEffectCountOk returns a tuple with the ReservecouponEffectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservecouponEffectCount

`func (o *AdditionalCampaignProperties) SetReservecouponEffectCount(v int64)`

SetReservecouponEffectCount sets ReservecouponEffectCount field to given value.

### HasReservecouponEffectCount

`func (o *AdditionalCampaignProperties) HasReservecouponEffectCount() bool`

HasReservecouponEffectCount returns a boolean if a field has been set.

### GetLastActivity

`func (o *AdditionalCampaignProperties) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *AdditionalCampaignProperties) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *AdditionalCampaignProperties) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *AdditionalCampaignProperties) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetUpdated

`func (o *AdditionalCampaignProperties) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AdditionalCampaignProperties) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AdditionalCampaignProperties) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AdditionalCampaignProperties) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AdditionalCampaignProperties) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AdditionalCampaignProperties) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AdditionalCampaignProperties) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AdditionalCampaignProperties) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AdditionalCampaignProperties) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AdditionalCampaignProperties) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AdditionalCampaignProperties) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AdditionalCampaignProperties) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetTemplateId

`func (o *AdditionalCampaignProperties) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *AdditionalCampaignProperties) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *AdditionalCampaignProperties) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *AdditionalCampaignProperties) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetFrontendState

`func (o *AdditionalCampaignProperties) GetFrontendState() string`

GetFrontendState returns the FrontendState field if non-nil, zero value otherwise.

### GetFrontendStateOk

`func (o *AdditionalCampaignProperties) GetFrontendStateOk() (*string, bool)`

GetFrontendStateOk returns a tuple with the FrontendState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendState

`func (o *AdditionalCampaignProperties) SetFrontendState(v string)`

SetFrontendState sets FrontendState field to given value.


### GetStoresImported

`func (o *AdditionalCampaignProperties) GetStoresImported() bool`

GetStoresImported returns the StoresImported field if non-nil, zero value otherwise.

### GetStoresImportedOk

`func (o *AdditionalCampaignProperties) GetStoresImportedOk() (*bool, bool)`

GetStoresImportedOk returns a tuple with the StoresImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresImported

`func (o *AdditionalCampaignProperties) SetStoresImported(v bool)`

SetStoresImported sets StoresImported field to given value.


### GetValueMapsIds

`func (o *AdditionalCampaignProperties) GetValueMapsIds() []int64`

GetValueMapsIds returns the ValueMapsIds field if non-nil, zero value otherwise.

### GetValueMapsIdsOk

`func (o *AdditionalCampaignProperties) GetValueMapsIdsOk() (*[]int64, bool)`

GetValueMapsIdsOk returns a tuple with the ValueMapsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMapsIds

`func (o *AdditionalCampaignProperties) SetValueMapsIds(v []int64)`

SetValueMapsIds sets ValueMapsIds field to given value.

### HasValueMapsIds

`func (o *AdditionalCampaignProperties) HasValueMapsIds() bool`

HasValueMapsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


