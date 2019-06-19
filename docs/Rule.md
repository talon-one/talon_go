# Rule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | A short description of the rule. | [default to null]
**Description** | **string** | A longer, more detailed description of the rule. | [optional] [default to null]
**Bindings** | [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | [optional] [default to null]
**Condition** | [**[]interface{}**](interface{}.md) | A Talang expression that will be evaluated in the context of the given event. | [default to null]
**Effects** | [**[]interface{}**](interface{}.md) | An array of effectful Talang expressions in arrays that will be evaluated when a rule matches. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


