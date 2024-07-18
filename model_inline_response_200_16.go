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

// InlineResponse20016 struct for InlineResponse20016
type InlineResponse20016 struct {
	// true means there is more data in the source collection to request..
	HasMore bool `json:"hasMore"`
	// List of loyalty card transaction logs.
	Data []CardLedgerTransactionLogEntry `json:"data"`
}

// GetHasMore returns the HasMore field value
func (o *InlineResponse20016) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// SetHasMore sets field value
func (o *InlineResponse20016) SetHasMore(v bool) {
	o.HasMore = v
}

// GetData returns the Data field value
func (o *InlineResponse20016) GetData() []CardLedgerTransactionLogEntry {
	if o == nil {
		var ret []CardLedgerTransactionLogEntry
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20016) SetData(v []CardLedgerTransactionLogEntry) {
	o.Data = v
}

type NullableInlineResponse20016 struct {
	Value InlineResponse20016
	ExplicitNull bool
}

func (v NullableInlineResponse20016) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20016) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
