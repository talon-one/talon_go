# UpdateCouponBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | Pointer to **int32** | The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.  | [optional] 
**DiscountLimit** | Pointer to **float32** | The amount of discounts that can be given with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item | [optional] 
**BatchID** | Pointer to **string** | The id of the batch the coupon belongs to. | 

## Methods

### GetUsageLimit

`func (o *UpdateCouponBatch) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *UpdateCouponBatch) GetUsageLimitOk() (int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLimit

`func (o *UpdateCouponBatch) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimit

`func (o *UpdateCouponBatch) SetUsageLimit(v int32)`

SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.

### GetDiscountLimit

`func (o *UpdateCouponBatch) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *UpdateCouponBatch) GetDiscountLimitOk() (float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountLimit

`func (o *UpdateCouponBatch) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### SetDiscountLimit

`func (o *UpdateCouponBatch) SetDiscountLimit(v float32)`

SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.

### GetStartDate

`func (o *UpdateCouponBatch) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateCouponBatch) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *UpdateCouponBatch) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *UpdateCouponBatch) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *UpdateCouponBatch) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *UpdateCouponBatch) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *UpdateCouponBatch) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *UpdateCouponBatch) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetAttributes

`func (o *UpdateCouponBatch) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateCouponBatch) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *UpdateCouponBatch) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *UpdateCouponBatch) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetBatchID

`func (o *UpdateCouponBatch) GetBatchID() string`

GetBatchID returns the BatchID field if non-nil, zero value otherwise.

### GetBatchIDOk

`func (o *UpdateCouponBatch) GetBatchIDOk() (string, bool)`

GetBatchIDOk returns a tuple with the BatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchID

`func (o *UpdateCouponBatch) HasBatchID() bool`

HasBatchID returns a boolean if a field has been set.

### SetBatchID

`func (o *UpdateCouponBatch) SetBatchID(v string)`

SetBatchID gets a reference to the given string and assigns it to the BatchID field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


