/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// CustomerProfileAudienceRequestItem struct for CustomerProfileAudienceRequestItem
type CustomerProfileAudienceRequestItem struct {
	Action               string `json:"action"`
	ProfileIntegrationId string `json:"profileIntegrationId"`
	AudienceId           int32  `json:"audienceId"`
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
