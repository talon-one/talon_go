# CatalogActionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A catalog sync action discriminator of type &#x60;PATCH&#x60;. | 
**Payload** | Pointer to [**PatchItemCatalogAction**](PatchItemCatalogAction.md) |  | 

## Methods

### NewCatalogActionPatch

`func NewCatalogActionPatch(type_ string, payload PatchItemCatalogAction, ) *CatalogActionPatch`

NewCatalogActionPatch instantiates a new CatalogActionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogActionPatchWithDefaults

`func NewCatalogActionPatchWithDefaults() *CatalogActionPatch`

NewCatalogActionPatchWithDefaults instantiates a new CatalogActionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogActionPatch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogActionPatch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogActionPatch) SetType(v string)`

SetType sets Type field to given value.


### GetPayload

`func (o *CatalogActionPatch) GetPayload() PatchItemCatalogAction`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CatalogActionPatch) GetPayloadOk() (*PatchItemCatalogAction, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CatalogActionPatch) SetPayload(v PatchItemCatalogAction)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


