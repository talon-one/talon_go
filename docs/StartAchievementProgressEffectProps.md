# StartAchievementProgressEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementId** | Pointer to **int64** | The ID of the achievement. | 
**AchievementName** | Pointer to **string** | The name of the achievement. | 
**ProgressTrackerId** | Pointer to **int64** | The ID of the customer&#39;s progress tracker for this achievement.  For [on-completion achievements](https://docs.talon.one/docs/product/campaigns/achievements/overview#recurring-on-completion-achievements), this effect generates a unique ID for each iteration. | [optional] 
**Target** | Pointer to **float32** | The target value to complete the achievement. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which the customer&#39;s progress started. | 
**EndDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which this progress period ends.  Only returned for achievements that have a fixed end date. [On-completion achievements](https://docs.talon.one/docs/product/campaigns/achievements/overview#recurring-on-completion-achievements) have no end date. | [optional] 

## Methods

### NewStartAchievementProgressEffectProps

`func NewStartAchievementProgressEffectProps(achievementId int64, achievementName string, target float32, startDate time.Time, ) *StartAchievementProgressEffectProps`

NewStartAchievementProgressEffectProps instantiates a new StartAchievementProgressEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartAchievementProgressEffectPropsWithDefaults

`func NewStartAchievementProgressEffectPropsWithDefaults() *StartAchievementProgressEffectProps`

NewStartAchievementProgressEffectPropsWithDefaults instantiates a new StartAchievementProgressEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementId

`func (o *StartAchievementProgressEffectProps) GetAchievementId() int64`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *StartAchievementProgressEffectProps) GetAchievementIdOk() (*int64, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementId

`func (o *StartAchievementProgressEffectProps) SetAchievementId(v int64)`

SetAchievementId sets AchievementId field to given value.


### GetAchievementName

`func (o *StartAchievementProgressEffectProps) GetAchievementName() string`

GetAchievementName returns the AchievementName field if non-nil, zero value otherwise.

### GetAchievementNameOk

`func (o *StartAchievementProgressEffectProps) GetAchievementNameOk() (*string, bool)`

GetAchievementNameOk returns a tuple with the AchievementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementName

`func (o *StartAchievementProgressEffectProps) SetAchievementName(v string)`

SetAchievementName sets AchievementName field to given value.


### GetProgressTrackerId

`func (o *StartAchievementProgressEffectProps) GetProgressTrackerId() int64`

GetProgressTrackerId returns the ProgressTrackerId field if non-nil, zero value otherwise.

### GetProgressTrackerIdOk

`func (o *StartAchievementProgressEffectProps) GetProgressTrackerIdOk() (*int64, bool)`

GetProgressTrackerIdOk returns a tuple with the ProgressTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTrackerId

`func (o *StartAchievementProgressEffectProps) SetProgressTrackerId(v int64)`

SetProgressTrackerId sets ProgressTrackerId field to given value.

### HasProgressTrackerId

`func (o *StartAchievementProgressEffectProps) HasProgressTrackerId() bool`

HasProgressTrackerId returns a boolean if a field has been set.

### GetTarget

`func (o *StartAchievementProgressEffectProps) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *StartAchievementProgressEffectProps) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *StartAchievementProgressEffectProps) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetStartDate

`func (o *StartAchievementProgressEffectProps) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *StartAchievementProgressEffectProps) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *StartAchievementProgressEffectProps) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *StartAchievementProgressEffectProps) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *StartAchievementProgressEffectProps) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *StartAchievementProgressEffectProps) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *StartAchievementProgressEffectProps) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


