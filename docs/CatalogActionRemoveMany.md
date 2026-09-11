# CatalogActionRemoveMany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A catalog sync action discriminator of type &#x60;REMOVE_MANY&#x60;. | 
**Payload** | Pointer to [**RemoveManyItemsCatalogAction**](RemoveManyItemsCatalogAction.md) |  | 

## Methods

### NewCatalogActionRemoveMany

`func NewCatalogActionRemoveMany(type_ string, payload RemoveManyItemsCatalogAction, ) *CatalogActionRemoveMany`

NewCatalogActionRemoveMany instantiates a new CatalogActionRemoveMany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogActionRemoveManyWithDefaults

`func NewCatalogActionRemoveManyWithDefaults() *CatalogActionRemoveMany`

NewCatalogActionRemoveManyWithDefaults instantiates a new CatalogActionRemoveMany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogActionRemoveMany) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogActionRemoveMany) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogActionRemoveMany) SetType(v string)`

SetType sets Type field to given value.


### GetPayload

`func (o *CatalogActionRemoveMany) GetPayload() RemoveManyItemsCatalogAction`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CatalogActionRemoveMany) GetPayloadOk() (*RemoveManyItemsCatalogAction, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CatalogActionRemoveMany) SetPayload(v RemoveManyItemsCatalogAction)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


