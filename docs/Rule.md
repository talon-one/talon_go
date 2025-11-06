# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the rule. | [optional] 
**ParentId** | Pointer to **string** | The ID of the rule that was copied to create this rule. | [optional] 
**Title** | Pointer to **string** | A short description of the rule. | 
**Description** | Pointer to **string** | A longer, more detailed description of the rule. | [optional] 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | [optional] 
**Condition** | Pointer to **[]map[string]interface{}** | A Talang expression that will be evaluated in the context of the given event. | 
**Effects** | Pointer to **[]map[string]interface{}** | An array of effectful Talang expressions in arrays that will be evaluated when a rule matches. | 

## Methods

### NewRule

`func NewRule(title string, condition []map[string]interface{}, effects []map[string]interface{}, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Rule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Rule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *Rule) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Rule) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Rule) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Rule) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTitle

`func (o *Rule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Rule) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Rule) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Rule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Rule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Rule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Rule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBindings

`func (o *Rule) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *Rule) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *Rule) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *Rule) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetCondition

`func (o *Rule) GetCondition() []interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Rule) GetConditionOk() (*[]map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Rule) SetCondition(v []interface{})`

SetCondition sets Condition field to given value.


### GetEffects

`func (o *Rule) GetEffects() [][]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *Rule) GetEffectsOk() (*[]map[string]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *Rule) SetEffects(v [][]interface{})`

SetEffects sets Effects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


