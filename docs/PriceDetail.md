# PriceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** | The value of this price type. | [optional] 
**AdjustmentReferenceId** | Pointer to **string** | The reference identifier of the selected price adjustment for this SKU. | [optional] 
**AdjustmentEffectiveFrom** | Pointer to [**time.Time**](time.Time.md) | The date and time from which the price adjustment is effective. | [optional] 
**AdjustmentEffectiveUntil** | Pointer to [**time.Time**](time.Time.md) | The date and time until which the price adjustment is effective. | [optional] 

## Methods

### GetPrice

`func (o *PriceDetail) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceDetail) GetPriceOk() (float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *PriceDetail) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *PriceDetail) SetPrice(v float32)`

SetPrice gets a reference to the given float32 and assigns it to the Price field.

### GetAdjustmentReferenceId

`func (o *PriceDetail) GetAdjustmentReferenceId() string`

GetAdjustmentReferenceId returns the AdjustmentReferenceId field if non-nil, zero value otherwise.

### GetAdjustmentReferenceIdOk

`func (o *PriceDetail) GetAdjustmentReferenceIdOk() (string, bool)`

GetAdjustmentReferenceIdOk returns a tuple with the AdjustmentReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdjustmentReferenceId

`func (o *PriceDetail) HasAdjustmentReferenceId() bool`

HasAdjustmentReferenceId returns a boolean if a field has been set.

### SetAdjustmentReferenceId

`func (o *PriceDetail) SetAdjustmentReferenceId(v string)`

SetAdjustmentReferenceId gets a reference to the given string and assigns it to the AdjustmentReferenceId field.

### GetAdjustmentEffectiveFrom

`func (o *PriceDetail) GetAdjustmentEffectiveFrom() time.Time`

GetAdjustmentEffectiveFrom returns the AdjustmentEffectiveFrom field if non-nil, zero value otherwise.

### GetAdjustmentEffectiveFromOk

`func (o *PriceDetail) GetAdjustmentEffectiveFromOk() (time.Time, bool)`

GetAdjustmentEffectiveFromOk returns a tuple with the AdjustmentEffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdjustmentEffectiveFrom

`func (o *PriceDetail) HasAdjustmentEffectiveFrom() bool`

HasAdjustmentEffectiveFrom returns a boolean if a field has been set.

### SetAdjustmentEffectiveFrom

`func (o *PriceDetail) SetAdjustmentEffectiveFrom(v time.Time)`

SetAdjustmentEffectiveFrom gets a reference to the given time.Time and assigns it to the AdjustmentEffectiveFrom field.

### GetAdjustmentEffectiveUntil

`func (o *PriceDetail) GetAdjustmentEffectiveUntil() time.Time`

GetAdjustmentEffectiveUntil returns the AdjustmentEffectiveUntil field if non-nil, zero value otherwise.

### GetAdjustmentEffectiveUntilOk

`func (o *PriceDetail) GetAdjustmentEffectiveUntilOk() (time.Time, bool)`

GetAdjustmentEffectiveUntilOk returns a tuple with the AdjustmentEffectiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdjustmentEffectiveUntil

`func (o *PriceDetail) HasAdjustmentEffectiveUntil() bool`

HasAdjustmentEffectiveUntil returns a boolean if a field has been set.

### SetAdjustmentEffectiveUntil

`func (o *PriceDetail) SetAdjustmentEffectiveUntil(v time.Time)`

SetAdjustmentEffectiveUntil gets a reference to the given time.Time and assigns it to the AdjustmentEffectiveUntil field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


