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

// UpdateRole struct for UpdateRole
type UpdateRole struct {
	// Name of the role.
	Name *string `json:"name,omitempty"`
	// Description of the role.
	Description *string `json:"description,omitempty"`
	// The `Access Control List` json defining the role of the user. This represents the access control on the user level.
	Acl *string `json:"acl,omitempty"`
	// An array of user identifiers.
	Members *[]int32 `json:"members,omitempty"`
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateRole) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateRole) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateRole) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateRole) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateRole) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateRole) SetDescription(v string) {
	o.Description = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *UpdateRole) GetAcl() string {
	if o == nil || o.Acl == nil {
		var ret string
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetAclOk() (string, bool) {
	if o == nil || o.Acl == nil {
		var ret string
		return ret, false
	}
	return *o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *UpdateRole) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *UpdateRole) SetAcl(v string) {
	o.Acl = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *UpdateRole) GetMembers() []int32 {
	if o == nil || o.Members == nil {
		var ret []int32
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRole) GetMembersOk() ([]int32, bool) {
	if o == nil || o.Members == nil {
		var ret []int32
		return ret, false
	}
	return *o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *UpdateRole) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []int32 and assigns it to the Members field.
func (o *UpdateRole) SetMembers(v []int32) {
	o.Members = &v
}

type NullableUpdateRole struct {
	Value        UpdateRole
	ExplicitNull bool
}

func (v NullableUpdateRole) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateRole) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
