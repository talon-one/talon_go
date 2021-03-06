/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// MultipleCustomerProfileIntegrationResponseV2 struct for MultipleCustomerProfileIntegrationResponseV2
type MultipleCustomerProfileIntegrationResponseV2 struct {
	IntegrationStates *[]IntegrationStateV2 `json:"integrationStates,omitempty"`
}

// GetIntegrationStates returns the IntegrationStates field value if set, zero value otherwise.
func (o *MultipleCustomerProfileIntegrationResponseV2) GetIntegrationStates() []IntegrationStateV2 {
	if o == nil || o.IntegrationStates == nil {
		var ret []IntegrationStateV2
		return ret
	}
	return *o.IntegrationStates
}

// GetIntegrationStatesOk returns a tuple with the IntegrationStates field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MultipleCustomerProfileIntegrationResponseV2) GetIntegrationStatesOk() ([]IntegrationStateV2, bool) {
	if o == nil || o.IntegrationStates == nil {
		var ret []IntegrationStateV2
		return ret, false
	}
	return *o.IntegrationStates, true
}

// HasIntegrationStates returns a boolean if a field has been set.
func (o *MultipleCustomerProfileIntegrationResponseV2) HasIntegrationStates() bool {
	if o != nil && o.IntegrationStates != nil {
		return true
	}

	return false
}

// SetIntegrationStates gets a reference to the given []IntegrationStateV2 and assigns it to the IntegrationStates field.
func (o *MultipleCustomerProfileIntegrationResponseV2) SetIntegrationStates(v []IntegrationStateV2) {
	o.IntegrationStates = &v
}

type NullableMultipleCustomerProfileIntegrationResponseV2 struct {
	Value        MultipleCustomerProfileIntegrationResponseV2
	ExplicitNull bool
}

func (v NullableMultipleCustomerProfileIntegrationResponseV2) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMultipleCustomerProfileIntegrationResponseV2) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
