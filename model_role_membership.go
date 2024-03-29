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

// RoleMembership struct for RoleMembership
type RoleMembership struct {
	// ID of role.
	RoleID int32 `json:"RoleID"`
	// ID of User.
	UserID int32 `json:"UserID"`
}

// GetRoleID returns the RoleID field value
func (o *RoleMembership) GetRoleID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RoleID
}

// SetRoleID sets field value
func (o *RoleMembership) SetRoleID(v int32) {
	o.RoleID = v
}

// GetUserID returns the UserID field value
func (o *RoleMembership) GetUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserID
}

// SetUserID sets field value
func (o *RoleMembership) SetUserID(v int32) {
	o.UserID = v
}

type NullableRoleMembership struct {
	Value        RoleMembership
	ExplicitNull bool
}

func (v NullableRoleMembership) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRoleMembership) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
