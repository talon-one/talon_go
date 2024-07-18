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

// EffectAllOf A generic effect that is fired by a triggered campaign. The props property will contain information specific to the specific effect type.
type EffectAllOf struct {
	Props EffectProps `json:"props"`
}

// GetProps returns the Props field value
func (o *EffectAllOf) GetProps() EffectProps {
	if o == nil {
		var ret EffectProps
		return ret
	}

	return o.Props
}

// SetProps sets field value
func (o *EffectAllOf) SetProps(v EffectProps) {
	o.Props = v
}

type NullableEffectAllOf struct {
	Value EffectAllOf
	ExplicitNull bool
}

func (v NullableEffectAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEffectAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
