# StrikethroughChangedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogId** | Pointer to **int32** | The ID of the catalog that the changed item belongs to. | 
**Effects** | Pointer to [**[]StrikethroughEffect**](StrikethroughEffect.md) |  | [optional] 
**EvaluatedAt** | Pointer to [**time.Time**](time.Time.md) | The evaluation time of the changed item. | 
**Id** | Pointer to **int32** | The ID of the event that triggered the strikethrough labeling. | 
**Price** | Pointer to **float32** | The price of the changed item. | 
**Sku** | Pointer to **string** | The unique SKU of the changed item. | 
**Version** | Pointer to **int32** | The version of the changed item. | 

## Methods

### GetCatalogId

`func (o *StrikethroughChangedItem) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *StrikethroughChangedItem) GetCatalogIdOk() (int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCatalogId

`func (o *StrikethroughChangedItem) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### SetCatalogId

`func (o *StrikethroughChangedItem) SetCatalogId(v int32)`

SetCatalogId gets a reference to the given int32 and assigns it to the CatalogId field.

### GetEffects

`func (o *StrikethroughChangedItem) GetEffects() []StrikethroughEffect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *StrikethroughChangedItem) GetEffectsOk() ([]StrikethroughEffect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffects

`func (o *StrikethroughChangedItem) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffects

`func (o *StrikethroughChangedItem) SetEffects(v []StrikethroughEffect)`

SetEffects gets a reference to the given []StrikethroughEffect and assigns it to the Effects field.

### GetEvaluatedAt

`func (o *StrikethroughChangedItem) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *StrikethroughChangedItem) GetEvaluatedAtOk() (time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluatedAt

`func (o *StrikethroughChangedItem) HasEvaluatedAt() bool`

HasEvaluatedAt returns a boolean if a field has been set.

### SetEvaluatedAt

`func (o *StrikethroughChangedItem) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt gets a reference to the given time.Time and assigns it to the EvaluatedAt field.

### GetId

`func (o *StrikethroughChangedItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StrikethroughChangedItem) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *StrikethroughChangedItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *StrikethroughChangedItem) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetPrice

`func (o *StrikethroughChangedItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *StrikethroughChangedItem) GetPriceOk() (float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *StrikethroughChangedItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *StrikethroughChangedItem) SetPrice(v float32)`

SetPrice gets a reference to the given float32 and assigns it to the Price field.

### GetSku

`func (o *StrikethroughChangedItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StrikethroughChangedItem) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *StrikethroughChangedItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *StrikethroughChangedItem) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.

### GetVersion

`func (o *StrikethroughChangedItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StrikethroughChangedItem) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *StrikethroughChangedItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *StrikethroughChangedItem) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


