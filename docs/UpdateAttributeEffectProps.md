# UpdateAttributeEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | The exact path of the attribute that was updated. | 
**Value** | Pointer to [**map[string]interface{}**](.md) | The new value of this attribute. The value can be of the following types: - boolean - location - number - string - time - list of any of those types  | 

## Methods

### NewUpdateAttributeEffectProps

`func NewUpdateAttributeEffectProps(path string, value map[string]interface{}, ) *UpdateAttributeEffectProps`

NewUpdateAttributeEffectProps instantiates a new UpdateAttributeEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAttributeEffectPropsWithDefaults

`func NewUpdateAttributeEffectPropsWithDefaults() *UpdateAttributeEffectProps`

NewUpdateAttributeEffectPropsWithDefaults instantiates a new UpdateAttributeEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *UpdateAttributeEffectProps) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateAttributeEffectProps) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateAttributeEffectProps) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *UpdateAttributeEffectProps) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateAttributeEffectProps) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateAttributeEffectProps) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


