# AddItemCatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | The unique SKU of the item to add. | 
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | The attributes of the item to add. | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**ReplaceIfExists** | Pointer to **bool** | Indicates whether to replace the attributes of the item if the same SKU exists.  **Note**: When set to &#x60;true&#x60;:   - If you do not provide a new &#x60;price&#x60; value, the existing &#x60;price&#x60; value is retained.   - If you do not provide a new &#x60;product&#x60; value, the &#x60;product&#x60; value is set to &#x60;null&#x60;.  | [optional] [default to false]

## Methods

### NewAddItemCatalogAction

`func NewAddItemCatalogAction(sku string, ) *AddItemCatalogAction`

NewAddItemCatalogAction instantiates a new AddItemCatalogAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddItemCatalogActionWithDefaults

`func NewAddItemCatalogActionWithDefaults() *AddItemCatalogAction`

NewAddItemCatalogActionWithDefaults instantiates a new AddItemCatalogAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *AddItemCatalogAction) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AddItemCatalogAction) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *AddItemCatalogAction) SetSku(v string)`

SetSku sets Sku field to given value.


### GetPrice

`func (o *AddItemCatalogAction) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddItemCatalogAction) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AddItemCatalogAction) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AddItemCatalogAction) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAttributes

`func (o *AddItemCatalogAction) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AddItemCatalogAction) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AddItemCatalogAction) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AddItemCatalogAction) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetProduct

`func (o *AddItemCatalogAction) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AddItemCatalogAction) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AddItemCatalogAction) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AddItemCatalogAction) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetReplaceIfExists

`func (o *AddItemCatalogAction) GetReplaceIfExists() bool`

GetReplaceIfExists returns the ReplaceIfExists field if non-nil, zero value otherwise.

### GetReplaceIfExistsOk

`func (o *AddItemCatalogAction) GetReplaceIfExistsOk() (*bool, bool)`

GetReplaceIfExistsOk returns a tuple with the ReplaceIfExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceIfExists

`func (o *AddItemCatalogAction) SetReplaceIfExists(v bool)`

SetReplaceIfExists sets ReplaceIfExists field to given value.

### HasReplaceIfExists

`func (o *AddItemCatalogAction) HasReplaceIfExists() bool`

HasReplaceIfExists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


