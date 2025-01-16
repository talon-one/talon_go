# ProductSkuUnitAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The start of the aggregation time frame in UTC. | 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the aggregation time frame in UTC. | 
**PurchasedUnits** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | 
**Sku** | Pointer to **string** | The SKU linked to the analytics-level product. | 

## Methods

### GetStartTime

`func (o *ProductSkuUnitAnalytics) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProductSkuUnitAnalytics) GetStartTimeOk() (time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *ProductSkuUnitAnalytics) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *ProductSkuUnitAnalytics) SetStartTime(v time.Time)`

SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.

### GetEndTime

`func (o *ProductSkuUnitAnalytics) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ProductSkuUnitAnalytics) GetEndTimeOk() (time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *ProductSkuUnitAnalytics) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *ProductSkuUnitAnalytics) SetEndTime(v time.Time)`

SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.

### GetPurchasedUnits

`func (o *ProductSkuUnitAnalytics) GetPurchasedUnits() AnalyticsDataPointWithTrend`

GetPurchasedUnits returns the PurchasedUnits field if non-nil, zero value otherwise.

### GetPurchasedUnitsOk

`func (o *ProductSkuUnitAnalytics) GetPurchasedUnitsOk() (AnalyticsDataPointWithTrend, bool)`

GetPurchasedUnitsOk returns a tuple with the PurchasedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPurchasedUnits

`func (o *ProductSkuUnitAnalytics) HasPurchasedUnits() bool`

HasPurchasedUnits returns a boolean if a field has been set.

### SetPurchasedUnits

`func (o *ProductSkuUnitAnalytics) SetPurchasedUnits(v AnalyticsDataPointWithTrend)`

SetPurchasedUnits gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the PurchasedUnits field.

### GetSku

`func (o *ProductSkuUnitAnalytics) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductSkuUnitAnalytics) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *ProductSkuUnitAnalytics) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *ProductSkuUnitAnalytics) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


