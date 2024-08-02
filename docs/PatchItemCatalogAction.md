# PatchItemCatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | The attributes of the item to patch. | [optional] 
**CreateIfNotExists** | Pointer to **bool** | Indicates whether to create an item if the SKU does not exist. | [optional] [default to false]
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**Sku** | Pointer to **string** | The unique SKU of the item to patch. | 

## Methods

### GetAttributes

`func (o *PatchItemCatalogAction) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchItemCatalogAction) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *PatchItemCatalogAction) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *PatchItemCatalogAction) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetCreateIfNotExists

`func (o *PatchItemCatalogAction) GetCreateIfNotExists() bool`

GetCreateIfNotExists returns the CreateIfNotExists field if non-nil, zero value otherwise.

### GetCreateIfNotExistsOk

`func (o *PatchItemCatalogAction) GetCreateIfNotExistsOk() (bool, bool)`

GetCreateIfNotExistsOk returns a tuple with the CreateIfNotExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreateIfNotExists

`func (o *PatchItemCatalogAction) HasCreateIfNotExists() bool`

HasCreateIfNotExists returns a boolean if a field has been set.

### SetCreateIfNotExists

`func (o *PatchItemCatalogAction) SetCreateIfNotExists(v bool)`

SetCreateIfNotExists gets a reference to the given bool and assigns it to the CreateIfNotExists field.

### GetPrice

`func (o *PatchItemCatalogAction) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PatchItemCatalogAction) GetPriceOk() (float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *PatchItemCatalogAction) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *PatchItemCatalogAction) SetPrice(v float32)`

SetPrice gets a reference to the given float32 and assigns it to the Price field.

### GetProduct

`func (o *PatchItemCatalogAction) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PatchItemCatalogAction) GetProductOk() (Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProduct

`func (o *PatchItemCatalogAction) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProduct

`func (o *PatchItemCatalogAction) SetProduct(v Product)`

SetProduct gets a reference to the given Product and assigns it to the Product field.

### GetSku

`func (o *PatchItemCatalogAction) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PatchItemCatalogAction) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *PatchItemCatalogAction) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *PatchItemCatalogAction) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


