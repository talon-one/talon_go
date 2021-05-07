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

// Audience struct for Audience
type Audience struct {
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The human-friendly display name for this Audience.
	Name string `json:"name"`
	// Integration that this audience was created in.
	Integration string `json:"integration"`
	// The ID of this Audience in the third-party integration
	IntegrationId string `json:"integrationId"`
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

// GetIntegration returns the Integration field value
func (o *Audience) GetIntegration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Integration
}

// SetIntegration sets field value
func (o *Audience) SetIntegration(v string) {
	o.Integration = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *Audience) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *Audience) SetIntegrationId(v string) {
	o.IntegrationId = v
}

type NullableAudience struct {
	Value        Audience
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
