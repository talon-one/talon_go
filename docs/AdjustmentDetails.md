# AdjustmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | Pointer to **string** | The reference identifier used during an &#x60;ADD_PRICE_ADJUSTMENT&#x60; action. | 
**SelectedPriceType** | Pointer to **string** | The selected price type for the SKU targeted by this effect. | 
**Value** | Pointer to **float32** | The value of the applied price adjustment. | 

## Methods

### NewAdjustmentDetails

`func NewAdjustmentDetails(referenceId string, selectedPriceType string, value float32, ) *AdjustmentDetails`

NewAdjustmentDetails instantiates a new AdjustmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentDetailsWithDefaults

`func NewAdjustmentDetailsWithDefaults() *AdjustmentDetails`

NewAdjustmentDetailsWithDefaults instantiates a new AdjustmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *AdjustmentDetails) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AdjustmentDetails) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AdjustmentDetails) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetSelectedPriceType

`func (o *AdjustmentDetails) GetSelectedPriceType() string`

GetSelectedPriceType returns the SelectedPriceType field if non-nil, zero value otherwise.

### GetSelectedPriceTypeOk

`func (o *AdjustmentDetails) GetSelectedPriceTypeOk() (*string, bool)`

GetSelectedPriceTypeOk returns a tuple with the SelectedPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPriceType

`func (o *AdjustmentDetails) SetSelectedPriceType(v string)`

SetSelectedPriceType sets SelectedPriceType field to given value.


### GetValue

`func (o *AdjustmentDetails) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AdjustmentDetails) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AdjustmentDetails) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


