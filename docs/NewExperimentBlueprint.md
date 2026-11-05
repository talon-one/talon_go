# NewExperimentBlueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The display name for the experiment blueprint. | 
**Description** | Pointer to **string** | A longer, more detailed description of the experiment blueprint. | [optional] 
**Category** | Pointer to **string** | Category used to group experiment blueprints. | [optional] [default to "custom"]
**Variants** | Pointer to [**[]ExperimentBlueprintVariant**](ExperimentBlueprintVariant.md) | Array of experiment variant templates to store in this experiment blueprint. At least one variant or cart item filter is required. | [optional] 
**CartItemFilters** | Pointer to [**[]CartItemFilterTemplate**](CartItemFilterTemplate.md) | Array of cart item filters to store in this experiment blueprint. If not provided, will be extracted from the variant rulesets. Cart item filters should only contain name (no description, as description is at the experiment blueprint level). | [optional] 

## Methods

### NewNewExperimentBlueprint

`func NewNewExperimentBlueprint(title string, ) *NewExperimentBlueprint`

NewNewExperimentBlueprint instantiates a new NewExperimentBlueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewExperimentBlueprintWithDefaults

`func NewNewExperimentBlueprintWithDefaults() *NewExperimentBlueprint`

NewNewExperimentBlueprintWithDefaults instantiates a new NewExperimentBlueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewExperimentBlueprint) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewExperimentBlueprint) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewExperimentBlueprint) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewExperimentBlueprint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewExperimentBlueprint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewExperimentBlueprint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewExperimentBlueprint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *NewExperimentBlueprint) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NewExperimentBlueprint) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NewExperimentBlueprint) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NewExperimentBlueprint) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVariants

`func (o *NewExperimentBlueprint) GetVariants() []ExperimentBlueprintVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *NewExperimentBlueprint) GetVariantsOk() (*[]ExperimentBlueprintVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *NewExperimentBlueprint) SetVariants(v []ExperimentBlueprintVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *NewExperimentBlueprint) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetCartItemFilters

`func (o *NewExperimentBlueprint) GetCartItemFilters() []CartItemFilterTemplate`

GetCartItemFilters returns the CartItemFilters field if non-nil, zero value otherwise.

### GetCartItemFiltersOk

`func (o *NewExperimentBlueprint) GetCartItemFiltersOk() (*[]CartItemFilterTemplate, bool)`

GetCartItemFiltersOk returns a tuple with the CartItemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemFilters

`func (o *NewExperimentBlueprint) SetCartItemFilters(v []CartItemFilterTemplate)`

SetCartItemFilters sets CartItemFilters field to given value.

### HasCartItemFilters

`func (o *NewExperimentBlueprint) HasCartItemFilters() bool`

HasCartItemFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


