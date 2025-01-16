# ProductUnitAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start of the aggregation time frame in UTC. | 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the aggregation time frame in UTC. | 
**PurchasedUnits** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | 
**ProductId** | Pointer to **int32** | The ID of the analytics-level product. | 
**ProductName** | Pointer to **string** | The name of the analytics-level product. | 

## Methods

### GetStartTime

`func (o *ProductUnitAnalytics) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProductUnitAnalytics) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *ProductUnitAnalytics) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *ProductUnitAnalytics) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *ProductUnitAnalytics) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ProductUnitAnalytics) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *ProductUnitAnalytics) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *ProductUnitAnalytics) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetPurchasedUnits

`func (o *ProductUnitAnalytics) GetPurchasedUnits() AnalyticsDataPointWithTrend`

GetPurchasedUnits returns the PurchasedUnits field if non-nil, zero value otherwise.

### GetPurchasedUnitsOk

`func (o *ProductUnitAnalytics) GetPurchasedUnitsOk() (AnalyticsDataPointWithTrend, bool)`

GetPurchasedUnitsOk returns a tuple with the PurchasedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPurchasedUnits

`func (o *ProductUnitAnalytics) HasPurchasedUnits() bool`

HasPurchasedUnits returns a boolean if a field has been set.

### SetPurchasedUnits

`func (o *ProductUnitAnalytics) SetPurchasedUnits(v AnalyticsDataPointWithTrend)`

SetPurchasedUnits gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the PurchasedUnits field.

### GetProductId

`func (o *ProductUnitAnalytics) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductUnitAnalytics) GetProductIdOk() (int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductId

`func (o *ProductUnitAnalytics) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductId

`func (o *ProductUnitAnalytics) SetProductId(v int32)`

SetProductId gets a reference to the given int32 and assigns it to the ProductId field.

### GetProductName

`func (o *ProductUnitAnalytics) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ProductUnitAnalytics) GetProductNameOk() (string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductName

`func (o *ProductUnitAnalytics) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductName

`func (o *ProductUnitAnalytics) SetProductName(v string)`

SetProductName gets a reference to the given string and assigns it to the ProductName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


