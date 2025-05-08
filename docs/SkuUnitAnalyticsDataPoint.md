# SkuUnitAnalyticsDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start of the aggregation time frame in UTC. | 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the aggregation time frame in UTC. | 
**UnitsSold** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | 
**Sku** | Pointer to **string** | The SKU linked to the application. | 

## Methods

### GetStartTime

`func (o *SkuUnitAnalyticsDataPoint) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SkuUnitAnalyticsDataPoint) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *SkuUnitAnalyticsDataPoint) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *SkuUnitAnalyticsDataPoint) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *SkuUnitAnalyticsDataPoint) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SkuUnitAnalyticsDataPoint) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *SkuUnitAnalyticsDataPoint) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *SkuUnitAnalyticsDataPoint) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetUnitsSold

`func (o *SkuUnitAnalyticsDataPoint) GetUnitsSold() AnalyticsDataPointWithTrend`

GetUnitsSold returns the UnitsSold field if non-nil, zero value otherwise.

### GetUnitsSoldOk

`func (o *SkuUnitAnalyticsDataPoint) GetUnitsSoldOk() (AnalyticsDataPointWithTrend, bool)`

GetUnitsSoldOk returns a tuple with the UnitsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnitsSold

`func (o *SkuUnitAnalyticsDataPoint) HasUnitsSold() bool`

HasUnitsSold returns a boolean if a field has been set.

### SetUnitsSold

`func (o *SkuUnitAnalyticsDataPoint) SetUnitsSold(v AnalyticsDataPointWithTrend)`

SetUnitsSold gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the UnitsSold field.

### GetSku

`func (o *SkuUnitAnalyticsDataPoint) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *SkuUnitAnalyticsDataPoint) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *SkuUnitAnalyticsDataPoint) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *SkuUnitAnalyticsDataPoint) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


