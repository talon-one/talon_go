# CartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | 
**Sku** | Pointer to **string** |  | 
**Quantity** | Pointer to **int32** |  | 
**Price** | Pointer to **float32** |  | 
**Category** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **float32** | Weight of item in mm | [optional] 
**Height** | Pointer to **float32** | Height of item in mm | [optional] 
**Width** | Pointer to **float32** | Width of item in mm | [optional] 
**Length** | Pointer to **float32** | Length of item in mm | [optional] 
**Position** | Pointer to **float32** | Position of the Cart Item in the Cart (calculated internally) | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item | [optional] 

## Methods

### GetName

`func (o *CartItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartItem) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CartItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CartItem) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetSku

`func (o *CartItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CartItem) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *CartItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *CartItem) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.

### GetQuantity

`func (o *CartItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartItem) GetQuantityOk() (int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuantity

`func (o *CartItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantity

`func (o *CartItem) SetQuantity(v int32)`

SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.

### GetPrice

`func (o *CartItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CartItem) GetPriceOk() (float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *CartItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *CartItem) SetPrice(v float32)`

SetPrice gets a reference to the given float32 and assigns it to the Price field.

### GetCategory

`func (o *CartItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CartItem) GetCategoryOk() (string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategory

`func (o *CartItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategory

`func (o *CartItem) SetCategory(v string)`

SetCategory gets a reference to the given string and assigns it to the Category field.

### GetWeight

`func (o *CartItem) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CartItem) GetWeightOk() (float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWeight

`func (o *CartItem) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeight

`func (o *CartItem) SetWeight(v float32)`

SetWeight gets a reference to the given float32 and assigns it to the Weight field.

### GetHeight

`func (o *CartItem) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *CartItem) GetHeightOk() (float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeight

`func (o *CartItem) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeight

`func (o *CartItem) SetHeight(v float32)`

SetHeight gets a reference to the given float32 and assigns it to the Height field.

### GetWidth

`func (o *CartItem) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CartItem) GetWidthOk() (float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWidth

`func (o *CartItem) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidth

`func (o *CartItem) SetWidth(v float32)`

SetWidth gets a reference to the given float32 and assigns it to the Width field.

### GetLength

`func (o *CartItem) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *CartItem) GetLengthOk() (float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLength

`func (o *CartItem) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLength

`func (o *CartItem) SetLength(v float32)`

SetLength gets a reference to the given float32 and assigns it to the Length field.

### GetPosition

`func (o *CartItem) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CartItem) GetPositionOk() (float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosition

`func (o *CartItem) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPosition

`func (o *CartItem) SetPosition(v float32)`

SetPosition gets a reference to the given float32 and assigns it to the Position field.

### GetAttributes

`func (o *CartItem) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CartItem) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CartItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CartItem) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


