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
)

// RemoveItemCatalogAction The specific properties of the \"REMOVE\" catalog sync action.
type RemoveItemCatalogAction struct {
	// The unique SKU of the item to remove.
	Sku string `json:"sku"`
}

// GetSku returns the Sku field value
func (o *RemoveItemCatalogAction) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// SetSku sets field value
func (o *RemoveItemCatalogAction) SetSku(v string) {
	o.Sku = v
}

type NullableRemoveItemCatalogAction struct {
	Value RemoveItemCatalogAction
	ExplicitNull bool
}

func (v NullableRemoveItemCatalogAction) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRemoveItemCatalogAction) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
