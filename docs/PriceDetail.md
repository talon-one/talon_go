# PriceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** | The value of this price type. | [optional] 
**AdjustmentContextId** | Pointer to **string** | The context identifier of the selected price adjustment. | [optional] 
**AdjustmentReferenceId** | Pointer to **string** | The reference identifier of the selected price adjustment for this SKU. | [optional] 
**AdjustmentEffectiveFrom** | Pointer to [**time.Time**](time.Time.md) | The date and time from which the price adjustment is effective. | [optional] 
**AdjustmentEffectiveUntil** | Pointer to [**time.Time**](time.Time.md) | The date and time until which the price adjustment is effective. | [optional] 

## Methods

### NewPriceDetail

`func NewPriceDetail() *PriceDetail`

NewPriceDetail instantiates a new PriceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceDetailWithDefaults

`func NewPriceDetailWithDefaults() *PriceDetail`

NewPriceDetailWithDefaults instantiates a new PriceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *PriceDetail) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceDetail) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceDetail) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PriceDetail) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAdjustmentContextId

`func (o *PriceDetail) GetAdjustmentContextId() string`

GetAdjustmentContextId returns the AdjustmentContextId field if non-nil, zero value otherwise.

### GetAdjustmentContextIdOk

`func (o *PriceDetail) GetAdjustmentContextIdOk() (*string, bool)`

GetAdjustmentContextIdOk returns a tuple with the AdjustmentContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentContextId

`func (o *PriceDetail) SetAdjustmentContextId(v string)`

SetAdjustmentContextId sets AdjustmentContextId field to given value.

### HasAdjustmentContextId

`func (o *PriceDetail) HasAdjustmentContextId() bool`

HasAdjustmentContextId returns a boolean if a field has been set.

### GetAdjustmentReferenceId

`func (o *PriceDetail) GetAdjustmentReferenceId() string`

GetAdjustmentReferenceId returns the AdjustmentReferenceId field if non-nil, zero value otherwise.

### GetAdjustmentReferenceIdOk

`func (o *PriceDetail) GetAdjustmentReferenceIdOk() (*string, bool)`

GetAdjustmentReferenceIdOk returns a tuple with the AdjustmentReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentReferenceId

`func (o *PriceDetail) SetAdjustmentReferenceId(v string)`

SetAdjustmentReferenceId sets AdjustmentReferenceId field to given value.

### HasAdjustmentReferenceId

`func (o *PriceDetail) HasAdjustmentReferenceId() bool`

HasAdjustmentReferenceId returns a boolean if a field has been set.

### GetAdjustmentEffectiveFrom

`func (o *PriceDetail) GetAdjustmentEffectiveFrom() time.Time`

GetAdjustmentEffectiveFrom returns the AdjustmentEffectiveFrom field if non-nil, zero value otherwise.

### GetAdjustmentEffectiveFromOk

`func (o *PriceDetail) GetAdjustmentEffectiveFromOk() (*time.Time, bool)`

GetAdjustmentEffectiveFromOk returns a tuple with the AdjustmentEffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentEffectiveFrom

`func (o *PriceDetail) SetAdjustmentEffectiveFrom(v time.Time)`

SetAdjustmentEffectiveFrom sets AdjustmentEffectiveFrom field to given value.

### HasAdjustmentEffectiveFrom

`func (o *PriceDetail) HasAdjustmentEffectiveFrom() bool`

HasAdjustmentEffectiveFrom returns a boolean if a field has been set.

### GetAdjustmentEffectiveUntil

`func (o *PriceDetail) GetAdjustmentEffectiveUntil() time.Time`

GetAdjustmentEffectiveUntil returns the AdjustmentEffectiveUntil field if non-nil, zero value otherwise.

### GetAdjustmentEffectiveUntilOk

`func (o *PriceDetail) GetAdjustmentEffectiveUntilOk() (*time.Time, bool)`

GetAdjustmentEffectiveUntilOk returns a tuple with the AdjustmentEffectiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentEffectiveUntil

`func (o *PriceDetail) SetAdjustmentEffectiveUntil(v time.Time)`

SetAdjustmentEffectiveUntil sets AdjustmentEffectiveUntil field to given value.

### HasAdjustmentEffectiveUntil

`func (o *PriceDetail) HasAdjustmentEffectiveUntil() bool`

HasAdjustmentEffectiveUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


