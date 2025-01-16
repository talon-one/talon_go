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

// IntegrationEvent struct for IntegrationEvent
type IntegrationEvent struct {
	// ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known `profileId`, we recommend you use a guest `profileId`.
	ProfileId *string `json:"profileId,omitempty"`
	// The integration ID of the store. You choose this ID when you create a store.
	StoreIntegrationId *string `json:"storeIntegrationId,omitempty"`
	// A string representing the event. Must not be a reserved event name.
	Type string `json:"type"`
	// Arbitrary additional JSON data associated with the event.
	Attributes map[string]interface{} `json:"attributes"`
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *IntegrationEvent) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEvent) GetProfileIdOk() (string, bool) {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *IntegrationEvent) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *IntegrationEvent) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetStoreIntegrationId returns the StoreIntegrationId field value if set, zero value otherwise.
func (o *IntegrationEvent) GetStoreIntegrationId() string {
	if o == nil || o.StoreIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.StoreIntegrationId
}

// GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEvent) GetStoreIntegrationIdOk() (string, bool) {
	if o == nil || o.StoreIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.StoreIntegrationId, true
}

// HasStoreIntegrationId returns a boolean if a field has been set.
func (o *IntegrationEvent) HasStoreIntegrationId() bool {
	if o != nil && o.StoreIntegrationId != nil {
		return true
	}

	return false
}

// SetStoreIntegrationId gets a reference to the given string and assigns it to the StoreIntegrationId field.
func (o *IntegrationEvent) SetStoreIntegrationId(v string) {
	o.StoreIntegrationId = &v
}

// GetType returns the Type field value
func (o *IntegrationEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *IntegrationEvent) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *IntegrationEvent) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// SetAttributes sets field value
func (o *IntegrationEvent) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

type NullableIntegrationEvent struct {
	Value        IntegrationEvent
	ExplicitNull bool
}

func (v NullableIntegrationEvent) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableIntegrationEvent) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
