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

// NewPassword struct for NewPassword
type NewPassword struct {
	// The new password for your account.
	Password string `json:"password"`
	ResetToken string `json:"resetToken"`
}

// GetPassword returns the Password field value
func (o *NewPassword) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// SetPassword sets field value
func (o *NewPassword) SetPassword(v string) {
	o.Password = v
}

// GetResetToken returns the ResetToken field value
func (o *NewPassword) GetResetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetToken
}

// SetResetToken sets field value
func (o *NewPassword) SetResetToken(v string) {
	o.ResetToken = v
}

type NullableNewPassword struct {
	Value NewPassword
	ExplicitNull bool
}

func (v NullableNewPassword) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewPassword) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
