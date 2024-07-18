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

// RoleV2PermissionSet struct for RoleV2PermissionSet
type RoleV2PermissionSet struct {
	// Name of the permission set.
	Name string `json:"name"`
	// List of logical operations in the permission set. Each logical operation must be shown under the `x-permission` tag on an endpoint level. 
	LogicalOperations []string `json:"logicalOperations"`
}

// GetName returns the Name field value
func (o *RoleV2PermissionSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *RoleV2PermissionSet) SetName(v string) {
	o.Name = v
}

// GetLogicalOperations returns the LogicalOperations field value
func (o *RoleV2PermissionSet) GetLogicalOperations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LogicalOperations
}

// SetLogicalOperations sets field value
func (o *RoleV2PermissionSet) SetLogicalOperations(v []string) {
	o.LogicalOperations = v
}

type NullableRoleV2PermissionSet struct {
	Value RoleV2PermissionSet
	ExplicitNull bool
}

func (v NullableRoleV2PermissionSet) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRoleV2PermissionSet) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
