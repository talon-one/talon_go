# IncreaseAchievementProgressPerItemEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementId** | Pointer to **int64** | The internal ID of the achievement. | 
**AchievementName** | Pointer to **string** | The name of the achievement. | 
**ProgressTrackerId** | Pointer to **int64** | The internal ID of the customer progress tracker. For [on-completion achievements](https://docs.talon.one/docs/product/achievements/overview#recurring-on-completion-achievements), this effect generates a unique ID for each iteration. | [optional] 
**Delta** | Pointer to **float32** | The value by which the customer&#39;s current progress in the achievement has increased. | 
**Value** | Pointer to **float32** | The current progress of the customer in the achievement. | 
**Target** | Pointer to **float32** | The target value to complete the achievement. | 
**IsJustCompleted** | Pointer to **bool** | Indicates if the customer has completed the achievement in the current session. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the &#x60;cartItem&#x60; object on which the customer&#39;s progress in the achievement was updated. | 
**CartItemSubPosition** | Pointer to **float32** | For cart items with &#x60;quantity&#x60; &gt; 1, indicates which specific unit of the item the effect applies to. | [optional] 
**BundleIndex** | Pointer to **int64** | _(With bundles only)_ The position of the specific bundle in the list of bundles created from the same bundle definition. For example, if a customer creates two identical bundles, the first has index &#x60;0&#x60; and the second has index &#x60;1&#x60;. | [optional] 
**BundleName** | Pointer to **string** | _(With bundles only)_ The name of the bundle definition. | [optional] 

## Methods

### NewIncreaseAchievementProgressPerItemEffectProps

`func NewIncreaseAchievementProgressPerItemEffectProps(achievementId int64, achievementName string, delta float32, value float32, target float32, isJustCompleted bool, cartItemPosition float32, ) *IncreaseAchievementProgressPerItemEffectProps`

NewIncreaseAchievementProgressPerItemEffectProps instantiates a new IncreaseAchievementProgressPerItemEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncreaseAchievementProgressPerItemEffectPropsWithDefaults

`func NewIncreaseAchievementProgressPerItemEffectPropsWithDefaults() *IncreaseAchievementProgressPerItemEffectProps`

NewIncreaseAchievementProgressPerItemEffectPropsWithDefaults instantiates a new IncreaseAchievementProgressPerItemEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementId

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetAchievementId() int64`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetAchievementIdOk() (*int64, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementId

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetAchievementId(v int64)`

SetAchievementId sets AchievementId field to given value.


### GetAchievementName

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetAchievementName() string`

GetAchievementName returns the AchievementName field if non-nil, zero value otherwise.

### GetAchievementNameOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetAchievementNameOk() (*string, bool)`

GetAchievementNameOk returns a tuple with the AchievementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementName

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetAchievementName(v string)`

SetAchievementName sets AchievementName field to given value.


### GetProgressTrackerId

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetProgressTrackerId() int64`

GetProgressTrackerId returns the ProgressTrackerId field if non-nil, zero value otherwise.

### GetProgressTrackerIdOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetProgressTrackerIdOk() (*int64, bool)`

GetProgressTrackerIdOk returns a tuple with the ProgressTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTrackerId

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetProgressTrackerId(v int64)`

SetProgressTrackerId sets ProgressTrackerId field to given value.

### HasProgressTrackerId

`func (o *IncreaseAchievementProgressPerItemEffectProps) HasProgressTrackerId() bool`

HasProgressTrackerId returns a boolean if a field has been set.

### GetDelta

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetDelta() float32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetDeltaOk() (*float32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetDelta(v float32)`

SetDelta sets Delta field to given value.


### GetValue

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetTarget

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetIsJustCompleted

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetIsJustCompleted() bool`

GetIsJustCompleted returns the IsJustCompleted field if non-nil, zero value otherwise.

### GetIsJustCompletedOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetIsJustCompletedOk() (*bool, bool)`

GetIsJustCompletedOk returns a tuple with the IsJustCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJustCompleted

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetIsJustCompleted(v bool)`

SetIsJustCompleted sets IsJustCompleted field to given value.


### GetCartItemPosition

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetCartItemPositionOk() (*float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemPosition

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition sets CartItemPosition field to given value.


### GetCartItemSubPosition

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetCartItemSubPositionOk() (*float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemSubPosition

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition sets CartItemSubPosition field to given value.

### HasCartItemSubPosition

`func (o *IncreaseAchievementProgressPerItemEffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### GetBundleIndex

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetBundleIndex() int64`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetBundleIndexOk() (*int64, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleIndex

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetBundleIndex(v int64)`

SetBundleIndex sets BundleIndex field to given value.

### HasBundleIndex

`func (o *IncreaseAchievementProgressPerItemEffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### GetBundleName

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *IncreaseAchievementProgressPerItemEffectProps) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *IncreaseAchievementProgressPerItemEffectProps) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *IncreaseAchievementProgressPerItemEffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


