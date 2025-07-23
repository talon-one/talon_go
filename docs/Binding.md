# Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the value to be bound. | 
**Type** | Pointer to **string** | The kind of binding. Possible values are: - &#x60;bundle&#x60; - &#x60;cartItemFilter&#x60; - &#x60;subledgerBalance&#x60; - &#x60;templateParameter&#x60;  | [optional] 
**Expression** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | A Talang expression that will be evaluated and its result attached to the name of the binding. | 
**ValueType** | Pointer to **string** | Can be one of the following: - &#x60;string&#x60; - &#x60;number&#x60; - &#x60;boolean&#x60;  | [optional] 
**MinValue** | Pointer to **float32** | The minimum value allowed for this placeholder. | [optional] 
**MaxValue** | Pointer to **float32** | The maximum value allowed for this placeholder. | [optional] 
**AttributeId** | Pointer to **int32** | Id of the attribute attached to the placeholder. | [optional] 
**Description** | Pointer to **string** | Describes the placeholder field and value in the template. This description can be used when creating campaigns from this template. | [optional] 

## Methods

### GetName

`func (o *Binding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Binding) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Binding) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Binding) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetType

`func (o *Binding) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Binding) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Binding) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Binding) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetExpression

`func (o *Binding) GetExpression() []map[string]interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Binding) GetExpressionOk() ([]map[string]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpression

`func (o *Binding) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpression

`func (o *Binding) SetExpression(v []map[string]interface{})`

SetExpression gets a reference to the given []map[string]interface{} and assigns it to the Expression field.

### GetValueType

`func (o *Binding) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *Binding) GetValueTypeOk() (string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValueType

`func (o *Binding) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueType

`func (o *Binding) SetValueType(v string)`

SetValueType gets a reference to the given string and assigns it to the ValueType field.

### GetMinValue

`func (o *Binding) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *Binding) GetMinValueOk() (float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinValue

`func (o *Binding) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValue

`func (o *Binding) SetMinValue(v float32)`

SetMinValue gets a reference to the given float32 and assigns it to the MinValue field.

### GetMaxValue

`func (o *Binding) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *Binding) GetMaxValueOk() (float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxValue

`func (o *Binding) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValue

`func (o *Binding) SetMaxValue(v float32)`

SetMaxValue gets a reference to the given float32 and assigns it to the MaxValue field.

### GetAttributeId

`func (o *Binding) GetAttributeId() int32`

GetAttributeId returns the AttributeId field if non-nil, zero value otherwise.

### GetAttributeIdOk

`func (o *Binding) GetAttributeIdOk() (int32, bool)`

GetAttributeIdOk returns a tuple with the AttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributeId

`func (o *Binding) HasAttributeId() bool`

HasAttributeId returns a boolean if a field has been set.

### SetAttributeId

`func (o *Binding) SetAttributeId(v int32)`

SetAttributeId gets a reference to the given int32 and assigns it to the AttributeId field.

### GetDescription

`func (o *Binding) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Binding) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Binding) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Binding) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


