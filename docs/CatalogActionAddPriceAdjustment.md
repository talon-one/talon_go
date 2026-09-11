# CatalogActionAddPriceAdjustment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A catalog sync action discriminator of type &#x60;ADD_PRICE_ADJUSTMENT&#x60;. | 
**Payload** | Pointer to [**AddPriceAdjustmentCatalogAction**](AddPriceAdjustmentCatalogAction.md) |  | 

## Methods

### NewCatalogActionAddPriceAdjustment

`func NewCatalogActionAddPriceAdjustment(type_ string, payload AddPriceAdjustmentCatalogAction, ) *CatalogActionAddPriceAdjustment`

NewCatalogActionAddPriceAdjustment instantiates a new CatalogActionAddPriceAdjustment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogActionAddPriceAdjustmentWithDefaults

`func NewCatalogActionAddPriceAdjustmentWithDefaults() *CatalogActionAddPriceAdjustment`

NewCatalogActionAddPriceAdjustmentWithDefaults instantiates a new CatalogActionAddPriceAdjustment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogActionAddPriceAdjustment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogActionAddPriceAdjustment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogActionAddPriceAdjustment) SetType(v string)`

SetType sets Type field to given value.


### GetPayload

`func (o *CatalogActionAddPriceAdjustment) GetPayload() AddPriceAdjustmentCatalogAction`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CatalogActionAddPriceAdjustment) GetPayloadOk() (*AddPriceAdjustmentCatalogAction, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CatalogActionAddPriceAdjustment) SetPayload(v AddPriceAdjustmentCatalogAction)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


