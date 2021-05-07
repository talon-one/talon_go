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

// FeatureFlags struct for FeatureFlags
type FeatureFlags struct {
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// Whether the account has access to the loyalty features or not
	Loyalty *bool `json:"loyalty,omitempty"`
	// Whether the account queries coupons with or without total result size
	CouponsWithoutCount *bool `json:"coupons_without_count,omitempty"`
	// Whether the account can test beta effects or not
	BetaEffects *bool `json:"betaEffects,omitempty"`
}

// GetAccountId returns the AccountId field value
func (o *FeatureFlags) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *FeatureFlags) SetAccountId(v int32) {
	o.AccountId = v
}

// GetLoyalty returns the Loyalty field value if set, zero value otherwise.
func (o *FeatureFlags) GetLoyalty() bool {
	if o == nil || o.Loyalty == nil {
		var ret bool
		return ret
	}
	return *o.Loyalty
}

// GetLoyaltyOk returns a tuple with the Loyalty field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlags) GetLoyaltyOk() (bool, bool) {
	if o == nil || o.Loyalty == nil {
		var ret bool
		return ret, false
	}
	return *o.Loyalty, true
}

// HasLoyalty returns a boolean if a field has been set.
func (o *FeatureFlags) HasLoyalty() bool {
	if o != nil && o.Loyalty != nil {
		return true
	}

	return false
}

// SetLoyalty gets a reference to the given bool and assigns it to the Loyalty field.
func (o *FeatureFlags) SetLoyalty(v bool) {
	o.Loyalty = &v
}

// GetCouponsWithoutCount returns the CouponsWithoutCount field value if set, zero value otherwise.
func (o *FeatureFlags) GetCouponsWithoutCount() bool {
	if o == nil || o.CouponsWithoutCount == nil {
		var ret bool
		return ret
	}
	return *o.CouponsWithoutCount
}

// GetCouponsWithoutCountOk returns a tuple with the CouponsWithoutCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlags) GetCouponsWithoutCountOk() (bool, bool) {
	if o == nil || o.CouponsWithoutCount == nil {
		var ret bool
		return ret, false
	}
	return *o.CouponsWithoutCount, true
}

// HasCouponsWithoutCount returns a boolean if a field has been set.
func (o *FeatureFlags) HasCouponsWithoutCount() bool {
	if o != nil && o.CouponsWithoutCount != nil {
		return true
	}

	return false
}

// SetCouponsWithoutCount gets a reference to the given bool and assigns it to the CouponsWithoutCount field.
func (o *FeatureFlags) SetCouponsWithoutCount(v bool) {
	o.CouponsWithoutCount = &v
}

// GetBetaEffects returns the BetaEffects field value if set, zero value otherwise.
func (o *FeatureFlags) GetBetaEffects() bool {
	if o == nil || o.BetaEffects == nil {
		var ret bool
		return ret
	}
	return *o.BetaEffects
}

// GetBetaEffectsOk returns a tuple with the BetaEffects field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlags) GetBetaEffectsOk() (bool, bool) {
	if o == nil || o.BetaEffects == nil {
		var ret bool
		return ret, false
	}
	return *o.BetaEffects, true
}

// HasBetaEffects returns a boolean if a field has been set.
func (o *FeatureFlags) HasBetaEffects() bool {
	if o != nil && o.BetaEffects != nil {
		return true
	}

	return false
}

// SetBetaEffects gets a reference to the given bool and assigns it to the BetaEffects field.
func (o *FeatureFlags) SetBetaEffects(v bool) {
	o.BetaEffects = &v
}

type NullableFeatureFlags struct {
	Value        FeatureFlags
	ExplicitNull bool
}

func (v NullableFeatureFlags) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFeatureFlags) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
