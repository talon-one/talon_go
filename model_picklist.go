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

// Picklist struct for Picklist
type Picklist struct {
	// The ID of the account that owns this entity.
	AccountId *int32 `json:"accountId,omitempty"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// ID of the user who created this effect.
	CreatedBy int32 `json:"createdBy"`
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// Imported flag shows that a picklist is imported by a CSV file or not
	Imported *bool `json:"imported,omitempty"`
	// ID of the user who last updated this effect if available.
	ModifiedBy *int32 `json:"modifiedBy,omitempty"`
	// The type of allowed values in the picklist. If the type `time` is chosen, it must be an RFC3339 timestamp string.
	Type string `json:"type"`
	// The list of allowed values provided by this picklist.
	Values []string `json:"values"`
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Picklist) GetAccountId() int32 {
	if o == nil || o.AccountId == nil {
		var ret int32
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Picklist) GetAccountIdOk() (int32, bool) {
	if o == nil || o.AccountId == nil {
		var ret int32
		return ret, false
	}
	return *o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Picklist) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.
func (o *Picklist) SetAccountId(v int32) {
	o.AccountId = &v
}

// GetCreated returns the Created field value
func (o *Picklist) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Picklist) SetCreated(v time.Time) {
	o.Created = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Picklist) GetCreatedBy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedBy
}

// SetCreatedBy sets field value
func (o *Picklist) SetCreatedBy(v int32) {
	o.CreatedBy = v
}

// GetId returns the Id field value
func (o *Picklist) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Picklist) SetId(v int32) {
	o.Id = v
}

// GetImported returns the Imported field value if set, zero value otherwise.
func (o *Picklist) GetImported() bool {
	if o == nil || o.Imported == nil {
		var ret bool
		return ret
	}
	return *o.Imported
}

// GetImportedOk returns a tuple with the Imported field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Picklist) GetImportedOk() (bool, bool) {
	if o == nil || o.Imported == nil {
		var ret bool
		return ret, false
	}
	return *o.Imported, true
}

// HasImported returns a boolean if a field has been set.
func (o *Picklist) HasImported() bool {
	if o != nil && o.Imported != nil {
		return true
	}

	return false
}

// SetImported gets a reference to the given bool and assigns it to the Imported field.
func (o *Picklist) SetImported(v bool) {
	o.Imported = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *Picklist) GetModifiedBy() int32 {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Picklist) GetModifiedByOk() (int32, bool) {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret, false
	}
	return *o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *Picklist) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.
func (o *Picklist) SetModifiedBy(v int32) {
	o.ModifiedBy = &v
}

// GetType returns the Type field value
func (o *Picklist) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *Picklist) SetType(v string) {
	o.Type = v
}

// GetValues returns the Values field value
func (o *Picklist) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// SetValues sets field value
func (o *Picklist) SetValues(v []string) {
	o.Values = v
}

type NullablePicklist struct {
	Value Picklist
	ExplicitNull bool
}

func (v NullablePicklist) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullablePicklist) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
