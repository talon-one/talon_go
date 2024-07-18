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
	"time"
)

// RulesetAllOf struct for RulesetAllOf
type RulesetAllOf struct {
	// The ID of the campaign that owns this entity.
	CampaignId *int32 `json:"campaignId,omitempty"`
	// The ID of the campaign template that owns this entity.
	TemplateId *int32 `json:"templateId,omitempty"`
	// Timestamp indicating when this Ruleset was activated.
	ActivatedAt *time.Time `json:"activatedAt,omitempty"`
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *RulesetAllOf) GetCampaignId() int32 {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RulesetAllOf) GetCampaignIdOk() (int32, bool) {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret, false
	}
	return *o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *RulesetAllOf) HasCampaignId() bool {
	if o != nil && o.CampaignId != nil {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.
func (o *RulesetAllOf) SetCampaignId(v int32) {
	o.CampaignId = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *RulesetAllOf) GetTemplateId() int32 {
	if o == nil || o.TemplateId == nil {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RulesetAllOf) GetTemplateIdOk() (int32, bool) {
	if o == nil || o.TemplateId == nil {
		var ret int32
		return ret, false
	}
	return *o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *RulesetAllOf) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *RulesetAllOf) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise.
func (o *RulesetAllOf) GetActivatedAt() time.Time {
	if o == nil || o.ActivatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ActivatedAt
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RulesetAllOf) GetActivatedAtOk() (time.Time, bool) {
	if o == nil || o.ActivatedAt == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ActivatedAt, true
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *RulesetAllOf) HasActivatedAt() bool {
	if o != nil && o.ActivatedAt != nil {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given time.Time and assigns it to the ActivatedAt field.
func (o *RulesetAllOf) SetActivatedAt(v time.Time) {
	o.ActivatedAt = &v
}

type NullableRulesetAllOf struct {
	Value RulesetAllOf
	ExplicitNull bool
}

func (v NullableRulesetAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRulesetAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
