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

// CreateTemplateCampaign struct for CreateTemplateCampaign
type CreateTemplateCampaign struct {
	// A user-facing name for this campaign.
	Name string `json:"name"`
	// A detailed description of the campaign.
	Description *string `json:"description,omitempty"`
	// The ID of the Campaign Template which will be used in order to create the Campaign.
	TemplateId int32 `json:"templateId"`
	// Custom Campaign Attributes. If the Campaign Template defines the same values, they will be overridden.
	CampaignAttributesOverrides *map[string]map[string]interface{} `json:"campaignAttributesOverrides,omitempty"`
	// Actual values to replace the template placeholder values in the Ruleset bindings. Values for all Template Parameters must be provided.
	TemplateParamValues *[]Binding `json:"templateParamValues,omitempty"`
	// Limits for this Campaign. If the Campaign Template or Application define default values for the same limits, they will be overridden.
	LimitOverrides *[]LimitConfig `json:"limitOverrides,omitempty"`
	// The IDs of the [campaign groups](https://docs.talon.one/docs/product/account/account-settings/managing-campaign-groups) this campaign belongs to. 
	CampaignGroups *[]int32 `json:"campaignGroups,omitempty"`
	// A list of tags for the campaign. If the campaign template has tags, they will be overridden by this list.
	Tags *[]string `json:"tags,omitempty"`
	// The ID of the campaign evaluation group the campaign belongs to.
	EvaluationGroupId *int32 `json:"evaluationGroupId,omitempty"`
	// A list of store IDs that are linked to the campaign.  **Note:** Campaigns with linked store IDs will only be evaluated when there is a [customer session update](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) that references a linked store. 
	LinkedStoreIds *[]int32 `json:"linkedStoreIds,omitempty"`
}

// GetName returns the Name field value
func (o *CreateTemplateCampaign) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CreateTemplateCampaign) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateTemplateCampaign) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateCampaign) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateTemplateCampaign) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateTemplateCampaign) SetDescription(v string) {
	o.Description = &v
}

// GetTemplateId returns the TemplateId field value
func (o *CreateTemplateCampaign) GetTemplateId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TemplateId
}

// SetTemplateId sets field value
func (o *CreateTemplateCampaign) SetTemplateId(v int32) {
	o.TemplateId = v
}

// GetCampaignAttributesOverrides returns the CampaignAttributesOverrides field value if set, zero value otherwise.
func (o *CreateTemplateCampaign) GetCampaignAttributesOverrides() map[string]map[string]interface{} {
	if o == nil || o.CampaignAttributesOverrides == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.CampaignAttributesOverrides
}

// GetCampaignAttributesOverridesOk returns a tuple with the CampaignAttributesOverrides field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateCampaign) GetCampaignAttributesOverridesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.CampaignAttributesOverrides == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.CampaignAttributesOverrides, true
}

// HasCampaignAttributesOverrides returns a boolean if a field has been set.
func (o *CreateTemplateCampaign) HasCampaignAttributesOverrides() bool {
	if o != nil && o.CampaignAttributesOverrides != nil {
		return true
	}

	return false
}

// SetCampaignAttributesOverrides gets a reference to the given map[string]map[string]interface{} and assigns it to the CampaignAttributesOverrides field.
func (o *CreateTemplateCampaign) SetCampaignAttributesOverrides(v map[string]map[string]interface{}) {
	o.CampaignAttributesOverrides = &v
}

// GetTemplateParamValues returns the TemplateParamValues field value if set, zero value otherwise.
func (o *CreateTemplateCampaign) GetTemplateParamValues() []Binding {
	if o == nil || o.TemplateParamValues == nil {
		var ret []Binding
		return ret
	}
	return *o.TemplateParamValues
}

// GetTemplateParamValuesOk returns a tuple with the TemplateParamValues field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateCampaign) GetTemplateParamValuesOk() ([]Binding, bool) {
	if o == nil || o.TemplateParamValues == nil {
		var ret []Binding
		return ret, false
	}
	return *o.TemplateParamValues, true
}

// HasTemplateParamValues returns a boolean if a field has been set.
func (o *CreateTemplateCampaign) HasTemplateParamValues() bool {
	if o != nil && o.TemplateParamValues != nil {
		return true
	}

	return false
}

// SetTemplateParamValues gets a reference to the given []Binding and assigns it to the TemplateParamValues field.
func (o *CreateTemplateCampaign) SetTemplateParamValues(v []Binding) {
	o.TemplateParamValues = &v
}

// GetLimitOverrides returns the LimitOverrides field value if set, zero value otherwise.
func (o *CreateTemplateCampaign) GetLimitOverrides() []LimitConfig {
	if o == nil || o.LimitOverrides == nil {
		var ret []LimitConfig
		return ret
	}
	return *o.LimitOverrides
}

// GetLimitOverridesOk returns a tuple with the LimitOverrides field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateCampaign) GetLimitOverridesOk() ([]LimitConfig, bool) {
	if o == nil || o.LimitOverrides == nil {
		var ret []LimitConfig
		return ret, false
	}
	return *o.LimitOverrides, true
}

// HasLimitOverrides returns a boolean if a field has been set.
func (o *CreateTemplateCampaign) HasLimitOverrides() bool {
	if o != nil && o.LimitOverrides != nil {
		return true
	}

	return false
}

// SetLimitOverrides gets a reference to the given []LimitConfig and assigns it to the LimitOverrides field.
func (o *CreateTemplateCampaign) SetLimitOverrides(v []LimitConfig) {
	o.LimitOverrides = &v
}

// GetCampaignGroups returns the CampaignGroups field value if set, zero value otherwise.
func (o *CreateTemplateCampaign) GetCampaignGroups() []int32 {
	if o == nil || o.CampaignGroups == nil {
		var ret []int32
		return ret
	}
	return *o.CampaignGroups
}

// GetCampaignGroupsOk returns a tuple with the CampaignGroups field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateCampaign) GetCampaignGroupsOk() ([]int32, bool) {
	if o == nil || o.CampaignGroups == nil {
		var ret []int32
		return ret, false
	}
	return *o.CampaignGroups, true
}

// HasCampaignGroups returns a boolean if a field has been set.
func (o *CreateTemplateCampaign) HasCampaignGroups() bool {
	if o != nil && o.CampaignGroups != nil {
		return true
	}

	return false
}

// SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.
func (o *CreateTemplateCampaign) SetCampaignGroups(v []int32) {
	o.CampaignGroups = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateTemplateCampaign) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateCampaign) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateTemplateCampaign) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateTemplateCampaign) SetTags(v []string) {
	o.Tags = &v
}

// GetEvaluationGroupId returns the EvaluationGroupId field value if set, zero value otherwise.
func (o *CreateTemplateCampaign) GetEvaluationGroupId() int32 {
	if o == nil || o.EvaluationGroupId == nil {
		var ret int32
		return ret
	}
	return *o.EvaluationGroupId
}

// GetEvaluationGroupIdOk returns a tuple with the EvaluationGroupId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateCampaign) GetEvaluationGroupIdOk() (int32, bool) {
	if o == nil || o.EvaluationGroupId == nil {
		var ret int32
		return ret, false
	}
	return *o.EvaluationGroupId, true
}

// HasEvaluationGroupId returns a boolean if a field has been set.
func (o *CreateTemplateCampaign) HasEvaluationGroupId() bool {
	if o != nil && o.EvaluationGroupId != nil {
		return true
	}

	return false
}

// SetEvaluationGroupId gets a reference to the given int32 and assigns it to the EvaluationGroupId field.
func (o *CreateTemplateCampaign) SetEvaluationGroupId(v int32) {
	o.EvaluationGroupId = &v
}

// GetLinkedStoreIds returns the LinkedStoreIds field value if set, zero value otherwise.
func (o *CreateTemplateCampaign) GetLinkedStoreIds() []int32 {
	if o == nil || o.LinkedStoreIds == nil {
		var ret []int32
		return ret
	}
	return *o.LinkedStoreIds
}

// GetLinkedStoreIdsOk returns a tuple with the LinkedStoreIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateTemplateCampaign) GetLinkedStoreIdsOk() ([]int32, bool) {
	if o == nil || o.LinkedStoreIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.LinkedStoreIds, true
}

// HasLinkedStoreIds returns a boolean if a field has been set.
func (o *CreateTemplateCampaign) HasLinkedStoreIds() bool {
	if o != nil && o.LinkedStoreIds != nil {
		return true
	}

	return false
}

// SetLinkedStoreIds gets a reference to the given []int32 and assigns it to the LinkedStoreIds field.
func (o *CreateTemplateCampaign) SetLinkedStoreIds(v []int32) {
	o.LinkedStoreIds = &v
}

type NullableCreateTemplateCampaign struct {
	Value CreateTemplateCampaign
	ExplicitNull bool
}

func (v NullableCreateTemplateCampaign) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateTemplateCampaign) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
