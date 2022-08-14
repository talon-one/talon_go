# CatalogSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]CatalogAction**](CatalogAction.md) |  | 
**Version** | Pointer to **int32** | The version number of the catalog to apply the actions on. | [optional] 

## Methods

### GetActions

`func (o *CatalogSyncRequest) GetActions() []CatalogAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CatalogSyncRequest) GetActionsOk() ([]CatalogAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActions

`func (o *CatalogSyncRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActions

`func (o *CatalogSyncRequest) SetActions(v []CatalogAction)`

SetActions gets a reference to the given []CatalogAction and assigns it to the Actions field.

### GetVersion

`func (o *CatalogSyncRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogSyncRequest) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *CatalogSyncRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *CatalogSyncRequest) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


