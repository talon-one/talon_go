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

// LoyaltyProgramLedgers Customer specific information about loyalty points.
type LoyaltyProgramLedgers struct {
	// Visible name of loyalty program
	Title string `json:"title"`
	// Internal name of loyalty program
	Name string `json:"name"`
	Ledger LoyaltyProgramBalance `json:"ledger"`
	// A map containing a list of all loyalty subledger balances
	SubLedgers *map[string]LoyaltyProgramBalance `json:"subLedgers,omitempty"`
}

// GetTitle returns the Title field value
func (o *LoyaltyProgramLedgers) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *LoyaltyProgramLedgers) SetTitle(v string) {
	o.Title = v
}

// GetName returns the Name field value
func (o *LoyaltyProgramLedgers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *LoyaltyProgramLedgers) SetName(v string) {
	o.Name = v
}

// GetLedger returns the Ledger field value
func (o *LoyaltyProgramLedgers) GetLedger() LoyaltyProgramBalance {
	if o == nil {
		var ret LoyaltyProgramBalance
		return ret
	}

	return o.Ledger
}

// SetLedger sets field value
func (o *LoyaltyProgramLedgers) SetLedger(v LoyaltyProgramBalance) {
	o.Ledger = v
}

// GetSubLedgers returns the SubLedgers field value if set, zero value otherwise.
func (o *LoyaltyProgramLedgers) GetSubLedgers() map[string]LoyaltyProgramBalance {
	if o == nil || o.SubLedgers == nil {
		var ret map[string]LoyaltyProgramBalance
		return ret
	}
	return *o.SubLedgers
}

// GetSubLedgersOk returns a tuple with the SubLedgers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyProgramLedgers) GetSubLedgersOk() (map[string]LoyaltyProgramBalance, bool) {
	if o == nil || o.SubLedgers == nil {
		var ret map[string]LoyaltyProgramBalance
		return ret, false
	}
	return *o.SubLedgers, true
}

// HasSubLedgers returns a boolean if a field has been set.
func (o *LoyaltyProgramLedgers) HasSubLedgers() bool {
	if o != nil && o.SubLedgers != nil {
		return true
	}

	return false
}

// SetSubLedgers gets a reference to the given map[string]LoyaltyProgramBalance and assigns it to the SubLedgers field.
func (o *LoyaltyProgramLedgers) SetSubLedgers(v map[string]LoyaltyProgramBalance) {
	o.SubLedgers = &v
}

type NullableLoyaltyProgramLedgers struct {
	Value LoyaltyProgramLedgers
	ExplicitNull bool
}

func (v NullableLoyaltyProgramLedgers) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyProgramLedgers) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
