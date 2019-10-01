# UpdateUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address associated with your account. | [default to null]
**Name** | **string** | Your name. | [optional] [default to null]
**Password** | **string** | Your old password. | [optional] [default to null]
**NewPassword** | **string** | Your new password. | [optional] [default to null]
**Policy** | **string** | a blob of acl json | [optional] [default to null]
**State** | **string** | New state (\&quot;deactivated\&quot; or \&quot;active\&quot;) for the user. Only usable by admins for the user. | [optional] [default to null]
**ReleaseUpdate** | **bool** | Update the user via email | [optional] [default to null]
**LatestFeature** | **string** | The latest feature you&#39;ve been notified. | [optional] [default to null]
**Roles** | **[]int32** | Update | [optional] [default to null]
**ApplicationNotificationSubscriptions** | [***interface{}**](interface{}.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


