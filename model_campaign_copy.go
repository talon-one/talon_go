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

// CampaignCopy struct for CampaignCopy
type CampaignCopy struct {
	// A list of tags for the campaign.
	Tags *[]string `json:"tags,omitempty"`
	// Application IDs of the applications to which a campaign should be copied to.
	ApplicationIds []int32 `json:"applicationIds"`
	// A detailed description of the campaign.
	Description *string `json:"description,omitempty"`
	// Timestamp when the campaign will become inactive.
	EndTime *time.Time `json:"endTime,omitempty"`
	// The ID of the campaign evaluation group the campaign belongs to.
	EvaluationGroupId *int32 `json:"evaluationGroupId,omitempty"`
	// Name of the copied campaign (Defaults to \"Copy of original campaign name\").
	Name *string `json:"name,omitempty"`
	// Timestamp when the campaign will become active.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CampaignCopy) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCopy) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CampaignCopy) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CampaignCopy) SetTags(v []string) {
	o.Tags = &v
}

// GetApplicationIds returns the ApplicationIds field value
func (o *CampaignCopy) GetApplicationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationIds
}

// SetApplicationIds sets field value
func (o *CampaignCopy) SetApplicationIds(v []int32) {
	o.ApplicationIds = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignCopy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCopy) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignCopy) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignCopy) SetDescription(v string) {
	o.Description = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *CampaignCopy) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCopy) GetEndTimeOk() (time.Time, bool) {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *CampaignCopy) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *CampaignCopy) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetEvaluationGroupId returns the EvaluationGroupId field value if set, zero value otherwise.
func (o *CampaignCopy) GetEvaluationGroupId() int32 {
	if o == nil || o.EvaluationGroupId == nil {
		var ret int32
		return ret
	}
	return *o.EvaluationGroupId
}

// GetEvaluationGroupIdOk returns a tuple with the EvaluationGroupId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCopy) GetEvaluationGroupIdOk() (int32, bool) {
	if o == nil || o.EvaluationGroupId == nil {
		var ret int32
		return ret, false
	}
	return *o.EvaluationGroupId, true
}

// HasEvaluationGroupId returns a boolean if a field has been set.
func (o *CampaignCopy) HasEvaluationGroupId() bool {
	if o != nil && o.EvaluationGroupId != nil {
		return true
	}

	return false
}

// SetEvaluationGroupId gets a reference to the given int32 and assigns it to the EvaluationGroupId field.
func (o *CampaignCopy) SetEvaluationGroupId(v int32) {
	o.EvaluationGroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CampaignCopy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCopy) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CampaignCopy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CampaignCopy) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CampaignCopy) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCopy) GetStartTimeOk() (time.Time, bool) {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CampaignCopy) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *CampaignCopy) SetStartTime(v time.Time) {
	o.StartTime = &v
}

type NullableCampaignCopy struct {
	Value CampaignCopy
	ExplicitNull bool
}

func (v NullableCampaignCopy) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignCopy) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
