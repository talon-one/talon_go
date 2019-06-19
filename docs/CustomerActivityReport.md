# CustomerActivityReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **string** | The ID used for this entity in the application system. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**Name** | **string** | The name for this customer profile. | [default to null]
**CustomerId** | **int32** | The internal Talon.One ID of the customer. | [default to null]
**LastActivity** | [**time.Time**](time.Time.md) | The last activity of the customer. | [optional] [default to null]
**CouponRedemptions** | **int32** | Number of coupon redemptions in all customer campaigns. | [default to null]
**CouponUseAttempts** | **int32** | Number of coupon use attempts in all customer campaigns. | [default to null]
**CouponFailedAttempts** | **int32** | Number of failed coupon use attempts in all customer campaigns. | [default to null]
**AccruedDiscounts** | **float32** | Number of accrued discounts in all customer campaigns. | [default to null]
**AccruedRevenue** | **float32** | Amount of accrued revenue in all customer campaigns. | [default to null]
**TotalOrders** | **int32** | Number of orders in all customer campaigns. | [default to null]
**TotalOrdersNoCoupon** | **int32** | Number of orders without coupon used in all customer campaigns. | [default to null]
**CampaignName** | **string** | The name of the campaign this customer belongs to. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


