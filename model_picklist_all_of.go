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
)

// PicklistAllOf struct for PicklistAllOf
type PicklistAllOf struct {
	// ID of the user who last updated this effect if available.
	ModifiedBy *int32 `json:"modifiedBy,omitempty"`
	// ID of the user who created this effect.
	CreatedBy int32 `json:"createdBy"`
	// The ID of the account that owns this entity.
	AccountId *int32 `json:"accountId,omitempty"`
	// Imported flag shows that a picklist is imported by a CSV file or not
	Imported *bool `json:"imported,omitempty"`
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *PicklistAllOf) GetModifiedBy() int32 {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PicklistAllOf) GetModifiedByOk() (int32, bool) {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret, false
	}
	return *o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *PicklistAllOf) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.
func (o *PicklistAllOf) SetModifiedBy(v int32) {
	o.ModifiedBy = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *PicklistAllOf) GetCreatedBy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedBy
}

// SetCreatedBy sets field value
func (o *PicklistAllOf) SetCreatedBy(v int32) {
	o.CreatedBy = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PicklistAllOf) GetAccountId() int32 {
	if o == nil || o.AccountId == nil {
		var ret int32
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PicklistAllOf) GetAccountIdOk() (int32, bool) {
	if o == nil || o.AccountId == nil {
		var ret int32
		return ret, false
	}
	return *o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PicklistAllOf) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.
func (o *PicklistAllOf) SetAccountId(v int32) {
	o.AccountId = &v
}

// GetImported returns the Imported field value if set, zero value otherwise.
func (o *PicklistAllOf) GetImported() bool {
	if o == nil || o.Imported == nil {
		var ret bool
		return ret
	}
	return *o.Imported
}

// GetImportedOk returns a tuple with the Imported field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PicklistAllOf) GetImportedOk() (bool, bool) {
	if o == nil || o.Imported == nil {
		var ret bool
		return ret, false
	}
	return *o.Imported, true
}

// HasImported returns a boolean if a field has been set.
func (o *PicklistAllOf) HasImported() bool {
	if o != nil && o.Imported != nil {
		return true
	}

	return false
}

// SetImported gets a reference to the given bool and assigns it to the Imported field.
func (o *PicklistAllOf) SetImported(v bool) {
	o.Imported = &v
}

type NullablePicklistAllOf struct {
	Value PicklistAllOf
	ExplicitNull bool
}

func (v NullablePicklistAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullablePicklistAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
