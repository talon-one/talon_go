# CatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Sku** | Pointer to **string** | The stock keeping unit of the item. | 
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Catalogid** | Pointer to **int32** | The ID of the catalog the item belongs to. | 
**Version** | Pointer to **int32** | The version of the catalog item. | 
**Attributes** | Pointer to [**[]ItemAttribute**](ItemAttribute.md) |  | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 

## Methods

### GetId

`func (o *CatalogItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogItem) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CatalogItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CatalogItem) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CatalogItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CatalogItem) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CatalogItem) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CatalogItem) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetSku

`func (o *CatalogItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CatalogItem) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *CatalogItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *CatalogItem) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.

### GetPrice

`func (o *CatalogItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogItem) GetPriceOk() (float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *CatalogItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *CatalogItem) SetPrice(v float32)`

SetPrice gets a reference to the given float32 and assigns it to the Price field.

### GetCatalogid

`func (o *CatalogItem) GetCatalogid() int32`

GetCatalogid returns the Catalogid field if non-nil, zero value otherwise.

### GetCatalogidOk

`func (o *CatalogItem) GetCatalogidOk() (int32, bool)`

GetCatalogidOk returns a tuple with the Catalogid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCatalogid

`func (o *CatalogItem) HasCatalogid() bool`

HasCatalogid returns a boolean if a field has been set.

### SetCatalogid

`func (o *CatalogItem) SetCatalogid(v int32)`

SetCatalogid gets a reference to the given int32 and assigns it to the Catalogid field.

### GetVersion

`func (o *CatalogItem) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogItem) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *CatalogItem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *CatalogItem) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAttributes

`func (o *CatalogItem) GetAttributes() []ItemAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogItem) GetAttributesOk() ([]ItemAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CatalogItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CatalogItem) SetAttributes(v []ItemAttribute)`

SetAttributes gets a reference to the given []ItemAttribute and assigns it to the Attributes field.

### GetProduct

`func (o *CatalogItem) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CatalogItem) GetProductOk() (Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProduct

`func (o *CatalogItem) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProduct

`func (o *CatalogItem) SetProduct(v Product)`

SetProduct gets a reference to the given Product and assigns it to the Product field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


