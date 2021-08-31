# LedgerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentBalance** | Pointer to **float32** | Sum of current active points amounts | 
**PendingBalance** | Pointer to **float32** | Sum of pending points amounts | 
**ExpiredBalance** | Pointer to **float32** | Sum of expired points amounts | 
**SpentBalance** | Pointer to **float32** | Sum of spent points amounts | 
**TentativeCurrentBalance** | Pointer to **float32** | Sum of current active points amounts, including additions and deductions on open sessions | 
**CurrentTier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**PointsToNextTier** | Pointer to **float32** | Points required to move up a tier. | [optional] 

## Methods

### GetCurrentBalance

`func (o *LedgerInfo) GetCurrentBalance() float32`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *LedgerInfo) GetCurrentBalanceOk() (float32, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentBalance

`func (o *LedgerInfo) HasCurrentBalance() bool`

HasCurrentBalance returns a boolean if a field has been set.

### SetCurrentBalance

`func (o *LedgerInfo) SetCurrentBalance(v float32)`

SetCurrentBalance gets a reference to the given float32 and assigns it to the CurrentBalance field.

### GetPendingBalance

`func (o *LedgerInfo) GetPendingBalance() float32`

GetPendingBalance returns the PendingBalance field if non-nil, zero value otherwise.

### GetPendingBalanceOk

`func (o *LedgerInfo) GetPendingBalanceOk() (float32, bool)`

GetPendingBalanceOk returns a tuple with the PendingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingBalance

`func (o *LedgerInfo) HasPendingBalance() bool`

HasPendingBalance returns a boolean if a field has been set.

### SetPendingBalance

`func (o *LedgerInfo) SetPendingBalance(v float32)`

SetPendingBalance gets a reference to the given float32 and assigns it to the PendingBalance field.

### GetExpiredBalance

`func (o *LedgerInfo) GetExpiredBalance() float32`

GetExpiredBalance returns the ExpiredBalance field if non-nil, zero value otherwise.

### GetExpiredBalanceOk

`func (o *LedgerInfo) GetExpiredBalanceOk() (float32, bool)`

GetExpiredBalanceOk returns a tuple with the ExpiredBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiredBalance

`func (o *LedgerInfo) HasExpiredBalance() bool`

HasExpiredBalance returns a boolean if a field has been set.

### SetExpiredBalance

`func (o *LedgerInfo) SetExpiredBalance(v float32)`

SetExpiredBalance gets a reference to the given float32 and assigns it to the ExpiredBalance field.

### GetSpentBalance

`func (o *LedgerInfo) GetSpentBalance() float32`

GetSpentBalance returns the SpentBalance field if non-nil, zero value otherwise.

### GetSpentBalanceOk

`func (o *LedgerInfo) GetSpentBalanceOk() (float32, bool)`

GetSpentBalanceOk returns a tuple with the SpentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpentBalance

`func (o *LedgerInfo) HasSpentBalance() bool`

HasSpentBalance returns a boolean if a field has been set.

### SetSpentBalance

`func (o *LedgerInfo) SetSpentBalance(v float32)`

SetSpentBalance gets a reference to the given float32 and assigns it to the SpentBalance field.

### GetTentativeCurrentBalance

`func (o *LedgerInfo) GetTentativeCurrentBalance() float32`

GetTentativeCurrentBalance returns the TentativeCurrentBalance field if non-nil, zero value otherwise.

### GetTentativeCurrentBalanceOk

`func (o *LedgerInfo) GetTentativeCurrentBalanceOk() (float32, bool)`

GetTentativeCurrentBalanceOk returns a tuple with the TentativeCurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTentativeCurrentBalance

`func (o *LedgerInfo) HasTentativeCurrentBalance() bool`

HasTentativeCurrentBalance returns a boolean if a field has been set.

### SetTentativeCurrentBalance

`func (o *LedgerInfo) SetTentativeCurrentBalance(v float32)`

SetTentativeCurrentBalance gets a reference to the given float32 and assigns it to the TentativeCurrentBalance field.

### GetCurrentTier

`func (o *LedgerInfo) GetCurrentTier() Tier`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *LedgerInfo) GetCurrentTierOk() (Tier, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentTier

`func (o *LedgerInfo) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### SetCurrentTier

`func (o *LedgerInfo) SetCurrentTier(v Tier)`

SetCurrentTier gets a reference to the given Tier and assigns it to the CurrentTier field.

### GetPointsToNextTier

`func (o *LedgerInfo) GetPointsToNextTier() float32`

GetPointsToNextTier returns the PointsToNextTier field if non-nil, zero value otherwise.

### GetPointsToNextTierOk

`func (o *LedgerInfo) GetPointsToNextTierOk() (float32, bool)`

GetPointsToNextTierOk returns a tuple with the PointsToNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPointsToNextTier

`func (o *LedgerInfo) HasPointsToNextTier() bool`

HasPointsToNextTier returns a boolean if a field has been set.

### SetPointsToNextTier

`func (o *LedgerInfo) SetPointsToNextTier(v float32)`

SetPointsToNextTier gets a reference to the given float32 and assigns it to the PointsToNextTier field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


