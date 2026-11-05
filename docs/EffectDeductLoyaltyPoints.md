# EffectDeductLoyaltyPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentId** | Pointer to **int64** | The ID of the experiment that campaign belongs to. | [optional] 
**CampaignId** | Pointer to **int64** | The ID of the campaign that triggered this effect. | 
**RulesetId** | Pointer to **int64** | The ID of the ruleset that was active in the campaign when this effect was triggered. | 
**RuleIndex** | Pointer to **int64** | The position of the rule that triggered this effect within the ruleset. | 
**RuleName** | Pointer to **string** | The name of the rule that triggered this effect. | 
**EffectType** | Pointer to **string** | An effect discriminator of type &#x60;deductLoyaltyPoints&#x60;. | 
**TriggeredByCoupon** | Pointer to **int64** | The ID of the coupon that was being evaluated when this effect was triggered. | [optional] 
**TriggeredForCatalogItem** | Pointer to **int64** | The ID of the catalog item that was being evaluated when this effect was triggered. | [optional] 
**ConditionIndex** | Pointer to **int64** | The index of the condition that was triggered. | [optional] 
**EvaluationGroupID** | Pointer to **int64** | The ID of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation). | [optional] 
**EvaluationGroupMode** | Pointer to **string** | The evaluation mode of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation). | [optional] 
**CampaignRevisionId** | Pointer to **int64** | The revision ID of the campaign that was used when triggering the effect. | [optional] 
**CampaignRevisionVersionId** | Pointer to **int64** | The revision version ID of the campaign that was used when triggering the effect. | [optional] 
**SelectedPriceType** | Pointer to **string** | The selected price type for the SKU targeted by this effect. | [optional] 
**SelectedPrice** | Pointer to **float32** | The value of the selected price type to apply to the SKU targeted by this effect, before any discounts are applied. | [optional] 
**AdjustmentReferenceId** | Pointer to **string** | The reference identifier of the selected price adjustment for this SKU. This is only returned if the &#x60;selectedPrice&#x60; resulted from a price adjustment. | [optional] 
**RewardId** | Pointer to **int64** | The ID of the reward that was being evaluated when this effect was triggered. | [optional] 
**RewardIntegrationId** | Pointer to **string** | The integration ID of the specific customer reward whose usage produced this effect. | [optional] 
**Props** | Pointer to [**DeductLoyaltyPointsEffectProps**](DeductLoyaltyPointsEffectProps.md) |  | 

## Methods

### NewEffectDeductLoyaltyPoints

`func NewEffectDeductLoyaltyPoints(campaignId int64, rulesetId int64, ruleIndex int64, ruleName string, effectType string, props DeductLoyaltyPointsEffectProps, ) *EffectDeductLoyaltyPoints`

NewEffectDeductLoyaltyPoints instantiates a new EffectDeductLoyaltyPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectDeductLoyaltyPointsWithDefaults

`func NewEffectDeductLoyaltyPointsWithDefaults() *EffectDeductLoyaltyPoints`

NewEffectDeductLoyaltyPointsWithDefaults instantiates a new EffectDeductLoyaltyPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentId

`func (o *EffectDeductLoyaltyPoints) GetExperimentId() int64`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *EffectDeductLoyaltyPoints) GetExperimentIdOk() (*int64, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *EffectDeductLoyaltyPoints) SetExperimentId(v int64)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *EffectDeductLoyaltyPoints) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### GetCampaignId

`func (o *EffectDeductLoyaltyPoints) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *EffectDeductLoyaltyPoints) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *EffectDeductLoyaltyPoints) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetRulesetId

`func (o *EffectDeductLoyaltyPoints) GetRulesetId() int64`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *EffectDeductLoyaltyPoints) GetRulesetIdOk() (*int64, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *EffectDeductLoyaltyPoints) SetRulesetId(v int64)`

SetRulesetId sets RulesetId field to given value.


### GetRuleIndex

`func (o *EffectDeductLoyaltyPoints) GetRuleIndex() int64`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *EffectDeductLoyaltyPoints) GetRuleIndexOk() (*int64, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIndex

`func (o *EffectDeductLoyaltyPoints) SetRuleIndex(v int64)`

SetRuleIndex sets RuleIndex field to given value.


### GetRuleName

`func (o *EffectDeductLoyaltyPoints) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *EffectDeductLoyaltyPoints) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *EffectDeductLoyaltyPoints) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetEffectType

`func (o *EffectDeductLoyaltyPoints) GetEffectType() string`

GetEffectType returns the EffectType field if non-nil, zero value otherwise.

### GetEffectTypeOk

`func (o *EffectDeductLoyaltyPoints) GetEffectTypeOk() (*string, bool)`

GetEffectTypeOk returns a tuple with the EffectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectType

`func (o *EffectDeductLoyaltyPoints) SetEffectType(v string)`

SetEffectType sets EffectType field to given value.


### GetTriggeredByCoupon

`func (o *EffectDeductLoyaltyPoints) GetTriggeredByCoupon() int64`

GetTriggeredByCoupon returns the TriggeredByCoupon field if non-nil, zero value otherwise.

### GetTriggeredByCouponOk

`func (o *EffectDeductLoyaltyPoints) GetTriggeredByCouponOk() (*int64, bool)`

GetTriggeredByCouponOk returns a tuple with the TriggeredByCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredByCoupon

`func (o *EffectDeductLoyaltyPoints) SetTriggeredByCoupon(v int64)`

SetTriggeredByCoupon sets TriggeredByCoupon field to given value.

### HasTriggeredByCoupon

`func (o *EffectDeductLoyaltyPoints) HasTriggeredByCoupon() bool`

HasTriggeredByCoupon returns a boolean if a field has been set.

### GetTriggeredForCatalogItem

`func (o *EffectDeductLoyaltyPoints) GetTriggeredForCatalogItem() int64`

GetTriggeredForCatalogItem returns the TriggeredForCatalogItem field if non-nil, zero value otherwise.

### GetTriggeredForCatalogItemOk

`func (o *EffectDeductLoyaltyPoints) GetTriggeredForCatalogItemOk() (*int64, bool)`

GetTriggeredForCatalogItemOk returns a tuple with the TriggeredForCatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredForCatalogItem

`func (o *EffectDeductLoyaltyPoints) SetTriggeredForCatalogItem(v int64)`

SetTriggeredForCatalogItem sets TriggeredForCatalogItem field to given value.

### HasTriggeredForCatalogItem

`func (o *EffectDeductLoyaltyPoints) HasTriggeredForCatalogItem() bool`

HasTriggeredForCatalogItem returns a boolean if a field has been set.

### GetConditionIndex

`func (o *EffectDeductLoyaltyPoints) GetConditionIndex() int64`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *EffectDeductLoyaltyPoints) GetConditionIndexOk() (*int64, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *EffectDeductLoyaltyPoints) SetConditionIndex(v int64)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *EffectDeductLoyaltyPoints) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### GetEvaluationGroupID

`func (o *EffectDeductLoyaltyPoints) GetEvaluationGroupID() int64`

GetEvaluationGroupID returns the EvaluationGroupID field if non-nil, zero value otherwise.

### GetEvaluationGroupIDOk

`func (o *EffectDeductLoyaltyPoints) GetEvaluationGroupIDOk() (*int64, bool)`

GetEvaluationGroupIDOk returns a tuple with the EvaluationGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupID

`func (o *EffectDeductLoyaltyPoints) SetEvaluationGroupID(v int64)`

SetEvaluationGroupID sets EvaluationGroupID field to given value.

### HasEvaluationGroupID

`func (o *EffectDeductLoyaltyPoints) HasEvaluationGroupID() bool`

HasEvaluationGroupID returns a boolean if a field has been set.

### GetEvaluationGroupMode

`func (o *EffectDeductLoyaltyPoints) GetEvaluationGroupMode() string`

GetEvaluationGroupMode returns the EvaluationGroupMode field if non-nil, zero value otherwise.

### GetEvaluationGroupModeOk

`func (o *EffectDeductLoyaltyPoints) GetEvaluationGroupModeOk() (*string, bool)`

GetEvaluationGroupModeOk returns a tuple with the EvaluationGroupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupMode

`func (o *EffectDeductLoyaltyPoints) SetEvaluationGroupMode(v string)`

SetEvaluationGroupMode sets EvaluationGroupMode field to given value.

### HasEvaluationGroupMode

`func (o *EffectDeductLoyaltyPoints) HasEvaluationGroupMode() bool`

HasEvaluationGroupMode returns a boolean if a field has been set.

### GetCampaignRevisionId

`func (o *EffectDeductLoyaltyPoints) GetCampaignRevisionId() int64`

GetCampaignRevisionId returns the CampaignRevisionId field if non-nil, zero value otherwise.

### GetCampaignRevisionIdOk

`func (o *EffectDeductLoyaltyPoints) GetCampaignRevisionIdOk() (*int64, bool)`

GetCampaignRevisionIdOk returns a tuple with the CampaignRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRevisionId

`func (o *EffectDeductLoyaltyPoints) SetCampaignRevisionId(v int64)`

SetCampaignRevisionId sets CampaignRevisionId field to given value.

### HasCampaignRevisionId

`func (o *EffectDeductLoyaltyPoints) HasCampaignRevisionId() bool`

HasCampaignRevisionId returns a boolean if a field has been set.

### GetCampaignRevisionVersionId

`func (o *EffectDeductLoyaltyPoints) GetCampaignRevisionVersionId() int64`

GetCampaignRevisionVersionId returns the CampaignRevisionVersionId field if non-nil, zero value otherwise.

### GetCampaignRevisionVersionIdOk

`func (o *EffectDeductLoyaltyPoints) GetCampaignRevisionVersionIdOk() (*int64, bool)`

GetCampaignRevisionVersionIdOk returns a tuple with the CampaignRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRevisionVersionId

`func (o *EffectDeductLoyaltyPoints) SetCampaignRevisionVersionId(v int64)`

SetCampaignRevisionVersionId sets CampaignRevisionVersionId field to given value.

### HasCampaignRevisionVersionId

`func (o *EffectDeductLoyaltyPoints) HasCampaignRevisionVersionId() bool`

HasCampaignRevisionVersionId returns a boolean if a field has been set.

### GetSelectedPriceType

`func (o *EffectDeductLoyaltyPoints) GetSelectedPriceType() string`

GetSelectedPriceType returns the SelectedPriceType field if non-nil, zero value otherwise.

### GetSelectedPriceTypeOk

`func (o *EffectDeductLoyaltyPoints) GetSelectedPriceTypeOk() (*string, bool)`

GetSelectedPriceTypeOk returns a tuple with the SelectedPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPriceType

`func (o *EffectDeductLoyaltyPoints) SetSelectedPriceType(v string)`

SetSelectedPriceType sets SelectedPriceType field to given value.

### HasSelectedPriceType

`func (o *EffectDeductLoyaltyPoints) HasSelectedPriceType() bool`

HasSelectedPriceType returns a boolean if a field has been set.

### GetSelectedPrice

`func (o *EffectDeductLoyaltyPoints) GetSelectedPrice() float32`

GetSelectedPrice returns the SelectedPrice field if non-nil, zero value otherwise.

### GetSelectedPriceOk

`func (o *EffectDeductLoyaltyPoints) GetSelectedPriceOk() (*float32, bool)`

GetSelectedPriceOk returns a tuple with the SelectedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPrice

`func (o *EffectDeductLoyaltyPoints) SetSelectedPrice(v float32)`

SetSelectedPrice sets SelectedPrice field to given value.

### HasSelectedPrice

`func (o *EffectDeductLoyaltyPoints) HasSelectedPrice() bool`

HasSelectedPrice returns a boolean if a field has been set.

### GetAdjustmentReferenceId

`func (o *EffectDeductLoyaltyPoints) GetAdjustmentReferenceId() string`

GetAdjustmentReferenceId returns the AdjustmentReferenceId field if non-nil, zero value otherwise.

### GetAdjustmentReferenceIdOk

`func (o *EffectDeductLoyaltyPoints) GetAdjustmentReferenceIdOk() (*string, bool)`

GetAdjustmentReferenceIdOk returns a tuple with the AdjustmentReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentReferenceId

`func (o *EffectDeductLoyaltyPoints) SetAdjustmentReferenceId(v string)`

SetAdjustmentReferenceId sets AdjustmentReferenceId field to given value.

### HasAdjustmentReferenceId

`func (o *EffectDeductLoyaltyPoints) HasAdjustmentReferenceId() bool`

HasAdjustmentReferenceId returns a boolean if a field has been set.

### GetRewardId

`func (o *EffectDeductLoyaltyPoints) GetRewardId() int64`

GetRewardId returns the RewardId field if non-nil, zero value otherwise.

### GetRewardIdOk

`func (o *EffectDeductLoyaltyPoints) GetRewardIdOk() (*int64, bool)`

GetRewardIdOk returns a tuple with the RewardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardId

`func (o *EffectDeductLoyaltyPoints) SetRewardId(v int64)`

SetRewardId sets RewardId field to given value.

### HasRewardId

`func (o *EffectDeductLoyaltyPoints) HasRewardId() bool`

HasRewardId returns a boolean if a field has been set.

### GetRewardIntegrationId

`func (o *EffectDeductLoyaltyPoints) GetRewardIntegrationId() string`

GetRewardIntegrationId returns the RewardIntegrationId field if non-nil, zero value otherwise.

### GetRewardIntegrationIdOk

`func (o *EffectDeductLoyaltyPoints) GetRewardIntegrationIdOk() (*string, bool)`

GetRewardIntegrationIdOk returns a tuple with the RewardIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardIntegrationId

`func (o *EffectDeductLoyaltyPoints) SetRewardIntegrationId(v string)`

SetRewardIntegrationId sets RewardIntegrationId field to given value.

### HasRewardIntegrationId

`func (o *EffectDeductLoyaltyPoints) HasRewardIntegrationId() bool`

HasRewardIntegrationId returns a boolean if a field has been set.

### GetProps

`func (o *EffectDeductLoyaltyPoints) GetProps() DeductLoyaltyPointsEffectProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *EffectDeductLoyaltyPoints) GetPropsOk() (*DeductLoyaltyPointsEffectProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *EffectDeductLoyaltyPoints) SetProps(v DeductLoyaltyPointsEffectProps)`

SetProps sets Props field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


