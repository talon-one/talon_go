/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// Effect struct for Effect
type Effect struct {
	// The ID of the campaign that triggered this effect.
	CampaignId int32 `json:"campaignId"`
	// The ID of the ruleset that was active in the campaign when this effect was triggered.
	RulesetId int32 `json:"rulesetId"`
	// The position of the rule that triggered this effect within the ruleset.
	RuleIndex int32 `json:"ruleIndex"`
	// The name of the rule that triggered this effect.
	RuleName string `json:"ruleName"`
	// The type of effect that was triggered. See [API effects](https://docs.talon.one/docs/dev/integration-api/api-effects).
	EffectType string `json:"effectType"`
	// The ID of the coupon that was being evaluated when this effect was triggered.
	TriggeredByCoupon *int32 `json:"triggeredByCoupon,omitempty"`
	// The ID of the catalog item that was being evaluated when this effect was triggered.
	TriggeredForCatalogItem *int32 `json:"triggeredForCatalogItem,omitempty"`
	// The index of the condition that was triggered.
	ConditionIndex *int32 `json:"conditionIndex,omitempty"`
	// The ID of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation).
	EvaluationGroupID *int32 `json:"evaluationGroupID,omitempty"`
	// The evaluation mode of the evaluation group. For more information, see [Managing campaign evaluation](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation).
	EvaluationGroupMode *string `json:"evaluationGroupMode,omitempty"`
	Props EffectProps `json:"props"`
}

// GetCampaignId returns the CampaignId field value
func (o *Effect) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *Effect) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetRulesetId returns the RulesetId field value
func (o *Effect) GetRulesetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RulesetId
}

// SetRulesetId sets field value
func (o *Effect) SetRulesetId(v int32) {
	o.RulesetId = v
}

// GetRuleIndex returns the RuleIndex field value
func (o *Effect) GetRuleIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RuleIndex
}

// SetRuleIndex sets field value
func (o *Effect) SetRuleIndex(v int32) {
	o.RuleIndex = v
}

// GetRuleName returns the RuleName field value
func (o *Effect) GetRuleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleName
}

// SetRuleName sets field value
func (o *Effect) SetRuleName(v string) {
	o.RuleName = v
}

// GetEffectType returns the EffectType field value
func (o *Effect) GetEffectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectType
}

// SetEffectType sets field value
func (o *Effect) SetEffectType(v string) {
	o.EffectType = v
}

// GetTriggeredByCoupon returns the TriggeredByCoupon field value if set, zero value otherwise.
func (o *Effect) GetTriggeredByCoupon() int32 {
	if o == nil || o.TriggeredByCoupon == nil {
		var ret int32
		return ret
	}
	return *o.TriggeredByCoupon
}

// GetTriggeredByCouponOk returns a tuple with the TriggeredByCoupon field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetTriggeredByCouponOk() (int32, bool) {
	if o == nil || o.TriggeredByCoupon == nil {
		var ret int32
		return ret, false
	}
	return *o.TriggeredByCoupon, true
}

// HasTriggeredByCoupon returns a boolean if a field has been set.
func (o *Effect) HasTriggeredByCoupon() bool {
	if o != nil && o.TriggeredByCoupon != nil {
		return true
	}

	return false
}

// SetTriggeredByCoupon gets a reference to the given int32 and assigns it to the TriggeredByCoupon field.
func (o *Effect) SetTriggeredByCoupon(v int32) {
	o.TriggeredByCoupon = &v
}

// GetTriggeredForCatalogItem returns the TriggeredForCatalogItem field value if set, zero value otherwise.
func (o *Effect) GetTriggeredForCatalogItem() int32 {
	if o == nil || o.TriggeredForCatalogItem == nil {
		var ret int32
		return ret
	}
	return *o.TriggeredForCatalogItem
}

// GetTriggeredForCatalogItemOk returns a tuple with the TriggeredForCatalogItem field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetTriggeredForCatalogItemOk() (int32, bool) {
	if o == nil || o.TriggeredForCatalogItem == nil {
		var ret int32
		return ret, false
	}
	return *o.TriggeredForCatalogItem, true
}

// HasTriggeredForCatalogItem returns a boolean if a field has been set.
func (o *Effect) HasTriggeredForCatalogItem() bool {
	if o != nil && o.TriggeredForCatalogItem != nil {
		return true
	}

	return false
}

// SetTriggeredForCatalogItem gets a reference to the given int32 and assigns it to the TriggeredForCatalogItem field.
func (o *Effect) SetTriggeredForCatalogItem(v int32) {
	o.TriggeredForCatalogItem = &v
}

// GetConditionIndex returns the ConditionIndex field value if set, zero value otherwise.
func (o *Effect) GetConditionIndex() int32 {
	if o == nil || o.ConditionIndex == nil {
		var ret int32
		return ret
	}
	return *o.ConditionIndex
}

// GetConditionIndexOk returns a tuple with the ConditionIndex field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetConditionIndexOk() (int32, bool) {
	if o == nil || o.ConditionIndex == nil {
		var ret int32
		return ret, false
	}
	return *o.ConditionIndex, true
}

// HasConditionIndex returns a boolean if a field has been set.
func (o *Effect) HasConditionIndex() bool {
	if o != nil && o.ConditionIndex != nil {
		return true
	}

	return false
}

// SetConditionIndex gets a reference to the given int32 and assigns it to the ConditionIndex field.
func (o *Effect) SetConditionIndex(v int32) {
	o.ConditionIndex = &v
}

// GetEvaluationGroupID returns the EvaluationGroupID field value if set, zero value otherwise.
func (o *Effect) GetEvaluationGroupID() int32 {
	if o == nil || o.EvaluationGroupID == nil {
		var ret int32
		return ret
	}
	return *o.EvaluationGroupID
}

// GetEvaluationGroupIDOk returns a tuple with the EvaluationGroupID field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetEvaluationGroupIDOk() (int32, bool) {
	if o == nil || o.EvaluationGroupID == nil {
		var ret int32
		return ret, false
	}
	return *o.EvaluationGroupID, true
}

// HasEvaluationGroupID returns a boolean if a field has been set.
func (o *Effect) HasEvaluationGroupID() bool {
	if o != nil && o.EvaluationGroupID != nil {
		return true
	}

	return false
}

// SetEvaluationGroupID gets a reference to the given int32 and assigns it to the EvaluationGroupID field.
func (o *Effect) SetEvaluationGroupID(v int32) {
	o.EvaluationGroupID = &v
}

// GetEvaluationGroupMode returns the EvaluationGroupMode field value if set, zero value otherwise.
func (o *Effect) GetEvaluationGroupMode() string {
	if o == nil || o.EvaluationGroupMode == nil {
		var ret string
		return ret
	}
	return *o.EvaluationGroupMode
}

// GetEvaluationGroupModeOk returns a tuple with the EvaluationGroupMode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Effect) GetEvaluationGroupModeOk() (string, bool) {
	if o == nil || o.EvaluationGroupMode == nil {
		var ret string
		return ret, false
	}
	return *o.EvaluationGroupMode, true
}

// HasEvaluationGroupMode returns a boolean if a field has been set.
func (o *Effect) HasEvaluationGroupMode() bool {
	if o != nil && o.EvaluationGroupMode != nil {
		return true
	}

	return false
}

// SetEvaluationGroupMode gets a reference to the given string and assigns it to the EvaluationGroupMode field.
func (o *Effect) SetEvaluationGroupMode(v string) {
	o.EvaluationGroupMode = &v
}

// GetProps returns the Props field value
func (o *Effect) GetProps() EffectProps {
	if o == nil {
		var ret EffectProps
		return ret
	}

	return o.Props
}

// SetProps sets field value
func (o *Effect) SetProps(v EffectProps) {
	o.Props = v
}

type NullableEffect struct {
	Value Effect
	ExplicitNull bool
}

func (v NullableEffect) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEffect) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
