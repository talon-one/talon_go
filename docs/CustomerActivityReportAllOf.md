# CustomerActivityReportAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for this customer profile. | 
**CustomerId** | Pointer to **int32** | The internal Talon.One ID of the customer. | 
**LastActivity** | Pointer to [**time.Time**](time.Time.md) | The last activity of the customer. | [optional] 
**CouponRedemptions** | Pointer to **int32** | Number of coupon redemptions in all customer campaigns. | 
**CouponUseAttempts** | Pointer to **int32** | Number of coupon use attempts in all customer campaigns. | 
**CouponFailedAttempts** | Pointer to **int32** | Number of failed coupon use attempts in all customer campaigns. | 
**AccruedDiscounts** | Pointer to **float32** | Number of accrued discounts in all customer campaigns. | 
**AccruedRevenue** | Pointer to **float32** | Amount of accrued revenue in all customer campaigns. | 
**TotalOrders** | Pointer to **int32** | Number of orders in all customer campaigns. | 
**TotalOrdersNoCoupon** | Pointer to **int32** | Number of orders without coupon used in all customer campaigns. | 
**CampaignName** | Pointer to **string** | The name of the campaign this customer belongs to. | 

## Methods

### GetName

`func (o *CustomerActivityReportAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerActivityReportAllOf) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CustomerActivityReportAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CustomerActivityReportAllOf) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetCustomerId

`func (o *CustomerActivityReportAllOf) GetCustomerId() int32`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerActivityReportAllOf) GetCustomerIdOk() (int32, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerId

`func (o *CustomerActivityReportAllOf) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerId

`func (o *CustomerActivityReportAllOf) SetCustomerId(v int32)`

SetCustomerId gets a reference to the given int32 and assigns it to the CustomerId field.

### GetLastActivity

`func (o *CustomerActivityReportAllOf) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CustomerActivityReportAllOf) GetLastActivityOk() (time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActivity

`func (o *CustomerActivityReportAllOf) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivity

`func (o *CustomerActivityReportAllOf) SetLastActivity(v time.Time)`

SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.

### GetCouponRedemptions

`func (o *CustomerActivityReportAllOf) GetCouponRedemptions() int32`

GetCouponRedemptions returns the CouponRedemptions field if non-nil, zero value otherwise.

### GetCouponRedemptionsOk

`func (o *CustomerActivityReportAllOf) GetCouponRedemptionsOk() (int32, bool)`

GetCouponRedemptionsOk returns a tuple with the CouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRedemptions

`func (o *CustomerActivityReportAllOf) HasCouponRedemptions() bool`

HasCouponRedemptions returns a boolean if a field has been set.

### SetCouponRedemptions

`func (o *CustomerActivityReportAllOf) SetCouponRedemptions(v int32)`

SetCouponRedemptions gets a reference to the given int32 and assigns it to the CouponRedemptions field.

### GetCouponUseAttempts

`func (o *CustomerActivityReportAllOf) GetCouponUseAttempts() int32`

GetCouponUseAttempts returns the CouponUseAttempts field if non-nil, zero value otherwise.

### GetCouponUseAttemptsOk

`func (o *CustomerActivityReportAllOf) GetCouponUseAttemptsOk() (int32, bool)`

GetCouponUseAttemptsOk returns a tuple with the CouponUseAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponUseAttempts

`func (o *CustomerActivityReportAllOf) HasCouponUseAttempts() bool`

HasCouponUseAttempts returns a boolean if a field has been set.

### SetCouponUseAttempts

`func (o *CustomerActivityReportAllOf) SetCouponUseAttempts(v int32)`

SetCouponUseAttempts gets a reference to the given int32 and assigns it to the CouponUseAttempts field.

### GetCouponFailedAttempts

`func (o *CustomerActivityReportAllOf) GetCouponFailedAttempts() int32`

GetCouponFailedAttempts returns the CouponFailedAttempts field if non-nil, zero value otherwise.

### GetCouponFailedAttemptsOk

`func (o *CustomerActivityReportAllOf) GetCouponFailedAttemptsOk() (int32, bool)`

GetCouponFailedAttemptsOk returns a tuple with the CouponFailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponFailedAttempts

`func (o *CustomerActivityReportAllOf) HasCouponFailedAttempts() bool`

HasCouponFailedAttempts returns a boolean if a field has been set.

### SetCouponFailedAttempts

`func (o *CustomerActivityReportAllOf) SetCouponFailedAttempts(v int32)`

SetCouponFailedAttempts gets a reference to the given int32 and assigns it to the CouponFailedAttempts field.

### GetAccruedDiscounts

`func (o *CustomerActivityReportAllOf) GetAccruedDiscounts() float32`

GetAccruedDiscounts returns the AccruedDiscounts field if non-nil, zero value otherwise.

### GetAccruedDiscountsOk

`func (o *CustomerActivityReportAllOf) GetAccruedDiscountsOk() (float32, bool)`

GetAccruedDiscountsOk returns a tuple with the AccruedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccruedDiscounts

`func (o *CustomerActivityReportAllOf) HasAccruedDiscounts() bool`

HasAccruedDiscounts returns a boolean if a field has been set.

### SetAccruedDiscounts

`func (o *CustomerActivityReportAllOf) SetAccruedDiscounts(v float32)`

SetAccruedDiscounts gets a reference to the given float32 and assigns it to the AccruedDiscounts field.

### GetAccruedRevenue

`func (o *CustomerActivityReportAllOf) GetAccruedRevenue() float32`

GetAccruedRevenue returns the AccruedRevenue field if non-nil, zero value otherwise.

### GetAccruedRevenueOk

`func (o *CustomerActivityReportAllOf) GetAccruedRevenueOk() (float32, bool)`

GetAccruedRevenueOk returns a tuple with the AccruedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccruedRevenue

`func (o *CustomerActivityReportAllOf) HasAccruedRevenue() bool`

HasAccruedRevenue returns a boolean if a field has been set.

### SetAccruedRevenue

`func (o *CustomerActivityReportAllOf) SetAccruedRevenue(v float32)`

SetAccruedRevenue gets a reference to the given float32 and assigns it to the AccruedRevenue field.

### GetTotalOrders

`func (o *CustomerActivityReportAllOf) GetTotalOrders() int32`

GetTotalOrders returns the TotalOrders field if non-nil, zero value otherwise.

### GetTotalOrdersOk

`func (o *CustomerActivityReportAllOf) GetTotalOrdersOk() (int32, bool)`

GetTotalOrdersOk returns a tuple with the TotalOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalOrders

`func (o *CustomerActivityReportAllOf) HasTotalOrders() bool`

HasTotalOrders returns a boolean if a field has been set.

### SetTotalOrders

`func (o *CustomerActivityReportAllOf) SetTotalOrders(v int32)`

SetTotalOrders gets a reference to the given int32 and assigns it to the TotalOrders field.

### GetTotalOrdersNoCoupon

`func (o *CustomerActivityReportAllOf) GetTotalOrdersNoCoupon() int32`

GetTotalOrdersNoCoupon returns the TotalOrdersNoCoupon field if non-nil, zero value otherwise.

### GetTotalOrdersNoCouponOk

`func (o *CustomerActivityReportAllOf) GetTotalOrdersNoCouponOk() (int32, bool)`

GetTotalOrdersNoCouponOk returns a tuple with the TotalOrdersNoCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalOrdersNoCoupon

`func (o *CustomerActivityReportAllOf) HasTotalOrdersNoCoupon() bool`

HasTotalOrdersNoCoupon returns a boolean if a field has been set.

### SetTotalOrdersNoCoupon

`func (o *CustomerActivityReportAllOf) SetTotalOrdersNoCoupon(v int32)`

SetTotalOrdersNoCoupon gets a reference to the given int32 and assigns it to the TotalOrdersNoCoupon field.

### GetCampaignName

`func (o *CustomerActivityReportAllOf) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *CustomerActivityReportAllOf) GetCampaignNameOk() (string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignName

`func (o *CustomerActivityReportAllOf) HasCampaignName() bool`

HasCampaignName returns a boolean if a field has been set.

### SetCampaignName

`func (o *CustomerActivityReportAllOf) SetCampaignName(v string)`

SetCampaignName gets a reference to the given string and assigns it to the CampaignName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


