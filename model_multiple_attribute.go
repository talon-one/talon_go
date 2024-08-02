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

// MultipleAttribute struct for MultipleAttribute
type MultipleAttribute struct {
	Attributes *[]Attribute `json:"attributes,omitempty"`
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MultipleAttribute) GetAttributes() []Attribute {
	if o == nil || o.Attributes == nil {
		var ret []Attribute
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MultipleAttribute) GetAttributesOk() ([]Attribute, bool) {
	if o == nil || o.Attributes == nil {
		var ret []Attribute
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MultipleAttribute) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Attribute and assigns it to the Attributes field.
func (o *MultipleAttribute) SetAttributes(v []Attribute) {
	o.Attributes = &v
}

type NullableMultipleAttribute struct {
	Value MultipleAttribute
	ExplicitNull bool
}

func (v NullableMultipleAttribute) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMultipleAttribute) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
