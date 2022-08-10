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

// CampaignStateChangedNotification A notification regarding a campaign whose state changed.
type CampaignStateChangedNotification struct {
	Campaign Campaign `json:"campaign"`
	// The campaign's old state. Can be one of the following: ['running', 'disabled', 'scheduled', 'expired', 'draft', 'archived']
	OldState string `json:"oldState"`
	// The campaign's new state. Can be one of the following: ['running', 'disabled', 'scheduled', 'expired', 'draft', 'archived']
	NewState string `json:"newState"`
}

// GetCampaign returns the Campaign field value
func (o *CampaignStateChangedNotification) GetCampaign() Campaign {
	if o == nil {
		var ret Campaign
		return ret
	}

	return o.Campaign
}

// SetCampaign sets field value
func (o *CampaignStateChangedNotification) SetCampaign(v Campaign) {
	o.Campaign = v
}

// GetOldState returns the OldState field value
func (o *CampaignStateChangedNotification) GetOldState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldState
}

// SetOldState sets field value
func (o *CampaignStateChangedNotification) SetOldState(v string) {
	o.OldState = v
}

// GetNewState returns the NewState field value
func (o *CampaignStateChangedNotification) GetNewState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewState
}

// SetNewState sets field value
func (o *CampaignStateChangedNotification) SetNewState(v string) {
	o.NewState = v
}

type NullableCampaignStateChangedNotification struct {
	Value        CampaignStateChangedNotification
	ExplicitNull bool
}

func (v NullableCampaignStateChangedNotification) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignStateChangedNotification) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
