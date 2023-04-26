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

// ItemAttribute
type ItemAttribute struct {
	// The ID of the attribute of the item.
	Attributeid int32 `json:"attributeid"`
	// The name of the attribute.
	Name string `json:"name"`
	// The value of the attribute.
	Value map[string]interface{} `json:"value"`
}

// GetAttributeid returns the Attributeid field value
func (o *ItemAttribute) GetAttributeid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Attributeid
}

// SetAttributeid sets field value
func (o *ItemAttribute) SetAttributeid(v int32) {
	o.Attributeid = v
}

// GetName returns the Name field value
func (o *ItemAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *ItemAttribute) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *ItemAttribute) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *ItemAttribute) SetValue(v map[string]interface{}) {
	o.Value = v
}

type NullableItemAttribute struct {
	Value        ItemAttribute
	ExplicitNull bool
}

func (v NullableItemAttribute) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableItemAttribute) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
