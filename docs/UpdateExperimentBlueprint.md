# UpdateExperimentBlueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The display name for the experiment blueprint. | [optional] 
**Description** | Pointer to **string** | A longer, more detailed description of the experiment blueprint. | [optional] 
**Category** | Pointer to **string** | Category used to group experiment blueprints. | [optional] 
**Variants** | Pointer to [**[]ExperimentBlueprintVariant**](ExperimentBlueprintVariant.md) | Replaces the stored experiment variant templates. | [optional] 
**CartItemFilters** | Pointer to [**[]CartItemFilterTemplate**](CartItemFilterTemplate.md) | Replaces the stored cart item filters. Cart item filters should only contain name (no description, as description is at the experiment blueprint level). | [optional] 

## Methods

### NewUpdateExperimentBlueprint

`func NewUpdateExperimentBlueprint() *UpdateExperimentBlueprint`

NewUpdateExperimentBlueprint instantiates a new UpdateExperimentBlueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExperimentBlueprintWithDefaults

`func NewUpdateExperimentBlueprintWithDefaults() *UpdateExperimentBlueprint`

NewUpdateExperimentBlueprintWithDefaults instantiates a new UpdateExperimentBlueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateExperimentBlueprint) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateExperimentBlueprint) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateExperimentBlueprint) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateExperimentBlueprint) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateExperimentBlueprint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateExperimentBlueprint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateExperimentBlueprint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateExperimentBlueprint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateExperimentBlueprint) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateExperimentBlueprint) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateExperimentBlueprint) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateExperimentBlueprint) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVariants

`func (o *UpdateExperimentBlueprint) GetVariants() []ExperimentBlueprintVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *UpdateExperimentBlueprint) GetVariantsOk() (*[]ExperimentBlueprintVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *UpdateExperimentBlueprint) SetVariants(v []ExperimentBlueprintVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *UpdateExperimentBlueprint) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetCartItemFilters

`func (o *UpdateExperimentBlueprint) GetCartItemFilters() []CartItemFilterTemplate`

GetCartItemFilters returns the CartItemFilters field if non-nil, zero value otherwise.

### GetCartItemFiltersOk

`func (o *UpdateExperimentBlueprint) GetCartItemFiltersOk() (*[]CartItemFilterTemplate, bool)`

GetCartItemFiltersOk returns a tuple with the CartItemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemFilters

`func (o *UpdateExperimentBlueprint) SetCartItemFilters(v []CartItemFilterTemplate)`

SetCartItemFilters sets CartItemFilters field to given value.

### HasCartItemFilters

`func (o *UpdateExperimentBlueprint) HasCartItemFilters() bool`

HasCartItemFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


