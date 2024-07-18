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

// OutgoingIntegrationTypes struct for OutgoingIntegrationTypes
type OutgoingIntegrationTypes struct {
	// List of outgoing integrations.
	Data *[]OutgoingIntegrationType `json:"data,omitempty"`
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OutgoingIntegrationTypes) GetData() []OutgoingIntegrationType {
	if o == nil || o.Data == nil {
		var ret []OutgoingIntegrationType
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingIntegrationTypes) GetDataOk() ([]OutgoingIntegrationType, bool) {
	if o == nil || o.Data == nil {
		var ret []OutgoingIntegrationType
		return ret, false
	}
	return *o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OutgoingIntegrationTypes) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []OutgoingIntegrationType and assigns it to the Data field.
func (o *OutgoingIntegrationTypes) SetData(v []OutgoingIntegrationType) {
	o.Data = &v
}

type NullableOutgoingIntegrationTypes struct {
	Value OutgoingIntegrationTypes
	ExplicitNull bool
}

func (v NullableOutgoingIntegrationTypes) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOutgoingIntegrationTypes) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
