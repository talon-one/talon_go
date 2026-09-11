# UpdateAchievementProgressBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | Pointer to **string** |  | 
**Value** | Pointer to **string** | The value to update the progress by. Supports template placeholders (e.g. \&quot;{{$Session.Total / 2}}\&quot;) for dynamic quantities. | 
**Achievement** | Pointer to [**UpdateAchievementProgressBlockAchievement**](UpdateAchievementProgressBlock_achievement.md) |  | 

## Methods

### NewUpdateAchievementProgressBlock

`func NewUpdateAchievementProgressBlock(type_ string, operator string, value string, achievement UpdateAchievementProgressBlockAchievement, ) *UpdateAchievementProgressBlock`

NewUpdateAchievementProgressBlock instantiates a new UpdateAchievementProgressBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAchievementProgressBlockWithDefaults

`func NewUpdateAchievementProgressBlockWithDefaults() *UpdateAchievementProgressBlock`

NewUpdateAchievementProgressBlockWithDefaults instantiates a new UpdateAchievementProgressBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAchievementProgressBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAchievementProgressBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAchievementProgressBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateAchievementProgressBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateAchievementProgressBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAchievementProgressBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAchievementProgressBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *UpdateAchievementProgressBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateAchievementProgressBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateAchievementProgressBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateAchievementProgressBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *UpdateAchievementProgressBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UpdateAchievementProgressBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UpdateAchievementProgressBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *UpdateAchievementProgressBlock) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateAchievementProgressBlock) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateAchievementProgressBlock) SetValue(v string)`

SetValue sets Value field to given value.


### GetAchievement

`func (o *UpdateAchievementProgressBlock) GetAchievement() UpdateAchievementProgressBlockAchievement`

GetAchievement returns the Achievement field if non-nil, zero value otherwise.

### GetAchievementOk

`func (o *UpdateAchievementProgressBlock) GetAchievementOk() (*UpdateAchievementProgressBlockAchievement, bool)`

GetAchievementOk returns a tuple with the Achievement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievement

`func (o *UpdateAchievementProgressBlock) SetAchievement(v UpdateAchievementProgressBlockAchievement)`

SetAchievement sets Achievement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


