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
	"time"
)

// StoreAllOf struct for StoreAllOf
type StoreAllOf struct {
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Timestamp of the most recent update on this entity.
	Updated time.Time `json:"updated"`
	// A list of IDs of the campaigns that are linked with current store.
	LinkedCampaignIds *[]int32 `json:"linkedCampaignIds,omitempty"`
}

// GetCreated returns the Created field value
func (o *StoreAllOf) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *StoreAllOf) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *StoreAllOf) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *StoreAllOf) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetUpdated returns the Updated field value
func (o *StoreAllOf) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// SetUpdated sets field value
func (o *StoreAllOf) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetLinkedCampaignIds returns the LinkedCampaignIds field value if set, zero value otherwise.
func (o *StoreAllOf) GetLinkedCampaignIds() []int32 {
	if o == nil || o.LinkedCampaignIds == nil {
		var ret []int32
		return ret
	}
	return *o.LinkedCampaignIds
}

// GetLinkedCampaignIdsOk returns a tuple with the LinkedCampaignIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StoreAllOf) GetLinkedCampaignIdsOk() ([]int32, bool) {
	if o == nil || o.LinkedCampaignIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.LinkedCampaignIds, true
}

// HasLinkedCampaignIds returns a boolean if a field has been set.
func (o *StoreAllOf) HasLinkedCampaignIds() bool {
	if o != nil && o.LinkedCampaignIds != nil {
		return true
	}

	return false
}

// SetLinkedCampaignIds gets a reference to the given []int32 and assigns it to the LinkedCampaignIds field.
func (o *StoreAllOf) SetLinkedCampaignIds(v []int32) {
	o.LinkedCampaignIds = &v
}

type NullableStoreAllOf struct {
	Value StoreAllOf
	ExplicitNull bool
}

func (v NullableStoreAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableStoreAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
