# SetDiscountPerItemEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the discount. Contains a hashtag character indicating the index of the position of the item the discount applies to. It is identical to the value of the &#x60;position&#x60; property.  | 
**Value** | Pointer to **float32** | The total monetary value of the discount. | 
**Position** | Pointer to **float32** | The index of the item in the cart items list on which this discount should be applied. | 
**SubPosition** | Pointer to **float32** | Only used when [cart item flattening](https://docs.talon.one/docs/product/campaigns/campaign-evaluation#flattening) is enabled. Indicates which item the discount applies to for cart items with &#x60;quantity&#x60; &gt; 1.  | [optional] 
**DesiredValue** | Pointer to **float32** | The original value of the discount. | [optional] 
**Scope** | Pointer to **string** | The scope of the discount: - &#x60;additionalCosts&#x60;: The discount applies to all the additional costs of the item. - &#x60;itemTotal&#x60;: The discount applies to the price of the item + the additional costs of the item. - &#x60;price&#x60;: The discount applies to the price of the item.  | [optional] 
**TotalDiscount** | Pointer to **float32** | The total discount given if this effect is a result of a prorated discount. | [optional] 
**DesiredTotalDiscount** | Pointer to **float32** | The original total discount to give if this effect is a result of a prorated discount. | [optional] 
**BundleIndex** | Pointer to **int32** | The position of the bundle in a list of item bundles created from the same bundle definition. | [optional] 
**BundleName** | Pointer to **string** | The name of the bundle definition. | [optional] 

## Methods

### GetName

`func (o *SetDiscountPerItemEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetDiscountPerItemEffectProps) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *SetDiscountPerItemEffectProps) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *SetDiscountPerItemEffectProps) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetValue

`func (o *SetDiscountPerItemEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SetDiscountPerItemEffectProps) GetValueOk() (float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *SetDiscountPerItemEffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *SetDiscountPerItemEffectProps) SetValue(v float32)`

SetValue gets a reference to the given float32 and assigns it to the Value field.

### GetPosition

`func (o *SetDiscountPerItemEffectProps) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SetDiscountPerItemEffectProps) GetPositionOk() (float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosition

`func (o *SetDiscountPerItemEffectProps) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPosition

`func (o *SetDiscountPerItemEffectProps) SetPosition(v float32)`

SetPosition gets a reference to the given float32 and assigns it to the Position field.

### GetSubPosition

`func (o *SetDiscountPerItemEffectProps) GetSubPosition() float32`

GetSubPosition returns the SubPosition field if non-nil, zero value otherwise.

### GetSubPositionOk

`func (o *SetDiscountPerItemEffectProps) GetSubPositionOk() (float32, bool)`

GetSubPositionOk returns a tuple with the SubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubPosition

`func (o *SetDiscountPerItemEffectProps) HasSubPosition() bool`

HasSubPosition returns a boolean if a field has been set.

### SetSubPosition

`func (o *SetDiscountPerItemEffectProps) SetSubPosition(v float32)`

SetSubPosition gets a reference to the given float32 and assigns it to the SubPosition field.

### GetDesiredValue

`func (o *SetDiscountPerItemEffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *SetDiscountPerItemEffectProps) GetDesiredValueOk() (float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDesiredValue

`func (o *SetDiscountPerItemEffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.

### SetDesiredValue

`func (o *SetDiscountPerItemEffectProps) SetDesiredValue(v float32)`

SetDesiredValue gets a reference to the given float32 and assigns it to the DesiredValue field.

### GetScope

`func (o *SetDiscountPerItemEffectProps) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SetDiscountPerItemEffectProps) GetScopeOk() (string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScope

`func (o *SetDiscountPerItemEffectProps) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScope

`func (o *SetDiscountPerItemEffectProps) SetScope(v string)`

SetScope gets a reference to the given string and assigns it to the Scope field.

### GetTotalDiscount

`func (o *SetDiscountPerItemEffectProps) GetTotalDiscount() float32`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *SetDiscountPerItemEffectProps) GetTotalDiscountOk() (float32, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalDiscount

`func (o *SetDiscountPerItemEffectProps) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### SetTotalDiscount

`func (o *SetDiscountPerItemEffectProps) SetTotalDiscount(v float32)`

SetTotalDiscount gets a reference to the given float32 and assigns it to the TotalDiscount field.

### GetDesiredTotalDiscount

`func (o *SetDiscountPerItemEffectProps) GetDesiredTotalDiscount() float32`

GetDesiredTotalDiscount returns the DesiredTotalDiscount field if non-nil, zero value otherwise.

### GetDesiredTotalDiscountOk

`func (o *SetDiscountPerItemEffectProps) GetDesiredTotalDiscountOk() (float32, bool)`

GetDesiredTotalDiscountOk returns a tuple with the DesiredTotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDesiredTotalDiscount

`func (o *SetDiscountPerItemEffectProps) HasDesiredTotalDiscount() bool`

HasDesiredTotalDiscount returns a boolean if a field has been set.

### SetDesiredTotalDiscount

`func (o *SetDiscountPerItemEffectProps) SetDesiredTotalDiscount(v float32)`

SetDesiredTotalDiscount gets a reference to the given float32 and assigns it to the DesiredTotalDiscount field.

### GetBundleIndex

`func (o *SetDiscountPerItemEffectProps) GetBundleIndex() int32`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *SetDiscountPerItemEffectProps) GetBundleIndexOk() (int32, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleIndex

`func (o *SetDiscountPerItemEffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### SetBundleIndex

`func (o *SetDiscountPerItemEffectProps) SetBundleIndex(v int32)`

SetBundleIndex gets a reference to the given int32 and assigns it to the BundleIndex field.

### GetBundleName

`func (o *SetDiscountPerItemEffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *SetDiscountPerItemEffectProps) GetBundleNameOk() (string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleName

`func (o *SetDiscountPerItemEffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### SetBundleName

`func (o *SetDiscountPerItemEffectProps) SetBundleName(v string)`

SetBundleName gets a reference to the given string and assigns it to the BundleName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


