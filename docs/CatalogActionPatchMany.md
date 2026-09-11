# CatalogActionPatchMany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A catalog sync action discriminator of type &#x60;PATCH_MANY&#x60;. | 
**Payload** | Pointer to [**PatchManyItemsCatalogAction**](PatchManyItemsCatalogAction.md) |  | 

## Methods

### NewCatalogActionPatchMany

`func NewCatalogActionPatchMany(type_ string, payload PatchManyItemsCatalogAction, ) *CatalogActionPatchMany`

NewCatalogActionPatchMany instantiates a new CatalogActionPatchMany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogActionPatchManyWithDefaults

`func NewCatalogActionPatchManyWithDefaults() *CatalogActionPatchMany`

NewCatalogActionPatchManyWithDefaults instantiates a new CatalogActionPatchMany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogActionPatchMany) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogActionPatchMany) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogActionPatchMany) SetType(v string)`

SetType sets Type field to given value.


### GetPayload

`func (o *CatalogActionPatchMany) GetPayload() PatchManyItemsCatalogAction`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CatalogActionPatchMany) GetPayloadOk() (*PatchManyItemsCatalogAction, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CatalogActionPatchMany) SetPayload(v PatchManyItemsCatalogAction)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


