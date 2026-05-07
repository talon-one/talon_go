# ExperimentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variants** | Pointer to [**[]ExperimentVariantResult**](ExperimentVariantResult.md) |  | 
**Confidence** | Pointer to [**ExperimentVariantResultConfidence**](ExperimentVariantResultConfidence.md) |  | 
**ExperimentId** | Pointer to **int64** |  | 

## Methods

### NewExperimentResult

`func NewExperimentResult(variants []ExperimentVariantResult, confidence ExperimentVariantResultConfidence, experimentId int64, ) *ExperimentResult`

NewExperimentResult instantiates a new ExperimentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentResultWithDefaults

`func NewExperimentResultWithDefaults() *ExperimentResult`

NewExperimentResultWithDefaults instantiates a new ExperimentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariants

`func (o *ExperimentResult) GetVariants() []ExperimentVariantResult`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExperimentResult) GetVariantsOk() (*[]ExperimentVariantResult, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExperimentResult) SetVariants(v []ExperimentVariantResult)`

SetVariants sets Variants field to given value.


### GetConfidence

`func (o *ExperimentResult) GetConfidence() ExperimentVariantResultConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ExperimentResult) GetConfidenceOk() (*ExperimentVariantResultConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ExperimentResult) SetConfidence(v ExperimentVariantResultConfidence)`

SetConfidence sets Confidence field to given value.


### GetExperimentId

`func (o *ExperimentResult) GetExperimentId() int64`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *ExperimentResult) GetExperimentIdOk() (*int64, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *ExperimentResult) SetExperimentId(v int64)`

SetExperimentId sets ExperimentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


