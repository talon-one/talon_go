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

// CampaignPrioritiesV2 struct for CampaignPrioritiesV2
type CampaignPrioritiesV2 struct {
	Exclusive *[]CampaignSetIDs `json:"exclusive,omitempty"`
	Stackable *[]CampaignSetIDs `json:"stackable,omitempty"`
	Universal *[]CampaignSetIDs `json:"universal,omitempty"`
}

// GetExclusive returns the Exclusive field value if set, zero value otherwise.
func (o *CampaignPrioritiesV2) GetExclusive() []CampaignSetIDs {
	if o == nil || o.Exclusive == nil {
		var ret []CampaignSetIDs
		return ret
	}
	return *o.Exclusive
}

// GetExclusiveOk returns a tuple with the Exclusive field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignPrioritiesV2) GetExclusiveOk() ([]CampaignSetIDs, bool) {
	if o == nil || o.Exclusive == nil {
		var ret []CampaignSetIDs
		return ret, false
	}
	return *o.Exclusive, true
}

// HasExclusive returns a boolean if a field has been set.
func (o *CampaignPrioritiesV2) HasExclusive() bool {
	if o != nil && o.Exclusive != nil {
		return true
	}

	return false
}

// SetExclusive gets a reference to the given []CampaignSetIDs and assigns it to the Exclusive field.
func (o *CampaignPrioritiesV2) SetExclusive(v []CampaignSetIDs) {
	o.Exclusive = &v
}

// GetStackable returns the Stackable field value if set, zero value otherwise.
func (o *CampaignPrioritiesV2) GetStackable() []CampaignSetIDs {
	if o == nil || o.Stackable == nil {
		var ret []CampaignSetIDs
		return ret
	}
	return *o.Stackable
}

// GetStackableOk returns a tuple with the Stackable field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignPrioritiesV2) GetStackableOk() ([]CampaignSetIDs, bool) {
	if o == nil || o.Stackable == nil {
		var ret []CampaignSetIDs
		return ret, false
	}
	return *o.Stackable, true
}

// HasStackable returns a boolean if a field has been set.
func (o *CampaignPrioritiesV2) HasStackable() bool {
	if o != nil && o.Stackable != nil {
		return true
	}

	return false
}

// SetStackable gets a reference to the given []CampaignSetIDs and assigns it to the Stackable field.
func (o *CampaignPrioritiesV2) SetStackable(v []CampaignSetIDs) {
	o.Stackable = &v
}

// GetUniversal returns the Universal field value if set, zero value otherwise.
func (o *CampaignPrioritiesV2) GetUniversal() []CampaignSetIDs {
	if o == nil || o.Universal == nil {
		var ret []CampaignSetIDs
		return ret
	}
	return *o.Universal
}

// GetUniversalOk returns a tuple with the Universal field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignPrioritiesV2) GetUniversalOk() ([]CampaignSetIDs, bool) {
	if o == nil || o.Universal == nil {
		var ret []CampaignSetIDs
		return ret, false
	}
	return *o.Universal, true
}

// HasUniversal returns a boolean if a field has been set.
func (o *CampaignPrioritiesV2) HasUniversal() bool {
	if o != nil && o.Universal != nil {
		return true
	}

	return false
}

// SetUniversal gets a reference to the given []CampaignSetIDs and assigns it to the Universal field.
func (o *CampaignPrioritiesV2) SetUniversal(v []CampaignSetIDs) {
	o.Universal = &v
}

type NullableCampaignPrioritiesV2 struct {
	Value        CampaignPrioritiesV2
	ExplicitNull bool
}

func (v NullableCampaignPrioritiesV2) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignPrioritiesV2) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}