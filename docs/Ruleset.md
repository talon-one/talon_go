# Ruleset

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**CampaignId** | **int32** | The ID of the campaign that owns this entity. | [default to null]
**UserId** | **int32** | The ID of the account that owns this entity. | [default to null]
**Rules** | [**[]Rule**](Rule.md) | Set of rules to apply. | [default to null]
**Bindings** | [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | [default to null]
**RbVersion** | **string** | A string indicating which version of the rulebuilder was used to create this ruleset. | [optional] [default to null]
**Activate** | **bool** | A boolean indicating whether this newly created ruleset should also be activated for the campaign owns it | [optional] [default to null]
**ActivatedAt** | [**time.Time**](time.Time.md) | Timestamp indicating when this Ruleset was activated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


