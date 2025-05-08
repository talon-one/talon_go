# ProductSearchMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **int32** | The ID of the product. | [optional] 
**Value** | Pointer to **string** | The string matching the given value. Either a product name or SKU. | 
**ProductSkuId** | Pointer to **int32** | The ID of the SKU linked to a product. If empty, this is an product. | [optional] 

## Methods

### GetProductId

`func (o *ProductSearchMatch) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductSearchMatch) GetProductIdOk() (int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductId

`func (o *ProductSearchMatch) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductId

`func (o *ProductSearchMatch) SetProductId(v int32)`

SetProductId gets a reference to the given int32 and assigns it to the ProductId field.

### GetValue

`func (o *ProductSearchMatch) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductSearchMatch) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *ProductSearchMatch) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *ProductSearchMatch) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.

### GetProductSkuId

`func (o *ProductSearchMatch) GetProductSkuId() int32`

GetProductSkuId returns the ProductSkuId field if non-nil, zero value otherwise.

### GetProductSkuIdOk

`func (o *ProductSearchMatch) GetProductSkuIdOk() (int32, bool)`

GetProductSkuIdOk returns a tuple with the ProductSkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductSkuId

`func (o *ProductSearchMatch) HasProductSkuId() bool`

HasProductSkuId returns a boolean if a field has been set.

### SetProductSkuId

`func (o *ProductSearchMatch) SetProductSkuId(v int32)`

SetProductSkuId gets a reference to the given int32 and assigns it to the ProductSkuId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


