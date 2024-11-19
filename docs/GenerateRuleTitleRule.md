# GenerateRuleTitleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effects** | Pointer to [**[][]interface{}**]([][]interface{}.md) | An array of effectful Talang expressions in arrays that will be evaluated when a rule matches. | [optional] 
**Condition** | Pointer to [**[]interface{}**]([]interface{}.md) | A Talang expression that will be evaluated in the context of the given event. | [optional] 

## Methods

### GetEffects

`func (o *GenerateRuleTitleRule) GetEffects() [][]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *GenerateRuleTitleRule) GetEffectsOk() ([][]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffects

`func (o *GenerateRuleTitleRule) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffects

`func (o *GenerateRuleTitleRule) SetEffects(v [][]interface{})`

SetEffects gets a reference to the given [][]interface{} and assigns it to the Effects field.

### GetCondition

`func (o *GenerateRuleTitleRule) GetCondition() []interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GenerateRuleTitleRule) GetConditionOk() ([]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCondition

`func (o *GenerateRuleTitleRule) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetCondition

`func (o *GenerateRuleTitleRule) SetCondition(v []interface{})`

SetCondition gets a reference to the given []interface{} and assigns it to the Condition field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


