# CampaignEvaluationGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**Description** | Pointer to **string** | A description of the campaign evaluation group. | [optional] 
**EvaluationMode** | Pointer to **string** | The mode by which campaigns in the campaign evaluation group are evaluated. | 
**EvaluationScope** | Pointer to **string** | The evaluation scope of the campaign evaluation group. | 
**Id** | Pointer to **int32** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 
**Locked** | Pointer to **bool** | An indicator of whether the campaign evaluation group is locked for modification. | 
**Name** | Pointer to **string** | The name of the campaign evaluation group. | 
**ParentId** | Pointer to **int32** | The ID of the parent group that contains the campaign evaluation group. | 

## Methods

### GetApplicationId

`func (o *CampaignEvaluationGroup) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignEvaluationGroup) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CampaignEvaluationGroup) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CampaignEvaluationGroup) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetDescription

`func (o *CampaignEvaluationGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignEvaluationGroup) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CampaignEvaluationGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CampaignEvaluationGroup) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetEvaluationMode

`func (o *CampaignEvaluationGroup) GetEvaluationMode() string`

GetEvaluationMode returns the EvaluationMode field if non-nil, zero value otherwise.

### GetEvaluationModeOk

`func (o *CampaignEvaluationGroup) GetEvaluationModeOk() (string, bool)`

GetEvaluationModeOk returns a tuple with the EvaluationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationMode

`func (o *CampaignEvaluationGroup) HasEvaluationMode() bool`

HasEvaluationMode returns a boolean if a field has been set.

### SetEvaluationMode

`func (o *CampaignEvaluationGroup) SetEvaluationMode(v string)`

SetEvaluationMode gets a reference to the given string and assigns it to the EvaluationMode field.

### GetEvaluationScope

`func (o *CampaignEvaluationGroup) GetEvaluationScope() string`

GetEvaluationScope returns the EvaluationScope field if non-nil, zero value otherwise.

### GetEvaluationScopeOk

`func (o *CampaignEvaluationGroup) GetEvaluationScopeOk() (string, bool)`

GetEvaluationScopeOk returns a tuple with the EvaluationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationScope

`func (o *CampaignEvaluationGroup) HasEvaluationScope() bool`

HasEvaluationScope returns a boolean if a field has been set.

### SetEvaluationScope

`func (o *CampaignEvaluationGroup) SetEvaluationScope(v string)`

SetEvaluationScope gets a reference to the given string and assigns it to the EvaluationScope field.

### GetId

`func (o *CampaignEvaluationGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignEvaluationGroup) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CampaignEvaluationGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CampaignEvaluationGroup) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetLocked

`func (o *CampaignEvaluationGroup) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CampaignEvaluationGroup) GetLockedOk() (bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocked

`func (o *CampaignEvaluationGroup) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLocked

`func (o *CampaignEvaluationGroup) SetLocked(v bool)`

SetLocked gets a reference to the given bool and assigns it to the Locked field.

### GetName

`func (o *CampaignEvaluationGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignEvaluationGroup) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CampaignEvaluationGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CampaignEvaluationGroup) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetParentId

`func (o *CampaignEvaluationGroup) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CampaignEvaluationGroup) GetParentIdOk() (int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentId

`func (o *CampaignEvaluationGroup) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentId

`func (o *CampaignEvaluationGroup) SetParentId(v int32)`

SetParentId gets a reference to the given int32 and assigns it to the ParentId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


