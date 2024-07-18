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
	"time"
)

// LoyaltyCardProfileRegistration struct for LoyaltyCardProfileRegistration
type LoyaltyCardProfileRegistration struct {
	// Integration ID of the customer profile linked to the card.
	IntegrationId string `json:"integrationId"`
	// Timestamp the customer profile was linked to the card.
	Timestamp time.Time `json:"timestamp"`
}

// GetIntegrationId returns the IntegrationId field value
func (o *LoyaltyCardProfileRegistration) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *LoyaltyCardProfileRegistration) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetTimestamp returns the Timestamp field value
func (o *LoyaltyCardProfileRegistration) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// SetTimestamp sets field value
func (o *LoyaltyCardProfileRegistration) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

type NullableLoyaltyCardProfileRegistration struct {
	Value LoyaltyCardProfileRegistration
	ExplicitNull bool
}

func (v NullableLoyaltyCardProfileRegistration) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyCardProfileRegistration) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
