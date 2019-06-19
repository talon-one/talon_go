# LedgerEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**ProfileId** | **string** | ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID. | [default to null]
**AccountId** | **int32** | The ID of the Talon.One account that owns this profile. | [default to null]
**LoyaltyProgramId** | **int32** | ID of the ledger | [default to null]
**EventId** | **int32** | ID of the related event | [default to null]
**Amount** | **int32** | Amount of loyalty points | [default to null]
**Reason** | **string** | reason for awarding/deducting points | [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiry date of the points | [default to null]
**ReferenceId** | **int32** | The ID of the balancing ledgerEntry | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


