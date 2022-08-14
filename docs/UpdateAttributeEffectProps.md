# UpdateAttributeEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The exact path of the attribute that was updated. | 
**Value** | Pointer to [**map[string]interface{}**](.md) | The new value of this attribute. The value can be of the following types: - boolean - location - number - string - time - list of any of those types  | 

## Methods

### GetPath

`func (o *UpdateAttributeEffectProps) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateAttributeEffectProps) GetPathOk() (string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPath

`func (o *UpdateAttributeEffectProps) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPath

`func (o *UpdateAttributeEffectProps) SetPath(v string)`

SetPath gets a reference to the given string and assigns it to the Path field.

### GetValue

`func (o *UpdateAttributeEffectProps) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateAttributeEffectProps) GetValueOk() (map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *UpdateAttributeEffectProps) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *UpdateAttributeEffectProps) SetValue(v map[string]interface{})`

SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


