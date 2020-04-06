# NewRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]Rule**](Rule.md) | Set of rules to apply. | 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | 
**RbVersion** | Pointer to **string** | A string indicating which version of the rulebuilder was used to create this ruleset. | [optional] 
**Activate** | Pointer to **bool** | A boolean indicating whether this newly created ruleset should also be activated for the campaign owns it | [optional] 

## Methods

### GetRules

`func (o *NewRuleset) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NewRuleset) GetRulesOk() ([]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRules

`func (o *NewRuleset) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRules

`func (o *NewRuleset) SetRules(v []Rule)`

SetRules gets a reference to the given []Rule and assigns it to the Rules field.

### GetBindings

`func (o *NewRuleset) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *NewRuleset) GetBindingsOk() ([]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBindings

`func (o *NewRuleset) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### SetBindings

`func (o *NewRuleset) SetBindings(v []Binding)`

SetBindings gets a reference to the given []Binding and assigns it to the Bindings field.

### GetRbVersion

`func (o *NewRuleset) GetRbVersion() string`

GetRbVersion returns the RbVersion field if non-nil, zero value otherwise.

### GetRbVersionOk

`func (o *NewRuleset) GetRbVersionOk() (string, bool)`

GetRbVersionOk returns a tuple with the RbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRbVersion

`func (o *NewRuleset) HasRbVersion() bool`

HasRbVersion returns a boolean if a field has been set.

### SetRbVersion

`func (o *NewRuleset) SetRbVersion(v string)`

SetRbVersion gets a reference to the given string and assigns it to the RbVersion field.

### GetActivate

`func (o *NewRuleset) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *NewRuleset) GetActivateOk() (bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivate

`func (o *NewRuleset) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### SetActivate

`func (o *NewRuleset) SetActivate(v bool)`

SetActivate gets a reference to the given bool and assigns it to the Activate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


