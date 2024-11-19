# RuleFailureReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignID** | Pointer to **int32** | The ID of the campaign that contains the rule that failed. | 
**CampaignName** | Pointer to **string** | The name of the campaign that contains the rule that failed. | 
**RulesetID** | Pointer to **int32** | The ID of the ruleset that contains the rule that failed. | 
**CouponID** | Pointer to **int32** | The ID of the coupon that was being evaluated at the time of the rule failure. | [optional] 
**CouponValue** | Pointer to **string** | The code of the coupon that was being evaluated at the time of the rule failure. | [optional] 
**ReferralID** | Pointer to **int32** | The ID of the referral that was being evaluated at the time of the rule failure. | [optional] 
**ReferralValue** | Pointer to **string** | The code of the referral that was being evaluated at the time of the rule failure. | [optional] 
**RuleIndex** | Pointer to **int32** | The index of the rule that failed within the ruleset. | 
**RuleName** | Pointer to **string** | The name of the rule that failed within the ruleset. | 
**ConditionIndex** | Pointer to **int32** | The index of the condition that failed. | [optional] 
**EffectIndex** | Pointer to **int32** | The index of the effect that failed. | [optional] 
**Details** | Pointer to **string** | More details about the failure. | [optional] 
**EvaluationGroupID** | Pointer to **int32** | The ID of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation). | [optional] 
**EvaluationGroupMode** | Pointer to **string** | The evaluation mode of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign- | [optional] 

## Methods

### GetCampaignID

`func (o *RuleFailureReason) GetCampaignID() int32`

GetCampaignID returns the CampaignID field if non-nil, zero value otherwise.

### GetCampaignIDOk

`func (o *RuleFailureReason) GetCampaignIDOk() (int32, bool)`

GetCampaignIDOk returns a tuple with the CampaignID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignID

`func (o *RuleFailureReason) HasCampaignID() bool`

HasCampaignID returns a boolean if a field has been set.

### SetCampaignID

`func (o *RuleFailureReason) SetCampaignID(v int32)`

SetCampaignID gets a reference to the given int32 and assigns it to the CampaignID field.

### GetCampaignName

`func (o *RuleFailureReason) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *RuleFailureReason) GetCampaignNameOk() (string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignName

`func (o *RuleFailureReason) HasCampaignName() bool`

HasCampaignName returns a boolean if a field has been set.

### SetCampaignName

`func (o *RuleFailureReason) SetCampaignName(v string)`

SetCampaignName gets a reference to the given string and assigns it to the CampaignName field.

### GetRulesetID

`func (o *RuleFailureReason) GetRulesetID() int32`

GetRulesetID returns the RulesetID field if non-nil, zero value otherwise.

### GetRulesetIDOk

`func (o *RuleFailureReason) GetRulesetIDOk() (int32, bool)`

GetRulesetIDOk returns a tuple with the RulesetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRulesetID

`func (o *RuleFailureReason) HasRulesetID() bool`

HasRulesetID returns a boolean if a field has been set.

### SetRulesetID

`func (o *RuleFailureReason) SetRulesetID(v int32)`

SetRulesetID gets a reference to the given int32 and assigns it to the RulesetID field.

### GetCouponID

`func (o *RuleFailureReason) GetCouponID() int32`

GetCouponID returns the CouponID field if non-nil, zero value otherwise.

### GetCouponIDOk

`func (o *RuleFailureReason) GetCouponIDOk() (int32, bool)`

GetCouponIDOk returns a tuple with the CouponID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponID

`func (o *RuleFailureReason) HasCouponID() bool`

HasCouponID returns a boolean if a field has been set.

### SetCouponID

`func (o *RuleFailureReason) SetCouponID(v int32)`

SetCouponID gets a reference to the given int32 and assigns it to the CouponID field.

### GetCouponValue

`func (o *RuleFailureReason) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *RuleFailureReason) GetCouponValueOk() (string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponValue

`func (o *RuleFailureReason) HasCouponValue() bool`

HasCouponValue returns a boolean if a field has been set.

### SetCouponValue

`func (o *RuleFailureReason) SetCouponValue(v string)`

SetCouponValue gets a reference to the given string and assigns it to the CouponValue field.

### GetReferralID

`func (o *RuleFailureReason) GetReferralID() int32`

GetReferralID returns the ReferralID field if non-nil, zero value otherwise.

### GetReferralIDOk

`func (o *RuleFailureReason) GetReferralIDOk() (int32, bool)`

GetReferralIDOk returns a tuple with the ReferralID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralID

`func (o *RuleFailureReason) HasReferralID() bool`

HasReferralID returns a boolean if a field has been set.

### SetReferralID

`func (o *RuleFailureReason) SetReferralID(v int32)`

SetReferralID gets a reference to the given int32 and assigns it to the ReferralID field.

### GetReferralValue

`func (o *RuleFailureReason) GetReferralValue() string`

GetReferralValue returns the ReferralValue field if non-nil, zero value otherwise.

### GetReferralValueOk

`func (o *RuleFailureReason) GetReferralValueOk() (string, bool)`

GetReferralValueOk returns a tuple with the ReferralValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralValue

`func (o *RuleFailureReason) HasReferralValue() bool`

HasReferralValue returns a boolean if a field has been set.

### SetReferralValue

`func (o *RuleFailureReason) SetReferralValue(v string)`

SetReferralValue gets a reference to the given string and assigns it to the ReferralValue field.

### GetRuleIndex

`func (o *RuleFailureReason) GetRuleIndex() int32`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *RuleFailureReason) GetRuleIndexOk() (int32, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleIndex

`func (o *RuleFailureReason) HasRuleIndex() bool`

HasRuleIndex returns a boolean if a field has been set.

### SetRuleIndex

`func (o *RuleFailureReason) SetRuleIndex(v int32)`

SetRuleIndex gets a reference to the given int32 and assigns it to the RuleIndex field.

### GetRuleName

`func (o *RuleFailureReason) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *RuleFailureReason) GetRuleNameOk() (string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleName

`func (o *RuleFailureReason) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleName

`func (o *RuleFailureReason) SetRuleName(v string)`

SetRuleName gets a reference to the given string and assigns it to the RuleName field.

### GetConditionIndex

`func (o *RuleFailureReason) GetConditionIndex() int32`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *RuleFailureReason) GetConditionIndexOk() (int32, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditionIndex

`func (o *RuleFailureReason) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### SetConditionIndex

`func (o *RuleFailureReason) SetConditionIndex(v int32)`

SetConditionIndex gets a reference to the given int32 and assigns it to the ConditionIndex field.

### GetEffectIndex

`func (o *RuleFailureReason) GetEffectIndex() int32`

GetEffectIndex returns the EffectIndex field if non-nil, zero value otherwise.

### GetEffectIndexOk

`func (o *RuleFailureReason) GetEffectIndexOk() (int32, bool)`

GetEffectIndexOk returns a tuple with the EffectIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffectIndex

`func (o *RuleFailureReason) HasEffectIndex() bool`

HasEffectIndex returns a boolean if a field has been set.

### SetEffectIndex

`func (o *RuleFailureReason) SetEffectIndex(v int32)`

SetEffectIndex gets a reference to the given int32 and assigns it to the EffectIndex field.

### GetDetails

`func (o *RuleFailureReason) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RuleFailureReason) GetDetailsOk() (string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *RuleFailureReason) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *RuleFailureReason) SetDetails(v string)`

SetDetails gets a reference to the given string and assigns it to the Details field.

### GetEvaluationGroupID

`func (o *RuleFailureReason) GetEvaluationGroupID() int32`

GetEvaluationGroupID returns the EvaluationGroupID field if non-nil, zero value otherwise.

### GetEvaluationGroupIDOk

`func (o *RuleFailureReason) GetEvaluationGroupIDOk() (int32, bool)`

GetEvaluationGroupIDOk returns a tuple with the EvaluationGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationGroupID

`func (o *RuleFailureReason) HasEvaluationGroupID() bool`

HasEvaluationGroupID returns a boolean if a field has been set.

### SetEvaluationGroupID

`func (o *RuleFailureReason) SetEvaluationGroupID(v int32)`

SetEvaluationGroupID gets a reference to the given int32 and assigns it to the EvaluationGroupID field.

### GetEvaluationGroupMode

`func (o *RuleFailureReason) GetEvaluationGroupMode() string`

GetEvaluationGroupMode returns the EvaluationGroupMode field if non-nil, zero value otherwise.

### GetEvaluationGroupModeOk

`func (o *RuleFailureReason) GetEvaluationGroupModeOk() (string, bool)`

GetEvaluationGroupModeOk returns a tuple with the EvaluationGroupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationGroupMode

`func (o *RuleFailureReason) HasEvaluationGroupMode() bool`

HasEvaluationGroupMode returns a boolean if a field has been set.

### SetEvaluationGroupMode

`func (o *RuleFailureReason) SetEvaluationGroupMode(v string)`

SetEvaluationGroupMode gets a reference to the given string and assigns it to the EvaluationGroupMode field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


