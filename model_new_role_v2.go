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

// NewRoleV2 struct for NewRoleV2
type NewRoleV2 struct {
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description string             `json:"description"`
	Permissions *RoleV2Permissions `json:"permissions,omitempty"`
	// A list of user IDs the role is assigned to.
	Members *[]int32 `json:"members,omitempty"`
}

// GetName returns the Name field value
func (o *NewRoleV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewRoleV2) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *NewRoleV2) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *NewRoleV2) SetDescription(v string) {
	o.Description = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *NewRoleV2) GetPermissions() RoleV2Permissions {
	if o == nil || o.Permissions == nil {
		var ret RoleV2Permissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRoleV2) GetPermissionsOk() (RoleV2Permissions, bool) {
	if o == nil || o.Permissions == nil {
		var ret RoleV2Permissions
		return ret, false
	}
	return *o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *NewRoleV2) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given RoleV2Permissions and assigns it to the Permissions field.
func (o *NewRoleV2) SetPermissions(v RoleV2Permissions) {
	o.Permissions = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *NewRoleV2) GetMembers() []int32 {
	if o == nil || o.Members == nil {
		var ret []int32
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewRoleV2) GetMembersOk() ([]int32, bool) {
	if o == nil || o.Members == nil {
		var ret []int32
		return ret, false
	}
	return *o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *NewRoleV2) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []int32 and assigns it to the Members field.
func (o *NewRoleV2) SetMembers(v []int32) {
	o.Members = &v
}

type NullableNewRoleV2 struct {
	Value        NewRoleV2
	ExplicitNull bool
}

func (v NullableNewRoleV2) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewRoleV2) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
