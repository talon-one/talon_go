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
**ContextId** | Pointer to **string** | Identifier of the context of this price adjustment (e.g. summer sale). | [optional] 

## Methods

### NewNewPriceAdjustment

`func NewNewPriceAdjustment(priceType string, referenceId string, ) *NewPriceAdjustment`

NewNewPriceAdjustment instantiates a new NewPriceAdjustment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewPriceAdjustmentWithDefaults

`func NewNewPriceAdjustmentWithDefaults() *NewPriceAdjustment`

NewNewPriceAdjustmentWithDefaults instantiates a new NewPriceAdjustment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceType

`func (o *NewPriceAdjustment) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *NewPriceAdjustment) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *NewPriceAdjustment) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.


### GetPrice

`func (o *NewPriceAdjustment) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NewPriceAdjustment) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NewPriceAdjustment) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NewPriceAdjustment) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *NewPriceAdjustment) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *NewPriceAdjustment) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetReferenceId

`func (o *NewPriceAdjustment) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *NewPriceAdjustment) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *NewPriceAdjustment) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetCalculatedAt

`func (o *NewPriceAdjustment) GetCalculatedAt() time.Time`

GetCalculatedAt returns the CalculatedAt field if non-nil, zero value otherwise.

### GetCalculatedAtOk

`func (o *NewPriceAdjustment) GetCalculatedAtOk() (*time.Time, bool)`

GetCalculatedAtOk returns a tuple with the CalculatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedAt

`func (o *NewPriceAdjustment) SetCalculatedAt(v time.Time)`

SetCalculatedAt sets CalculatedAt field to given value.

### HasCalculatedAt

`func (o *NewPriceAdjustment) HasCalculatedAt() bool`

HasCalculatedAt returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *NewPriceAdjustment) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *NewPriceAdjustment) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *NewPriceAdjustment) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *NewPriceAdjustment) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveUntil

`func (o *NewPriceAdjustment) GetEffectiveUntil() time.Time`

GetEffectiveUntil returns the EffectiveUntil field if non-nil, zero value otherwise.

### GetEffectiveUntilOk

`func (o *NewPriceAdjustment) GetEffectiveUntilOk() (*time.Time, bool)`

GetEffectiveUntilOk returns a tuple with the EffectiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveUntil

`func (o *NewPriceAdjustment) SetEffectiveUntil(v time.Time)`

SetEffectiveUntil sets EffectiveUntil field to given value.

### HasEffectiveUntil

`func (o *NewPriceAdjustment) HasEffectiveUntil() bool`

HasEffectiveUntil returns a boolean if a field has been set.

### GetContextId

`func (o *NewPriceAdjustment) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *NewPriceAdjustment) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *NewPriceAdjustment) SetContextId(v string)`

SetContextId sets ContextId field to given value.

### HasContextId

`func (o *NewPriceAdjustment) HasContextId() bool`

HasContextId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


