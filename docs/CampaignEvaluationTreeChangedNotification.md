# CampaignEvaluationTreeChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int64** | The ID of the Application whose campaign evaluation tree changed. | 
**OldEvaluationTree** | Pointer to [**CampaignSet**](CampaignSet.md) |  | [optional] 
**EvaluationTree** | Pointer to [**CampaignSet**](CampaignSet.md) |  | 

## Methods

### NewCampaignEvaluationTreeChangedNotification

`func NewCampaignEvaluationTreeChangedNotification(applicationId int64, evaluationTree CampaignSet, ) *CampaignEvaluationTreeChangedNotification`

NewCampaignEvaluationTreeChangedNotification instantiates a new CampaignEvaluationTreeChangedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEvaluationTreeChangedNotificationWithDefaults

`func NewCampaignEvaluationTreeChangedNotificationWithDefaults() *CampaignEvaluationTreeChangedNotification`

NewCampaignEvaluationTreeChangedNotificationWithDefaults instantiates a new CampaignEvaluationTreeChangedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CampaignEvaluationTreeChangedNotification) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignEvaluationTreeChangedNotification) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CampaignEvaluationTreeChangedNotification) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetOldEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) GetOldEvaluationTree() CampaignSet`

GetOldEvaluationTree returns the OldEvaluationTree field if non-nil, zero value otherwise.

### GetOldEvaluationTreeOk

`func (o *CampaignEvaluationTreeChangedNotification) GetOldEvaluationTreeOk() (*CampaignSet, bool)`

GetOldEvaluationTreeOk returns a tuple with the OldEvaluationTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) SetOldEvaluationTree(v CampaignSet)`

SetOldEvaluationTree sets OldEvaluationTree field to given value.

### HasOldEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) HasOldEvaluationTree() bool`

HasOldEvaluationTree returns a boolean if a field has been set.

### GetEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) GetEvaluationTree() CampaignSet`

GetEvaluationTree returns the EvaluationTree field if non-nil, zero value otherwise.

### GetEvaluationTreeOk

`func (o *CampaignEvaluationTreeChangedNotification) GetEvaluationTreeOk() (*CampaignSet, bool)`

GetEvaluationTreeOk returns a tuple with the EvaluationTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotification) SetEvaluationTree(v CampaignSet)`

SetEvaluationTree sets EvaluationTree field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


