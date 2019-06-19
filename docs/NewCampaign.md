# NewCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A friendly name for this campaign. | [default to null]
**Description** | **string** | A detailed description of the campaign. | [optional] [default to null]
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


