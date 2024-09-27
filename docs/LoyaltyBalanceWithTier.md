# LoyaltyBalanceWithTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoints** | Pointer to **float32** | Total amount of points awarded to this customer and available to spend. | [optional] 
**PendingPoints** | Pointer to **float32** | Total amount of points awarded to this customer but not available until their start date. | [optional] 
**SpentPoints** | Pointer to **float32** | Total amount of points already spent by this customer. | [optional] 
**ExpiredPoints** | Pointer to **float32** | Total amount of points awarded but never redeemed. They cannot be used anymore. | [optional] 
**CurrentTier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**ProjectedTier** | Pointer to [**ProjectedTier**](ProjectedTier.md) |  | [optional] 
**PointsToNextTier** | Pointer to **float32** | The number of points required to move up a tier. | [optional] 
**NextTierName** | Pointer to **string** | The name of the tier consecutive to the current tier. | [optional] 

## Methods

### GetActivePoints

`func (o *LoyaltyBalanceWithTier) GetActivePoints() float32`

GetActivePoints returns the ActivePoints field if non-nil, zero value otherwise.

### GetActivePointsOk

`func (o *LoyaltyBalanceWithTier) GetActivePointsOk() (float32, bool)`

GetActivePointsOk returns a tuple with the ActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivePoints

`func (o *LoyaltyBalanceWithTier) HasActivePoints() bool`

HasActivePoints returns a boolean if a field has been set.

### SetActivePoints

`func (o *LoyaltyBalanceWithTier) SetActivePoints(v float32)`

SetActivePoints gets a reference to the given float32 and assigns it to the ActivePoints field.

### GetPendingPoints

`func (o *LoyaltyBalanceWithTier) GetPendingPoints() float32`

GetPendingPoints returns the PendingPoints field if non-nil, zero value otherwise.

### GetPendingPointsOk

`func (o *LoyaltyBalanceWithTier) GetPendingPointsOk() (float32, bool)`

GetPendingPointsOk returns a tuple with the PendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingPoints

`func (o *LoyaltyBalanceWithTier) HasPendingPoints() bool`

HasPendingPoints returns a boolean if a field has been set.

### SetPendingPoints

`func (o *LoyaltyBalanceWithTier) SetPendingPoints(v float32)`

SetPendingPoints gets a reference to the given float32 and assigns it to the PendingPoints field.

### GetSpentPoints

`func (o *LoyaltyBalanceWithTier) GetSpentPoints() float32`

GetSpentPoints returns the SpentPoints field if non-nil, zero value otherwise.

### GetSpentPointsOk

`func (o *LoyaltyBalanceWithTier) GetSpentPointsOk() (float32, bool)`

GetSpentPointsOk returns a tuple with the SpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpentPoints

`func (o *LoyaltyBalanceWithTier) HasSpentPoints() bool`

HasSpentPoints returns a boolean if a field has been set.

### SetSpentPoints

`func (o *LoyaltyBalanceWithTier) SetSpentPoints(v float32)`

SetSpentPoints gets a reference to the given float32 and assigns it to the SpentPoints field.

### GetExpiredPoints

`func (o *LoyaltyBalanceWithTier) GetExpiredPoints() float32`

GetExpiredPoints returns the ExpiredPoints field if non-nil, zero value otherwise.

### GetExpiredPointsOk

`func (o *LoyaltyBalanceWithTier) GetExpiredPointsOk() (float32, bool)`

GetExpiredPointsOk returns a tuple with the ExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiredPoints

`func (o *LoyaltyBalanceWithTier) HasExpiredPoints() bool`

HasExpiredPoints returns a boolean if a field has been set.

### SetExpiredPoints

`func (o *LoyaltyBalanceWithTier) SetExpiredPoints(v float32)`

SetExpiredPoints gets a reference to the given float32 and assigns it to the ExpiredPoints field.

### GetCurrentTier

`func (o *LoyaltyBalanceWithTier) GetCurrentTier() Tier`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *LoyaltyBalanceWithTier) GetCurrentTierOk() (Tier, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentTier

`func (o *LoyaltyBalanceWithTier) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### SetCurrentTier

`func (o *LoyaltyBalanceWithTier) SetCurrentTier(v Tier)`

SetCurrentTier gets a reference to the given Tier and assigns it to the CurrentTier field.

### GetProjectedTier

`func (o *LoyaltyBalanceWithTier) GetProjectedTier() ProjectedTier`

GetProjectedTier returns the ProjectedTier field if non-nil, zero value otherwise.

### GetProjectedTierOk

`func (o *LoyaltyBalanceWithTier) GetProjectedTierOk() (ProjectedTier, bool)`

GetProjectedTierOk returns a tuple with the ProjectedTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProjectedTier

`func (o *LoyaltyBalanceWithTier) HasProjectedTier() bool`

HasProjectedTier returns a boolean if a field has been set.

### SetProjectedTier

`func (o *LoyaltyBalanceWithTier) SetProjectedTier(v ProjectedTier)`

SetProjectedTier gets a reference to the given ProjectedTier and assigns it to the ProjectedTier field.

### GetPointsToNextTier

`func (o *LoyaltyBalanceWithTier) GetPointsToNextTier() float32`

GetPointsToNextTier returns the PointsToNextTier field if non-nil, zero value otherwise.

### GetPointsToNextTierOk

`func (o *LoyaltyBalanceWithTier) GetPointsToNextTierOk() (float32, bool)`

GetPointsToNextTierOk returns a tuple with the PointsToNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPointsToNextTier

`func (o *LoyaltyBalanceWithTier) HasPointsToNextTier() bool`

HasPointsToNextTier returns a boolean if a field has been set.

### SetPointsToNextTier

`func (o *LoyaltyBalanceWithTier) SetPointsToNextTier(v float32)`

SetPointsToNextTier gets a reference to the given float32 and assigns it to the PointsToNextTier field.

### GetNextTierName

`func (o *LoyaltyBalanceWithTier) GetNextTierName() string`

GetNextTierName returns the NextTierName field if non-nil, zero value otherwise.

### GetNextTierNameOk

`func (o *LoyaltyBalanceWithTier) GetNextTierNameOk() (string, bool)`

GetNextTierNameOk returns a tuple with the NextTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextTierName

`func (o *LoyaltyBalanceWithTier) HasNextTierName() bool`

HasNextTierName returns a boolean if a field has been set.

### SetNextTierName

`func (o *LoyaltyBalanceWithTier) SetNextTierName(v string)`

SetNextTierName gets a reference to the given string and assigns it to the NextTierName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


