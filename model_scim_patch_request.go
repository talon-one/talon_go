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

// ScimPatchRequest SCIM Patch request
type ScimPatchRequest struct {
	// SCIM schema for the given resource.
	Schemas    *[]string            `json:"schemas,omitempty"`
	Operations []ScimPatchOperation `json:"Operations"`
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimPatchRequest) GetSchemas() []string {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret
	}
	return *o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimPatchRequest) GetSchemasOk() ([]string, bool) {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret, false
	}
	return *o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimPatchRequest) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *ScimPatchRequest) SetSchemas(v []string) {
	o.Schemas = &v
}

// GetOperations returns the Operations field value
func (o *ScimPatchRequest) GetOperations() []ScimPatchOperation {
	if o == nil {
		var ret []ScimPatchOperation
		return ret
	}

	return o.Operations
}

// SetOperations sets field value
func (o *ScimPatchRequest) SetOperations(v []ScimPatchOperation) {
	o.Operations = v
}

type NullableScimPatchRequest struct {
	Value        ScimPatchRequest
	ExplicitNull bool
}

func (v NullableScimPatchRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableScimPatchRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
