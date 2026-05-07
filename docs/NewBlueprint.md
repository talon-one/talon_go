# NewBlueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The display name for the blueprint. | 
**Description** | Pointer to **string** | A longer, more detailed description of the blueprint. | [optional] 
**Category** | Pointer to **string** | Category used to group blueprints. | [optional] [default to "custom"]
**Rules** | Pointer to [**[]CatalogRule**](CatalogRule.md) | Array of rules to store in this blueprint. Rules should only contain title (no description, as description is at the blueprint level). At least one rule or cart item filter is required. | [optional] 
**CartItemFilters** | Pointer to [**[]CartItemFilterTemplate**](CartItemFilterTemplate.md) | Array of cart item filters to store in this blueprint. If not provided, will be extracted from the rules. Cart item filters should only contain name (no description, as description is at the blueprint level). | [optional] 

## Methods

### NewNewBlueprint

`func NewNewBlueprint(title string, ) *NewBlueprint`

NewNewBlueprint instantiates a new NewBlueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewBlueprintWithDefaults

`func NewNewBlueprintWithDefaults() *NewBlueprint`

NewNewBlueprintWithDefaults instantiates a new NewBlueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewBlueprint) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewBlueprint) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewBlueprint) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewBlueprint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewBlueprint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewBlueprint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewBlueprint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *NewBlueprint) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NewBlueprint) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NewBlueprint) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NewBlueprint) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetRules

`func (o *NewBlueprint) GetRules() []CatalogRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NewBlueprint) GetRulesOk() (*[]CatalogRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NewBlueprint) SetRules(v []CatalogRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *NewBlueprint) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetCartItemFilters

`func (o *NewBlueprint) GetCartItemFilters() []CartItemFilterTemplate`

GetCartItemFilters returns the CartItemFilters field if non-nil, zero value otherwise.

### GetCartItemFiltersOk

`func (o *NewBlueprint) GetCartItemFiltersOk() (*[]CartItemFilterTemplate, bool)`

GetCartItemFiltersOk returns a tuple with the CartItemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemFilters

`func (o *NewBlueprint) SetCartItemFilters(v []CartItemFilterTemplate)`

SetCartItemFilters sets CartItemFilters field to given value.

### HasCartItemFilters

`func (o *NewBlueprint) HasCartItemFilters() bool`

HasCartItemFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


