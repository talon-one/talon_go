# RollbackIncreasedAchievementProgressPerItemEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementId** | Pointer to **int64** | The internal ID of the achievement. | 
**AchievementName** | Pointer to **string** | The name of the achievement. | 
**ProgressTrackerId** | Pointer to **int64** | The internal ID of the achievement progress tracker. | 
**DecreaseProgressBy** | Pointer to **float32** | The value by which the customer&#39;s current progress in the achievement has decreased. | 
**CurrentProgress** | Pointer to **float32** | The current progress of the customer in the achievement. | 
**Target** | Pointer to **float32** | The target value to complete the achievement. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the &#x60;cartItem&#x60; object on which the customer&#39;s progress in the achievement was updated. | 
**CartItemSubPosition** | Pointer to **float32** | For cart items with &#x60;quantity&#x60; &gt; 1, indicates which specific unit of the item the effect applies to. | [optional] 
**BundleIndex** | Pointer to **int64** | _(With bundles only)_ The position of the specific bundle in the list of bundles created from the same bundle definition. For example, if a customer creates two identical bundles, the first has index &#x60;0&#x60; and the second has index &#x60;1&#x60;. | [optional] 
**BundleName** | Pointer to **string** | _(With bundles only)_ The name of the bundle definition. | [optional] 

## Methods

### NewRollbackIncreasedAchievementProgressPerItemEffectProps

`func NewRollbackIncreasedAchievementProgressPerItemEffectProps(achievementId int64, achievementName string, progressTrackerId int64, decreaseProgressBy float32, currentProgress float32, target float32, cartItemPosition float32, ) *RollbackIncreasedAchievementProgressPerItemEffectProps`

NewRollbackIncreasedAchievementProgressPerItemEffectProps instantiates a new RollbackIncreasedAchievementProgressPerItemEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackIncreasedAchievementProgressPerItemEffectPropsWithDefaults

`func NewRollbackIncreasedAchievementProgressPerItemEffectPropsWithDefaults() *RollbackIncreasedAchievementProgressPerItemEffectProps`

NewRollbackIncreasedAchievementProgressPerItemEffectPropsWithDefaults instantiates a new RollbackIncreasedAchievementProgressPerItemEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementId

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetAchievementId() int64`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetAchievementIdOk() (*int64, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementId

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetAchievementId(v int64)`

SetAchievementId sets AchievementId field to given value.


### GetAchievementName

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetAchievementName() string`

GetAchievementName returns the AchievementName field if non-nil, zero value otherwise.

### GetAchievementNameOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetAchievementNameOk() (*string, bool)`

GetAchievementNameOk returns a tuple with the AchievementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementName

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetAchievementName(v string)`

SetAchievementName sets AchievementName field to given value.


### GetProgressTrackerId

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetProgressTrackerId() int64`

GetProgressTrackerId returns the ProgressTrackerId field if non-nil, zero value otherwise.

### GetProgressTrackerIdOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetProgressTrackerIdOk() (*int64, bool)`

GetProgressTrackerIdOk returns a tuple with the ProgressTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTrackerId

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetProgressTrackerId(v int64)`

SetProgressTrackerId sets ProgressTrackerId field to given value.


### GetDecreaseProgressBy

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetDecreaseProgressBy() float32`

GetDecreaseProgressBy returns the DecreaseProgressBy field if non-nil, zero value otherwise.

### GetDecreaseProgressByOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetDecreaseProgressByOk() (*float32, bool)`

GetDecreaseProgressByOk returns a tuple with the DecreaseProgressBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecreaseProgressBy

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetDecreaseProgressBy(v float32)`

SetDecreaseProgressBy sets DecreaseProgressBy field to given value.


### GetCurrentProgress

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetCurrentProgress() float32`

GetCurrentProgress returns the CurrentProgress field if non-nil, zero value otherwise.

### GetCurrentProgressOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetCurrentProgressOk() (*float32, bool)`

GetCurrentProgressOk returns a tuple with the CurrentProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgress

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetCurrentProgress(v float32)`

SetCurrentProgress sets CurrentProgress field to given value.


### GetTarget

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetCartItemPosition

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetCartItemPositionOk() (*float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemPosition

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition sets CartItemPosition field to given value.


### GetCartItemSubPosition

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetCartItemSubPositionOk() (*float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemSubPosition

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition sets CartItemSubPosition field to given value.

### HasCartItemSubPosition

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### GetBundleIndex

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetBundleIndex() int64`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetBundleIndexOk() (*int64, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleIndex

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetBundleIndex(v int64)`

SetBundleIndex sets BundleIndex field to given value.

### HasBundleIndex

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### GetBundleName

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *RollbackIncreasedAchievementProgressPerItemEffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


