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

// RejectReferralEffectProps The properties specific to the \"rejectReferral\" effect. This gets triggered whenever the referral code was rejected. See rejectionReason for more info on why.
type RejectReferralEffectProps struct {
	// The referral code that was rejected.
	Value string `json:"value"`
	// The reason why this referral code was rejected.
	RejectionReason string `json:"rejectionReason"`
	// The index of the condition that caused the rejection of the referral.
	ConditionIndex *int32 `json:"conditionIndex,omitempty"`
	// The index of the effect that caused the rejection of the referral.
	EffectIndex *int32 `json:"effectIndex,omitempty"`
	// More details about the failure.
	Details *string `json:"details,omitempty"`
	// The reason why the campaign was not applied.
	CampaignExclusionReason *string `json:"campaignExclusionReason,omitempty"`
}

// GetValue returns the Value field value
func (o *RejectReferralEffectProps) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *RejectReferralEffectProps) SetValue(v string) {
	o.Value = v
}

// GetRejectionReason returns the RejectionReason field value
func (o *RejectReferralEffectProps) GetRejectionReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RejectionReason
}

// SetRejectionReason sets field value
func (o *RejectReferralEffectProps) SetRejectionReason(v string) {
	o.RejectionReason = v
}

// GetConditionIndex returns the ConditionIndex field value if set, zero value otherwise.
func (o *RejectReferralEffectProps) GetConditionIndex() int32 {
	if o == nil || o.ConditionIndex == nil {
		var ret int32
		return ret
	}
	return *o.ConditionIndex
}

// GetConditionIndexOk returns a tuple with the ConditionIndex field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RejectReferralEffectProps) GetConditionIndexOk() (int32, bool) {
	if o == nil || o.ConditionIndex == nil {
		var ret int32
		return ret, false
	}
	return *o.ConditionIndex, true
}

// HasConditionIndex returns a boolean if a field has been set.
func (o *RejectReferralEffectProps) HasConditionIndex() bool {
	if o != nil && o.ConditionIndex != nil {
		return true
	}

	return false
}

// SetConditionIndex gets a reference to the given int32 and assigns it to the ConditionIndex field.
func (o *RejectReferralEffectProps) SetConditionIndex(v int32) {
	o.ConditionIndex = &v
}

// GetEffectIndex returns the EffectIndex field value if set, zero value otherwise.
func (o *RejectReferralEffectProps) GetEffectIndex() int32 {
	if o == nil || o.EffectIndex == nil {
		var ret int32
		return ret
	}
	return *o.EffectIndex
}

// GetEffectIndexOk returns a tuple with the EffectIndex field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RejectReferralEffectProps) GetEffectIndexOk() (int32, bool) {
	if o == nil || o.EffectIndex == nil {
		var ret int32
		return ret, false
	}
	return *o.EffectIndex, true
}

// HasEffectIndex returns a boolean if a field has been set.
func (o *RejectReferralEffectProps) HasEffectIndex() bool {
	if o != nil && o.EffectIndex != nil {
		return true
	}

	return false
}

// SetEffectIndex gets a reference to the given int32 and assigns it to the EffectIndex field.
func (o *RejectReferralEffectProps) SetEffectIndex(v int32) {
	o.EffectIndex = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RejectReferralEffectProps) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RejectReferralEffectProps) GetDetailsOk() (string, bool) {
	if o == nil || o.Details == nil {
		var ret string
		return ret, false
	}
	return *o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RejectReferralEffectProps) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *RejectReferralEffectProps) SetDetails(v string) {
	o.Details = &v
}

// GetCampaignExclusionReason returns the CampaignExclusionReason field value if set, zero value otherwise.
func (o *RejectReferralEffectProps) GetCampaignExclusionReason() string {
	if o == nil || o.CampaignExclusionReason == nil {
		var ret string
		return ret
	}
	return *o.CampaignExclusionReason
}

// GetCampaignExclusionReasonOk returns a tuple with the CampaignExclusionReason field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RejectReferralEffectProps) GetCampaignExclusionReasonOk() (string, bool) {
	if o == nil || o.CampaignExclusionReason == nil {
		var ret string
		return ret, false
	}
	return *o.CampaignExclusionReason, true
}

// HasCampaignExclusionReason returns a boolean if a field has been set.
func (o *RejectReferralEffectProps) HasCampaignExclusionReason() bool {
	if o != nil && o.CampaignExclusionReason != nil {
		return true
	}

	return false
}

// SetCampaignExclusionReason gets a reference to the given string and assigns it to the CampaignExclusionReason field.
func (o *RejectReferralEffectProps) SetCampaignExclusionReason(v string) {
	o.CampaignExclusionReason = &v
}

type NullableRejectReferralEffectProps struct {
	Value RejectReferralEffectProps
	ExplicitNull bool
}

func (v NullableRejectReferralEffectProps) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRejectReferralEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
