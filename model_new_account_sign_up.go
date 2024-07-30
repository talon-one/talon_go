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

// NewAccountSignUp 
type NewAccountSignUp struct {
	// The email address associated with the user profile.
	Email string `json:"email"`
	// The password for your account.
	Password string `json:"password"`
	CompanyName string `json:"companyName"`
}

// GetEmail returns the Email field value
func (o *NewAccountSignUp) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// SetEmail sets field value
func (o *NewAccountSignUp) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *NewAccountSignUp) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// SetPassword sets field value
func (o *NewAccountSignUp) SetPassword(v string) {
	o.Password = v
}

// GetCompanyName returns the CompanyName field value
func (o *NewAccountSignUp) GetCompanyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyName
}

// SetCompanyName sets field value
func (o *NewAccountSignUp) SetCompanyName(v string) {
	o.CompanyName = v
}

type NullableNewAccountSignUp struct {
	Value NewAccountSignUp
	ExplicitNull bool
}

func (v NullableNewAccountSignUp) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewAccountSignUp) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
