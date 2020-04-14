# FeatureFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Loyalty** | Pointer to **bool** | Whether the account has access to the loyalty features or not | [optional] 
**CouponsWithoutCount** | Pointer to **bool** | Whether the account queries coupons with or without total result size | [optional] 
**BetaEffects** | Pointer to **bool** | Whether the account can test beta effects or not | [optional] 

## Methods

### GetAccountId

`func (o *FeatureFlags) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *FeatureFlags) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *FeatureFlags) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *FeatureFlags) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetLoyalty

`func (o *FeatureFlags) GetLoyalty() bool`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *FeatureFlags) GetLoyaltyOk() (bool, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyalty

`func (o *FeatureFlags) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### SetLoyalty

`func (o *FeatureFlags) SetLoyalty(v bool)`

SetLoyalty gets a reference to the given bool and assigns it to the Loyalty field.

### GetCouponsWithoutCount

`func (o *FeatureFlags) GetCouponsWithoutCount() bool`

GetCouponsWithoutCount returns the CouponsWithoutCount field if non-nil, zero value otherwise.

### GetCouponsWithoutCountOk

`func (o *FeatureFlags) GetCouponsWithoutCountOk() (bool, bool)`

GetCouponsWithoutCountOk returns a tuple with the CouponsWithoutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponsWithoutCount

`func (o *FeatureFlags) HasCouponsWithoutCount() bool`

HasCouponsWithoutCount returns a boolean if a field has been set.

### SetCouponsWithoutCount

`func (o *FeatureFlags) SetCouponsWithoutCount(v bool)`

SetCouponsWithoutCount gets a reference to the given bool and assigns it to the CouponsWithoutCount field.

### GetBetaEffects

`func (o *FeatureFlags) GetBetaEffects() bool`

GetBetaEffects returns the BetaEffects field if non-nil, zero value otherwise.

### GetBetaEffectsOk

`func (o *FeatureFlags) GetBetaEffectsOk() (bool, bool)`

GetBetaEffectsOk returns a tuple with the BetaEffects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBetaEffects

`func (o *FeatureFlags) HasBetaEffects() bool`

HasBetaEffects returns a boolean if a field has been set.

### SetBetaEffects

`func (o *FeatureFlags) SetBetaEffects(v bool)`

SetBetaEffects gets a reference to the given bool and assigns it to the BetaEffects field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


