# AddFreeItemEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | SKU of the item that needs to be added. | 
**Name** | Pointer to **string** | The name / description of the effect | 
**DesiredQuantity** | Pointer to **int64** | The original quantity in case a partial reward was applied. | [optional] 

## Methods

### NewAddFreeItemEffectProps

`func NewAddFreeItemEffectProps(sku string, name string, ) *AddFreeItemEffectProps`

NewAddFreeItemEffectProps instantiates a new AddFreeItemEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFreeItemEffectPropsWithDefaults

`func NewAddFreeItemEffectPropsWithDefaults() *AddFreeItemEffectProps`

NewAddFreeItemEffectPropsWithDefaults instantiates a new AddFreeItemEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *AddFreeItemEffectProps) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AddFreeItemEffectProps) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *AddFreeItemEffectProps) SetSku(v string)`

SetSku sets Sku field to given value.


### GetName

`func (o *AddFreeItemEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddFreeItemEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddFreeItemEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetDesiredQuantity

`func (o *AddFreeItemEffectProps) GetDesiredQuantity() int64`

GetDesiredQuantity returns the DesiredQuantity field if non-nil, zero value otherwise.

### GetDesiredQuantityOk

`func (o *AddFreeItemEffectProps) GetDesiredQuantityOk() (*int64, bool)`

GetDesiredQuantityOk returns a tuple with the DesiredQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredQuantity

`func (o *AddFreeItemEffectProps) SetDesiredQuantity(v int64)`

SetDesiredQuantity sets DesiredQuantity field to given value.

### HasDesiredQuantity

`func (o *AddFreeItemEffectProps) HasDesiredQuantity() bool`

HasDesiredQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


