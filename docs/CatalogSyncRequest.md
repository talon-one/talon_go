# CatalogSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]CatalogAction**](CatalogAction.md) |  | 
**Version** | Pointer to **int64** | The version number of the catalog to apply the actions on. | [optional] 

## Methods

### NewCatalogSyncRequest

`func NewCatalogSyncRequest(actions []CatalogAction, ) *CatalogSyncRequest`

NewCatalogSyncRequest instantiates a new CatalogSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogSyncRequestWithDefaults

`func NewCatalogSyncRequestWithDefaults() *CatalogSyncRequest`

NewCatalogSyncRequestWithDefaults instantiates a new CatalogSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *CatalogSyncRequest) GetActions() []CatalogAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CatalogSyncRequest) GetActionsOk() (*[]CatalogAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CatalogSyncRequest) SetActions(v []CatalogAction)`

SetActions sets Actions field to given value.


### GetVersion

`func (o *CatalogSyncRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogSyncRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogSyncRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogSyncRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


