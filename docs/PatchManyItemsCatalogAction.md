# PatchManyItemsCatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Filters** | Pointer to [**[]CatalogActionFilter**](CatalogActionFilter.md) | The list of filters used to select the items to patch, joined by &#x60;AND&#x60;.  **Note:** Every item in the catalog will be modified if there are no filters.  | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | The attributes of the items to patch. | [optional] 

## Methods

### NewPatchManyItemsCatalogAction

`func NewPatchManyItemsCatalogAction() *PatchManyItemsCatalogAction`

NewPatchManyItemsCatalogAction instantiates a new PatchManyItemsCatalogAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchManyItemsCatalogActionWithDefaults

`func NewPatchManyItemsCatalogActionWithDefaults() *PatchManyItemsCatalogAction`

NewPatchManyItemsCatalogActionWithDefaults instantiates a new PatchManyItemsCatalogAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *PatchManyItemsCatalogAction) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PatchManyItemsCatalogAction) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PatchManyItemsCatalogAction) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PatchManyItemsCatalogAction) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFilters

`func (o *PatchManyItemsCatalogAction) GetFilters() []CatalogActionFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchManyItemsCatalogAction) GetFiltersOk() (*[]CatalogActionFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchManyItemsCatalogAction) SetFilters(v []CatalogActionFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchManyItemsCatalogAction) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetAttributes

`func (o *PatchManyItemsCatalogAction) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchManyItemsCatalogAction) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PatchManyItemsCatalogAction) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PatchManyItemsCatalogAction) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


