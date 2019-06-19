# UpdateCoupon

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | **int32** | The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.  | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] [default to null]
**RecipientIntegrationId** | **string** | The integration ID for this coupon&#39;s beneficiary&#39;s profile | [optional] [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Arbitrary properties associated with this item | [optional] [default to null]
**Reservation** | **bool** | This value controls what reservations mean to a coupon. If set to true the coupon reservation is used to mark it as a favourite, if set to false the coupon reservation is used as a requirement of usage. This value defaults to false if not specified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


