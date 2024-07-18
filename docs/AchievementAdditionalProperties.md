# AchievementAdditionalProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int32** | ID of the campaign, to which the achievement belongs to | 
**UserId** | Pointer to **int32** | ID of the user that created this achievement. | 
**CreatedBy** | Pointer to **string** | Name of the user that created the achievement.  **Note**: This is not available if the user has been deleted.  | 
**HasProgress** | Pointer to **bool** | Indicates if a customer has made progress in the achievement. | [optional] 

## Methods

### GetCampaignId

`func (o *AchievementAdditionalProperties) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *AchievementAdditionalProperties) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *AchievementAdditionalProperties) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *AchievementAdditionalProperties) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetUserId

`func (o *AchievementAdditionalProperties) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AchievementAdditionalProperties) GetUserIdOk() (int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *AchievementAdditionalProperties) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *AchievementAdditionalProperties) SetUserId(v int32)`

SetUserId gets a reference to the given int32 and assigns it to the UserId field.

### GetCreatedBy

`func (o *AchievementAdditionalProperties) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AchievementAdditionalProperties) GetCreatedByOk() (string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *AchievementAdditionalProperties) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *AchievementAdditionalProperties) SetCreatedBy(v string)`

SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.

### GetHasProgress

`func (o *AchievementAdditionalProperties) GetHasProgress() bool`

GetHasProgress returns the HasProgress field if non-nil, zero value otherwise.

### GetHasProgressOk

`func (o *AchievementAdditionalProperties) GetHasProgressOk() (bool, bool)`

GetHasProgressOk returns a tuple with the HasProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasProgress

`func (o *AchievementAdditionalProperties) HasHasProgress() bool`

HasHasProgress returns a boolean if a field has been set.

### SetHasProgress

`func (o *AchievementAdditionalProperties) SetHasProgress(v bool)`

SetHasProgress gets a reference to the given bool and assigns it to the HasProgress field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


