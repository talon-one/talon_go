# Change

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**UserId** | **int32** | The ID of the account that owns this entity. | [default to null]
**Entity** | **string** | API endpoint on which the change was initiated. | [default to null]
**Old** | [***interface{}**](interface{}.md) | Resource before the change occurred. | [optional] [default to null]
**New** | [***interface{}**](interface{}.md) | Resource after the change occurred. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


