# ProductSearchMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **int64** | The ID of the product. | [optional] 
**Value** | Pointer to **string** | The string matching the given value. Either a product name or SKU. | 
**ProductSkuId** | Pointer to **int64** | The ID of the SKU linked to a product. If empty, this is an product. | [optional] 

## Methods

### NewProductSearchMatch

`func NewProductSearchMatch(value string, ) *ProductSearchMatch`

NewProductSearchMatch instantiates a new ProductSearchMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSearchMatchWithDefaults

`func NewProductSearchMatchWithDefaults() *ProductSearchMatch`

NewProductSearchMatchWithDefaults instantiates a new ProductSearchMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductSearchMatch) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductSearchMatch) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductSearchMatch) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductSearchMatch) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetValue

`func (o *ProductSearchMatch) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductSearchMatch) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductSearchMatch) SetValue(v string)`

SetValue sets Value field to given value.


### GetProductSkuId

`func (o *ProductSearchMatch) GetProductSkuId() int64`

GetProductSkuId returns the ProductSkuId field if non-nil, zero value otherwise.

### GetProductSkuIdOk

`func (o *ProductSearchMatch) GetProductSkuIdOk() (*int64, bool)`

GetProductSkuIdOk returns a tuple with the ProductSkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSkuId

`func (o *ProductSearchMatch) SetProductSkuId(v int64)`

SetProductSkuId sets ProductSkuId field to given value.

### HasProductSkuId

`func (o *ProductSearchMatch) HasProductSkuId() bool`

HasProductSkuId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


