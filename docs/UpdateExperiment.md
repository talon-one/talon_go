# UpdateExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsVariantAssignmentExternal** | Pointer to **bool** | The source of the assignment. - false - The variant assignment is handled internally by Talon.One. - true - The variant assignment is handled externally.  | 
**Campaign** | Pointer to [**UpdateCampaign**](UpdateCampaign.md) |  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


