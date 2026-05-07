# ExperimentVariantResultConfidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgSessionValue** | Pointer to **float32** | The calculated confidence value of the average customer session value. | 
**AvgDiscountedSessionValue** | Pointer to **float32** | The calculated confidence value of the average customer discounted session value. | 
**AvgItemsPerSession** | Pointer to **float32** | The calculated confidence value of the number of items from sessions value. | 

## Methods

### NewExperimentVariantResultConfidence

`func NewExperimentVariantResultConfidence(avgSessionValue float32, avgDiscountedSessionValue float32, avgItemsPerSession float32, ) *ExperimentVariantResultConfidence`

NewExperimentVariantResultConfidence instantiates a new ExperimentVariantResultConfidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentVariantResultConfidenceWithDefaults

`func NewExperimentVariantResultConfidenceWithDefaults() *ExperimentVariantResultConfidence`

NewExperimentVariantResultConfidenceWithDefaults instantiates a new ExperimentVariantResultConfidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgSessionValue

`func (o *ExperimentVariantResultConfidence) GetAvgSessionValue() float32`

GetAvgSessionValue returns the AvgSessionValue field if non-nil, zero value otherwise.

### GetAvgSessionValueOk

`func (o *ExperimentVariantResultConfidence) GetAvgSessionValueOk() (*float32, bool)`

GetAvgSessionValueOk returns a tuple with the AvgSessionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSessionValue

`func (o *ExperimentVariantResultConfidence) SetAvgSessionValue(v float32)`

SetAvgSessionValue sets AvgSessionValue field to given value.


### GetAvgDiscountedSessionValue

`func (o *ExperimentVariantResultConfidence) GetAvgDiscountedSessionValue() float32`

GetAvgDiscountedSessionValue returns the AvgDiscountedSessionValue field if non-nil, zero value otherwise.

### GetAvgDiscountedSessionValueOk

`func (o *ExperimentVariantResultConfidence) GetAvgDiscountedSessionValueOk() (*float32, bool)`

GetAvgDiscountedSessionValueOk returns a tuple with the AvgDiscountedSessionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDiscountedSessionValue

`func (o *ExperimentVariantResultConfidence) SetAvgDiscountedSessionValue(v float32)`

SetAvgDiscountedSessionValue sets AvgDiscountedSessionValue field to given value.


### GetAvgItemsPerSession

`func (o *ExperimentVariantResultConfidence) GetAvgItemsPerSession() float32`

GetAvgItemsPerSession returns the AvgItemsPerSession field if non-nil, zero value otherwise.

### GetAvgItemsPerSessionOk

`func (o *ExperimentVariantResultConfidence) GetAvgItemsPerSessionOk() (*float32, bool)`

GetAvgItemsPerSessionOk returns a tuple with the AvgItemsPerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgItemsPerSession

`func (o *ExperimentVariantResultConfidence) SetAvgItemsPerSession(v float32)`

SetAvgItemsPerSession sets AvgItemsPerSession field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


