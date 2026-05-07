# ExperimentSegmentInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **string** | The segmentation dimension used to group customers or purchases for analysis. | 
**Bucket** | Pointer to **string** | The specific group within the segmentation dimension. | 
**Confidence** | Pointer to **float64** | The raw (unadjusted) confidence score expressed as a percentage. Only segments with a confidence score greater than or equal to 95% are returned.  | 
**WinnerVariantId** | Pointer to **int64** | The ID of the variant that performed better in this segment. | 
**Variants** | Pointer to [**[]ExperimentSegmentInsightVariant**](ExperimentSegmentInsightVariant.md) | Per-variant metric values for this segment. | 

## Methods

### NewExperimentSegmentInsight

`func NewExperimentSegmentInsight(dimension string, bucket string, confidence float64, winnerVariantId int64, variants []ExperimentSegmentInsightVariant, ) *ExperimentSegmentInsight`

NewExperimentSegmentInsight instantiates a new ExperimentSegmentInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentSegmentInsightWithDefaults

`func NewExperimentSegmentInsightWithDefaults() *ExperimentSegmentInsight`

NewExperimentSegmentInsightWithDefaults instantiates a new ExperimentSegmentInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *ExperimentSegmentInsight) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *ExperimentSegmentInsight) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *ExperimentSegmentInsight) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetBucket

`func (o *ExperimentSegmentInsight) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ExperimentSegmentInsight) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ExperimentSegmentInsight) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetConfidence

`func (o *ExperimentSegmentInsight) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ExperimentSegmentInsight) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ExperimentSegmentInsight) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.


### GetWinnerVariantId

`func (o *ExperimentSegmentInsight) GetWinnerVariantId() int64`

GetWinnerVariantId returns the WinnerVariantId field if non-nil, zero value otherwise.

### GetWinnerVariantIdOk

`func (o *ExperimentSegmentInsight) GetWinnerVariantIdOk() (*int64, bool)`

GetWinnerVariantIdOk returns a tuple with the WinnerVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerVariantId

`func (o *ExperimentSegmentInsight) SetWinnerVariantId(v int64)`

SetWinnerVariantId sets WinnerVariantId field to given value.


### GetVariants

`func (o *ExperimentSegmentInsight) GetVariants() []ExperimentSegmentInsightVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExperimentSegmentInsight) GetVariantsOk() (*[]ExperimentSegmentInsightVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExperimentSegmentInsight) SetVariants(v []ExperimentSegmentInsightVariant)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


