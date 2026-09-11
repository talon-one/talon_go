# UpdateExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsVariantAssignmentExternal** | Pointer to **bool** | The source of the assignment. - false - The variant assignment is handled internally by Talon.One. - true - The variant assignment is handled externally.  | 
**Campaign** | Pointer to [**UpdateCampaign**](UpdateCampaign.md) |  | 
**GoalType** | Pointer to **string** | The goal of the experiment. Determines which single metric is used to decide the winning variant. When set to &#x60;other&#x60;, multiple metrics are used. If omitted, the current value is preserved.  | [optional] 
**GoalDescription** | Pointer to **string** | A description of the experiment goal. Provides context for the AI summary and helps it interpret the outcome of the experiment against the stated goal. If omitted, the current value is preserved.  | [optional] 

## Methods

### NewUpdateExperiment

`func NewUpdateExperiment(isVariantAssignmentExternal bool, campaign UpdateCampaign, ) *UpdateExperiment`

NewUpdateExperiment instantiates a new UpdateExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExperimentWithDefaults

`func NewUpdateExperimentWithDefaults() *UpdateExperiment`

NewUpdateExperimentWithDefaults instantiates a new UpdateExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsVariantAssignmentExternal

`func (o *UpdateExperiment) GetIsVariantAssignmentExternal() bool`

GetIsVariantAssignmentExternal returns the IsVariantAssignmentExternal field if non-nil, zero value otherwise.

### GetIsVariantAssignmentExternalOk

`func (o *UpdateExperiment) GetIsVariantAssignmentExternalOk() (*bool, bool)`

GetIsVariantAssignmentExternalOk returns a tuple with the IsVariantAssignmentExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariantAssignmentExternal

`func (o *UpdateExperiment) SetIsVariantAssignmentExternal(v bool)`

SetIsVariantAssignmentExternal sets IsVariantAssignmentExternal field to given value.


### GetCampaign

`func (o *UpdateExperiment) GetCampaign() UpdateCampaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *UpdateExperiment) GetCampaignOk() (*UpdateCampaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *UpdateExperiment) SetCampaign(v UpdateCampaign)`

SetCampaign sets Campaign field to given value.


### GetGoalType

`func (o *UpdateExperiment) GetGoalType() string`

GetGoalType returns the GoalType field if non-nil, zero value otherwise.

### GetGoalTypeOk

`func (o *UpdateExperiment) GetGoalTypeOk() (*string, bool)`

GetGoalTypeOk returns a tuple with the GoalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalType

`func (o *UpdateExperiment) SetGoalType(v string)`

SetGoalType sets GoalType field to given value.

### HasGoalType

`func (o *UpdateExperiment) HasGoalType() bool`

HasGoalType returns a boolean if a field has been set.

### GetGoalDescription

`func (o *UpdateExperiment) GetGoalDescription() string`

GetGoalDescription returns the GoalDescription field if non-nil, zero value otherwise.

### GetGoalDescriptionOk

`func (o *UpdateExperiment) GetGoalDescriptionOk() (*string, bool)`

GetGoalDescriptionOk returns a tuple with the GoalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalDescription

`func (o *UpdateExperiment) SetGoalDescription(v string)`

SetGoalDescription sets GoalDescription field to given value.

### HasGoalDescription

`func (o *UpdateExperiment) HasGoalDescription() bool`

HasGoalDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


