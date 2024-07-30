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

// InlineResponse20047 struct for InlineResponse20047
type InlineResponse20047 struct {
	HasMore bool `json:"hasMore"`
	Data []AchievementProgress `json:"data"`
}

// GetHasMore returns the HasMore field value
func (o *InlineResponse20047) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// SetHasMore sets field value
func (o *InlineResponse20047) SetHasMore(v bool) {
	o.HasMore = v
}

// GetData returns the Data field value
func (o *InlineResponse20047) GetData() []AchievementProgress {
	if o == nil {
		var ret []AchievementProgress
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20047) SetData(v []AchievementProgress) {
	o.Data = v
}

type NullableInlineResponse20047 struct {
	Value InlineResponse20047
	ExplicitNull bool
}

func (v NullableInlineResponse20047) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20047) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
