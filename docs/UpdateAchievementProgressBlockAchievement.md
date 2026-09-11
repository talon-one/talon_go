# UpdateAchievementProgressBlockAchievement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the achievement. | 
**Name** | Pointer to **string** | The internal name of the achievement used in API requests. | 
**Title** | Pointer to **string** | The display name of the achievement in the Campaign Manager. | 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | 

## Methods

### NewUpdateAchievementProgressBlockAchievement

`func NewUpdateAchievementProgressBlockAchievement(id int64, name string, title string, target float32, ) *UpdateAchievementProgressBlockAchievement`

NewUpdateAchievementProgressBlockAchievement instantiates a new UpdateAchievementProgressBlockAchievement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAchievementProgressBlockAchievementWithDefaults

`func NewUpdateAchievementProgressBlockAchievementWithDefaults() *UpdateAchievementProgressBlockAchievement`

NewUpdateAchievementProgressBlockAchievementWithDefaults instantiates a new UpdateAchievementProgressBlockAchievement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAchievementProgressBlockAchievement) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAchievementProgressBlockAchievement) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAchievementProgressBlockAchievement) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateAchievementProgressBlockAchievement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAchievementProgressBlockAchievement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAchievementProgressBlockAchievement) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *UpdateAchievementProgressBlockAchievement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateAchievementProgressBlockAchievement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateAchievementProgressBlockAchievement) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTarget

`func (o *UpdateAchievementProgressBlockAchievement) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *UpdateAchievementProgressBlockAchievement) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *UpdateAchievementProgressBlockAchievement) SetTarget(v float32)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


