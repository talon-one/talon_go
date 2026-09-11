# Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the value to be bound. | 
**Type** | Pointer to **string** | The kind of binding. Possible values are: - &#x60;bundle&#x60; - &#x60;cartItemFilter&#x60; - &#x60;subledgerBalance&#x60; - &#x60;templateParameter&#x60;  | [optional] 
**Expression** | Pointer to **[]map[string]interface{}** | A Talang expression that is evaluated, and its result is bound to the name of the binding. The first element must be one of the functions or operators supported by Talang, followed by its arguments. The arguments can be strings, numbers, or nested expressions. For example: - &#x60;[\&quot;list\&quot;, \&quot;10014\&quot;, \&quot;10015\&quot;]&#x60; calls the &#x60;list&#x60; function to build a list of strings. - &#x60;[\&quot;+\&quot;, 2, 0]&#x60; uses the &#x60;+&#x60; operator to add two numbers.  | 
**ValueType** | Pointer to **string** | The data type of the value. One of the following: - &#x60;string&#x60; - &#x60;number&#x60; - &#x60;boolean&#x60;  | [optional] 
**MinValue** | Pointer to **float32** | The minimum value allowed for this placeholder. | [optional] 
**MaxValue** | Pointer to **float32** | The maximum value allowed for this placeholder. | [optional] 
**AttributeId** | Pointer to **int64** | Identifier of the attribute attached to the placeholder. | [optional] 
**Description** | Pointer to **string** | Description of the placeholder field and its value in the template. This text can be shown when creating campaigns from this template. | [optional] 

## Methods

### NewBinding

`func NewBinding(name string, expression []map[string]interface{}, ) *Binding`

NewBinding instantiates a new Binding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindingWithDefaults

`func NewBindingWithDefaults() *Binding`

NewBindingWithDefaults instantiates a new Binding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Binding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Binding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Binding) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Binding) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Binding) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Binding) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Binding) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExpression

`func (o *Binding) GetExpression() []interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Binding) GetExpressionOk() (*[]map[string]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Binding) SetExpression(v []interface{})`

SetExpression sets Expression field to given value.


### GetValueType

`func (o *Binding) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *Binding) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *Binding) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *Binding) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetMinValue

`func (o *Binding) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *Binding) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *Binding) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *Binding) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *Binding) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *Binding) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *Binding) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *Binding) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetAttributeId

`func (o *Binding) GetAttributeId() int64`

GetAttributeId returns the AttributeId field if non-nil, zero value otherwise.

### GetAttributeIdOk

`func (o *Binding) GetAttributeIdOk() (*int64, bool)`

GetAttributeIdOk returns a tuple with the AttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeId

`func (o *Binding) SetAttributeId(v int64)`

SetAttributeId sets AttributeId field to given value.

### HasAttributeId

`func (o *Binding) HasAttributeId() bool`

HasAttributeId returns a boolean if a field has been set.

### GetDescription

`func (o *Binding) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Binding) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Binding) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Binding) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


