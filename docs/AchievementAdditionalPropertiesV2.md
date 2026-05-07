# AchievementAdditionalPropertiesV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** | The ID of the user that created this achievement. | 
**CreatedBy** | Pointer to **string** | Name of the user that created the achievement.  **Note**: This is not available if the user has been deleted.  | [optional] 
**HasProgress** | Pointer to **bool** | Indicates if a customer has made progress in the achievement. | [optional] 
**Status** | Pointer to **string** | The status of the achievement. | [optional] 

## Methods

### NewAchievementAdditionalPropertiesV2

`func NewAchievementAdditionalPropertiesV2(userId int64, ) *AchievementAdditionalPropertiesV2`

NewAchievementAdditionalPropertiesV2 instantiates a new AchievementAdditionalPropertiesV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchievementAdditionalPropertiesV2WithDefaults

`func NewAchievementAdditionalPropertiesV2WithDefaults() *AchievementAdditionalPropertiesV2`

NewAchievementAdditionalPropertiesV2WithDefaults instantiates a new AchievementAdditionalPropertiesV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AchievementAdditionalPropertiesV2) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AchievementAdditionalPropertiesV2) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AchievementAdditionalPropertiesV2) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetCreatedBy

`func (o *AchievementAdditionalPropertiesV2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AchievementAdditionalPropertiesV2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AchievementAdditionalPropertiesV2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AchievementAdditionalPropertiesV2) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHasProgress

`func (o *AchievementAdditionalPropertiesV2) GetHasProgress() bool`

GetHasProgress returns the HasProgress field if non-nil, zero value otherwise.

### GetHasProgressOk

`func (o *AchievementAdditionalPropertiesV2) GetHasProgressOk() (*bool, bool)`

GetHasProgressOk returns a tuple with the HasProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProgress

`func (o *AchievementAdditionalPropertiesV2) SetHasProgress(v bool)`

SetHasProgress sets HasProgress field to given value.

### HasHasProgress

`func (o *AchievementAdditionalPropertiesV2) HasHasProgress() bool`

HasHasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *AchievementAdditionalPropertiesV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchievementAdditionalPropertiesV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AchievementAdditionalPropertiesV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AchievementAdditionalPropertiesV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


