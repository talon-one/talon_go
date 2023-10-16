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

// TwoFaConfig struct for TwoFaConfig
type TwoFaConfig struct {
	// An indication of whether two-factor authentication is enabled for the account.
	Enabled bool `json:"enabled"`
	// Can be `true` or `false`. - `true`: Two-factor authentication is required each time a user signs in to their Talon.One account. - `false`: Two-factor authentication is only required when a user signs in to their Talon.One account on a new device, and every 30 days after that.
	RequireEverySignIn *bool `json:"requireEverySignIn,omitempty"`
}

// GetEnabled returns the Enabled field value
func (o *TwoFaConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// SetEnabled sets field value
func (o *TwoFaConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRequireEverySignIn returns the RequireEverySignIn field value if set, zero value otherwise.
func (o *TwoFaConfig) GetRequireEverySignIn() bool {
	if o == nil || o.RequireEverySignIn == nil {
		var ret bool
		return ret
	}
	return *o.RequireEverySignIn
}

// GetRequireEverySignInOk returns a tuple with the RequireEverySignIn field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TwoFaConfig) GetRequireEverySignInOk() (bool, bool) {
	if o == nil || o.RequireEverySignIn == nil {
		var ret bool
		return ret, false
	}
	return *o.RequireEverySignIn, true
}

// HasRequireEverySignIn returns a boolean if a field has been set.
func (o *TwoFaConfig) HasRequireEverySignIn() bool {
	if o != nil && o.RequireEverySignIn != nil {
		return true
	}

	return false
}

// SetRequireEverySignIn gets a reference to the given bool and assigns it to the RequireEverySignIn field.
func (o *TwoFaConfig) SetRequireEverySignIn(v bool) {
	o.RequireEverySignIn = &v
}

type NullableTwoFaConfig struct {
	Value        TwoFaConfig
	ExplicitNull bool
}

func (v NullableTwoFaConfig) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTwoFaConfig) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
