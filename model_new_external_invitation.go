/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// NewExternalInvitation Parameters for inviting a new user from an external identity provider.
type NewExternalInvitation struct {
	// Name of the user.
	Name *string `json:"name,omitempty"`
	// List of user groups in the external identity provider.  If there are roles in Talon.One whose names match these user groups, those roles will be automatically assigned to the user upon invitation. 
	UserGroups *[]string `json:"userGroups,omitempty"`
	// Email address of the user.
	Email string `json:"email"`
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NewExternalInvitation) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewExternalInvitation) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NewExternalInvitation) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NewExternalInvitation) SetName(v string) {
	o.Name = &v
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise.
func (o *NewExternalInvitation) GetUserGroups() []string {
	if o == nil || o.UserGroups == nil {
		var ret []string
		return ret
	}
	return *o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewExternalInvitation) GetUserGroupsOk() ([]string, bool) {
	if o == nil || o.UserGroups == nil {
		var ret []string
		return ret, false
	}
	return *o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *NewExternalInvitation) HasUserGroups() bool {
	if o != nil && o.UserGroups != nil {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given []string and assigns it to the UserGroups field.
func (o *NewExternalInvitation) SetUserGroups(v []string) {
	o.UserGroups = &v
}

// GetEmail returns the Email field value
func (o *NewExternalInvitation) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// SetEmail sets field value
func (o *NewExternalInvitation) SetEmail(v string) {
	o.Email = v
}

type NullableNewExternalInvitation struct {
	Value NewExternalInvitation
	ExplicitNull bool
}

func (v NullableNewExternalInvitation) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewExternalInvitation) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}