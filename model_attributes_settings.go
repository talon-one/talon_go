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

// AttributesSettings Arbitrary settings associated with attributes.
type AttributesSettings struct {
	Mandatory *AttributesMandatory `json:"mandatory,omitempty"`
}

// GetMandatory returns the Mandatory field value if set, zero value otherwise.
func (o *AttributesSettings) GetMandatory() AttributesMandatory {
	if o == nil || o.Mandatory == nil {
		var ret AttributesMandatory
		return ret
	}
	return *o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AttributesSettings) GetMandatoryOk() (AttributesMandatory, bool) {
	if o == nil || o.Mandatory == nil {
		var ret AttributesMandatory
		return ret, false
	}
	return *o.Mandatory, true
}

// HasMandatory returns a boolean if a field has been set.
func (o *AttributesSettings) HasMandatory() bool {
	if o != nil && o.Mandatory != nil {
		return true
	}

	return false
}

// SetMandatory gets a reference to the given AttributesMandatory and assigns it to the Mandatory field.
func (o *AttributesSettings) SetMandatory(v AttributesMandatory) {
	o.Mandatory = &v
}

type NullableAttributesSettings struct {
	Value        AttributesSettings
	ExplicitNull bool
}

func (v NullableAttributesSettings) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAttributesSettings) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
