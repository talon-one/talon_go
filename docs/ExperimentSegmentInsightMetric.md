# ExperimentSegmentInsightMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** | The metric being measured. | 
**Segments** | Pointer to [**[]ExperimentSegmentInsight**](ExperimentSegmentInsight.md) | Segments with statistically significant results for this metric. An empty array means no significant segments were found. Segments are sorted by confidence score from highest to lowest.  | 

## Methods

### NewExperimentSegmentInsightMetric

`func NewExperimentSegmentInsightMetric(metric string, segments []ExperimentSegmentInsight, ) *ExperimentSegmentInsightMetric`

NewExperimentSegmentInsightMetric instantiates a new ExperimentSegmentInsightMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentSegmentInsightMetricWithDefaults

`func NewExperimentSegmentInsightMetricWithDefaults() *ExperimentSegmentInsightMetric`

NewExperimentSegmentInsightMetricWithDefaults instantiates a new ExperimentSegmentInsightMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *ExperimentSegmentInsightMetric) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ExperimentSegmentInsightMetric) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ExperimentSegmentInsightMetric) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetSegments

`func (o *ExperimentSegmentInsightMetric) GetSegments() []ExperimentSegmentInsight`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *ExperimentSegmentInsightMetric) GetSegmentsOk() (*[]ExperimentSegmentInsight, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *ExperimentSegmentInsightMetric) SetSegments(v []ExperimentSegmentInsight)`

SetSegments sets Segments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


