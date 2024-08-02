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

// GenerateItemFilterDescription struct for GenerateItemFilterDescription
type GenerateItemFilterDescription struct {
	// An array of item filter Talang expressions.
	ItemFilter [][]map[string]interface{} `json:"itemFilter"`
}

// GetItemFilter returns the ItemFilter field value
func (o *GenerateItemFilterDescription) GetItemFilter() [][]map[string]interface{} {
	if o == nil {
		var ret [][]map[string]interface{}
		return ret
	}

	return o.ItemFilter
}

// SetItemFilter sets field value
func (o *GenerateItemFilterDescription) SetItemFilter(v [][]map[string]interface{}) {
	o.ItemFilter = v
}

type NullableGenerateItemFilterDescription struct {
	Value GenerateItemFilterDescription
	ExplicitNull bool
}

func (v NullableGenerateItemFilterDescription) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableGenerateItemFilterDescription) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
