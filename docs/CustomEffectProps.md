# CustomEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectId** | Pointer to **int32** | The ID of the custom effect that was triggered | 
**Name** | Pointer to **string** | The type of the custom effect. | 
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


