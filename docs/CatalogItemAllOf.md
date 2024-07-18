# CatalogItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | The stock keeping unit of the item. | 
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Catalogid** | Pointer to **int32** | The ID of the catalog the item belongs to. | 
**Version** | Pointer to **int32** | The version of the catalog item. | 
**Attributes** | Pointer to [**[]ItemAttribute**](ItemAttribute.md) |  | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 

## Methods

### GetSku

`func (o *CatalogItemAllOf) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CatalogItemAllOf) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *CatalogItemAllOf) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *CatalogItemAllOf) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.

### GetPrice

`func (o *CatalogItemAllOf) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogItemAllOf) GetPriceOk() (float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *CatalogItemAllOf) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *CatalogItemAllOf) SetPrice(v float32)`

SetPrice gets a reference to the given float32 and assigns it to the Price field.

### GetCatalogid

`func (o *CatalogItemAllOf) GetCatalogid() int32`

GetCatalogid returns the Catalogid field if non-nil, zero value otherwise.

### GetCatalogidOk

`func (o *CatalogItemAllOf) GetCatalogidOk() (int32, bool)`

GetCatalogidOk returns a tuple with the Catalogid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCatalogid

`func (o *CatalogItemAllOf) HasCatalogid() bool`

HasCatalogid returns a boolean if a field has been set.

### SetCatalogid

`func (o *CatalogItemAllOf) SetCatalogid(v int32)`

SetCatalogid gets a reference to the given int32 and assigns it to the Catalogid field.

### GetVersion

`func (o *CatalogItemAllOf) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogItemAllOf) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *CatalogItemAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *CatalogItemAllOf) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAttributes

`func (o *CatalogItemAllOf) GetAttributes() []ItemAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogItemAllOf) GetAttributesOk() ([]ItemAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CatalogItemAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CatalogItemAllOf) SetAttributes(v []ItemAttribute)`

SetAttributes gets a reference to the given []ItemAttribute and assigns it to the Attributes field.

### GetProduct

`func (o *CatalogItemAllOf) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CatalogItemAllOf) GetProductOk() (Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProduct

`func (o *CatalogItemAllOf) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProduct

`func (o *CatalogItemAllOf) SetProduct(v Product)`

SetProduct gets a reference to the given Product and assigns it to the Product field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


