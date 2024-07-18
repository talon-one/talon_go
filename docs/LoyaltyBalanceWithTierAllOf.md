# LoyaltyBalanceWithTierAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentTier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**ProjectedTier** | Pointer to [**ProjectedTier**](ProjectedTier.md) |  | [optional] 
**PointsToNextTier** | Pointer to **float32** | The number of points required to move up a tier. | [optional] 
**NextTierName** | Pointer to **string** | The name of the tier consecutive to the current tier. | [optional] 

## Methods

### GetCurrentTier

`func (o *LoyaltyBalanceWithTierAllOf) GetCurrentTier() Tier`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *LoyaltyBalanceWithTierAllOf) GetCurrentTierOk() (Tier, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentTier

`func (o *LoyaltyBalanceWithTierAllOf) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### SetCurrentTier

`func (o *LoyaltyBalanceWithTierAllOf) SetCurrentTier(v Tier)`

SetCurrentTier gets a reference to the given Tier and assigns it to the CurrentTier field.

### GetProjectedTier

`func (o *LoyaltyBalanceWithTierAllOf) GetProjectedTier() ProjectedTier`

GetProjectedTier returns the ProjectedTier field if non-nil, zero value otherwise.

### GetProjectedTierOk

`func (o *LoyaltyBalanceWithTierAllOf) GetProjectedTierOk() (ProjectedTier, bool)`

GetProjectedTierOk returns a tuple with the ProjectedTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProjectedTier

`func (o *LoyaltyBalanceWithTierAllOf) HasProjectedTier() bool`

HasProjectedTier returns a boolean if a field has been set.

### SetProjectedTier

`func (o *LoyaltyBalanceWithTierAllOf) SetProjectedTier(v ProjectedTier)`

SetProjectedTier gets a reference to the given ProjectedTier and assigns it to the ProjectedTier field.

### GetPointsToNextTier

`func (o *LoyaltyBalanceWithTierAllOf) GetPointsToNextTier() float32`

GetPointsToNextTier returns the PointsToNextTier field if non-nil, zero value otherwise.

### GetPointsToNextTierOk

`func (o *LoyaltyBalanceWithTierAllOf) GetPointsToNextTierOk() (float32, bool)`

GetPointsToNextTierOk returns a tuple with the PointsToNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPointsToNextTier

`func (o *LoyaltyBalanceWithTierAllOf) HasPointsToNextTier() bool`

HasPointsToNextTier returns a boolean if a field has been set.

### SetPointsToNextTier

`func (o *LoyaltyBalanceWithTierAllOf) SetPointsToNextTier(v float32)`

SetPointsToNextTier gets a reference to the given float32 and assigns it to the PointsToNextTier field.

### GetNextTierName

`func (o *LoyaltyBalanceWithTierAllOf) GetNextTierName() string`

GetNextTierName returns the NextTierName field if non-nil, zero value otherwise.

### GetNextTierNameOk

`func (o *LoyaltyBalanceWithTierAllOf) GetNextTierNameOk() (string, bool)`

GetNextTierNameOk returns a tuple with the NextTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextTierName

`func (o *LoyaltyBalanceWithTierAllOf) HasNextTierName() bool`

HasNextTierName returns a boolean if a field has been set.

### SetNextTierName

`func (o *LoyaltyBalanceWithTierAllOf) SetNextTierName(v string)`

SetNextTierName gets a reference to the given string and assigns it to the NextTierName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


