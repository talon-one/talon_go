# AchievementBlockReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the achievement. | 
**Title** | Pointer to **string** | The display name for the achievement in the Campaign Manager. | [optional] [readonly] 
**Name** | Pointer to **string** | The internal name of the achievement used in API requests. | [optional] [readonly] 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | [optional] [readonly] 

## Methods

### NewAchievementBlockReference

`func NewAchievementBlockReference(id int64, ) *AchievementBlockReference`

NewAchievementBlockReference instantiates a new AchievementBlockReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchievementBlockReferenceWithDefaults

`func NewAchievementBlockReferenceWithDefaults() *AchievementBlockReference`

NewAchievementBlockReferenceWithDefaults instantiates a new AchievementBlockReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AchievementBlockReference) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AchievementBlockReference) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AchievementBlockReference) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *AchievementBlockReference) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AchievementBlockReference) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AchievementBlockReference) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AchievementBlockReference) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *AchievementBlockReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AchievementBlockReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AchievementBlockReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AchievementBlockReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *AchievementBlockReference) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AchievementBlockReference) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AchievementBlockReference) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AchievementBlockReference) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


