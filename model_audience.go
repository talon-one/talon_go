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

// Audience struct for Audience
type Audience struct {
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The human-friendly display name for this audience.
	Name string `json:"name"`
	// Indicates if this is a live or sandbox Application.
	Sandbox *bool `json:"sandbox,omitempty"`
	// A description of the audience.
	Description *string `json:"description,omitempty"`
	// The Talon.One-supported [3rd-party platform](https://docs.talon.one/docs/dev/technology-partners/overview) that this audience was created in.  For example, `mParticle`, `Segment`, `Selligent`, `Braze`, or `Iterable`.  **Note:** If you do not integrate with any of these platforms, do not use this property. 
	Integration *string `json:"integration,omitempty"`
	// The ID of this audience in the third-party integration.  **Note:** To create an audience that doesn't come from a 3rd party platform, do not use this property. 
	IntegrationId *string `json:"integrationId,omitempty"`
	// Determines if this audience is a 3rd party audience or not.
	CreatedIn3rdParty *bool `json:"createdIn3rdParty,omitempty"`
	// The last time that the audience memberships changed.
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`
}

// GetAccountId returns the AccountId field value
func (o *Audience) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *Audience) SetAccountId(v int32) {
	o.AccountId = v
}

// GetId returns the Id field value
func (o *Audience) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Audience) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Audience) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Audience) SetCreated(v time.Time) {
	o.Created = v
}

// GetName returns the Name field value
func (o *Audience) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *Audience) SetName(v string) {
	o.Name = v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *Audience) GetSandbox() bool {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Audience) GetSandboxOk() (bool, bool) {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret, false
	}
	return *o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *Audience) HasSandbox() bool {
	if o != nil && o.Sandbox != nil {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *Audience) SetSandbox(v bool) {
	o.Sandbox = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Audience) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Audience) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Audience) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Audience) SetDescription(v string) {
	o.Description = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *Audience) GetIntegration() string {
	if o == nil || o.Integration == nil {
		var ret string
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Audience) GetIntegrationOk() (string, bool) {
	if o == nil || o.Integration == nil {
		var ret string
		return ret, false
	}
	return *o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *Audience) HasIntegration() bool {
	if o != nil && o.Integration != nil {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given string and assigns it to the Integration field.
func (o *Audience) SetIntegration(v string) {
	o.Integration = &v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *Audience) GetIntegrationId() string {
	if o == nil || o.IntegrationId == nil {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Audience) GetIntegrationIdOk() (string, bool) {
	if o == nil || o.IntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *Audience) HasIntegrationId() bool {
	if o != nil && o.IntegrationId != nil {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *Audience) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetCreatedIn3rdParty returns the CreatedIn3rdParty field value if set, zero value otherwise.
func (o *Audience) GetCreatedIn3rdParty() bool {
	if o == nil || o.CreatedIn3rdParty == nil {
		var ret bool
		return ret
	}
	return *o.CreatedIn3rdParty
}

// GetCreatedIn3rdPartyOk returns a tuple with the CreatedIn3rdParty field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Audience) GetCreatedIn3rdPartyOk() (bool, bool) {
	if o == nil || o.CreatedIn3rdParty == nil {
		var ret bool
		return ret, false
	}
	return *o.CreatedIn3rdParty, true
}

// HasCreatedIn3rdParty returns a boolean if a field has been set.
func (o *Audience) HasCreatedIn3rdParty() bool {
	if o != nil && o.CreatedIn3rdParty != nil {
		return true
	}

	return false
}

// SetCreatedIn3rdParty gets a reference to the given bool and assigns it to the CreatedIn3rdParty field.
func (o *Audience) SetCreatedIn3rdParty(v bool) {
	o.CreatedIn3rdParty = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *Audience) GetLastUpdate() time.Time {
	if o == nil || o.LastUpdate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Audience) GetLastUpdateOk() (time.Time, bool) {
	if o == nil || o.LastUpdate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *Audience) HasLastUpdate() bool {
	if o != nil && o.LastUpdate != nil {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.
func (o *Audience) SetLastUpdate(v time.Time) {
	o.LastUpdate = &v
}

type NullableAudience struct {
	Value Audience
	ExplicitNull bool
}

func (v NullableAudience) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAudience) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
