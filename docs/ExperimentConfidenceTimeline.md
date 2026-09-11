# ExperimentConfidenceTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ExperimentConfidenceTimelineDataPoint**](ExperimentConfidenceTimelineDataPoint.md) | Daily cumulative confidence values ordered chronologically from experiment start to end, or to today if the experiment is still running. Empty if the experiment has no data yet.  | 

## Methods

### NewExperimentConfidenceTimeline

`func NewExperimentConfidenceTimeline(data []ExperimentConfidenceTimelineDataPoint, ) *ExperimentConfidenceTimeline`

NewExperimentConfidenceTimeline instantiates a new ExperimentConfidenceTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentConfidenceTimelineWithDefaults

`func NewExperimentConfidenceTimelineWithDefaults() *ExperimentConfidenceTimeline`

NewExperimentConfidenceTimelineWithDefaults instantiates a new ExperimentConfidenceTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExperimentConfidenceTimeline) GetData() []ExperimentConfidenceTimelineDataPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExperimentConfidenceTimeline) GetDataOk() (*[]ExperimentConfidenceTimelineDataPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExperimentConfidenceTimeline) SetData(v []ExperimentConfidenceTimelineDataPoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


