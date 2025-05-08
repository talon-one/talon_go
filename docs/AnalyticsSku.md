# AnalyticsSku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the SKU linked to the application. | 
**Sku** | Pointer to **string** | The SKU linked to the application. | 
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) | Values in UTC for the date the SKU linked to the product was last updated. | [optional] 
**UnitsSold** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | [optional] 

## Methods

### GetId

`func (o *AnalyticsSku) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsSku) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *AnalyticsSku) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *AnalyticsSku) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetSku

`func (o *AnalyticsSku) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AnalyticsSku) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *AnalyticsSku) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *AnalyticsSku) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.

### GetLastUpdated

`func (o *AnalyticsSku) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AnalyticsSku) GetLastUpdatedOk() (time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdated

`func (o *AnalyticsSku) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdated

`func (o *AnalyticsSku) SetLastUpdated(v time.Time)`

SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.

### GetUnitsSold

`func (o *AnalyticsSku) GetUnitsSold() AnalyticsDataPointWithTrend`

GetUnitsSold returns the UnitsSold field if non-nil, zero value otherwise.

### GetUnitsSoldOk

`func (o *AnalyticsSku) GetUnitsSoldOk() (AnalyticsDataPointWithTrend, bool)`

GetUnitsSoldOk returns a tuple with the UnitsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnitsSold

`func (o *AnalyticsSku) HasUnitsSold() bool`

HasUnitsSold returns a boolean if a field has been set.

### SetUnitsSold

`func (o *AnalyticsSku) SetUnitsSold(v AnalyticsDataPointWithTrend)`

SetUnitsSold gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the UnitsSold field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


