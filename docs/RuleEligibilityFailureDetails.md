# RuleEligibilityFailureDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureCode** | Pointer to **string** | A code identifying why the customer was not eligible for the rule in the current session. | 
**CouponID** | Pointer to **int64** | The ID of the coupon that was being evaluated when the rule failed.  | [optional] 
**CouponValue** | Pointer to **string** | The coupon code that was being evaluated when the rule failed.  | [optional] 
**ReferralID** | Pointer to **int64** | The ID of the referral that was being evaluated when the rule failed.  | [optional] 
**ReferralValue** | Pointer to **string** | The referral code that was being evaluated when the rule failed.  | [optional] 
**ConditionIndex** | Pointer to **int64** | The index of the condition that caused the rule to fail. | [optional] 
**EffectIndex** | Pointer to **int64** | The index of the effect that caused the rule to fail. | [optional] 
**Details** | Pointer to **string** | Additional details about the failure. | 

## Methods

### NewRuleEligibilityFailureDetails

`func NewRuleEligibilityFailureDetails(failureCode string, details string, ) *RuleEligibilityFailureDetails`

NewRuleEligibilityFailureDetails instantiates a new RuleEligibilityFailureDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleEligibilityFailureDetailsWithDefaults

`func NewRuleEligibilityFailureDetailsWithDefaults() *RuleEligibilityFailureDetails`

NewRuleEligibilityFailureDetailsWithDefaults instantiates a new RuleEligibilityFailureDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureCode

`func (o *RuleEligibilityFailureDetails) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *RuleEligibilityFailureDetails) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *RuleEligibilityFailureDetails) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.


### GetCouponID

`func (o *RuleEligibilityFailureDetails) GetCouponID() int64`

GetCouponID returns the CouponID field if non-nil, zero value otherwise.

### GetCouponIDOk

`func (o *RuleEligibilityFailureDetails) GetCouponIDOk() (*int64, bool)`

GetCouponIDOk returns a tuple with the CouponID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponID

`func (o *RuleEligibilityFailureDetails) SetCouponID(v int64)`

SetCouponID sets CouponID field to given value.

### HasCouponID

`func (o *RuleEligibilityFailureDetails) HasCouponID() bool`

HasCouponID returns a boolean if a field has been set.

### GetCouponValue

`func (o *RuleEligibilityFailureDetails) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *RuleEligibilityFailureDetails) GetCouponValueOk() (*string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponValue

`func (o *RuleEligibilityFailureDetails) SetCouponValue(v string)`

SetCouponValue sets CouponValue field to given value.

### HasCouponValue

`func (o *RuleEligibilityFailureDetails) HasCouponValue() bool`

HasCouponValue returns a boolean if a field has been set.

### GetReferralID

`func (o *RuleEligibilityFailureDetails) GetReferralID() int64`

GetReferralID returns the ReferralID field if non-nil, zero value otherwise.

### GetReferralIDOk

`func (o *RuleEligibilityFailureDetails) GetReferralIDOk() (*int64, bool)`

GetReferralIDOk returns a tuple with the ReferralID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralID

`func (o *RuleEligibilityFailureDetails) SetReferralID(v int64)`

SetReferralID sets ReferralID field to given value.

### HasReferralID

`func (o *RuleEligibilityFailureDetails) HasReferralID() bool`

HasReferralID returns a boolean if a field has been set.

### GetReferralValue

`func (o *RuleEligibilityFailureDetails) GetReferralValue() string`

GetReferralValue returns the ReferralValue field if non-nil, zero value otherwise.

### GetReferralValueOk

`func (o *RuleEligibilityFailureDetails) GetReferralValueOk() (*string, bool)`

GetReferralValueOk returns a tuple with the ReferralValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralValue

`func (o *RuleEligibilityFailureDetails) SetReferralValue(v string)`

SetReferralValue sets ReferralValue field to given value.

### HasReferralValue

`func (o *RuleEligibilityFailureDetails) HasReferralValue() bool`

HasReferralValue returns a boolean if a field has been set.

### GetConditionIndex

`func (o *RuleEligibilityFailureDetails) GetConditionIndex() int64`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *RuleEligibilityFailureDetails) GetConditionIndexOk() (*int64, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *RuleEligibilityFailureDetails) SetConditionIndex(v int64)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *RuleEligibilityFailureDetails) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### GetEffectIndex

`func (o *RuleEligibilityFailureDetails) GetEffectIndex() int64`

GetEffectIndex returns the EffectIndex field if non-nil, zero value otherwise.

### GetEffectIndexOk

`func (o *RuleEligibilityFailureDetails) GetEffectIndexOk() (*int64, bool)`

GetEffectIndexOk returns a tuple with the EffectIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectIndex

`func (o *RuleEligibilityFailureDetails) SetEffectIndex(v int64)`

SetEffectIndex sets EffectIndex field to given value.

### HasEffectIndex

`func (o *RuleEligibilityFailureDetails) HasEffectIndex() bool`

HasEffectIndex returns a boolean if a field has been set.

### GetDetails

`func (o *RuleEligibilityFailureDetails) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RuleEligibilityFailureDetails) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RuleEligibilityFailureDetails) SetDetails(v string)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


