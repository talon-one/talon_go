# AwardDiscountAllItemsTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A target discriminator of type &#x60;allItems&#x60;. | 
**Prorated** | Pointer to **bool** | Whether to distribute the discount proportionally across the targeted items. | [optional] 

## Methods

### NewAwardDiscountAllItemsTarget

`func NewAwardDiscountAllItemsTarget(type_ string, ) *AwardDiscountAllItemsTarget`

NewAwardDiscountAllItemsTarget instantiates a new AwardDiscountAllItemsTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountAllItemsTargetWithDefaults

`func NewAwardDiscountAllItemsTargetWithDefaults() *AwardDiscountAllItemsTarget`

NewAwardDiscountAllItemsTargetWithDefaults instantiates a new AwardDiscountAllItemsTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwardDiscountAllItemsTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountAllItemsTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountAllItemsTarget) SetType(v string)`

SetType sets Type field to given value.


### GetProrated

`func (o *AwardDiscountAllItemsTarget) GetProrated() bool`

GetProrated returns the Prorated field if non-nil, zero value otherwise.

### GetProratedOk

`func (o *AwardDiscountAllItemsTarget) GetProratedOk() (*bool, bool)`

GetProratedOk returns a tuple with the Prorated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrated

`func (o *AwardDiscountAllItemsTarget) SetProrated(v bool)`

SetProrated sets Prorated field to given value.

### HasProrated

`func (o *AwardDiscountAllItemsTarget) HasProrated() bool`

HasProrated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


