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

// RollbackReferralEffectProps The properties specific to the \"rollbackReferral\" effect. This gets triggered whenever previously closed session is now cancelled and a referral redemption was cancelled on our internal usage limit counters.
type RollbackReferralEffectProps struct {
	// The referral code whose usage has been rolled back.
	Value string `json:"value"`
}

// GetValue returns the Value field value
func (o *RollbackReferralEffectProps) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *RollbackReferralEffectProps) SetValue(v string) {
	o.Value = v
}

type NullableRollbackReferralEffectProps struct {
	Value RollbackReferralEffectProps
	ExplicitNull bool
}

func (v NullableRollbackReferralEffectProps) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRollbackReferralEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
