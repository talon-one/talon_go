# NewExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsVariantAssignmentExternal** | Pointer to **bool** | The source of the assignment. - false - The variant assignment is handled internally by Talon.One. - true - The variant assignment is handled externally.  | 
**Campaign** | Pointer to [**NewCampaign**](NewCampaign.md) |  | 
**GoalType** | Pointer to **string** | The goal of the experiment. Determines which single metric is used to decide the winning variant. When set to &#x60;other&#x60;, multiple metrics are used.  | [default to "other"]
**GoalDescription** | Pointer to **string** | A description of the experiment goal. Provides context for the AI summary and helps it interpret the outcome of the experiment against the stated goal.  | [optional] 

## Methods

### NewNewExperiment

`func NewNewExperiment(isVariantAssignmentExternal bool, campaign NewCampaign, goalType string, ) *NewExperiment`

NewNewExperiment instantiates a new NewExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewExperimentWithDefaults

`func NewNewExperimentWithDefaults() *NewExperiment`

NewNewExperimentWithDefaults instantiates a new NewExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsVariantAssignmentExternal

`func (o *NewExperiment) GetIsVariantAssignmentExternal() bool`

GetIsVariantAssignmentExternal returns the IsVariantAssignmentExternal field if non-nil, zero value otherwise.

### GetIsVariantAssignmentExternalOk

`func (o *NewExperiment) GetIsVariantAssignmentExternalOk() (*bool, bool)`

GetIsVariantAssignmentExternalOk returns a tuple with the IsVariantAssignmentExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariantAssignmentExternal

`func (o *NewExperiment) SetIsVariantAssignmentExternal(v bool)`

SetIsVariantAssignmentExternal sets IsVariantAssignmentExternal field to given value.


### GetCampaign

`func (o *NewExperiment) GetCampaign() NewCampaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *NewExperiment) GetCampaignOk() (*NewCampaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *NewExperiment) SetCampaign(v NewCampaign)`

SetCampaign sets Campaign field to given value.


### GetGoalType

`func (o *NewExperiment) GetGoalType() string`

GetGoalType returns the GoalType field if non-nil, zero value otherwise.

### GetGoalTypeOk

`func (o *NewExperiment) GetGoalTypeOk() (*string, bool)`

GetGoalTypeOk returns a tuple with the GoalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalType

`func (o *NewExperiment) SetGoalType(v string)`

SetGoalType sets GoalType field to given value.


### GetGoalDescription

`func (o *NewExperiment) GetGoalDescription() string`

GetGoalDescription returns the GoalDescription field if non-nil, zero value otherwise.

### GetGoalDescriptionOk

`func (o *NewExperiment) GetGoalDescriptionOk() (*string, bool)`

GetGoalDescriptionOk returns a tuple with the GoalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalDescription

`func (o *NewExperiment) SetGoalDescription(v string)`

SetGoalDescription sets GoalDescription field to given value.

### HasGoalDescription

`func (o *NewExperiment) HasGoalDescription() bool`

HasGoalDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


