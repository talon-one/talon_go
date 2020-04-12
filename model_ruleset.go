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
	"time"
)

// Ruleset
type Ruleset struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the campaign that owns this entity.
	CampaignId int32 `json:"campaignId"`
	// The ID of the account that owns this entity.
	UserId int32 `json:"userId"`
	// Set of rules to apply.
	Rules []Rule `json:"rules"`
	// An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array.
	Bindings []Binding `json:"bindings"`
	// A string indicating which version of the rulebuilder was used to create this ruleset.
	RbVersion *string `json:"rbVersion,omitempty"`
	// A boolean indicating whether this newly created ruleset should also be activated for the campaign owns it
	Activate *bool `json:"activate,omitempty"`
	// Timestamp indicating when this Ruleset was activated.
	ActivatedAt *time.Time `json:"activatedAt,omitempty"`
}

// GetId returns the Id field value
func (o *Ruleset) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Ruleset) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Ruleset) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Ruleset) SetCreated(v time.Time) {
	o.Created = v
}

// GetCampaignId returns the CampaignId field value
func (o *Ruleset) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *Ruleset) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetUserId returns the UserId field value
func (o *Ruleset) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// SetUserId sets field value
func (o *Ruleset) SetUserId(v int32) {
	o.UserId = v
}

// GetRules returns the Rules field value
func (o *Ruleset) GetRules() []Rule {
	if o == nil {
		var ret []Rule
		return ret
	}

	return o.Rules
}

// SetRules sets field value
func (o *Ruleset) SetRules(v []Rule) {
	o.Rules = v
}

// GetBindings returns the Bindings field value
func (o *Ruleset) GetBindings() []Binding {
	if o == nil {
		var ret []Binding
		return ret
	}

	return o.Bindings
}

// SetBindings sets field value
func (o *Ruleset) SetBindings(v []Binding) {
	o.Bindings = v
}

// GetRbVersion returns the RbVersion field value if set, zero value otherwise.
func (o *Ruleset) GetRbVersion() string {
	if o == nil || o.RbVersion == nil {
		var ret string
		return ret
	}
	return *o.RbVersion
}

// GetRbVersionOk returns a tuple with the RbVersion field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ruleset) GetRbVersionOk() (string, bool) {
	if o == nil || o.RbVersion == nil {
		var ret string
		return ret, false
	}
	return *o.RbVersion, true
}

// HasRbVersion returns a boolean if a field has been set.
func (o *Ruleset) HasRbVersion() bool {
	if o != nil && o.RbVersion != nil {
		return true
	}

	return false
}

// SetRbVersion gets a reference to the given string and assigns it to the RbVersion field.
func (o *Ruleset) SetRbVersion(v string) {
	o.RbVersion = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *Ruleset) GetActivate() bool {
	if o == nil || o.Activate == nil {
		var ret bool
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ruleset) GetActivateOk() (bool, bool) {
	if o == nil || o.Activate == nil {
		var ret bool
		return ret, false
	}
	return *o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *Ruleset) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given bool and assigns it to the Activate field.
func (o *Ruleset) SetActivate(v bool) {
	o.Activate = &v
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise.
func (o *Ruleset) GetActivatedAt() time.Time {
	if o == nil || o.ActivatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ActivatedAt
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Ruleset) GetActivatedAtOk() (time.Time, bool) {
	if o == nil || o.ActivatedAt == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ActivatedAt, true
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *Ruleset) HasActivatedAt() bool {
	if o != nil && o.ActivatedAt != nil {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given time.Time and assigns it to the ActivatedAt field.
func (o *Ruleset) SetActivatedAt(v time.Time) {
	o.ActivatedAt = &v
}

type NullableRuleset struct {
	Value        Ruleset
	ExplicitNull bool
}

func (v NullableRuleset) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRuleset) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
