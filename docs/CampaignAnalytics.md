# CampaignAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedLoyaltyPoints** | Pointer to **float32** | Number of added loyalty points in the campaign in a specific interval. | 
**CampaignDiscountCosts** | Pointer to **float32** | Amount of cost caused by discounts given in the campaign. | 
**CampaignFreeItems** | Pointer to **int32** | Amount of free items given in the campaign. | 
**CampaignRefund** | Pointer to **float32** | Amount of refunds in this campaign (for coupon or discount sessions). | 
**CampaignRefundedDiscounts** | Pointer to **float32** | Amount of discounts rolledback due to refund in the campaign. | 
**CampaignRevenue** | Pointer to **float32** | Amount of revenue in this campaign (for coupon or discount sessions). | 
**CouponRedemptions** | Pointer to **int32** | Number of coupon redemptions in the campaign. | 
**CouponRolledbackRedemptions** | Pointer to **int32** | Number of coupon redemptions that have been rolled back (due to canceling closed session) in the campaign. | 
**CouponsCreated** | Pointer to **int32** | Number of coupons created in the campaign by the rule engine. | 
**Date** | Pointer to [**time.Time**](time.Time.md) |  | 
**DeductedLoyaltyPoints** | Pointer to **float32** | Number of deducted loyalty points in the campaign in a specific interval. | 
**ReferralRedemptions** | Pointer to **int32** | Number of referral redemptions in the campaign. | 
**ReferralsCreated** | Pointer to **int32** | Number of referrals created in the campaign by the rule engine. | 
**TotalAddedLoyaltyPoints** | Pointer to **float32** | Number of added loyalty points in the campaign since it began. | 
**TotalCampaignDiscountCosts** | Pointer to **float32** | Amount of cost caused by discounts given in the campaign since it began. | 
**TotalCampaignFreeItems** | Pointer to **int32** | Amount of free items given in the campaign since it began. | 
**TotalCampaignRefund** | Pointer to **float32** | Amount of refunds in this campaign since it began (for coupon or discount sessions). | 
**TotalCampaignRefundedDiscounts** | Pointer to **float32** | Amount of discounts rolledback due to refund in the campaign since it began. | 
**TotalCampaignRevenue** | Pointer to **float32** | Amount of revenue in this campaign since it began (for coupon or discount sessions). | 
**TotalCouponRedemptions** | Pointer to **int32** | Number of coupon redemptions in the campaign since it began. | 
**TotalCouponRolledbackRedemptions** | Pointer to **int32** | Number of coupon redemptions that have been rolled back (due to canceling closed session) in the campaign since it began. | 
**TotalCouponsCreated** | Pointer to **int32** | Number of coupons created in the campaign by the rule engine since it began. | 
**TotalDeductedLoyaltyPoints** | Pointer to **float32** | Number of deducted loyalty points in the campaign since it began. | 
**TotalReferralRedemptions** | Pointer to **int32** | Number of referral redemptions in the campaign since it began. | 
**TotalReferralsCreated** | Pointer to **int32** | Number of referrals created in the campaign by the rule engine since it began. | 

## Methods

### GetAddedLoyaltyPoints

`func (o *CampaignAnalytics) GetAddedLoyaltyPoints() float32`

GetAddedLoyaltyPoints returns the AddedLoyaltyPoints field if non-nil, zero value otherwise.

### GetAddedLoyaltyPointsOk

`func (o *CampaignAnalytics) GetAddedLoyaltyPointsOk() (float32, bool)`

GetAddedLoyaltyPointsOk returns a tuple with the AddedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddedLoyaltyPoints

`func (o *CampaignAnalytics) HasAddedLoyaltyPoints() bool`

HasAddedLoyaltyPoints returns a boolean if a field has been set.

### SetAddedLoyaltyPoints

`func (o *CampaignAnalytics) SetAddedLoyaltyPoints(v float32)`

SetAddedLoyaltyPoints gets a reference to the given float32 and assigns it to the AddedLoyaltyPoints field.

### GetCampaignDiscountCosts

`func (o *CampaignAnalytics) GetCampaignDiscountCosts() float32`

GetCampaignDiscountCosts returns the CampaignDiscountCosts field if non-nil, zero value otherwise.

### GetCampaignDiscountCostsOk

`func (o *CampaignAnalytics) GetCampaignDiscountCostsOk() (float32, bool)`

GetCampaignDiscountCostsOk returns a tuple with the CampaignDiscountCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignDiscountCosts

`func (o *CampaignAnalytics) HasCampaignDiscountCosts() bool`

HasCampaignDiscountCosts returns a boolean if a field has been set.

### SetCampaignDiscountCosts

`func (o *CampaignAnalytics) SetCampaignDiscountCosts(v float32)`

SetCampaignDiscountCosts gets a reference to the given float32 and assigns it to the CampaignDiscountCosts field.

### GetCampaignFreeItems

`func (o *CampaignAnalytics) GetCampaignFreeItems() int32`

GetCampaignFreeItems returns the CampaignFreeItems field if non-nil, zero value otherwise.

### GetCampaignFreeItemsOk

`func (o *CampaignAnalytics) GetCampaignFreeItemsOk() (int32, bool)`

GetCampaignFreeItemsOk returns a tuple with the CampaignFreeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignFreeItems

`func (o *CampaignAnalytics) HasCampaignFreeItems() bool`

HasCampaignFreeItems returns a boolean if a field has been set.

### SetCampaignFreeItems

`func (o *CampaignAnalytics) SetCampaignFreeItems(v int32)`

SetCampaignFreeItems gets a reference to the given int32 and assigns it to the CampaignFreeItems field.

### GetCampaignRefund

`func (o *CampaignAnalytics) GetCampaignRefund() float32`

GetCampaignRefund returns the CampaignRefund field if non-nil, zero value otherwise.

### GetCampaignRefundOk

`func (o *CampaignAnalytics) GetCampaignRefundOk() (float32, bool)`

GetCampaignRefundOk returns a tuple with the CampaignRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignRefund

`func (o *CampaignAnalytics) HasCampaignRefund() bool`

HasCampaignRefund returns a boolean if a field has been set.

### SetCampaignRefund

`func (o *CampaignAnalytics) SetCampaignRefund(v float32)`

SetCampaignRefund gets a reference to the given float32 and assigns it to the CampaignRefund field.

### GetCampaignRefundedDiscounts

`func (o *CampaignAnalytics) GetCampaignRefundedDiscounts() float32`

GetCampaignRefundedDiscounts returns the CampaignRefundedDiscounts field if non-nil, zero value otherwise.

### GetCampaignRefundedDiscountsOk

`func (o *CampaignAnalytics) GetCampaignRefundedDiscountsOk() (float32, bool)`

GetCampaignRefundedDiscountsOk returns a tuple with the CampaignRefundedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignRefundedDiscounts

`func (o *CampaignAnalytics) HasCampaignRefundedDiscounts() bool`

HasCampaignRefundedDiscounts returns a boolean if a field has been set.

### SetCampaignRefundedDiscounts

`func (o *CampaignAnalytics) SetCampaignRefundedDiscounts(v float32)`

SetCampaignRefundedDiscounts gets a reference to the given float32 and assigns it to the CampaignRefundedDiscounts field.

### GetCampaignRevenue

`func (o *CampaignAnalytics) GetCampaignRevenue() float32`

GetCampaignRevenue returns the CampaignRevenue field if non-nil, zero value otherwise.

### GetCampaignRevenueOk

`func (o *CampaignAnalytics) GetCampaignRevenueOk() (float32, bool)`

GetCampaignRevenueOk returns a tuple with the CampaignRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignRevenue

`func (o *CampaignAnalytics) HasCampaignRevenue() bool`

HasCampaignRevenue returns a boolean if a field has been set.

### SetCampaignRevenue

`func (o *CampaignAnalytics) SetCampaignRevenue(v float32)`

SetCampaignRevenue gets a reference to the given float32 and assigns it to the CampaignRevenue field.

### GetCouponRedemptions

`func (o *CampaignAnalytics) GetCouponRedemptions() int32`

GetCouponRedemptions returns the CouponRedemptions field if non-nil, zero value otherwise.

### GetCouponRedemptionsOk

`func (o *CampaignAnalytics) GetCouponRedemptionsOk() (int32, bool)`

GetCouponRedemptionsOk returns a tuple with the CouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRedemptions

`func (o *CampaignAnalytics) HasCouponRedemptions() bool`

HasCouponRedemptions returns a boolean if a field has been set.

### SetCouponRedemptions

`func (o *CampaignAnalytics) SetCouponRedemptions(v int32)`

SetCouponRedemptions gets a reference to the given int32 and assigns it to the CouponRedemptions field.

### GetCouponRolledbackRedemptions

`func (o *CampaignAnalytics) GetCouponRolledbackRedemptions() int32`

GetCouponRolledbackRedemptions returns the CouponRolledbackRedemptions field if non-nil, zero value otherwise.

### GetCouponRolledbackRedemptionsOk

`func (o *CampaignAnalytics) GetCouponRolledbackRedemptionsOk() (int32, bool)`

GetCouponRolledbackRedemptionsOk returns a tuple with the CouponRolledbackRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRolledbackRedemptions

`func (o *CampaignAnalytics) HasCouponRolledbackRedemptions() bool`

HasCouponRolledbackRedemptions returns a boolean if a field has been set.

### SetCouponRolledbackRedemptions

`func (o *CampaignAnalytics) SetCouponRolledbackRedemptions(v int32)`

SetCouponRolledbackRedemptions gets a reference to the given int32 and assigns it to the CouponRolledbackRedemptions field.

### GetCouponsCreated

`func (o *CampaignAnalytics) GetCouponsCreated() int32`

GetCouponsCreated returns the CouponsCreated field if non-nil, zero value otherwise.

### GetCouponsCreatedOk

`func (o *CampaignAnalytics) GetCouponsCreatedOk() (int32, bool)`

GetCouponsCreatedOk returns a tuple with the CouponsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponsCreated

`func (o *CampaignAnalytics) HasCouponsCreated() bool`

HasCouponsCreated returns a boolean if a field has been set.

### SetCouponsCreated

`func (o *CampaignAnalytics) SetCouponsCreated(v int32)`

SetCouponsCreated gets a reference to the given int32 and assigns it to the CouponsCreated field.

### GetDate

`func (o *CampaignAnalytics) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CampaignAnalytics) GetDateOk() (time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDate

`func (o *CampaignAnalytics) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDate

`func (o *CampaignAnalytics) SetDate(v time.Time)`

SetDate gets a reference to the given time.Time and assigns it to the Date field.

### GetDeductedLoyaltyPoints

`func (o *CampaignAnalytics) GetDeductedLoyaltyPoints() float32`

GetDeductedLoyaltyPoints returns the DeductedLoyaltyPoints field if non-nil, zero value otherwise.

### GetDeductedLoyaltyPointsOk

`func (o *CampaignAnalytics) GetDeductedLoyaltyPointsOk() (float32, bool)`

GetDeductedLoyaltyPointsOk returns a tuple with the DeductedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeductedLoyaltyPoints

`func (o *CampaignAnalytics) HasDeductedLoyaltyPoints() bool`

HasDeductedLoyaltyPoints returns a boolean if a field has been set.

### SetDeductedLoyaltyPoints

`func (o *CampaignAnalytics) SetDeductedLoyaltyPoints(v float32)`

SetDeductedLoyaltyPoints gets a reference to the given float32 and assigns it to the DeductedLoyaltyPoints field.

### GetReferralRedemptions

`func (o *CampaignAnalytics) GetReferralRedemptions() int32`

GetReferralRedemptions returns the ReferralRedemptions field if non-nil, zero value otherwise.

### GetReferralRedemptionsOk

`func (o *CampaignAnalytics) GetReferralRedemptionsOk() (int32, bool)`

GetReferralRedemptionsOk returns a tuple with the ReferralRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralRedemptions

`func (o *CampaignAnalytics) HasReferralRedemptions() bool`

HasReferralRedemptions returns a boolean if a field has been set.

### SetReferralRedemptions

`func (o *CampaignAnalytics) SetReferralRedemptions(v int32)`

SetReferralRedemptions gets a reference to the given int32 and assigns it to the ReferralRedemptions field.

### GetReferralsCreated

`func (o *CampaignAnalytics) GetReferralsCreated() int32`

GetReferralsCreated returns the ReferralsCreated field if non-nil, zero value otherwise.

### GetReferralsCreatedOk

`func (o *CampaignAnalytics) GetReferralsCreatedOk() (int32, bool)`

GetReferralsCreatedOk returns a tuple with the ReferralsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralsCreated

`func (o *CampaignAnalytics) HasReferralsCreated() bool`

HasReferralsCreated returns a boolean if a field has been set.

### SetReferralsCreated

`func (o *CampaignAnalytics) SetReferralsCreated(v int32)`

SetReferralsCreated gets a reference to the given int32 and assigns it to the ReferralsCreated field.

### GetTotalAddedLoyaltyPoints

`func (o *CampaignAnalytics) GetTotalAddedLoyaltyPoints() float32`

GetTotalAddedLoyaltyPoints returns the TotalAddedLoyaltyPoints field if non-nil, zero value otherwise.

### GetTotalAddedLoyaltyPointsOk

`func (o *CampaignAnalytics) GetTotalAddedLoyaltyPointsOk() (float32, bool)`

GetTotalAddedLoyaltyPointsOk returns a tuple with the TotalAddedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalAddedLoyaltyPoints

`func (o *CampaignAnalytics) HasTotalAddedLoyaltyPoints() bool`

HasTotalAddedLoyaltyPoints returns a boolean if a field has been set.

### SetTotalAddedLoyaltyPoints

`func (o *CampaignAnalytics) SetTotalAddedLoyaltyPoints(v float32)`

SetTotalAddedLoyaltyPoints gets a reference to the given float32 and assigns it to the TotalAddedLoyaltyPoints field.

### GetTotalCampaignDiscountCosts

`func (o *CampaignAnalytics) GetTotalCampaignDiscountCosts() float32`

GetTotalCampaignDiscountCosts returns the TotalCampaignDiscountCosts field if non-nil, zero value otherwise.

### GetTotalCampaignDiscountCostsOk

`func (o *CampaignAnalytics) GetTotalCampaignDiscountCostsOk() (float32, bool)`

GetTotalCampaignDiscountCostsOk returns a tuple with the TotalCampaignDiscountCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignDiscountCosts

`func (o *CampaignAnalytics) HasTotalCampaignDiscountCosts() bool`

HasTotalCampaignDiscountCosts returns a boolean if a field has been set.

### SetTotalCampaignDiscountCosts

`func (o *CampaignAnalytics) SetTotalCampaignDiscountCosts(v float32)`

SetTotalCampaignDiscountCosts gets a reference to the given float32 and assigns it to the TotalCampaignDiscountCosts field.

### GetTotalCampaignFreeItems

`func (o *CampaignAnalytics) GetTotalCampaignFreeItems() int32`

GetTotalCampaignFreeItems returns the TotalCampaignFreeItems field if non-nil, zero value otherwise.

### GetTotalCampaignFreeItemsOk

`func (o *CampaignAnalytics) GetTotalCampaignFreeItemsOk() (int32, bool)`

GetTotalCampaignFreeItemsOk returns a tuple with the TotalCampaignFreeItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignFreeItems

`func (o *CampaignAnalytics) HasTotalCampaignFreeItems() bool`

HasTotalCampaignFreeItems returns a boolean if a field has been set.

### SetTotalCampaignFreeItems

`func (o *CampaignAnalytics) SetTotalCampaignFreeItems(v int32)`

SetTotalCampaignFreeItems gets a reference to the given int32 and assigns it to the TotalCampaignFreeItems field.

### GetTotalCampaignRefund

`func (o *CampaignAnalytics) GetTotalCampaignRefund() float32`

GetTotalCampaignRefund returns the TotalCampaignRefund field if non-nil, zero value otherwise.

### GetTotalCampaignRefundOk

`func (o *CampaignAnalytics) GetTotalCampaignRefundOk() (float32, bool)`

GetTotalCampaignRefundOk returns a tuple with the TotalCampaignRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignRefund

`func (o *CampaignAnalytics) HasTotalCampaignRefund() bool`

HasTotalCampaignRefund returns a boolean if a field has been set.

### SetTotalCampaignRefund

`func (o *CampaignAnalytics) SetTotalCampaignRefund(v float32)`

SetTotalCampaignRefund gets a reference to the given float32 and assigns it to the TotalCampaignRefund field.

### GetTotalCampaignRefundedDiscounts

`func (o *CampaignAnalytics) GetTotalCampaignRefundedDiscounts() float32`

GetTotalCampaignRefundedDiscounts returns the TotalCampaignRefundedDiscounts field if non-nil, zero value otherwise.

### GetTotalCampaignRefundedDiscountsOk

`func (o *CampaignAnalytics) GetTotalCampaignRefundedDiscountsOk() (float32, bool)`

GetTotalCampaignRefundedDiscountsOk returns a tuple with the TotalCampaignRefundedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignRefundedDiscounts

`func (o *CampaignAnalytics) HasTotalCampaignRefundedDiscounts() bool`

HasTotalCampaignRefundedDiscounts returns a boolean if a field has been set.

### SetTotalCampaignRefundedDiscounts

`func (o *CampaignAnalytics) SetTotalCampaignRefundedDiscounts(v float32)`

SetTotalCampaignRefundedDiscounts gets a reference to the given float32 and assigns it to the TotalCampaignRefundedDiscounts field.

### GetTotalCampaignRevenue

`func (o *CampaignAnalytics) GetTotalCampaignRevenue() float32`

GetTotalCampaignRevenue returns the TotalCampaignRevenue field if non-nil, zero value otherwise.

### GetTotalCampaignRevenueOk

`func (o *CampaignAnalytics) GetTotalCampaignRevenueOk() (float32, bool)`

GetTotalCampaignRevenueOk returns a tuple with the TotalCampaignRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCampaignRevenue

`func (o *CampaignAnalytics) HasTotalCampaignRevenue() bool`

HasTotalCampaignRevenue returns a boolean if a field has been set.

### SetTotalCampaignRevenue

`func (o *CampaignAnalytics) SetTotalCampaignRevenue(v float32)`

SetTotalCampaignRevenue gets a reference to the given float32 and assigns it to the TotalCampaignRevenue field.

### GetTotalCouponRedemptions

`func (o *CampaignAnalytics) GetTotalCouponRedemptions() int32`

GetTotalCouponRedemptions returns the TotalCouponRedemptions field if non-nil, zero value otherwise.

### GetTotalCouponRedemptionsOk

`func (o *CampaignAnalytics) GetTotalCouponRedemptionsOk() (int32, bool)`

GetTotalCouponRedemptionsOk returns a tuple with the TotalCouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCouponRedemptions

`func (o *CampaignAnalytics) HasTotalCouponRedemptions() bool`

HasTotalCouponRedemptions returns a boolean if a field has been set.

### SetTotalCouponRedemptions

`func (o *CampaignAnalytics) SetTotalCouponRedemptions(v int32)`

SetTotalCouponRedemptions gets a reference to the given int32 and assigns it to the TotalCouponRedemptions field.

### GetTotalCouponRolledbackRedemptions

`func (o *CampaignAnalytics) GetTotalCouponRolledbackRedemptions() int32`

GetTotalCouponRolledbackRedemptions returns the TotalCouponRolledbackRedemptions field if non-nil, zero value otherwise.

### GetTotalCouponRolledbackRedemptionsOk

`func (o *CampaignAnalytics) GetTotalCouponRolledbackRedemptionsOk() (int32, bool)`

GetTotalCouponRolledbackRedemptionsOk returns a tuple with the TotalCouponRolledbackRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCouponRolledbackRedemptions

`func (o *CampaignAnalytics) HasTotalCouponRolledbackRedemptions() bool`

HasTotalCouponRolledbackRedemptions returns a boolean if a field has been set.

### SetTotalCouponRolledbackRedemptions

`func (o *CampaignAnalytics) SetTotalCouponRolledbackRedemptions(v int32)`

SetTotalCouponRolledbackRedemptions gets a reference to the given int32 and assigns it to the TotalCouponRolledbackRedemptions field.

### GetTotalCouponsCreated

`func (o *CampaignAnalytics) GetTotalCouponsCreated() int32`

GetTotalCouponsCreated returns the TotalCouponsCreated field if non-nil, zero value otherwise.

### GetTotalCouponsCreatedOk

`func (o *CampaignAnalytics) GetTotalCouponsCreatedOk() (int32, bool)`

GetTotalCouponsCreatedOk returns a tuple with the TotalCouponsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalCouponsCreated

`func (o *CampaignAnalytics) HasTotalCouponsCreated() bool`

HasTotalCouponsCreated returns a boolean if a field has been set.

### SetTotalCouponsCreated

`func (o *CampaignAnalytics) SetTotalCouponsCreated(v int32)`

SetTotalCouponsCreated gets a reference to the given int32 and assigns it to the TotalCouponsCreated field.

### GetTotalDeductedLoyaltyPoints

`func (o *CampaignAnalytics) GetTotalDeductedLoyaltyPoints() float32`

GetTotalDeductedLoyaltyPoints returns the TotalDeductedLoyaltyPoints field if non-nil, zero value otherwise.

### GetTotalDeductedLoyaltyPointsOk

`func (o *CampaignAnalytics) GetTotalDeductedLoyaltyPointsOk() (float32, bool)`

GetTotalDeductedLoyaltyPointsOk returns a tuple with the TotalDeductedLoyaltyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalDeductedLoyaltyPoints

`func (o *CampaignAnalytics) HasTotalDeductedLoyaltyPoints() bool`

HasTotalDeductedLoyaltyPoints returns a boolean if a field has been set.

### SetTotalDeductedLoyaltyPoints

`func (o *CampaignAnalytics) SetTotalDeductedLoyaltyPoints(v float32)`

SetTotalDeductedLoyaltyPoints gets a reference to the given float32 and assigns it to the TotalDeductedLoyaltyPoints field.

### GetTotalReferralRedemptions

`func (o *CampaignAnalytics) GetTotalReferralRedemptions() int32`

GetTotalReferralRedemptions returns the TotalReferralRedemptions field if non-nil, zero value otherwise.

### GetTotalReferralRedemptionsOk

`func (o *CampaignAnalytics) GetTotalReferralRedemptionsOk() (int32, bool)`

GetTotalReferralRedemptionsOk returns a tuple with the TotalReferralRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalReferralRedemptions

`func (o *CampaignAnalytics) HasTotalReferralRedemptions() bool`

HasTotalReferralRedemptions returns a boolean if a field has been set.

### SetTotalReferralRedemptions

`func (o *CampaignAnalytics) SetTotalReferralRedemptions(v int32)`

SetTotalReferralRedemptions gets a reference to the given int32 and assigns it to the TotalReferralRedemptions field.

### GetTotalReferralsCreated

`func (o *CampaignAnalytics) GetTotalReferralsCreated() int32`

GetTotalReferralsCreated returns the TotalReferralsCreated field if non-nil, zero value otherwise.

### GetTotalReferralsCreatedOk

`func (o *CampaignAnalytics) GetTotalReferralsCreatedOk() (int32, bool)`

GetTotalReferralsCreatedOk returns a tuple with the TotalReferralsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalReferralsCreated

`func (o *CampaignAnalytics) HasTotalReferralsCreated() bool`

HasTotalReferralsCreated returns a boolean if a field has been set.

### SetTotalReferralsCreated

`func (o *CampaignAnalytics) SetTotalReferralsCreated(v int32)`

SetTotalReferralsCreated gets a reference to the given int32 and assigns it to the TotalReferralsCreated field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


