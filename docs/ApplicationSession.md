# ApplicationSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. The exact moment this entity was created. | [default to null]
**ApplicationId** | **int32** | The ID of the application that owns this entity. | [default to null]
**ProfileId** | **int32** | The globally unique Talon.One ID of the customer that created this entity. | [optional] [default to null]
**IntegrationId** | **string** | The ID used for this entity in the application system. | [default to null]
**Coupon** | **string** | Any coupon code entered. | [default to null]
**Referral** | **string** | Any referal code entered. | [default to null]
**State** | **string** | Indicating if the customer session is in progress (\&quot;open\&quot;), \&quot;closed\&quot;, or \&quot;cancelled\&quot;. | [default to null]
**CartItems** | [**[]CartItem**](CartItem.md) | Serialized JSON representation. | [default to null]
**Discounts** | **map[string]float32** | A map of labelled discount values, in the same currency as the session. | [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Arbitrary properties associated with this item | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


