# CustomerActivityReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The integration ID set by your integration layer. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Name** | Pointer to **string** | The name for this customer profile. | 
**CustomerId** | Pointer to **int64** | The internal Talon.One ID of the customer. | 
**LastActivity** | Pointer to [**time.Time**](time.Time.md) | The last activity of the customer. | [optional] 
**CouponRedemptions** | Pointer to **int64** | Number of coupon redemptions in all customer campaigns. | 
**CouponUseAttempts** | Pointer to **int64** | Number of coupon use attempts in all customer campaigns. | 
**CouponFailedAttempts** | Pointer to **int64** | Number of failed coupon use attempts in all customer campaigns. | 
**AccruedDiscounts** | Pointer to **float32** | Number of accrued discounts in all customer campaigns. | 
**AccruedRevenue** | Pointer to **float32** | Amount of accrued revenue in all customer campaigns. | 
**TotalOrders** | Pointer to **int64** | Number of orders in all customer campaigns. | 
**TotalOrdersNoCoupon** | Pointer to **int64** | Number of orders without coupon used in all customer campaigns. | 
**CampaignName** | Pointer to **string** | The name of the campaign this customer belongs to. | 

## Methods

### NewCustomerActivityReport

`func NewCustomerActivityReport(integrationId string, created time.Time, name string, customerId int64, couponRedemptions int64, couponUseAttempts int64, couponFailedAttempts int64, accruedDiscounts float32, accruedRevenue float32, totalOrders int64, totalOrdersNoCoupon int64, campaignName string, ) *CustomerActivityReport`

NewCustomerActivityReport instantiates a new CustomerActivityReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerActivityReportWithDefaults

`func NewCustomerActivityReportWithDefaults() *CustomerActivityReport`

NewCustomerActivityReportWithDefaults instantiates a new CustomerActivityReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *CustomerActivityReport) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerActivityReport) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CustomerActivityReport) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetCreated

`func (o *CustomerActivityReport) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerActivityReport) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerActivityReport) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetName

`func (o *CustomerActivityReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerActivityReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerActivityReport) SetName(v string)`

SetName sets Name field to given value.


### GetCustomerId

`func (o *CustomerActivityReport) GetCustomerId() int64`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerActivityReport) GetCustomerIdOk() (*int64, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerActivityReport) SetCustomerId(v int64)`

SetCustomerId sets CustomerId field to given value.


### GetLastActivity

`func (o *CustomerActivityReport) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CustomerActivityReport) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *CustomerActivityReport) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *CustomerActivityReport) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetCouponRedemptions

`func (o *CustomerActivityReport) GetCouponRedemptions() int64`

GetCouponRedemptions returns the CouponRedemptions field if non-nil, zero value otherwise.

### GetCouponRedemptionsOk

`func (o *CustomerActivityReport) GetCouponRedemptionsOk() (*int64, bool)`

GetCouponRedemptionsOk returns a tuple with the CouponRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRedemptions

`func (o *CustomerActivityReport) SetCouponRedemptions(v int64)`

SetCouponRedemptions sets CouponRedemptions field to given value.


### GetCouponUseAttempts

`func (o *CustomerActivityReport) GetCouponUseAttempts() int64`

GetCouponUseAttempts returns the CouponUseAttempts field if non-nil, zero value otherwise.

### GetCouponUseAttemptsOk

`func (o *CustomerActivityReport) GetCouponUseAttemptsOk() (*int64, bool)`

GetCouponUseAttemptsOk returns a tuple with the CouponUseAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponUseAttempts

`func (o *CustomerActivityReport) SetCouponUseAttempts(v int64)`

SetCouponUseAttempts sets CouponUseAttempts field to given value.


### GetCouponFailedAttempts

`func (o *CustomerActivityReport) GetCouponFailedAttempts() int64`

GetCouponFailedAttempts returns the CouponFailedAttempts field if non-nil, zero value otherwise.

### GetCouponFailedAttemptsOk

`func (o *CustomerActivityReport) GetCouponFailedAttemptsOk() (*int64, bool)`

GetCouponFailedAttemptsOk returns a tuple with the CouponFailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponFailedAttempts

`func (o *CustomerActivityReport) SetCouponFailedAttempts(v int64)`

SetCouponFailedAttempts sets CouponFailedAttempts field to given value.


### GetAccruedDiscounts

`func (o *CustomerActivityReport) GetAccruedDiscounts() float32`

GetAccruedDiscounts returns the AccruedDiscounts field if non-nil, zero value otherwise.

### GetAccruedDiscountsOk

`func (o *CustomerActivityReport) GetAccruedDiscountsOk() (*float32, bool)`

GetAccruedDiscountsOk returns a tuple with the AccruedDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedDiscounts

`func (o *CustomerActivityReport) SetAccruedDiscounts(v float32)`

SetAccruedDiscounts sets AccruedDiscounts field to given value.


### GetAccruedRevenue

`func (o *CustomerActivityReport) GetAccruedRevenue() float32`

GetAccruedRevenue returns the AccruedRevenue field if non-nil, zero value otherwise.

### GetAccruedRevenueOk

`func (o *CustomerActivityReport) GetAccruedRevenueOk() (*float32, bool)`

GetAccruedRevenueOk returns a tuple with the AccruedRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedRevenue

`func (o *CustomerActivityReport) SetAccruedRevenue(v float32)`

SetAccruedRevenue sets AccruedRevenue field to given value.


### GetTotalOrders

`func (o *CustomerActivityReport) GetTotalOrders() int64`

GetTotalOrders returns the TotalOrders field if non-nil, zero value otherwise.

### GetTotalOrdersOk

`func (o *CustomerActivityReport) GetTotalOrdersOk() (*int64, bool)`

GetTotalOrdersOk returns a tuple with the TotalOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrders

`func (o *CustomerActivityReport) SetTotalOrders(v int64)`

SetTotalOrders sets TotalOrders field to given value.


### GetTotalOrdersNoCoupon

`func (o *CustomerActivityReport) GetTotalOrdersNoCoupon() int64`

GetTotalOrdersNoCoupon returns the TotalOrdersNoCoupon field if non-nil, zero value otherwise.

### GetTotalOrdersNoCouponOk

`func (o *CustomerActivityReport) GetTotalOrdersNoCouponOk() (*int64, bool)`

GetTotalOrdersNoCouponOk returns a tuple with the TotalOrdersNoCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrdersNoCoupon

`func (o *CustomerActivityReport) SetTotalOrdersNoCoupon(v int64)`

SetTotalOrdersNoCoupon sets TotalOrdersNoCoupon field to given value.


### GetCampaignName

`func (o *CustomerActivityReport) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *CustomerActivityReport) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *CustomerActivityReport) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


