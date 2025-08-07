# ScimGroupsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ScimGroup**](ScimGroup.md) |  | 
**Schemas** | Pointer to **[]string** | SCIM schema for the given resource. | [optional] 
**TotalResults** | Pointer to **int32** | Number of results in the response. | [optional] 

## Methods

### GetResources

`func (o *ScimGroupsListResponse) GetResources() []ScimGroup`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ScimGroupsListResponse) GetResourcesOk() ([]ScimGroup, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResources

`func (o *ScimGroupsListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResources

`func (o *ScimGroupsListResponse) SetResources(v []ScimGroup)`

SetResources gets a reference to the given []ScimGroup and assigns it to the Resources field.

### GetSchemas

`func (o *ScimGroupsListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimGroupsListResponse) GetSchemasOk() ([]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchemas

`func (o *ScimGroupsListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### SetSchemas

`func (o *ScimGroupsListResponse) SetSchemas(v []string)`

SetSchemas gets a reference to the given []string and assigns it to the Schemas field.

### GetTotalResults

`func (o *ScimGroupsListResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ScimGroupsListResponse) GetTotalResultsOk() (int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalResults

`func (o *ScimGroupsListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### SetTotalResults

`func (o *ScimGroupsListResponse) SetTotalResults(v int32)`

SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


