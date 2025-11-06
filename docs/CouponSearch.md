# CouponSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Properties to match against a coupon. All provided attributes will be exactly matched against attributes. | 

## Methods

### NewCouponSearch

`func NewCouponSearch(attributes map[string]interface{}, ) *CouponSearch`

NewCouponSearch instantiates a new CouponSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponSearchWithDefaults

`func NewCouponSearchWithDefaults() *CouponSearch`

NewCouponSearchWithDefaults instantiates a new CouponSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CouponSearch) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponSearch) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponSearch) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


