# AddPriceAdjustmentCatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **string** | The SKU of the item for which the price is being adjusted. | 
**Adjustments** | Pointer to [**[]NewPriceAdjustment**](NewPriceAdjustment.md) | A list of adjustments to apply to a given item. | 

## Methods

### GetSku

`func (o *AddPriceAdjustmentCatalogAction) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AddPriceAdjustmentCatalogAction) GetSkuOk() (string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSku

`func (o *AddPriceAdjustmentCatalogAction) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSku

`func (o *AddPriceAdjustmentCatalogAction) SetSku(v string)`

SetSku gets a reference to the given string and assigns it to the Sku field.

### GetAdjustments

`func (o *AddPriceAdjustmentCatalogAction) GetAdjustments() []NewPriceAdjustment`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### GetAdjustmentsOk

`func (o *AddPriceAdjustmentCatalogAction) GetAdjustmentsOk() ([]NewPriceAdjustment, bool)`

GetAdjustmentsOk returns a tuple with the Adjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdjustments

`func (o *AddPriceAdjustmentCatalogAction) HasAdjustments() bool`

HasAdjustments returns a boolean if a field has been set.

### SetAdjustments

`func (o *AddPriceAdjustmentCatalogAction) SetAdjustments(v []NewPriceAdjustment)`

SetAdjustments gets a reference to the given []NewPriceAdjustment and assigns it to the Adjustments field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


