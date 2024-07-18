/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
	"time"
)

// CampaignAnalytics struct for CampaignAnalytics
type CampaignAnalytics struct {
	Date time.Time `json:"date"`
	// Amount of revenue in this campaign (for coupon or discount sessions).
	CampaignRevenue float32 `json:"campaignRevenue"`
	// Amount of revenue in this campaign since it began (for coupon or discount sessions).
	TotalCampaignRevenue float32 `json:"totalCampaignRevenue"`
	// Amount of refunds in this campaign (for coupon or discount sessions).
	CampaignRefund float32 `json:"campaignRefund"`
	// Amount of refunds in this campaign since it began (for coupon or discount sessions).
	TotalCampaignRefund float32 `json:"totalCampaignRefund"`
	// Amount of cost caused by discounts given in the campaign.
	CampaignDiscountCosts float32 `json:"campaignDiscountCosts"`
	// Amount of cost caused by discounts given in the campaign since it began.
	TotalCampaignDiscountCosts float32 `json:"totalCampaignDiscountCosts"`
	// Amount of discounts rolledback due to refund in the campaign.
	CampaignRefundedDiscounts float32 `json:"campaignRefundedDiscounts"`
	// Amount of discounts rolledback due to refund in the campaign since it began.
	TotalCampaignRefundedDiscounts float32 `json:"totalCampaignRefundedDiscounts"`
	// Amount of free items given in the campaign.
	CampaignFreeItems int32 `json:"campaignFreeItems"`
	// Amount of free items given in the campaign since it began.
	TotalCampaignFreeItems int32 `json:"totalCampaignFreeItems"`
	// Number of coupon redemptions in the campaign.
	CouponRedemptions int32 `json:"couponRedemptions"`
	// Number of coupon redemptions in the campaign since it began.
	TotalCouponRedemptions int32 `json:"totalCouponRedemptions"`
	// Number of coupon redemptions that have been rolled back (due to canceling closed session) in the campaign.
	CouponRolledbackRedemptions int32 `json:"couponRolledbackRedemptions"`
	// Number of coupon redemptions that have been rolled back (due to canceling closed session) in the campaign since it began.
	TotalCouponRolledbackRedemptions int32 `json:"totalCouponRolledbackRedemptions"`
	// Number of referral redemptions in the campaign.
	ReferralRedemptions int32 `json:"referralRedemptions"`
	// Number of referral redemptions in the campaign since it began.
	TotalReferralRedemptions int32 `json:"totalReferralRedemptions"`
	// Number of coupons created in the campaign by the rule engine.
	CouponsCreated int32 `json:"couponsCreated"`
	// Number of coupons created in the campaign by the rule engine since it began.
	TotalCouponsCreated int32 `json:"totalCouponsCreated"`
	// Number of referrals created in the campaign by the rule engine.
	ReferralsCreated int32 `json:"referralsCreated"`
	// Number of referrals created in the campaign by the rule engine since it began.
	TotalReferralsCreated int32 `json:"totalReferralsCreated"`
	// Number of added loyalty points in the campaign in a specific interval.
	AddedLoyaltyPoints float32 `json:"addedLoyaltyPoints"`
	// Number of added loyalty points in the campaign since it began.
	TotalAddedLoyaltyPoints float32 `json:"totalAddedLoyaltyPoints"`
	// Number of deducted loyalty points in the campaign in a specific interval.
	DeductedLoyaltyPoints float32 `json:"deductedLoyaltyPoints"`
	// Number of deducted loyalty points in the campaign since it began.
	TotalDeductedLoyaltyPoints float32 `json:"totalDeductedLoyaltyPoints"`
}

// GetDate returns the Date field value
func (o *CampaignAnalytics) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// SetDate sets field value
func (o *CampaignAnalytics) SetDate(v time.Time) {
	o.Date = v
}

// GetCampaignRevenue returns the CampaignRevenue field value
func (o *CampaignAnalytics) GetCampaignRevenue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CampaignRevenue
}

// SetCampaignRevenue sets field value
func (o *CampaignAnalytics) SetCampaignRevenue(v float32) {
	o.CampaignRevenue = v
}

// GetTotalCampaignRevenue returns the TotalCampaignRevenue field value
func (o *CampaignAnalytics) GetTotalCampaignRevenue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCampaignRevenue
}

// SetTotalCampaignRevenue sets field value
func (o *CampaignAnalytics) SetTotalCampaignRevenue(v float32) {
	o.TotalCampaignRevenue = v
}

// GetCampaignRefund returns the CampaignRefund field value
func (o *CampaignAnalytics) GetCampaignRefund() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CampaignRefund
}

// SetCampaignRefund sets field value
func (o *CampaignAnalytics) SetCampaignRefund(v float32) {
	o.CampaignRefund = v
}

// GetTotalCampaignRefund returns the TotalCampaignRefund field value
func (o *CampaignAnalytics) GetTotalCampaignRefund() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCampaignRefund
}

// SetTotalCampaignRefund sets field value
func (o *CampaignAnalytics) SetTotalCampaignRefund(v float32) {
	o.TotalCampaignRefund = v
}

// GetCampaignDiscountCosts returns the CampaignDiscountCosts field value
func (o *CampaignAnalytics) GetCampaignDiscountCosts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CampaignDiscountCosts
}

// SetCampaignDiscountCosts sets field value
func (o *CampaignAnalytics) SetCampaignDiscountCosts(v float32) {
	o.CampaignDiscountCosts = v
}

// GetTotalCampaignDiscountCosts returns the TotalCampaignDiscountCosts field value
func (o *CampaignAnalytics) GetTotalCampaignDiscountCosts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCampaignDiscountCosts
}

// SetTotalCampaignDiscountCosts sets field value
func (o *CampaignAnalytics) SetTotalCampaignDiscountCosts(v float32) {
	o.TotalCampaignDiscountCosts = v
}

// GetCampaignRefundedDiscounts returns the CampaignRefundedDiscounts field value
func (o *CampaignAnalytics) GetCampaignRefundedDiscounts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CampaignRefundedDiscounts
}

// SetCampaignRefundedDiscounts sets field value
func (o *CampaignAnalytics) SetCampaignRefundedDiscounts(v float32) {
	o.CampaignRefundedDiscounts = v
}

// GetTotalCampaignRefundedDiscounts returns the TotalCampaignRefundedDiscounts field value
func (o *CampaignAnalytics) GetTotalCampaignRefundedDiscounts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCampaignRefundedDiscounts
}

// SetTotalCampaignRefundedDiscounts sets field value
func (o *CampaignAnalytics) SetTotalCampaignRefundedDiscounts(v float32) {
	o.TotalCampaignRefundedDiscounts = v
}

// GetCampaignFreeItems returns the CampaignFreeItems field value
func (o *CampaignAnalytics) GetCampaignFreeItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignFreeItems
}

// SetCampaignFreeItems sets field value
func (o *CampaignAnalytics) SetCampaignFreeItems(v int32) {
	o.CampaignFreeItems = v
}

// GetTotalCampaignFreeItems returns the TotalCampaignFreeItems field value
func (o *CampaignAnalytics) GetTotalCampaignFreeItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCampaignFreeItems
}

// SetTotalCampaignFreeItems sets field value
func (o *CampaignAnalytics) SetTotalCampaignFreeItems(v int32) {
	o.TotalCampaignFreeItems = v
}

// GetCouponRedemptions returns the CouponRedemptions field value
func (o *CampaignAnalytics) GetCouponRedemptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponRedemptions
}

// SetCouponRedemptions sets field value
func (o *CampaignAnalytics) SetCouponRedemptions(v int32) {
	o.CouponRedemptions = v
}

// GetTotalCouponRedemptions returns the TotalCouponRedemptions field value
func (o *CampaignAnalytics) GetTotalCouponRedemptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCouponRedemptions
}

// SetTotalCouponRedemptions sets field value
func (o *CampaignAnalytics) SetTotalCouponRedemptions(v int32) {
	o.TotalCouponRedemptions = v
}

// GetCouponRolledbackRedemptions returns the CouponRolledbackRedemptions field value
func (o *CampaignAnalytics) GetCouponRolledbackRedemptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponRolledbackRedemptions
}

// SetCouponRolledbackRedemptions sets field value
func (o *CampaignAnalytics) SetCouponRolledbackRedemptions(v int32) {
	o.CouponRolledbackRedemptions = v
}

// GetTotalCouponRolledbackRedemptions returns the TotalCouponRolledbackRedemptions field value
func (o *CampaignAnalytics) GetTotalCouponRolledbackRedemptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCouponRolledbackRedemptions
}

// SetTotalCouponRolledbackRedemptions sets field value
func (o *CampaignAnalytics) SetTotalCouponRolledbackRedemptions(v int32) {
	o.TotalCouponRolledbackRedemptions = v
}

// GetReferralRedemptions returns the ReferralRedemptions field value
func (o *CampaignAnalytics) GetReferralRedemptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferralRedemptions
}

// SetReferralRedemptions sets field value
func (o *CampaignAnalytics) SetReferralRedemptions(v int32) {
	o.ReferralRedemptions = v
}

// GetTotalReferralRedemptions returns the TotalReferralRedemptions field value
func (o *CampaignAnalytics) GetTotalReferralRedemptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalReferralRedemptions
}

// SetTotalReferralRedemptions sets field value
func (o *CampaignAnalytics) SetTotalReferralRedemptions(v int32) {
	o.TotalReferralRedemptions = v
}

// GetCouponsCreated returns the CouponsCreated field value
func (o *CampaignAnalytics) GetCouponsCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponsCreated
}

// SetCouponsCreated sets field value
func (o *CampaignAnalytics) SetCouponsCreated(v int32) {
	o.CouponsCreated = v
}

// GetTotalCouponsCreated returns the TotalCouponsCreated field value
func (o *CampaignAnalytics) GetTotalCouponsCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCouponsCreated
}

// SetTotalCouponsCreated sets field value
func (o *CampaignAnalytics) SetTotalCouponsCreated(v int32) {
	o.TotalCouponsCreated = v
}

// GetReferralsCreated returns the ReferralsCreated field value
func (o *CampaignAnalytics) GetReferralsCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferralsCreated
}

// SetReferralsCreated sets field value
func (o *CampaignAnalytics) SetReferralsCreated(v int32) {
	o.ReferralsCreated = v
}

// GetTotalReferralsCreated returns the TotalReferralsCreated field value
func (o *CampaignAnalytics) GetTotalReferralsCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalReferralsCreated
}

// SetTotalReferralsCreated sets field value
func (o *CampaignAnalytics) SetTotalReferralsCreated(v int32) {
	o.TotalReferralsCreated = v
}

// GetAddedLoyaltyPoints returns the AddedLoyaltyPoints field value
func (o *CampaignAnalytics) GetAddedLoyaltyPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AddedLoyaltyPoints
}

// SetAddedLoyaltyPoints sets field value
func (o *CampaignAnalytics) SetAddedLoyaltyPoints(v float32) {
	o.AddedLoyaltyPoints = v
}

// GetTotalAddedLoyaltyPoints returns the TotalAddedLoyaltyPoints field value
func (o *CampaignAnalytics) GetTotalAddedLoyaltyPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalAddedLoyaltyPoints
}

// SetTotalAddedLoyaltyPoints sets field value
func (o *CampaignAnalytics) SetTotalAddedLoyaltyPoints(v float32) {
	o.TotalAddedLoyaltyPoints = v
}

// GetDeductedLoyaltyPoints returns the DeductedLoyaltyPoints field value
func (o *CampaignAnalytics) GetDeductedLoyaltyPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DeductedLoyaltyPoints
}

// SetDeductedLoyaltyPoints sets field value
func (o *CampaignAnalytics) SetDeductedLoyaltyPoints(v float32) {
	o.DeductedLoyaltyPoints = v
}

// GetTotalDeductedLoyaltyPoints returns the TotalDeductedLoyaltyPoints field value
func (o *CampaignAnalytics) GetTotalDeductedLoyaltyPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalDeductedLoyaltyPoints
}

// SetTotalDeductedLoyaltyPoints sets field value
func (o *CampaignAnalytics) SetTotalDeductedLoyaltyPoints(v float32) {
	o.TotalDeductedLoyaltyPoints = v
}

type NullableCampaignAnalytics struct {
	Value CampaignAnalytics
	ExplicitNull bool
}

func (v NullableCampaignAnalytics) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignAnalytics) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
