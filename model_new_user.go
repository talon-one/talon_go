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

// NewUser
type NewUser struct {
	// The email address associated with your account.
	Email string `json:"email"`
	// The password for your account.
	Password string `json:"password"`
	// Your name.
	Name        *string `json:"name,omitempty"`
	InviteToken string  `json:"inviteToken"`
}

// GetEmail returns the Email field value
func (o *NewUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// SetEmail sets field value
func (o *NewUser) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *NewUser) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// SetPassword sets field value
func (o *NewUser) SetPassword(v string) {
	o.Password = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NewUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NewUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NewUser) SetName(v string) {
	o.Name = &v
}

// GetInviteToken returns the InviteToken field value
func (o *NewUser) GetInviteToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteToken
}

// SetInviteToken sets field value
func (o *NewUser) SetInviteToken(v string) {
	o.InviteToken = v
}

type NullableNewUser struct {
	Value        NewUser
	ExplicitNull bool
}

func (v NullableNewUser) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewUser) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
