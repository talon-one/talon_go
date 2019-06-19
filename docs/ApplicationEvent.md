# ApplicationEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**ApplicationId** | **int32** | The ID of the application that owns this entity. | [default to null]
**ProfileId** | **int32** | The globally unique Talon.One ID of the customer that created this entity. | [optional] [default to null]
**SessionId** | **int32** | The globally unique Talon.One ID of the session that contains this event. | [optional] [default to null]
**Type_** | **string** | A string representing the event. Must not be a reserved event name. | [default to null]
**Attributes** | [***interface{}**](interface{}.md) | Additional JSON serialized data associated with the event. | [default to null]
**Effects** | [**[]interface{}**](interface{}.md) | An array containing the effects that were applied as a result of this event. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


