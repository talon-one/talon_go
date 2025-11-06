# ScimPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** | SCIM schema for the given resource. | [optional] 
**Operations** | Pointer to [**[]ScimPatchOperation**](ScimPatchOperation.md) |  | 

## Methods

### NewScimPatchRequest

`func NewScimPatchRequest(operations []ScimPatchOperation, ) *ScimPatchRequest`

NewScimPatchRequest instantiates a new ScimPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimPatchRequestWithDefaults

`func NewScimPatchRequestWithDefaults() *ScimPatchRequest`

NewScimPatchRequestWithDefaults instantiates a new ScimPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimPatchRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimPatchRequest) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimPatchRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimPatchRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetOperations

`func (o *ScimPatchRequest) GetOperations() []ScimPatchOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ScimPatchRequest) GetOperationsOk() (*[]ScimPatchOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ScimPatchRequest) SetOperations(v []ScimPatchOperation)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


