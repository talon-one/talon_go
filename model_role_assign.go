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

// RoleAssign struct for RoleAssign
type RoleAssign struct {
	// An array of user IDs.
	Users []int32 `json:"users"`
	// An array of role IDs.
	Roles []int32 `json:"roles"`
}

// GetUsers returns the Users field value
func (o *RoleAssign) GetUsers() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Users
}

// SetUsers sets field value
func (o *RoleAssign) SetUsers(v []int32) {
	o.Users = v
}

// GetRoles returns the Roles field value
func (o *RoleAssign) GetRoles() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Roles
}

// SetRoles sets field value
func (o *RoleAssign) SetRoles(v []int32) {
	o.Roles = v
}

type NullableRoleAssign struct {
	Value        RoleAssign
	ExplicitNull bool
}

func (v NullableRoleAssign) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRoleAssign) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
