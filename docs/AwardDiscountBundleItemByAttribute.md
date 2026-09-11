# AwardDiscountBundleItemByAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A bundle-item selector of type &#x60;byAttribute&#x60;. | 
**Attribute** | Pointer to **string** | A per-item attribute expression used to rank bundle items. | 
**Direction** | Pointer to **string** | Ranking direction. &#x60;highest&#x60; picks the item with the largest attribute value, &#x60;lowest&#x60; the smallest. | 

## Methods

### NewAwardDiscountBundleItemByAttribute

`func NewAwardDiscountBundleItemByAttribute(type_ string, attribute string, direction string, ) *AwardDiscountBundleItemByAttribute`

NewAwardDiscountBundleItemByAttribute instantiates a new AwardDiscountBundleItemByAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountBundleItemByAttributeWithDefaults

`func NewAwardDiscountBundleItemByAttributeWithDefaults() *AwardDiscountBundleItemByAttribute`

NewAwardDiscountBundleItemByAttributeWithDefaults instantiates a new AwardDiscountBundleItemByAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwardDiscountBundleItemByAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountBundleItemByAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountBundleItemByAttribute) SetType(v string)`

SetType sets Type field to given value.


### GetAttribute

`func (o *AwardDiscountBundleItemByAttribute) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AwardDiscountBundleItemByAttribute) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AwardDiscountBundleItemByAttribute) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetDirection

`func (o *AwardDiscountBundleItemByAttribute) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AwardDiscountBundleItemByAttribute) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AwardDiscountBundleItemByAttribute) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


