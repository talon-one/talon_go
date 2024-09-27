# ScimSchemasListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ScimSchemaResource**](ScimSchemaResource.md) |  | 
**Schemas** | Pointer to **[]string** | SCIM schema for the given resource. | [optional] 
**TotalResults** | Pointer to **int32** | Number of total results in the response. | [optional] 

## Methods

### GetResources

`func (o *ScimSchemasListResponse) GetResources() []ScimSchemaResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ScimSchemasListResponse) GetResourcesOk() ([]ScimSchemaResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResources

`func (o *ScimSchemasListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResources

`func (o *ScimSchemasListResponse) SetResources(v []ScimSchemaResource)`

SetResources gets a reference to the given []ScimSchemaResource and assigns it to the Resources field.

### GetSchemas

`func (o *ScimSchemasListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimSchemasListResponse) GetSchemasOk() ([]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchemas

`func (o *ScimSchemasListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### SetSchemas

`func (o *ScimSchemasListResponse) SetSchemas(v []string)`

SetSchemas gets a reference to the given []string and assigns it to the Schemas field.

### GetTotalResults

`func (o *ScimSchemasListResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ScimSchemasListResponse) GetTotalResultsOk() (int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalResults

`func (o *ScimSchemasListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### SetTotalResults

`func (o *ScimSchemasListResponse) SetTotalResults(v int32)`

SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


