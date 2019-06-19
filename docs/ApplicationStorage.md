# ApplicationStorage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**Modified** | [**time.Time**](time.Time.md) | The exact moment this entity was last modified. | [default to null]
**ApplicationId** | **int32** | The ID of the application that owns this entity. | [default to null]
**Name** | **string** | Identifier for the information to be saved, e.g. \&quot;Loyalty points for SKU\&quot;. | [default to null]
**Datatype** | [***interface{}**](interface{}.md) | A JSON Schema describing the information to be saved in the storage | [default to null]
**Description** | **string** | Description of the application store | [optional] [default to null]
**UsedAt** | **[]string** | array of rulesets where the application storage is used | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


