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

// GenerateCampaignDescription struct for GenerateCampaignDescription
type GenerateCampaignDescription struct {
	// ID of the campaign.
	CampaignID int32 `json:"campaignID"`
	// Currency for the campaign.
	Currency string `json:"currency"`
}

// GetCampaignID returns the CampaignID field value
func (o *GenerateCampaignDescription) GetCampaignID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignID
}

// SetCampaignID sets field value
func (o *GenerateCampaignDescription) SetCampaignID(v int32) {
	o.CampaignID = v
}

// GetCurrency returns the Currency field value
func (o *GenerateCampaignDescription) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// SetCurrency sets field value
func (o *GenerateCampaignDescription) SetCurrency(v string) {
	o.Currency = v
}

type NullableGenerateCampaignDescription struct {
	Value        GenerateCampaignDescription
	ExplicitNull bool
}

func (v NullableGenerateCampaignDescription) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableGenerateCampaignDescription) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
