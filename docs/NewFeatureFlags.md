# NewFeatureFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loyalty** | Pointer to **bool** | Whether the account has access to the loyalty features or not | [optional] 
**CouponsWithoutCount** | Pointer to **bool** | Whether the account queries coupons with or without total result size | [optional] 
**BetaEffects** | Pointer to **bool** | Whether the account can test beta effects or not | [optional] 

## Methods

### GetLoyalty

`func (o *NewFeatureFlags) GetLoyalty() bool`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *NewFeatureFlags) GetLoyaltyOk() (bool, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyalty

`func (o *NewFeatureFlags) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### SetLoyalty

`func (o *NewFeatureFlags) SetLoyalty(v bool)`

SetLoyalty gets a reference to the given bool and assigns it to the Loyalty field.

### GetCouponsWithoutCount

`func (o *NewFeatureFlags) GetCouponsWithoutCount() bool`

GetCouponsWithoutCount returns the CouponsWithoutCount field if non-nil, zero value otherwise.

### GetCouponsWithoutCountOk

`func (o *NewFeatureFlags) GetCouponsWithoutCountOk() (bool, bool)`

GetCouponsWithoutCountOk returns a tuple with the CouponsWithoutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponsWithoutCount

`func (o *NewFeatureFlags) HasCouponsWithoutCount() bool`

HasCouponsWithoutCount returns a boolean if a field has been set.

### SetCouponsWithoutCount

`func (o *NewFeatureFlags) SetCouponsWithoutCount(v bool)`

SetCouponsWithoutCount gets a reference to the given bool and assigns it to the CouponsWithoutCount field.

### GetBetaEffects

`func (o *NewFeatureFlags) GetBetaEffects() bool`

GetBetaEffects returns the BetaEffects field if non-nil, zero value otherwise.

### GetBetaEffectsOk

`func (o *NewFeatureFlags) GetBetaEffectsOk() (bool, bool)`

GetBetaEffectsOk returns a tuple with the BetaEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetaEffects

`func (o *NewFeatureFlags) HasBetaEffects() bool`

HasBetaEffects returns a boolean if a field has been set.

### SetBetaEffects

`func (o *NewFeatureFlags) SetBetaEffects(v bool)`

SetBetaEffects gets a reference to the given bool and assigns it to the BetaEffects field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


