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
	"time"
)

// LedgerTransactionLogEntryIntegrationApi Log entry for a given loyalty card transaction.
type LedgerTransactionLogEntryIntegrationApi struct {
	// Date and time the loyalty transaction occurred.
	Created time.Time `json:"created"`
	// ID of the loyalty program.
	ProgramId int32 `json:"programId"`
	// ID of the customer session where the transaction occurred.
	CustomerSessionId *string `json:"customerSessionId,omitempty"`
	// Type of transaction. Possible values:   - `addition`: Signifies added points.   - `subtraction`: Signifies deducted points.
	Type string `json:"type"`
	// Name or reason of the loyalty ledger transaction.
	Name string `json:"name"`
	// When points become active. Possible values:   - `immediate`: Points are immediately active.   - a timestamp value: Points become active at a given date and time.
	StartDate string `json:"startDate"`
	// Date when points expire. Possible values are:   - `unlimited`: Points have no expiration date.   - `timestamp value`: Points expire on the given date.
	ExpiryDate string `json:"expiryDate"`
	// ID of the subledger.
	SubledgerId string `json:"subledgerId"`
	// Amount of loyalty points added or deducted in the transaction.
	Amount float32 `json:"amount"`
	// ID of the loyalty ledger transaction.
	Id int32 `json:"id"`
	// The ID of the ruleset containing the rule that triggered this effect.
	RulesetId *int32 `json:"rulesetId,omitempty"`
	// The name of the rule that triggered this effect.
	RuleName *string                  `json:"ruleName,omitempty"`
	Flags    *LoyaltyLedgerEntryFlags `json:"flags,omitempty"`
}

// GetCreated returns the Created field value
func (o *LedgerTransactionLogEntryIntegrationApi) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *LedgerTransactionLogEntryIntegrationApi) SetCreated(v time.Time) {
	o.Created = v
}

// GetProgramId returns the ProgramId field value
func (o *LedgerTransactionLogEntryIntegrationApi) GetProgramId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramId
}

// SetProgramId sets field value
func (o *LedgerTransactionLogEntryIntegrationApi) SetProgramId(v int32) {
	o.ProgramId = v
}

// GetCustomerSessionId returns the CustomerSessionId field value if set, zero value otherwise.
func (o *LedgerTransactionLogEntryIntegrationApi) GetCustomerSessionId() string {
	if o == nil || o.CustomerSessionId == nil {
		var ret string
		return ret
	}
	return *o.CustomerSessionId
}

// GetCustomerSessionIdOk returns a tuple with the CustomerSessionId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionLogEntryIntegrationApi) GetCustomerSessionIdOk() (string, bool) {
	if o == nil || o.CustomerSessionId == nil {
		var ret string
		return ret, false
	}
	return *o.CustomerSessionId, true
}

// HasCustomerSessionId returns a boolean if a field has been set.
func (o *LedgerTransactionLogEntryIntegrationApi) HasCustomerSessionId() bool {
	if o != nil && o.CustomerSessionId != nil {
		return true
	}

	return false
}

// SetCustomerSessionId gets a reference to the given string and assigns it to the CustomerSessionId field.
func (o *LedgerTransactionLogEntryIntegrationApi) SetCustomerSessionId(v string) {
	o.CustomerSessionId = &v
}

// GetType returns the Type field value
func (o *LedgerTransactionLogEntryIntegrationApi) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *LedgerTransactionLogEntryIntegrationApi) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *LedgerTransactionLogEntryIntegrationApi) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *LedgerTransactionLogEntryIntegrationApi) SetName(v string) {
	o.Name = v
}

// GetStartDate returns the StartDate field value
func (o *LedgerTransactionLogEntryIntegrationApi) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// SetStartDate sets field value
func (o *LedgerTransactionLogEntryIntegrationApi) SetStartDate(v string) {
	o.StartDate = v
}

// GetExpiryDate returns the ExpiryDate field value
func (o *LedgerTransactionLogEntryIntegrationApi) GetExpiryDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryDate
}

// SetExpiryDate sets field value
func (o *LedgerTransactionLogEntryIntegrationApi) SetExpiryDate(v string) {
	o.ExpiryDate = v
}

// GetSubledgerId returns the SubledgerId field value
func (o *LedgerTransactionLogEntryIntegrationApi) GetSubledgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubledgerId
}

// SetSubledgerId sets field value
func (o *LedgerTransactionLogEntryIntegrationApi) SetSubledgerId(v string) {
	o.SubledgerId = v
}

// GetAmount returns the Amount field value
func (o *LedgerTransactionLogEntryIntegrationApi) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// SetAmount sets field value
func (o *LedgerTransactionLogEntryIntegrationApi) SetAmount(v float32) {
	o.Amount = v
}

// GetId returns the Id field value
func (o *LedgerTransactionLogEntryIntegrationApi) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *LedgerTransactionLogEntryIntegrationApi) SetId(v int32) {
	o.Id = v
}

// GetRulesetId returns the RulesetId field value if set, zero value otherwise.
func (o *LedgerTransactionLogEntryIntegrationApi) GetRulesetId() int32 {
	if o == nil || o.RulesetId == nil {
		var ret int32
		return ret
	}
	return *o.RulesetId
}

// GetRulesetIdOk returns a tuple with the RulesetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionLogEntryIntegrationApi) GetRulesetIdOk() (int32, bool) {
	if o == nil || o.RulesetId == nil {
		var ret int32
		return ret, false
	}
	return *o.RulesetId, true
}

// HasRulesetId returns a boolean if a field has been set.
func (o *LedgerTransactionLogEntryIntegrationApi) HasRulesetId() bool {
	if o != nil && o.RulesetId != nil {
		return true
	}

	return false
}

// SetRulesetId gets a reference to the given int32 and assigns it to the RulesetId field.
func (o *LedgerTransactionLogEntryIntegrationApi) SetRulesetId(v int32) {
	o.RulesetId = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *LedgerTransactionLogEntryIntegrationApi) GetRuleName() string {
	if o == nil || o.RuleName == nil {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionLogEntryIntegrationApi) GetRuleNameOk() (string, bool) {
	if o == nil || o.RuleName == nil {
		var ret string
		return ret, false
	}
	return *o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *LedgerTransactionLogEntryIntegrationApi) HasRuleName() bool {
	if o != nil && o.RuleName != nil {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *LedgerTransactionLogEntryIntegrationApi) SetRuleName(v string) {
	o.RuleName = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *LedgerTransactionLogEntryIntegrationApi) GetFlags() LoyaltyLedgerEntryFlags {
	if o == nil || o.Flags == nil {
		var ret LoyaltyLedgerEntryFlags
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LedgerTransactionLogEntryIntegrationApi) GetFlagsOk() (LoyaltyLedgerEntryFlags, bool) {
	if o == nil || o.Flags == nil {
		var ret LoyaltyLedgerEntryFlags
		return ret, false
	}
	return *o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *LedgerTransactionLogEntryIntegrationApi) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given LoyaltyLedgerEntryFlags and assigns it to the Flags field.
func (o *LedgerTransactionLogEntryIntegrationApi) SetFlags(v LoyaltyLedgerEntryFlags) {
	o.Flags = &v
}

type NullableLedgerTransactionLogEntryIntegrationApi struct {
	Value        LedgerTransactionLogEntryIntegrationApi
	ExplicitNull bool
}

func (v NullableLedgerTransactionLogEntryIntegrationApi) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLedgerTransactionLogEntryIntegrationApi) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
