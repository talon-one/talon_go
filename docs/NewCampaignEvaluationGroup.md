# NewCampaignEvaluationGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the campaign evaluation group. | 
**ParentId** | Pointer to **int64** | The ID of the parent group that contains the campaign evaluation group. | 
**Description** | Pointer to **string** | A description of the campaign evaluation group. | [optional] 
**EvaluationMode** | Pointer to **string** | The mode by which campaigns in the campaign evaluation group are evaluated. | 
**EvaluationScope** | Pointer to **string** | The evaluation scope of the campaign evaluation group. | 
**Locked** | Pointer to **bool** | An indicator of whether the campaign evaluation group is locked for modification. | 

## Methods

### NewNewCampaignEvaluationGroup

`func NewNewCampaignEvaluationGroup(name string, parentId int64, evaluationMode string, evaluationScope string, locked bool, ) *NewCampaignEvaluationGroup`

NewNewCampaignEvaluationGroup instantiates a new NewCampaignEvaluationGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCampaignEvaluationGroupWithDefaults

`func NewNewCampaignEvaluationGroupWithDefaults() *NewCampaignEvaluationGroup`

NewNewCampaignEvaluationGroupWithDefaults instantiates a new NewCampaignEvaluationGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewCampaignEvaluationGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewCampaignEvaluationGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewCampaignEvaluationGroup) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *NewCampaignEvaluationGroup) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *NewCampaignEvaluationGroup) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *NewCampaignEvaluationGroup) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetDescription

`func (o *NewCampaignEvaluationGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewCampaignEvaluationGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewCampaignEvaluationGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewCampaignEvaluationGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluationMode

`func (o *NewCampaignEvaluationGroup) GetEvaluationMode() string`

GetEvaluationMode returns the EvaluationMode field if non-nil, zero value otherwise.

### GetEvaluationModeOk

`func (o *NewCampaignEvaluationGroup) GetEvaluationModeOk() (*string, bool)`

GetEvaluationModeOk returns a tuple with the EvaluationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMode

`func (o *NewCampaignEvaluationGroup) SetEvaluationMode(v string)`

SetEvaluationMode sets EvaluationMode field to given value.


### GetEvaluationScope

`func (o *NewCampaignEvaluationGroup) GetEvaluationScope() string`

GetEvaluationScope returns the EvaluationScope field if non-nil, zero value otherwise.

### GetEvaluationScopeOk

`func (o *NewCampaignEvaluationGroup) GetEvaluationScopeOk() (*string, bool)`

GetEvaluationScopeOk returns a tuple with the EvaluationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationScope

`func (o *NewCampaignEvaluationGroup) SetEvaluationScope(v string)`

SetEvaluationScope sets EvaluationScope field to given value.


### GetLocked

`func (o *NewCampaignEvaluationGroup) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *NewCampaignEvaluationGroup) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *NewCampaignEvaluationGroup) SetLocked(v bool)`

SetLocked sets Locked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


