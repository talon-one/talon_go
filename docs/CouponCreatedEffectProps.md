# CouponCreatedEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The coupon code that was created. | 
**ProfileId** | Pointer to **string** | The integration identifier of the customer for whom this coupon was created. | 

## Methods

### NewCouponCreatedEffectProps

`func NewCouponCreatedEffectProps(value string, profileId string, ) *CouponCreatedEffectProps`

NewCouponCreatedEffectProps instantiates a new CouponCreatedEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCreatedEffectPropsWithDefaults

`func NewCouponCreatedEffectPropsWithDefaults() *CouponCreatedEffectProps`

NewCouponCreatedEffectPropsWithDefaults instantiates a new CouponCreatedEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CouponCreatedEffectProps) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CouponCreatedEffectProps) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CouponCreatedEffectProps) SetValue(v string)`

SetValue sets Value field to given value.


### GetProfileId

`func (o *CouponCreatedEffectProps) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CouponCreatedEffectProps) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CouponCreatedEffectProps) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


