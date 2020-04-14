# LoyaltyLedger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ledger** | Pointer to [**LoyaltySubLedger**](LoyaltySubLedger.md) |  | 
**SubLedgers** | Pointer to [**map[string]LoyaltySubLedger**](LoyaltySubLedger.md) | A map containing a list of all loyalty subledger balances | [optional] 

## Methods

### GetLedger

`func (o *LoyaltyLedger) GetLedger() LoyaltySubLedger`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *LoyaltyLedger) GetLedgerOk() (LoyaltySubLedger, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLedger

`func (o *LoyaltyLedger) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### SetLedger

`func (o *LoyaltyLedger) SetLedger(v LoyaltySubLedger)`

SetLedger gets a reference to the given LoyaltySubLedger and assigns it to the Ledger field.

### GetSubLedgers

`func (o *LoyaltyLedger) GetSubLedgers() map[string]LoyaltySubLedger`

GetSubLedgers returns the SubLedgers field if non-nil, zero value otherwise.

### GetSubLedgersOk

`func (o *LoyaltyLedger) GetSubLedgersOk() (map[string]LoyaltySubLedger, bool)`

GetSubLedgersOk returns a tuple with the SubLedgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgers

`func (o *LoyaltyLedger) HasSubLedgers() bool`

HasSubLedgers returns a boolean if a field has been set.

### SetSubLedgers

`func (o *LoyaltyLedger) SetSubLedgers(v map[string]LoyaltySubLedger)`

SetSubLedgers gets a reference to the given map[string]LoyaltySubLedger and assigns it to the SubLedgers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


