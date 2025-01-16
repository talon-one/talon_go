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

// ValueMap struct for ValueMap
type ValueMap struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id      int32      `json:"id"`
	Created *time.Time `json:"created,omitempty"`
	// The ID of the user who created the value map.
	CreatedBy  *int32 `json:"createdBy,omitempty"`
	CampaignId int32  `json:"campaignId"`
}

// GetId returns the Id field value
func (o *ValueMap) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *ValueMap) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ValueMap) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ValueMap) GetCreatedOk() (time.Time, bool) {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ValueMap) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ValueMap) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ValueMap) GetCreatedBy() int32 {
	if o == nil || o.CreatedBy == nil {
		var ret int32
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ValueMap) GetCreatedByOk() (int32, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret int32
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ValueMap) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.
func (o *ValueMap) SetCreatedBy(v int32) {
	o.CreatedBy = &v
}

// GetCampaignId returns the CampaignId field value
func (o *ValueMap) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *ValueMap) SetCampaignId(v int32) {
	o.CampaignId = v
}

type NullableValueMap struct {
	Value        ValueMap
	ExplicitNull bool
}

func (v NullableValueMap) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableValueMap) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
