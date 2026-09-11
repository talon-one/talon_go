# ExperimentConfidenceTimelineDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**time.Time**](time.Time.md) | The date-time this data point represents. | 
**Confidence** | Pointer to [**ExperimentVariantResultConfidence**](ExperimentVariantResultConfidence.md) |  | 

## Methods

### NewExperimentConfidenceTimelineDataPoint

`func NewExperimentConfidenceTimelineDataPoint(date time.Time, confidence ExperimentVariantResultConfidence, ) *ExperimentConfidenceTimelineDataPoint`

NewExperimentConfidenceTimelineDataPoint instantiates a new ExperimentConfidenceTimelineDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentConfidenceTimelineDataPointWithDefaults

`func NewExperimentConfidenceTimelineDataPointWithDefaults() *ExperimentConfidenceTimelineDataPoint`

NewExperimentConfidenceTimelineDataPointWithDefaults instantiates a new ExperimentConfidenceTimelineDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ExperimentConfidenceTimelineDataPoint) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ExperimentConfidenceTimelineDataPoint) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ExperimentConfidenceTimelineDataPoint) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetConfidence

`func (o *ExperimentConfidenceTimelineDataPoint) GetConfidence() ExperimentVariantResultConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ExperimentConfidenceTimelineDataPoint) GetConfidenceOk() (*ExperimentVariantResultConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ExperimentConfidenceTimelineDataPoint) SetConfidence(v ExperimentVariantResultConfidence)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


