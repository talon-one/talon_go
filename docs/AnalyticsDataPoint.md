# AnalyticsDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** |  | 
**Influenced** | Pointer to **float32** |  | 

## Methods

### NewAnalyticsDataPoint

`func NewAnalyticsDataPoint(total float32, influenced float32, ) *AnalyticsDataPoint`

NewAnalyticsDataPoint instantiates a new AnalyticsDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsDataPointWithDefaults

`func NewAnalyticsDataPointWithDefaults() *AnalyticsDataPoint`

NewAnalyticsDataPointWithDefaults instantiates a new AnalyticsDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AnalyticsDataPoint) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AnalyticsDataPoint) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AnalyticsDataPoint) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetInfluenced

`func (o *AnalyticsDataPoint) GetInfluenced() float32`

GetInfluenced returns the Influenced field if non-nil, zero value otherwise.

### GetInfluencedOk

`func (o *AnalyticsDataPoint) GetInfluencedOk() (*float32, bool)`

GetInfluencedOk returns a tuple with the Influenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluenced

`func (o *AnalyticsDataPoint) SetInfluenced(v float32)`

SetInfluenced sets Influenced field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


