# ExpiringCouponsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponValue** | Pointer to **string** | The coupon code. | 
**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon was created. | [optional] 
**ValidFrom** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ValidUntil** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon expires. Coupon never expires if this is omitted, zero, or negative. | [optional] 
**CampaignId** | Pointer to **int64** | The ID of the campaign to which the coupon belongs. | 
**CustomerProfileId** | Pointer to **string** | The Integration ID of the customer that is allowed to redeem this coupon. | [optional] 
**UsageLimit** | Pointer to **int64** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**UsageCounter** | Pointer to **int64** | The number of times the coupon has been successfully redeemed. | 
**BatchId** | Pointer to **string** | The ID of the batch the coupon belongs to. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Custom attributes associated with this coupon. | 

## Methods

### NewExpiringCouponsData

`func NewExpiringCouponsData(couponValue string, campaignId int64, usageLimit int64, usageCounter int64, attributes map[string]interface{}, ) *ExpiringCouponsData`

NewExpiringCouponsData instantiates a new ExpiringCouponsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringCouponsDataWithDefaults

`func NewExpiringCouponsDataWithDefaults() *ExpiringCouponsData`

NewExpiringCouponsDataWithDefaults instantiates a new ExpiringCouponsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponValue

`func (o *ExpiringCouponsData) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *ExpiringCouponsData) GetCouponValueOk() (*string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponValue

`func (o *ExpiringCouponsData) SetCouponValue(v string)`

SetCouponValue sets CouponValue field to given value.


### GetCreatedDate

`func (o *ExpiringCouponsData) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ExpiringCouponsData) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ExpiringCouponsData) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ExpiringCouponsData) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetValidFrom

`func (o *ExpiringCouponsData) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *ExpiringCouponsData) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *ExpiringCouponsData) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *ExpiringCouponsData) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *ExpiringCouponsData) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *ExpiringCouponsData) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *ExpiringCouponsData) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *ExpiringCouponsData) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetCampaignId

`func (o *ExpiringCouponsData) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ExpiringCouponsData) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ExpiringCouponsData) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetCustomerProfileId

`func (o *ExpiringCouponsData) GetCustomerProfileId() string`

GetCustomerProfileId returns the CustomerProfileId field if non-nil, zero value otherwise.

### GetCustomerProfileIdOk

`func (o *ExpiringCouponsData) GetCustomerProfileIdOk() (*string, bool)`

GetCustomerProfileIdOk returns a tuple with the CustomerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileId

`func (o *ExpiringCouponsData) SetCustomerProfileId(v string)`

SetCustomerProfileId sets CustomerProfileId field to given value.

### HasCustomerProfileId

`func (o *ExpiringCouponsData) HasCustomerProfileId() bool`

HasCustomerProfileId returns a boolean if a field has been set.

### GetUsageLimit

`func (o *ExpiringCouponsData) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *ExpiringCouponsData) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *ExpiringCouponsData) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetUsageCounter

`func (o *ExpiringCouponsData) GetUsageCounter() int64`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *ExpiringCouponsData) GetUsageCounterOk() (*int64, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCounter

`func (o *ExpiringCouponsData) SetUsageCounter(v int64)`

SetUsageCounter sets UsageCounter field to given value.


### GetBatchId

`func (o *ExpiringCouponsData) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *ExpiringCouponsData) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *ExpiringCouponsData) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *ExpiringCouponsData) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetAttributes

`func (o *ExpiringCouponsData) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExpiringCouponsData) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExpiringCouponsData) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


