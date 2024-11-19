# Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the value to be bound. | 
**Type** | Pointer to **string** | The kind of binding. Possible values are: - &#x60;bundle&#x60; - &#x60;cartItemFilter&#x60; - &#x60;subledgerBalance&#x60; - &#x60;templateParameter&#x60;  | [optional] 
**Expression** | Pointer to [**[]interface{}**]([]interface{}.md) | A Talang expression that will be evaluated and its result attached to the name of the binding. | 
**ValueType** | Pointer to **string** | Can be one of the following: - &#x60;string&#x60; - &#x60;number&#x60; - &#x60;boolean&#x60;  | [optional] 

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

`func (o *Binding) GetExpression() []interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Binding) GetExpressionOk() ([]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpression

`func (o *Binding) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpression

`func (o *Binding) SetExpression(v []interface{})`

SetExpression gets a reference to the given []interface{} and assigns it to the Expression field.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


