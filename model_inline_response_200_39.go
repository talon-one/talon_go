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

// InlineResponse20039 struct for InlineResponse20039
type InlineResponse20039 struct {
	Data []CardLedgerTransactionLogEntryIntegrationApi `json:"data"`
	HasMore bool `json:"hasMore"`
}

// GetData returns the Data field value
func (o *InlineResponse20039) GetData() []CardLedgerTransactionLogEntryIntegrationApi {
	if o == nil {
		var ret []CardLedgerTransactionLogEntryIntegrationApi
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20039) SetData(v []CardLedgerTransactionLogEntryIntegrationApi) {
	o.Data = v
}

// GetHasMore returns the HasMore field value
func (o *InlineResponse20039) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// SetHasMore sets field value
func (o *InlineResponse20039) SetHasMore(v bool) {
	o.HasMore = v
}

type NullableInlineResponse20039 struct {
	Value InlineResponse20039
	ExplicitNull bool
}

func (v NullableInlineResponse20039) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20039) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
