# SetDiscountEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name / description of this discount | 
**Value** | Pointer to **float32** | The total monetary value of the discount. | 
**Scope** | Pointer to **string** | The scope which the discount was applied on, can be one of (cartItems,additionalCosts,sessionTotal). | [optional] 
**DesiredValue** | Pointer to **float32** | The original value of the discount. | [optional] 

## Methods

### NewSetDiscountEffectProps

`func NewSetDiscountEffectProps(name string, value float32, ) *SetDiscountEffectProps`

NewSetDiscountEffectProps instantiates a new SetDiscountEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDiscountEffectPropsWithDefaults

`func NewSetDiscountEffectPropsWithDefaults() *SetDiscountEffectProps`

NewSetDiscountEffectPropsWithDefaults instantiates a new SetDiscountEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SetDiscountEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SetDiscountEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SetDiscountEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *SetDiscountEffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SetDiscountEffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SetDiscountEffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetScope

`func (o *SetDiscountEffectProps) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SetDiscountEffectProps) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SetDiscountEffectProps) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SetDiscountEffectProps) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDesiredValue

`func (o *SetDiscountEffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *SetDiscountEffectProps) GetDesiredValueOk() (*float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredValue

`func (o *SetDiscountEffectProps) SetDesiredValue(v float32)`

SetDesiredValue sets DesiredValue field to given value.

### HasDesiredValue

`func (o *SetDiscountEffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


