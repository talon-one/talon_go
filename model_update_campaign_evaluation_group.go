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

// UpdateCampaignEvaluationGroup struct for UpdateCampaignEvaluationGroup
type UpdateCampaignEvaluationGroup struct {
	// A description of the campaign evaluation group.
	Description *string `json:"description,omitempty"`
	// The mode by which campaigns in the campaign evaluation group are evaluated.
	EvaluationMode string `json:"evaluationMode"`
	// The evaluation scope of the campaign evaluation group.
	EvaluationScope string `json:"evaluationScope"`
	// An indicator of whether the campaign evaluation group is locked for modification.
	Locked bool `json:"locked"`
	// The name of the campaign evaluation group.
	Name string `json:"name"`
	// The ID of the parent group that contains the campaign evaluation group.
	ParentId int32 `json:"parentId"`
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCampaignEvaluationGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCampaignEvaluationGroup) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateCampaignEvaluationGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCampaignEvaluationGroup) SetDescription(v string) {
	o.Description = &v
}

// GetEvaluationMode returns the EvaluationMode field value
func (o *UpdateCampaignEvaluationGroup) GetEvaluationMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationMode
}

// SetEvaluationMode sets field value
func (o *UpdateCampaignEvaluationGroup) SetEvaluationMode(v string) {
	o.EvaluationMode = v
}

// GetEvaluationScope returns the EvaluationScope field value
func (o *UpdateCampaignEvaluationGroup) GetEvaluationScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationScope
}

// SetEvaluationScope sets field value
func (o *UpdateCampaignEvaluationGroup) SetEvaluationScope(v string) {
	o.EvaluationScope = v
}

// GetLocked returns the Locked field value
func (o *UpdateCampaignEvaluationGroup) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// SetLocked sets field value
func (o *UpdateCampaignEvaluationGroup) SetLocked(v bool) {
	o.Locked = v
}

// GetName returns the Name field value
func (o *UpdateCampaignEvaluationGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *UpdateCampaignEvaluationGroup) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value
func (o *UpdateCampaignEvaluationGroup) GetParentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentId
}

// SetParentId sets field value
func (o *UpdateCampaignEvaluationGroup) SetParentId(v int32) {
	o.ParentId = v
}

type NullableUpdateCampaignEvaluationGroup struct {
	Value UpdateCampaignEvaluationGroup
	ExplicitNull bool
}

func (v NullableUpdateCampaignEvaluationGroup) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateCampaignEvaluationGroup) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
