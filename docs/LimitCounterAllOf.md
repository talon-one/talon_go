# LimitCounterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Action** | Pointer to **string** | The limitable action of the limit counter. | 
**ProfileId** | Pointer to **int32** | The profile ID for which this limit counter is used. | [optional] 
**ProfileIntegrationId** | Pointer to **string** | The profile integration ID for which this limit counter is used. | [optional] 
**CouponId** | Pointer to **int32** | The internal coupon ID for which this limit counter is used. | [optional] 
**CouponValue** | Pointer to **string** | The coupon value for which this limit counter is used. | [optional] 
**ReferralId** | Pointer to **int32** | The referral ID for which this limit counter is used. | [optional] 
**ReferralValue** | Pointer to **string** | The referral value for which this limit counter is used. | [optional] 
**Identifier** | Pointer to **string** | The arbitrary identifier for which this limit counter is used. | [optional] 
**Period** | Pointer to **string** | The time period for which this limit counter is used. | [optional] 
**Limit** | Pointer to **float32** | The highest possible value for this limit counter. | 
**Counter** | Pointer to **float32** | The current value for this limit counter. | 

## Methods

### GetId

`func (o *LimitCounterAllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitCounterAllOf) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LimitCounterAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LimitCounterAllOf) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetAction

`func (o *LimitCounterAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LimitCounterAllOf) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *LimitCounterAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *LimitCounterAllOf) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetProfileId

`func (o *LimitCounterAllOf) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *LimitCounterAllOf) GetProfileIdOk() (int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *LimitCounterAllOf) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *LimitCounterAllOf) SetProfileId(v int32)`

SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.

### GetProfileIntegrationId

`func (o *LimitCounterAllOf) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *LimitCounterAllOf) GetProfileIntegrationIdOk() (string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileIntegrationId

`func (o *LimitCounterAllOf) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### SetProfileIntegrationId

`func (o *LimitCounterAllOf) SetProfileIntegrationId(v string)`

SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.

### GetCouponId

`func (o *LimitCounterAllOf) GetCouponId() int32`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *LimitCounterAllOf) GetCouponIdOk() (int32, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponId

`func (o *LimitCounterAllOf) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponId

`func (o *LimitCounterAllOf) SetCouponId(v int32)`

SetCouponId gets a reference to the given int32 and assigns it to the CouponId field.

### GetCouponValue

`func (o *LimitCounterAllOf) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *LimitCounterAllOf) GetCouponValueOk() (string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponValue

`func (o *LimitCounterAllOf) HasCouponValue() bool`

HasCouponValue returns a boolean if a field has been set.

### SetCouponValue

`func (o *LimitCounterAllOf) SetCouponValue(v string)`

SetCouponValue gets a reference to the given string and assigns it to the CouponValue field.

### GetReferralId

`func (o *LimitCounterAllOf) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *LimitCounterAllOf) GetReferralIdOk() (int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralId

`func (o *LimitCounterAllOf) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### SetReferralId

`func (o *LimitCounterAllOf) SetReferralId(v int32)`

SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.

### GetReferralValue

`func (o *LimitCounterAllOf) GetReferralValue() string`

GetReferralValue returns the ReferralValue field if non-nil, zero value otherwise.

### GetReferralValueOk

`func (o *LimitCounterAllOf) GetReferralValueOk() (string, bool)`

GetReferralValueOk returns a tuple with the ReferralValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralValue

`func (o *LimitCounterAllOf) HasReferralValue() bool`

HasReferralValue returns a boolean if a field has been set.

### SetReferralValue

`func (o *LimitCounterAllOf) SetReferralValue(v string)`

SetReferralValue gets a reference to the given string and assigns it to the ReferralValue field.

### GetIdentifier

`func (o *LimitCounterAllOf) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LimitCounterAllOf) GetIdentifierOk() (string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentifier

`func (o *LimitCounterAllOf) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifier

`func (o *LimitCounterAllOf) SetIdentifier(v string)`

SetIdentifier gets a reference to the given string and assigns it to the Identifier field.

### GetPeriod

`func (o *LimitCounterAllOf) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LimitCounterAllOf) GetPeriodOk() (string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriod

`func (o *LimitCounterAllOf) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriod

`func (o *LimitCounterAllOf) SetPeriod(v string)`

SetPeriod gets a reference to the given string and assigns it to the Period field.

### GetLimit

`func (o *LimitCounterAllOf) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LimitCounterAllOf) GetLimitOk() (float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *LimitCounterAllOf) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *LimitCounterAllOf) SetLimit(v float32)`

SetLimit gets a reference to the given float32 and assigns it to the Limit field.

### GetCounter

`func (o *LimitCounterAllOf) GetCounter() float32`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *LimitCounterAllOf) GetCounterOk() (float32, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCounter

`func (o *LimitCounterAllOf) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### SetCounter

`func (o *LimitCounterAllOf) SetCounter(v float32)`

SetCounter gets a reference to the given float32 and assigns it to the Counter field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


