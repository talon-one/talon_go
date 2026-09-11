# CatalogActionRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A catalog sync action discriminator of type &#x60;REMOVE&#x60;. | 
**Payload** | Pointer to [**RemoveItemCatalogAction**](RemoveItemCatalogAction.md) |  | 

## Methods

### NewCatalogActionRemove

`func NewCatalogActionRemove(type_ string, payload RemoveItemCatalogAction, ) *CatalogActionRemove`

NewCatalogActionRemove instantiates a new CatalogActionRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogActionRemoveWithDefaults

`func NewCatalogActionRemoveWithDefaults() *CatalogActionRemove`

NewCatalogActionRemoveWithDefaults instantiates a new CatalogActionRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogActionRemove) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogActionRemove) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogActionRemove) SetType(v string)`

SetType sets Type field to given value.


### GetPayload

`func (o *CatalogActionRemove) GetPayload() RemoveItemCatalogAction`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CatalogActionRemove) GetPayloadOk() (*RemoveItemCatalogAction, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CatalogActionRemove) SetPayload(v RemoveItemCatalogAction)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


