# EffectAddNegativeLoyaltyPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentId** | Pointer to **int64** | The ID of the experiment that campaign belongs to. | [optional] 
**CampaignId** | Pointer to **int64** | The ID of the campaign that triggered this effect. | 
**RulesetId** | Pointer to **int64** | The ID of the ruleset that was active in the campaign when this effect was triggered. | 
**RuleIndex** | Pointer to **int64** | The position of the rule that triggered this effect within the ruleset. | 
**RuleName** | Pointer to **string** | The name of the rule that triggered this effect. | 
**EffectType** | Pointer to **string** | An effect discriminator of type &#x60;addNegativeLoyaltyPoints&#x60;. | 
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

## Methods

### NewEffectAddNegativeLoyaltyPoints

`func NewEffectAddNegativeLoyaltyPoints(campaignId int64, rulesetId int64, ruleIndex int64, ruleName string, effectType string, ) *EffectAddNegativeLoyaltyPoints`

NewEffectAddNegativeLoyaltyPoints instantiates a new EffectAddNegativeLoyaltyPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectAddNegativeLoyaltyPointsWithDefaults

`func NewEffectAddNegativeLoyaltyPointsWithDefaults() *EffectAddNegativeLoyaltyPoints`

NewEffectAddNegativeLoyaltyPointsWithDefaults instantiates a new EffectAddNegativeLoyaltyPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentId

`func (o *EffectAddNegativeLoyaltyPoints) GetExperimentId() int64`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *EffectAddNegativeLoyaltyPoints) GetExperimentIdOk() (*int64, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *EffectAddNegativeLoyaltyPoints) SetExperimentId(v int64)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *EffectAddNegativeLoyaltyPoints) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### GetCampaignId

`func (o *EffectAddNegativeLoyaltyPoints) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *EffectAddNegativeLoyaltyPoints) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *EffectAddNegativeLoyaltyPoints) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetRulesetId

`func (o *EffectAddNegativeLoyaltyPoints) GetRulesetId() int64`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *EffectAddNegativeLoyaltyPoints) GetRulesetIdOk() (*int64, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *EffectAddNegativeLoyaltyPoints) SetRulesetId(v int64)`

SetRulesetId sets RulesetId field to given value.


### GetRuleIndex

`func (o *EffectAddNegativeLoyaltyPoints) GetRuleIndex() int64`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *EffectAddNegativeLoyaltyPoints) GetRuleIndexOk() (*int64, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIndex

`func (o *EffectAddNegativeLoyaltyPoints) SetRuleIndex(v int64)`

SetRuleIndex sets RuleIndex field to given value.


### GetRuleName

`func (o *EffectAddNegativeLoyaltyPoints) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *EffectAddNegativeLoyaltyPoints) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *EffectAddNegativeLoyaltyPoints) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetEffectType

`func (o *EffectAddNegativeLoyaltyPoints) GetEffectType() string`

GetEffectType returns the EffectType field if non-nil, zero value otherwise.

### GetEffectTypeOk

`func (o *EffectAddNegativeLoyaltyPoints) GetEffectTypeOk() (*string, bool)`

GetEffectTypeOk returns a tuple with the EffectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectType

`func (o *EffectAddNegativeLoyaltyPoints) SetEffectType(v string)`

SetEffectType sets EffectType field to given value.


### GetTriggeredByCoupon

`func (o *EffectAddNegativeLoyaltyPoints) GetTriggeredByCoupon() int64`

GetTriggeredByCoupon returns the TriggeredByCoupon field if non-nil, zero value otherwise.

### GetTriggeredByCouponOk

`func (o *EffectAddNegativeLoyaltyPoints) GetTriggeredByCouponOk() (*int64, bool)`

GetTriggeredByCouponOk returns a tuple with the TriggeredByCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredByCoupon

`func (o *EffectAddNegativeLoyaltyPoints) SetTriggeredByCoupon(v int64)`

SetTriggeredByCoupon sets TriggeredByCoupon field to given value.

### HasTriggeredByCoupon

`func (o *EffectAddNegativeLoyaltyPoints) HasTriggeredByCoupon() bool`

HasTriggeredByCoupon returns a boolean if a field has been set.

### GetTriggeredForCatalogItem

`func (o *EffectAddNegativeLoyaltyPoints) GetTriggeredForCatalogItem() int64`

GetTriggeredForCatalogItem returns the TriggeredForCatalogItem field if non-nil, zero value otherwise.

### GetTriggeredForCatalogItemOk

`func (o *EffectAddNegativeLoyaltyPoints) GetTriggeredForCatalogItemOk() (*int64, bool)`

GetTriggeredForCatalogItemOk returns a tuple with the TriggeredForCatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredForCatalogItem

`func (o *EffectAddNegativeLoyaltyPoints) SetTriggeredForCatalogItem(v int64)`

SetTriggeredForCatalogItem sets TriggeredForCatalogItem field to given value.

### HasTriggeredForCatalogItem

`func (o *EffectAddNegativeLoyaltyPoints) HasTriggeredForCatalogItem() bool`

HasTriggeredForCatalogItem returns a boolean if a field has been set.

### GetConditionIndex

`func (o *EffectAddNegativeLoyaltyPoints) GetConditionIndex() int64`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *EffectAddNegativeLoyaltyPoints) GetConditionIndexOk() (*int64, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *EffectAddNegativeLoyaltyPoints) SetConditionIndex(v int64)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *EffectAddNegativeLoyaltyPoints) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### GetEvaluationGroupID

`func (o *EffectAddNegativeLoyaltyPoints) GetEvaluationGroupID() int64`

GetEvaluationGroupID returns the EvaluationGroupID field if non-nil, zero value otherwise.

### GetEvaluationGroupIDOk

`func (o *EffectAddNegativeLoyaltyPoints) GetEvaluationGroupIDOk() (*int64, bool)`

GetEvaluationGroupIDOk returns a tuple with the EvaluationGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupID

`func (o *EffectAddNegativeLoyaltyPoints) SetEvaluationGroupID(v int64)`

SetEvaluationGroupID sets EvaluationGroupID field to given value.

### HasEvaluationGroupID

`func (o *EffectAddNegativeLoyaltyPoints) HasEvaluationGroupID() bool`

HasEvaluationGroupID returns a boolean if a field has been set.

### GetEvaluationGroupMode

`func (o *EffectAddNegativeLoyaltyPoints) GetEvaluationGroupMode() string`

GetEvaluationGroupMode returns the EvaluationGroupMode field if non-nil, zero value otherwise.

### GetEvaluationGroupModeOk

`func (o *EffectAddNegativeLoyaltyPoints) GetEvaluationGroupModeOk() (*string, bool)`

GetEvaluationGroupModeOk returns a tuple with the EvaluationGroupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationGroupMode

`func (o *EffectAddNegativeLoyaltyPoints) SetEvaluationGroupMode(v string)`

SetEvaluationGroupMode sets EvaluationGroupMode field to given value.

### HasEvaluationGroupMode

`func (o *EffectAddNegativeLoyaltyPoints) HasEvaluationGroupMode() bool`

HasEvaluationGroupMode returns a boolean if a field has been set.

### GetCampaignRevisionId

`func (o *EffectAddNegativeLoyaltyPoints) GetCampaignRevisionId() int64`

GetCampaignRevisionId returns the CampaignRevisionId field if non-nil, zero value otherwise.

### GetCampaignRevisionIdOk

`func (o *EffectAddNegativeLoyaltyPoints) GetCampaignRevisionIdOk() (*int64, bool)`

GetCampaignRevisionIdOk returns a tuple with the CampaignRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRevisionId

`func (o *EffectAddNegativeLoyaltyPoints) SetCampaignRevisionId(v int64)`

SetCampaignRevisionId sets CampaignRevisionId field to given value.

### HasCampaignRevisionId

`func (o *EffectAddNegativeLoyaltyPoints) HasCampaignRevisionId() bool`

HasCampaignRevisionId returns a boolean if a field has been set.

### GetCampaignRevisionVersionId

`func (o *EffectAddNegativeLoyaltyPoints) GetCampaignRevisionVersionId() int64`

GetCampaignRevisionVersionId returns the CampaignRevisionVersionId field if non-nil, zero value otherwise.

### GetCampaignRevisionVersionIdOk

`func (o *EffectAddNegativeLoyaltyPoints) GetCampaignRevisionVersionIdOk() (*int64, bool)`

GetCampaignRevisionVersionIdOk returns a tuple with the CampaignRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRevisionVersionId

`func (o *EffectAddNegativeLoyaltyPoints) SetCampaignRevisionVersionId(v int64)`

SetCampaignRevisionVersionId sets CampaignRevisionVersionId field to given value.

### HasCampaignRevisionVersionId

`func (o *EffectAddNegativeLoyaltyPoints) HasCampaignRevisionVersionId() bool`

HasCampaignRevisionVersionId returns a boolean if a field has been set.

### GetSelectedPriceType

`func (o *EffectAddNegativeLoyaltyPoints) GetSelectedPriceType() string`

GetSelectedPriceType returns the SelectedPriceType field if non-nil, zero value otherwise.

### GetSelectedPriceTypeOk

`func (o *EffectAddNegativeLoyaltyPoints) GetSelectedPriceTypeOk() (*string, bool)`

GetSelectedPriceTypeOk returns a tuple with the SelectedPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPriceType

`func (o *EffectAddNegativeLoyaltyPoints) SetSelectedPriceType(v string)`

SetSelectedPriceType sets SelectedPriceType field to given value.

### HasSelectedPriceType

`func (o *EffectAddNegativeLoyaltyPoints) HasSelectedPriceType() bool`

HasSelectedPriceType returns a boolean if a field has been set.

### GetSelectedPrice

`func (o *EffectAddNegativeLoyaltyPoints) GetSelectedPrice() float32`

GetSelectedPrice returns the SelectedPrice field if non-nil, zero value otherwise.

### GetSelectedPriceOk

`func (o *EffectAddNegativeLoyaltyPoints) GetSelectedPriceOk() (*float32, bool)`

GetSelectedPriceOk returns a tuple with the SelectedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPrice

`func (o *EffectAddNegativeLoyaltyPoints) SetSelectedPrice(v float32)`

SetSelectedPrice sets SelectedPrice field to given value.

### HasSelectedPrice

`func (o *EffectAddNegativeLoyaltyPoints) HasSelectedPrice() bool`

HasSelectedPrice returns a boolean if a field has been set.

### GetAdjustmentReferenceId

`func (o *EffectAddNegativeLoyaltyPoints) GetAdjustmentReferenceId() string`

GetAdjustmentReferenceId returns the AdjustmentReferenceId field if non-nil, zero value otherwise.

### GetAdjustmentReferenceIdOk

`func (o *EffectAddNegativeLoyaltyPoints) GetAdjustmentReferenceIdOk() (*string, bool)`

GetAdjustmentReferenceIdOk returns a tuple with the AdjustmentReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentReferenceId

`func (o *EffectAddNegativeLoyaltyPoints) SetAdjustmentReferenceId(v string)`

SetAdjustmentReferenceId sets AdjustmentReferenceId field to given value.

### HasAdjustmentReferenceId

`func (o *EffectAddNegativeLoyaltyPoints) HasAdjustmentReferenceId() bool`

HasAdjustmentReferenceId returns a boolean if a field has been set.

### GetRewardId

`func (o *EffectAddNegativeLoyaltyPoints) GetRewardId() int64`

GetRewardId returns the RewardId field if non-nil, zero value otherwise.

### GetRewardIdOk

`func (o *EffectAddNegativeLoyaltyPoints) GetRewardIdOk() (*int64, bool)`

GetRewardIdOk returns a tuple with the RewardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardId

`func (o *EffectAddNegativeLoyaltyPoints) SetRewardId(v int64)`

SetRewardId sets RewardId field to given value.

### HasRewardId

`func (o *EffectAddNegativeLoyaltyPoints) HasRewardId() bool`

HasRewardId returns a boolean if a field has been set.

### GetRewardIntegrationId

`func (o *EffectAddNegativeLoyaltyPoints) GetRewardIntegrationId() string`

GetRewardIntegrationId returns the RewardIntegrationId field if non-nil, zero value otherwise.

### GetRewardIntegrationIdOk

`func (o *EffectAddNegativeLoyaltyPoints) GetRewardIntegrationIdOk() (*string, bool)`

GetRewardIntegrationIdOk returns a tuple with the RewardIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardIntegrationId

`func (o *EffectAddNegativeLoyaltyPoints) SetRewardIntegrationId(v string)`

SetRewardIntegrationId sets RewardIntegrationId field to given value.

### HasRewardIntegrationId

`func (o *EffectAddNegativeLoyaltyPoints) HasRewardIntegrationId() bool`

HasRewardIntegrationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


