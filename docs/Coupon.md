# Coupon

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**CampaignId** | **int32** | The ID of the campaign that owns this entity. | [default to null]
**Value** | **string** | The actual coupon code. | [default to null]
**UsageLimit** | **int32** | The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.  | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] [default to null]
**UsageCounter** | **int32** | The number of times this coupon has been successfully used. | [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Arbitrary properties associated with this item | [optional] [default to null]
**ReferralId** | **int32** | The integration ID of the referring customer (if any) for whom this coupon was created as an effect. | [optional] [default to null]
**RecipientIntegrationId** | **string** | The integration ID of a referred customer profile. | [optional] [default to null]
**ImportId** | **int32** | The ID of the Import which created this coupon. | [optional] [default to null]
**Reservation** | **bool** | This value controls what reservations mean to a coupon. If set to true the coupon reservation is used to mark it as a favourite, if set to false the coupon reservation is used as a requirement of usage. This value defaults to false if not specified. | [optional] [default to null]
**BatchId** | **string** | The id of the batch the coupon belongs to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


