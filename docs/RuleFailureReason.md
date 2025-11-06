# RuleFailureReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignID** | Pointer to **int64** | The ID of the campaign that contains the rule that failed. | 
**CampaignName** | Pointer to **string** | The name of the campaign that contains the rule that failed. | 
**RulesetID** | Pointer to **int64** | The ID of the ruleset that contains the rule that failed. | 
**CouponID** | Pointer to **int64** | The ID of the coupon that was being evaluated at the time of the rule failure. | [optional] 
**CouponValue** | Pointer to **string** | The code of the coupon that was being evaluated at the time of the rule failure. | [optional] 
**ReferralID** | Pointer to **int64** | The ID of the referral that was being evaluated at the time of the rule failure. | [optional] 
**ReferralValue** | Pointer to **string** | The code of the referral that was being evaluated at the time of the rule failure. | [optional] 
**RuleIndex** | Pointer to **int64** | The index of the rule that failed within the ruleset. | 
**RuleName** | Pointer to **string** | The name of the rule that failed within the ruleset. | 
**ConditionIndex** | Pointer to **int64** | The index of the condition that failed. | [optional] 
**EffectIndex** | Pointer to **int64** | The index of the effect that failed. | [optional] 
**Details** | Pointer to **string** | More details about the failure. | [optional] 
**EvaluationGroupID** | Pointer to **int64** | The ID of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation). | [optional] 
**EvaluationGroupMode** | Pointer to **string** | The evaluation mode of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign- | [optional] 

## Methods

### NewRuleFailureReason

`func NewRuleFailureReason(campaignID int64, campaignName string, rulesetID int64, ruleIndex int64, ruleName string, ) *RuleFailureReason`

NewRuleFailureReason instantiates a new RuleFailureReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleFailureReasonWithDefaults

`func NewRuleFailureReasonWithDefaults() *RuleFailureReason`

NewRuleFailureReasonWithDefaults instantiates a new RuleFailureReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignID

`func (o *RuleFailureReason) GetCampaignID() int64`

GetCampaignID returns the CampaignID field if non-nil, zero value otherwise.

### GetCampaignIDOk

`func (o *RuleFailureReason) GetCampaignIDOk() (*int64, bool)`

GetCampaignIDOk returns a tuple with the CampaignID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignID

`func (o *RuleFailureReason) SetCampaignID(v int64)`

SetCampaignID sets CampaignID field to given value.


### GetCampaignName

`func (o *RuleFailureReason) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *RuleFailureReason) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *RuleFailureReason) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.


### GetRulesetID

`func (o *RuleFailureReason) GetRulesetID() int64`

GetRulesetID returns the RulesetID field if non-nil, zero value otherwise.

### GetRulesetIDOk

`func (o *RuleFailureReason) GetRulesetIDOk() (*int64, bool)`

GetRulesetIDOk returns a tuple with the RulesetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetID

`func (o *RuleFailureReason) SetRulesetID(v int64)`

SetRulesetID sets RulesetID field to given value.


### GetCouponID

`func (o *RuleFailureReason) GetCouponID() int64`

GetCouponID returns the CouponID field if non-nil, zero value otherwise.

### GetCouponIDOk

`func (o *RuleFailureReason) GetCouponIDOk() (*int64, bool)`

GetCouponIDOk returns a tuple with the CouponID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponID

`func (o *RuleFailureReason) SetCouponID(v int64)`

SetCouponID sets CouponID field to given value.

### HasCouponID

`func (o *RuleFailureReason) HasCouponID() bool`

HasCouponID returns a boolean if a field has been set.

### GetCouponValue

`func (o *RuleFailureReason) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *RuleFailureReason) GetCouponValueOk() (*string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponValue

`func (o *RuleFailureReason) SetCouponValue(v string)`

SetCouponValue sets CouponValue field to given value.

### HasCouponValue

`func (o *RuleFailureReason) HasCouponValue() bool`

HasCouponValue returns a boolean if a field has been set.

### GetReferralID

`func (o *RuleFailureReason) GetReferralID() int64`

GetReferralID returns the ReferralID field if non-nil, zero value otherwise.

### GetReferralIDOk

`func (o *RuleFailureReason) GetReferralIDOk() (*int64, bool)`

GetReferralIDOk returns a tuple with the ReferralID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralID

`func (o *RuleFailureReason) SetReferralID(v int64)`

SetReferralID sets ReferralID field to given value.

### HasReferralID

`func (o *RuleFailureReason) HasReferralID() bool`

HasReferralID returns a boolean if a field has been set.

### GetReferralValue

`func (o *RuleFailureReason) GetReferralValue() string`

GetReferralValue returns the ReferralValue field if non-nil, zero value otherwise.

### GetReferralValueOk

`func (o *RuleFailureReason) GetReferralValueOk() (*string, bool)`

GetReferralValueOk returns a tuple with the ReferralValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralValue

`func (o *RuleFailureReason) SetReferralValue(v string)`

SetReferralValue sets ReferralValue field to given value.

### HasReferralValue

`func (o *RuleFailureReason) HasReferralValue() bool`

HasReferralValue returns a boolean if a field has been set.

### GetRuleIndex

`func (o *RuleFailureReason) GetRuleIndex() int64`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *RuleFailureReason) GetRuleIndexOk() (*int64, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIndex

`func (o *RuleFailureReason) SetRuleIndex(v int64)`

SetRuleIndex sets RuleIndex field to given value.


### GetRuleName

`func (o *RuleFailureReason) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *RuleFailureReason) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *RuleFailureReason) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetConditionIndex

`func (o *RuleFailureReason) GetConditionIndex() int64`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *RuleFailureReason) GetConditionIndexOk() (*int64, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *RuleFailureReason) SetConditionIndex(v int64)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *RuleFailureReason) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### GetEffectIndex

`func (o *RuleFailureReason) GetEffectIndex() int64`

GetEffectIndex returns the EffectIndex field if non-nil, zero value otherwise.

### GetEffectIndexOk

`func (o *RuleFailureReason) GetEffectIndexOk() (*int64, bool)`

GetEffectIndexOk returns a tuple with the EffectIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectIndex

`func (o *RuleFailureReason) SetEffectIndex(v int64)`

SetEffectIndex sets EffectIndex field to given value.

### HasEffectIndex

`func (o *RuleFailureReason) HasEffectIndex() bool`

HasEffectIndex returns a boolean if a field has been set.

### GetDetails

`func (o *RuleFailureReason) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RuleFailureReason) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RuleFailureReason) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RuleFailureReason) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEvaluationGroupID

`func (o *RuleFailureReason) GetEvaluationGroupID() int64`

GetEvaluationGroupID returns the EvaluationGroupID field if non-nil, zero value otherwise.

### GetEvaluationGroupIDOk

`func (o *RuleFailureReason) GetEvaluationGroupIDOk() (*int64, bool)`

GetEvaluationGroupIDOk returns a tuple with the EvaluationGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupID

`func (o *RuleFailureReason) SetEvaluationGroupID(v int64)`

SetEvaluationGroupID sets EvaluationGroupID field to given value.

### HasEvaluationGroupID

`func (o *RuleFailureReason) HasEvaluationGroupID() bool`

HasEvaluationGroupID returns a boolean if a field has been set.

### GetEvaluationGroupMode

`func (o *RuleFailureReason) GetEvaluationGroupMode() string`

GetEvaluationGroupMode returns the EvaluationGroupMode field if non-nil, zero value otherwise.

### GetEvaluationGroupModeOk

`func (o *RuleFailureReason) GetEvaluationGroupModeOk() (*string, bool)`

GetEvaluationGroupModeOk returns a tuple with the EvaluationGroupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupMode

`func (o *RuleFailureReason) SetEvaluationGroupMode(v string)`

SetEvaluationGroupMode sets EvaluationGroupMode field to given value.

### HasEvaluationGroupMode

`func (o *RuleFailureReason) HasEvaluationGroupMode() bool`

HasEvaluationGroupMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


