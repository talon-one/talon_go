# CampaignTemplateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationsIds** | Pointer to **[]int32** |  | [optional] 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the campaign template or any of its elements. | [optional] 
**UpdatedBy** | Pointer to **string** | Name of the user who last updated this campaign template, if available. | [optional] 
**ValidApplicationIds** | Pointer to **[]int32** | The IDs of the Applications that are related to this entity. | 

## Methods

### GetApplicationsIds

`func (o *CampaignTemplateAllOf) GetApplicationsIds() []int32`

GetApplicationsIds returns the ApplicationsIds field if non-nil, zero value otherwise.

### GetApplicationsIdsOk

`func (o *CampaignTemplateAllOf) GetApplicationsIdsOk() ([]int32, bool)`

GetApplicationsIdsOk returns a tuple with the ApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationsIds

`func (o *CampaignTemplateAllOf) HasApplicationsIds() bool`

HasApplicationsIds returns a boolean if a field has been set.

### SetApplicationsIds

`func (o *CampaignTemplateAllOf) SetApplicationsIds(v []int32)`

SetApplicationsIds gets a reference to the given []int32 and assigns it to the ApplicationsIds field.

### GetUpdated

`func (o *CampaignTemplateAllOf) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CampaignTemplateAllOf) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *CampaignTemplateAllOf) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *CampaignTemplateAllOf) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.

### GetUpdatedBy

`func (o *CampaignTemplateAllOf) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CampaignTemplateAllOf) GetUpdatedByOk() (string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedBy

`func (o *CampaignTemplateAllOf) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedBy

`func (o *CampaignTemplateAllOf) SetUpdatedBy(v string)`

SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.

### GetValidApplicationIds

`func (o *CampaignTemplateAllOf) GetValidApplicationIds() []int32`

GetValidApplicationIds returns the ValidApplicationIds field if non-nil, zero value otherwise.

### GetValidApplicationIdsOk

`func (o *CampaignTemplateAllOf) GetValidApplicationIdsOk() ([]int32, bool)`

GetValidApplicationIdsOk returns a tuple with the ValidApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidApplicationIds

`func (o *CampaignTemplateAllOf) HasValidApplicationIds() bool`

HasValidApplicationIds returns a boolean if a field has been set.

### SetValidApplicationIds

`func (o *CampaignTemplateAllOf) SetValidApplicationIds(v []int32)`

SetValidApplicationIds gets a reference to the given []int32 and assigns it to the ValidApplicationIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


