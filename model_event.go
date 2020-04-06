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

// Event
type Event struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID.
	ProfileId *string `json:"profileId,omitempty"`
	// A string representing the event. Must not be a reserved event name.
	Type string `json:"type"`
	// Arbitrary additional JSON data associated with the event.
	Attributes map[string]interface{} `json:"attributes"`
	// The ID of the session that this event occurred in.
	SessionId *string `json:"sessionId,omitempty"`
	// An array of \"effects\" that must be applied in response to this event. Example effects include `addItemToCart` or `setDiscount`.
	Effects []interface{} `json:"effects"`
	// Ledger entries for the event.
	LedgerEntries []LedgerEntry `json:"ledgerEntries"`
	Meta          *Meta         `json:"meta,omitempty"`
}

// GetId returns the Id field value
func (o *Event) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Event) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Event) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Event) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *Event) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *Event) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *Event) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetProfileIdOk() (string, bool) {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *Event) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *Event) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetType returns the Type field value
func (o *Event) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *Event) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *Event) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// SetAttributes sets field value
func (o *Event) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *Event) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSessionIdOk() (string, bool) {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret, false
	}
	return *o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *Event) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *Event) SetSessionId(v string) {
	o.SessionId = &v
}

// GetEffects returns the Effects field value
func (o *Event) GetEffects() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Effects
}

// SetEffects sets field value
func (o *Event) SetEffects(v []interface{}) {
	o.Effects = v
}

// GetLedgerEntries returns the LedgerEntries field value
func (o *Event) GetLedgerEntries() []LedgerEntry {
	if o == nil {
		var ret []LedgerEntry
		return ret
	}

	return o.LedgerEntries
}

// SetLedgerEntries sets field value
func (o *Event) SetLedgerEntries(v []LedgerEntry) {
	o.LedgerEntries = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Event) GetMeta() Meta {
	if o == nil || o.Meta == nil {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetMetaOk() (Meta, bool) {
	if o == nil || o.Meta == nil {
		var ret Meta
		return ret, false
	}
	return *o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Event) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *Event) SetMeta(v Meta) {
	o.Meta = &v
}

type NullableEvent struct {
	Value        Event
	ExplicitNull bool
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
