# SetDiscountPerAdditionalCostPerItemEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name / description of this discount | 
**AdditionalCostId** | Pointer to **int32** | The ID of the additional cost. | 
**Value** | Pointer to **float32** | The total monetary value of the discount. | 
**Position** | Pointer to **float32** | The index of the item in the cart item list containing the additional cost to be discounted. | 
**SubPosition** | Pointer to **float32** | For cart items with &#x60;quantity&#x60; &gt; 1, the sub position indicates which item the discount applies to.  | [optional] 
**AdditionalCost** | Pointer to **string** | The name of the additional cost. | 
**DesiredValue** | Pointer to **float32** | Only with [partial discounts enabled](https://docs.talon.one/docs/product/campaigns/campaign-evaluation/#partial-discounts). Represents the monetary value of the discount to be applied to additional discount without considering budget limitations.  | [optional] 

## Methods

### GetName

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetAdditionalCostId

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCostId() int32`

GetAdditionalCostId returns the AdditionalCostId field if non-nil, zero value otherwise.

### GetAdditionalCostIdOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCostIdOk() (int32, bool)`

GetAdditionalCostIdOk returns a tuple with the AdditionalCostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCostId

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasAdditionalCostId() bool`

HasAdditionalCostId returns a boolean if a field has been set.

### SetAdditionalCostId

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetAdditionalCostId(v int32)`

SetAdditionalCostId gets a reference to the given int32 and assigns it to the AdditionalCostId field.

### GetValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetValueOk() (float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetValue(v float32)`

SetValue gets a reference to the given float32 and assigns it to the Value field.

### GetPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetPositionOk() (float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetPosition(v float32)`

SetPosition gets a reference to the given float32 and assigns it to the Position field.

### GetSubPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetSubPosition() float32`

GetSubPosition returns the SubPosition field if non-nil, zero value otherwise.

### GetSubPositionOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetSubPositionOk() (float32, bool)`

GetSubPositionOk returns a tuple with the SubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasSubPosition() bool`

HasSubPosition returns a boolean if a field has been set.

### SetSubPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetSubPosition(v float32)`

SetSubPosition gets a reference to the given float32 and assigns it to the SubPosition field.

### GetAdditionalCost

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCost() string`

GetAdditionalCost returns the AdditionalCost field if non-nil, zero value otherwise.

### GetAdditionalCostOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCostOk() (string, bool)`

GetAdditionalCostOk returns a tuple with the AdditionalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalCost

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasAdditionalCost() bool`

HasAdditionalCost returns a boolean if a field has been set.

### SetAdditionalCost

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetAdditionalCost(v string)`

SetAdditionalCost gets a reference to the given string and assigns it to the AdditionalCost field.

### GetDesiredValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetDesiredValueOk() (float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDesiredValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.

### SetDesiredValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetDesiredValue(v float32)`

SetDesiredValue gets a reference to the given float32 and assigns it to the DesiredValue field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


