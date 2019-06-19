# Campaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**ApplicationId** | **int32** | The ID of the application that owns this entity. | [default to null]
**UserId** | **int32** | The ID of the account that owns this entity. | [default to null]
**Name** | **string** | A friendly name for this campaign. | [default to null]
**Description** | **string** | A detailed description of the campaign. | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | Datetime when the campaign will become active. | [optional] [default to null]
**EndTime** | [**time.Time**](time.Time.md) | Datetime when the campaign will become in-active. | [optional] [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Arbitrary properties associated with this campaign | [optional] [default to null]
**State** | **string** | A disabled or archived campaign is not evaluated for rules or coupons.  | [default to null]
**ActiveRulesetId** | **int32** | ID of Ruleset this campaign applies on customer session evaluation. | [optional] [default to null]
**Tags** | **[]string** | A list of tags for the campaign. | [default to null]
**Features** | **[]string** | A list of features for the campaign. | [default to null]
**CouponSettings** | [***CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] [default to null]
**ReferralSettings** | [***CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] [default to null]
**Limits** | [**[]LimitConfig**](LimitConfig.md) | The set of limits that will operate for this campaign | [default to null]
**CouponRedemptionCount** | **int32** | Number of coupons redeemed in the campaign. | [optional] [default to null]
**ReferralRedemptionCount** | **int32** | Number of referral codes redeemed in the campaign. | [optional] [default to null]
**DiscountCount** | **int32** | Total amount of discounts redeemed in the campaign. | [optional] [default to null]
**LastActivity** | [**time.Time**](time.Time.md) | Timestamp of the most recent event received by this campaign. | [optional] [default to null]
**Updated** | [**time.Time**](time.Time.md) | Timestamp of the most recent update to the campaign or any of its elements. | [optional] [default to null]
**CreatedBy** | **string** | Name of the user who created this campaign if available. | [optional] [default to null]
**UpdatedBy** | **string** | Name of the user who last updated this campaign if available. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


