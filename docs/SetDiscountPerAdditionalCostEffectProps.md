# SetDiscountPerAdditionalCostEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name / description of this discount | 
**AdditionalCostId** | Pointer to **int64** | The ID of the additional cost. | 
**AdditionalCost** | Pointer to **string** | The name of the additional cost. | 
**Value** | Pointer to **float32** | The total monetary value of the discount. | 
**DesiredValue** | Pointer to **float32** | The original value of the discount. | [optional] 

## Methods

### NewSetDiscountPerAdditionalCostEffectProps

`func NewSetDiscountPerAdditionalCostEffectProps(name string, additionalCostId int64, additionalCost string, value float32, ) *SetDiscountPerAdditionalCostEffectProps`

NewSetDiscountPerAdditionalCostEffectProps instantiates a new SetDiscountPerAdditionalCostEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDiscountPerAdditionalCostEffectPropsWithDefaults

`func NewSetDiscountPerAdditionalCostEffectPropsWithDefaults() *SetDiscountPerAdditionalCostEffectProps`

NewSetDiscountPerAdditionalCostEffectPropsWithDefaults instantiates a new SetDiscountPerAdditionalCostEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SetDiscountPerAdditionalCostEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetDiscountPerAdditionalCostEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetDiscountPerAdditionalCostEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetAdditionalCostId

`func (o *SetDiscountPerAdditionalCostEffectProps) GetAdditionalCostId() int64`

GetAdditionalCostId returns the AdditionalCostId field if non-nil, zero value otherwise.

### GetAdditionalCostIdOk

`func (o *SetDiscountPerAdditionalCostEffectProps) GetAdditionalCostIdOk() (*int64, bool)`

GetAdditionalCostIdOk returns a tuple with the AdditionalCostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCostId

`func (o *SetDiscountPerAdditionalCostEffectProps) SetAdditionalCostId(v int64)`

SetAdditionalCostId sets AdditionalCostId field to given value.


### GetAdditionalCost

`func (o *SetDiscountPerAdditionalCostEffectProps) GetAdditionalCost() string`

GetAdditionalCost returns the AdditionalCost field if non-nil, zero value otherwise.

### GetAdditionalCostOk

`func (o *SetDiscountPerAdditionalCostEffectProps) GetAdditionalCostOk() (*string, bool)`

GetAdditionalCostOk returns a tuple with the AdditionalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCost

`func (o *SetDiscountPerAdditionalCostEffectProps) SetAdditionalCost(v string)`

SetAdditionalCost sets AdditionalCost field to given value.


### GetValue

`func (o *SetDiscountPerAdditionalCostEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SetDiscountPerAdditionalCostEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SetDiscountPerAdditionalCostEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetDesiredValue

`func (o *SetDiscountPerAdditionalCostEffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *SetDiscountPerAdditionalCostEffectProps) GetDesiredValueOk() (*float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredValue

`func (o *SetDiscountPerAdditionalCostEffectProps) SetDesiredValue(v float32)`

SetDesiredValue sets DesiredValue field to given value.

### HasDesiredValue

`func (o *SetDiscountPerAdditionalCostEffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


