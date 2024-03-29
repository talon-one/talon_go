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

// CampaignGroupEntity struct for CampaignGroupEntity
type CampaignGroupEntity struct {
	// The IDs of the campaign groups that own this entity.
	CampaignGroups *[]int32 `json:"campaignGroups,omitempty"`
}

// GetCampaignGroups returns the CampaignGroups field value if set, zero value otherwise.
func (o *CampaignGroupEntity) GetCampaignGroups() []int32 {
	if o == nil || o.CampaignGroups == nil {
		var ret []int32
		return ret
	}
	return *o.CampaignGroups
}

// GetCampaignGroupsOk returns a tuple with the CampaignGroups field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignGroupEntity) GetCampaignGroupsOk() ([]int32, bool) {
	if o == nil || o.CampaignGroups == nil {
		var ret []int32
		return ret, false
	}
	return *o.CampaignGroups, true
}

// HasCampaignGroups returns a boolean if a field has been set.
func (o *CampaignGroupEntity) HasCampaignGroups() bool {
	if o != nil && o.CampaignGroups != nil {
		return true
	}

	return false
}

// SetCampaignGroups gets a reference to the given []int32 and assigns it to the CampaignGroups field.
func (o *CampaignGroupEntity) SetCampaignGroups(v []int32) {
	o.CampaignGroups = &v
}

type NullableCampaignGroupEntity struct {
	Value        CampaignGroupEntity
	ExplicitNull bool
}

func (v NullableCampaignGroupEntity) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignGroupEntity) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
