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

// NewMultipleAudiencesItem struct for NewMultipleAudiencesItem
type NewMultipleAudiencesItem struct {
	// The human-friendly display name for this audience.
	Name string `json:"name"`
	// The ID of this audience in the third-party integration.
	IntegrationId *string `json:"integrationId,omitempty"`
}

// GetName returns the Name field value
func (o *NewMultipleAudiencesItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewMultipleAudiencesItem) SetName(v string) {
	o.Name = v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *NewMultipleAudiencesItem) GetIntegrationId() string {
	if o == nil || o.IntegrationId == nil {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewMultipleAudiencesItem) GetIntegrationIdOk() (string, bool) {
	if o == nil || o.IntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *NewMultipleAudiencesItem) HasIntegrationId() bool {
	if o != nil && o.IntegrationId != nil {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *NewMultipleAudiencesItem) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

type NullableNewMultipleAudiencesItem struct {
	Value        NewMultipleAudiencesItem
	ExplicitNull bool
}

func (v NullableNewMultipleAudiencesItem) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewMultipleAudiencesItem) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
