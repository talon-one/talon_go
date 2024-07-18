# PatchManyItemsCatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** | Price of the item. | [optional] 
**Filters** | Pointer to [**[]CatalogActionFilter**](CatalogActionFilter.md) | The list of filters used to select the items to patch, joined by &#x60;AND&#x60;.  **Note:** Every item in the catalog will be modified if there are no filters.  | [optional] 
**Attributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | The attributes of the items to patch. | [optional] 

## Methods

### GetPrice

`func (o *PatchManyItemsCatalogAction) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PatchManyItemsCatalogAction) GetPriceOk() (float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *PatchManyItemsCatalogAction) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *PatchManyItemsCatalogAction) SetPrice(v float32)`

SetPrice gets a reference to the given float32 and assigns it to the Price field.

### GetFilters

`func (o *PatchManyItemsCatalogAction) GetFilters() []CatalogActionFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchManyItemsCatalogAction) GetFiltersOk() ([]CatalogActionFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilters

`func (o *PatchManyItemsCatalogAction) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFilters

`func (o *PatchManyItemsCatalogAction) SetFilters(v []CatalogActionFilter)`

SetFilters gets a reference to the given []CatalogActionFilter and assigns it to the Filters field.

### GetAttributes

`func (o *PatchManyItemsCatalogAction) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PatchManyItemsCatalogAction) GetAttributesOk() (map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *PatchManyItemsCatalogAction) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *PatchManyItemsCatalogAction) SetAttributes(v map[string]map[string]interface{})`

SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


