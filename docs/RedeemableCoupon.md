# RedeemableCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponId** | Pointer to **int64** | The internal ID of the coupon. | 
**CouponCode** | Pointer to **string** | The coupon code. | 
**UsageCounter** | Pointer to **int64** | The number of times the coupon has been successfully redeemed. | 
**UsageLimit** | Pointer to **int64** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits still apply. | 
**CampaignName** | Pointer to **string** | The name of the campaign that owns the coupon. | 

## Methods

### NewRedeemableCoupon

`func NewRedeemableCoupon(couponId int64, couponCode string, usageCounter int64, usageLimit int64, campaignName string, ) *RedeemableCoupon`

NewRedeemableCoupon instantiates a new RedeemableCoupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemableCouponWithDefaults

`func NewRedeemableCouponWithDefaults() *RedeemableCoupon`

NewRedeemableCouponWithDefaults instantiates a new RedeemableCoupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponId

`func (o *RedeemableCoupon) GetCouponId() int64`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *RedeemableCoupon) GetCouponIdOk() (*int64, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *RedeemableCoupon) SetCouponId(v int64)`

SetCouponId sets CouponId field to given value.


### GetCouponCode

`func (o *RedeemableCoupon) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *RedeemableCoupon) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *RedeemableCoupon) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.


### GetUsageCounter

`func (o *RedeemableCoupon) GetUsageCounter() int64`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *RedeemableCoupon) GetUsageCounterOk() (*int64, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCounter

`func (o *RedeemableCoupon) SetUsageCounter(v int64)`

SetUsageCounter sets UsageCounter field to given value.


### GetUsageLimit

`func (o *RedeemableCoupon) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *RedeemableCoupon) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *RedeemableCoupon) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetCampaignName

`func (o *RedeemableCoupon) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *RedeemableCoupon) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *RedeemableCoupon) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


