# CampaignCopy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the copied campaign (Defaults to \&quot;Copy of original campaign name\&quot;). | [optional] 
**ApplicationIds** | Pointer to **[]int64** | Application IDs of the applications to which a campaign should be copied to. | 
**Description** | Pointer to **string** | A detailed description of the campaign. | [optional] 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the campaign will become inactive. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags for the campaign. | [optional] 
**EvaluationGroupId** | Pointer to **int64** | The ID of the campaign evaluation group the campaign belongs to. | [optional] 

## Methods

### NewCampaignCopy

`func NewCampaignCopy(applicationIds []int64, ) *CampaignCopy`

NewCampaignCopy instantiates a new CampaignCopy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCopyWithDefaults

`func NewCampaignCopyWithDefaults() *CampaignCopy`

NewCampaignCopyWithDefaults instantiates a new CampaignCopy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignCopy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignCopy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignCopy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignCopy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApplicationIds

`func (o *CampaignCopy) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *CampaignCopy) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *CampaignCopy) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetDescription

`func (o *CampaignCopy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignCopy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignCopy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignCopy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartTime

`func (o *CampaignCopy) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CampaignCopy) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CampaignCopy) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CampaignCopy) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *CampaignCopy) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CampaignCopy) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CampaignCopy) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CampaignCopy) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetTags

`func (o *CampaignCopy) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CampaignCopy) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CampaignCopy) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CampaignCopy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvaluationGroupId

`func (o *CampaignCopy) GetEvaluationGroupId() int64`

GetEvaluationGroupId returns the EvaluationGroupId field if non-nil, zero value otherwise.

### GetEvaluationGroupIdOk

`func (o *CampaignCopy) GetEvaluationGroupIdOk() (*int64, bool)`

GetEvaluationGroupIdOk returns a tuple with the EvaluationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupId

`func (o *CampaignCopy) SetEvaluationGroupId(v int64)`

SetEvaluationGroupId sets EvaluationGroupId field to given value.

### HasEvaluationGroupId

`func (o *CampaignCopy) HasEvaluationGroupId() bool`

HasEvaluationGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


