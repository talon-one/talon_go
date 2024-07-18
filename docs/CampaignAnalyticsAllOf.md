# CampaignAnalyticsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**time.Time**](time.Time.md) |  | 
**CampaignRevenue** | Pointer to **float32** | Amount of revenue in this campaign (for coupon or discount sessions). | 
**TotalCampaignRevenue** | Pointer to **float32** | Amount of revenue in this campaign since it began (for coupon or discount sessions). | 
**CampaignRefund** | Pointer to **float32** | Amount of refunds in this campaign (for coupon or discount sessions). | 
**TotalCampaignRefund** | Pointer to **float32** | Amount of refunds in this campaign since it began (for coupon or discount sessions). | 
**CampaignDiscountCosts** | Pointer to **float32** | Amount of cost caused by discounts given in the campaign. | 
**TotalCampaignDiscountCosts** | Pointer to **float32** | Amount of cost caused by discounts given in the campaign since it began. | 
**CampaignRefundedDiscounts** | Pointer to **float32** | Amount of discounts rolledback due to refund in the campaign. | 
**TotalCampaignRefundedDiscounts** | Pointer to **float32** | Amount of discounts rolledback due to refund in the campaign since it began. | 
**CampaignFreeItems** | Pointer to **int32** | Amount of free items given in the campaign. | 
**TotalCampaignFreeItems** | Pointer to **int32** | Amount of free items given in the campaign since it began. | 
**CouponRedemptions** | Pointer to **int32** | Number of coupon redemptions in the campaign. | 
**TotalCouponRedemptions** | Pointer to **int32** | Number of coupon redemptions in the campaign since it began. | 
**CouponRolledbackRedemptions** | Pointer to **int32** | Number of coupon redemptions that have been rolled back (due to canceling closed session) in the campaign. | 
**TotalCouponRolledbackRedemptions** | Pointer to **int32** | Number of coupon redemptions that have been rolled back (due to canceling closed session) in the campaign since it began. | 
**ReferralRedemptions** | Pointer to **int32** | Number of referral redemptions in the campaign. | 
**TotalReferralRedemptions** | Pointer to **int32** | Number of referral redemptions in the campaign since it began. | 
**CouponsCreated** | Pointer to **int32** | Number of coupons created in the campaign by the rule engine. | 
**TotalCouponsCreated** | Pointer to **int32** | Number of coupons created in the campaign by the rule engine since it began. | 
**ReferralsCreated** | Pointer to **int32** | Number of referrals created in the campaign by the rule engine. | 
**TotalReferralsCreated** | Pointer to **int32** | Number of referrals created in the campaign by the rule engine since it began. | 
**AddedLoyaltyPoints** | Pointer to **float32** | Number of added loyalty points in the campaign in a specific interval. | 
**TotalAddedLoyaltyPoints** | Pointer to **float32** | Number of added loyalty points in the campaign since it began. | 
**DeductedLoyaltyPoints** | Pointer to **float32** | Number of deducted loyalty points in the campaign in a specific interval. | 
**TotalDeductedLoyaltyPoints** | Pointer to **float32** | Number of deducted loyalty points in the campaign since it began. | 

## Methods

### GetDate

`func (o *CampaignAnalyticsAllOf) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CampaignAnalyticsAllOf) GetDateOk() (time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDate

`func (o *CampaignAnalyticsAllOf) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDate

`func (o *CampaignAnalyticsAllOf) SetDate(v time.Time)`

SetDate gets a reference to the given time.Time and assigns it to the Date field.

### GetCampaignRevenue

`func (o *CampaignAnalyticsAllOf) GetCampaignRevenue() float32`

GetCampaignRevenue returns the CampaignRevenue field if non-nil, zero value otherwise.

### GetCampaignRevenueOk

`func (o *CampaignAnalyticsAllOf) GetCampaignRevenueOk() (float32, bool)`

GetCampaignRevenueOk returns a tuple with the CampaignRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignRevenue

`func (o *CampaignAnalyticsAllOf) HasCampaignRevenue() bool`

HasCampaignRevenue returns a boolean if a field has been set.

### SetCampaignRevenue

`func (o *CampaignAnalyticsAllOf) SetCampaignRevenue(v float32)`

SetCampaignRevenue gets a reference to the given float32 and assigns it to the CampaignRevenue field.

### GetTotalCampaignRevenue

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignRevenue() float32`

GetTotalCampaignRevenue returns the TotalCampaignRevenue field if non-nil, zero value otherwise.

### GetTotalCampaignRevenueOk

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignRevenueOk() (float32, bool)`

GetTotalCampaignRevenueOk returns a tuple with the TotalCampaignRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignRevenue

`func (o *CampaignAnalyticsAllOf) HasTotalCampaignRevenue() bool`

HasTotalCampaignRevenue returns a boolean if a field has been set.

### SetTotalCampaignRevenue

`func (o *CampaignAnalyticsAllOf) SetTotalCampaignRevenue(v float32)`

SetTotalCampaignRevenue gets a reference to the given float32 and assigns it to the TotalCampaignRevenue field.

### GetCampaignRefund

`func (o *CampaignAnalyticsAllOf) GetCampaignRefund() float32`

GetCampaignRefund returns the CampaignRefund field if non-nil, zero value otherwise.

### GetCampaignRefundOk

`func (o *CampaignAnalyticsAllOf) GetCampaignRefundOk() (float32, bool)`

GetCampaignRefundOk returns a tuple with the CampaignRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignRefund

`func (o *CampaignAnalyticsAllOf) HasCampaignRefund() bool`

HasCampaignRefund returns a boolean if a field has been set.

### SetCampaignRefund

`func (o *CampaignAnalyticsAllOf) SetCampaignRefund(v float32)`

SetCampaignRefund gets a reference to the given float32 and assigns it to the CampaignRefund field.

### GetTotalCampaignRefund

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignRefund() float32`

GetTotalCampaignRefund returns the TotalCampaignRefund field if non-nil, zero value otherwise.

### GetTotalCampaignRefundOk

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignRefundOk() (float32, bool)`

GetTotalCampaignRefundOk returns a tuple with the TotalCampaignRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignRefund

`func (o *CampaignAnalyticsAllOf) HasTotalCampaignRefund() bool`

HasTotalCampaignRefund returns a boolean if a field has been set.

### SetTotalCampaignRefund

`func (o *CampaignAnalyticsAllOf) SetTotalCampaignRefund(v float32)`

SetTotalCampaignRefund gets a reference to the given float32 and assigns it to the TotalCampaignRefund field.

### GetCampaignDiscountCosts

`func (o *CampaignAnalyticsAllOf) GetCampaignDiscountCosts() float32`

GetCampaignDiscountCosts returns the CampaignDiscountCosts field if non-nil, zero value otherwise.

### GetCampaignDiscountCostsOk

`func (o *CampaignAnalyticsAllOf) GetCampaignDiscountCostsOk() (float32, bool)`

GetCampaignDiscountCostsOk returns a tuple with the CampaignDiscountCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignDiscountCosts

`func (o *CampaignAnalyticsAllOf) HasCampaignDiscountCosts() bool`

HasCampaignDiscountCosts returns a boolean if a field has been set.

### SetCampaignDiscountCosts

`func (o *CampaignAnalyticsAllOf) SetCampaignDiscountCosts(v float32)`

SetCampaignDiscountCosts gets a reference to the given float32 and assigns it to the CampaignDiscountCosts field.

### GetTotalCampaignDiscountCosts

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignDiscountCosts() float32`

GetTotalCampaignDiscountCosts returns the TotalCampaignDiscountCosts field if non-nil, zero value otherwise.

### GetTotalCampaignDiscountCostsOk

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignDiscountCostsOk() (float32, bool)`

GetTotalCampaignDiscountCostsOk returns a tuple with the TotalCampaignDiscountCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignDiscountCosts

`func (o *CampaignAnalyticsAllOf) HasTotalCampaignDiscountCosts() bool`

HasTotalCampaignDiscountCosts returns a boolean if a field has been set.

### SetTotalCampaignDiscountCosts

`func (o *CampaignAnalyticsAllOf) SetTotalCampaignDiscountCosts(v float32)`

SetTotalCampaignDiscountCosts gets a reference to the given float32 and assigns it to the TotalCampaignDiscountCosts field.

### GetCampaignRefundedDiscounts

`func (o *CampaignAnalyticsAllOf) GetCampaignRefundedDiscounts() float32`

GetCampaignRefundedDiscounts returns the CampaignRefundedDiscounts field if non-nil, zero value otherwise.

### GetCampaignRefundedDiscountsOk

`func (o *CampaignAnalyticsAllOf) GetCampaignRefundedDiscountsOk() (float32, bool)`

GetCampaignRefundedDiscountsOk returns a tuple with the CampaignRefundedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignRefundedDiscounts

`func (o *CampaignAnalyticsAllOf) HasCampaignRefundedDiscounts() bool`

HasCampaignRefundedDiscounts returns a boolean if a field has been set.

### SetCampaignRefundedDiscounts

`func (o *CampaignAnalyticsAllOf) SetCampaignRefundedDiscounts(v float32)`

SetCampaignRefundedDiscounts gets a reference to the given float32 and assigns it to the CampaignRefundedDiscounts field.

### GetTotalCampaignRefundedDiscounts

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignRefundedDiscounts() float32`

GetTotalCampaignRefundedDiscounts returns the TotalCampaignRefundedDiscounts field if non-nil, zero value otherwise.

### GetTotalCampaignRefundedDiscountsOk

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignRefundedDiscountsOk() (float32, bool)`

GetTotalCampaignRefundedDiscountsOk returns a tuple with the TotalCampaignRefundedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignRefundedDiscounts

`func (o *CampaignAnalyticsAllOf) HasTotalCampaignRefundedDiscounts() bool`

HasTotalCampaignRefundedDiscounts returns a boolean if a field has been set.

### SetTotalCampaignRefundedDiscounts

`func (o *CampaignAnalyticsAllOf) SetTotalCampaignRefundedDiscounts(v float32)`

SetTotalCampaignRefundedDiscounts gets a reference to the given float32 and assigns it to the TotalCampaignRefundedDiscounts field.

### GetCampaignFreeItems

`func (o *CampaignAnalyticsAllOf) GetCampaignFreeItems() int32`

GetCampaignFreeItems returns the CampaignFreeItems field if non-nil, zero value otherwise.

### GetCampaignFreeItemsOk

`func (o *CampaignAnalyticsAllOf) GetCampaignFreeItemsOk() (int32, bool)`

GetCampaignFreeItemsOk returns a tuple with the CampaignFreeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignFreeItems

`func (o *CampaignAnalyticsAllOf) HasCampaignFreeItems() bool`

HasCampaignFreeItems returns a boolean if a field has been set.

### SetCampaignFreeItems

`func (o *CampaignAnalyticsAllOf) SetCampaignFreeItems(v int32)`

SetCampaignFreeItems gets a reference to the given int32 and assigns it to the CampaignFreeItems field.

### GetTotalCampaignFreeItems

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignFreeItems() int32`

GetTotalCampaignFreeItems returns the TotalCampaignFreeItems field if non-nil, zero value otherwise.

### GetTotalCampaignFreeItemsOk

`func (o *CampaignAnalyticsAllOf) GetTotalCampaignFreeItemsOk() (int32, bool)`

GetTotalCampaignFreeItemsOk returns a tuple with the TotalCampaignFreeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignFreeItems

`func (o *CampaignAnalyticsAllOf) HasTotalCampaignFreeItems() bool`

HasTotalCampaignFreeItems returns a boolean if a field has been set.

### SetTotalCampaignFreeItems

`func (o *CampaignAnalyticsAllOf) SetTotalCampaignFreeItems(v int32)`

SetTotalCampaignFreeItems gets a reference to the given int32 and assigns it to the TotalCampaignFreeItems field.

### GetCouponRedemptions

`func (o *CampaignAnalyticsAllOf) GetCouponRedemptions() int32`

GetCouponRedemptions returns the CouponRedemptions field if non-nil, zero value otherwise.

### GetCouponRedemptionsOk

`func (o *CampaignAnalyticsAllOf) GetCouponRedemptionsOk() (int32, bool)`

GetCouponRedemptionsOk returns a tuple with the CouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRedemptions

`func (o *CampaignAnalyticsAllOf) HasCouponRedemptions() bool`

HasCouponRedemptions returns a boolean if a field has been set.

### SetCouponRedemptions

`func (o *CampaignAnalyticsAllOf) SetCouponRedemptions(v int32)`

SetCouponRedemptions gets a reference to the given int32 and assigns it to the CouponRedemptions field.

### GetTotalCouponRedemptions

`func (o *CampaignAnalyticsAllOf) GetTotalCouponRedemptions() int32`

GetTotalCouponRedemptions returns the TotalCouponRedemptions field if non-nil, zero value otherwise.

### GetTotalCouponRedemptionsOk

`func (o *CampaignAnalyticsAllOf) GetTotalCouponRedemptionsOk() (int32, bool)`

GetTotalCouponRedemptionsOk returns a tuple with the TotalCouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCouponRedemptions

`func (o *CampaignAnalyticsAllOf) HasTotalCouponRedemptions() bool`

HasTotalCouponRedemptions returns a boolean if a field has been set.

### SetTotalCouponRedemptions

`func (o *CampaignAnalyticsAllOf) SetTotalCouponRedemptions(v int32)`

SetTotalCouponRedemptions gets a reference to the given int32 and assigns it to the TotalCouponRedemptions field.

### GetCouponRolledbackRedemptions

`func (o *CampaignAnalyticsAllOf) GetCouponRolledbackRedemptions() int32`

GetCouponRolledbackRedemptions returns the CouponRolledbackRedemptions field if non-nil, zero value otherwise.

### GetCouponRolledbackRedemptionsOk

`func (o *CampaignAnalyticsAllOf) GetCouponRolledbackRedemptionsOk() (int32, bool)`

GetCouponRolledbackRedemptionsOk returns a tuple with the CouponRolledbackRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRolledbackRedemptions

`func (o *CampaignAnalyticsAllOf) HasCouponRolledbackRedemptions() bool`

HasCouponRolledbackRedemptions returns a boolean if a field has been set.

### SetCouponRolledbackRedemptions

`func (o *CampaignAnalyticsAllOf) SetCouponRolledbackRedemptions(v int32)`

SetCouponRolledbackRedemptions gets a reference to the given int32 and assigns it to the CouponRolledbackRedemptions field.

### GetTotalCouponRolledbackRedemptions

`func (o *CampaignAnalyticsAllOf) GetTotalCouponRolledbackRedemptions() int32`

GetTotalCouponRolledbackRedemptions returns the TotalCouponRolledbackRedemptions field if non-nil, zero value otherwise.

### GetTotalCouponRolledbackRedemptionsOk

`func (o *CampaignAnalyticsAllOf) GetTotalCouponRolledbackRedemptionsOk() (int32, bool)`

GetTotalCouponRolledbackRedemptionsOk returns a tuple with the TotalCouponRolledbackRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCouponRolledbackRedemptions

`func (o *CampaignAnalyticsAllOf) HasTotalCouponRolledbackRedemptions() bool`

HasTotalCouponRolledbackRedemptions returns a boolean if a field has been set.

### SetTotalCouponRolledbackRedemptions

`func (o *CampaignAnalyticsAllOf) SetTotalCouponRolledbackRedemptions(v int32)`

SetTotalCouponRolledbackRedemptions gets a reference to the given int32 and assigns it to the TotalCouponRolledbackRedemptions field.

### GetReferralRedemptions

`func (o *CampaignAnalyticsAllOf) GetReferralRedemptions() int32`

GetReferralRedemptions returns the ReferralRedemptions field if non-nil, zero value otherwise.

### GetReferralRedemptionsOk

`func (o *CampaignAnalyticsAllOf) GetReferralRedemptionsOk() (int32, bool)`

GetReferralRedemptionsOk returns a tuple with the ReferralRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralRedemptions

`func (o *CampaignAnalyticsAllOf) HasReferralRedemptions() bool`

HasReferralRedemptions returns a boolean if a field has been set.

### SetReferralRedemptions

`func (o *CampaignAnalyticsAllOf) SetReferralRedemptions(v int32)`

SetReferralRedemptions gets a reference to the given int32 and assigns it to the ReferralRedemptions field.

### GetTotalReferralRedemptions

`func (o *CampaignAnalyticsAllOf) GetTotalReferralRedemptions() int32`

GetTotalReferralRedemptions returns the TotalReferralRedemptions field if non-nil, zero value otherwise.

### GetTotalReferralRedemptionsOk

`func (o *CampaignAnalyticsAllOf) GetTotalReferralRedemptionsOk() (int32, bool)`

GetTotalReferralRedemptionsOk returns a tuple with the TotalReferralRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalReferralRedemptions

`func (o *CampaignAnalyticsAllOf) HasTotalReferralRedemptions() bool`

HasTotalReferralRedemptions returns a boolean if a field has been set.

### SetTotalReferralRedemptions

`func (o *CampaignAnalyticsAllOf) SetTotalReferralRedemptions(v int32)`

SetTotalReferralRedemptions gets a reference to the given int32 and assigns it to the TotalReferralRedemptions field.

### GetCouponsCreated

`func (o *CampaignAnalyticsAllOf) GetCouponsCreated() int32`

GetCouponsCreated returns the CouponsCreated field if non-nil, zero value otherwise.

### GetCouponsCreatedOk

`func (o *CampaignAnalyticsAllOf) GetCouponsCreatedOk() (int32, bool)`

GetCouponsCreatedOk returns a tuple with the CouponsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponsCreated

`func (o *CampaignAnalyticsAllOf) HasCouponsCreated() bool`

HasCouponsCreated returns a boolean if a field has been set.

### SetCouponsCreated

`func (o *CampaignAnalyticsAllOf) SetCouponsCreated(v int32)`

SetCouponsCreated gets a reference to the given int32 and assigns it to the CouponsCreated field.

### GetTotalCouponsCreated

`func (o *CampaignAnalyticsAllOf) GetTotalCouponsCreated() int32`

GetTotalCouponsCreated returns the TotalCouponsCreated field if non-nil, zero value otherwise.

### GetTotalCouponsCreatedOk

`func (o *CampaignAnalyticsAllOf) GetTotalCouponsCreatedOk() (int32, bool)`

GetTotalCouponsCreatedOk returns a tuple with the TotalCouponsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCouponsCreated

`func (o *CampaignAnalyticsAllOf) HasTotalCouponsCreated() bool`

HasTotalCouponsCreated returns a boolean if a field has been set.

### SetTotalCouponsCreated

`func (o *CampaignAnalyticsAllOf) SetTotalCouponsCreated(v int32)`

SetTotalCouponsCreated gets a reference to the given int32 and assigns it to the TotalCouponsCreated field.

### GetReferralsCreated

`func (o *CampaignAnalyticsAllOf) GetReferralsCreated() int32`

GetReferralsCreated returns the ReferralsCreated field if non-nil, zero value otherwise.

### GetReferralsCreatedOk

`func (o *CampaignAnalyticsAllOf) GetReferralsCreatedOk() (int32, bool)`

GetReferralsCreatedOk returns a tuple with the ReferralsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralsCreated

`func (o *CampaignAnalyticsAllOf) HasReferralsCreated() bool`

HasReferralsCreated returns a boolean if a field has been set.

### SetReferralsCreated

`func (o *CampaignAnalyticsAllOf) SetReferralsCreated(v int32)`

SetReferralsCreated gets a reference to the given int32 and assigns it to the ReferralsCreated field.

### GetTotalReferralsCreated

`func (o *CampaignAnalyticsAllOf) GetTotalReferralsCreated() int32`

GetTotalReferralsCreated returns the TotalReferralsCreated field if non-nil, zero value otherwise.

### GetTotalReferralsCreatedOk

`func (o *CampaignAnalyticsAllOf) GetTotalReferralsCreatedOk() (int32, bool)`

GetTotalReferralsCreatedOk returns a tuple with the TotalReferralsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalReferralsCreated

`func (o *CampaignAnalyticsAllOf) HasTotalReferralsCreated() bool`

HasTotalReferralsCreated returns a boolean if a field has been set.

### SetTotalReferralsCreated

`func (o *CampaignAnalyticsAllOf) SetTotalReferralsCreated(v int32)`

SetTotalReferralsCreated gets a reference to the given int32 and assigns it to the TotalReferralsCreated field.

### GetAddedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) GetAddedLoyaltyPoints() float32`

GetAddedLoyaltyPoints returns the AddedLoyaltyPoints field if non-nil, zero value otherwise.

### GetAddedLoyaltyPointsOk

`func (o *CampaignAnalyticsAllOf) GetAddedLoyaltyPointsOk() (float32, bool)`

GetAddedLoyaltyPointsOk returns a tuple with the AddedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) HasAddedLoyaltyPoints() bool`

HasAddedLoyaltyPoints returns a boolean if a field has been set.

### SetAddedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) SetAddedLoyaltyPoints(v float32)`

SetAddedLoyaltyPoints gets a reference to the given float32 and assigns it to the AddedLoyaltyPoints field.

### GetTotalAddedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) GetTotalAddedLoyaltyPoints() float32`

GetTotalAddedLoyaltyPoints returns the TotalAddedLoyaltyPoints field if non-nil, zero value otherwise.

### GetTotalAddedLoyaltyPointsOk

`func (o *CampaignAnalyticsAllOf) GetTotalAddedLoyaltyPointsOk() (float32, bool)`

GetTotalAddedLoyaltyPointsOk returns a tuple with the TotalAddedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalAddedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) HasTotalAddedLoyaltyPoints() bool`

HasTotalAddedLoyaltyPoints returns a boolean if a field has been set.

### SetTotalAddedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) SetTotalAddedLoyaltyPoints(v float32)`

SetTotalAddedLoyaltyPoints gets a reference to the given float32 and assigns it to the TotalAddedLoyaltyPoints field.

### GetDeductedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) GetDeductedLoyaltyPoints() float32`

GetDeductedLoyaltyPoints returns the DeductedLoyaltyPoints field if non-nil, zero value otherwise.

### GetDeductedLoyaltyPointsOk

`func (o *CampaignAnalyticsAllOf) GetDeductedLoyaltyPointsOk() (float32, bool)`

GetDeductedLoyaltyPointsOk returns a tuple with the DeductedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeductedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) HasDeductedLoyaltyPoints() bool`

HasDeductedLoyaltyPoints returns a boolean if a field has been set.

### SetDeductedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) SetDeductedLoyaltyPoints(v float32)`

SetDeductedLoyaltyPoints gets a reference to the given float32 and assigns it to the DeductedLoyaltyPoints field.

### GetTotalDeductedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) GetTotalDeductedLoyaltyPoints() float32`

GetTotalDeductedLoyaltyPoints returns the TotalDeductedLoyaltyPoints field if non-nil, zero value otherwise.

### GetTotalDeductedLoyaltyPointsOk

`func (o *CampaignAnalyticsAllOf) GetTotalDeductedLoyaltyPointsOk() (float32, bool)`

GetTotalDeductedLoyaltyPointsOk returns a tuple with the TotalDeductedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalDeductedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) HasTotalDeductedLoyaltyPoints() bool`

HasTotalDeductedLoyaltyPoints returns a boolean if a field has been set.

### SetTotalDeductedLoyaltyPoints

`func (o *CampaignAnalyticsAllOf) SetTotalDeductedLoyaltyPoints(v float32)`

SetTotalDeductedLoyaltyPoints gets a reference to the given float32 and assigns it to the TotalDeductedLoyaltyPoints field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


