/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// Effect
type Effect struct {
	// The ID of the campaign that triggered this effect
	CampaignId int32 `json:"campaignId"`
	// The ID of the ruleset that was active in the campaign when this effect was triggered
	RulesetId int32 `json:"rulesetId"`
	// The position of the rule that triggered this effect within the ruleset
	RuleIndex int32 `json:"ruleIndex"`
	// The name of the rule that triggered this effect
	RuleName string `json:"ruleName"`
	// The type of effect that was triggered
	EffectType string `json:"effectType"`
	// The ID of the coupon that was being evaluated when this effect was triggered
	TriggeredByCoupon *int32                 `json:"triggeredByCoupon,omitempty"`
	Props             map[string]interface{} `json:"props"`
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

// GetProps returns the Props field value
func (o *Effect) GetProps() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Props
}

// SetProps sets field value
func (o *Effect) SetProps(v map[string]interface{}) {
	o.Props = v
}

type NullableEffect struct {
	Value        Effect
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
