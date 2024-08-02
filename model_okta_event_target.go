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

// OktaEventTarget Target of the specific Okta event.
type OktaEventTarget struct {
	// Identifier of the event target, depending on its type.
	AlternateId string `json:"alternateId"`
	// Display name of the event target.
	DisplayName string `json:"displayName"`
	// Type of the event target.
	Type string `json:"type"`
}

// GetAlternateId returns the AlternateId field value
func (o *OktaEventTarget) GetAlternateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternateId
}

// SetAlternateId sets field value
func (o *OktaEventTarget) SetAlternateId(v string) {
	o.AlternateId = v
}

// GetDisplayName returns the DisplayName field value
func (o *OktaEventTarget) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// SetDisplayName sets field value
func (o *OktaEventTarget) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetType returns the Type field value
func (o *OktaEventTarget) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *OktaEventTarget) SetType(v string) {
	o.Type = v
}

type NullableOktaEventTarget struct {
	Value OktaEventTarget
	ExplicitNull bool
}

func (v NullableOktaEventTarget) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOktaEventTarget) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
