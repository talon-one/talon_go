# StrikethroughChangedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the event that triggered the strikethrough labeling. | 
**CatalogId** | Pointer to **int64** | The ID of the catalog that the changed item belongs to. | 
**Sku** | Pointer to **string** | The unique SKU of the changed item. | 
**Version** | Pointer to **int64** | The version of the changed item. | 
**Price** | Pointer to **float32** | The price of the changed item. | 
**Prices** | Pointer to [**map[string]PriceDetail**](PriceDetail.md) | A map of keys and values representing the price types and related price adjustment details for this cart item.       The keys correspond to the &#x60;priceType&#x60; names.  | [optional] 
**EvaluatedAt** | Pointer to [**time.Time**](time.Time.md) | The evaluation time of the changed item. | 
**Effects** | Pointer to [**[]StrikethroughEffect**](StrikethroughEffect.md) |  | [optional] 

## Methods

### NewStrikethroughChangedItem

`func NewStrikethroughChangedItem(id int64, catalogId int64, sku string, version int64, price float32, evaluatedAt time.Time, ) *StrikethroughChangedItem`

NewStrikethroughChangedItem instantiates a new StrikethroughChangedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughChangedItemWithDefaults

`func NewStrikethroughChangedItemWithDefaults() *StrikethroughChangedItem`

NewStrikethroughChangedItemWithDefaults instantiates a new StrikethroughChangedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StrikethroughChangedItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StrikethroughChangedItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StrikethroughChangedItem) SetId(v int64)`

SetId sets Id field to given value.


### GetCatalogId

`func (o *StrikethroughChangedItem) GetCatalogId() int64`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *StrikethroughChangedItem) GetCatalogIdOk() (*int64, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *StrikethroughChangedItem) SetCatalogId(v int64)`

SetCatalogId sets CatalogId field to given value.


### GetSku

`func (o *StrikethroughChangedItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StrikethroughChangedItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StrikethroughChangedItem) SetSku(v string)`

SetSku sets Sku field to given value.


### GetVersion

`func (o *StrikethroughChangedItem) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StrikethroughChangedItem) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StrikethroughChangedItem) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetPrice

`func (o *StrikethroughChangedItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *StrikethroughChangedItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *StrikethroughChangedItem) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetPrices

`func (o *StrikethroughChangedItem) GetPrices() map[string]PriceDetail`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *StrikethroughChangedItem) GetPricesOk() (*map[string]PriceDetail, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *StrikethroughChangedItem) SetPrices(v map[string]PriceDetail)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *StrikethroughChangedItem) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetEvaluatedAt

`func (o *StrikethroughChangedItem) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *StrikethroughChangedItem) GetEvaluatedAtOk() (*time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAt

`func (o *StrikethroughChangedItem) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt sets EvaluatedAt field to given value.


### GetEffects

`func (o *StrikethroughChangedItem) GetEffects() []StrikethroughEffect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *StrikethroughChangedItem) GetEffectsOk() (*[]StrikethroughEffect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *StrikethroughChangedItem) SetEffects(v []StrikethroughEffect)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *StrikethroughChangedItem) HasEffects() bool`

HasEffects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


