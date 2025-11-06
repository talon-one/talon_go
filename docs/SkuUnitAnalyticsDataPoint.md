# SkuUnitAnalyticsDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start of the aggregation time frame in UTC. | 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the aggregation time frame in UTC. | 
**UnitsSold** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | 
**Sku** | Pointer to **string** | The SKU linked to the application. | 

## Methods

### NewSkuUnitAnalyticsDataPoint

`func NewSkuUnitAnalyticsDataPoint(startTime time.Time, endTime time.Time, unitsSold AnalyticsDataPointWithTrend, sku string, ) *SkuUnitAnalyticsDataPoint`

NewSkuUnitAnalyticsDataPoint instantiates a new SkuUnitAnalyticsDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuUnitAnalyticsDataPointWithDefaults

`func NewSkuUnitAnalyticsDataPointWithDefaults() *SkuUnitAnalyticsDataPoint`

NewSkuUnitAnalyticsDataPointWithDefaults instantiates a new SkuUnitAnalyticsDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *SkuUnitAnalyticsDataPoint) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SkuUnitAnalyticsDataPoint) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SkuUnitAnalyticsDataPoint) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SkuUnitAnalyticsDataPoint) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SkuUnitAnalyticsDataPoint) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SkuUnitAnalyticsDataPoint) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetUnitsSold

`func (o *SkuUnitAnalyticsDataPoint) GetUnitsSold() AnalyticsDataPointWithTrend`

GetUnitsSold returns the UnitsSold field if non-nil, zero value otherwise.

### GetUnitsSoldOk

`func (o *SkuUnitAnalyticsDataPoint) GetUnitsSoldOk() (*AnalyticsDataPointWithTrend, bool)`

GetUnitsSoldOk returns a tuple with the UnitsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsSold

`func (o *SkuUnitAnalyticsDataPoint) SetUnitsSold(v AnalyticsDataPointWithTrend)`

SetUnitsSold sets UnitsSold field to given value.


### GetSku

`func (o *SkuUnitAnalyticsDataPoint) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *SkuUnitAnalyticsDataPoint) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *SkuUnitAnalyticsDataPoint) SetSku(v string)`

SetSku sets Sku field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


