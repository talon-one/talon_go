# CampaignPrioritiesChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | The ID of the Application whose campaigns&#39; priorities changed. | 
**OldPriorities** | Pointer to [**CampaignSet**](CampaignSet.md) |  | [optional] 
**Priorities** | Pointer to [**CampaignSet**](CampaignSet.md) |  | 

## Methods

### GetApplicationId

`func (o *CampaignPrioritiesChangedNotification) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignPrioritiesChangedNotification) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CampaignPrioritiesChangedNotification) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CampaignPrioritiesChangedNotification) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetOldPriorities

`func (o *CampaignPrioritiesChangedNotification) GetOldPriorities() CampaignSet`

GetOldPriorities returns the OldPriorities field if non-nil, zero value otherwise.

### GetOldPrioritiesOk

`func (o *CampaignPrioritiesChangedNotification) GetOldPrioritiesOk() (CampaignSet, bool)`

GetOldPrioritiesOk returns a tuple with the OldPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldPriorities

`func (o *CampaignPrioritiesChangedNotification) HasOldPriorities() bool`

HasOldPriorities returns a boolean if a field has been set.

### SetOldPriorities

`func (o *CampaignPrioritiesChangedNotification) SetOldPriorities(v CampaignSet)`

SetOldPriorities gets a reference to the given CampaignSet and assigns it to the OldPriorities field.

### GetPriorities

`func (o *CampaignPrioritiesChangedNotification) GetPriorities() CampaignSet`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *CampaignPrioritiesChangedNotification) GetPrioritiesOk() (CampaignSet, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriorities

`func (o *CampaignPrioritiesChangedNotification) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPriorities

`func (o *CampaignPrioritiesChangedNotification) SetPriorities(v CampaignSet)`

SetPriorities gets a reference to the given CampaignSet and assigns it to the Priorities field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


