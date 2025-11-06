# CustomEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectId** | Pointer to **int64** | The ID of the custom effect that was triggered. | 
**Name** | Pointer to **string** | The type of the custom effect. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the cart item list to which the custom effect is applied. | [optional] 
**CartItemSubPosition** | Pointer to **float32** | For cart items with quantity &gt; 1, the sub position indicates to which item unit the custom effect is applied.  | [optional] 
**BundleIndex** | Pointer to **int64** | The position of the bundle in a list of item bundles created from the same bundle definition. | [optional] 
**BundleName** | Pointer to **string** | The name of the bundle definition. | [optional] 
**Payload** | Pointer to [**map[string]interface{}**](.md) | The JSON payload of the custom effect. | 

## Methods

### NewCustomEffectProps

`func NewCustomEffectProps(effectId int64, name string, payload map[string]interface{}, ) *CustomEffectProps`

NewCustomEffectProps instantiates a new CustomEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEffectPropsWithDefaults

`func NewCustomEffectPropsWithDefaults() *CustomEffectProps`

NewCustomEffectPropsWithDefaults instantiates a new CustomEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectId

`func (o *CustomEffectProps) GetEffectId() int64`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *CustomEffectProps) GetEffectIdOk() (*int64, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectId

`func (o *CustomEffectProps) SetEffectId(v int64)`

SetEffectId sets EffectId field to given value.


### GetName

`func (o *CustomEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetCartItemPosition

`func (o *CustomEffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *CustomEffectProps) GetCartItemPositionOk() (*float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemPosition

`func (o *CustomEffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition sets CartItemPosition field to given value.

### HasCartItemPosition

`func (o *CustomEffectProps) HasCartItemPosition() bool`

HasCartItemPosition returns a boolean if a field has been set.

### GetCartItemSubPosition

`func (o *CustomEffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *CustomEffectProps) GetCartItemSubPositionOk() (*float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemSubPosition

`func (o *CustomEffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition sets CartItemSubPosition field to given value.

### HasCartItemSubPosition

`func (o *CustomEffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### GetBundleIndex

`func (o *CustomEffectProps) GetBundleIndex() int64`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *CustomEffectProps) GetBundleIndexOk() (*int64, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleIndex

`func (o *CustomEffectProps) SetBundleIndex(v int64)`

SetBundleIndex sets BundleIndex field to given value.

### HasBundleIndex

`func (o *CustomEffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### GetBundleName

`func (o *CustomEffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *CustomEffectProps) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *CustomEffectProps) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *CustomEffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### GetPayload

`func (o *CustomEffectProps) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomEffectProps) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CustomEffectProps) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


