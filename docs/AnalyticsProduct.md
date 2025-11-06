# AnalyticsProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the product. | 
**Name** | Pointer to **string** | The name of the product. | 
**CatalogId** | Pointer to **int64** | The ID of the catalog. You can find the ID in the Campaign Manager in **Account** &gt; **Tools** &gt; **Cart item catalogs**.  | 
**UnitsSold** | Pointer to [**AnalyticsDataPointWithTrend**](AnalyticsDataPointWithTrend.md) |  | [optional] 

## Methods

### NewAnalyticsProduct

`func NewAnalyticsProduct(id int64, name string, catalogId int64, ) *AnalyticsProduct`

NewAnalyticsProduct instantiates a new AnalyticsProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsProductWithDefaults

`func NewAnalyticsProductWithDefaults() *AnalyticsProduct`

NewAnalyticsProductWithDefaults instantiates a new AnalyticsProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsProduct) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsProduct) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsProduct) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *AnalyticsProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalyticsProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalyticsProduct) SetName(v string)`

SetName sets Name field to given value.


### GetCatalogId

`func (o *AnalyticsProduct) GetCatalogId() int64`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *AnalyticsProduct) GetCatalogIdOk() (*int64, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *AnalyticsProduct) SetCatalogId(v int64)`

SetCatalogId sets CatalogId field to given value.


### GetUnitsSold

`func (o *AnalyticsProduct) GetUnitsSold() AnalyticsDataPointWithTrend`

GetUnitsSold returns the UnitsSold field if non-nil, zero value otherwise.

### GetUnitsSoldOk

`func (o *AnalyticsProduct) GetUnitsSoldOk() (*AnalyticsDataPointWithTrend, bool)`

GetUnitsSoldOk returns a tuple with the UnitsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsSold

`func (o *AnalyticsProduct) SetUnitsSold(v AnalyticsDataPointWithTrend)`

SetUnitsSold sets UnitsSold field to given value.

### HasUnitsSold

`func (o *AnalyticsProduct) HasUnitsSold() bool`

HasUnitsSold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


