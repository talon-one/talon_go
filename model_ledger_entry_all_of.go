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
	"time"
)

// LedgerEntryAllOf Entry in the point ledger.
type LedgerEntryAllOf struct {
	// The ID of the Talon.One account that owns this profile.
	AccountId int32 `json:"accountId"`
	// ID of the ledger.
	LoyaltyProgramId int32 `json:"loyaltyProgramId"`
	// ID of the related event.
	EventId int32 `json:"eventId"`
	// Amount of loyalty points.
	Amount int32 `json:"amount"`
	// reason for awarding/deducting points.
	Reason string `json:"reason"`
	// Expiration date of the points.
	ExpiryDate time.Time `json:"expiryDate"`
	// The ID of the balancing ledgerEntry.
	ReferenceId *int32 `json:"referenceId,omitempty"`
}

// GetAccountId returns the AccountId field value
func (o *LedgerEntryAllOf) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *LedgerEntryAllOf) SetAccountId(v int32) {
	o.AccountId = v
}

// GetLoyaltyProgramId returns the LoyaltyProgramId field value
func (o *LedgerEntryAllOf) GetLoyaltyProgramId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoyaltyProgramId
}

// SetLoyaltyProgramId sets field value
func (o *LedgerEntryAllOf) SetLoyaltyProgramId(v int32) {
	o.LoyaltyProgramId = v
}

// GetEventId returns the EventId field value
func (o *LedgerEntryAllOf) GetEventId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventId
}

// SetEventId sets field value
func (o *LedgerEntryAllOf) SetEventId(v int32) {
	o.EventId = v
}

// GetAmount returns the Amount field value
func (o *LedgerEntryAllOf) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// SetAmount sets field value
func (o *LedgerEntryAllOf) SetAmount(v int32) {
	o.Amount = v
}

// GetReason returns the Reason field value
func (o *LedgerEntryAllOf) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// SetReason sets field value
func (o *LedgerEntryAllOf) SetReason(v string) {
	o.Reason = v
}

// GetExpiryDate returns the ExpiryDate field value
func (o *LedgerEntryAllOf) GetExpiryDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiryDate
}

// SetExpiryDate sets field value
func (o *LedgerEntryAllOf) SetExpiryDate(v time.Time) {
	o.ExpiryDate = v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *LedgerEntryAllOf) GetReferenceId() int32 {
	if o == nil || o.ReferenceId == nil {
		var ret int32
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LedgerEntryAllOf) GetReferenceIdOk() (int32, bool) {
	if o == nil || o.ReferenceId == nil {
		var ret int32
		return ret, false
	}
	return *o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *LedgerEntryAllOf) HasReferenceId() bool {
	if o != nil && o.ReferenceId != nil {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given int32 and assigns it to the ReferenceId field.
func (o *LedgerEntryAllOf) SetReferenceId(v int32) {
	o.ReferenceId = &v
}

type NullableLedgerEntryAllOf struct {
	Value LedgerEntryAllOf
	ExplicitNull bool
}

func (v NullableLedgerEntryAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLedgerEntryAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}