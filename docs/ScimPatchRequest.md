# ScimPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** | SCIM schema for the given resource. | [optional] 
**Operations** | Pointer to [**[]ScimPatchOperation**](ScimPatchOperation.md) |  | 

## Methods

### GetSchemas

`func (o *ScimPatchRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimPatchRequest) GetSchemasOk() ([]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchemas

`func (o *ScimPatchRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### SetSchemas

`func (o *ScimPatchRequest) SetSchemas(v []string)`

SetSchemas gets a reference to the given []string and assigns it to the Schemas field.

### GetOperations

`func (o *ScimPatchRequest) GetOperations() []ScimPatchOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ScimPatchRequest) GetOperationsOk() ([]ScimPatchOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperations

`func (o *ScimPatchRequest) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperations

`func (o *ScimPatchRequest) SetOperations(v []ScimPatchOperation)`

SetOperations gets a reference to the given []ScimPatchOperation and assigns it to the Operations field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


