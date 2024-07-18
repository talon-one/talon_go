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

// CampaignSetBranchNode struct for CampaignSetBranchNode
type CampaignSetBranchNode struct {
	// Indicates the node type.
	Type string `json:"type"`
	// Name of the set.
	Name string `json:"name"`
	// An indicator of how the set operates on its elements.
	Operator string `json:"operator"`
	// Child elements of this set.
	Elements []CampaignSetNode `json:"elements"`
	// The ID of the campaign set.
	GroupId int32 `json:"groupId"`
	// An indicator of whether the campaign set is locked for modification.
	Locked bool `json:"locked"`
	// A description of the campaign set.
	Description *string `json:"description,omitempty"`
	// The mode by which campaigns in the campaign evaluation group are evaluated.
	EvaluationMode string `json:"evaluationMode"`
	// The evaluation scope of the campaign evaluation group.
	EvaluationScope string `json:"evaluationScope"`
}

// GetType returns the Type field value
func (o *CampaignSetBranchNode) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *CampaignSetBranchNode) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *CampaignSetBranchNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CampaignSetBranchNode) SetName(v string) {
	o.Name = v
}

// GetOperator returns the Operator field value
func (o *CampaignSetBranchNode) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// SetOperator sets field value
func (o *CampaignSetBranchNode) SetOperator(v string) {
	o.Operator = v
}

// GetElements returns the Elements field value
func (o *CampaignSetBranchNode) GetElements() []CampaignSetNode {
	if o == nil {
		var ret []CampaignSetNode
		return ret
	}

	return o.Elements
}

// SetElements sets field value
func (o *CampaignSetBranchNode) SetElements(v []CampaignSetNode) {
	o.Elements = v
}

// GetGroupId returns the GroupId field value
func (o *CampaignSetBranchNode) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// SetGroupId sets field value
func (o *CampaignSetBranchNode) SetGroupId(v int32) {
	o.GroupId = v
}

// GetLocked returns the Locked field value
func (o *CampaignSetBranchNode) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// SetLocked sets field value
func (o *CampaignSetBranchNode) SetLocked(v bool) {
	o.Locked = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignSetBranchNode) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSetBranchNode) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignSetBranchNode) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignSetBranchNode) SetDescription(v string) {
	o.Description = &v
}

// GetEvaluationMode returns the EvaluationMode field value
func (o *CampaignSetBranchNode) GetEvaluationMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationMode
}

// SetEvaluationMode sets field value
func (o *CampaignSetBranchNode) SetEvaluationMode(v string) {
	o.EvaluationMode = v
}

// GetEvaluationScope returns the EvaluationScope field value
func (o *CampaignSetBranchNode) GetEvaluationScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationScope
}

// SetEvaluationScope sets field value
func (o *CampaignSetBranchNode) SetEvaluationScope(v string) {
	o.EvaluationScope = v
}

type NullableCampaignSetBranchNode struct {
	Value CampaignSetBranchNode
	ExplicitNull bool
}

func (v NullableCampaignSetBranchNode) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignSetBranchNode) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
