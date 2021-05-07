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

// Import struct for Import
type Import struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The ID of the account that owns this entity.
	UserId int32 `json:"userId"`
	// The name of the entity that was imported. Possible values are Coupons and LoyaltyPoints.
	Entity string `json:"entity"`
	// The number of members that imported.
	Amount int32 `json:"amount"`
}

// GetId returns the Id field value
func (o *Import) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Import) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Import) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Import) SetCreated(v time.Time) {
	o.Created = v
}

// GetAccountId returns the AccountId field value
func (o *Import) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *Import) SetAccountId(v int32) {
	o.AccountId = v
}

// GetUserId returns the UserId field value
func (o *Import) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// SetUserId sets field value
func (o *Import) SetUserId(v int32) {
	o.UserId = v
}

// GetEntity returns the Entity field value
func (o *Import) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// SetEntity sets field value
func (o *Import) SetEntity(v string) {
	o.Entity = v
}

// GetAmount returns the Amount field value
func (o *Import) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// SetAmount sets field value
func (o *Import) SetAmount(v int32) {
	o.Amount = v
}

type NullableImport struct {
	Value        Import
	ExplicitNull bool
}

func (v NullableImport) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableImport) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
