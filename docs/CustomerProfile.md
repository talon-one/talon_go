# CustomerProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **string** | The ID used for this entity in the application system. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Arbitrary properties associated with this item | [default to null]
**AccountId** | **int32** | The ID of the Talon.One account that owns this profile. | [default to null]
**ClosedSessions** | **int32** | The total amount of closed sessions by a customer. A closed session is a successful purchase. | [default to null]
**TotalSales** | **float32** | Sum of all purchases made by this customer | [default to null]
**LoyaltyMemberships** | [**[]LoyaltyMembership**](LoyaltyMembership.md) | A list of loyalty programs joined by the customer | [default to null]
**LastActivity** | [**time.Time**](time.Time.md) | Timestamp of the most recent event received from this customer | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


