# ExperimentCopyExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsVariantAssignmentExternal** | Pointer to **bool** | The source of the assignment. - false - The variant assignment is handled internally by Talon.One. - true - The variant assignment is handled externally.  | 
**Campaign** | Pointer to [**ExperimentCampaignCopy**](ExperimentCampaignCopy.md) |  | 

## Methods

### NewExperimentCopyExperiment

`func NewExperimentCopyExperiment(isVariantAssignmentExternal bool, campaign ExperimentCampaignCopy, ) *ExperimentCopyExperiment`

NewExperimentCopyExperiment instantiates a new ExperimentCopyExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentCopyExperimentWithDefaults

`func NewExperimentCopyExperimentWithDefaults() *ExperimentCopyExperiment`

NewExperimentCopyExperimentWithDefaults instantiates a new ExperimentCopyExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


