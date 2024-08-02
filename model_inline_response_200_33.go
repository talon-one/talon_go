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

// InlineResponse20033 struct for InlineResponse20033
type InlineResponse20033 struct {
	Data []EventType `json:"data"`
	TotalResultSize int32 `json:"totalResultSize"`
}

// GetData returns the Data field value
func (o *InlineResponse20033) GetData() []EventType {
	if o == nil {
		var ret []EventType
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20033) SetData(v []EventType) {
	o.Data = v
}

// GetTotalResultSize returns the TotalResultSize field value
func (o *InlineResponse20033) GetTotalResultSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResultSize
}

// SetTotalResultSize sets field value
func (o *InlineResponse20033) SetTotalResultSize(v int32) {
	o.TotalResultSize = v
}

type NullableInlineResponse20033 struct {
	Value InlineResponse20033
	ExplicitNull bool
}

func (v NullableInlineResponse20033) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20033) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
