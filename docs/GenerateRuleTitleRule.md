# GenerateRuleTitleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**[]map[string]interface{}**](map[string]interface{}.md) | A Talang expression that will be evaluated in the context of the given event. | [optional] 
**Effects** | Pointer to [**[][]map[string]interface{}**](array.md) | An array of effectful Talang expressions in arrays that will be evaluated when a rule matches. | [optional] 

## Methods

### GetCondition

`func (o *GenerateRuleTitleRule) GetCondition() []map[string]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GenerateRuleTitleRule) GetConditionOk() ([]map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCondition

`func (o *GenerateRuleTitleRule) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetCondition

`func (o *GenerateRuleTitleRule) SetCondition(v []map[string]interface{})`

SetCondition gets a reference to the given []map[string]interface{} and assigns it to the Condition field.

### GetEffects

`func (o *GenerateRuleTitleRule) GetEffects() [][]map[string]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *GenerateRuleTitleRule) GetEffectsOk() ([][]map[string]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffects

`func (o *GenerateRuleTitleRule) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffects

`func (o *GenerateRuleTitleRule) SetEffects(v [][]map[string]interface{})`

SetEffects gets a reference to the given [][]map[string]interface{} and assigns it to the Effects field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


