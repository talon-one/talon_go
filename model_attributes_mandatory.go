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

// AttributesMandatory Arbitrary settings associated with attributes.
type AttributesMandatory struct {
	// List of mandatory attributes for campaigns.
	Campaigns *[]string `json:"campaigns,omitempty"`
	// List of mandatory attributes for campaigns.
	Coupons *[]string `json:"coupons,omitempty"`
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *AttributesMandatory) GetCampaigns() []string {
	if o == nil || o.Campaigns == nil {
		var ret []string
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AttributesMandatory) GetCampaignsOk() ([]string, bool) {
	if o == nil || o.Campaigns == nil {
		var ret []string
		return ret, false
	}
	return *o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *AttributesMandatory) HasCampaigns() bool {
	if o != nil && o.Campaigns != nil {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []string and assigns it to the Campaigns field.
func (o *AttributesMandatory) SetCampaigns(v []string) {
	o.Campaigns = &v
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *AttributesMandatory) GetCoupons() []string {
	if o == nil || o.Coupons == nil {
		var ret []string
		return ret
	}
	return *o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AttributesMandatory) GetCouponsOk() ([]string, bool) {
	if o == nil || o.Coupons == nil {
		var ret []string
		return ret, false
	}
	return *o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *AttributesMandatory) HasCoupons() bool {
	if o != nil && o.Coupons != nil {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given []string and assigns it to the Coupons field.
func (o *AttributesMandatory) SetCoupons(v []string) {
	o.Coupons = &v
}

type NullableAttributesMandatory struct {
	Value        AttributesMandatory
	ExplicitNull bool
}

func (v NullableAttributesMandatory) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAttributesMandatory) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
