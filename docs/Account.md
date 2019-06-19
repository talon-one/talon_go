# Account

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**Modified** | [**time.Time**](time.Time.md) | The exact moment this entity was last modified. | [default to null]
**CompanyName** | **string** |  | [default to null]
**DomainName** | **string** | Subdomain Name for yourcompany.talon.one | [default to null]
**State** | **string** | State of the account (trial, active, trial_expired) | [default to null]
**BillingEmail** | **string** | The billing email address associated with your company account. | [default to null]
**PlanName** | **string** | The name of your booked plan. | [optional] [default to null]
**PlanExpires** | [**time.Time**](time.Time.md) | The point in time at which your current plan expires. | [optional] [default to null]
**ApplicationLimit** | **int32** | The maximum number of Applications covered by your plan. | [optional] [default to null]
**UserLimit** | **int32** | The maximum number of Campaign Manager Users covered by your plan. | [optional] [default to null]
**CampaignLimit** | **int32** | The maximum number of Campaigns covered by your plan. | [optional] [default to null]
**ApiLimit** | **int32** | The maximum number of Integration API calls covered by your plan per billing period. | [optional] [default to null]
**ApplicationCount** | **int32** | The current number of Applications in your account. | [default to null]
**UserCount** | **int32** | The current number of Campaign Manager Users in your account. | [default to null]
**CampaignsActiveCount** | **int32** | The current number of active Campaigns in your account. | [default to null]
**CampaignsInactiveCount** | **int32** | The current number of inactive Campaigns in your account. | [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Arbitrary properties associated with this campaign | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


