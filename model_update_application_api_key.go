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

// UpdateApplicationApiKey struct for UpdateApplicationApiKey
type UpdateApplicationApiKey struct {
	// A time offset in nanoseconds associated with the API key. When making a request using the API key, rule evaluation is based on a date that is calculated by adding the offset to the current date. 
	TimeOffset int32 `json:"timeOffset"`
}

// GetTimeOffset returns the TimeOffset field value
func (o *UpdateApplicationApiKey) GetTimeOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeOffset
}

// SetTimeOffset sets field value
func (o *UpdateApplicationApiKey) SetTimeOffset(v int32) {
	o.TimeOffset = v
}

type NullableUpdateApplicationApiKey struct {
	Value UpdateApplicationApiKey
	ExplicitNull bool
}

func (v NullableUpdateApplicationApiKey) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateApplicationApiKey) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
