# LoyaltySubLedger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePoints** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of all currently active points. | [optional] 
**CurrentTier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**ExpiredPoints** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of expired points. | [optional] 
**ExpiringPoints** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of all points that will expire. | [optional] 
**PendingPoints** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of all points pending activation. | [optional] 
**Total** | Pointer to **float32** | **DEPRECATED** Use &#x60;totalActivePoints&#x60; property instead. Total amount of currently active and available points in the customer&#39;s balance.  | 
**TotalActivePoints** | Pointer to **float32** | Total amount of currently active and available points in the customer&#39;s balance. | 
**TotalExpiredPoints** | Pointer to **float32** | Total amount of points, that expired without ever being spent. | 
**TotalPendingPoints** | Pointer to **float32** | Total amount of pending points, which are not active yet but will become active in the future. | 
**TotalSpentPoints** | Pointer to **float32** | Total amount of points already spent by this customer. | 
**Transactions** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of all events that have happened such as additions, subtractions and expiries. | [optional] 

## Methods

### GetActivePoints

`func (o *LoyaltySubLedger) GetActivePoints() []LoyaltyLedgerEntry`

GetActivePoints returns the ActivePoints field if non-nil, zero value otherwise.

### GetActivePointsOk

`func (o *LoyaltySubLedger) GetActivePointsOk() ([]LoyaltyLedgerEntry, bool)`

GetActivePointsOk returns a tuple with the ActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivePoints

`func (o *LoyaltySubLedger) HasActivePoints() bool`

HasActivePoints returns a boolean if a field has been set.

### SetActivePoints

`func (o *LoyaltySubLedger) SetActivePoints(v []LoyaltyLedgerEntry)`

SetActivePoints gets a reference to the given []LoyaltyLedgerEntry and assigns it to the ActivePoints field.

### GetCurrentTier

`func (o *LoyaltySubLedger) GetCurrentTier() Tier`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *LoyaltySubLedger) GetCurrentTierOk() (Tier, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentTier

`func (o *LoyaltySubLedger) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### SetCurrentTier

`func (o *LoyaltySubLedger) SetCurrentTier(v Tier)`

SetCurrentTier gets a reference to the given Tier and assigns it to the CurrentTier field.

### GetExpiredPoints

`func (o *LoyaltySubLedger) GetExpiredPoints() []LoyaltyLedgerEntry`

GetExpiredPoints returns the ExpiredPoints field if non-nil, zero value otherwise.

### GetExpiredPointsOk

`func (o *LoyaltySubLedger) GetExpiredPointsOk() ([]LoyaltyLedgerEntry, bool)`

GetExpiredPointsOk returns a tuple with the ExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiredPoints

`func (o *LoyaltySubLedger) HasExpiredPoints() bool`

HasExpiredPoints returns a boolean if a field has been set.

### SetExpiredPoints

`func (o *LoyaltySubLedger) SetExpiredPoints(v []LoyaltyLedgerEntry)`

SetExpiredPoints gets a reference to the given []LoyaltyLedgerEntry and assigns it to the ExpiredPoints field.

### GetExpiringPoints

`func (o *LoyaltySubLedger) GetExpiringPoints() []LoyaltyLedgerEntry`

GetExpiringPoints returns the ExpiringPoints field if non-nil, zero value otherwise.

### GetExpiringPointsOk

`func (o *LoyaltySubLedger) GetExpiringPointsOk() ([]LoyaltyLedgerEntry, bool)`

GetExpiringPointsOk returns a tuple with the ExpiringPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiringPoints

`func (o *LoyaltySubLedger) HasExpiringPoints() bool`

HasExpiringPoints returns a boolean if a field has been set.

### SetExpiringPoints

`func (o *LoyaltySubLedger) SetExpiringPoints(v []LoyaltyLedgerEntry)`

SetExpiringPoints gets a reference to the given []LoyaltyLedgerEntry and assigns it to the ExpiringPoints field.

### GetPendingPoints

`func (o *LoyaltySubLedger) GetPendingPoints() []LoyaltyLedgerEntry`

GetPendingPoints returns the PendingPoints field if non-nil, zero value otherwise.

### GetPendingPointsOk

`func (o *LoyaltySubLedger) GetPendingPointsOk() ([]LoyaltyLedgerEntry, bool)`

GetPendingPointsOk returns a tuple with the PendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingPoints

`func (o *LoyaltySubLedger) HasPendingPoints() bool`

HasPendingPoints returns a boolean if a field has been set.

### SetPendingPoints

`func (o *LoyaltySubLedger) SetPendingPoints(v []LoyaltyLedgerEntry)`

SetPendingPoints gets a reference to the given []LoyaltyLedgerEntry and assigns it to the PendingPoints field.

### GetTotal

`func (o *LoyaltySubLedger) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LoyaltySubLedger) GetTotalOk() (float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotal

`func (o *LoyaltySubLedger) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotal

`func (o *LoyaltySubLedger) SetTotal(v float32)`

SetTotal gets a reference to the given float32 and assigns it to the Total field.

### GetTotalActivePoints

`func (o *LoyaltySubLedger) GetTotalActivePoints() float32`

GetTotalActivePoints returns the TotalActivePoints field if non-nil, zero value otherwise.

### GetTotalActivePointsOk

`func (o *LoyaltySubLedger) GetTotalActivePointsOk() (float32, bool)`

GetTotalActivePointsOk returns a tuple with the TotalActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalActivePoints

`func (o *LoyaltySubLedger) HasTotalActivePoints() bool`

HasTotalActivePoints returns a boolean if a field has been set.

### SetTotalActivePoints

`func (o *LoyaltySubLedger) SetTotalActivePoints(v float32)`

SetTotalActivePoints gets a reference to the given float32 and assigns it to the TotalActivePoints field.

### GetTotalExpiredPoints

`func (o *LoyaltySubLedger) GetTotalExpiredPoints() float32`

GetTotalExpiredPoints returns the TotalExpiredPoints field if non-nil, zero value otherwise.

### GetTotalExpiredPointsOk

`func (o *LoyaltySubLedger) GetTotalExpiredPointsOk() (float32, bool)`

GetTotalExpiredPointsOk returns a tuple with the TotalExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalExpiredPoints

`func (o *LoyaltySubLedger) HasTotalExpiredPoints() bool`

HasTotalExpiredPoints returns a boolean if a field has been set.

### SetTotalExpiredPoints

`func (o *LoyaltySubLedger) SetTotalExpiredPoints(v float32)`

SetTotalExpiredPoints gets a reference to the given float32 and assigns it to the TotalExpiredPoints field.

### GetTotalPendingPoints

`func (o *LoyaltySubLedger) GetTotalPendingPoints() float32`

GetTotalPendingPoints returns the TotalPendingPoints field if non-nil, zero value otherwise.

### GetTotalPendingPointsOk

`func (o *LoyaltySubLedger) GetTotalPendingPointsOk() (float32, bool)`

GetTotalPendingPointsOk returns a tuple with the TotalPendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalPendingPoints

`func (o *LoyaltySubLedger) HasTotalPendingPoints() bool`

HasTotalPendingPoints returns a boolean if a field has been set.

### SetTotalPendingPoints

`func (o *LoyaltySubLedger) SetTotalPendingPoints(v float32)`

SetTotalPendingPoints gets a reference to the given float32 and assigns it to the TotalPendingPoints field.

### GetTotalSpentPoints

`func (o *LoyaltySubLedger) GetTotalSpentPoints() float32`

GetTotalSpentPoints returns the TotalSpentPoints field if non-nil, zero value otherwise.

### GetTotalSpentPointsOk

`func (o *LoyaltySubLedger) GetTotalSpentPointsOk() (float32, bool)`

GetTotalSpentPointsOk returns a tuple with the TotalSpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalSpentPoints

`func (o *LoyaltySubLedger) HasTotalSpentPoints() bool`

HasTotalSpentPoints returns a boolean if a field has been set.

### SetTotalSpentPoints

`func (o *LoyaltySubLedger) SetTotalSpentPoints(v float32)`

SetTotalSpentPoints gets a reference to the given float32 and assigns it to the TotalSpentPoints field.

### GetTransactions

`func (o *LoyaltySubLedger) GetTransactions() []LoyaltyLedgerEntry`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *LoyaltySubLedger) GetTransactionsOk() ([]LoyaltyLedgerEntry, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransactions

`func (o *LoyaltySubLedger) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### SetTransactions

`func (o *LoyaltySubLedger) SetTransactions(v []LoyaltyLedgerEntry)`

SetTransactions gets a reference to the given []LoyaltyLedgerEntry and assigns it to the Transactions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


