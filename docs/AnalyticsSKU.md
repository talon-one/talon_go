# AnalyticsSKU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the SKU linked to the Application. | 
**Sku** | Pointer to **string** | The SKU linked to the Application. | 
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) | Values in UTC for the date the SKU linked to the product was last updated. | [optional] 
**CatalogId** | Pointer to **int64** | The ID of the catalog that contains the SKU. | [optional] 
**ProductId** | Pointer to **int64** | The ID of the product that the SKU belongs to. | [optional] 
**UnitsSold** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | [optional] 

## Methods

### NewAnalyticsSKU

`func NewAnalyticsSKU(id int64, sku string, ) *AnalyticsSKU`

NewAnalyticsSKU instantiates a new AnalyticsSKU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsSKUWithDefaults

`func NewAnalyticsSKUWithDefaults() *AnalyticsSKU`

NewAnalyticsSKUWithDefaults instantiates a new AnalyticsSKU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsSKU) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsSKU) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsSKU) SetId(v int64)`

SetId sets Id field to given value.


### GetSku

`func (o *AnalyticsSKU) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AnalyticsSKU) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *AnalyticsSKU) SetSku(v string)`

SetSku sets Sku field to given value.


### GetLastUpdated

`func (o *AnalyticsSKU) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AnalyticsSKU) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AnalyticsSKU) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AnalyticsSKU) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCatalogId

`func (o *AnalyticsSKU) GetCatalogId() int64`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *AnalyticsSKU) GetCatalogIdOk() (*int64, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *AnalyticsSKU) SetCatalogId(v int64)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *AnalyticsSKU) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetProductId

`func (o *AnalyticsSKU) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AnalyticsSKU) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AnalyticsSKU) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *AnalyticsSKU) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUnitsSold

`func (o *AnalyticsSKU) GetUnitsSold() AnalyticsDataPointWithTrend`

GetUnitsSold returns the UnitsSold field if non-nil, zero value otherwise.

### GetUnitsSoldOk

`func (o *AnalyticsSKU) GetUnitsSoldOk() (*AnalyticsDataPointWithTrend, bool)`

GetUnitsSoldOk returns a tuple with the UnitsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsSold

`func (o *AnalyticsSKU) SetUnitsSold(v AnalyticsDataPointWithTrend)`

SetUnitsSold sets UnitsSold field to given value.

### HasUnitsSold

`func (o *AnalyticsSKU) HasUnitsSold() bool`

HasUnitsSold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


