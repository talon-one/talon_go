# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A short description of the rule. | 
**Description** | Pointer to **string** | A longer, more detailed description of the rule. | [optional] 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | [optional] 
**Condition** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | A Talang expression that will be evaluated in the context of the given event. | 
**Effects** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | An array of effectful Talang expressions in arrays that will be evaluated when a rule matches. | 

## Methods

### GetTitle

`func (o *Rule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Rule) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *Rule) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *Rule) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *Rule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Rule) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Rule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Rule) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetBindings

`func (o *Rule) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *Rule) GetBindingsOk() ([]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBindings

`func (o *Rule) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### SetBindings

`func (o *Rule) SetBindings(v []Binding)`

SetBindings gets a reference to the given []Binding and assigns it to the Bindings field.

### GetCondition

`func (o *Rule) GetCondition() []map[string]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Rule) GetConditionOk() ([]map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCondition

`func (o *Rule) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetCondition

`func (o *Rule) SetCondition(v []map[string]interface{})`

SetCondition gets a reference to the given []map[string]interface{} and assigns it to the Condition field.

### GetEffects

`func (o *Rule) GetEffects() []map[string]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *Rule) GetEffectsOk() ([]map[string]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffects

`func (o *Rule) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffects

`func (o *Rule) SetEffects(v []map[string]interface{})`

SetEffects gets a reference to the given []map[string]interface{} and assigns it to the Effects field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


