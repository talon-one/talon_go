# ReturnedCartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to **int64** | The index of the cart item in the provided customer session&#39;s &#x60;cartItems&#x60; property. | [optional] 
**Quantity** | Pointer to **int64** | Number of cart items to return.  | [optional] 
**Sku** | Pointer to **string** | The SKU of the cart item in the provided customer session&#39;s &#x60;cartItems&#x60; property. | [optional] 

## Methods

### NewReturnedCartItem

`func NewReturnedCartItem() *ReturnedCartItem`

NewReturnedCartItem instantiates a new ReturnedCartItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnedCartItemWithDefaults

`func NewReturnedCartItemWithDefaults() *ReturnedCartItem`

NewReturnedCartItemWithDefaults instantiates a new ReturnedCartItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *ReturnedCartItem) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ReturnedCartItem) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ReturnedCartItem) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ReturnedCartItem) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetQuantity

`func (o *ReturnedCartItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnedCartItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnedCartItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReturnedCartItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSku

`func (o *ReturnedCartItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ReturnedCartItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ReturnedCartItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ReturnedCartItem) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


