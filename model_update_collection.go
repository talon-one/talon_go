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

// UpdateCollection struct for UpdateCollection
type UpdateCollection struct {
	// A short description of the purpose of this collection.
	Description *string `json:"description,omitempty"`
	// A list of the IDs of the Applications where this collection is enabled.
	SubscribedApplicationsIds *[]int32 `json:"subscribedApplicationsIds,omitempty"`
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCollection) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCollection) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateCollection) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCollection) SetDescription(v string) {
	o.Description = &v
}

// GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field value if set, zero value otherwise.
func (o *UpdateCollection) GetSubscribedApplicationsIds() []int32 {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret
	}
	return *o.SubscribedApplicationsIds
}

// GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCollection) GetSubscribedApplicationsIdsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.SubscribedApplicationsIds, true
}

// HasSubscribedApplicationsIds returns a boolean if a field has been set.
func (o *UpdateCollection) HasSubscribedApplicationsIds() bool {
	if o != nil && o.SubscribedApplicationsIds != nil {
		return true
	}

	return false
}

// SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.
func (o *UpdateCollection) SetSubscribedApplicationsIds(v []int32) {
	o.SubscribedApplicationsIds = &v
}

type NullableUpdateCollection struct {
	Value UpdateCollection
	ExplicitNull bool
}

func (v NullableUpdateCollection) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateCollection) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
