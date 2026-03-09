# PromoteExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetApplicationId** | Pointer to **int64** | The ID of the Application to copy the experiment. It is displayed in your Talon.One deployment URL.  | 
**VariantId** | Pointer to **int64** | The ID of the Experiment Variant to build the new campaign.  | 
**DisableExperiment** | Pointer to **bool** | Force disable the experiment.  | [optional] 
**Campaign** | Pointer to [**ExperimentCampaignCopy**](ExperimentCampaignCopy.md) |  | 

## Methods

### NewPromoteExperiment

`func NewPromoteExperiment(targetApplicationId int64, variantId int64, campaign ExperimentCampaignCopy, ) *PromoteExperiment`

NewPromoteExperiment instantiates a new PromoteExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoteExperimentWithDefaults

`func NewPromoteExperimentWithDefaults() *PromoteExperiment`

NewPromoteExperimentWithDefaults instantiates a new PromoteExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetApplicationId

`func (o *PromoteExperiment) GetTargetApplicationId() int64`

GetTargetApplicationId returns the TargetApplicationId field if non-nil, zero value otherwise.

### GetTargetApplicationIdOk

`func (o *PromoteExperiment) GetTargetApplicationIdOk() (*int64, bool)`

GetTargetApplicationIdOk returns a tuple with the TargetApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplicationId

`func (o *PromoteExperiment) SetTargetApplicationId(v int64)`

SetTargetApplicationId sets TargetApplicationId field to given value.


### GetVariantId

`func (o *PromoteExperiment) GetVariantId() int64`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *PromoteExperiment) GetVariantIdOk() (*int64, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *PromoteExperiment) SetVariantId(v int64)`

SetVariantId sets VariantId field to given value.


### GetDisableExperiment

`func (o *PromoteExperiment) GetDisableExperiment() bool`

GetDisableExperiment returns the DisableExperiment field if non-nil, zero value otherwise.

### GetDisableExperimentOk

`func (o *PromoteExperiment) GetDisableExperimentOk() (*bool, bool)`

GetDisableExperimentOk returns a tuple with the DisableExperiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableExperiment

`func (o *PromoteExperiment) SetDisableExperiment(v bool)`

SetDisableExperiment sets DisableExperiment field to given value.

### HasDisableExperiment

`func (o *PromoteExperiment) HasDisableExperiment() bool`

HasDisableExperiment returns a boolean if a field has been set.

### GetCampaign

`func (o *PromoteExperiment) GetCampaign() ExperimentCampaignCopy`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *PromoteExperiment) GetCampaignOk() (*ExperimentCampaignCopy, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *PromoteExperiment) SetCampaign(v ExperimentCampaignCopy)`

SetCampaign sets Campaign field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


