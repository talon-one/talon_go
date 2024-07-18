# ApplicationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | Event type. It can be one of the following: [&#39;campaign_evaluation_tree_changed&#39;]  | 
**ApplicationId** | Pointer to **int32** | The ID of the Application whose campaign evaluation tree changed. | 
**OldEvaluationTree** | Pointer to [**CampaignSet**](.md) |  | [optional] 
**EvaluationTree** | Pointer to [**CampaignSet**](.md) |  | 

## Methods

### GetEvent

`func (o *ApplicationNotification) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ApplicationNotification) GetEventOk() (string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvent

`func (o *ApplicationNotification) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEvent

`func (o *ApplicationNotification) SetEvent(v string)`

SetEvent gets a reference to the given string and assigns it to the Event field.

### GetApplicationId

`func (o *ApplicationNotification) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationNotification) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *ApplicationNotification) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *ApplicationNotification) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetOldEvaluationTree

`func (o *ApplicationNotification) GetOldEvaluationTree() CampaignSet`

GetOldEvaluationTree returns the OldEvaluationTree field if non-nil, zero value otherwise.

### GetOldEvaluationTreeOk

`func (o *ApplicationNotification) GetOldEvaluationTreeOk() (CampaignSet, bool)`

GetOldEvaluationTreeOk returns a tuple with the OldEvaluationTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOldEvaluationTree

`func (o *ApplicationNotification) HasOldEvaluationTree() bool`

HasOldEvaluationTree returns a boolean if a field has been set.

### SetOldEvaluationTree

`func (o *ApplicationNotification) SetOldEvaluationTree(v CampaignSet)`

SetOldEvaluationTree gets a reference to the given CampaignSet and assigns it to the OldEvaluationTree field.

### GetEvaluationTree

`func (o *ApplicationNotification) GetEvaluationTree() CampaignSet`

GetEvaluationTree returns the EvaluationTree field if non-nil, zero value otherwise.

### GetEvaluationTreeOk

`func (o *ApplicationNotification) GetEvaluationTreeOk() (CampaignSet, bool)`

GetEvaluationTreeOk returns a tuple with the EvaluationTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationTree

`func (o *ApplicationNotification) HasEvaluationTree() bool`

HasEvaluationTree returns a boolean if a field has been set.

### SetEvaluationTree

`func (o *ApplicationNotification) SetEvaluationTree(v CampaignSet)`

SetEvaluationTree gets a reference to the given CampaignSet and assigns it to the EvaluationTree field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


