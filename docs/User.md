# User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**Modified** | [**time.Time**](time.Time.md) | The exact moment this entity was last modified. | [default to null]
**Email** | **string** | The email address associated with your account. | [default to null]
**AccountId** | **int32** | The ID of the account that owns this entity. | [default to null]
**InviteToken** | **string** | Invite token, empty if the user as already accepted their invite. | [default to null]
**State** | **string** | Current user state. | [default to null]
**Name** | **string** | Full name | [default to null]
**Policy** | **string** | A blob of ACL JSON | [default to null]
**ReleaseUpdate** | **bool** | Update the user via email | [default to null]
**LatestFeature** | **string** | Latest feature the user has been notified. | [optional] [default to null]
**Roles** | **[]int32** | Contains a list of all roles a user is a memeber of | [optional] [default to null]
**ApplicationNotificationSubscriptions** | [***interface{}**](interface{}.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


