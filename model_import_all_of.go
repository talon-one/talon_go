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

// ImportAllOf struct for ImportAllOf
type ImportAllOf struct {
	// The name of the entity that was imported. 
	Entity string `json:"entity"`
	// The number of values that were imported.
	Amount int32 `json:"amount"`
}

// GetEntity returns the Entity field value
func (o *ImportAllOf) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// SetEntity sets field value
func (o *ImportAllOf) SetEntity(v string) {
	o.Entity = v
}

// GetAmount returns the Amount field value
func (o *ImportAllOf) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// SetAmount sets field value
func (o *ImportAllOf) SetAmount(v int32) {
	o.Amount = v
}

type NullableImportAllOf struct {
	Value ImportAllOf
	ExplicitNull bool
}

func (v NullableImportAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableImportAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
