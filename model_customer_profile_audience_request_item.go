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

// CustomerProfileAudienceRequestItem struct for CustomerProfileAudienceRequestItem
type CustomerProfileAudienceRequestItem struct {
	// Defines the action to perform: - `add`: Adds the customer profile to the audience.    **Note**: If the customer profile does not exist, it will be created. The profile will not be visible in any Application   until a session or profile update is received for that profile. - `delete`: Removes the customer profile from the audience.
	Action string `json:"action"`
	// The ID of this customer profile in the third-party integration.
	ProfileIntegrationId string `json:"profileIntegrationId"`
	// The ID of the audience. You get it via the `id` property when [creating an audience](#operation/createAudienceV2).
	AudienceId int32 `json:"audienceId"`
}

// GetAction returns the Action field value
func (o *CustomerProfileAudienceRequestItem) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// SetAction sets field value
func (o *CustomerProfileAudienceRequestItem) SetAction(v string) {
	o.Action = v
}

// GetProfileIntegrationId returns the ProfileIntegrationId field value
func (o *CustomerProfileAudienceRequestItem) GetProfileIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileIntegrationId
}

// SetProfileIntegrationId sets field value
func (o *CustomerProfileAudienceRequestItem) SetProfileIntegrationId(v string) {
	o.ProfileIntegrationId = v
}

// GetAudienceId returns the AudienceId field value
func (o *CustomerProfileAudienceRequestItem) GetAudienceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AudienceId
}

// SetAudienceId sets field value
func (o *CustomerProfileAudienceRequestItem) SetAudienceId(v int32) {
	o.AudienceId = v
}

type NullableCustomerProfileAudienceRequestItem struct {
	Value        CustomerProfileAudienceRequestItem
	ExplicitNull bool
}

func (v NullableCustomerProfileAudienceRequestItem) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomerProfileAudienceRequestItem) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
