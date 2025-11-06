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

### NewAchievementProgress

`func NewAchievementProgress(status string, progress float32, ) *AchievementProgress`

NewAchievementProgress instantiates a new AchievementProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchievementProgressWithDefaults

`func NewAchievementProgressWithDefaults() *AchievementProgress`

NewAchievementProgressWithDefaults instantiates a new AchievementProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AchievementProgress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchievementProgress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AchievementProgress) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *AchievementProgress) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *AchievementProgress) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *AchievementProgress) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetStartDate

`func (o *AchievementProgress) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AchievementProgress) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AchievementProgress) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AchievementProgress) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetCompletionDate

`func (o *AchievementProgress) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *AchievementProgress) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *AchievementProgress) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *AchievementProgress) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AchievementProgress) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AchievementProgress) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AchievementProgress) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AchievementProgress) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


