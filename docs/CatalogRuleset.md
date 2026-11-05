# CatalogRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]CatalogRule**](CatalogRule.md) | Array of rule templates in this ruleset. Rules only contain title (no description, as description is at the catalog entry level). | 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | [optional] 
**RbVersion** | Pointer to **string** | The version of the rulebuilder used to create this ruleset. | [optional] 

## Methods

### NewCatalogRuleset

`func NewCatalogRuleset(rules []CatalogRule, ) *CatalogRuleset`

NewCatalogRuleset instantiates a new CatalogRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogRulesetWithDefaults

`func NewCatalogRulesetWithDefaults() *CatalogRuleset`

NewCatalogRulesetWithDefaults instantiates a new CatalogRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *CatalogRuleset) GetRules() []CatalogRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CatalogRuleset) GetRulesOk() (*[]CatalogRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CatalogRuleset) SetRules(v []CatalogRule)`

SetRules sets Rules field to given value.


### GetBindings

`func (o *CatalogRuleset) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *CatalogRuleset) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *CatalogRuleset) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *CatalogRuleset) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetRbVersion

`func (o *CatalogRuleset) GetRbVersion() string`

GetRbVersion returns the RbVersion field if non-nil, zero value otherwise.

### GetRbVersionOk

`func (o *CatalogRuleset) GetRbVersionOk() (*string, bool)`

GetRbVersionOk returns a tuple with the RbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbVersion

`func (o *CatalogRuleset) SetRbVersion(v string)`

SetRbVersion sets RbVersion field to given value.

### HasRbVersion

`func (o *CatalogRuleset) HasRbVersion() bool`

HasRbVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


