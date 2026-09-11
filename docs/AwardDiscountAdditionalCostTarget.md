# AwardDiscountAdditionalCostTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A target discriminator of type &#x60;additionalCost&#x60;. | 
**AdditionalCost** | Pointer to [**AdditionalCostReference**](AdditionalCostReference.md) |  | 
**Target** | Pointer to [**map[string]interface{}**](.md) | A subset of cart items whose additional cost the discount applies to. Cannot be another &#x60;additionalCost&#x60; target. | 

## Methods

### NewAwardDiscountAdditionalCostTarget

`func NewAwardDiscountAdditionalCostTarget(type_ string, additionalCost AdditionalCostReference, target map[string]interface{}, ) *AwardDiscountAdditionalCostTarget`

NewAwardDiscountAdditionalCostTarget instantiates a new AwardDiscountAdditionalCostTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountAdditionalCostTargetWithDefaults

`func NewAwardDiscountAdditionalCostTargetWithDefaults() *AwardDiscountAdditionalCostTarget`

NewAwardDiscountAdditionalCostTargetWithDefaults instantiates a new AwardDiscountAdditionalCostTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwardDiscountAdditionalCostTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountAdditionalCostTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountAdditionalCostTarget) SetType(v string)`

SetType sets Type field to given value.


### GetAdditionalCost

`func (o *AwardDiscountAdditionalCostTarget) GetAdditionalCost() AdditionalCostReference`

GetAdditionalCost returns the AdditionalCost field if non-nil, zero value otherwise.

### GetAdditionalCostOk

`func (o *AwardDiscountAdditionalCostTarget) GetAdditionalCostOk() (*AdditionalCostReference, bool)`

GetAdditionalCostOk returns a tuple with the AdditionalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCost

`func (o *AwardDiscountAdditionalCostTarget) SetAdditionalCost(v AdditionalCostReference)`

SetAdditionalCost sets AdditionalCost field to given value.


### GetTarget

`func (o *AwardDiscountAdditionalCostTarget) GetTarget() map[string]interface{}`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AwardDiscountAdditionalCostTarget) GetTargetOk() (*map[string]interface{}, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AwardDiscountAdditionalCostTarget) SetTarget(v map[string]interface{})`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


