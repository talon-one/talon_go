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

// LoyaltyCardAllOf struct for LoyaltyCardAllOf
type LoyaltyCardAllOf struct {
	// The alphanumeric identifier of the loyalty card. 
	Identifier string `json:"identifier"`
	// The max amount of customer profiles that can be linked to the card. 0 means unlimited. 
	UsersPerCardLimit int32 `json:"usersPerCardLimit"`
	// Integration IDs of the customers profiles linked to the card.
	Profiles *[]LoyaltyCardProfileRegistration `json:"profiles,omitempty"`
	Ledger *LedgerInfo `json:"ledger,omitempty"`
	// Displays point balances of the card in the subledgers of the loyalty program.
	Subledgers *map[string]LedgerInfo `json:"subledgers,omitempty"`
	// Timestamp of the most recent update of the loyalty card.
	Modified *time.Time `json:"modified,omitempty"`
	// The alphanumeric identifier of the loyalty card. 
	OldCardIdentifier *string `json:"oldCardIdentifier,omitempty"`
	// The alphanumeric identifier of the loyalty card. 
	NewCardIdentifier *string `json:"newCardIdentifier,omitempty"`
}

// GetIdentifier returns the Identifier field value
func (o *LoyaltyCardAllOf) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// SetIdentifier sets field value
func (o *LoyaltyCardAllOf) SetIdentifier(v string) {
	o.Identifier = v
}

// GetUsersPerCardLimit returns the UsersPerCardLimit field value
func (o *LoyaltyCardAllOf) GetUsersPerCardLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsersPerCardLimit
}

// SetUsersPerCardLimit sets field value
func (o *LoyaltyCardAllOf) SetUsersPerCardLimit(v int32) {
	o.UsersPerCardLimit = v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *LoyaltyCardAllOf) GetProfiles() []LoyaltyCardProfileRegistration {
	if o == nil || o.Profiles == nil {
		var ret []LoyaltyCardProfileRegistration
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyCardAllOf) GetProfilesOk() ([]LoyaltyCardProfileRegistration, bool) {
	if o == nil || o.Profiles == nil {
		var ret []LoyaltyCardProfileRegistration
		return ret, false
	}
	return *o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *LoyaltyCardAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []LoyaltyCardProfileRegistration and assigns it to the Profiles field.
func (o *LoyaltyCardAllOf) SetProfiles(v []LoyaltyCardProfileRegistration) {
	o.Profiles = &v
}

// GetLedger returns the Ledger field value if set, zero value otherwise.
func (o *LoyaltyCardAllOf) GetLedger() LedgerInfo {
	if o == nil || o.Ledger == nil {
		var ret LedgerInfo
		return ret
	}
	return *o.Ledger
}

// GetLedgerOk returns a tuple with the Ledger field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyCardAllOf) GetLedgerOk() (LedgerInfo, bool) {
	if o == nil || o.Ledger == nil {
		var ret LedgerInfo
		return ret, false
	}
	return *o.Ledger, true
}

// HasLedger returns a boolean if a field has been set.
func (o *LoyaltyCardAllOf) HasLedger() bool {
	if o != nil && o.Ledger != nil {
		return true
	}

	return false
}

// SetLedger gets a reference to the given LedgerInfo and assigns it to the Ledger field.
func (o *LoyaltyCardAllOf) SetLedger(v LedgerInfo) {
	o.Ledger = &v
}

// GetSubledgers returns the Subledgers field value if set, zero value otherwise.
func (o *LoyaltyCardAllOf) GetSubledgers() map[string]LedgerInfo {
	if o == nil || o.Subledgers == nil {
		var ret map[string]LedgerInfo
		return ret
	}
	return *o.Subledgers
}

// GetSubledgersOk returns a tuple with the Subledgers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyCardAllOf) GetSubledgersOk() (map[string]LedgerInfo, bool) {
	if o == nil || o.Subledgers == nil {
		var ret map[string]LedgerInfo
		return ret, false
	}
	return *o.Subledgers, true
}

// HasSubledgers returns a boolean if a field has been set.
func (o *LoyaltyCardAllOf) HasSubledgers() bool {
	if o != nil && o.Subledgers != nil {
		return true
	}

	return false
}

// SetSubledgers gets a reference to the given map[string]LedgerInfo and assigns it to the Subledgers field.
func (o *LoyaltyCardAllOf) SetSubledgers(v map[string]LedgerInfo) {
	o.Subledgers = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *LoyaltyCardAllOf) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyCardAllOf) GetModifiedOk() (time.Time, bool) {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *LoyaltyCardAllOf) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *LoyaltyCardAllOf) SetModified(v time.Time) {
	o.Modified = &v
}

// GetOldCardIdentifier returns the OldCardIdentifier field value if set, zero value otherwise.
func (o *LoyaltyCardAllOf) GetOldCardIdentifier() string {
	if o == nil || o.OldCardIdentifier == nil {
		var ret string
		return ret
	}
	return *o.OldCardIdentifier
}

// GetOldCardIdentifierOk returns a tuple with the OldCardIdentifier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyCardAllOf) GetOldCardIdentifierOk() (string, bool) {
	if o == nil || o.OldCardIdentifier == nil {
		var ret string
		return ret, false
	}
	return *o.OldCardIdentifier, true
}

// HasOldCardIdentifier returns a boolean if a field has been set.
func (o *LoyaltyCardAllOf) HasOldCardIdentifier() bool {
	if o != nil && o.OldCardIdentifier != nil {
		return true
	}

	return false
}

// SetOldCardIdentifier gets a reference to the given string and assigns it to the OldCardIdentifier field.
func (o *LoyaltyCardAllOf) SetOldCardIdentifier(v string) {
	o.OldCardIdentifier = &v
}

// GetNewCardIdentifier returns the NewCardIdentifier field value if set, zero value otherwise.
func (o *LoyaltyCardAllOf) GetNewCardIdentifier() string {
	if o == nil || o.NewCardIdentifier == nil {
		var ret string
		return ret
	}
	return *o.NewCardIdentifier
}

// GetNewCardIdentifierOk returns a tuple with the NewCardIdentifier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyCardAllOf) GetNewCardIdentifierOk() (string, bool) {
	if o == nil || o.NewCardIdentifier == nil {
		var ret string
		return ret, false
	}
	return *o.NewCardIdentifier, true
}

// HasNewCardIdentifier returns a boolean if a field has been set.
func (o *LoyaltyCardAllOf) HasNewCardIdentifier() bool {
	if o != nil && o.NewCardIdentifier != nil {
		return true
	}

	return false
}

// SetNewCardIdentifier gets a reference to the given string and assigns it to the NewCardIdentifier field.
func (o *LoyaltyCardAllOf) SetNewCardIdentifier(v string) {
	o.NewCardIdentifier = &v
}

type NullableLoyaltyCardAllOf struct {
	Value LoyaltyCardAllOf
	ExplicitNull bool
}

func (v NullableLoyaltyCardAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyCardAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}