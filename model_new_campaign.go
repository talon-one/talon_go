/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}`
 *
 * API version:
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
	"time"
)

// NewCampaign struct for NewCampaign
type NewCampaign struct {
	// A user-facing name for this campaign.
	Name string `json:"name"`
	// A detailed description of the campaign.
	Description *string `json:"description,omitempty"`
	// Timestamp when the campaign will become active.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Timestamp when the campaign will become inactive.
	EndTime *time.Time `json:"endTime,omitempty"`
	// Arbitrary properties associated with this campaign.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// A disabled or archived campaign is not evaluated for rules or coupons.
	State string `json:"state"`
	// [ID of Ruleset](https://docs.talon.one/management-api#operation/getRulesets) this campaign applies on customer session evaluation.
	ActiveRulesetId *int32 `json:"activeRulesetId,omitempty"`
	// A list of tags for the campaign.
	Tags []string `json:"tags"`
	// The features enabled in this campaign.
	Features         []string               `json:"features"`
	CouponSettings   *CodeGeneratorSettings `json:"couponSettings,omitempty"`
	ReferralSettings *CodeGeneratorSettings `json:"referralSettings,omitempty"`
	// The set of [budget limits](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets) for this campaign.
	Limits []LimitConfig `json:"limits"`
	// The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/managing-campaign-groups) this campaign belongs to.
	CampaignGroups *[]int32 `json:"campaignGroups,omitempty"`
	// The campaign type. Possible type values:   - `cartItem`: Type of campaign that can apply effects only to cart items.   - `advanced`: Type of campaign that can apply effects to customer sessions and cart items.
	Type *string `json:"type,omitempty"`
	// A list of store IDs that you want to link to the campaign.  **Note:** Campaigns with linked store IDs will only be evaluated when there is a [customer session update](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) that references a linked store.
	LinkedStoreIds *[]int32 `json:"linkedStoreIds,omitempty"`
	// The ID of the campaign evaluation group the campaign belongs to.
	EvaluationGroupId *int32 `json:"evaluationGroupId,omitempty"`
}

// GetName returns the Name field value
func (o *NewCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewCampaign) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewCampaign) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewCampaign) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewCampaign) SetDescription(v string) {
	o.Description = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *NewCampaign) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetStartTimeOk() (time.Time, bool) {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *NewCampaign) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *NewCampaign) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *NewCampaign) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetEndTimeOk() (time.Time, bool) {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *NewCampaign) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *NewCampaign) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewCampaign) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewCampaign) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewCampaign) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetState returns the State field value
func (o *NewCampaign) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// SetState sets field value
func (o *NewCampaign) SetState(v string) {
	o.State = v
}

// GetActiveRulesetId returns the ActiveRulesetId field value if set, zero value otherwise.
func (o *NewCampaign) GetActiveRulesetId() int32 {
	if o == nil || o.ActiveRulesetId == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesetId
}

// GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetActiveRulesetIdOk() (int32, bool) {
	if o == nil || o.ActiveRulesetId == nil {
		var ret int32
		return ret, false
	}
	return *o.ActiveRulesetId, true
}

// HasActiveRulesetId returns a boolean if a field has been set.
func (o *NewCampaign) HasActiveRulesetId() bool {
	if o != nil && o.ActiveRulesetId != nil {
		return true
	}

	return false
}

// SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.
func (o *NewCampaign) SetActiveRulesetId(v int32) {
	o.ActiveRulesetId = &v
}

// GetTags returns the Tags field value
func (o *NewCampaign) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// SetTags sets field value
func (o *NewCampaign) SetTags(v []string) {
	o.Tags = v
}

// GetFeatures returns the Features field value
func (o *NewCampaign) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Features
}

// SetFeatures sets field value
func (o *NewCampaign) SetFeatures(v []string) {
	o.Features = v
}

// GetCouponSettings returns the CouponSettings field value if set, zero value otherwise.
func (o *NewCampaign) GetCouponSettings() CodeGeneratorSettings {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.CouponSettings
}

// GetCouponSettingsOk returns a tuple with the CouponSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetCouponSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.CouponSettings, true
}

// HasCouponSettings returns a boolean if a field has been set.
func (o *NewCampaign) HasCouponSettings() bool {
	if o != nil && o.CouponSettings != nil {
		return true
	}

	return false
}

// SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.
func (o *NewCampaign) SetCouponSettings(v CodeGeneratorSettings) {
	o.CouponSettings = &v
}

// GetReferralSettings returns the ReferralSettings field value if set, zero value otherwise.
func (o *NewCampaign) GetReferralSettings() CodeGeneratorSettings {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.ReferralSettings
}

// GetReferralSettingsOk returns a tuple with the ReferralSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetReferralSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.ReferralSettings, true
}

// HasReferralSettings returns a boolean if a field has been set.
func (o *NewCampaign) HasReferralSettings() bool {
	if o != nil && o.ReferralSettings != nil {
		return true
	}

	return false
}

// SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.
func (o *NewCampaign) SetReferralSettings(v CodeGeneratorSettings) {
	o.ReferralSettings = &v
}

// GetLimits returns the Limits field value
func (o *NewCampaign) GetLimits() []LimitConfig {
	if o == nil {
		var ret []LimitConfig
		return ret
	}

	return o.Limits
}

// SetLimits sets field value
func (o *NewCampaign) SetLimits(v []LimitConfig) {
	o.Limits = v
}

// GetCampaignGroups returns the CampaignGroups field value if set, zero value otherwise.
func (o *NewCampaign) GetCampaignGroups() []int32 {
	if o == nil || o.CampaignGroups == nil {
		var ret []int32
		return ret
	}
	return *o.CampaignGroups
}

// GetCampaignGroupsOk returns a tuple with the CampaignGroups field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetCampaignGroupsOk() ([]int32, bool) {
	if o == nil || o.CampaignGroups == nil {
		var ret []int32
		return ret, false
	}
	return *o.CampaignGroups, true
}

// HasCampaignGroups returns a boolean if a field has been set.
func (o *NewCampaign) HasCampaignGroups() bool {
	if o != nil && o.CampaignGroups != nil {
		return true
	}

	return false
}

// SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.
func (o *NewCampaign) SetCampaignGroups(v []int32) {
	o.CampaignGroups = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NewCampaign) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NewCampaign) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NewCampaign) SetType(v string) {
	o.Type = &v
}

// GetLinkedStoreIds returns the LinkedStoreIds field value if set, zero value otherwise.
func (o *NewCampaign) GetLinkedStoreIds() []int32 {
	if o == nil || o.LinkedStoreIds == nil {
		var ret []int32
		return ret
	}
	return *o.LinkedStoreIds
}

// GetLinkedStoreIdsOk returns a tuple with the LinkedStoreIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetLinkedStoreIdsOk() ([]int32, bool) {
	if o == nil || o.LinkedStoreIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.LinkedStoreIds, true
}

// HasLinkedStoreIds returns a boolean if a field has been set.
func (o *NewCampaign) HasLinkedStoreIds() bool {
	if o != nil && o.LinkedStoreIds != nil {
		return true
	}

	return false
}

// SetLinkedStoreIds gets a reference to the given []int32 and assigns it to the LinkedStoreIds field.
func (o *NewCampaign) SetLinkedStoreIds(v []int32) {
	o.LinkedStoreIds = &v
}

// GetEvaluationGroupId returns the EvaluationGroupId field value if set, zero value otherwise.
func (o *NewCampaign) GetEvaluationGroupId() int32 {
	if o == nil || o.EvaluationGroupId == nil {
		var ret int32
		return ret
	}
	return *o.EvaluationGroupId
}

// GetEvaluationGroupIdOk returns a tuple with the EvaluationGroupId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCampaign) GetEvaluationGroupIdOk() (int32, bool) {
	if o == nil || o.EvaluationGroupId == nil {
		var ret int32
		return ret, false
	}
	return *o.EvaluationGroupId, true
}

// HasEvaluationGroupId returns a boolean if a field has been set.
func (o *NewCampaign) HasEvaluationGroupId() bool {
	if o != nil && o.EvaluationGroupId != nil {
		return true
	}

	return false
}

// SetEvaluationGroupId gets a reference to the given int32 and assigns it to the EvaluationGroupId field.
func (o *NewCampaign) SetEvaluationGroupId(v int32) {
	o.EvaluationGroupId = &v
}

type NullableNewCampaign struct {
	Value        NewCampaign
	ExplicitNull bool
}

func (v NullableNewCampaign) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewCampaign) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
