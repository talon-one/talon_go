# LoyaltyBalancesWithTiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to [**LoyaltyBalanceWithTier**](LoyaltyBalanceWithTier.md) |  | [optional] 
**SubledgerBalances** | Pointer to [**map[string]LoyaltyBalanceWithTier**](LoyaltyBalanceWithTier.md) | Map of the loyalty balances of the subledgers of a ledger. | [optional] 

## Methods

### GetBalance

`func (o *LoyaltyBalancesWithTiers) GetBalance() LoyaltyBalanceWithTier`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *LoyaltyBalancesWithTiers) GetBalanceOk() (LoyaltyBalanceWithTier, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBalance

`func (o *LoyaltyBalancesWithTiers) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalance

`func (o *LoyaltyBalancesWithTiers) SetBalance(v LoyaltyBalanceWithTier)`

SetBalance gets a reference to the given LoyaltyBalanceWithTier and assigns it to the Balance field.

### GetSubledgerBalances

`func (o *LoyaltyBalancesWithTiers) GetSubledgerBalances() map[string]LoyaltyBalanceWithTier`

GetSubledgerBalances returns the SubledgerBalances field if non-nil, zero value otherwise.

### GetSubledgerBalancesOk

`func (o *LoyaltyBalancesWithTiers) GetSubledgerBalancesOk() (map[string]LoyaltyBalanceWithTier, bool)`

GetSubledgerBalancesOk returns a tuple with the SubledgerBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgerBalances

`func (o *LoyaltyBalancesWithTiers) HasSubledgerBalances() bool`

HasSubledgerBalances returns a boolean if a field has been set.

### SetSubledgerBalances

`func (o *LoyaltyBalancesWithTiers) SetSubledgerBalances(v map[string]LoyaltyBalanceWithTier)`

SetSubledgerBalances gets a reference to the given map[string]LoyaltyBalanceWithTier and assigns it to the SubledgerBalances field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


