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

// LoyaltyBalancesWithTiers List of loyalty balances for a ledger and its subledgers.
type LoyaltyBalancesWithTiers struct {
	Balance *LoyaltyBalanceWithTier `json:"balance,omitempty"`
	// Map of the loyalty balances of the subledgers of a ledger.
	SubledgerBalances *map[string]LoyaltyBalanceWithTier `json:"subledgerBalances,omitempty"`
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *LoyaltyBalancesWithTiers) GetBalance() LoyaltyBalanceWithTier {
	if o == nil || o.Balance == nil {
		var ret LoyaltyBalanceWithTier
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalancesWithTiers) GetBalanceOk() (LoyaltyBalanceWithTier, bool) {
	if o == nil || o.Balance == nil {
		var ret LoyaltyBalanceWithTier
		return ret, false
	}
	return *o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *LoyaltyBalancesWithTiers) HasBalance() bool {
	if o != nil && o.Balance != nil {
		return true
	}

	return false
}

// SetBalance gets a reference to the given LoyaltyBalanceWithTier and assigns it to the Balance field.
func (o *LoyaltyBalancesWithTiers) SetBalance(v LoyaltyBalanceWithTier) {
	o.Balance = &v
}

// GetSubledgerBalances returns the SubledgerBalances field value if set, zero value otherwise.
func (o *LoyaltyBalancesWithTiers) GetSubledgerBalances() map[string]LoyaltyBalanceWithTier {
	if o == nil || o.SubledgerBalances == nil {
		var ret map[string]LoyaltyBalanceWithTier
		return ret
	}
	return *o.SubledgerBalances
}

// GetSubledgerBalancesOk returns a tuple with the SubledgerBalances field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalancesWithTiers) GetSubledgerBalancesOk() (map[string]LoyaltyBalanceWithTier, bool) {
	if o == nil || o.SubledgerBalances == nil {
		var ret map[string]LoyaltyBalanceWithTier
		return ret, false
	}
	return *o.SubledgerBalances, true
}

// HasSubledgerBalances returns a boolean if a field has been set.
func (o *LoyaltyBalancesWithTiers) HasSubledgerBalances() bool {
	if o != nil && o.SubledgerBalances != nil {
		return true
	}

	return false
}

// SetSubledgerBalances gets a reference to the given map[string]LoyaltyBalanceWithTier and assigns it to the SubledgerBalances field.
func (o *LoyaltyBalancesWithTiers) SetSubledgerBalances(v map[string]LoyaltyBalanceWithTier) {
	o.SubledgerBalances = &v
}

type NullableLoyaltyBalancesWithTiers struct {
	Value        LoyaltyBalancesWithTiers
	ExplicitNull bool
}

func (v NullableLoyaltyBalancesWithTiers) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyBalancesWithTiers) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
