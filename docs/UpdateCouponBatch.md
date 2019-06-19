# UpdateCouponBatch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | **int32** | The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.  | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Arbitrary properties associated with this item | [optional] [default to null]
**BatchID** | **string** | The id of the batch the coupon belongs to. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


