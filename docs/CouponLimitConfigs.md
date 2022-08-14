# CouponLimitConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured.  | [optional] 

## Methods

### GetLimits

`func (o *CouponLimitConfigs) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CouponLimitConfigs) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *CouponLimitConfigs) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *CouponLimitConfigs) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


