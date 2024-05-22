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

// LoyaltyLedgerTransactions List of loyalty ledger transactions.
type LoyaltyLedgerTransactions struct {
	// If `true`, it means that there is more data to request in the source collection.
	HasMore *bool `json:"hasMore,omitempty"`
	// List of transaction entries from a loyalty ledger.
	Data *[]LoyaltyLedgerEntry `json:"data,omitempty"`
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *LoyaltyLedgerTransactions) GetHasMore() bool {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerTransactions) GetHasMoreOk() (bool, bool) {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret, false
	}
	return *o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *LoyaltyLedgerTransactions) HasHasMore() bool {
	if o != nil && o.HasMore != nil {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *LoyaltyLedgerTransactions) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoyaltyLedgerTransactions) GetData() []LoyaltyLedgerEntry {
	if o == nil || o.Data == nil {
		var ret []LoyaltyLedgerEntry
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyLedgerTransactions) GetDataOk() ([]LoyaltyLedgerEntry, bool) {
	if o == nil || o.Data == nil {
		var ret []LoyaltyLedgerEntry
		return ret, false
	}
	return *o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoyaltyLedgerTransactions) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []LoyaltyLedgerEntry and assigns it to the Data field.
func (o *LoyaltyLedgerTransactions) SetData(v []LoyaltyLedgerEntry) {
	o.Data = &v
}

type NullableLoyaltyLedgerTransactions struct {
	Value        LoyaltyLedgerTransactions
	ExplicitNull bool
}

func (v NullableLoyaltyLedgerTransactions) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyLedgerTransactions) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
