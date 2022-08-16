/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerSession](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2) endpoint is `https://mycompany.talon.one/v2/customer_sessions/{Id}`
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// InlineResponse20015 struct for InlineResponse20015
type InlineResponse20015 struct {
	TotalResultSize int32               `json:"totalResultSize"`
	Data            []CampaignAnalytics `json:"data"`
}

// GetTotalResultSize returns the TotalResultSize field value
func (o *InlineResponse20015) GetTotalResultSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResultSize
}

// SetTotalResultSize sets field value
func (o *InlineResponse20015) SetTotalResultSize(v int32) {
	o.TotalResultSize = v
}

// GetData returns the Data field value
func (o *InlineResponse20015) GetData() []CampaignAnalytics {
	if o == nil {
		var ret []CampaignAnalytics
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20015) SetData(v []CampaignAnalytics) {
	o.Data = v
}

type NullableInlineResponse20015 struct {
	Value        InlineResponse20015
	ExplicitNull bool
}

func (v NullableInlineResponse20015) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20015) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
