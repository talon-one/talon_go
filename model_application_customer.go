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

// ApplicationCustomer 
type ApplicationCustomer struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created. The exact moment this entity was created. The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID used for this entity in the application system. The ID used for this entity in the application system.
	IntegrationId string `json:"integrationId"`
	// Arbitrary properties associated with this item
	Attributes map[string]interface{} `json:"attributes"`
	// The ID of the Talon.One account that owns this profile. The ID of the Talon.One account that owns this profile.
	AccountId int32 `json:"accountId"`
	// The total amount of closed sessions by a customer. A closed session is a successful purchase.
	ClosedSessions int32 `json:"closedSessions"`
	// Sum of all purchases made by this customer
	TotalSales float32 `json:"totalSales"`
	// A list of loyalty programs joined by the customer
	LoyaltyMemberships *[]LoyaltyMembership `json:"loyaltyMemberships,omitempty"`
	// Timestamp of the most recent event received from this customer
	LastActivity time.Time `json:"lastActivity"`
}

// GetId returns the Id field value
func (o *ApplicationCustomer) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *ApplicationCustomer) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *ApplicationCustomer) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *ApplicationCustomer) SetCreated(v time.Time) {
	o.Created = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *ApplicationCustomer) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *ApplicationCustomer) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetAttributes returns the Attributes field value
func (o *ApplicationCustomer) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// SetAttributes sets field value
func (o *ApplicationCustomer) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetAccountId returns the AccountId field value
func (o *ApplicationCustomer) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *ApplicationCustomer) SetAccountId(v int32) {
	o.AccountId = v
}

// GetClosedSessions returns the ClosedSessions field value
func (o *ApplicationCustomer) GetClosedSessions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClosedSessions
}

// SetClosedSessions sets field value
func (o *ApplicationCustomer) SetClosedSessions(v int32) {
	o.ClosedSessions = v
}

// GetTotalSales returns the TotalSales field value
func (o *ApplicationCustomer) GetTotalSales() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalSales
}

// SetTotalSales sets field value
func (o *ApplicationCustomer) SetTotalSales(v float32) {
	o.TotalSales = v
}

// GetLoyaltyMemberships returns the LoyaltyMemberships field value if set, zero value otherwise.
func (o *ApplicationCustomer) GetLoyaltyMemberships() []LoyaltyMembership {
	if o == nil || o.LoyaltyMemberships == nil {
		var ret []LoyaltyMembership
		return ret
	}
	return *o.LoyaltyMemberships
}

// GetLoyaltyMembershipsOk returns a tuple with the LoyaltyMemberships field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCustomer) GetLoyaltyMembershipsOk() ([]LoyaltyMembership, bool) {
	if o == nil || o.LoyaltyMemberships == nil {
		var ret []LoyaltyMembership
		return ret, false
	}
	return *o.LoyaltyMemberships, true
}

// HasLoyaltyMemberships returns a boolean if a field has been set.
func (o *ApplicationCustomer) HasLoyaltyMemberships() bool {
	if o != nil && o.LoyaltyMemberships != nil {
		return true
	}

	return false
}

// SetLoyaltyMemberships gets a reference to the given []LoyaltyMembership and assigns it to the LoyaltyMemberships field.
func (o *ApplicationCustomer) SetLoyaltyMemberships(v []LoyaltyMembership) {
	o.LoyaltyMemberships = &v
}

// GetLastActivity returns the LastActivity field value
func (o *ApplicationCustomer) GetLastActivity() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastActivity
}

// SetLastActivity sets field value
func (o *ApplicationCustomer) SetLastActivity(v time.Time) {
	o.LastActivity = v
}

type NullableApplicationCustomer struct {
	Value ApplicationCustomer
	ExplicitNull bool
}

func (v NullableApplicationCustomer) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationCustomer) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
