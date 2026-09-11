# CheckAchievementBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | Pointer to **string** | The comparison operator applied to the achievement. | 
**Achievement** | Pointer to [**CheckAchievementBlockAchievement**](CheckAchievementBlock_achievement.md) |  | 
**OnFailure** | Pointer to **[]map[string]interface{}** | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewCheckAchievementBlock

`func NewCheckAchievementBlock(type_ string, operator string, achievement CheckAchievementBlockAchievement, ) *CheckAchievementBlock`

NewCheckAchievementBlock instantiates a new CheckAchievementBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAchievementBlockWithDefaults

`func NewCheckAchievementBlockWithDefaults() *CheckAchievementBlock`

NewCheckAchievementBlockWithDefaults instantiates a new CheckAchievementBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckAchievementBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckAchievementBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckAchievementBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckAchievementBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CheckAchievementBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckAchievementBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckAchievementBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CheckAchievementBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckAchievementBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckAchievementBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckAchievementBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *CheckAchievementBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CheckAchievementBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CheckAchievementBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetAchievement

`func (o *CheckAchievementBlock) GetAchievement() CheckAchievementBlockAchievement`

GetAchievement returns the Achievement field if non-nil, zero value otherwise.

### GetAchievementOk

`func (o *CheckAchievementBlock) GetAchievementOk() (*CheckAchievementBlockAchievement, bool)`

GetAchievementOk returns a tuple with the Achievement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievement

`func (o *CheckAchievementBlock) SetAchievement(v CheckAchievementBlockAchievement)`

SetAchievement sets Achievement field to given value.


### GetOnFailure

`func (o *CheckAchievementBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *CheckAchievementBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *CheckAchievementBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *CheckAchievementBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


