# ScimResourceTypesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ScimResource**](ScimResource.md) |  | 

## Methods

### NewScimResourceTypesListResponse

`func NewScimResourceTypesListResponse(resources []ScimResource, ) *ScimResourceTypesListResponse`

NewScimResourceTypesListResponse instantiates a new ScimResourceTypesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimResourceTypesListResponseWithDefaults

`func NewScimResourceTypesListResponseWithDefaults() *ScimResourceTypesListResponse`

NewScimResourceTypesListResponseWithDefaults instantiates a new ScimResourceTypesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ScimResourceTypesListResponse) GetResources() []ScimResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ScimResourceTypesListResponse) GetResourcesOk() (*[]ScimResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ScimResourceTypesListResponse) SetResources(v []ScimResource)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


