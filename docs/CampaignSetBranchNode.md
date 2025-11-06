# CampaignSetBranchNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Indicates the node type. | 
**Name** | Pointer to **string** | Name of the set. | 
**Operator** | Pointer to **string** | An indicator of how the set operates on its elements. | 
**Elements** | Pointer to [**[]CampaignSetNode**](CampaignSetNode.md) | Child elements of this set. | 
**GroupId** | Pointer to **int64** | The ID of the campaign set. | 
**Locked** | Pointer to **bool** | An indicator of whether the campaign set is locked for modification. | 
**Description** | Pointer to **string** | A description of the campaign set. | [optional] 
**EvaluationMode** | Pointer to **string** | The mode by which campaigns in the campaign evaluation group are evaluated. | 
**EvaluationScope** | Pointer to **string** | The evaluation scope of the campaign evaluation group. | 

## Methods

### NewCampaignSetBranchNode

`func NewCampaignSetBranchNode(type_ string, name string, operator string, elements []CampaignSetNode, groupId int64, locked bool, evaluationMode string, evaluationScope string, ) *CampaignSetBranchNode`

NewCampaignSetBranchNode instantiates a new CampaignSetBranchNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSetBranchNodeWithDefaults

`func NewCampaignSetBranchNodeWithDefaults() *CampaignSetBranchNode`

NewCampaignSetBranchNodeWithDefaults instantiates a new CampaignSetBranchNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignSetBranchNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignSetBranchNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignSetBranchNode) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CampaignSetBranchNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignSetBranchNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignSetBranchNode) SetName(v string)`

SetName sets Name field to given value.


### GetOperator

`func (o *CampaignSetBranchNode) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CampaignSetBranchNode) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CampaignSetBranchNode) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetElements

`func (o *CampaignSetBranchNode) GetElements() []CampaignSetNode`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CampaignSetBranchNode) GetElementsOk() (*[]CampaignSetNode, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CampaignSetBranchNode) SetElements(v []CampaignSetNode)`

SetElements sets Elements field to given value.


### GetGroupId

`func (o *CampaignSetBranchNode) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CampaignSetBranchNode) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CampaignSetBranchNode) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetLocked

`func (o *CampaignSetBranchNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CampaignSetBranchNode) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CampaignSetBranchNode) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetDescription

`func (o *CampaignSetBranchNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignSetBranchNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignSetBranchNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignSetBranchNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluationMode

`func (o *CampaignSetBranchNode) GetEvaluationMode() string`

GetEvaluationMode returns the EvaluationMode field if non-nil, zero value otherwise.

### GetEvaluationModeOk

`func (o *CampaignSetBranchNode) GetEvaluationModeOk() (*string, bool)`

GetEvaluationModeOk returns a tuple with the EvaluationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMode

`func (o *CampaignSetBranchNode) SetEvaluationMode(v string)`

SetEvaluationMode sets EvaluationMode field to given value.


### GetEvaluationScope

`func (o *CampaignSetBranchNode) GetEvaluationScope() string`

GetEvaluationScope returns the EvaluationScope field if non-nil, zero value otherwise.

### GetEvaluationScopeOk

`func (o *CampaignSetBranchNode) GetEvaluationScopeOk() (*string, bool)`

GetEvaluationScopeOk returns a tuple with the EvaluationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationScope

`func (o *CampaignSetBranchNode) SetEvaluationScope(v string)`

SetEvaluationScope sets EvaluationScope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


