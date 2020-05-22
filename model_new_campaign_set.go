/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// NewCampaignSet
type NewCampaignSet struct {
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Version of the campaign set
	Version int32                 `json:"version"`
	Set     CampaignSetBranchNode `json:"set"`
}

// GetApplicationId returns the ApplicationId field value
func (o *NewCampaignSet) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *NewCampaignSet) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetVersion returns the Version field value
func (o *NewCampaignSet) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// SetVersion sets field value
func (o *NewCampaignSet) SetVersion(v int32) {
	o.Version = v
}

// GetSet returns the Set field value
func (o *NewCampaignSet) GetSet() CampaignSetBranchNode {
	if o == nil {
		var ret CampaignSetBranchNode
		return ret
	}

	return o.Set
}

// SetSet sets field value
func (o *NewCampaignSet) SetSet(v CampaignSetBranchNode) {
	o.Set = v
}

type NullableNewCampaignSet struct {
	Value        NewCampaignSet
	ExplicitNull bool
}

func (v NullableNewCampaignSet) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewCampaignSet) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
