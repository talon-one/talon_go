# CampaignSetBranchNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the campaign set. | [optional] 
**Elements** | Pointer to [**[]CampaignSetNode**](CampaignSetNode.md) | Child elements of this set. | 
**EvaluationMode** | Pointer to **string** | The mode by which campaigns in the campaign evaluation group are evaluated. | 
**EvaluationScope** | Pointer to **string** | The evaluation scope of the campaign evaluation group. | 
**GroupId** | Pointer to **int32** | The ID of the campaign set. | 
**Locked** | Pointer to **bool** | An indicator of whether the campaign set is locked for modification. | 
**Name** | Pointer to **string** | Name of the set. | 
**Operator** | Pointer to **string** | An indicator of how the set operates on its elements. | 
**Type** | Pointer to **string** | Indicates the node type. | 

## Methods

### GetDescription

`func (o *CampaignSetBranchNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignSetBranchNode) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CampaignSetBranchNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CampaignSetBranchNode) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetElements

`func (o *CampaignSetBranchNode) GetElements() []CampaignSetNode`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CampaignSetBranchNode) GetElementsOk() ([]CampaignSetNode, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasElements

`func (o *CampaignSetBranchNode) HasElements() bool`

HasElements returns a boolean if a field has been set.

### SetElements

`func (o *CampaignSetBranchNode) SetElements(v []CampaignSetNode)`

SetElements gets a reference to the given []CampaignSetNode and assigns it to the Elements field.

### GetEvaluationMode

`func (o *CampaignSetBranchNode) GetEvaluationMode() string`

GetEvaluationMode returns the EvaluationMode field if non-nil, zero value otherwise.

### GetEvaluationModeOk

`func (o *CampaignSetBranchNode) GetEvaluationModeOk() (string, bool)`

GetEvaluationModeOk returns a tuple with the EvaluationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationMode

`func (o *CampaignSetBranchNode) HasEvaluationMode() bool`

HasEvaluationMode returns a boolean if a field has been set.

### SetEvaluationMode

`func (o *CampaignSetBranchNode) SetEvaluationMode(v string)`

SetEvaluationMode gets a reference to the given string and assigns it to the EvaluationMode field.

### GetEvaluationScope

`func (o *CampaignSetBranchNode) GetEvaluationScope() string`

GetEvaluationScope returns the EvaluationScope field if non-nil, zero value otherwise.

### GetEvaluationScopeOk

`func (o *CampaignSetBranchNode) GetEvaluationScopeOk() (string, bool)`

GetEvaluationScopeOk returns a tuple with the EvaluationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationScope

`func (o *CampaignSetBranchNode) HasEvaluationScope() bool`

HasEvaluationScope returns a boolean if a field has been set.

### SetEvaluationScope

`func (o *CampaignSetBranchNode) SetEvaluationScope(v string)`

SetEvaluationScope gets a reference to the given string and assigns it to the EvaluationScope field.

### GetGroupId

`func (o *CampaignSetBranchNode) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CampaignSetBranchNode) GetGroupIdOk() (int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroupId

`func (o *CampaignSetBranchNode) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupId

`func (o *CampaignSetBranchNode) SetGroupId(v int32)`

SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.

### GetLocked

`func (o *CampaignSetBranchNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CampaignSetBranchNode) GetLockedOk() (bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocked

`func (o *CampaignSetBranchNode) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLocked

`func (o *CampaignSetBranchNode) SetLocked(v bool)`

SetLocked gets a reference to the given bool and assigns it to the Locked field.

### GetName

`func (o *CampaignSetBranchNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignSetBranchNode) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CampaignSetBranchNode) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CampaignSetBranchNode) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetOperator

`func (o *CampaignSetBranchNode) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CampaignSetBranchNode) GetOperatorOk() (string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperator

`func (o *CampaignSetBranchNode) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperator

`func (o *CampaignSetBranchNode) SetOperator(v string)`

SetOperator gets a reference to the given string and assigns it to the Operator field.

### GetType

`func (o *CampaignSetBranchNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignSetBranchNode) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *CampaignSetBranchNode) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *CampaignSetBranchNode) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


