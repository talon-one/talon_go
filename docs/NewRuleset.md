# NewRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]Rule**](Rule.md) | Set of rules to apply. | 
**StrikethroughRules** | Pointer to [**[]Rule**](Rule.md) | Set of rules to apply for strikethrough. | [optional] 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | 
**RbVersion** | Pointer to **string** | The version of the rulebuilder used to create this ruleset. | [optional] 
**Activate** | Pointer to **bool** | Indicates whether this created ruleset should be activated for the campaign that owns it. | [optional] 

## Methods

### NewNewRuleset

`func NewNewRuleset(rules []Rule, bindings []Binding, ) *NewRuleset`

NewNewRuleset instantiates a new NewRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRulesetWithDefaults

`func NewNewRulesetWithDefaults() *NewRuleset`

NewNewRulesetWithDefaults instantiates a new NewRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *NewRuleset) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NewRuleset) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NewRuleset) SetRules(v []Rule)`

SetRules sets Rules field to given value.


### GetStrikethroughRules

`func (o *NewRuleset) GetStrikethroughRules() []Rule`

GetStrikethroughRules returns the StrikethroughRules field if non-nil, zero value otherwise.

### GetStrikethroughRulesOk

`func (o *NewRuleset) GetStrikethroughRulesOk() (*[]Rule, bool)`

GetStrikethroughRulesOk returns a tuple with the StrikethroughRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikethroughRules

`func (o *NewRuleset) SetStrikethroughRules(v []Rule)`

SetStrikethroughRules sets StrikethroughRules field to given value.

### HasStrikethroughRules

`func (o *NewRuleset) HasStrikethroughRules() bool`

HasStrikethroughRules returns a boolean if a field has been set.

### GetBindings

`func (o *NewRuleset) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *NewRuleset) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *NewRuleset) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.


### GetRbVersion

`func (o *NewRuleset) GetRbVersion() string`

GetRbVersion returns the RbVersion field if non-nil, zero value otherwise.

### GetRbVersionOk

`func (o *NewRuleset) GetRbVersionOk() (*string, bool)`

GetRbVersionOk returns a tuple with the RbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbVersion

`func (o *NewRuleset) SetRbVersion(v string)`

SetRbVersion sets RbVersion field to given value.

### HasRbVersion

`func (o *NewRuleset) HasRbVersion() bool`

HasRbVersion returns a boolean if a field has been set.

### GetActivate

`func (o *NewRuleset) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *NewRuleset) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *NewRuleset) SetActivate(v bool)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *NewRuleset) HasActivate() bool`

HasActivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


