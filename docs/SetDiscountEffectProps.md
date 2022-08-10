# SetDiscountEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name/description of this discount. | 
**Value** | Pointer to **float32** | The total monetary value of the discount. | 
**Scope** | Pointer to **string** | The scope which the discount was applied on, can be one of (cartItems,additionalCosts,sessionTotal). | [optional] 
**DesiredValue** | Pointer to **float32** | The original value of the discount. | [optional] 

## Methods

### GetName

`func (o *SetDiscountEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetDiscountEffectProps) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *SetDiscountEffectProps) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *SetDiscountEffectProps) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetValue

`func (o *SetDiscountEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SetDiscountEffectProps) GetValueOk() (float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *SetDiscountEffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *SetDiscountEffectProps) SetValue(v float32)`

SetValue gets a reference to the given float32 and assigns it to the Value field.

### GetScope

`func (o *SetDiscountEffectProps) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SetDiscountEffectProps) GetScopeOk() (string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScope

`func (o *SetDiscountEffectProps) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScope

`func (o *SetDiscountEffectProps) SetScope(v string)`

SetScope gets a reference to the given string and assigns it to the Scope field.

### GetDesiredValue

`func (o *SetDiscountEffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *SetDiscountEffectProps) GetDesiredValueOk() (float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDesiredValue

`func (o *SetDiscountEffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.

### SetDesiredValue

`func (o *SetDiscountEffectProps) SetDesiredValue(v float32)`

SetDesiredValue gets a reference to the given float32 and assigns it to the DesiredValue field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


