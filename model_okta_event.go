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

// OktaEvent Single event definition in the event data emitted by Okta.
type OktaEvent struct {
	// Event type defining an action.
	EventType string `json:"eventType"`
	Target []OktaEventTarget `json:"target"`
}

// GetEventType returns the EventType field value
func (o *OktaEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// SetEventType sets field value
func (o *OktaEvent) SetEventType(v string) {
	o.EventType = v
}

// GetTarget returns the Target field value
func (o *OktaEvent) GetTarget() []OktaEventTarget {
	if o == nil {
		var ret []OktaEventTarget
		return ret
	}

	return o.Target
}

// SetTarget sets field value
func (o *OktaEvent) SetTarget(v []OktaEventTarget) {
	o.Target = v
}

type NullableOktaEvent struct {
	Value OktaEvent
	ExplicitNull bool
}

func (v NullableOktaEvent) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOktaEvent) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
