# LimitCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | 
**ApplicationId** | Pointer to **int32** | The ID of the Application that owns this entity. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
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

### GetCampaignId

`func (o *LimitCounter) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *LimitCounter) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *LimitCounter) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *LimitCounter) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetApplicationId

`func (o *LimitCounter) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *LimitCounter) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *LimitCounter) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *LimitCounter) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetAccountId

`func (o *LimitCounter) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LimitCounter) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *LimitCounter) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *LimitCounter) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetId

`func (o *LimitCounter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitCounter) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LimitCounter) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LimitCounter) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetAction

`func (o *LimitCounter) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LimitCounter) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *LimitCounter) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *LimitCounter) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetProfileId

`func (o *LimitCounter) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *LimitCounter) GetProfileIdOk() (int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *LimitCounter) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *LimitCounter) SetProfileId(v int32)`

SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.

### GetProfileIntegrationId

`func (o *LimitCounter) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *LimitCounter) GetProfileIntegrationIdOk() (string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileIntegrationId

`func (o *LimitCounter) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### SetProfileIntegrationId

`func (o *LimitCounter) SetProfileIntegrationId(v string)`

SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.

### GetCouponId

`func (o *LimitCounter) GetCouponId() int32`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *LimitCounter) GetCouponIdOk() (int32, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponId

`func (o *LimitCounter) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### SetCouponId

`func (o *LimitCounter) SetCouponId(v int32)`

SetCouponId gets a reference to the given int32 and assigns it to the CouponId field.

### GetCouponValue

`func (o *LimitCounter) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *LimitCounter) GetCouponValueOk() (string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponValue

`func (o *LimitCounter) HasCouponValue() bool`

HasCouponValue returns a boolean if a field has been set.

### SetCouponValue

`func (o *LimitCounter) SetCouponValue(v string)`

SetCouponValue gets a reference to the given string and assigns it to the CouponValue field.

### GetReferralId

`func (o *LimitCounter) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *LimitCounter) GetReferralIdOk() (int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralId

`func (o *LimitCounter) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### SetReferralId

`func (o *LimitCounter) SetReferralId(v int32)`

SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.

### GetReferralValue

`func (o *LimitCounter) GetReferralValue() string`

GetReferralValue returns the ReferralValue field if non-nil, zero value otherwise.

### GetReferralValueOk

`func (o *LimitCounter) GetReferralValueOk() (string, bool)`

GetReferralValueOk returns a tuple with the ReferralValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralValue

`func (o *LimitCounter) HasReferralValue() bool`

HasReferralValue returns a boolean if a field has been set.

### SetReferralValue

`func (o *LimitCounter) SetReferralValue(v string)`

SetReferralValue gets a reference to the given string and assigns it to the ReferralValue field.

### GetIdentifier

`func (o *LimitCounter) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LimitCounter) GetIdentifierOk() (string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentifier

`func (o *LimitCounter) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifier

`func (o *LimitCounter) SetIdentifier(v string)`

SetIdentifier gets a reference to the given string and assigns it to the Identifier field.

### GetPeriod

`func (o *LimitCounter) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *LimitCounter) GetPeriodOk() (string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriod

`func (o *LimitCounter) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriod

`func (o *LimitCounter) SetPeriod(v string)`

SetPeriod gets a reference to the given string and assigns it to the Period field.

### GetLimit

`func (o *LimitCounter) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LimitCounter) GetLimitOk() (float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *LimitCounter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *LimitCounter) SetLimit(v float32)`

SetLimit gets a reference to the given float32 and assigns it to the Limit field.

### GetCounter

`func (o *LimitCounter) GetCounter() float32`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *LimitCounter) GetCounterOk() (float32, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCounter

`func (o *LimitCounter) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### SetCounter

`func (o *LimitCounter) SetCounter(v float32)`

SetCounter gets a reference to the given float32 and assigns it to the Counter field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


