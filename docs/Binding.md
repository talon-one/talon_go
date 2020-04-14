# Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the value to be bound. | 
**Expression** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | A Talang expression that will be evaluated and its result attached to the name of the binding. | 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


