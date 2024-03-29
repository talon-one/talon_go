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

// BulkCampaignNotification struct for BulkCampaignNotification
type BulkCampaignNotification struct {
	TotalResultSize int32                  `json:"totalResultSize"`
	Data            []CampaignNotification `json:"data"`
}

// GetTotalResultSize returns the TotalResultSize field value
func (o *BulkCampaignNotification) GetTotalResultSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResultSize
}

// SetTotalResultSize sets field value
func (o *BulkCampaignNotification) SetTotalResultSize(v int32) {
	o.TotalResultSize = v
}

// GetData returns the Data field value
func (o *BulkCampaignNotification) GetData() []CampaignNotification {
	if o == nil {
		var ret []CampaignNotification
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *BulkCampaignNotification) SetData(v []CampaignNotification) {
	o.Data = v
}

type NullableBulkCampaignNotification struct {
	Value        BulkCampaignNotification
	ExplicitNull bool
}

func (v NullableBulkCampaignNotification) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBulkCampaignNotification) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
