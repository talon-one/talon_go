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

// InlineResponse20011 struct for InlineResponse20011
type InlineResponse20011 struct {
	TotalResultSize int32 `json:"totalResultSize"`
	Data []CampaignGroup `json:"data"`
}

// GetTotalResultSize returns the TotalResultSize field value
func (o *InlineResponse20011) GetTotalResultSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResultSize
}

// SetTotalResultSize sets field value
func (o *InlineResponse20011) SetTotalResultSize(v int32) {
	o.TotalResultSize = v
}

// GetData returns the Data field value
func (o *InlineResponse20011) GetData() []CampaignGroup {
	if o == nil {
		var ret []CampaignGroup
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20011) SetData(v []CampaignGroup) {
	o.Data = v
}

type NullableInlineResponse20011 struct {
	Value InlineResponse20011
	ExplicitNull bool
}

func (v NullableInlineResponse20011) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20011) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
