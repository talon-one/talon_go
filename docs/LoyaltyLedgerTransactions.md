# LoyaltyLedgerTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | If &#x60;true&#x60;, it means that there is more data to request in the source collection. | [optional] 
**Data** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of transaction entries from a loyalty ledger. | [optional] 

## Methods

### GetHasMore

`func (o *LoyaltyLedgerTransactions) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *LoyaltyLedgerTransactions) GetHasMoreOk() (bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasMore

`func (o *LoyaltyLedgerTransactions) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMore

`func (o *LoyaltyLedgerTransactions) SetHasMore(v bool)`

SetHasMore gets a reference to the given bool and assigns it to the HasMore field.

### GetData

`func (o *LoyaltyLedgerTransactions) GetData() []LoyaltyLedgerEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoyaltyLedgerTransactions) GetDataOk() ([]LoyaltyLedgerEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasData

`func (o *LoyaltyLedgerTransactions) HasData() bool`

HasData returns a boolean if a field has been set.

### SetData

`func (o *LoyaltyLedgerTransactions) SetData(v []LoyaltyLedgerEntry)`

SetData gets a reference to the given []LoyaltyLedgerEntry and assigns it to the Data field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


