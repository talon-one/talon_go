# CatalogActionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | The unique SKU of the item to remove. | 
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Attributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | The attributes of the items to patch. | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**ReplaceIfExists** | Pointer to **bool** | Indicates whether to replace the attributes of the item if the same SKU exists.  **Note**: When set to &#x60;true&#x60;:   - If you do not provide a new &#x60;price&#x60; value, the existing &#x60;price&#x60; value is retained.   - If you do not provide a new &#x60;product&#x60; value, the &#x60;product&#x60; value is set to &#x60;null&#x60;.  | [optional] [default to false]
**CreateIfNotExists** | Pointer to **bool** | Indicates whether to create an item if the SKU does not exist. | [optional] [default to false]
**Filters** | Pointer to [**[]CatalogActionFilter**](CatalogActionFilter.md) | The list of filters used to select the items to patch, joined by &#x60;AND&#x60;.  **Note:** Every item in the catalog will be removed if there are no filters.  | [optional] 

## Methods

### GetSku

`func (o *CatalogActionPayload) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CatalogActionPayload) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *CatalogActionPayload) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *CatalogActionPayload) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.

### GetPrice

`func (o *CatalogActionPayload) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogActionPayload) GetPriceOk() (float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *CatalogActionPayload) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *CatalogActionPayload) SetPrice(v float32)`

SetPrice gets a reference to the given float32 and assigns it to the Price field.

### GetAttributes

`func (o *CatalogActionPayload) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogActionPayload) GetAttributesOk() (map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CatalogActionPayload) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CatalogActionPayload) SetAttributes(v map[string]map[string]interface{})`

SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.

### GetProduct

`func (o *CatalogActionPayload) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CatalogActionPayload) GetProductOk() (Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProduct

`func (o *CatalogActionPayload) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProduct

`func (o *CatalogActionPayload) SetProduct(v Product)`

SetProduct gets a reference to the given Product and assigns it to the Product field.

### GetReplaceIfExists

`func (o *CatalogActionPayload) GetReplaceIfExists() bool`

GetReplaceIfExists returns the ReplaceIfExists field if non-nil, zero value otherwise.

### GetReplaceIfExistsOk

`func (o *CatalogActionPayload) GetReplaceIfExistsOk() (bool, bool)`

GetReplaceIfExistsOk returns a tuple with the ReplaceIfExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReplaceIfExists

`func (o *CatalogActionPayload) HasReplaceIfExists() bool`

HasReplaceIfExists returns a boolean if a field has been set.

### SetReplaceIfExists

`func (o *CatalogActionPayload) SetReplaceIfExists(v bool)`

SetReplaceIfExists gets a reference to the given bool and assigns it to the ReplaceIfExists field.

### GetCreateIfNotExists

`func (o *CatalogActionPayload) GetCreateIfNotExists() bool`

GetCreateIfNotExists returns the CreateIfNotExists field if non-nil, zero value otherwise.

### GetCreateIfNotExistsOk

`func (o *CatalogActionPayload) GetCreateIfNotExistsOk() (bool, bool)`

GetCreateIfNotExistsOk returns a tuple with the CreateIfNotExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreateIfNotExists

`func (o *CatalogActionPayload) HasCreateIfNotExists() bool`

HasCreateIfNotExists returns a boolean if a field has been set.

### SetCreateIfNotExists

`func (o *CatalogActionPayload) SetCreateIfNotExists(v bool)`

SetCreateIfNotExists gets a reference to the given bool and assigns it to the CreateIfNotExists field.

### GetFilters

`func (o *CatalogActionPayload) GetFilters() []CatalogActionFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CatalogActionPayload) GetFiltersOk() ([]CatalogActionFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilters

`func (o *CatalogActionPayload) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFilters

`func (o *CatalogActionPayload) SetFilters(v []CatalogActionFilter)`

SetFilters gets a reference to the given []CatalogActionFilter and assigns it to the Filters field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


