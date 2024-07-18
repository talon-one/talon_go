/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// LedgerInfo struct for LedgerInfo
type LedgerInfo struct {
	// Sum of currently active points.
	CurrentBalance float32 `json:"currentBalance"`
	// Sum of pending points.
	PendingBalance float32 `json:"pendingBalance"`
	// **DEPRECATED** Value is shown as 0. 
	ExpiredBalance float32 `json:"expiredBalance"`
	// **DEPRECATED** Value is shown as 0. 
	SpentBalance float32 `json:"spentBalance"`
	// Sum of the tentative active points (including additions and deductions) inside the currently open session. The `currentBalance` is updated to this value when you close the session, and the effects are applied.
	TentativeCurrentBalance float32 `json:"tentativeCurrentBalance"`
	// Sum of pending points (including additions and deductions) inside the currently open session. The `pendingBalance` is updated to this value when you close the session, and the effects are applied.
	TentativePendingBalance *float32 `json:"tentativePendingBalance,omitempty"`
	CurrentTier *Tier `json:"currentTier,omitempty"`
	// Points required to move up a tier.
	PointsToNextTier *float32 `json:"pointsToNextTier,omitempty"`
}

// GetCurrentBalance returns the CurrentBalance field value
func (o *LedgerInfo) GetCurrentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentBalance
}

// SetCurrentBalance sets field value
func (o *LedgerInfo) SetCurrentBalance(v float32) {
	o.CurrentBalance = v
}

// GetPendingBalance returns the PendingBalance field value
func (o *LedgerInfo) GetPendingBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingBalance
}

// SetPendingBalance sets field value
func (o *LedgerInfo) SetPendingBalance(v float32) {
	o.PendingBalance = v
}

// GetExpiredBalance returns the ExpiredBalance field value
func (o *LedgerInfo) GetExpiredBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpiredBalance
}

// SetExpiredBalance sets field value
func (o *LedgerInfo) SetExpiredBalance(v float32) {
	o.ExpiredBalance = v
}

// GetSpentBalance returns the SpentBalance field value
func (o *LedgerInfo) GetSpentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpentBalance
}

// SetSpentBalance sets field value
func (o *LedgerInfo) SetSpentBalance(v float32) {
	o.SpentBalance = v
}

// GetTentativeCurrentBalance returns the TentativeCurrentBalance field value
func (o *LedgerInfo) GetTentativeCurrentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TentativeCurrentBalance
}

// SetTentativeCurrentBalance sets field value
func (o *LedgerInfo) SetTentativeCurrentBalance(v float32) {
	o.TentativeCurrentBalance = v
}

// GetTentativePendingBalance returns the TentativePendingBalance field value if set, zero value otherwise.
func (o *LedgerInfo) GetTentativePendingBalance() float32 {
	if o == nil || o.TentativePendingBalance == nil {
		var ret float32
		return ret
	}
	return *o.TentativePendingBalance
}

// GetTentativePendingBalanceOk returns a tuple with the TentativePendingBalance field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetTentativePendingBalanceOk() (float32, bool) {
	if o == nil || o.TentativePendingBalance == nil {
		var ret float32
		return ret, false
	}
	return *o.TentativePendingBalance, true
}

// HasTentativePendingBalance returns a boolean if a field has been set.
func (o *LedgerInfo) HasTentativePendingBalance() bool {
	if o != nil && o.TentativePendingBalance != nil {
		return true
	}

	return false
}

// SetTentativePendingBalance gets a reference to the given float32 and assigns it to the TentativePendingBalance field.
func (o *LedgerInfo) SetTentativePendingBalance(v float32) {
	o.TentativePendingBalance = &v
}

// GetCurrentTier returns the CurrentTier field value if set, zero value otherwise.
func (o *LedgerInfo) GetCurrentTier() Tier {
	if o == nil || o.CurrentTier == nil {
		var ret Tier
		return ret
	}
	return *o.CurrentTier
}

// GetCurrentTierOk returns a tuple with the CurrentTier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetCurrentTierOk() (Tier, bool) {
	if o == nil || o.CurrentTier == nil {
		var ret Tier
		return ret, false
	}
	return *o.CurrentTier, true
}

// HasCurrentTier returns a boolean if a field has been set.
func (o *LedgerInfo) HasCurrentTier() bool {
	if o != nil && o.CurrentTier != nil {
		return true
	}

	return false
}

// SetCurrentTier gets a reference to the given Tier and assigns it to the CurrentTier field.
func (o *LedgerInfo) SetCurrentTier(v Tier) {
	o.CurrentTier = &v
}

// GetPointsToNextTier returns the PointsToNextTier field value if set, zero value otherwise.
func (o *LedgerInfo) GetPointsToNextTier() float32 {
	if o == nil || o.PointsToNextTier == nil {
		var ret float32
		return ret
	}
	return *o.PointsToNextTier
}

// GetPointsToNextTierOk returns a tuple with the PointsToNextTier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LedgerInfo) GetPointsToNextTierOk() (float32, bool) {
	if o == nil || o.PointsToNextTier == nil {
		var ret float32
		return ret, false
	}
	return *o.PointsToNextTier, true
}

// HasPointsToNextTier returns a boolean if a field has been set.
func (o *LedgerInfo) HasPointsToNextTier() bool {
	if o != nil && o.PointsToNextTier != nil {
		return true
	}

	return false
}

// SetPointsToNextTier gets a reference to the given float32 and assigns it to the PointsToNextTier field.
func (o *LedgerInfo) SetPointsToNextTier(v float32) {
	o.PointsToNextTier = &v
}

type NullableLedgerInfo struct {
	Value LedgerInfo
	ExplicitNull bool
}

func (v NullableLedgerInfo) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLedgerInfo) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
