# AwardDiscountBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Name** | Pointer to **string** | The human-readable label attached to the discount. | 
**Value** | Pointer to [**map[string]interface{}**](.md) | Discount amount. Either a numeric scalar or a &#x60;{{expression}}&#x60; string that resolves to a number at evaluation time. | 
**Partial** | Pointer to **bool** | Whether to apply a partial discount when the requested value exceeds the configured budget. | 
**Target** | Pointer to [**map[string]interface{}**](.md) | Identifies the scope a discount applies to. The &#x60;type&#x60; field selects the concrete target variant. | 

## Methods

### NewAwardDiscountBlock

`func NewAwardDiscountBlock(type_ string, name string, value map[string]interface{}, partial bool, target map[string]interface{}, ) *AwardDiscountBlock`

NewAwardDiscountBlock instantiates a new AwardDiscountBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountBlockWithDefaults

`func NewAwardDiscountBlockWithDefaults() *AwardDiscountBlock`

NewAwardDiscountBlockWithDefaults instantiates a new AwardDiscountBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwardDiscountBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwardDiscountBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwardDiscountBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwardDiscountBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AwardDiscountBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *AwardDiscountBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwardDiscountBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwardDiscountBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwardDiscountBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *AwardDiscountBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwardDiscountBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwardDiscountBlock) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AwardDiscountBlock) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AwardDiscountBlock) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AwardDiscountBlock) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetPartial

`func (o *AwardDiscountBlock) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *AwardDiscountBlock) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *AwardDiscountBlock) SetPartial(v bool)`

SetPartial sets Partial field to given value.


### GetTarget

`func (o *AwardDiscountBlock) GetTarget() map[string]interface{}`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AwardDiscountBlock) GetTargetOk() (*map[string]interface{}, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AwardDiscountBlock) SetTarget(v map[string]interface{})`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


