# Effect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int32** | The ID of the campaign that triggered this effect. | 
**RulesetId** | Pointer to **int32** | The ID of the ruleset that was active in the campaign when this effect was triggered. | 
**RuleIndex** | Pointer to **int32** | The position of the rule that triggered this effect within the ruleset. | 
**RuleName** | Pointer to **string** | The name of the rule that triggered this effect. | 
**EffectType** | Pointer to **string** | The type of effect that was triggered. See [API effects](https://docs.talon.one/docs/dev/integration-api/api-effects). | 
**TriggeredByCoupon** | Pointer to **int32** | The ID of the coupon that was being evaluated when this effect was triggered. | [optional] 
**TriggeredForCatalogItem** | Pointer to **int32** | The ID of the catalog item that was being evaluated when this effect was triggered. | [optional] 
**ConditionIndex** | Pointer to **int32** | The index of the condition that was triggered. | [optional] 
**EvaluationGroupID** | Pointer to **int32** | The ID of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation). | [optional] 
**EvaluationGroupMode** | Pointer to **string** | The evaluation mode of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation). | [optional] 
**CampaignRevisionId** | Pointer to **int32** | The revision ID of the campaign that was used when triggering the effect. | [optional] 
**CampaignRevisionVersionId** | Pointer to **int32** | The revision version ID of the campaign that was used when triggering the effect. | [optional] 
**Props** | Pointer to [**map[string]interface{}**](.md) | The properties of the effect. See [API effects](https://docs.talon.one/docs/dev/integration-api/api-effects). | 

## Methods

### GetCampaignId

`func (o *Effect) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Effect) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *Effect) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *Effect) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetRulesetId

`func (o *Effect) GetRulesetId() int32`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *Effect) GetRulesetIdOk() (int32, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRulesetId

`func (o *Effect) HasRulesetId() bool`

HasRulesetId returns a boolean if a field has been set.

### SetRulesetId

`func (o *Effect) SetRulesetId(v int32)`

SetRulesetId gets a reference to the given int32 and assigns it to the RulesetId field.

### GetRuleIndex

`func (o *Effect) GetRuleIndex() int32`

GetRuleIndex returns the RuleIndex field if non-nil, zero value otherwise.

### GetRuleIndexOk

`func (o *Effect) GetRuleIndexOk() (int32, bool)`

GetRuleIndexOk returns a tuple with the RuleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleIndex

`func (o *Effect) HasRuleIndex() bool`

HasRuleIndex returns a boolean if a field has been set.

### SetRuleIndex

`func (o *Effect) SetRuleIndex(v int32)`

SetRuleIndex gets a reference to the given int32 and assigns it to the RuleIndex field.

### GetRuleName

`func (o *Effect) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *Effect) GetRuleNameOk() (string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleName

`func (o *Effect) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleName

`func (o *Effect) SetRuleName(v string)`

SetRuleName gets a reference to the given string and assigns it to the RuleName field.

### GetEffectType

`func (o *Effect) GetEffectType() string`

GetEffectType returns the EffectType field if non-nil, zero value otherwise.

### GetEffectTypeOk

`func (o *Effect) GetEffectTypeOk() (string, bool)`

GetEffectTypeOk returns a tuple with the EffectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffectType

`func (o *Effect) HasEffectType() bool`

HasEffectType returns a boolean if a field has been set.

### SetEffectType

`func (o *Effect) SetEffectType(v string)`

SetEffectType gets a reference to the given string and assigns it to the EffectType field.

### GetTriggeredByCoupon

`func (o *Effect) GetTriggeredByCoupon() int32`

GetTriggeredByCoupon returns the TriggeredByCoupon field if non-nil, zero value otherwise.

### GetTriggeredByCouponOk

`func (o *Effect) GetTriggeredByCouponOk() (int32, bool)`

GetTriggeredByCouponOk returns a tuple with the TriggeredByCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggeredByCoupon

`func (o *Effect) HasTriggeredByCoupon() bool`

HasTriggeredByCoupon returns a boolean if a field has been set.

### SetTriggeredByCoupon

`func (o *Effect) SetTriggeredByCoupon(v int32)`

SetTriggeredByCoupon gets a reference to the given int32 and assigns it to the TriggeredByCoupon field.

### GetTriggeredForCatalogItem

`func (o *Effect) GetTriggeredForCatalogItem() int32`

GetTriggeredForCatalogItem returns the TriggeredForCatalogItem field if non-nil, zero value otherwise.

### GetTriggeredForCatalogItemOk

`func (o *Effect) GetTriggeredForCatalogItemOk() (int32, bool)`

GetTriggeredForCatalogItemOk returns a tuple with the TriggeredForCatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggeredForCatalogItem

`func (o *Effect) HasTriggeredForCatalogItem() bool`

HasTriggeredForCatalogItem returns a boolean if a field has been set.

### SetTriggeredForCatalogItem

`func (o *Effect) SetTriggeredForCatalogItem(v int32)`

SetTriggeredForCatalogItem gets a reference to the given int32 and assigns it to the TriggeredForCatalogItem field.

### GetConditionIndex

`func (o *Effect) GetConditionIndex() int32`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *Effect) GetConditionIndexOk() (int32, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditionIndex

`func (o *Effect) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### SetConditionIndex

`func (o *Effect) SetConditionIndex(v int32)`

SetConditionIndex gets a reference to the given int32 and assigns it to the ConditionIndex field.

### GetEvaluationGroupID

`func (o *Effect) GetEvaluationGroupID() int32`

GetEvaluationGroupID returns the EvaluationGroupID field if non-nil, zero value otherwise.

### GetEvaluationGroupIDOk

`func (o *Effect) GetEvaluationGroupIDOk() (int32, bool)`

GetEvaluationGroupIDOk returns a tuple with the EvaluationGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationGroupID

`func (o *Effect) HasEvaluationGroupID() bool`

HasEvaluationGroupID returns a boolean if a field has been set.

### SetEvaluationGroupID

`func (o *Effect) SetEvaluationGroupID(v int32)`

SetEvaluationGroupID gets a reference to the given int32 and assigns it to the EvaluationGroupID field.

### GetEvaluationGroupMode

`func (o *Effect) GetEvaluationGroupMode() string`

GetEvaluationGroupMode returns the EvaluationGroupMode field if non-nil, zero value otherwise.

### GetEvaluationGroupModeOk

`func (o *Effect) GetEvaluationGroupModeOk() (string, bool)`

GetEvaluationGroupModeOk returns a tuple with the EvaluationGroupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationGroupMode

`func (o *Effect) HasEvaluationGroupMode() bool`

HasEvaluationGroupMode returns a boolean if a field has been set.

### SetEvaluationGroupMode

`func (o *Effect) SetEvaluationGroupMode(v string)`

SetEvaluationGroupMode gets a reference to the given string and assigns it to the EvaluationGroupMode field.

### GetCampaignRevisionId

`func (o *Effect) GetCampaignRevisionId() int32`

GetCampaignRevisionId returns the CampaignRevisionId field if non-nil, zero value otherwise.

### GetCampaignRevisionIdOk

`func (o *Effect) GetCampaignRevisionIdOk() (int32, bool)`

GetCampaignRevisionIdOk returns a tuple with the CampaignRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignRevisionId

`func (o *Effect) HasCampaignRevisionId() bool`

HasCampaignRevisionId returns a boolean if a field has been set.

### SetCampaignRevisionId

`func (o *Effect) SetCampaignRevisionId(v int32)`

SetCampaignRevisionId gets a reference to the given int32 and assigns it to the CampaignRevisionId field.

### GetCampaignRevisionVersionId

`func (o *Effect) GetCampaignRevisionVersionId() int32`

GetCampaignRevisionVersionId returns the CampaignRevisionVersionId field if non-nil, zero value otherwise.

### GetCampaignRevisionVersionIdOk

`func (o *Effect) GetCampaignRevisionVersionIdOk() (int32, bool)`

GetCampaignRevisionVersionIdOk returns a tuple with the CampaignRevisionVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignRevisionVersionId

`func (o *Effect) HasCampaignRevisionVersionId() bool`

HasCampaignRevisionVersionId returns a boolean if a field has been set.

### SetCampaignRevisionVersionId

`func (o *Effect) SetCampaignRevisionVersionId(v int32)`

SetCampaignRevisionVersionId gets a reference to the given int32 and assigns it to the CampaignRevisionVersionId field.

### GetProps

`func (o *Effect) GetProps() map[string]interface{}`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *Effect) GetPropsOk() (map[string]interface{}, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProps

`func (o *Effect) HasProps() bool`

HasProps returns a boolean if a field has been set.

### SetProps

`func (o *Effect) SetProps(v map[string]interface{})`

SetProps gets a reference to the given map[string]interface{} and assigns it to the Props field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


