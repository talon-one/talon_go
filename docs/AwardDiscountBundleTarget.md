# AwardDiscountBundleTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A target discriminator of type &#x60;bundle&#x60;. | 
**Name** | Pointer to **string** | Name of the bundle binding the discount targets. | 
**Item** | Pointer to [**map[string]interface{}**](.md) | Selects which slot inside a bundle a discount applies to. The &#x60;type&#x60; field picks the selection mode. | [optional] 
**Prorated** | Pointer to **bool** | Whether to distribute the discount proportionally across the bundle&#39;s items. | [optional] 

## Methods

### NewAwardDiscountBundleTarget

`func NewAwardDiscountBundleTarget(type_ string, name string, ) *AwardDiscountBundleTarget`

NewAwardDiscountBundleTarget instantiates a new AwardDiscountBundleTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountBundleTargetWithDefaults

`func NewAwardDiscountBundleTargetWithDefaults() *AwardDiscountBundleTarget`

NewAwardDiscountBundleTargetWithDefaults instantiates a new AwardDiscountBundleTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwardDiscountBundleTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountBundleTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountBundleTarget) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AwardDiscountBundleTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwardDiscountBundleTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwardDiscountBundleTarget) SetName(v string)`

SetName sets Name field to given value.


### GetItem

`func (o *AwardDiscountBundleTarget) GetItem() map[string]interface{}`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *AwardDiscountBundleTarget) GetItemOk() (*map[string]interface{}, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *AwardDiscountBundleTarget) SetItem(v map[string]interface{})`

SetItem sets Item field to given value.

### HasItem

`func (o *AwardDiscountBundleTarget) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetProrated

`func (o *AwardDiscountBundleTarget) GetProrated() bool`

GetProrated returns the Prorated field if non-nil, zero value otherwise.

### GetProratedOk

`func (o *AwardDiscountBundleTarget) GetProratedOk() (*bool, bool)`

GetProratedOk returns a tuple with the Prorated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrated

`func (o *AwardDiscountBundleTarget) SetProrated(v bool)`

SetProrated sets Prorated field to given value.

### HasProrated

`func (o *AwardDiscountBundleTarget) HasProrated() bool`

HasProrated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


