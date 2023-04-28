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

// LoyaltyDashboardPointsBreakdown struct for LoyaltyDashboardPointsBreakdown
type LoyaltyDashboardPointsBreakdown struct {
	CreatedManually      float32 `json:"createdManually"`
	CreatedViaRuleEngine float32 `json:"createdViaRuleEngine"`
}

// GetCreatedManually returns the CreatedManually field value
func (o *LoyaltyDashboardPointsBreakdown) GetCreatedManually() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedManually
}

// SetCreatedManually sets field value
func (o *LoyaltyDashboardPointsBreakdown) SetCreatedManually(v float32) {
	o.CreatedManually = v
}

// GetCreatedViaRuleEngine returns the CreatedViaRuleEngine field value
func (o *LoyaltyDashboardPointsBreakdown) GetCreatedViaRuleEngine() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedViaRuleEngine
}

// SetCreatedViaRuleEngine sets field value
func (o *LoyaltyDashboardPointsBreakdown) SetCreatedViaRuleEngine(v float32) {
	o.CreatedViaRuleEngine = v
}

type NullableLoyaltyDashboardPointsBreakdown struct {
	Value        LoyaltyDashboardPointsBreakdown
	ExplicitNull bool
}

func (v NullableLoyaltyDashboardPointsBreakdown) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyDashboardPointsBreakdown) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
