# NewReferral

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int32** | ID of the campaign from which the referral received the referral code. | [default to null]
**AdvocateProfileIntegrationId** | **string** | The Integration Id of the Advocate&#39;s Profile | [default to null]
**FriendProfileIntegrationId** | **string** | An optional Integration ID of the Friend&#39;s Profile | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Timestamp at which point the referral code becomes valid. | [optional] [default to null]
**ExpiryDate** | [**time.Time**](time.Time.md) | Expiry date of the referral code. Referral never expires if this is omitted, zero, or negative. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


