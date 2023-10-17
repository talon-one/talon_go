# CustomEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectId** | Pointer to **int32** | The ID of the custom effect that was triggered. | 
**Name** | Pointer to **string** | The type of the custom effect. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the cart item list to which the custom effect is applied. | [optional] 
**CartItemSubPosition** | Pointer to **float32** | For cart items with quantity &gt; 1, the sub position indicates to which item unit the custom effect is applied.  | [optional] 
**BundleIndex** | Pointer to **int32** | The position of the bundle in a list of item bundles created from the same bundle definition. | [optional] 
**BundleName** | Pointer to **string** | The name of the bundle definition. | [optional] 
**Payload** | Pointer to [**map[string]interface{}**](.md) | The JSON payload of the custom effect. | 

## Methods

### GetEffectId

`func (o *CustomEffectProps) GetEffectId() int32`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *CustomEffectProps) GetEffectIdOk() (int32, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffectId

`func (o *CustomEffectProps) HasEffectId() bool`

HasEffectId returns a boolean if a field has been set.

### SetEffectId

`func (o *CustomEffectProps) SetEffectId(v int32)`

SetEffectId gets a reference to the given int32 and assigns it to the EffectId field.

### GetName

`func (o *CustomEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEffectProps) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CustomEffectProps) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CustomEffectProps) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetCartItemPosition

`func (o *CustomEffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *CustomEffectProps) GetCartItemPositionOk() (float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemPosition

`func (o *CustomEffectProps) HasCartItemPosition() bool`

HasCartItemPosition returns a boolean if a field has been set.

### SetCartItemPosition

`func (o *CustomEffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition gets a reference to the given float32 and assigns it to the CartItemPosition field.

### GetCartItemSubPosition

`func (o *CustomEffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *CustomEffectProps) GetCartItemSubPositionOk() (float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCartItemSubPosition

`func (o *CustomEffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### SetCartItemSubPosition

`func (o *CustomEffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition gets a reference to the given float32 and assigns it to the CartItemSubPosition field.

### GetBundleIndex

`func (o *CustomEffectProps) GetBundleIndex() int32`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *CustomEffectProps) GetBundleIndexOk() (int32, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleIndex

`func (o *CustomEffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### SetBundleIndex

`func (o *CustomEffectProps) SetBundleIndex(v int32)`

SetBundleIndex gets a reference to the given int32 and assigns it to the BundleIndex field.

### GetBundleName

`func (o *CustomEffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *CustomEffectProps) GetBundleNameOk() (string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleName

`func (o *CustomEffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### SetBundleName

`func (o *CustomEffectProps) SetBundleName(v string)`

SetBundleName gets a reference to the given string and assigns it to the BundleName field.

### GetPayload

`func (o *CustomEffectProps) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomEffectProps) GetPayloadOk() (map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *CustomEffectProps) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *CustomEffectProps) SetPayload(v map[string]interface{})`

SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


