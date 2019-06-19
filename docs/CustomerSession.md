# CustomerSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **string** | The ID used for this entity in the application system. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**ApplicationId** | **int32** | The ID of the application that owns this entity. | [default to null]
**ProfileId** | **string** | ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID. | [default to null]
**Coupon** | **string** | Any coupon code entered. | [default to null]
**Referral** | **string** | Any referral code entered. | [default to null]
**State** | **string** | Indicates the current state of the session. All sessions must start in the \&quot;open\&quot; state, after which valid transitions are...  1. open -&gt; closed 2. open -&gt; cancelled 3. closed -&gt; cancelled  | [default to null]
**CartItems** | [**[]CartItem**](CartItem.md) | Serialized JSON representation. | [default to null]
**Total** | **float32** | The total sum of the cart in one session. | [default to null]
**Attributes** | [***interface{}**](interface{}.md) | A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.  | [default to null]
**FirstSession** | **bool** | Indicates whether this is the first session for the customer&#39;s profile. Will always be true for anonymous sessions. | [default to null]
**Discounts** | **map[string]float32** | A map of labelled discount values, values will be in the same currency as the application associated with the session. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


