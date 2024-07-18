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

// CatalogActionFilter The properties for a single filtering condition in a catalog sync action.
type CatalogActionFilter struct {
	// The name of the attribute to filter on.
	Attr string `json:"attr"`
	// The filtering operator.
	Op string `json:"op"`
	// The value to filter for.
	Value map[string]interface{} `json:"value"`
}

// GetAttr returns the Attr field value
func (o *CatalogActionFilter) GetAttr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attr
}

// SetAttr sets field value
func (o *CatalogActionFilter) SetAttr(v string) {
	o.Attr = v
}

// GetOp returns the Op field value
func (o *CatalogActionFilter) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// SetOp sets field value
func (o *CatalogActionFilter) SetOp(v string) {
	o.Op = v
}

// GetValue returns the Value field value
func (o *CatalogActionFilter) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *CatalogActionFilter) SetValue(v map[string]interface{}) {
	o.Value = v
}

type NullableCatalogActionFilter struct {
	Value CatalogActionFilter
	ExplicitNull bool
}

func (v NullableCatalogActionFilter) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCatalogActionFilter) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
