# NewPriceAdjustment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceType** | Pointer to **string** | The price type (e.g. the price for members only) to apply to a given SKU. | 
**Price** | Pointer to **NullableFloat32** | The value of the price type applied to the SKU. When set to &#x60;null&#x60;, the defined price type no longer applies to the SKU. | [optional] 
**ReferenceId** | Pointer to **string** | A unique reference identifier, e.g. a UUID. | 
**CalculatedAt** | Pointer to [**time.Time**](time.Time.md) | The time at which this price was calculated. If provided, this is used to determine the most recent price adjustment to choose if price adjustments overlap. Defaults to internal creation time if not provided. | [optional] 
**EffectiveFrom** | Pointer to [**time.Time**](time.Time.md) | The date and time from which the price adjustment is effective. | [optional] 
**EffectiveUntil** | Pointer to [**time.Time**](time.Time.md) | The date and time until which the price adjustment is effective. | [optional] 

## Methods

### GetPriceType

`func (o *NewPriceAdjustment) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *NewPriceAdjustment) GetPriceTypeOk() (string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriceType

`func (o *NewPriceAdjustment) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### SetPriceType

`func (o *NewPriceAdjustment) SetPriceType(v string)`

SetPriceType gets a reference to the given string and assigns it to the PriceType field.

### GetPrice

`func (o *NewPriceAdjustment) GetPrice() NullableFloat32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NewPriceAdjustment) GetPriceOk() (NullableFloat32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrice

`func (o *NewPriceAdjustment) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPrice

`func (o *NewPriceAdjustment) SetPrice(v NullableFloat32)`

SetPrice gets a reference to the given NullableFloat32 and assigns it to the Price field.

### SetPriceExplicitNull

`func (o *NewPriceAdjustment) SetPriceExplicitNull(b bool)`

SetPriceExplicitNull (un)sets Price to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Price value is set to nil even if false is passed
### GetReferenceId

`func (o *NewPriceAdjustment) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *NewPriceAdjustment) GetReferenceIdOk() (string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceId

`func (o *NewPriceAdjustment) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceId

`func (o *NewPriceAdjustment) SetReferenceId(v string)`

SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.

### GetCalculatedAt

`func (o *NewPriceAdjustment) GetCalculatedAt() time.Time`

GetCalculatedAt returns the CalculatedAt field if non-nil, zero value otherwise.

### GetCalculatedAtOk

`func (o *NewPriceAdjustment) GetCalculatedAtOk() (time.Time, bool)`

GetCalculatedAtOk returns a tuple with the CalculatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalculatedAt

`func (o *NewPriceAdjustment) HasCalculatedAt() bool`

HasCalculatedAt returns a boolean if a field has been set.

### SetCalculatedAt

`func (o *NewPriceAdjustment) SetCalculatedAt(v time.Time)`

SetCalculatedAt gets a reference to the given time.Time and assigns it to the CalculatedAt field.

### GetEffectiveFrom

`func (o *NewPriceAdjustment) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *NewPriceAdjustment) GetEffectiveFromOk() (time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffectiveFrom

`func (o *NewPriceAdjustment) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### SetEffectiveFrom

`func (o *NewPriceAdjustment) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom gets a reference to the given time.Time and assigns it to the EffectiveFrom field.

### GetEffectiveUntil

`func (o *NewPriceAdjustment) GetEffectiveUntil() time.Time`

GetEffectiveUntil returns the EffectiveUntil field if non-nil, zero value otherwise.

### GetEffectiveUntilOk

`func (o *NewPriceAdjustment) GetEffectiveUntilOk() (time.Time, bool)`

GetEffectiveUntilOk returns a tuple with the EffectiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffectiveUntil

`func (o *NewPriceAdjustment) HasEffectiveUntil() bool`

HasEffectiveUntil returns a boolean if a field has been set.

### SetEffectiveUntil

`func (o *NewPriceAdjustment) SetEffectiveUntil(v time.Time)`

SetEffectiveUntil gets a reference to the given time.Time and assigns it to the EffectiveUntil field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


