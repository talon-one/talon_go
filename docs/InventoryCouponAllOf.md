# InventoryCouponAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileRedemptionCount** | Pointer to **int32** | The number of times the coupon was redeemed by the profile. | 
**State** | Pointer to **string** | Can be:  - &#x60;active&#x60;: The coupon can be used. It is a reserved coupon that is not pending, used, or expired, and it has a non-exhausted limit counter.    **Note:** This coupon state is returned for [scheduled campaigns](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-schedule), but the coupon cannot be used until the campaign is **running**. - &#x60;used&#x60;: The coupon has been redeemed and cannot be used again. It is not pending and has reached its redemption limit or was redeemed by the profile before expiration. - &#x60;expired&#x60;: The coupon was never redeemed, and it is now expired. It is non-pending, non-active, and non-used by the profile. - &#x60;pending&#x60;: The coupon will be usable in the future. - &#x60;disabled&#x60;: The coupon is part of a non-active campaign.  | 

## Methods

### GetProfileRedemptionCount

`func (o *InventoryCouponAllOf) GetProfileRedemptionCount() int32`

GetProfileRedemptionCount returns the ProfileRedemptionCount field if non-nil, zero value otherwise.

### GetProfileRedemptionCountOk

`func (o *InventoryCouponAllOf) GetProfileRedemptionCountOk() (int32, bool)`

GetProfileRedemptionCountOk returns a tuple with the ProfileRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileRedemptionCount

`func (o *InventoryCouponAllOf) HasProfileRedemptionCount() bool`

HasProfileRedemptionCount returns a boolean if a field has been set.

### SetProfileRedemptionCount

`func (o *InventoryCouponAllOf) SetProfileRedemptionCount(v int32)`

SetProfileRedemptionCount gets a reference to the given int32 and assigns it to the ProfileRedemptionCount field.

### GetState

`func (o *InventoryCouponAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InventoryCouponAllOf) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *InventoryCouponAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *InventoryCouponAllOf) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


