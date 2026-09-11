# AwardDiscountGlobalFilterTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A target discriminator of type &#x60;globalFilter&#x60;. | 
**Name** | Pointer to **string** | The name of the Application-level cart-item filter the discount targets. | 
**Prorated** | Pointer to **bool** | Whether to distribute the discount proportionally across the matched items. | [optional] 

## Methods

### NewAwardDiscountGlobalFilterTarget

`func NewAwardDiscountGlobalFilterTarget(type_ string, name string, ) *AwardDiscountGlobalFilterTarget`

NewAwardDiscountGlobalFilterTarget instantiates a new AwardDiscountGlobalFilterTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountGlobalFilterTargetWithDefaults

`func NewAwardDiscountGlobalFilterTargetWithDefaults() *AwardDiscountGlobalFilterTarget`

NewAwardDiscountGlobalFilterTargetWithDefaults instantiates a new AwardDiscountGlobalFilterTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwardDiscountGlobalFilterTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountGlobalFilterTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountGlobalFilterTarget) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AwardDiscountGlobalFilterTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwardDiscountGlobalFilterTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwardDiscountGlobalFilterTarget) SetName(v string)`

SetName sets Name field to given value.


### GetProrated

`func (o *AwardDiscountGlobalFilterTarget) GetProrated() bool`

GetProrated returns the Prorated field if non-nil, zero value otherwise.

### GetProratedOk

`func (o *AwardDiscountGlobalFilterTarget) GetProratedOk() (*bool, bool)`

GetProratedOk returns a tuple with the Prorated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrated

`func (o *AwardDiscountGlobalFilterTarget) SetProrated(v bool)`

SetProrated sets Prorated field to given value.

### HasProrated

`func (o *AwardDiscountGlobalFilterTarget) HasProrated() bool`

HasProrated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


