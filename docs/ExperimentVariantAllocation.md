# ExperimentVariantAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentID** | Pointer to **int64** | The ID of the experiment. | 
**VariantID** | Pointer to **int64** | The ID of the variant to be allocated. | 

## Methods

### NewExperimentVariantAllocation

`func NewExperimentVariantAllocation(experimentID int64, variantID int64, ) *ExperimentVariantAllocation`

NewExperimentVariantAllocation instantiates a new ExperimentVariantAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentVariantAllocationWithDefaults

`func NewExperimentVariantAllocationWithDefaults() *ExperimentVariantAllocation`

NewExperimentVariantAllocationWithDefaults instantiates a new ExperimentVariantAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentID

`func (o *ExperimentVariantAllocation) GetExperimentID() int64`

GetExperimentID returns the ExperimentID field if non-nil, zero value otherwise.

### GetExperimentIDOk

`func (o *ExperimentVariantAllocation) GetExperimentIDOk() (*int64, bool)`

GetExperimentIDOk returns a tuple with the ExperimentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentID

`func (o *ExperimentVariantAllocation) SetExperimentID(v int64)`

SetExperimentID sets ExperimentID field to given value.


### GetVariantID

`func (o *ExperimentVariantAllocation) GetVariantID() int64`

GetVariantID returns the VariantID field if non-nil, zero value otherwise.

### GetVariantIDOk

`func (o *ExperimentVariantAllocation) GetVariantIDOk() (*int64, bool)`

GetVariantIDOk returns a tuple with the VariantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantID

`func (o *ExperimentVariantAllocation) SetVariantID(v int64)`

SetVariantID sets VariantID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


