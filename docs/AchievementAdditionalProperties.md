# AchievementAdditionalProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int64** | The ID of the campaign the achievement belongs to. | 
**UserId** | Pointer to **int64** | ID of the user that created this achievement. | 
**CreatedBy** | Pointer to **string** | Name of the user that created the achievement.  **Note**: This is not available if the user has been deleted.  | [optional] 
**HasProgress** | Pointer to **bool** | Indicates if a customer has made progress in the achievement. | [optional] 
**Status** | Pointer to **string** | The status of the achievement. | [optional] 

## Methods

### NewAchievementAdditionalProperties

`func NewAchievementAdditionalProperties(campaignId int64, userId int64, ) *AchievementAdditionalProperties`

NewAchievementAdditionalProperties instantiates a new AchievementAdditionalProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchievementAdditionalPropertiesWithDefaults

`func NewAchievementAdditionalPropertiesWithDefaults() *AchievementAdditionalProperties`

NewAchievementAdditionalPropertiesWithDefaults instantiates a new AchievementAdditionalProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *AchievementAdditionalProperties) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *AchievementAdditionalProperties) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *AchievementAdditionalProperties) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetUserId

`func (o *AchievementAdditionalProperties) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AchievementAdditionalProperties) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AchievementAdditionalProperties) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetCreatedBy

`func (o *AchievementAdditionalProperties) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AchievementAdditionalProperties) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AchievementAdditionalProperties) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AchievementAdditionalProperties) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHasProgress

`func (o *AchievementAdditionalProperties) GetHasProgress() bool`

GetHasProgress returns the HasProgress field if non-nil, zero value otherwise.

### GetHasProgressOk

`func (o *AchievementAdditionalProperties) GetHasProgressOk() (*bool, bool)`

GetHasProgressOk returns a tuple with the HasProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProgress

`func (o *AchievementAdditionalProperties) SetHasProgress(v bool)`

SetHasProgress sets HasProgress field to given value.

### HasHasProgress

`func (o *AchievementAdditionalProperties) HasHasProgress() bool`

HasHasProgress returns a boolean if a field has been set.

### GetStatus

`func (o *AchievementAdditionalProperties) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchievementAdditionalProperties) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AchievementAdditionalProperties) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AchievementAdditionalProperties) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


