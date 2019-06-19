# LoyaltyLedger

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | The current balance in the program. | [default to null]
**Transactions** | [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | Transactions contains a list of all events that have happened such as additions, subtractions and expiries | [default to null]
**ExpiringPoints** | [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | ExpiringPoints contains a list of all points that will expiry and when | [optional] [default to null]
**LoyaltyProgramId** | **int32** | The ID of the loyalty program this ledger belongs to. | [optional] [default to null]
**LoyaltyProgramName** | **string** | The name of the loyalty program this ledger belongs to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


