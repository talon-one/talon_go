# ExperimentCopyExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentType** | Pointer to **string** | Controls how customers are assigned to experiment variants in the copied experiment. - &#x60;random&#x60;: Talon.One assigns customers randomly based on variant weights. - &#x60;external&#x60;: The variant assignment is handled externally. - &#x60;audience&#x60;: Each variant targets a specific audience; customers are assigned based on audience membership. This is the source of truth. When omitted, it is derived from the deprecated &#x60;isVariantAssignmentExternal&#x60; flag (&#x60;true&#x60; maps to &#x60;external&#x60;, otherwise &#x60;random&#x60;).  | [optional] 
**IsVariantAssignmentExternal** | Pointer to **bool** | The source of the assignment. - false - The variant assignment is handled internally by Talon.One. - true - The variant assignment is handled externally. Deprecated: use &#x60;assignmentType&#x60; instead. Kept for backwards compatibility with older clients; when set and &#x60;assignmentType&#x60; is omitted, &#x60;true&#x60; maps to &#x60;external&#x60;.  | [optional] 
**Campaign** | Pointer to [**ExperimentCampaignCopy**](ExperimentCampaignCopy.md) |  | 
**GoalType** | Pointer to **string** | The goal of the experiment. Determines which single metric is used to decide the winning variant. When set to &#x60;other&#x60;, multiple metrics are used. If omitted, the value from the source experiment is used.  | [optional] 
**GoalDescription** | Pointer to **string** | A description of the experiment goal. Provides context for the AI summary and helps it interpret the outcome of the experiment against the stated goal. If omitted, the value from the source experiment is used.  | [optional] 

## Methods

### NewExperimentCopyExperiment

`func NewExperimentCopyExperiment(campaign ExperimentCampaignCopy, ) *ExperimentCopyExperiment`

NewExperimentCopyExperiment instantiates a new ExperimentCopyExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentCopyExperimentWithDefaults

`func NewExperimentCopyExperimentWithDefaults() *ExperimentCopyExperiment`

NewExperimentCopyExperimentWithDefaults instantiates a new ExperimentCopyExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentType

`func (o *ExperimentCopyExperiment) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *ExperimentCopyExperiment) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *ExperimentCopyExperiment) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *ExperimentCopyExperiment) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetIsVariantAssignmentExternal

`func (o *ExperimentCopyExperiment) GetIsVariantAssignmentExternal() bool`

GetIsVariantAssignmentExternal returns the IsVariantAssignmentExternal field if non-nil, zero value otherwise.

### GetIsVariantAssignmentExternalOk

`func (o *ExperimentCopyExperiment) GetIsVariantAssignmentExternalOk() (*bool, bool)`

GetIsVariantAssignmentExternalOk returns a tuple with the IsVariantAssignmentExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariantAssignmentExternal

`func (o *ExperimentCopyExperiment) SetIsVariantAssignmentExternal(v bool)`

SetIsVariantAssignmentExternal sets IsVariantAssignmentExternal field to given value.

### HasIsVariantAssignmentExternal

`func (o *ExperimentCopyExperiment) HasIsVariantAssignmentExternal() bool`

HasIsVariantAssignmentExternal returns a boolean if a field has been set.

### GetCampaign

`func (o *ExperimentCopyExperiment) GetCampaign() ExperimentCampaignCopy`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *ExperimentCopyExperiment) GetCampaignOk() (*ExperimentCampaignCopy, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *ExperimentCopyExperiment) SetCampaign(v ExperimentCampaignCopy)`

SetCampaign sets Campaign field to given value.


### GetGoalType

`func (o *ExperimentCopyExperiment) GetGoalType() string`

GetGoalType returns the GoalType field if non-nil, zero value otherwise.

### GetGoalTypeOk

`func (o *ExperimentCopyExperiment) GetGoalTypeOk() (*string, bool)`

GetGoalTypeOk returns a tuple with the GoalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalType

`func (o *ExperimentCopyExperiment) SetGoalType(v string)`

SetGoalType sets GoalType field to given value.

### HasGoalType

`func (o *ExperimentCopyExperiment) HasGoalType() bool`

HasGoalType returns a boolean if a field has been set.

### GetGoalDescription

`func (o *ExperimentCopyExperiment) GetGoalDescription() string`

GetGoalDescription returns the GoalDescription field if non-nil, zero value otherwise.

### GetGoalDescriptionOk

`func (o *ExperimentCopyExperiment) GetGoalDescriptionOk() (*string, bool)`

GetGoalDescriptionOk returns a tuple with the GoalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalDescription

`func (o *ExperimentCopyExperiment) SetGoalDescription(v string)`

SetGoalDescription sets GoalDescription field to given value.

### HasGoalDescription

`func (o *ExperimentCopyExperiment) HasGoalDescription() bool`

HasGoalDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


