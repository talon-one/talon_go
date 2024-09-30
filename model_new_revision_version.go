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
)

// NewRevisionVersion struct for NewRevisionVersion
type NewRevisionVersion struct {
	// A user-facing name for this campaign.
	Name *string `json:"name,omitempty"`
	// Timestamp when the campaign will become active.
	StartTime *NullableTime `json:"startTime,omitempty"`
	// Timestamp when the campaign will become inactive.
	EndTime *NullableTime `json:"endTime,omitempty"`
	// Arbitrary properties associated with this campaign.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// A detailed description of the campaign.
	Description *NullableString `json:"description,omitempty"`
	// The ID of the ruleset this campaign template will use.
	ActiveRulesetId *NullableInt32 `json:"activeRulesetId,omitempty"`
	// A list of tags for the campaign template.
	Tags             *[]string              `json:"tags,omitempty"`
	CouponSettings   *CodeGeneratorSettings `json:"couponSettings,omitempty"`
	ReferralSettings *CodeGeneratorSettings `json:"referralSettings,omitempty"`
	// The set of limits that will operate for this campaign version.
	Limits *[]LimitConfig `json:"limits,omitempty"`
	// A list of features for the campaign template.
	Features *[]string `json:"features,omitempty"`
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NewRevisionVersion) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetStartTime() NullableTime {
	if o == nil || o.StartTime == nil {
		var ret NullableTime
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetStartTimeOk() (NullableTime, bool) {
	if o == nil || o.StartTime == nil {
		var ret NullableTime
		return ret, false
	}
	return *o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *NewRevisionVersion) SetStartTime(v NullableTime) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetEndTime() NullableTime {
	if o == nil || o.EndTime == nil {
		var ret NullableTime
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetEndTimeOk() (NullableTime, bool) {
	if o == nil || o.EndTime == nil {
		var ret NullableTime
		return ret, false
	}
	return *o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given NullableTime and assigns it to the EndTime field.
func (o *NewRevisionVersion) SetEndTime(v NullableTime) {
	o.EndTime = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewRevisionVersion) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetDescription() NullableString {
	if o == nil || o.Description == nil {
		var ret NullableString
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetDescriptionOk() (NullableString, bool) {
	if o == nil || o.Description == nil {
		var ret NullableString
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NewRevisionVersion) SetDescription(v NullableString) {
	o.Description = &v
}

// GetActiveRulesetId returns the ActiveRulesetId field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetActiveRulesetId() NullableInt32 {
	if o == nil || o.ActiveRulesetId == nil {
		var ret NullableInt32
		return ret
	}
	return *o.ActiveRulesetId
}

// GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetActiveRulesetIdOk() (NullableInt32, bool) {
	if o == nil || o.ActiveRulesetId == nil {
		var ret NullableInt32
		return ret, false
	}
	return *o.ActiveRulesetId, true
}

// HasActiveRulesetId returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasActiveRulesetId() bool {
	if o != nil && o.ActiveRulesetId != nil {
		return true
	}

	return false
}

// SetActiveRulesetId gets a reference to the given NullableInt32 and assigns it to the ActiveRulesetId field.
func (o *NewRevisionVersion) SetActiveRulesetId(v NullableInt32) {
	o.ActiveRulesetId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *NewRevisionVersion) SetTags(v []string) {
	o.Tags = &v
}

// GetCouponSettings returns the CouponSettings field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetCouponSettings() CodeGeneratorSettings {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.CouponSettings
}

// GetCouponSettingsOk returns a tuple with the CouponSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetCouponSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.CouponSettings, true
}

// HasCouponSettings returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasCouponSettings() bool {
	if o != nil && o.CouponSettings != nil {
		return true
	}

	return false
}

// SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.
func (o *NewRevisionVersion) SetCouponSettings(v CodeGeneratorSettings) {
	o.CouponSettings = &v
}

// GetReferralSettings returns the ReferralSettings field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetReferralSettings() CodeGeneratorSettings {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.ReferralSettings
}

// GetReferralSettingsOk returns a tuple with the ReferralSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetReferralSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.ReferralSettings, true
}

// HasReferralSettings returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasReferralSettings() bool {
	if o != nil && o.ReferralSettings != nil {
		return true
	}

	return false
}

// SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.
func (o *NewRevisionVersion) SetReferralSettings(v CodeGeneratorSettings) {
	o.ReferralSettings = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetLimits() []LimitConfig {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetLimitsOk() ([]LimitConfig, bool) {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret, false
	}
	return *o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.
func (o *NewRevisionVersion) SetLimits(v []LimitConfig) {
	o.Limits = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *NewRevisionVersion) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRevisionVersion) GetFeaturesOk() ([]string, bool) {
	if o == nil || o.Features == nil {
		var ret []string
		return ret, false
	}
	return *o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *NewRevisionVersion) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *NewRevisionVersion) SetFeatures(v []string) {
	o.Features = &v
}

type NullableNewRevisionVersion struct {
	Value        NewRevisionVersion
	ExplicitNull bool
}

func (v NullableNewRevisionVersion) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewRevisionVersion) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}