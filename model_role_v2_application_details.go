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

// RoleV2ApplicationDetails struct for RoleV2ApplicationDetails
type RoleV2ApplicationDetails struct {
	// Name of the Application-related permission set for the given Application.
	Application *string `json:"application,omitempty"`
	// Name of the campaign-related permission set for the given Application.
	Campaign *string `json:"campaign,omitempty"`
	// Name of the draft campaign-related permission set for the given Application.
	DraftCampaign *string `json:"draftCampaign,omitempty"`
	// Name of the tools-related permission set.
	Tools *string `json:"tools,omitempty"`
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *RoleV2ApplicationDetails) GetApplication() string {
	if o == nil || o.Application == nil {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleV2ApplicationDetails) GetApplicationOk() (string, bool) {
	if o == nil || o.Application == nil {
		var ret string
		return ret, false
	}
	return *o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *RoleV2ApplicationDetails) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *RoleV2ApplicationDetails) SetApplication(v string) {
	o.Application = &v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *RoleV2ApplicationDetails) GetCampaign() string {
	if o == nil || o.Campaign == nil {
		var ret string
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleV2ApplicationDetails) GetCampaignOk() (string, bool) {
	if o == nil || o.Campaign == nil {
		var ret string
		return ret, false
	}
	return *o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *RoleV2ApplicationDetails) HasCampaign() bool {
	if o != nil && o.Campaign != nil {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given string and assigns it to the Campaign field.
func (o *RoleV2ApplicationDetails) SetCampaign(v string) {
	o.Campaign = &v
}

// GetDraftCampaign returns the DraftCampaign field value if set, zero value otherwise.
func (o *RoleV2ApplicationDetails) GetDraftCampaign() string {
	if o == nil || o.DraftCampaign == nil {
		var ret string
		return ret
	}
	return *o.DraftCampaign
}

// GetDraftCampaignOk returns a tuple with the DraftCampaign field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleV2ApplicationDetails) GetDraftCampaignOk() (string, bool) {
	if o == nil || o.DraftCampaign == nil {
		var ret string
		return ret, false
	}
	return *o.DraftCampaign, true
}

// HasDraftCampaign returns a boolean if a field has been set.
func (o *RoleV2ApplicationDetails) HasDraftCampaign() bool {
	if o != nil && o.DraftCampaign != nil {
		return true
	}

	return false
}

// SetDraftCampaign gets a reference to the given string and assigns it to the DraftCampaign field.
func (o *RoleV2ApplicationDetails) SetDraftCampaign(v string) {
	o.DraftCampaign = &v
}

// GetTools returns the Tools field value if set, zero value otherwise.
func (o *RoleV2ApplicationDetails) GetTools() string {
	if o == nil || o.Tools == nil {
		var ret string
		return ret
	}
	return *o.Tools
}

// GetToolsOk returns a tuple with the Tools field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleV2ApplicationDetails) GetToolsOk() (string, bool) {
	if o == nil || o.Tools == nil {
		var ret string
		return ret, false
	}
	return *o.Tools, true
}

// HasTools returns a boolean if a field has been set.
func (o *RoleV2ApplicationDetails) HasTools() bool {
	if o != nil && o.Tools != nil {
		return true
	}

	return false
}

// SetTools gets a reference to the given string and assigns it to the Tools field.
func (o *RoleV2ApplicationDetails) SetTools(v string) {
	o.Tools = &v
}

type NullableRoleV2ApplicationDetails struct {
	Value RoleV2ApplicationDetails
	ExplicitNull bool
}

func (v NullableRoleV2ApplicationDetails) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRoleV2ApplicationDetails) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
