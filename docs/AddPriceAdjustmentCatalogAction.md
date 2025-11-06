# AddPriceAdjustmentCatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | The SKU of the item for which the price is being adjusted. | 
**Adjustments** | Pointer to [**[]NewPriceAdjustment**](NewPriceAdjustment.md) | A list of adjustments to apply to a given item. | 

## Methods

### NewAddPriceAdjustmentCatalogAction

`func NewAddPriceAdjustmentCatalogAction(sku string, adjustments []NewPriceAdjustment, ) *AddPriceAdjustmentCatalogAction`

NewAddPriceAdjustmentCatalogAction instantiates a new AddPriceAdjustmentCatalogAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPriceAdjustmentCatalogActionWithDefaults

`func NewAddPriceAdjustmentCatalogActionWithDefaults() *AddPriceAdjustmentCatalogAction`

NewAddPriceAdjustmentCatalogActionWithDefaults instantiates a new AddPriceAdjustmentCatalogAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *AddPriceAdjustmentCatalogAction) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AddPriceAdjustmentCatalogAction) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *AddPriceAdjustmentCatalogAction) SetSku(v string)`

SetSku sets Sku field to given value.


### GetAdjustments

`func (o *AddPriceAdjustmentCatalogAction) GetAdjustments() []NewPriceAdjustment`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### GetAdjustmentsOk

`func (o *AddPriceAdjustmentCatalogAction) GetAdjustmentsOk() (*[]NewPriceAdjustment, bool)`

GetAdjustmentsOk returns a tuple with the Adjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustments

`func (o *AddPriceAdjustmentCatalogAction) SetAdjustments(v []NewPriceAdjustment)`

SetAdjustments sets Adjustments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


