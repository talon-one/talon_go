# TemplateParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the template parameter. | 
**Value** | Pointer to [**map[string]interface{}**](.md) | The parameter&#39;s bound value. Its type depends on the &#x60;valueType&#x60;. | 
**ValueType** | Pointer to **string** | The data type of the value, derived from the bound expression (for example &#x60;number&#x60;, &#x60;string&#x60;, &#x60;boolean&#x60;, &#x60;percent&#x60;, &#x60;time&#x60;, &#x60;(list string)&#x60;, or &#x60;(list number)&#x60;). | 
**MinValue** | Pointer to **float32** | The minimum value allowed for this parameter. | [optional] 
**MaxValue** | Pointer to **float32** | The maximum value allowed for this parameter. | [optional] 
**Description** | Pointer to **string** | A human-readable description of the parameter shown when creating campaigns from the template. | 
**Attribute** | Pointer to **int64** | The ID of the attribute linked to this parameter. Omitted when the parameter is not linked to an attribute. | [optional] 

## Methods

### NewTemplateParameter

`func NewTemplateParameter(name string, value map[string]interface{}, valueType string, description string, ) *TemplateParameter`

NewTemplateParameter instantiates a new TemplateParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateParameterWithDefaults

`func NewTemplateParameterWithDefaults() *TemplateParameter`

NewTemplateParameterWithDefaults instantiates a new TemplateParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateParameter) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *TemplateParameter) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TemplateParameter) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TemplateParameter) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetValueType

`func (o *TemplateParameter) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *TemplateParameter) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *TemplateParameter) SetValueType(v string)`

SetValueType sets ValueType field to given value.


### GetMinValue

`func (o *TemplateParameter) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *TemplateParameter) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *TemplateParameter) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *TemplateParameter) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *TemplateParameter) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *TemplateParameter) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *TemplateParameter) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *TemplateParameter) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateParameter) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAttribute

`func (o *TemplateParameter) GetAttribute() int64`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *TemplateParameter) GetAttributeOk() (*int64, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *TemplateParameter) SetAttribute(v int64)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *TemplateParameter) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


