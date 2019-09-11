# Event

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**ApplicationId** | **int32** | The ID of the application that owns this entity. | [default to null]
**ProfileId** | **string** | ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID. | [optional] [default to null]
**Type_** | **string** | A string representing the event. Must not be a reserved event name. | [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Arbitrary additional JSON data associated with the event. | [default to null]
**SessionId** | **string** | The ID of the session that this event occurred in. | [optional] [default to null]
**Effects** | [**[]interface{}**](interface{}.md) | An array of \&quot;effects\&quot; that must be applied in response to this event. Example effects include &#x60;addItemToCart&#x60; or &#x60;setDiscount&#x60;.  | [default to null]
**LedgerEntries** | [**[]LedgerEntry**](LedgerEntry.md) | Ledger entries for the event. | [default to null]
**Meta** | [***Meta**](Meta.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


