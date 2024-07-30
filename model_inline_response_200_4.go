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

// InlineResponse2004 struct for InlineResponse2004
type InlineResponse2004 struct {
	HasMore bool `json:"hasMore"`
	Data []LedgerPointsEntryIntegrationApi `json:"data"`
}

// GetHasMore returns the HasMore field value
func (o *InlineResponse2004) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// SetHasMore sets field value
func (o *InlineResponse2004) SetHasMore(v bool) {
	o.HasMore = v
}

// GetData returns the Data field value
func (o *InlineResponse2004) GetData() []LedgerPointsEntryIntegrationApi {
	if o == nil {
		var ret []LedgerPointsEntryIntegrationApi
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse2004) SetData(v []LedgerPointsEntryIntegrationApi) {
	o.Data = v
}

type NullableInlineResponse2004 struct {
	Value InlineResponse2004
	ExplicitNull bool
}

func (v NullableInlineResponse2004) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse2004) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
