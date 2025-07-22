# GenerateRuleTitleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effects** | Pointer to **[][]interface{}** | An array of effectful Talang expressions in arrays that will be evaluated when a rule matches. | [optional] 
**Condition** | Pointer to **[]interface{}** | A Talang expression that will be evaluated in the context of the given event. | [optional] 

## Methods

### NewGenerateRuleTitleRule

`func NewGenerateRuleTitleRule() *GenerateRuleTitleRule`

NewGenerateRuleTitleRule instantiates a new GenerateRuleTitleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateRuleTitleRuleWithDefaults

`func NewGenerateRuleTitleRuleWithDefaults() *GenerateRuleTitleRule`

NewGenerateRuleTitleRuleWithDefaults instantiates a new GenerateRuleTitleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffects

`func (o *GenerateRuleTitleRule) GetEffects() [][]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *GenerateRuleTitleRule) GetEffectsOk() (*[][]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *GenerateRuleTitleRule) SetEffects(v [][]interface{})`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *GenerateRuleTitleRule) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### GetCondition

`func (o *GenerateRuleTitleRule) GetCondition() []interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GenerateRuleTitleRule) GetConditionOk() (*[]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GenerateRuleTitleRule) SetCondition(v []interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GenerateRuleTitleRule) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


