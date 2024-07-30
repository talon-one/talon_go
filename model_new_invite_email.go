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

// NewInviteEmail struct for NewInviteEmail
type NewInviteEmail struct {
	// Email address of the user.
	Email string `json:"email"`
	// Invitation token of the user.
	Token string `json:"token"`
}

// GetEmail returns the Email field value
func (o *NewInviteEmail) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// SetEmail sets field value
func (o *NewInviteEmail) SetEmail(v string) {
	o.Email = v
}

// GetToken returns the Token field value
func (o *NewInviteEmail) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// SetToken sets field value
func (o *NewInviteEmail) SetToken(v string) {
	o.Token = v
}

type NullableNewInviteEmail struct {
	Value NewInviteEmail
	ExplicitNull bool
}

func (v NullableNewInviteEmail) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewInviteEmail) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
