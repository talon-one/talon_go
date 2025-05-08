# AchievementProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the achievement. | 
**Progress** | Pointer to **float32** | The current progress of the customer in the achievement. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the customer started the achievement. | [optional] 
**CompletionDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the customer completed the achievement. | [optional] 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the achievement ends and resets for the customer. | [optional] 

## Methods

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


