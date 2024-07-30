/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
	"time"
)

// AdditionalCampaignProperties struct for AdditionalCampaignProperties
type AdditionalCampaignProperties struct {
	// A list of all the budgets that are defined by this campaign and their usage.  **Note:** Budgets that are not defined do not appear in this list and their usage is not counted until they are defined. 
	Budgets []CampaignBudget `json:"budgets"`
	// This property is **deprecated**. The count should be available under *budgets* property. Number of coupons redeemed in the campaign. 
	CouponRedemptionCount *int32 `json:"couponRedemptionCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Number of referral codes redeemed in the campaign. 
	ReferralRedemptionCount *int32 `json:"referralRedemptionCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total amount of discounts redeemed in the campaign. 
	DiscountCount *float32 `json:"discountCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of times discounts were redeemed in this campaign. 
	DiscountEffectCount *int32 `json:"discountEffectCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of coupons created by rules in this campaign. 
	CouponCreationCount *int32 `json:"couponCreationCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of custom effects triggered by rules in this campaign. 
	CustomEffectCount *int32 `json:"customEffectCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of referrals created by rules in this campaign. 
	ReferralCreationCount *int32 `json:"referralCreationCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of times the [add free item effect](https://docs.talon.one/docs/dev/integration-api/api-effects#addfreeitem) can be triggered in this campaign. 
	AddFreeItemEffectCount *int32 `json:"addFreeItemEffectCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of giveaways awarded by rules in this campaign. 
	AwardedGiveawaysCount *int32 `json:"awardedGiveawaysCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty points created by rules in this campaign. 
	CreatedLoyaltyPointsCount *float32 `json:"createdLoyaltyPointsCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty point creation effects triggered by rules in this campaign. 
	CreatedLoyaltyPointsEffectCount *int32 `json:"createdLoyaltyPointsEffectCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty points redeemed by rules in this campaign. 
	RedeemedLoyaltyPointsCount *float32 `json:"redeemedLoyaltyPointsCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of loyalty point redemption effects triggered by rules in this campaign. 
	RedeemedLoyaltyPointsEffectCount *int32 `json:"redeemedLoyaltyPointsEffectCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of webhooks triggered by rules in this campaign. 
	CallApiEffectCount *int32 `json:"callApiEffectCount,omitempty"`
	// This property is **deprecated**. The count should be available under *budgets* property. Total number of reserve coupon effects triggered by rules in this campaign. 
	ReservecouponEffectCount *int32 `json:"reservecouponEffectCount,omitempty"`
	// Timestamp of the most recent event received by this campaign.
	LastActivity *time.Time `json:"lastActivity,omitempty"`
	// Timestamp of the most recent update to the campaign's property. Updates to external entities used in this campaign are **not** registered by this property, such as collection or coupon updates. 
	Updated *time.Time `json:"updated,omitempty"`
	// Name of the user who created this campaign if available.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Name of the user who last updated this campaign if available.
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// The ID of the Campaign Template this Campaign was created from.
	TemplateId *int32 `json:"templateId,omitempty"`
	// A campaign state described exactly as in the Campaign Manager.
	FrontendState string `json:"frontendState"`
	// Indicates whether the linked stores were imported via a CSV file.
	StoresImported bool `json:"storesImported"`
}

// GetBudgets returns the Budgets field value
func (o *AdditionalCampaignProperties) GetBudgets() []CampaignBudget {
	if o == nil {
		var ret []CampaignBudget
		return ret
	}

	return o.Budgets
}

// SetBudgets sets field value
func (o *AdditionalCampaignProperties) SetBudgets(v []CampaignBudget) {
	o.Budgets = v
}

// GetCouponRedemptionCount returns the CouponRedemptionCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetCouponRedemptionCount() int32 {
	if o == nil || o.CouponRedemptionCount == nil {
		var ret int32
		return ret
	}
	return *o.CouponRedemptionCount
}

// GetCouponRedemptionCountOk returns a tuple with the CouponRedemptionCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetCouponRedemptionCountOk() (int32, bool) {
	if o == nil || o.CouponRedemptionCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CouponRedemptionCount, true
}

// HasCouponRedemptionCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasCouponRedemptionCount() bool {
	if o != nil && o.CouponRedemptionCount != nil {
		return true
	}

	return false
}

// SetCouponRedemptionCount gets a reference to the given int32 and assigns it to the CouponRedemptionCount field.
func (o *AdditionalCampaignProperties) SetCouponRedemptionCount(v int32) {
	o.CouponRedemptionCount = &v
}

// GetReferralRedemptionCount returns the ReferralRedemptionCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetReferralRedemptionCount() int32 {
	if o == nil || o.ReferralRedemptionCount == nil {
		var ret int32
		return ret
	}
	return *o.ReferralRedemptionCount
}

// GetReferralRedemptionCountOk returns a tuple with the ReferralRedemptionCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetReferralRedemptionCountOk() (int32, bool) {
	if o == nil || o.ReferralRedemptionCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ReferralRedemptionCount, true
}

// HasReferralRedemptionCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasReferralRedemptionCount() bool {
	if o != nil && o.ReferralRedemptionCount != nil {
		return true
	}

	return false
}

// SetReferralRedemptionCount gets a reference to the given int32 and assigns it to the ReferralRedemptionCount field.
func (o *AdditionalCampaignProperties) SetReferralRedemptionCount(v int32) {
	o.ReferralRedemptionCount = &v
}

// GetDiscountCount returns the DiscountCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetDiscountCount() float32 {
	if o == nil || o.DiscountCount == nil {
		var ret float32
		return ret
	}
	return *o.DiscountCount
}

// GetDiscountCountOk returns a tuple with the DiscountCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetDiscountCountOk() (float32, bool) {
	if o == nil || o.DiscountCount == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountCount, true
}

// HasDiscountCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasDiscountCount() bool {
	if o != nil && o.DiscountCount != nil {
		return true
	}

	return false
}

// SetDiscountCount gets a reference to the given float32 and assigns it to the DiscountCount field.
func (o *AdditionalCampaignProperties) SetDiscountCount(v float32) {
	o.DiscountCount = &v
}

// GetDiscountEffectCount returns the DiscountEffectCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetDiscountEffectCount() int32 {
	if o == nil || o.DiscountEffectCount == nil {
		var ret int32
		return ret
	}
	return *o.DiscountEffectCount
}

// GetDiscountEffectCountOk returns a tuple with the DiscountEffectCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetDiscountEffectCountOk() (int32, bool) {
	if o == nil || o.DiscountEffectCount == nil {
		var ret int32
		return ret, false
	}
	return *o.DiscountEffectCount, true
}

// HasDiscountEffectCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasDiscountEffectCount() bool {
	if o != nil && o.DiscountEffectCount != nil {
		return true
	}

	return false
}

// SetDiscountEffectCount gets a reference to the given int32 and assigns it to the DiscountEffectCount field.
func (o *AdditionalCampaignProperties) SetDiscountEffectCount(v int32) {
	o.DiscountEffectCount = &v
}

// GetCouponCreationCount returns the CouponCreationCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetCouponCreationCount() int32 {
	if o == nil || o.CouponCreationCount == nil {
		var ret int32
		return ret
	}
	return *o.CouponCreationCount
}

// GetCouponCreationCountOk returns a tuple with the CouponCreationCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetCouponCreationCountOk() (int32, bool) {
	if o == nil || o.CouponCreationCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CouponCreationCount, true
}

// HasCouponCreationCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasCouponCreationCount() bool {
	if o != nil && o.CouponCreationCount != nil {
		return true
	}

	return false
}

// SetCouponCreationCount gets a reference to the given int32 and assigns it to the CouponCreationCount field.
func (o *AdditionalCampaignProperties) SetCouponCreationCount(v int32) {
	o.CouponCreationCount = &v
}

// GetCustomEffectCount returns the CustomEffectCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetCustomEffectCount() int32 {
	if o == nil || o.CustomEffectCount == nil {
		var ret int32
		return ret
	}
	return *o.CustomEffectCount
}

// GetCustomEffectCountOk returns a tuple with the CustomEffectCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetCustomEffectCountOk() (int32, bool) {
	if o == nil || o.CustomEffectCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CustomEffectCount, true
}

// HasCustomEffectCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasCustomEffectCount() bool {
	if o != nil && o.CustomEffectCount != nil {
		return true
	}

	return false
}

// SetCustomEffectCount gets a reference to the given int32 and assigns it to the CustomEffectCount field.
func (o *AdditionalCampaignProperties) SetCustomEffectCount(v int32) {
	o.CustomEffectCount = &v
}

// GetReferralCreationCount returns the ReferralCreationCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetReferralCreationCount() int32 {
	if o == nil || o.ReferralCreationCount == nil {
		var ret int32
		return ret
	}
	return *o.ReferralCreationCount
}

// GetReferralCreationCountOk returns a tuple with the ReferralCreationCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetReferralCreationCountOk() (int32, bool) {
	if o == nil || o.ReferralCreationCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ReferralCreationCount, true
}

// HasReferralCreationCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasReferralCreationCount() bool {
	if o != nil && o.ReferralCreationCount != nil {
		return true
	}

	return false
}

// SetReferralCreationCount gets a reference to the given int32 and assigns it to the ReferralCreationCount field.
func (o *AdditionalCampaignProperties) SetReferralCreationCount(v int32) {
	o.ReferralCreationCount = &v
}

// GetAddFreeItemEffectCount returns the AddFreeItemEffectCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetAddFreeItemEffectCount() int32 {
	if o == nil || o.AddFreeItemEffectCount == nil {
		var ret int32
		return ret
	}
	return *o.AddFreeItemEffectCount
}

// GetAddFreeItemEffectCountOk returns a tuple with the AddFreeItemEffectCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetAddFreeItemEffectCountOk() (int32, bool) {
	if o == nil || o.AddFreeItemEffectCount == nil {
		var ret int32
		return ret, false
	}
	return *o.AddFreeItemEffectCount, true
}

// HasAddFreeItemEffectCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasAddFreeItemEffectCount() bool {
	if o != nil && o.AddFreeItemEffectCount != nil {
		return true
	}

	return false
}

// SetAddFreeItemEffectCount gets a reference to the given int32 and assigns it to the AddFreeItemEffectCount field.
func (o *AdditionalCampaignProperties) SetAddFreeItemEffectCount(v int32) {
	o.AddFreeItemEffectCount = &v
}

// GetAwardedGiveawaysCount returns the AwardedGiveawaysCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetAwardedGiveawaysCount() int32 {
	if o == nil || o.AwardedGiveawaysCount == nil {
		var ret int32
		return ret
	}
	return *o.AwardedGiveawaysCount
}

// GetAwardedGiveawaysCountOk returns a tuple with the AwardedGiveawaysCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetAwardedGiveawaysCountOk() (int32, bool) {
	if o == nil || o.AwardedGiveawaysCount == nil {
		var ret int32
		return ret, false
	}
	return *o.AwardedGiveawaysCount, true
}

// HasAwardedGiveawaysCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasAwardedGiveawaysCount() bool {
	if o != nil && o.AwardedGiveawaysCount != nil {
		return true
	}

	return false
}

// SetAwardedGiveawaysCount gets a reference to the given int32 and assigns it to the AwardedGiveawaysCount field.
func (o *AdditionalCampaignProperties) SetAwardedGiveawaysCount(v int32) {
	o.AwardedGiveawaysCount = &v
}

// GetCreatedLoyaltyPointsCount returns the CreatedLoyaltyPointsCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsCount() float32 {
	if o == nil || o.CreatedLoyaltyPointsCount == nil {
		var ret float32
		return ret
	}
	return *o.CreatedLoyaltyPointsCount
}

// GetCreatedLoyaltyPointsCountOk returns a tuple with the CreatedLoyaltyPointsCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsCountOk() (float32, bool) {
	if o == nil || o.CreatedLoyaltyPointsCount == nil {
		var ret float32
		return ret, false
	}
	return *o.CreatedLoyaltyPointsCount, true
}

// HasCreatedLoyaltyPointsCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasCreatedLoyaltyPointsCount() bool {
	if o != nil && o.CreatedLoyaltyPointsCount != nil {
		return true
	}

	return false
}

// SetCreatedLoyaltyPointsCount gets a reference to the given float32 and assigns it to the CreatedLoyaltyPointsCount field.
func (o *AdditionalCampaignProperties) SetCreatedLoyaltyPointsCount(v float32) {
	o.CreatedLoyaltyPointsCount = &v
}

// GetCreatedLoyaltyPointsEffectCount returns the CreatedLoyaltyPointsEffectCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsEffectCount() int32 {
	if o == nil || o.CreatedLoyaltyPointsEffectCount == nil {
		var ret int32
		return ret
	}
	return *o.CreatedLoyaltyPointsEffectCount
}

// GetCreatedLoyaltyPointsEffectCountOk returns a tuple with the CreatedLoyaltyPointsEffectCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetCreatedLoyaltyPointsEffectCountOk() (int32, bool) {
	if o == nil || o.CreatedLoyaltyPointsEffectCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CreatedLoyaltyPointsEffectCount, true
}

// HasCreatedLoyaltyPointsEffectCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasCreatedLoyaltyPointsEffectCount() bool {
	if o != nil && o.CreatedLoyaltyPointsEffectCount != nil {
		return true
	}

	return false
}

// SetCreatedLoyaltyPointsEffectCount gets a reference to the given int32 and assigns it to the CreatedLoyaltyPointsEffectCount field.
func (o *AdditionalCampaignProperties) SetCreatedLoyaltyPointsEffectCount(v int32) {
	o.CreatedLoyaltyPointsEffectCount = &v
}

// GetRedeemedLoyaltyPointsCount returns the RedeemedLoyaltyPointsCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsCount() float32 {
	if o == nil || o.RedeemedLoyaltyPointsCount == nil {
		var ret float32
		return ret
	}
	return *o.RedeemedLoyaltyPointsCount
}

// GetRedeemedLoyaltyPointsCountOk returns a tuple with the RedeemedLoyaltyPointsCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsCountOk() (float32, bool) {
	if o == nil || o.RedeemedLoyaltyPointsCount == nil {
		var ret float32
		return ret, false
	}
	return *o.RedeemedLoyaltyPointsCount, true
}

// HasRedeemedLoyaltyPointsCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasRedeemedLoyaltyPointsCount() bool {
	if o != nil && o.RedeemedLoyaltyPointsCount != nil {
		return true
	}

	return false
}

// SetRedeemedLoyaltyPointsCount gets a reference to the given float32 and assigns it to the RedeemedLoyaltyPointsCount field.
func (o *AdditionalCampaignProperties) SetRedeemedLoyaltyPointsCount(v float32) {
	o.RedeemedLoyaltyPointsCount = &v
}

// GetRedeemedLoyaltyPointsEffectCount returns the RedeemedLoyaltyPointsEffectCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsEffectCount() int32 {
	if o == nil || o.RedeemedLoyaltyPointsEffectCount == nil {
		var ret int32
		return ret
	}
	return *o.RedeemedLoyaltyPointsEffectCount
}

// GetRedeemedLoyaltyPointsEffectCountOk returns a tuple with the RedeemedLoyaltyPointsEffectCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetRedeemedLoyaltyPointsEffectCountOk() (int32, bool) {
	if o == nil || o.RedeemedLoyaltyPointsEffectCount == nil {
		var ret int32
		return ret, false
	}
	return *o.RedeemedLoyaltyPointsEffectCount, true
}

// HasRedeemedLoyaltyPointsEffectCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasRedeemedLoyaltyPointsEffectCount() bool {
	if o != nil && o.RedeemedLoyaltyPointsEffectCount != nil {
		return true
	}

	return false
}

// SetRedeemedLoyaltyPointsEffectCount gets a reference to the given int32 and assigns it to the RedeemedLoyaltyPointsEffectCount field.
func (o *AdditionalCampaignProperties) SetRedeemedLoyaltyPointsEffectCount(v int32) {
	o.RedeemedLoyaltyPointsEffectCount = &v
}

// GetCallApiEffectCount returns the CallApiEffectCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetCallApiEffectCount() int32 {
	if o == nil || o.CallApiEffectCount == nil {
		var ret int32
		return ret
	}
	return *o.CallApiEffectCount
}

// GetCallApiEffectCountOk returns a tuple with the CallApiEffectCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetCallApiEffectCountOk() (int32, bool) {
	if o == nil || o.CallApiEffectCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CallApiEffectCount, true
}

// HasCallApiEffectCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasCallApiEffectCount() bool {
	if o != nil && o.CallApiEffectCount != nil {
		return true
	}

	return false
}

// SetCallApiEffectCount gets a reference to the given int32 and assigns it to the CallApiEffectCount field.
func (o *AdditionalCampaignProperties) SetCallApiEffectCount(v int32) {
	o.CallApiEffectCount = &v
}

// GetReservecouponEffectCount returns the ReservecouponEffectCount field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetReservecouponEffectCount() int32 {
	if o == nil || o.ReservecouponEffectCount == nil {
		var ret int32
		return ret
	}
	return *o.ReservecouponEffectCount
}

// GetReservecouponEffectCountOk returns a tuple with the ReservecouponEffectCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetReservecouponEffectCountOk() (int32, bool) {
	if o == nil || o.ReservecouponEffectCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ReservecouponEffectCount, true
}

// HasReservecouponEffectCount returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasReservecouponEffectCount() bool {
	if o != nil && o.ReservecouponEffectCount != nil {
		return true
	}

	return false
}

// SetReservecouponEffectCount gets a reference to the given int32 and assigns it to the ReservecouponEffectCount field.
func (o *AdditionalCampaignProperties) SetReservecouponEffectCount(v int32) {
	o.ReservecouponEffectCount = &v
}

// GetLastActivity returns the LastActivity field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetLastActivity() time.Time {
	if o == nil || o.LastActivity == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivity
}

// GetLastActivityOk returns a tuple with the LastActivity field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetLastActivityOk() (time.Time, bool) {
	if o == nil || o.LastActivity == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastActivity, true
}

// HasLastActivity returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasLastActivity() bool {
	if o != nil && o.LastActivity != nil {
		return true
	}

	return false
}

// SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.
func (o *AdditionalCampaignProperties) SetLastActivity(v time.Time) {
	o.LastActivity = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetUpdated() time.Time {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetUpdatedOk() (time.Time, bool) {
	if o == nil || o.Updated == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *AdditionalCampaignProperties) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetCreatedByOk() (string, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AdditionalCampaignProperties) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetUpdatedByOk() (string, bool) {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret, false
	}
	return *o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *AdditionalCampaignProperties) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *AdditionalCampaignProperties) GetTemplateId() int32 {
	if o == nil || o.TemplateId == nil {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCampaignProperties) GetTemplateIdOk() (int32, bool) {
	if o == nil || o.TemplateId == nil {
		var ret int32
		return ret, false
	}
	return *o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *AdditionalCampaignProperties) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *AdditionalCampaignProperties) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetFrontendState returns the FrontendState field value
func (o *AdditionalCampaignProperties) GetFrontendState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FrontendState
}

// SetFrontendState sets field value
func (o *AdditionalCampaignProperties) SetFrontendState(v string) {
	o.FrontendState = v
}

// GetStoresImported returns the StoresImported field value
func (o *AdditionalCampaignProperties) GetStoresImported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StoresImported
}

// SetStoresImported sets field value
func (o *AdditionalCampaignProperties) SetStoresImported(v bool) {
	o.StoresImported = v
}

type NullableAdditionalCampaignProperties struct {
	Value AdditionalCampaignProperties
	ExplicitNull bool
}

func (v NullableAdditionalCampaignProperties) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAdditionalCampaignProperties) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
