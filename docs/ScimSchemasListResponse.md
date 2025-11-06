# ScimSchemasListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ScimSchemaResource**](ScimSchemaResource.md) |  | 
**Schemas** | Pointer to **[]string** | SCIM schema for the given resource. | [optional] 
**TotalResults** | Pointer to **int64** | Number of total results in the response. | [optional] 

## Methods

### NewScimSchemasListResponse

`func NewScimSchemasListResponse(resources []ScimSchemaResource, ) *ScimSchemasListResponse`

NewScimSchemasListResponse instantiates a new ScimSchemasListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSchemasListResponseWithDefaults

`func NewScimSchemasListResponseWithDefaults() *ScimSchemasListResponse`

NewScimSchemasListResponseWithDefaults instantiates a new ScimSchemasListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ScimSchemasListResponse) GetResources() []ScimSchemaResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ScimSchemasListResponse) GetResourcesOk() (*[]ScimSchemaResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ScimSchemasListResponse) SetResources(v []ScimSchemaResource)`

SetResources sets Resources field to given value.


### GetSchemas

`func (o *ScimSchemasListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimSchemasListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimSchemasListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ScimSchemasListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *ScimSchemasListResponse) GetTotalResults() int64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ScimSchemasListResponse) GetTotalResultsOk() (*int64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ScimSchemasListResponse) SetTotalResults(v int64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ScimSchemasListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


