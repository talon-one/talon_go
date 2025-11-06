# PatchItemCatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | The unique SKU of the item to patch. | 
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | The attributes of the item to patch. | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**CreateIfNotExists** | Pointer to **bool** | Indicates whether to create an item if the SKU does not exist. | [optional] [default to false]

## Methods

### NewPatchItemCatalogAction

`func NewPatchItemCatalogAction(sku string, ) *PatchItemCatalogAction`

NewPatchItemCatalogAction instantiates a new PatchItemCatalogAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchItemCatalogActionWithDefaults

`func NewPatchItemCatalogActionWithDefaults() *PatchItemCatalogAction`

NewPatchItemCatalogActionWithDefaults instantiates a new PatchItemCatalogAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *PatchItemCatalogAction) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PatchItemCatalogAction) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PatchItemCatalogAction) SetSku(v string)`

SetSku sets Sku field to given value.


### GetPrice

`func (o *PatchItemCatalogAction) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PatchItemCatalogAction) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PatchItemCatalogAction) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PatchItemCatalogAction) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAttributes

`func (o *PatchItemCatalogAction) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchItemCatalogAction) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchItemCatalogAction) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchItemCatalogAction) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetProduct

`func (o *PatchItemCatalogAction) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PatchItemCatalogAction) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PatchItemCatalogAction) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PatchItemCatalogAction) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCreateIfNotExists

`func (o *PatchItemCatalogAction) GetCreateIfNotExists() bool`

GetCreateIfNotExists returns the CreateIfNotExists field if non-nil, zero value otherwise.

### GetCreateIfNotExistsOk

`func (o *PatchItemCatalogAction) GetCreateIfNotExistsOk() (*bool, bool)`

GetCreateIfNotExistsOk returns a tuple with the CreateIfNotExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfNotExists

`func (o *PatchItemCatalogAction) SetCreateIfNotExists(v bool)`

SetCreateIfNotExists sets CreateIfNotExists field to given value.

### HasCreateIfNotExists

`func (o *PatchItemCatalogAction) HasCreateIfNotExists() bool`

HasCreateIfNotExists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


