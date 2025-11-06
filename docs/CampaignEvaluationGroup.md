# CampaignEvaluationGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**Name** | Pointer to **string** | The name of the campaign evaluation group. | 
**ParentId** | Pointer to **int64** | The ID of the parent group that contains the campaign evaluation group. | 
**Description** | Pointer to **string** | A description of the campaign evaluation group. | [optional] 
**EvaluationMode** | Pointer to **string** | The mode by which campaigns in the campaign evaluation group are evaluated. | 
**EvaluationScope** | Pointer to **string** | The evaluation scope of the campaign evaluation group. | 
**Locked** | Pointer to **bool** | An indicator of whether the campaign evaluation group is locked for modification. | 
**Id** | Pointer to **int64** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 

## Methods

### NewCampaignEvaluationGroup

`func NewCampaignEvaluationGroup(applicationId int64, name string, parentId int64, evaluationMode string, evaluationScope string, locked bool, id int64, ) *CampaignEvaluationGroup`

NewCampaignEvaluationGroup instantiates a new CampaignEvaluationGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEvaluationGroupWithDefaults

`func NewCampaignEvaluationGroupWithDefaults() *CampaignEvaluationGroup`

NewCampaignEvaluationGroupWithDefaults instantiates a new CampaignEvaluationGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CampaignEvaluationGroup) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignEvaluationGroup) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CampaignEvaluationGroup) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetName

`func (o *CampaignEvaluationGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignEvaluationGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignEvaluationGroup) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *CampaignEvaluationGroup) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CampaignEvaluationGroup) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CampaignEvaluationGroup) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetDescription

`func (o *CampaignEvaluationGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignEvaluationGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignEvaluationGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignEvaluationGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluationMode

`func (o *CampaignEvaluationGroup) GetEvaluationMode() string`

GetEvaluationMode returns the EvaluationMode field if non-nil, zero value otherwise.

### GetEvaluationModeOk

`func (o *CampaignEvaluationGroup) GetEvaluationModeOk() (*string, bool)`

GetEvaluationModeOk returns a tuple with the EvaluationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMode

`func (o *CampaignEvaluationGroup) SetEvaluationMode(v string)`

SetEvaluationMode sets EvaluationMode field to given value.


### GetEvaluationScope

`func (o *CampaignEvaluationGroup) GetEvaluationScope() string`

GetEvaluationScope returns the EvaluationScope field if non-nil, zero value otherwise.

### GetEvaluationScopeOk

`func (o *CampaignEvaluationGroup) GetEvaluationScopeOk() (*string, bool)`

GetEvaluationScopeOk returns a tuple with the EvaluationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationScope

`func (o *CampaignEvaluationGroup) SetEvaluationScope(v string)`

SetEvaluationScope sets EvaluationScope field to given value.


### GetLocked

`func (o *CampaignEvaluationGroup) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CampaignEvaluationGroup) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CampaignEvaluationGroup) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetId

`func (o *CampaignEvaluationGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignEvaluationGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignEvaluationGroup) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


