# CheckAchievementBlockAchievement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the achievement. | 
**Title** | Pointer to **string** | The display name for the achievement in the Campaign Manager. | 
**Name** | Pointer to **string** | The internal name of the achievement used in API requests. | 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | 

## Methods

### NewCheckAchievementBlockAchievement

`func NewCheckAchievementBlockAchievement(id int64, title string, name string, target float32, ) *CheckAchievementBlockAchievement`

NewCheckAchievementBlockAchievement instantiates a new CheckAchievementBlockAchievement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAchievementBlockAchievementWithDefaults

`func NewCheckAchievementBlockAchievementWithDefaults() *CheckAchievementBlockAchievement`

NewCheckAchievementBlockAchievementWithDefaults instantiates a new CheckAchievementBlockAchievement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckAchievementBlockAchievement) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckAchievementBlockAchievement) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckAchievementBlockAchievement) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *CheckAchievementBlockAchievement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CheckAchievementBlockAchievement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CheckAchievementBlockAchievement) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetName

`func (o *CheckAchievementBlockAchievement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckAchievementBlockAchievement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckAchievementBlockAchievement) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *CheckAchievementBlockAchievement) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CheckAchievementBlockAchievement) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CheckAchievementBlockAchievement) SetTarget(v float32)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


