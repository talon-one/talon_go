# NewCustomerSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID. | [optional] [default to null]
**Coupon** | **string** | Any coupon code entered. | [optional] [default to null]
**Referral** | **string** | Any referral code entered. | [optional] [default to null]
**State** | **string** | Indicates the current state of the session. All sessions must start in the \&quot;open\&quot; state, after which valid transitions are...  1. open -&gt; closed 2. open -&gt; cancelled 3. closed -&gt; cancelled  | [optional] [default to null]
**CartItems** | [**[]CartItem**](CartItem.md) | Serialized JSON representation. | [optional] [default to null]
**Identifiers** | **[]string** | Identifiers for the customer, this can be used for limits on values such as device ID. | [optional] [default to null]
**Total** | **float32** | The total sum of the cart in one session. | [optional] [default to null]
**Attributes** | [***interface{}**](interface{}.md) | A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


