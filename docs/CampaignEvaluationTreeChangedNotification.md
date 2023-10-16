# CampaignEvaluationTreeChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | The ID of the Application whose campaign evaluation tree changed. | 
**OldEvaluationTree** | Pointer to [**CampaignSet**](CampaignSet.md) |  | [optional] 
**EvaluationTree** | Pointer to [**CampaignSet**](CampaignSet.md) |  | 

## Methods

### GetApplicationId

`func (o *CampaignEvaluationTreeChangedNotification) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignEvaluationTreeChangedNotification) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CampaignEvaluationTreeChangedNotification) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CampaignEvaluationTreeChangedNotification) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetOldEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) GetOldEvaluationTree() CampaignSet`

GetOldEvaluationTree returns the OldEvaluationTree field if non-nil, zero value otherwise.

### GetOldEvaluationTreeOk

`func (o *CampaignEvaluationTreeChangedNotification) GetOldEvaluationTreeOk() (CampaignSet, bool)`

GetOldEvaluationTreeOk returns a tuple with the OldEvaluationTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) HasOldEvaluationTree() bool`

HasOldEvaluationTree returns a boolean if a field has been set.

### SetOldEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) SetOldEvaluationTree(v CampaignSet)`

SetOldEvaluationTree gets a reference to the given CampaignSet and assigns it to the OldEvaluationTree field.

### GetEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) GetEvaluationTree() CampaignSet`

GetEvaluationTree returns the EvaluationTree field if non-nil, zero value otherwise.

### GetEvaluationTreeOk

`func (o *CampaignEvaluationTreeChangedNotification) GetEvaluationTreeOk() (CampaignSet, bool)`

GetEvaluationTreeOk returns a tuple with the EvaluationTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) HasEvaluationTree() bool`

HasEvaluationTree returns a boolean if a field has been set.

### SetEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) SetEvaluationTree(v CampaignSet)`

SetEvaluationTree gets a reference to the given CampaignSet and assigns it to the EvaluationTree field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


