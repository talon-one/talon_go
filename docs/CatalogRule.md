# CatalogRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A short description of the rule. | 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array. | [optional] 
**Condition** | Pointer to **[]map[string]interface{}** | A Talang expression that will be evaluated in the context of the given event. | 
**Effects** | Pointer to **[]map[string]interface{}** | An array of effectful Talang expressions in arrays that will be evaluated when a rule matches. | 

## Methods

### NewCatalogRule

`func NewCatalogRule(title string, condition []map[string]interface{}, effects []map[string]interface{}, ) *CatalogRule`

NewCatalogRule instantiates a new CatalogRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogRuleWithDefaults

`func NewCatalogRuleWithDefaults() *CatalogRule`

NewCatalogRuleWithDefaults instantiates a new CatalogRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CatalogRule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CatalogRule) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CatalogRule) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBindings

`func (o *CatalogRule) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *CatalogRule) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *CatalogRule) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *CatalogRule) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetCondition

`func (o *CatalogRule) GetCondition() []map[string]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CatalogRule) GetConditionOk() (*[]map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CatalogRule) SetCondition(v []map[string]interface{})`

SetCondition sets Condition field to given value.


### GetEffects

`func (o *CatalogRule) GetEffects() []map[string]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *CatalogRule) GetEffectsOk() (*[]map[string]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *CatalogRule) SetEffects(v []map[string]interface{})`

SetEffects sets Effects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


