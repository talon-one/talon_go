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
	"time"
)

// MultipleAudiencesItem 
type MultipleAudiencesItem struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The human-friendly display name for this audience.
	Name string `json:"name"`
	// The ID of this audience in the third-party integration.
	IntegrationId string `json:"integrationId"`
	// Indicates whether the audience is new, updated or unmodified by the request. 
	Status string `json:"status"`
}

// GetId returns the Id field value
func (o *MultipleAudiencesItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *MultipleAudiencesItem) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *MultipleAudiencesItem) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *MultipleAudiencesItem) SetCreated(v time.Time) {
	o.Created = v
}

// GetName returns the Name field value
func (o *MultipleAudiencesItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *MultipleAudiencesItem) SetName(v string) {
	o.Name = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *MultipleAudiencesItem) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *MultipleAudiencesItem) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetStatus returns the Status field value
func (o *MultipleAudiencesItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// SetStatus sets field value
func (o *MultipleAudiencesItem) SetStatus(v string) {
	o.Status = v
}

type NullableMultipleAudiencesItem struct {
	Value MultipleAudiencesItem
	ExplicitNull bool
}

func (v NullableMultipleAudiencesItem) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMultipleAudiencesItem) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
