# LedgerInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentTier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**PointsToNextTier** | Pointer to **float32** | Points required to move up a tier. | [optional] 

## Methods

### GetCurrentTier

`func (o *LedgerInfoAllOf) GetCurrentTier() Tier`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *LedgerInfoAllOf) GetCurrentTierOk() (Tier, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentTier

`func (o *LedgerInfoAllOf) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### SetCurrentTier

`func (o *LedgerInfoAllOf) SetCurrentTier(v Tier)`

SetCurrentTier gets a reference to the given Tier and assigns it to the CurrentTier field.

### GetPointsToNextTier

`func (o *LedgerInfoAllOf) GetPointsToNextTier() float32`

GetPointsToNextTier returns the PointsToNextTier field if non-nil, zero value otherwise.

### GetPointsToNextTierOk

`func (o *LedgerInfoAllOf) GetPointsToNextTierOk() (float32, bool)`

GetPointsToNextTierOk returns a tuple with the PointsToNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPointsToNextTier

`func (o *LedgerInfoAllOf) HasPointsToNextTier() bool`

HasPointsToNextTier returns a boolean if a field has been set.

### SetPointsToNextTier

`func (o *LedgerInfoAllOf) SetPointsToNextTier(v float32)`

SetPointsToNextTier gets a reference to the given float32 and assigns it to the PointsToNextTier field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


