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

// CampaignSetLeafNode struct for CampaignSetLeafNode
type CampaignSetLeafNode struct {
	// ID of the campaign
	CampaignId int32 `json:"campaignId"`
	// Indicates the node type.
	Type string `json:"type"`
}

// GetCampaignId returns the CampaignId field value
func (o *CampaignSetLeafNode) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *CampaignSetLeafNode) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetType returns the Type field value
func (o *CampaignSetLeafNode) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *CampaignSetLeafNode) SetType(v string) {
	o.Type = v
}

type NullableCampaignSetLeafNode struct {
	Value CampaignSetLeafNode
	ExplicitNull bool
}

func (v NullableCampaignSetLeafNode) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignSetLeafNode) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
