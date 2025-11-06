# RollbackIncreasedAchievementProgressEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementId** | Pointer to **int64** | The internal ID of the achievement. | 
**AchievementName** | Pointer to **string** | The name of the achievement. | 
**ProgressTrackerId** | Pointer to **int64** | The internal ID of the achievement progress tracker. | 
**DecreaseProgressBy** | Pointer to **float32** | The value by which the customer&#39;s current progress in the achievement is decreased. | 
**CurrentProgress** | Pointer to **float32** | The current progress of the customer in the achievement. | 
**Target** | Pointer to **float32** | The target value to complete the achievement. | 

## Methods

### NewRollbackIncreasedAchievementProgressEffectProps

`func NewRollbackIncreasedAchievementProgressEffectProps(achievementId int64, achievementName string, progressTrackerId int64, decreaseProgressBy float32, currentProgress float32, target float32, ) *RollbackIncreasedAchievementProgressEffectProps`

NewRollbackIncreasedAchievementProgressEffectProps instantiates a new RollbackIncreasedAchievementProgressEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackIncreasedAchievementProgressEffectPropsWithDefaults

`func NewRollbackIncreasedAchievementProgressEffectPropsWithDefaults() *RollbackIncreasedAchievementProgressEffectProps`

NewRollbackIncreasedAchievementProgressEffectPropsWithDefaults instantiates a new RollbackIncreasedAchievementProgressEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementId

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetAchievementId() int64`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetAchievementIdOk() (*int64, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementId

`func (o *RollbackIncreasedAchievementProgressEffectProps) SetAchievementId(v int64)`

SetAchievementId sets AchievementId field to given value.


### GetAchievementName

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetAchievementName() string`

GetAchievementName returns the AchievementName field if non-nil, zero value otherwise.

### GetAchievementNameOk

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetAchievementNameOk() (*string, bool)`

GetAchievementNameOk returns a tuple with the AchievementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementName

`func (o *RollbackIncreasedAchievementProgressEffectProps) SetAchievementName(v string)`

SetAchievementName sets AchievementName field to given value.


### GetProgressTrackerId

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetProgressTrackerId() int64`

GetProgressTrackerId returns the ProgressTrackerId field if non-nil, zero value otherwise.

### GetProgressTrackerIdOk

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetProgressTrackerIdOk() (*int64, bool)`

GetProgressTrackerIdOk returns a tuple with the ProgressTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTrackerId

`func (o *RollbackIncreasedAchievementProgressEffectProps) SetProgressTrackerId(v int64)`

SetProgressTrackerId sets ProgressTrackerId field to given value.


### GetDecreaseProgressBy

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetDecreaseProgressBy() float32`

GetDecreaseProgressBy returns the DecreaseProgressBy field if non-nil, zero value otherwise.

### GetDecreaseProgressByOk

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetDecreaseProgressByOk() (*float32, bool)`

GetDecreaseProgressByOk returns a tuple with the DecreaseProgressBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecreaseProgressBy

`func (o *RollbackIncreasedAchievementProgressEffectProps) SetDecreaseProgressBy(v float32)`

SetDecreaseProgressBy sets DecreaseProgressBy field to given value.


### GetCurrentProgress

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetCurrentProgress() float32`

GetCurrentProgress returns the CurrentProgress field if non-nil, zero value otherwise.

### GetCurrentProgressOk

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetCurrentProgressOk() (*float32, bool)`

GetCurrentProgressOk returns a tuple with the CurrentProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgress

`func (o *RollbackIncreasedAchievementProgressEffectProps) SetCurrentProgress(v float32)`

SetCurrentProgress sets CurrentProgress field to given value.


### GetTarget

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RollbackIncreasedAchievementProgressEffectProps) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RollbackIncreasedAchievementProgressEffectProps) SetTarget(v float32)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


