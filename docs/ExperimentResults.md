# ExperimentResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variants** | Pointer to [**[]ExperimentVariantResult**](ExperimentVariantResult.md) |  | [optional] 
**Confidence** | Pointer to [**ExperimentVariantResultConfidence**](ExperimentVariantResultConfidence.md) |  | 

## Methods

### NewExperimentResults

`func NewExperimentResults(confidence ExperimentVariantResultConfidence, ) *ExperimentResults`

NewExperimentResults instantiates a new ExperimentResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentResultsWithDefaults

`func NewExperimentResultsWithDefaults() *ExperimentResults`

NewExperimentResultsWithDefaults instantiates a new ExperimentResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariants

`func (o *ExperimentResults) GetVariants() []ExperimentVariantResult`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExperimentResults) GetVariantsOk() (*[]ExperimentVariantResult, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExperimentResults) SetVariants(v []ExperimentVariantResult)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ExperimentResults) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetConfidence

`func (o *ExperimentResults) GetConfidence() ExperimentVariantResultConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ExperimentResults) GetConfidenceOk() (*ExperimentVariantResultConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ExperimentResults) SetConfidence(v ExperimentVariantResultConfidence)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


