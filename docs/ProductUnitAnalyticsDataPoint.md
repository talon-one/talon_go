# ProductUnitAnalyticsDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start of the aggregation time frame in UTC. | 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the aggregation time frame in UTC. | 
**UnitsSold** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | 
**ProductId** | Pointer to **int32** | The ID of the product. | 
**ProductName** | Pointer to **string** | The name of the product. | 

## Methods

### GetStartTime

`func (o *ProductUnitAnalyticsDataPoint) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProductUnitAnalyticsDataPoint) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *ProductUnitAnalyticsDataPoint) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *ProductUnitAnalyticsDataPoint) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *ProductUnitAnalyticsDataPoint) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ProductUnitAnalyticsDataPoint) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *ProductUnitAnalyticsDataPoint) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *ProductUnitAnalyticsDataPoint) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetUnitsSold

`func (o *ProductUnitAnalyticsDataPoint) GetUnitsSold() AnalyticsDataPointWithTrend`

GetUnitsSold returns the UnitsSold field if non-nil, zero value otherwise.

### GetUnitsSoldOk

`func (o *ProductUnitAnalyticsDataPoint) GetUnitsSoldOk() (AnalyticsDataPointWithTrend, bool)`

GetUnitsSoldOk returns a tuple with the UnitsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnitsSold

`func (o *ProductUnitAnalyticsDataPoint) HasUnitsSold() bool`

HasUnitsSold returns a boolean if a field has been set.

### SetUnitsSold

`func (o *ProductUnitAnalyticsDataPoint) SetUnitsSold(v AnalyticsDataPointWithTrend)`

SetUnitsSold gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the UnitsSold field.

### GetProductId

`func (o *ProductUnitAnalyticsDataPoint) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductUnitAnalyticsDataPoint) GetProductIdOk() (int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductId

`func (o *ProductUnitAnalyticsDataPoint) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductId

`func (o *ProductUnitAnalyticsDataPoint) SetProductId(v int32)`

SetProductId gets a reference to the given int32 and assigns it to the ProductId field.

### GetProductName

`func (o *ProductUnitAnalyticsDataPoint) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ProductUnitAnalyticsDataPoint) GetProductNameOk() (string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductName

`func (o *ProductUnitAnalyticsDataPoint) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductName

`func (o *ProductUnitAnalyticsDataPoint) SetProductName(v string)`

SetProductName gets a reference to the given string and assigns it to the ProductName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


