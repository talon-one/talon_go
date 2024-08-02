# AddItemCatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | The attributes of the item to add. | [optional] 
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**ReplaceIfExists** | Pointer to **bool** | Indicates whether to replace the attributes of the item if the same SKU exists.  **Note**: When set to &#x60;true&#x60;:   - If you do not provide a new &#x60;price&#x60; value, the existing &#x60;price&#x60; value is retained.   - If you do not provide a new &#x60;product&#x60; value, the &#x60;product&#x60; value is set to &#x60;null&#x60;.  | [optional] [default to false]
**Sku** | Pointer to **string** | The unique SKU of the item to add. | 

## Methods

### GetAttributes

`func (o *AddItemCatalogAction) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AddItemCatalogAction) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *AddItemCatalogAction) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *AddItemCatalogAction) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetPrice

`func (o *AddItemCatalogAction) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddItemCatalogAction) GetPriceOk() (float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *AddItemCatalogAction) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *AddItemCatalogAction) SetPrice(v float32)`

SetPrice gets a reference to the given float32 and assigns it to the Price field.

### GetProduct

`func (o *AddItemCatalogAction) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AddItemCatalogAction) GetProductOk() (Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProduct

`func (o *AddItemCatalogAction) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProduct

`func (o *AddItemCatalogAction) SetProduct(v Product)`

SetProduct gets a reference to the given Product and assigns it to the Product field.

### GetReplaceIfExists

`func (o *AddItemCatalogAction) GetReplaceIfExists() bool`

GetReplaceIfExists returns the ReplaceIfExists field if non-nil, zero value otherwise.

### GetReplaceIfExistsOk

`func (o *AddItemCatalogAction) GetReplaceIfExistsOk() (bool, bool)`

GetReplaceIfExistsOk returns a tuple with the ReplaceIfExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReplaceIfExists

`func (o *AddItemCatalogAction) HasReplaceIfExists() bool`

HasReplaceIfExists returns a boolean if a field has been set.

### SetReplaceIfExists

`func (o *AddItemCatalogAction) SetReplaceIfExists(v bool)`

SetReplaceIfExists gets a reference to the given bool and assigns it to the ReplaceIfExists field.

### GetSku

`func (o *AddItemCatalogAction) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AddItemCatalogAction) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *AddItemCatalogAction) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *AddItemCatalogAction) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


