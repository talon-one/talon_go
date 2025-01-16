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

// UpdateCampaignTemplate struct for UpdateCampaignTemplate
type UpdateCampaignTemplate struct {
	// The campaign template name.
	Name string `json:"name"`
	// Customer-facing text that explains the objective of the template.
	Description string `json:"description"`
	// Customer-facing text that explains how to use the template. For example, you can use this property to explain the available attributes of this template, and how they can be modified when a user uses this template to create a new campaign.
	Instructions string `json:"instructions"`
	// The campaign attributes that campaigns created from this template will have by default.
	CampaignAttributes *map[string]interface{} `json:"campaignAttributes,omitempty"`
	// The campaign attributes that coupons created from this template will have by default.
	CouponAttributes *map[string]interface{} `json:"couponAttributes,omitempty"`
	// Only campaign templates in 'available' state may be used to create campaigns.
	State string `json:"state"`
	// The ID of the ruleset this campaign template will use.
	ActiveRulesetId *int32 `json:"activeRulesetId,omitempty"`
	// A list of tags for the campaign template.
	Tags *[]string `json:"tags,omitempty"`
	// A list of features for the campaign template.
	Features                  *[]string                                  `json:"features,omitempty"`
	CouponSettings            *CodeGeneratorSettings                     `json:"couponSettings,omitempty"`
	CouponReservationSettings *CampaignTemplateCouponReservationSettings `json:"couponReservationSettings,omitempty"`
	ReferralSettings          *CodeGeneratorSettings                     `json:"referralSettings,omitempty"`
	// The set of limits that operate for this campaign template.
	Limits *[]TemplateLimitConfig `json:"limits,omitempty"`
	// Fields which can be used to replace values in a rule.
	TemplateParams *[]CampaignTemplateParams `json:"templateParams,omitempty"`
	// A list of IDs of the Applications that are subscribed to this campaign template.
	ApplicationsIds []int32 `json:"applicationsIds"`
	// The campaign collections from the blueprint campaign for the template.
	CampaignCollections *[]CampaignTemplateCollection `json:"campaignCollections,omitempty"`
	// The default campaign group ID.
	DefaultCampaignGroupId *int32 `json:"defaultCampaignGroupId,omitempty"`
	// The campaign type. Possible type values:   - `cartItem`: Type of campaign that can apply effects only to cart items.   - `advanced`: Type of campaign that can apply effects to customer sessions and cart items.
	CampaignType *string `json:"campaignType,omitempty"`
}

// GetName returns the Name field value
func (o *UpdateCampaignTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *UpdateCampaignTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *UpdateCampaignTemplate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *UpdateCampaignTemplate) SetDescription(v string) {
	o.Description = v
}

// GetInstructions returns the Instructions field value
func (o *UpdateCampaignTemplate) GetInstructions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instructions
}

// SetInstructions sets field value
func (o *UpdateCampaignTemplate) SetInstructions(v string) {
	o.Instructions = v
}

// GetCampaignAttributes returns the CampaignAttributes field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCampaignAttributes() map[string]interface{} {
	if o == nil || o.CampaignAttributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CampaignAttributes
}

// GetCampaignAttributesOk returns a tuple with the CampaignAttributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCampaignAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.CampaignAttributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.CampaignAttributes, true
}

// HasCampaignAttributes returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCampaignAttributes() bool {
	if o != nil && o.CampaignAttributes != nil {
		return true
	}

	return false
}

// SetCampaignAttributes gets a reference to the given map[string]interface{} and assigns it to the CampaignAttributes field.
func (o *UpdateCampaignTemplate) SetCampaignAttributes(v map[string]interface{}) {
	o.CampaignAttributes = &v
}

// GetCouponAttributes returns the CouponAttributes field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCouponAttributes() map[string]interface{} {
	if o == nil || o.CouponAttributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CouponAttributes
}

// GetCouponAttributesOk returns a tuple with the CouponAttributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCouponAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.CouponAttributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.CouponAttributes, true
}

// HasCouponAttributes returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCouponAttributes() bool {
	if o != nil && o.CouponAttributes != nil {
		return true
	}

	return false
}

// SetCouponAttributes gets a reference to the given map[string]interface{} and assigns it to the CouponAttributes field.
func (o *UpdateCampaignTemplate) SetCouponAttributes(v map[string]interface{}) {
	o.CouponAttributes = &v
}

// GetState returns the State field value
func (o *UpdateCampaignTemplate) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// SetState sets field value
func (o *UpdateCampaignTemplate) SetState(v string) {
	o.State = v
}

// GetActiveRulesetId returns the ActiveRulesetId field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetActiveRulesetId() int32 {
	if o == nil || o.ActiveRulesetId == nil {
		var ret int32
		return ret
	}
	return *o.ActiveRulesetId
}

// GetActiveRulesetIdOk returns a tuple with the ActiveRulesetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetActiveRulesetIdOk() (int32, bool) {
	if o == nil || o.ActiveRulesetId == nil {
		var ret int32
		return ret, false
	}
	return *o.ActiveRulesetId, true
}

// HasActiveRulesetId returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasActiveRulesetId() bool {
	if o != nil && o.ActiveRulesetId != nil {
		return true
	}

	return false
}

// SetActiveRulesetId gets a reference to the given int32 and assigns it to the ActiveRulesetId field.
func (o *UpdateCampaignTemplate) SetActiveRulesetId(v int32) {
	o.ActiveRulesetId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateCampaignTemplate) SetTags(v []string) {
	o.Tags = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetFeaturesOk() ([]string, bool) {
	if o == nil || o.Features == nil {
		var ret []string
		return ret, false
	}
	return *o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *UpdateCampaignTemplate) SetFeatures(v []string) {
	o.Features = &v
}

// GetCouponSettings returns the CouponSettings field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCouponSettings() CodeGeneratorSettings {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.CouponSettings
}

// GetCouponSettingsOk returns a tuple with the CouponSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCouponSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.CouponSettings, true
}

// HasCouponSettings returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCouponSettings() bool {
	if o != nil && o.CouponSettings != nil {
		return true
	}

	return false
}

// SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.
func (o *UpdateCampaignTemplate) SetCouponSettings(v CodeGeneratorSettings) {
	o.CouponSettings = &v
}

// GetCouponReservationSettings returns the CouponReservationSettings field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCouponReservationSettings() CampaignTemplateCouponReservationSettings {
	if o == nil || o.CouponReservationSettings == nil {
		var ret CampaignTemplateCouponReservationSettings
		return ret
	}
	return *o.CouponReservationSettings
}

// GetCouponReservationSettingsOk returns a tuple with the CouponReservationSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCouponReservationSettingsOk() (CampaignTemplateCouponReservationSettings, bool) {
	if o == nil || o.CouponReservationSettings == nil {
		var ret CampaignTemplateCouponReservationSettings
		return ret, false
	}
	return *o.CouponReservationSettings, true
}

// HasCouponReservationSettings returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCouponReservationSettings() bool {
	if o != nil && o.CouponReservationSettings != nil {
		return true
	}

	return false
}

// SetCouponReservationSettings gets a reference to the given CampaignTemplateCouponReservationSettings and assigns it to the CouponReservationSettings field.
func (o *UpdateCampaignTemplate) SetCouponReservationSettings(v CampaignTemplateCouponReservationSettings) {
	o.CouponReservationSettings = &v
}

// GetReferralSettings returns the ReferralSettings field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetReferralSettings() CodeGeneratorSettings {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.ReferralSettings
}

// GetReferralSettingsOk returns a tuple with the ReferralSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetReferralSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.ReferralSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.ReferralSettings, true
}

// HasReferralSettings returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasReferralSettings() bool {
	if o != nil && o.ReferralSettings != nil {
		return true
	}

	return false
}

// SetReferralSettings gets a reference to the given CodeGeneratorSettings and assigns it to the ReferralSettings field.
func (o *UpdateCampaignTemplate) SetReferralSettings(v CodeGeneratorSettings) {
	o.ReferralSettings = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetLimits() []TemplateLimitConfig {
	if o == nil || o.Limits == nil {
		var ret []TemplateLimitConfig
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetLimitsOk() ([]TemplateLimitConfig, bool) {
	if o == nil || o.Limits == nil {
		var ret []TemplateLimitConfig
		return ret, false
	}
	return *o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []TemplateLimitConfig and assigns it to the Limits field.
func (o *UpdateCampaignTemplate) SetLimits(v []TemplateLimitConfig) {
	o.Limits = &v
}

// GetTemplateParams returns the TemplateParams field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetTemplateParams() []CampaignTemplateParams {
	if o == nil || o.TemplateParams == nil {
		var ret []CampaignTemplateParams
		return ret
	}
	return *o.TemplateParams
}

// GetTemplateParamsOk returns a tuple with the TemplateParams field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetTemplateParamsOk() ([]CampaignTemplateParams, bool) {
	if o == nil || o.TemplateParams == nil {
		var ret []CampaignTemplateParams
		return ret, false
	}
	return *o.TemplateParams, true
}

// HasTemplateParams returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasTemplateParams() bool {
	if o != nil && o.TemplateParams != nil {
		return true
	}

	return false
}

// SetTemplateParams gets a reference to the given []CampaignTemplateParams and assigns it to the TemplateParams field.
func (o *UpdateCampaignTemplate) SetTemplateParams(v []CampaignTemplateParams) {
	o.TemplateParams = &v
}

// GetApplicationsIds returns the ApplicationsIds field value
func (o *UpdateCampaignTemplate) GetApplicationsIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationsIds
}

// SetApplicationsIds sets field value
func (o *UpdateCampaignTemplate) SetApplicationsIds(v []int32) {
	o.ApplicationsIds = v
}

// GetCampaignCollections returns the CampaignCollections field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCampaignCollections() []CampaignTemplateCollection {
	if o == nil || o.CampaignCollections == nil {
		var ret []CampaignTemplateCollection
		return ret
	}
	return *o.CampaignCollections
}

// GetCampaignCollectionsOk returns a tuple with the CampaignCollections field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCampaignCollectionsOk() ([]CampaignTemplateCollection, bool) {
	if o == nil || o.CampaignCollections == nil {
		var ret []CampaignTemplateCollection
		return ret, false
	}
	return *o.CampaignCollections, true
}

// HasCampaignCollections returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCampaignCollections() bool {
	if o != nil && o.CampaignCollections != nil {
		return true
	}

	return false
}

// SetCampaignCollections gets a reference to the given []CampaignTemplateCollection and assigns it to the CampaignCollections field.
func (o *UpdateCampaignTemplate) SetCampaignCollections(v []CampaignTemplateCollection) {
	o.CampaignCollections = &v
}

// GetDefaultCampaignGroupId returns the DefaultCampaignGroupId field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetDefaultCampaignGroupId() int32 {
	if o == nil || o.DefaultCampaignGroupId == nil {
		var ret int32
		return ret
	}
	return *o.DefaultCampaignGroupId
}

// GetDefaultCampaignGroupIdOk returns a tuple with the DefaultCampaignGroupId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetDefaultCampaignGroupIdOk() (int32, bool) {
	if o == nil || o.DefaultCampaignGroupId == nil {
		var ret int32
		return ret, false
	}
	return *o.DefaultCampaignGroupId, true
}

// HasDefaultCampaignGroupId returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasDefaultCampaignGroupId() bool {
	if o != nil && o.DefaultCampaignGroupId != nil {
		return true
	}

	return false
}

// SetDefaultCampaignGroupId gets a reference to the given int32 and assigns it to the DefaultCampaignGroupId field.
func (o *UpdateCampaignTemplate) SetDefaultCampaignGroupId(v int32) {
	o.DefaultCampaignGroupId = &v
}

// GetCampaignType returns the CampaignType field value if set, zero value otherwise.
func (o *UpdateCampaignTemplate) GetCampaignType() string {
	if o == nil || o.CampaignType == nil {
		var ret string
		return ret
	}
	return *o.CampaignType
}

// GetCampaignTypeOk returns a tuple with the CampaignType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignTemplate) GetCampaignTypeOk() (string, bool) {
	if o == nil || o.CampaignType == nil {
		var ret string
		return ret, false
	}
	return *o.CampaignType, true
}

// HasCampaignType returns a boolean if a field has been set.
func (o *UpdateCampaignTemplate) HasCampaignType() bool {
	if o != nil && o.CampaignType != nil {
		return true
	}

	return false
}

// SetCampaignType gets a reference to the given string and assigns it to the CampaignType field.
func (o *UpdateCampaignTemplate) SetCampaignType(v string) {
	o.CampaignType = &v
}

type NullableUpdateCampaignTemplate struct {
	Value        UpdateCampaignTemplate
	ExplicitNull bool
}

func (v NullableUpdateCampaignTemplate) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateCampaignTemplate) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
