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

// LoyaltyProgramBalance The balance in a Loyalty Program for some Customer.
type LoyaltyProgramBalance struct {
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
}

// GetCurrentBalance returns the CurrentBalance field value
func (o *LoyaltyProgramBalance) GetCurrentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentBalance
}

// SetCurrentBalance sets field value
func (o *LoyaltyProgramBalance) SetCurrentBalance(v float32) {
	o.CurrentBalance = v
}

// GetPendingBalance returns the PendingBalance field value
func (o *LoyaltyProgramBalance) GetPendingBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingBalance
}

// SetPendingBalance sets field value
func (o *LoyaltyProgramBalance) SetPendingBalance(v float32) {
	o.PendingBalance = v
}

// GetExpiredBalance returns the ExpiredBalance field value
func (o *LoyaltyProgramBalance) GetExpiredBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpiredBalance
}

// SetExpiredBalance sets field value
func (o *LoyaltyProgramBalance) SetExpiredBalance(v float32) {
	o.ExpiredBalance = v
}

// GetSpentBalance returns the SpentBalance field value
func (o *LoyaltyProgramBalance) GetSpentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpentBalance
}

// SetSpentBalance sets field value
func (o *LoyaltyProgramBalance) SetSpentBalance(v float32) {
	o.SpentBalance = v
}

// GetTentativeCurrentBalance returns the TentativeCurrentBalance field value
func (o *LoyaltyProgramBalance) GetTentativeCurrentBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TentativeCurrentBalance
}

// SetTentativeCurrentBalance sets field value
func (o *LoyaltyProgramBalance) SetTentativeCurrentBalance(v float32) {
	o.TentativeCurrentBalance = v
}

// GetTentativePendingBalance returns the TentativePendingBalance field value if set, zero value otherwise.
func (o *LoyaltyProgramBalance) GetTentativePendingBalance() float32 {
	if o == nil || o.TentativePendingBalance == nil {
		var ret float32
		return ret
	}
	return *o.TentativePendingBalance
}

// GetTentativePendingBalanceOk returns a tuple with the TentativePendingBalance field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramBalance) GetTentativePendingBalanceOk() (float32, bool) {
	if o == nil || o.TentativePendingBalance == nil {
		var ret float32
		return ret, false
	}
	return *o.TentativePendingBalance, true
}

// HasTentativePendingBalance returns a boolean if a field has been set.
func (o *LoyaltyProgramBalance) HasTentativePendingBalance() bool {
	if o != nil && o.TentativePendingBalance != nil {
		return true
	}

	return false
}

// SetTentativePendingBalance gets a reference to the given float32 and assigns it to the TentativePendingBalance field.
func (o *LoyaltyProgramBalance) SetTentativePendingBalance(v float32) {
	o.TentativePendingBalance = &v
}

type NullableLoyaltyProgramBalance struct {
	Value        LoyaltyProgramBalance
	ExplicitNull bool
}

func (v NullableLoyaltyProgramBalance) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyProgramBalance) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
