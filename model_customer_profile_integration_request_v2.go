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

// CustomerProfileIntegrationRequestV2 struct for CustomerProfileIntegrationRequestV2
type CustomerProfileIntegrationRequestV2 struct {
	// Arbitrary properties associated with this item.
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	// When using the `dry` query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them. 
	EvaluableCampaignIds *[]int32 `json:"evaluableCampaignIds,omitempty"`
	AudiencesChanges *ProfileAudiencesChanges `json:"audiencesChanges,omitempty"`
	// Extends the response with the chosen data entities. Use this property to get as much data as you need in one _Update customer profile_ request instead of sending extra requests to other endpoints. 
	ResponseContent *[]string `json:"responseContent,omitempty"`
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerProfileIntegrationRequestV2) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileIntegrationRequestV2) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerProfileIntegrationRequestV2) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *CustomerProfileIntegrationRequestV2) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

// GetEvaluableCampaignIds returns the EvaluableCampaignIds field value if set, zero value otherwise.
func (o *CustomerProfileIntegrationRequestV2) GetEvaluableCampaignIds() []int32 {
	if o == nil || o.EvaluableCampaignIds == nil {
		var ret []int32
		return ret
	}
	return *o.EvaluableCampaignIds
}

// GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileIntegrationRequestV2) GetEvaluableCampaignIdsOk() ([]int32, bool) {
	if o == nil || o.EvaluableCampaignIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.EvaluableCampaignIds, true
}

// HasEvaluableCampaignIds returns a boolean if a field has been set.
func (o *CustomerProfileIntegrationRequestV2) HasEvaluableCampaignIds() bool {
	if o != nil && o.EvaluableCampaignIds != nil {
		return true
	}

	return false
}

// SetEvaluableCampaignIds gets a reference to the given []int32 and assigns it to the EvaluableCampaignIds field.
func (o *CustomerProfileIntegrationRequestV2) SetEvaluableCampaignIds(v []int32) {
	o.EvaluableCampaignIds = &v
}

// GetAudiencesChanges returns the AudiencesChanges field value if set, zero value otherwise.
func (o *CustomerProfileIntegrationRequestV2) GetAudiencesChanges() ProfileAudiencesChanges {
	if o == nil || o.AudiencesChanges == nil {
		var ret ProfileAudiencesChanges
		return ret
	}
	return *o.AudiencesChanges
}

// GetAudiencesChangesOk returns a tuple with the AudiencesChanges field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileIntegrationRequestV2) GetAudiencesChangesOk() (ProfileAudiencesChanges, bool) {
	if o == nil || o.AudiencesChanges == nil {
		var ret ProfileAudiencesChanges
		return ret, false
	}
	return *o.AudiencesChanges, true
}

// HasAudiencesChanges returns a boolean if a field has been set.
func (o *CustomerProfileIntegrationRequestV2) HasAudiencesChanges() bool {
	if o != nil && o.AudiencesChanges != nil {
		return true
	}

	return false
}

// SetAudiencesChanges gets a reference to the given ProfileAudiencesChanges and assigns it to the AudiencesChanges field.
func (o *CustomerProfileIntegrationRequestV2) SetAudiencesChanges(v ProfileAudiencesChanges) {
	o.AudiencesChanges = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *CustomerProfileIntegrationRequestV2) GetResponseContent() []string {
	if o == nil || o.ResponseContent == nil {
		var ret []string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileIntegrationRequestV2) GetResponseContentOk() ([]string, bool) {
	if o == nil || o.ResponseContent == nil {
		var ret []string
		return ret, false
	}
	return *o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *CustomerProfileIntegrationRequestV2) HasResponseContent() bool {
	if o != nil && o.ResponseContent != nil {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given []string and assigns it to the ResponseContent field.
func (o *CustomerProfileIntegrationRequestV2) SetResponseContent(v []string) {
	o.ResponseContent = &v
}

type NullableCustomerProfileIntegrationRequestV2 struct {
	Value CustomerProfileIntegrationRequestV2
	ExplicitNull bool
}

func (v NullableCustomerProfileIntegrationRequestV2) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomerProfileIntegrationRequestV2) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
