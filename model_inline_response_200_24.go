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

// InlineResponse20024 struct for InlineResponse20024
type InlineResponse20024 struct {
	HasMore bool              `json:"hasMore"`
	Data    []CustomerProfile `json:"data"`
}

// GetHasMore returns the HasMore field value
func (o *InlineResponse20024) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// SetHasMore sets field value
func (o *InlineResponse20024) SetHasMore(v bool) {
	o.HasMore = v
}

// GetData returns the Data field value
func (o *InlineResponse20024) GetData() []CustomerProfile {
	if o == nil {
		var ret []CustomerProfile
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20024) SetData(v []CustomerProfile) {
	o.Data = v
}

type NullableInlineResponse20024 struct {
	Value        InlineResponse20024
	ExplicitNull bool
}

func (v NullableInlineResponse20024) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20024) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
