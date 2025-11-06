# RollbackDiscountEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the \&quot;setDiscount\&quot; effect that was rolled back. | 
**Value** | Pointer to **float32** | The value of the discount that was rolled back. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the cart items for which the discount was rolled back. | [optional] 
**CartItemSubPosition** | Pointer to **float32** | For cart items with &#x60;quantity&#x60; &gt; 1, the subposition returns the index of the item unit in its line item.  | [optional] 
**AdditionalCostId** | Pointer to **int64** | The ID of the additional cost that was rolled back. | [optional] 
**AdditionalCost** | Pointer to **string** | The name of the additional cost that was rolled back. | [optional] 
**Scope** | Pointer to **string** | The scope of the rolled back discount - For a discount per session, it can be one of &#x60;cartItems&#x60;, &#x60;additionalCosts&#x60; or &#x60;sessionTotal&#x60; - For a discount per item, it can be one of &#x60;price&#x60;, &#x60;additionalCosts&#x60; or &#x60;itemTotal&#x60;  | [optional] 

## Methods

### NewRollbackDiscountEffectProps

`func NewRollbackDiscountEffectProps(name string, value float32, ) *RollbackDiscountEffectProps`

NewRollbackDiscountEffectProps instantiates a new RollbackDiscountEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackDiscountEffectPropsWithDefaults

`func NewRollbackDiscountEffectPropsWithDefaults() *RollbackDiscountEffectProps`

NewRollbackDiscountEffectPropsWithDefaults instantiates a new RollbackDiscountEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RollbackDiscountEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RollbackDiscountEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RollbackDiscountEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *RollbackDiscountEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RollbackDiscountEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RollbackDiscountEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetCartItemPosition

`func (o *RollbackDiscountEffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *RollbackDiscountEffectProps) GetCartItemPositionOk() (*float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemPosition

`func (o *RollbackDiscountEffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition sets CartItemPosition field to given value.

### HasCartItemPosition

`func (o *RollbackDiscountEffectProps) HasCartItemPosition() bool`

HasCartItemPosition returns a boolean if a field has been set.

### GetCartItemSubPosition

`func (o *RollbackDiscountEffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *RollbackDiscountEffectProps) GetCartItemSubPositionOk() (*float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemSubPosition

`func (o *RollbackDiscountEffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition sets CartItemSubPosition field to given value.

### HasCartItemSubPosition

`func (o *RollbackDiscountEffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### GetAdditionalCostId

`func (o *RollbackDiscountEffectProps) GetAdditionalCostId() int64`

GetAdditionalCostId returns the AdditionalCostId field if non-nil, zero value otherwise.

### GetAdditionalCostIdOk

`func (o *RollbackDiscountEffectProps) GetAdditionalCostIdOk() (*int64, bool)`

GetAdditionalCostIdOk returns a tuple with the AdditionalCostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCostId

`func (o *RollbackDiscountEffectProps) SetAdditionalCostId(v int64)`

SetAdditionalCostId sets AdditionalCostId field to given value.

### HasAdditionalCostId

`func (o *RollbackDiscountEffectProps) HasAdditionalCostId() bool`

HasAdditionalCostId returns a boolean if a field has been set.

### GetAdditionalCost

`func (o *RollbackDiscountEffectProps) GetAdditionalCost() string`

GetAdditionalCost returns the AdditionalCost field if non-nil, zero value otherwise.

### GetAdditionalCostOk

`func (o *RollbackDiscountEffectProps) GetAdditionalCostOk() (*string, bool)`

GetAdditionalCostOk returns a tuple with the AdditionalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCost

`func (o *RollbackDiscountEffectProps) SetAdditionalCost(v string)`

SetAdditionalCost sets AdditionalCost field to given value.

### HasAdditionalCost

`func (o *RollbackDiscountEffectProps) HasAdditionalCost() bool`

HasAdditionalCost returns a boolean if a field has been set.

### GetScope

`func (o *RollbackDiscountEffectProps) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RollbackDiscountEffectProps) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RollbackDiscountEffectProps) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RollbackDiscountEffectProps) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


