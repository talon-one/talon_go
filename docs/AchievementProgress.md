# AchievementProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementId** | Pointer to **int32** | The internal ID of the achievement. | 
**Name** | Pointer to **string** | The internal name of the achievement used in API requests.  | 
**Title** | Pointer to **string** | The display name of the achievement in the Campaign Manager. | 
**Description** | Pointer to **string** | The description of the achievement in the Campaign Manager. | 
**CampaignId** | Pointer to **int32** | The ID of the campaign the achievement belongs to. | 
**Status** | Pointer to **string** | The status of the achievement. | 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | [optional] 
**Progress** | Pointer to **float32** | The current progress of the customer in the achievement. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the customer started the achievement. | 
**CompletionDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the customer completed the achievement. | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the achievement ends and resets for the customer. | 

## Methods

### GetAchievementId

`func (o *AchievementProgress) GetAchievementId() int32`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *AchievementProgress) GetAchievementIdOk() (int32, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievementId

`func (o *AchievementProgress) HasAchievementId() bool`

HasAchievementId returns a boolean if a field has been set.

### SetAchievementId

`func (o *AchievementProgress) SetAchievementId(v int32)`

SetAchievementId gets a reference to the given int32 and assigns it to the AchievementId field.

### GetName

`func (o *AchievementProgress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AchievementProgress) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *AchievementProgress) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *AchievementProgress) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *AchievementProgress) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AchievementProgress) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *AchievementProgress) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *AchievementProgress) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *AchievementProgress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AchievementProgress) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *AchievementProgress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *AchievementProgress) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCampaignId

`func (o *AchievementProgress) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *AchievementProgress) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *AchievementProgress) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *AchievementProgress) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetStatus

`func (o *AchievementProgress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchievementProgress) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *AchievementProgress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *AchievementProgress) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetTarget

`func (o *AchievementProgress) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AchievementProgress) GetTargetOk() (float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *AchievementProgress) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *AchievementProgress) SetTarget(v float32)`

SetTarget gets a reference to the given float32 and assigns it to the Target field.

### GetProgress

`func (o *AchievementProgress) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *AchievementProgress) GetProgressOk() (float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgress

`func (o *AchievementProgress) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgress

`func (o *AchievementProgress) SetProgress(v float32)`

SetProgress gets a reference to the given float32 and assigns it to the Progress field.

### GetStartDate

`func (o *AchievementProgress) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AchievementProgress) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *AchievementProgress) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *AchievementProgress) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetCompletionDate

`func (o *AchievementProgress) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *AchievementProgress) GetCompletionDateOk() (time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompletionDate

`func (o *AchievementProgress) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDate

`func (o *AchievementProgress) SetCompletionDate(v time.Time)`

SetCompletionDate gets a reference to the given time.Time and assigns it to the CompletionDate field.

### GetEndDate

`func (o *AchievementProgress) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AchievementProgress) GetEndDateOk() (time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndDate

`func (o *AchievementProgress) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDate

`func (o *AchievementProgress) SetEndDate(v time.Time)`

SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


