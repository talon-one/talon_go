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

// Event struct for Event
type Event struct {
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Arbitrary additional JSON data associated with the event.
	Attributes map[string]interface{} `json:"attributes"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// An array of effects generated by the rules of the enabled campaigns of the Application.  You decide how to apply them in your system. See the list of [API effects](https://docs.talon.one/docs/dev/integration-api/api-effects). 
	Effects [][]map[string]interface{} `json:"effects"`
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// Ledger entries for the event.
	LedgerEntries []LedgerEntry `json:"ledgerEntries"`
	Meta *Meta `json:"meta,omitempty"`
	// ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known `profileId`, we recommend you use a guest `profileId`. 
	ProfileId *string `json:"profileId,omitempty"`
	// The ID of the session that this event occurred in.
	SessionId *string `json:"sessionId,omitempty"`
	// The integration ID of the store. You choose this ID when you create a store.
	StoreIntegrationId *string `json:"storeIntegrationId,omitempty"`
	// A string representing the event. Must not be a reserved event name.
	Type string `json:"type"`
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

// GetEffects returns the Effects field value
func (o *Event) GetEffects() [][]map[string]interface{} {
	if o == nil {
		var ret [][]map[string]interface{}
		return ret
	}

	return o.Effects
}

// SetEffects sets field value
func (o *Event) SetEffects(v [][]map[string]interface{}) {
	o.Effects = v
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

// GetStoreIntegrationId returns the StoreIntegrationId field value if set, zero value otherwise.
func (o *Event) GetStoreIntegrationId() string {
	if o == nil || o.StoreIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.StoreIntegrationId
}

// GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetStoreIntegrationIdOk() (string, bool) {
	if o == nil || o.StoreIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.StoreIntegrationId, true
}

// HasStoreIntegrationId returns a boolean if a field has been set.
func (o *Event) HasStoreIntegrationId() bool {
	if o != nil && o.StoreIntegrationId != nil {
		return true
	}

	return false
}

// SetStoreIntegrationId gets a reference to the given string and assigns it to the StoreIntegrationId field.
func (o *Event) SetStoreIntegrationId(v string) {
	o.StoreIntegrationId = &v
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

type NullableEvent struct {
	Value Event
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
