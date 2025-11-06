# IncreaseAchievementProgressEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementId** | Pointer to **int64** | The internal ID of the achievement. | 
**AchievementName** | Pointer to **string** | The name of the achievement. | 
**ProgressTrackerId** | Pointer to **int64** | The internal ID of the achievement progress tracker. | [optional] 
**Delta** | Pointer to **float32** | The value by which the customer&#39;s current progress in the achievement is increased. | 
**Value** | Pointer to **float32** | The current progress of the customer in the achievement. | 
**Target** | Pointer to **float32** | The target value to complete the achievement. | 
**IsJustCompleted** | Pointer to **bool** | Indicates if the customer has completed the achievement in the current session. | 

## Methods

### NewIncreaseAchievementProgressEffectProps

`func NewIncreaseAchievementProgressEffectProps(achievementId int64, achievementName string, delta float32, value float32, target float32, isJustCompleted bool, ) *IncreaseAchievementProgressEffectProps`

NewIncreaseAchievementProgressEffectProps instantiates a new IncreaseAchievementProgressEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncreaseAchievementProgressEffectPropsWithDefaults

`func NewIncreaseAchievementProgressEffectPropsWithDefaults() *IncreaseAchievementProgressEffectProps`

NewIncreaseAchievementProgressEffectPropsWithDefaults instantiates a new IncreaseAchievementProgressEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementId

`func (o *IncreaseAchievementProgressEffectProps) GetAchievementId() int64`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *IncreaseAchievementProgressEffectProps) GetAchievementIdOk() (*int64, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementId

`func (o *IncreaseAchievementProgressEffectProps) SetAchievementId(v int64)`

SetAchievementId sets AchievementId field to given value.


### GetAchievementName

`func (o *IncreaseAchievementProgressEffectProps) GetAchievementName() string`

GetAchievementName returns the AchievementName field if non-nil, zero value otherwise.

### GetAchievementNameOk

`func (o *IncreaseAchievementProgressEffectProps) GetAchievementNameOk() (*string, bool)`

GetAchievementNameOk returns a tuple with the AchievementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementName

`func (o *IncreaseAchievementProgressEffectProps) SetAchievementName(v string)`

SetAchievementName sets AchievementName field to given value.


### GetProgressTrackerId

`func (o *IncreaseAchievementProgressEffectProps) GetProgressTrackerId() int64`

GetProgressTrackerId returns the ProgressTrackerId field if non-nil, zero value otherwise.

### GetProgressTrackerIdOk

`func (o *IncreaseAchievementProgressEffectProps) GetProgressTrackerIdOk() (*int64, bool)`

GetProgressTrackerIdOk returns a tuple with the ProgressTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTrackerId

`func (o *IncreaseAchievementProgressEffectProps) SetProgressTrackerId(v int64)`

SetProgressTrackerId sets ProgressTrackerId field to given value.

### HasProgressTrackerId

`func (o *IncreaseAchievementProgressEffectProps) HasProgressTrackerId() bool`

HasProgressTrackerId returns a boolean if a field has been set.

### GetDelta

`func (o *IncreaseAchievementProgressEffectProps) GetDelta() float32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *IncreaseAchievementProgressEffectProps) GetDeltaOk() (*float32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *IncreaseAchievementProgressEffectProps) SetDelta(v float32)`

SetDelta sets Delta field to given value.


### GetValue

`func (o *IncreaseAchievementProgressEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IncreaseAchievementProgressEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IncreaseAchievementProgressEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetTarget

`func (o *IncreaseAchievementProgressEffectProps) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IncreaseAchievementProgressEffectProps) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IncreaseAchievementProgressEffectProps) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetIsJustCompleted

`func (o *IncreaseAchievementProgressEffectProps) GetIsJustCompleted() bool`

GetIsJustCompleted returns the IsJustCompleted field if non-nil, zero value otherwise.

### GetIsJustCompletedOk

`func (o *IncreaseAchievementProgressEffectProps) GetIsJustCompletedOk() (*bool, bool)`

GetIsJustCompletedOk returns a tuple with the IsJustCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJustCompleted

`func (o *IncreaseAchievementProgressEffectProps) SetIsJustCompleted(v bool)`

SetIsJustCompleted sets IsJustCompleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


