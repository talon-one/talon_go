# ExperimentSegmentInsights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**[]ExperimentSegmentInsightMetric**](ExperimentSegmentInsightMetric.md) | Segment insights grouped by metric. This array always contains exactly three metric objects. Each metric includes a segments array, which is empty if no significant results were found. The metrics array itself is empty if the &#x60;reason&#x60; property is populated.  | 
**TotalSegmentsTested** | Pointer to **int64** | Total number of segment-metric combinations that were tested for statistical significance.  | 
**DimensionsTested** | Pointer to **int64** | Number of segmentation dimensions that had sufficient data variance to test. Dimensions where all sessions fall into a single bucket are excluded.  | 
**Reason** | Pointer to **string** | Empty string when segment insights are available. Contains a reason code when insights could not be computed (e.g., \&quot;insufficient_data\&quot; when the experiment has fewer than 100 sessions per variant).  | 

## Methods

### NewExperimentSegmentInsights

`func NewExperimentSegmentInsights(metrics []ExperimentSegmentInsightMetric, totalSegmentsTested int64, dimensionsTested int64, reason string, ) *ExperimentSegmentInsights`

NewExperimentSegmentInsights instantiates a new ExperimentSegmentInsights object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentSegmentInsightsWithDefaults

`func NewExperimentSegmentInsightsWithDefaults() *ExperimentSegmentInsights`

NewExperimentSegmentInsightsWithDefaults instantiates a new ExperimentSegmentInsights object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *ExperimentSegmentInsights) GetMetrics() []ExperimentSegmentInsightMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ExperimentSegmentInsights) GetMetricsOk() (*[]ExperimentSegmentInsightMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ExperimentSegmentInsights) SetMetrics(v []ExperimentSegmentInsightMetric)`

SetMetrics sets Metrics field to given value.


### GetTotalSegmentsTested

`func (o *ExperimentSegmentInsights) GetTotalSegmentsTested() int64`

GetTotalSegmentsTested returns the TotalSegmentsTested field if non-nil, zero value otherwise.

### GetTotalSegmentsTestedOk

`func (o *ExperimentSegmentInsights) GetTotalSegmentsTestedOk() (*int64, bool)`

GetTotalSegmentsTestedOk returns a tuple with the TotalSegmentsTested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSegmentsTested

`func (o *ExperimentSegmentInsights) SetTotalSegmentsTested(v int64)`

SetTotalSegmentsTested sets TotalSegmentsTested field to given value.


### GetDimensionsTested

`func (o *ExperimentSegmentInsights) GetDimensionsTested() int64`

GetDimensionsTested returns the DimensionsTested field if non-nil, zero value otherwise.

### GetDimensionsTestedOk

`func (o *ExperimentSegmentInsights) GetDimensionsTestedOk() (*int64, bool)`

GetDimensionsTestedOk returns a tuple with the DimensionsTested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsTested

`func (o *ExperimentSegmentInsights) SetDimensionsTested(v int64)`

SetDimensionsTested sets DimensionsTested field to given value.


### GetReason

`func (o *ExperimentSegmentInsights) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ExperimentSegmentInsights) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ExperimentSegmentInsights) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


