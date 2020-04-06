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

// CustomerProfileSearchQuery struct for CustomerProfileSearchQuery
type CustomerProfileSearchQuery struct {
	// Properties to match against a customer profile. All provided attributes will be exactly matched against profile attributes
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	IntegrationIDs *[]string `json:"integrationIDs,omitempty"`
	ProfileIDs *[]int32 `json:"profileIDs,omitempty"`
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerProfileSearchQuery) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileSearchQuery) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerProfileSearchQuery) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *CustomerProfileSearchQuery) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetIntegrationIDs returns the IntegrationIDs field value if set, zero value otherwise.
func (o *CustomerProfileSearchQuery) GetIntegrationIDs() []string {
	if o == nil || o.IntegrationIDs == nil {
		var ret []string
		return ret
	}
	return *o.IntegrationIDs
}

// GetIntegrationIDsOk returns a tuple with the IntegrationIDs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileSearchQuery) GetIntegrationIDsOk() ([]string, bool) {
	if o == nil || o.IntegrationIDs == nil {
		var ret []string
		return ret, false
	}
	return *o.IntegrationIDs, true
}

// HasIntegrationIDs returns a boolean if a field has been set.
func (o *CustomerProfileSearchQuery) HasIntegrationIDs() bool {
	if o != nil && o.IntegrationIDs != nil {
		return true
	}

	return false
}

// SetIntegrationIDs gets a reference to the given []string and assigns it to the IntegrationIDs field.
func (o *CustomerProfileSearchQuery) SetIntegrationIDs(v []string) {
	o.IntegrationIDs = &v
}

// GetProfileIDs returns the ProfileIDs field value if set, zero value otherwise.
func (o *CustomerProfileSearchQuery) GetProfileIDs() []int32 {
	if o == nil || o.ProfileIDs == nil {
		var ret []int32
		return ret
	}
	return *o.ProfileIDs
}

// GetProfileIDsOk returns a tuple with the ProfileIDs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileSearchQuery) GetProfileIDsOk() ([]int32, bool) {
	if o == nil || o.ProfileIDs == nil {
		var ret []int32
		return ret, false
	}
	return *o.ProfileIDs, true
}

// HasProfileIDs returns a boolean if a field has been set.
func (o *CustomerProfileSearchQuery) HasProfileIDs() bool {
	if o != nil && o.ProfileIDs != nil {
		return true
	}

	return false
}

// SetProfileIDs gets a reference to the given []int32 and assigns it to the ProfileIDs field.
func (o *CustomerProfileSearchQuery) SetProfileIDs(v []int32) {
	o.ProfileIDs = &v
}

type NullableCustomerProfileSearchQuery struct {
	Value CustomerProfileSearchQuery
	ExplicitNull bool
}

func (v NullableCustomerProfileSearchQuery) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomerProfileSearchQuery) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
