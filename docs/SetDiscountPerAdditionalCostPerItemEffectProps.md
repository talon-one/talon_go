# SetDiscountPerAdditionalCostPerItemEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name / description of this discount | 
**AdditionalCostId** | Pointer to **int64** | The ID of the additional cost. | 
**Value** | Pointer to **float32** | The total monetary value of the discount. | 
**Position** | Pointer to **float32** | The index of the item in the cart item list containing the additional cost to be discounted. | 
**SubPosition** | Pointer to **float32** | For cart items with &#x60;quantity&#x60; &gt; 1, the sub position indicates which item the discount applies to.  | [optional] 
**AdditionalCost** | Pointer to **string** | The name of the additional cost. | 
**DesiredValue** | Pointer to **float32** | Only with [partial discounts enabled](https://docs.talon.one/docs/product/campaigns/campaign-evaluation/#partial-discounts). Represents the monetary value of the discount to be applied to additional discount without considering budget limitations.  | [optional] 

## Methods

### NewSetDiscountPerAdditionalCostPerItemEffectProps

`func NewSetDiscountPerAdditionalCostPerItemEffectProps(name string, additionalCostId int64, value float32, position float32, additionalCost string, ) *SetDiscountPerAdditionalCostPerItemEffectProps`

NewSetDiscountPerAdditionalCostPerItemEffectProps instantiates a new SetDiscountPerAdditionalCostPerItemEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDiscountPerAdditionalCostPerItemEffectPropsWithDefaults

`func NewSetDiscountPerAdditionalCostPerItemEffectPropsWithDefaults() *SetDiscountPerAdditionalCostPerItemEffectProps`

NewSetDiscountPerAdditionalCostPerItemEffectPropsWithDefaults instantiates a new SetDiscountPerAdditionalCostPerItemEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetAdditionalCostId

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCostId() int64`

GetAdditionalCostId returns the AdditionalCostId field if non-nil, zero value otherwise.

### GetAdditionalCostIdOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCostIdOk() (*int64, bool)`

GetAdditionalCostIdOk returns a tuple with the AdditionalCostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCostId

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetAdditionalCostId(v int64)`

SetAdditionalCostId sets AdditionalCostId field to given value.


### GetValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetPosition(v float32)`

SetPosition sets Position field to given value.


### GetSubPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetSubPosition() float32`

GetSubPosition returns the SubPosition field if non-nil, zero value otherwise.

### GetSubPositionOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetSubPositionOk() (*float32, bool)`

GetSubPositionOk returns a tuple with the SubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetSubPosition(v float32)`

SetSubPosition sets SubPosition field to given value.

### HasSubPosition

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasSubPosition() bool`

HasSubPosition returns a boolean if a field has been set.

### GetAdditionalCost

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCost() string`

GetAdditionalCost returns the AdditionalCost field if non-nil, zero value otherwise.

### GetAdditionalCostOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCostOk() (*string, bool)`

GetAdditionalCostOk returns a tuple with the AdditionalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCost

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetAdditionalCost(v string)`

SetAdditionalCost sets AdditionalCost field to given value.


### GetDesiredValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetDesiredValueOk() (*float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetDesiredValue(v float32)`

SetDesiredValue sets DesiredValue field to given value.

### HasDesiredValue

`func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


