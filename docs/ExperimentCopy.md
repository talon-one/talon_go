# ExperimentCopy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetApplicationId** | Pointer to **int64** | The ID of the Application to copy the experiment. It is displayed in your Talon.One deployment URL.  | 
**Experiment** | Pointer to [**ExperimentCopyExperiment**](ExperimentCopy_experiment.md) |  | 

## Methods

### NewExperimentCopy

`func NewExperimentCopy(targetApplicationId int64, experiment ExperimentCopyExperiment, ) *ExperimentCopy`

NewExperimentCopy instantiates a new ExperimentCopy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentCopyWithDefaults

`func NewExperimentCopyWithDefaults() *ExperimentCopy`

NewExperimentCopyWithDefaults instantiates a new ExperimentCopy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetApplicationId

`func (o *ExperimentCopy) GetTargetApplicationId() int64`

GetTargetApplicationId returns the TargetApplicationId field if non-nil, zero value otherwise.

### GetTargetApplicationIdOk

`func (o *ExperimentCopy) GetTargetApplicationIdOk() (*int64, bool)`

GetTargetApplicationIdOk returns a tuple with the TargetApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplicationId

`func (o *ExperimentCopy) SetTargetApplicationId(v int64)`

SetTargetApplicationId sets TargetApplicationId field to given value.


### GetExperiment

`func (o *ExperimentCopy) GetExperiment() ExperimentCopyExperiment`

GetExperiment returns the Experiment field if non-nil, zero value otherwise.

### GetExperimentOk

`func (o *ExperimentCopy) GetExperimentOk() (*ExperimentCopyExperiment, bool)`

GetExperimentOk returns a tuple with the Experiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiment

`func (o *ExperimentCopy) SetExperiment(v ExperimentCopyExperiment)`

SetExperiment sets Experiment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


