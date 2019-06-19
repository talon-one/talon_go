# Referral

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique ID for this entity. | [default to null]
**Created** | [**time.Time**](time.Time.md) | The exact moment this entity was created. | [default to null]
**CampaignId** | **int32** | ID of the campaign from which the referral received the referral code. | [default to null]
**AdvocateProfileIntegrationId** | **string** | The Integration Id of the Advocate&#39;s Profile | [default to null]
**FriendProfileIntegrationId** | **string** | An optional Integration ID of the Friend&#39;s Profile | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Timestamp at which point the referral code becomes valid. | [optional] [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiry date of the referral code. Referral never expires if this is omitted, zero, or negative. | [optional] [default to null]
**Code** | **string** | The actual referral code. | [default to null]
**UsageCounter** | **int32** | The number of times this referral code has been successfully used. | [default to null]
**UsageLimit** | **int32** | The number of times a referral code can be used. This can be set to 0 for no limit, but any campaign usage limits will still apply.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


