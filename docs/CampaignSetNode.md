# CampaignSetNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Name** | Pointer to **string** | Name of the set. | 
**Operator** | Pointer to **string** | An indicator of how the set operates on its elements. | 
**Elements** | Pointer to [**[]CampaignSetNode**](CampaignSetNode.md) | Child elements of this set. | 
**GroupId** | Pointer to **int32** | The ID of the campaign set. | 
**Locked** | Pointer to **bool** | An indicator of whether the campaign set is locked for modification. | 
**Description** | Pointer to **string** | A description of the campaign set. | [optional] 
**EvaluationMode** | Pointer to **string** | The mode by which campaigns in the campaign evaluation group are evaluated. | 
**EvaluationScope** | Pointer to **string** | The evaluation scope of the campaign evaluation group. | 
**CampaignId** | Pointer to **int32** | ID of the campaign | 

## Methods

### GetType

`func (o *CampaignSetNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignSetNode) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *CampaignSetNode) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *CampaignSetNode) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetName

`func (o *CampaignSetNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignSetNode) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CampaignSetNode) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CampaignSetNode) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetOperator

`func (o *CampaignSetNode) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CampaignSetNode) GetOperatorOk() (string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperator

`func (o *CampaignSetNode) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperator

`func (o *CampaignSetNode) SetOperator(v string)`

SetOperator gets a reference to the given string and assigns it to the Operator field.

### GetElements

`func (o *CampaignSetNode) GetElements() []CampaignSetNode`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CampaignSetNode) GetElementsOk() ([]CampaignSetNode, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasElements

`func (o *CampaignSetNode) HasElements() bool`

HasElements returns a boolean if a field has been set.

### SetElements

`func (o *CampaignSetNode) SetElements(v []CampaignSetNode)`

SetElements gets a reference to the given []CampaignSetNode and assigns it to the Elements field.

### GetGroupId

`func (o *CampaignSetNode) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CampaignSetNode) GetGroupIdOk() (int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroupId

`func (o *CampaignSetNode) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupId

`func (o *CampaignSetNode) SetGroupId(v int32)`

SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.

### GetLocked

`func (o *CampaignSetNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CampaignSetNode) GetLockedOk() (bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocked

`func (o *CampaignSetNode) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLocked

`func (o *CampaignSetNode) SetLocked(v bool)`

SetLocked gets a reference to the given bool and assigns it to the Locked field.

### GetDescription

`func (o *CampaignSetNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignSetNode) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CampaignSetNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CampaignSetNode) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetEvaluationMode

`func (o *CampaignSetNode) GetEvaluationMode() string`

GetEvaluationMode returns the EvaluationMode field if non-nil, zero value otherwise.

### GetEvaluationModeOk

`func (o *CampaignSetNode) GetEvaluationModeOk() (string, bool)`

GetEvaluationModeOk returns a tuple with the EvaluationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationMode

`func (o *CampaignSetNode) HasEvaluationMode() bool`

HasEvaluationMode returns a boolean if a field has been set.

### SetEvaluationMode

`func (o *CampaignSetNode) SetEvaluationMode(v string)`

SetEvaluationMode gets a reference to the given string and assigns it to the EvaluationMode field.

### GetEvaluationScope

`func (o *CampaignSetNode) GetEvaluationScope() string`

GetEvaluationScope returns the EvaluationScope field if non-nil, zero value otherwise.

### GetEvaluationScopeOk

`func (o *CampaignSetNode) GetEvaluationScopeOk() (string, bool)`

GetEvaluationScopeOk returns a tuple with the EvaluationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationScope

`func (o *CampaignSetNode) HasEvaluationScope() bool`

HasEvaluationScope returns a boolean if a field has been set.

### SetEvaluationScope

`func (o *CampaignSetNode) SetEvaluationScope(v string)`

SetEvaluationScope gets a reference to the given string and assigns it to the EvaluationScope field.

### GetCampaignId

`func (o *CampaignSetNode) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignSetNode) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *CampaignSetNode) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *CampaignSetNode) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


