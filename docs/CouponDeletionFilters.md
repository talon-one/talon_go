# CouponDeletionFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBefore** | Pointer to [**time.Time**](time.Time.md) | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | [optional] 
**CreatedAfter** | Pointer to [**time.Time**](time.Time.md) | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | [optional] 
**StartsAfter** | Pointer to [**time.Time**](time.Time.md) | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | [optional] 
**StartsBefore** | Pointer to [**time.Time**](time.Time.md) | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | [optional] 
**Valid** | Pointer to **string** | - &#x60;expired&#x60;: Matches coupons in which the expiration date is set and in the past. - &#x60;validNow&#x60;: Matches coupons in which the start date is null or in the past and the expiration date is null or in the future. - &#x60;validFuture&#x60;: Matches coupons in which the start date is set and in the future.  | [optional] 
**Usable** | Pointer to **bool** | - &#x60;true&#x60;: only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned. - &#x60;false&#x60;: only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60; will be returned. - This field cannot be used in conjunction with the &#x60;usable&#x60; query parameter.  | [optional] 
**Redeemed** | Pointer to **bool** | - &#x60;true&#x60;: only coupons where &#x60;usageCounter &gt; 0&#x60; will be returned. - &#x60;false&#x60;: only coupons where &#x60;usageCounter &#x3D; 0&#x60; will be returned.  **Note:** This field cannot be used in conjunction with the &#x60;usable&#x60; query parameter.  | [optional] 
**RecipientIntegrationId** | Pointer to **string** | Filter results by match with a profile id specified in the coupon&#39;s &#x60;RecipientIntegrationId&#x60; field.  | [optional] 
**ExactMatch** | Pointer to **bool** | Filter results to an exact case-insensitive matching against the coupon code | [optional] [default to false]
**Value** | Pointer to **string** | Filter results by the coupon code | [optional] [default to false]
**BatchId** | Pointer to **string** | Filter results by batches of coupons | [optional] 
**ReferralId** | Pointer to **int32** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | [optional] 
**ExpiresAfter** | Pointer to [**time.Time**](time.Time.md) | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | [optional] 
**ExpiresBefore** | Pointer to [**time.Time**](time.Time.md) | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | [optional] 

## Methods

### GetCreatedBefore

`func (o *CouponDeletionFilters) GetCreatedBefore() time.Time`

GetCreatedBefore returns the CreatedBefore field if non-nil, zero value otherwise.

### GetCreatedBeforeOk

`func (o *CouponDeletionFilters) GetCreatedBeforeOk() (time.Time, bool)`

GetCreatedBeforeOk returns a tuple with the CreatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBefore

`func (o *CouponDeletionFilters) HasCreatedBefore() bool`

HasCreatedBefore returns a boolean if a field has been set.

### SetCreatedBefore

`func (o *CouponDeletionFilters) SetCreatedBefore(v time.Time)`

SetCreatedBefore gets a reference to the given time.Time and assigns it to the CreatedBefore field.

### GetCreatedAfter

`func (o *CouponDeletionFilters) GetCreatedAfter() time.Time`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *CouponDeletionFilters) GetCreatedAfterOk() (time.Time, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAfter

`func (o *CouponDeletionFilters) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### SetCreatedAfter

`func (o *CouponDeletionFilters) SetCreatedAfter(v time.Time)`

SetCreatedAfter gets a reference to the given time.Time and assigns it to the CreatedAfter field.

### GetStartsAfter

`func (o *CouponDeletionFilters) GetStartsAfter() time.Time`

GetStartsAfter returns the StartsAfter field if non-nil, zero value otherwise.

### GetStartsAfterOk

`func (o *CouponDeletionFilters) GetStartsAfterOk() (time.Time, bool)`

GetStartsAfterOk returns a tuple with the StartsAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartsAfter

`func (o *CouponDeletionFilters) HasStartsAfter() bool`

HasStartsAfter returns a boolean if a field has been set.

### SetStartsAfter

`func (o *CouponDeletionFilters) SetStartsAfter(v time.Time)`

SetStartsAfter gets a reference to the given time.Time and assigns it to the StartsAfter field.

### GetStartsBefore

`func (o *CouponDeletionFilters) GetStartsBefore() time.Time`

GetStartsBefore returns the StartsBefore field if non-nil, zero value otherwise.

### GetStartsBeforeOk

`func (o *CouponDeletionFilters) GetStartsBeforeOk() (time.Time, bool)`

GetStartsBeforeOk returns a tuple with the StartsBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartsBefore

`func (o *CouponDeletionFilters) HasStartsBefore() bool`

HasStartsBefore returns a boolean if a field has been set.

### SetStartsBefore

`func (o *CouponDeletionFilters) SetStartsBefore(v time.Time)`

SetStartsBefore gets a reference to the given time.Time and assigns it to the StartsBefore field.

### GetValid

`func (o *CouponDeletionFilters) GetValid() string`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CouponDeletionFilters) GetValidOk() (string, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValid

`func (o *CouponDeletionFilters) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValid

`func (o *CouponDeletionFilters) SetValid(v string)`

SetValid gets a reference to the given string and assigns it to the Valid field.

### GetUsable

`func (o *CouponDeletionFilters) GetUsable() bool`

GetUsable returns the Usable field if non-nil, zero value otherwise.

### GetUsableOk

`func (o *CouponDeletionFilters) GetUsableOk() (bool, bool)`

GetUsableOk returns a tuple with the Usable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsable

`func (o *CouponDeletionFilters) HasUsable() bool`

HasUsable returns a boolean if a field has been set.

### SetUsable

`func (o *CouponDeletionFilters) SetUsable(v bool)`

SetUsable gets a reference to the given bool and assigns it to the Usable field.

### GetRedeemed

`func (o *CouponDeletionFilters) GetRedeemed() bool`

GetRedeemed returns the Redeemed field if non-nil, zero value otherwise.

### GetRedeemedOk

`func (o *CouponDeletionFilters) GetRedeemedOk() (bool, bool)`

GetRedeemedOk returns a tuple with the Redeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRedeemed

`func (o *CouponDeletionFilters) HasRedeemed() bool`

HasRedeemed returns a boolean if a field has been set.

### SetRedeemed

`func (o *CouponDeletionFilters) SetRedeemed(v bool)`

SetRedeemed gets a reference to the given bool and assigns it to the Redeemed field.

### GetRecipientIntegrationId

`func (o *CouponDeletionFilters) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *CouponDeletionFilters) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *CouponDeletionFilters) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *CouponDeletionFilters) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.

### GetExactMatch

`func (o *CouponDeletionFilters) GetExactMatch() bool`

GetExactMatch returns the ExactMatch field if non-nil, zero value otherwise.

### GetExactMatchOk

`func (o *CouponDeletionFilters) GetExactMatchOk() (bool, bool)`

GetExactMatchOk returns a tuple with the ExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExactMatch

`func (o *CouponDeletionFilters) HasExactMatch() bool`

HasExactMatch returns a boolean if a field has been set.

### SetExactMatch

`func (o *CouponDeletionFilters) SetExactMatch(v bool)`

SetExactMatch gets a reference to the given bool and assigns it to the ExactMatch field.

### GetValue

`func (o *CouponDeletionFilters) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CouponDeletionFilters) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *CouponDeletionFilters) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *CouponDeletionFilters) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.

### GetBatchId

`func (o *CouponDeletionFilters) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *CouponDeletionFilters) GetBatchIdOk() (string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchId

`func (o *CouponDeletionFilters) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchId

`func (o *CouponDeletionFilters) SetBatchId(v string)`

SetBatchId gets a reference to the given string and assigns it to the BatchId field.

### GetReferralId

`func (o *CouponDeletionFilters) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *CouponDeletionFilters) GetReferralIdOk() (int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralId

`func (o *CouponDeletionFilters) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### SetReferralId

`func (o *CouponDeletionFilters) SetReferralId(v int32)`

SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.

### GetExpiresAfter

`func (o *CouponDeletionFilters) GetExpiresAfter() time.Time`

GetExpiresAfter returns the ExpiresAfter field if non-nil, zero value otherwise.

### GetExpiresAfterOk

`func (o *CouponDeletionFilters) GetExpiresAfterOk() (time.Time, bool)`

GetExpiresAfterOk returns a tuple with the ExpiresAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiresAfter

`func (o *CouponDeletionFilters) HasExpiresAfter() bool`

HasExpiresAfter returns a boolean if a field has been set.

### SetExpiresAfter

`func (o *CouponDeletionFilters) SetExpiresAfter(v time.Time)`

SetExpiresAfter gets a reference to the given time.Time and assigns it to the ExpiresAfter field.

### GetExpiresBefore

`func (o *CouponDeletionFilters) GetExpiresBefore() time.Time`

GetExpiresBefore returns the ExpiresBefore field if non-nil, zero value otherwise.

### GetExpiresBeforeOk

`func (o *CouponDeletionFilters) GetExpiresBeforeOk() (time.Time, bool)`

GetExpiresBeforeOk returns a tuple with the ExpiresBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiresBefore

`func (o *CouponDeletionFilters) HasExpiresBefore() bool`

HasExpiresBefore returns a boolean if a field has been set.

### SetExpiresBefore

`func (o *CouponDeletionFilters) SetExpiresBefore(v time.Time)`

SetExpiresBefore gets a reference to the given time.Time and assigns it to the ExpiresBefore field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


