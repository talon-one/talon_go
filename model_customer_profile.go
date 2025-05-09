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

// CustomerProfile struct for CustomerProfile
type CustomerProfile struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The integration ID set by your integration layer.
	IntegrationId string `json:"integrationId"`
	// Arbitrary properties associated with this item.
	Attributes map[string]interface{} `json:"attributes"`
	// The ID of the Talon.One account that owns this profile.
	AccountId int32 `json:"accountId"`
	// The total amount of closed sessions by a customer. A closed session is a successful purchase.
	ClosedSessions int32 `json:"closedSessions"`
	// The total amount of money spent by the customer **before** discounts are applied.  The total sales amount excludes the following: - Cancelled or reopened sessions. - Returned items.
	TotalSales float32 `json:"totalSales"`
	// **DEPRECATED** A list of loyalty programs joined by the customer.
	LoyaltyMemberships *[]LoyaltyMembership `json:"loyaltyMemberships,omitempty"`
	// The audiences the customer belongs to.
	AudienceMemberships *[]AudienceMembership `json:"audienceMemberships,omitempty"`
	// Timestamp of the most recent event received from this customer. This field is updated on calls that trigger the Rule Engine and that are not [dry requests](https://docs.talon.one/docs/dev/integration-api/dry-requests/#overlay).  For example, [reserving a coupon](https://docs.talon.one/integration-api#operation/createCouponReservation) for a customer doesn't impact this field.
	LastActivity time.Time `json:"lastActivity"`
	// An indicator of whether the customer is part of a sandbox or live Application. See the [docs](https://docs.talon.one/docs/product/applications/overview#application-environments).
	Sandbox *bool `json:"sandbox,omitempty"`
}

// GetId returns the Id field value
func (o *CustomerProfile) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *CustomerProfile) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *CustomerProfile) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *CustomerProfile) SetCreated(v time.Time) {
	o.Created = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *CustomerProfile) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *CustomerProfile) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerProfile) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// SetAttributes sets field value
func (o *CustomerProfile) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetAccountId returns the AccountId field value
func (o *CustomerProfile) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *CustomerProfile) SetAccountId(v int32) {
	o.AccountId = v
}

// GetClosedSessions returns the ClosedSessions field value
func (o *CustomerProfile) GetClosedSessions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClosedSessions
}

// SetClosedSessions sets field value
func (o *CustomerProfile) SetClosedSessions(v int32) {
	o.ClosedSessions = v
}

// GetTotalSales returns the TotalSales field value
func (o *CustomerProfile) GetTotalSales() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalSales
}

// SetTotalSales sets field value
func (o *CustomerProfile) SetTotalSales(v float32) {
	o.TotalSales = v
}

// GetLoyaltyMemberships returns the LoyaltyMemberships field value if set, zero value otherwise.
func (o *CustomerProfile) GetLoyaltyMemberships() []LoyaltyMembership {
	if o == nil || o.LoyaltyMemberships == nil {
		var ret []LoyaltyMembership
		return ret
	}
	return *o.LoyaltyMemberships
}

// GetLoyaltyMembershipsOk returns a tuple with the LoyaltyMemberships field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfile) GetLoyaltyMembershipsOk() ([]LoyaltyMembership, bool) {
	if o == nil || o.LoyaltyMemberships == nil {
		var ret []LoyaltyMembership
		return ret, false
	}
	return *o.LoyaltyMemberships, true
}

// HasLoyaltyMemberships returns a boolean if a field has been set.
func (o *CustomerProfile) HasLoyaltyMemberships() bool {
	if o != nil && o.LoyaltyMemberships != nil {
		return true
	}

	return false
}

// SetLoyaltyMemberships gets a reference to the given []LoyaltyMembership and assigns it to the LoyaltyMemberships field.
func (o *CustomerProfile) SetLoyaltyMemberships(v []LoyaltyMembership) {
	o.LoyaltyMemberships = &v
}

// GetAudienceMemberships returns the AudienceMemberships field value if set, zero value otherwise.
func (o *CustomerProfile) GetAudienceMemberships() []AudienceMembership {
	if o == nil || o.AudienceMemberships == nil {
		var ret []AudienceMembership
		return ret
	}
	return *o.AudienceMemberships
}

// GetAudienceMembershipsOk returns a tuple with the AudienceMemberships field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfile) GetAudienceMembershipsOk() ([]AudienceMembership, bool) {
	if o == nil || o.AudienceMemberships == nil {
		var ret []AudienceMembership
		return ret, false
	}
	return *o.AudienceMemberships, true
}

// HasAudienceMemberships returns a boolean if a field has been set.
func (o *CustomerProfile) HasAudienceMemberships() bool {
	if o != nil && o.AudienceMemberships != nil {
		return true
	}

	return false
}

// SetAudienceMemberships gets a reference to the given []AudienceMembership and assigns it to the AudienceMemberships field.
func (o *CustomerProfile) SetAudienceMemberships(v []AudienceMembership) {
	o.AudienceMemberships = &v
}

// GetLastActivity returns the LastActivity field value
func (o *CustomerProfile) GetLastActivity() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastActivity
}

// SetLastActivity sets field value
func (o *CustomerProfile) SetLastActivity(v time.Time) {
	o.LastActivity = v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *CustomerProfile) GetSandbox() bool {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfile) GetSandboxOk() (bool, bool) {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret, false
	}
	return *o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *CustomerProfile) HasSandbox() bool {
	if o != nil && o.Sandbox != nil {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *CustomerProfile) SetSandbox(v bool) {
	o.Sandbox = &v
}

type NullableCustomerProfile struct {
	Value        CustomerProfile
	ExplicitNull bool
}

func (v NullableCustomerProfile) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomerProfile) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
