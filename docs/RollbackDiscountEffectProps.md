# RollbackDiscountEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the \&quot;setDiscount\&quot; effect that was rolled back. | 
**Value** | Pointer to **float32** | The value of the discount that was rolled back. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the cart items for which the discount was rolled back. | [optional] 
**CartItemSubPosition** | Pointer to **float32** | For cart items with &#x60;quantity&#x60; &gt; 1, the subposition returns the index of the item unit in its line item.  | [optional] 
**AdditionalCostId** | Pointer to **int32** | The ID of the additional cost that was rolled back. | [optional] 
**AdditionalCost** | Pointer to **string** | The name of the additional cost that was rolled back. | [optional] 
**Scope** | Pointer to **string** | The scope of the rolled back discount - For a discount per session, it can be one of &#x60;cartItems&#x60;, &#x60;additionalCosts&#x60; or &#x60;sessionTotal&#x60; - For a discount per item, it can be one of &#x60;price&#x60;, &#x60;additionalCosts&#x60; or &#x60;itemTotal&#x60;  | [optional] 

## Methods

### GetName

`func (o *RollbackDiscountEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RollbackDiscountEffectProps) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *RollbackDiscountEffectProps) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *RollbackDiscountEffectProps) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetValue

`func (o *RollbackDiscountEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RollbackDiscountEffectProps) GetValueOk() (float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *RollbackDiscountEffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *RollbackDiscountEffectProps) SetValue(v float32)`

SetValue gets a reference to the given float32 and assigns it to the Value field.

### GetCartItemPosition

`func (o *RollbackDiscountEffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *RollbackDiscountEffectProps) GetCartItemPositionOk() (float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemPosition

`func (o *RollbackDiscountEffectProps) HasCartItemPosition() bool`

HasCartItemPosition returns a boolean if a field has been set.

### SetCartItemPosition

`func (o *RollbackDiscountEffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition gets a reference to the given float32 and assigns it to the CartItemPosition field.

### GetCartItemSubPosition

`func (o *RollbackDiscountEffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *RollbackDiscountEffectProps) GetCartItemSubPositionOk() (float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemSubPosition

`func (o *RollbackDiscountEffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### SetCartItemSubPosition

`func (o *RollbackDiscountEffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition gets a reference to the given float32 and assigns it to the CartItemSubPosition field.

### GetAdditionalCostId

`func (o *RollbackDiscountEffectProps) GetAdditionalCostId() int32`

GetAdditionalCostId returns the AdditionalCostId field if non-nil, zero value otherwise.

### GetAdditionalCostIdOk

`func (o *RollbackDiscountEffectProps) GetAdditionalCostIdOk() (int32, bool)`

GetAdditionalCostIdOk returns a tuple with the AdditionalCostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCostId

`func (o *RollbackDiscountEffectProps) HasAdditionalCostId() bool`

HasAdditionalCostId returns a boolean if a field has been set.

### SetAdditionalCostId

`func (o *RollbackDiscountEffectProps) SetAdditionalCostId(v int32)`

SetAdditionalCostId gets a reference to the given int32 and assigns it to the AdditionalCostId field.

### GetAdditionalCost

`func (o *RollbackDiscountEffectProps) GetAdditionalCost() string`

GetAdditionalCost returns the AdditionalCost field if non-nil, zero value otherwise.

### GetAdditionalCostOk

`func (o *RollbackDiscountEffectProps) GetAdditionalCostOk() (string, bool)`

GetAdditionalCostOk returns a tuple with the AdditionalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCost

`func (o *RollbackDiscountEffectProps) HasAdditionalCost() bool`

HasAdditionalCost returns a boolean if a field has been set.

### SetAdditionalCost

`func (o *RollbackDiscountEffectProps) SetAdditionalCost(v string)`

SetAdditionalCost gets a reference to the given string and assigns it to the AdditionalCost field.

### GetScope

`func (o *RollbackDiscountEffectProps) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RollbackDiscountEffectProps) GetScopeOk() (string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScope

`func (o *RollbackDiscountEffectProps) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScope

`func (o *RollbackDiscountEffectProps) SetScope(v string)`

SetScope gets a reference to the given string and assigns it to the Scope field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


