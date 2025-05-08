# AnalyticsProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the product. | 
**Name** | Pointer to **string** | The name of the product. | 
**CatalogId** | Pointer to **int32** | The ID of the catalog. You can find the ID in the Campaign Manager in **Account** &gt; **Tools** &gt; **Cart item catalogs**.  | 
**UnitsSold** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | [optional] 

## Methods

### GetId

`func (o *AnalyticsProduct) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsProduct) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *AnalyticsProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *AnalyticsProduct) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetName

`func (o *AnalyticsProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalyticsProduct) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *AnalyticsProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *AnalyticsProduct) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetCatalogId

`func (o *AnalyticsProduct) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *AnalyticsProduct) GetCatalogIdOk() (int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCatalogId

`func (o *AnalyticsProduct) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### SetCatalogId

`func (o *AnalyticsProduct) SetCatalogId(v int32)`

SetCatalogId gets a reference to the given int32 and assigns it to the CatalogId field.

### GetUnitsSold

`func (o *AnalyticsProduct) GetUnitsSold() AnalyticsDataPointWithTrend`

GetUnitsSold returns the UnitsSold field if non-nil, zero value otherwise.

### GetUnitsSoldOk

`func (o *AnalyticsProduct) GetUnitsSoldOk() (AnalyticsDataPointWithTrend, bool)`

GetUnitsSoldOk returns a tuple with the UnitsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnitsSold

`func (o *AnalyticsProduct) HasUnitsSold() bool`

HasUnitsSold returns a boolean if a field has been set.

### SetUnitsSold

`func (o *AnalyticsProduct) SetUnitsSold(v AnalyticsDataPointWithTrend)`

SetUnitsSold gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the UnitsSold field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


