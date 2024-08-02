# CouponCreationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with coupons. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**NumberOfCoupons** | Pointer to **int32** | The number of new coupon codes to generate for the campaign. | 

## Methods

### GetAttributes

`func (o *CouponCreationJob) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponCreationJob) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CouponCreationJob) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CouponCreationJob) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetCouponSettings

`func (o *CouponCreationJob) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *CouponCreationJob) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *CouponCreationJob) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *CouponCreationJob) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetNumberOfCoupons

`func (o *CouponCreationJob) GetNumberOfCoupons() int32`

GetNumberOfCoupons returns the NumberOfCoupons field if non-nil, zero value otherwise.

### GetNumberOfCouponsOk

`func (o *CouponCreationJob) GetNumberOfCouponsOk() (int32, bool)`

GetNumberOfCouponsOk returns a tuple with the NumberOfCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberOfCoupons

`func (o *CouponCreationJob) HasNumberOfCoupons() bool`

HasNumberOfCoupons returns a boolean if a field has been set.

### SetNumberOfCoupons

`func (o *CouponCreationJob) SetNumberOfCoupons(v int32)`

SetNumberOfCoupons gets a reference to the given int32 and assigns it to the NumberOfCoupons field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


