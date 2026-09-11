# AwardDiscountSelectorTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A target discriminator of type &#x60;selector&#x60;. | 
**Name** | Pointer to **string** | The name of the selector binding the discount targets. | 
**Prorated** | Pointer to **bool** | Whether to distribute the discount proportionally across the selected items. | [optional] 

## Methods

### NewAwardDiscountSelectorTarget

`func NewAwardDiscountSelectorTarget(type_ string, name string, ) *AwardDiscountSelectorTarget`

NewAwardDiscountSelectorTarget instantiates a new AwardDiscountSelectorTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountSelectorTargetWithDefaults

`func NewAwardDiscountSelectorTargetWithDefaults() *AwardDiscountSelectorTarget`

NewAwardDiscountSelectorTargetWithDefaults instantiates a new AwardDiscountSelectorTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwardDiscountSelectorTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountSelectorTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountSelectorTarget) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AwardDiscountSelectorTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwardDiscountSelectorTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwardDiscountSelectorTarget) SetName(v string)`

SetName sets Name field to given value.


### GetProrated

`func (o *AwardDiscountSelectorTarget) GetProrated() bool`

GetProrated returns the Prorated field if non-nil, zero value otherwise.

### GetProratedOk

`func (o *AwardDiscountSelectorTarget) GetProratedOk() (*bool, bool)`

GetProratedOk returns a tuple with the Prorated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrated

`func (o *AwardDiscountSelectorTarget) SetProrated(v bool)`

SetProrated sets Prorated field to given value.

### HasProrated

`func (o *AwardDiscountSelectorTarget) HasProrated() bool`

HasProrated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


