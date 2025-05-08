# AchievementProgressWithDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the achievement. | 
**Progress** | Pointer to **float32** | The current progress of the customer in the achievement. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the customer started the achievement. | [optional] 
**CompletionDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the customer completed the achievement. | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the achievement ends and resets for the customer. | [optional] 
**AchievementId** | Pointer to **int32** | The internal ID of the achievement. | 
**Name** | Pointer to **string** | The internal name of the achievement used in API requests.  | 
**Title** | Pointer to **string** | The display name of the achievement in the Campaign Manager. | 
**Description** | Pointer to **string** | The description of the achievement in the Campaign Manager. | 
**CampaignId** | Pointer to **int32** | The ID of the campaign the achievement belongs to. | 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | [optional] 
**AchievementRecurrencePolicy** | Pointer to **string** | The policy that determines if and how the achievement recurs. - &#x60;no_recurrence&#x60;: The achievement can be completed only once. - &#x60;on_expiration&#x60;: The achievement resets after it expires and becomes available again.  | 
**AchievementActivationPolicy** | Pointer to **string** | The policy that determines how the achievement starts, ends, or resets. - &#x60;user_action&#x60;: The achievement ends or resets relative to when the customer started the achievement. - &#x60;fixed_schedule&#x60;: The achievement starts, ends, or resets for all customers following a fixed schedule.  | 
**AchievementFixedStartDate** | Pointer to [**time.Time**](time.Time.md) | The achievement&#39;s start date when &#x60;achievementActivationPolicy&#x60; is equal to &#x60;fixed_schedule&#x60;.  **Note:** It is an RFC3339 timestamp string.  | [optional] 
**AchievementEndDate** | Pointer to [**time.Time**](time.Time.md) | The achievement&#39;s end date. If defined, customers cannot participate in the achievement after this date.  **Note:** It is an RFC3339 timestamp string.  | [optional] 

## Methods

### GetStatus

`func (o *AchievementProgressWithDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchievementProgressWithDefinition) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *AchievementProgressWithDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *AchievementProgressWithDefinition) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetProgress

`func (o *AchievementProgressWithDefinition) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *AchievementProgressWithDefinition) GetProgressOk() (float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgress

`func (o *AchievementProgressWithDefinition) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgress

`func (o *AchievementProgressWithDefinition) SetProgress(v float32)`

SetProgress gets a reference to the given float32 and assigns it to the Progress field.

### GetStartDate

`func (o *AchievementProgressWithDefinition) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AchievementProgressWithDefinition) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *AchievementProgressWithDefinition) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *AchievementProgressWithDefinition) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetCompletionDate

`func (o *AchievementProgressWithDefinition) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *AchievementProgressWithDefinition) GetCompletionDateOk() (time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompletionDate

`func (o *AchievementProgressWithDefinition) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### SetCompletionDate

`func (o *AchievementProgressWithDefinition) SetCompletionDate(v time.Time)`

SetCompletionDate gets a reference to the given time.Time and assigns it to the CompletionDate field.

### GetEndDate

`func (o *AchievementProgressWithDefinition) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AchievementProgressWithDefinition) GetEndDateOk() (time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndDate

`func (o *AchievementProgressWithDefinition) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDate

`func (o *AchievementProgressWithDefinition) SetEndDate(v time.Time)`

SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.

### GetAchievementId

`func (o *AchievementProgressWithDefinition) GetAchievementId() int32`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *AchievementProgressWithDefinition) GetAchievementIdOk() (int32, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievementId

`func (o *AchievementProgressWithDefinition) HasAchievementId() bool`

HasAchievementId returns a boolean if a field has been set.

### SetAchievementId

`func (o *AchievementProgressWithDefinition) SetAchievementId(v int32)`

SetAchievementId gets a reference to the given int32 and assigns it to the AchievementId field.

### GetName

`func (o *AchievementProgressWithDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AchievementProgressWithDefinition) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *AchievementProgressWithDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *AchievementProgressWithDefinition) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *AchievementProgressWithDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AchievementProgressWithDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *AchievementProgressWithDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *AchievementProgressWithDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *AchievementProgressWithDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AchievementProgressWithDefinition) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *AchievementProgressWithDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *AchievementProgressWithDefinition) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetCampaignId

`func (o *AchievementProgressWithDefinition) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *AchievementProgressWithDefinition) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *AchievementProgressWithDefinition) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *AchievementProgressWithDefinition) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetTarget

`func (o *AchievementProgressWithDefinition) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AchievementProgressWithDefinition) GetTargetOk() (float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *AchievementProgressWithDefinition) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *AchievementProgressWithDefinition) SetTarget(v float32)`

SetTarget gets a reference to the given float32 and assigns it to the Target field.

### GetAchievementRecurrencePolicy

`func (o *AchievementProgressWithDefinition) GetAchievementRecurrencePolicy() string`

GetAchievementRecurrencePolicy returns the AchievementRecurrencePolicy field if non-nil, zero value otherwise.

### GetAchievementRecurrencePolicyOk

`func (o *AchievementProgressWithDefinition) GetAchievementRecurrencePolicyOk() (string, bool)`

GetAchievementRecurrencePolicyOk returns a tuple with the AchievementRecurrencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievementRecurrencePolicy

`func (o *AchievementProgressWithDefinition) HasAchievementRecurrencePolicy() bool`

HasAchievementRecurrencePolicy returns a boolean if a field has been set.

### SetAchievementRecurrencePolicy

`func (o *AchievementProgressWithDefinition) SetAchievementRecurrencePolicy(v string)`

SetAchievementRecurrencePolicy gets a reference to the given string and assigns it to the AchievementRecurrencePolicy field.

### GetAchievementActivationPolicy

`func (o *AchievementProgressWithDefinition) GetAchievementActivationPolicy() string`

GetAchievementActivationPolicy returns the AchievementActivationPolicy field if non-nil, zero value otherwise.

### GetAchievementActivationPolicyOk

`func (o *AchievementProgressWithDefinition) GetAchievementActivationPolicyOk() (string, bool)`

GetAchievementActivationPolicyOk returns a tuple with the AchievementActivationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievementActivationPolicy

`func (o *AchievementProgressWithDefinition) HasAchievementActivationPolicy() bool`

HasAchievementActivationPolicy returns a boolean if a field has been set.

### SetAchievementActivationPolicy

`func (o *AchievementProgressWithDefinition) SetAchievementActivationPolicy(v string)`

SetAchievementActivationPolicy gets a reference to the given string and assigns it to the AchievementActivationPolicy field.

### GetAchievementFixedStartDate

`func (o *AchievementProgressWithDefinition) GetAchievementFixedStartDate() time.Time`

GetAchievementFixedStartDate returns the AchievementFixedStartDate field if non-nil, zero value otherwise.

### GetAchievementFixedStartDateOk

`func (o *AchievementProgressWithDefinition) GetAchievementFixedStartDateOk() (time.Time, bool)`

GetAchievementFixedStartDateOk returns a tuple with the AchievementFixedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievementFixedStartDate

`func (o *AchievementProgressWithDefinition) HasAchievementFixedStartDate() bool`

HasAchievementFixedStartDate returns a boolean if a field has been set.

### SetAchievementFixedStartDate

`func (o *AchievementProgressWithDefinition) SetAchievementFixedStartDate(v time.Time)`

SetAchievementFixedStartDate gets a reference to the given time.Time and assigns it to the AchievementFixedStartDate field.

### GetAchievementEndDate

`func (o *AchievementProgressWithDefinition) GetAchievementEndDate() time.Time`

GetAchievementEndDate returns the AchievementEndDate field if non-nil, zero value otherwise.

### GetAchievementEndDateOk

`func (o *AchievementProgressWithDefinition) GetAchievementEndDateOk() (time.Time, bool)`

GetAchievementEndDateOk returns a tuple with the AchievementEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAchievementEndDate

`func (o *AchievementProgressWithDefinition) HasAchievementEndDate() bool`

HasAchievementEndDate returns a boolean if a field has been set.

### SetAchievementEndDate

`func (o *AchievementProgressWithDefinition) SetAchievementEndDate(v time.Time)`

SetAchievementEndDate gets a reference to the given time.Time and assigns it to the AchievementEndDate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


