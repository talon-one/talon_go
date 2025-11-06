# LoyaltySubLedger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | **DEPRECATED** Use &#x60;totalActivePoints&#x60; property instead. Total amount of currently active and available points in the customer&#39;s balance.  | 
**TotalActivePoints** | Pointer to **float32** | Total amount of currently active and available points in the customer&#39;s balance. | 
**TotalPendingPoints** | Pointer to **float32** | Total amount of pending points, which are not active yet but will become active in the future. | 
**TotalSpentPoints** | Pointer to **float32** | Total amount of points already spent by this customer. | 
**TotalExpiredPoints** | Pointer to **float32** | Total amount of points, that expired without ever being spent. | 
**TotalNegativePoints** | Pointer to **float32** | Total amount of negative points. This implies that &#x60;totalActivePoints&#x60; is &#x60;0&#x60;. | 
**Transactions** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of all events that have happened such as additions, subtractions and expiries. | [optional] 
**ExpiringPoints** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of all points that will expire. | [optional] 
**ActivePoints** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of all currently active points. | [optional] 
**PendingPoints** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of all points pending activation. | [optional] 
**ExpiredPoints** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of expired points. | [optional] 
**CurrentTier** | Pointer to [**Tier**](Tier.md) |  | [optional] 

## Methods

### NewLoyaltySubLedger

`func NewLoyaltySubLedger(total float32, totalActivePoints float32, totalPendingPoints float32, totalSpentPoints float32, totalExpiredPoints float32, totalNegativePoints float32, ) *LoyaltySubLedger`

NewLoyaltySubLedger instantiates a new LoyaltySubLedger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltySubLedgerWithDefaults

`func NewLoyaltySubLedgerWithDefaults() *LoyaltySubLedger`

NewLoyaltySubLedgerWithDefaults instantiates a new LoyaltySubLedger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *LoyaltySubLedger) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LoyaltySubLedger) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LoyaltySubLedger) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetTotalActivePoints

`func (o *LoyaltySubLedger) GetTotalActivePoints() float32`

GetTotalActivePoints returns the TotalActivePoints field if non-nil, zero value otherwise.

### GetTotalActivePointsOk

`func (o *LoyaltySubLedger) GetTotalActivePointsOk() (*float32, bool)`

GetTotalActivePointsOk returns a tuple with the TotalActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActivePoints

`func (o *LoyaltySubLedger) SetTotalActivePoints(v float32)`

SetTotalActivePoints sets TotalActivePoints field to given value.


### GetTotalPendingPoints

`func (o *LoyaltySubLedger) GetTotalPendingPoints() float32`

GetTotalPendingPoints returns the TotalPendingPoints field if non-nil, zero value otherwise.

### GetTotalPendingPointsOk

`func (o *LoyaltySubLedger) GetTotalPendingPointsOk() (*float32, bool)`

GetTotalPendingPointsOk returns a tuple with the TotalPendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPendingPoints

`func (o *LoyaltySubLedger) SetTotalPendingPoints(v float32)`

SetTotalPendingPoints sets TotalPendingPoints field to given value.


### GetTotalSpentPoints

`func (o *LoyaltySubLedger) GetTotalSpentPoints() float32`

GetTotalSpentPoints returns the TotalSpentPoints field if non-nil, zero value otherwise.

### GetTotalSpentPointsOk

`func (o *LoyaltySubLedger) GetTotalSpentPointsOk() (*float32, bool)`

GetTotalSpentPointsOk returns a tuple with the TotalSpentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpentPoints

`func (o *LoyaltySubLedger) SetTotalSpentPoints(v float32)`

SetTotalSpentPoints sets TotalSpentPoints field to given value.


### GetTotalExpiredPoints

`func (o *LoyaltySubLedger) GetTotalExpiredPoints() float32`

GetTotalExpiredPoints returns the TotalExpiredPoints field if non-nil, zero value otherwise.

### GetTotalExpiredPointsOk

`func (o *LoyaltySubLedger) GetTotalExpiredPointsOk() (*float32, bool)`

GetTotalExpiredPointsOk returns a tuple with the TotalExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExpiredPoints

`func (o *LoyaltySubLedger) SetTotalExpiredPoints(v float32)`

SetTotalExpiredPoints sets TotalExpiredPoints field to given value.


### GetTotalNegativePoints

`func (o *LoyaltySubLedger) GetTotalNegativePoints() float32`

GetTotalNegativePoints returns the TotalNegativePoints field if non-nil, zero value otherwise.

### GetTotalNegativePointsOk

`func (o *LoyaltySubLedger) GetTotalNegativePointsOk() (*float32, bool)`

GetTotalNegativePointsOk returns a tuple with the TotalNegativePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNegativePoints

`func (o *LoyaltySubLedger) SetTotalNegativePoints(v float32)`

SetTotalNegativePoints sets TotalNegativePoints field to given value.


### GetTransactions

`func (o *LoyaltySubLedger) GetTransactions() []LoyaltyLedgerEntry`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *LoyaltySubLedger) GetTransactionsOk() (*[]LoyaltyLedgerEntry, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *LoyaltySubLedger) SetTransactions(v []LoyaltyLedgerEntry)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *LoyaltySubLedger) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetExpiringPoints

`func (o *LoyaltySubLedger) GetExpiringPoints() []LoyaltyLedgerEntry`

GetExpiringPoints returns the ExpiringPoints field if non-nil, zero value otherwise.

### GetExpiringPointsOk

`func (o *LoyaltySubLedger) GetExpiringPointsOk() (*[]LoyaltyLedgerEntry, bool)`

GetExpiringPointsOk returns a tuple with the ExpiringPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringPoints

`func (o *LoyaltySubLedger) SetExpiringPoints(v []LoyaltyLedgerEntry)`

SetExpiringPoints sets ExpiringPoints field to given value.

### HasExpiringPoints

`func (o *LoyaltySubLedger) HasExpiringPoints() bool`

HasExpiringPoints returns a boolean if a field has been set.

### GetActivePoints

`func (o *LoyaltySubLedger) GetActivePoints() []LoyaltyLedgerEntry`

GetActivePoints returns the ActivePoints field if non-nil, zero value otherwise.

### GetActivePointsOk

`func (o *LoyaltySubLedger) GetActivePointsOk() (*[]LoyaltyLedgerEntry, bool)`

GetActivePointsOk returns a tuple with the ActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePoints

`func (o *LoyaltySubLedger) SetActivePoints(v []LoyaltyLedgerEntry)`

SetActivePoints sets ActivePoints field to given value.

### HasActivePoints

`func (o *LoyaltySubLedger) HasActivePoints() bool`

HasActivePoints returns a boolean if a field has been set.

### GetPendingPoints

`func (o *LoyaltySubLedger) GetPendingPoints() []LoyaltyLedgerEntry`

GetPendingPoints returns the PendingPoints field if non-nil, zero value otherwise.

### GetPendingPointsOk

`func (o *LoyaltySubLedger) GetPendingPointsOk() (*[]LoyaltyLedgerEntry, bool)`

GetPendingPointsOk returns a tuple with the PendingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPoints

`func (o *LoyaltySubLedger) SetPendingPoints(v []LoyaltyLedgerEntry)`

SetPendingPoints sets PendingPoints field to given value.

### HasPendingPoints

`func (o *LoyaltySubLedger) HasPendingPoints() bool`

HasPendingPoints returns a boolean if a field has been set.

### GetExpiredPoints

`func (o *LoyaltySubLedger) GetExpiredPoints() []LoyaltyLedgerEntry`

GetExpiredPoints returns the ExpiredPoints field if non-nil, zero value otherwise.

### GetExpiredPointsOk

`func (o *LoyaltySubLedger) GetExpiredPointsOk() (*[]LoyaltyLedgerEntry, bool)`

GetExpiredPointsOk returns a tuple with the ExpiredPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredPoints

`func (o *LoyaltySubLedger) SetExpiredPoints(v []LoyaltyLedgerEntry)`

SetExpiredPoints sets ExpiredPoints field to given value.

### HasExpiredPoints

`func (o *LoyaltySubLedger) HasExpiredPoints() bool`

HasExpiredPoints returns a boolean if a field has been set.

### GetCurrentTier

`func (o *LoyaltySubLedger) GetCurrentTier() Tier`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *LoyaltySubLedger) GetCurrentTierOk() (*Tier, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *LoyaltySubLedger) SetCurrentTier(v Tier)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *LoyaltySubLedger) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


