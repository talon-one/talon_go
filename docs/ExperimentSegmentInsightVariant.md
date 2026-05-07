# ExperimentSegmentInsightVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariantId** | Pointer to **int64** | The ID of the experiment variant. | 
**VariantName** | Pointer to **string** | The name of the experiment variant. | 
**SessionsCount** | Pointer to **int64** | The number of sessions in this segment for this variant. | 
**Value** | Pointer to **float64** | The metric value for this variant in the segment. | 

## Methods

### NewExperimentSegmentInsightVariant

`func NewExperimentSegmentInsightVariant(variantId int64, variantName string, sessionsCount int64, value float64, ) *ExperimentSegmentInsightVariant`

NewExperimentSegmentInsightVariant instantiates a new ExperimentSegmentInsightVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentSegmentInsightVariantWithDefaults

`func NewExperimentSegmentInsightVariantWithDefaults() *ExperimentSegmentInsightVariant`

NewExperimentSegmentInsightVariantWithDefaults instantiates a new ExperimentSegmentInsightVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariantId

`func (o *ExperimentSegmentInsightVariant) GetVariantId() int64`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *ExperimentSegmentInsightVariant) GetVariantIdOk() (*int64, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *ExperimentSegmentInsightVariant) SetVariantId(v int64)`

SetVariantId sets VariantId field to given value.


### GetVariantName

`func (o *ExperimentSegmentInsightVariant) GetVariantName() string`

GetVariantName returns the VariantName field if non-nil, zero value otherwise.

### GetVariantNameOk

`func (o *ExperimentSegmentInsightVariant) GetVariantNameOk() (*string, bool)`

GetVariantNameOk returns a tuple with the VariantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantName

`func (o *ExperimentSegmentInsightVariant) SetVariantName(v string)`

SetVariantName sets VariantName field to given value.


### GetSessionsCount

`func (o *ExperimentSegmentInsightVariant) GetSessionsCount() int64`

GetSessionsCount returns the SessionsCount field if non-nil, zero value otherwise.

### GetSessionsCountOk

`func (o *ExperimentSegmentInsightVariant) GetSessionsCountOk() (*int64, bool)`

GetSessionsCountOk returns a tuple with the SessionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsCount

`func (o *ExperimentSegmentInsightVariant) SetSessionsCount(v int64)`

SetSessionsCount sets SessionsCount field to given value.


### GetValue

`func (o *ExperimentSegmentInsightVariant) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExperimentSegmentInsightVariant) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExperimentSegmentInsightVariant) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


