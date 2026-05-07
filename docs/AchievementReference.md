# AchievementReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchievementId** | Pointer to **int64** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievementsV2) endpoint. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application associated with the campaign that references this achievement. | 
**ApplicationName** | Pointer to **string** | The name of the Application associated with the campaign that references this achievement. | 
**CampaignId** | Pointer to **int64** | The ID of the campaign that references this achievement. | 

## Methods

### NewAchievementReference

`func NewAchievementReference(achievementId int64, applicationId int64, applicationName string, campaignId int64, ) *AchievementReference`

NewAchievementReference instantiates a new AchievementReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchievementReferenceWithDefaults

`func NewAchievementReferenceWithDefaults() *AchievementReference`

NewAchievementReferenceWithDefaults instantiates a new AchievementReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchievementId

`func (o *AchievementReference) GetAchievementId() int64`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *AchievementReference) GetAchievementIdOk() (*int64, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementId

`func (o *AchievementReference) SetAchievementId(v int64)`

SetAchievementId sets AchievementId field to given value.


### GetApplicationId

`func (o *AchievementReference) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AchievementReference) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AchievementReference) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationName

`func (o *AchievementReference) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AchievementReference) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AchievementReference) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetCampaignId

`func (o *AchievementReference) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *AchievementReference) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *AchievementReference) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


