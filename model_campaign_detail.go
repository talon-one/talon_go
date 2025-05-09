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

// CampaignDetail struct for CampaignDetail
type CampaignDetail struct {
	// The ID of the campaign that references the application cart item filter.
	CampaignId *int32 `json:"campaignId,omitempty"`
	// A user-facing name for this campaign.
	CampaignName *string `json:"campaignName,omitempty"`
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *CampaignDetail) GetCampaignId() int32 {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDetail) GetCampaignIdOk() (int32, bool) {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret, false
	}
	return *o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *CampaignDetail) HasCampaignId() bool {
	if o != nil && o.CampaignId != nil {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.
func (o *CampaignDetail) SetCampaignId(v int32) {
	o.CampaignId = &v
}

// GetCampaignName returns the CampaignName field value if set, zero value otherwise.
func (o *CampaignDetail) GetCampaignName() string {
	if o == nil || o.CampaignName == nil {
		var ret string
		return ret
	}
	return *o.CampaignName
}

// GetCampaignNameOk returns a tuple with the CampaignName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDetail) GetCampaignNameOk() (string, bool) {
	if o == nil || o.CampaignName == nil {
		var ret string
		return ret, false
	}
	return *o.CampaignName, true
}

// HasCampaignName returns a boolean if a field has been set.
func (o *CampaignDetail) HasCampaignName() bool {
	if o != nil && o.CampaignName != nil {
		return true
	}

	return false
}

// SetCampaignName gets a reference to the given string and assigns it to the CampaignName field.
func (o *CampaignDetail) SetCampaignName(v string) {
	o.CampaignName = &v
}

type NullableCampaignDetail struct {
	Value        CampaignDetail
	ExplicitNull bool
}

func (v NullableCampaignDetail) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignDetail) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
