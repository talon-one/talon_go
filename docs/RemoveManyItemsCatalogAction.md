# RemoveManyItemsCatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]CatalogActionFilter**](CatalogActionFilter.md) | The list of filters used to select the items to patch, joined by &#x60;AND&#x60;.  **Note:** Every item in the catalog will be removed if there are no filters.  | [optional] 

## Methods

### GetFilters

`func (o *RemoveManyItemsCatalogAction) GetFilters() []CatalogActionFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RemoveManyItemsCatalogAction) GetFiltersOk() ([]CatalogActionFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilters

`func (o *RemoveManyItemsCatalogAction) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFilters

`func (o *RemoveManyItemsCatalogAction) SetFilters(v []CatalogActionFilter)`

SetFilters gets a reference to the given []CatalogActionFilter and assigns it to the Filters field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


