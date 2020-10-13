# CustomerActivityReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The integration ID for this entity sent to and used in the Talon.One system. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
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

### GetIntegrationId

`func (o *CustomerActivityReport) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerActivityReport) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *CustomerActivityReport) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *CustomerActivityReport) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetCreated

`func (o *CustomerActivityReport) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerActivityReport) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CustomerActivityReport) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CustomerActivityReport) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetName

`func (o *CustomerActivityReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerActivityReport) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CustomerActivityReport) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CustomerActivityReport) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetCustomerId

`func (o *CustomerActivityReport) GetCustomerId() int32`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerActivityReport) GetCustomerIdOk() (int32, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerId

`func (o *CustomerActivityReport) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerId

`func (o *CustomerActivityReport) SetCustomerId(v int32)`

SetCustomerId gets a reference to the given int32 and assigns it to the CustomerId field.

### GetLastActivity

`func (o *CustomerActivityReport) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CustomerActivityReport) GetLastActivityOk() (time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActivity

`func (o *CustomerActivityReport) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### SetLastActivity

`func (o *CustomerActivityReport) SetLastActivity(v time.Time)`

SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.

### GetCouponRedemptions

`func (o *CustomerActivityReport) GetCouponRedemptions() int32`

GetCouponRedemptions returns the CouponRedemptions field if non-nil, zero value otherwise.

### GetCouponRedemptionsOk

`func (o *CustomerActivityReport) GetCouponRedemptionsOk() (int32, bool)`

GetCouponRedemptionsOk returns a tuple with the CouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponRedemptions

`func (o *CustomerActivityReport) HasCouponRedemptions() bool`

HasCouponRedemptions returns a boolean if a field has been set.

### SetCouponRedemptions

`func (o *CustomerActivityReport) SetCouponRedemptions(v int32)`

SetCouponRedemptions gets a reference to the given int32 and assigns it to the CouponRedemptions field.

### GetCouponUseAttempts

`func (o *CustomerActivityReport) GetCouponUseAttempts() int32`

GetCouponUseAttempts returns the CouponUseAttempts field if non-nil, zero value otherwise.

### GetCouponUseAttemptsOk

`func (o *CustomerActivityReport) GetCouponUseAttemptsOk() (int32, bool)`

GetCouponUseAttemptsOk returns a tuple with the CouponUseAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponUseAttempts

`func (o *CustomerActivityReport) HasCouponUseAttempts() bool`

HasCouponUseAttempts returns a boolean if a field has been set.

### SetCouponUseAttempts

`func (o *CustomerActivityReport) SetCouponUseAttempts(v int32)`

SetCouponUseAttempts gets a reference to the given int32 and assigns it to the CouponUseAttempts field.

### GetCouponFailedAttempts

`func (o *CustomerActivityReport) GetCouponFailedAttempts() int32`

GetCouponFailedAttempts returns the CouponFailedAttempts field if non-nil, zero value otherwise.

### GetCouponFailedAttemptsOk

`func (o *CustomerActivityReport) GetCouponFailedAttemptsOk() (int32, bool)`

GetCouponFailedAttemptsOk returns a tuple with the CouponFailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponFailedAttempts

`func (o *CustomerActivityReport) HasCouponFailedAttempts() bool`

HasCouponFailedAttempts returns a boolean if a field has been set.

### SetCouponFailedAttempts

`func (o *CustomerActivityReport) SetCouponFailedAttempts(v int32)`

SetCouponFailedAttempts gets a reference to the given int32 and assigns it to the CouponFailedAttempts field.

### GetAccruedDiscounts

`func (o *CustomerActivityReport) GetAccruedDiscounts() float32`

GetAccruedDiscounts returns the AccruedDiscounts field if non-nil, zero value otherwise.

### GetAccruedDiscountsOk

`func (o *CustomerActivityReport) GetAccruedDiscountsOk() (float32, bool)`

GetAccruedDiscountsOk returns a tuple with the AccruedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccruedDiscounts

`func (o *CustomerActivityReport) HasAccruedDiscounts() bool`

HasAccruedDiscounts returns a boolean if a field has been set.

### SetAccruedDiscounts

`func (o *CustomerActivityReport) SetAccruedDiscounts(v float32)`

SetAccruedDiscounts gets a reference to the given float32 and assigns it to the AccruedDiscounts field.

### GetAccruedRevenue

`func (o *CustomerActivityReport) GetAccruedRevenue() float32`

GetAccruedRevenue returns the AccruedRevenue field if non-nil, zero value otherwise.

### GetAccruedRevenueOk

`func (o *CustomerActivityReport) GetAccruedRevenueOk() (float32, bool)`

GetAccruedRevenueOk returns a tuple with the AccruedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccruedRevenue

`func (o *CustomerActivityReport) HasAccruedRevenue() bool`

HasAccruedRevenue returns a boolean if a field has been set.

### SetAccruedRevenue

`func (o *CustomerActivityReport) SetAccruedRevenue(v float32)`

SetAccruedRevenue gets a reference to the given float32 and assigns it to the AccruedRevenue field.

### GetTotalOrders

`func (o *CustomerActivityReport) GetTotalOrders() int32`

GetTotalOrders returns the TotalOrders field if non-nil, zero value otherwise.

### GetTotalOrdersOk

`func (o *CustomerActivityReport) GetTotalOrdersOk() (int32, bool)`

GetTotalOrdersOk returns a tuple with the TotalOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalOrders

`func (o *CustomerActivityReport) HasTotalOrders() bool`

HasTotalOrders returns a boolean if a field has been set.

### SetTotalOrders

`func (o *CustomerActivityReport) SetTotalOrders(v int32)`

SetTotalOrders gets a reference to the given int32 and assigns it to the TotalOrders field.

### GetTotalOrdersNoCoupon

`func (o *CustomerActivityReport) GetTotalOrdersNoCoupon() int32`

GetTotalOrdersNoCoupon returns the TotalOrdersNoCoupon field if non-nil, zero value otherwise.

### GetTotalOrdersNoCouponOk

`func (o *CustomerActivityReport) GetTotalOrdersNoCouponOk() (int32, bool)`

GetTotalOrdersNoCouponOk returns a tuple with the TotalOrdersNoCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalOrdersNoCoupon

`func (o *CustomerActivityReport) HasTotalOrdersNoCoupon() bool`

HasTotalOrdersNoCoupon returns a boolean if a field has been set.

### SetTotalOrdersNoCoupon

`func (o *CustomerActivityReport) SetTotalOrdersNoCoupon(v int32)`

SetTotalOrdersNoCoupon gets a reference to the given int32 and assigns it to the TotalOrdersNoCoupon field.

### GetCampaignName

`func (o *CustomerActivityReport) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *CustomerActivityReport) GetCampaignNameOk() (string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignName

`func (o *CustomerActivityReport) HasCampaignName() bool`

HasCampaignName returns a boolean if a field has been set.

### SetCampaignName

`func (o *CustomerActivityReport) SetCampaignName(v string)`

SetCampaignName gets a reference to the given string and assigns it to the CampaignName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


