# LoyaltySubLedger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** |  | 
**Transactions** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | Transactions contains a list of all events that have happened such as additions, subtractions and expiries | [optional] 
**ExpiringPoints** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | ExpiringPoints contains a list of all points that will expiry and when | [optional] 

## Methods

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


