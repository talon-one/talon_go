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

// EntityWithTalangVisibleId struct for EntityWithTalangVisibleId
type EntityWithTalangVisibleId struct {
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// Unique ID for this entity.
	Id int32 `json:"id"`
}

// GetCreated returns the Created field value
func (o *EntityWithTalangVisibleId) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *EntityWithTalangVisibleId) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *EntityWithTalangVisibleId) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *EntityWithTalangVisibleId) SetId(v int32) {
	o.Id = v
}

type NullableEntityWithTalangVisibleId struct {
	Value EntityWithTalangVisibleId
	ExplicitNull bool
}

func (v NullableEntityWithTalangVisibleId) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableEntityWithTalangVisibleId) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
