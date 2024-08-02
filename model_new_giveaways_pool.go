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

// NewGiveawaysPool struct for NewGiveawaysPool
type NewGiveawaysPool struct {
	// The description of this giveaways pool.
	Description *string `json:"description,omitempty"`
	// The name of this giveaways pool.
	Name string `json:"name"`
	// Indicates if this program is a live or sandbox program. Programs of a given type can only be connected to Applications of the same type.
	Sandbox bool `json:"sandbox"`
	// A list of the IDs of the applications that this giveaways pool is enabled for.
	SubscribedApplicationsIds *[]int32 `json:"subscribedApplicationsIds,omitempty"`
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewGiveawaysPool) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewGiveawaysPool) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewGiveawaysPool) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewGiveawaysPool) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *NewGiveawaysPool) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewGiveawaysPool) SetName(v string) {
	o.Name = v
}

// GetSandbox returns the Sandbox field value
func (o *NewGiveawaysPool) GetSandbox() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sandbox
}

// SetSandbox sets field value
func (o *NewGiveawaysPool) SetSandbox(v bool) {
	o.Sandbox = v
}

// GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field value if set, zero value otherwise.
func (o *NewGiveawaysPool) GetSubscribedApplicationsIds() []int32 {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret
	}
	return *o.SubscribedApplicationsIds
}

// GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewGiveawaysPool) GetSubscribedApplicationsIdsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.SubscribedApplicationsIds, true
}

// HasSubscribedApplicationsIds returns a boolean if a field has been set.
func (o *NewGiveawaysPool) HasSubscribedApplicationsIds() bool {
	if o != nil && o.SubscribedApplicationsIds != nil {
		return true
	}

	return false
}

// SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.
func (o *NewGiveawaysPool) SetSubscribedApplicationsIds(v []int32) {
	o.SubscribedApplicationsIds = &v
}

type NullableNewGiveawaysPool struct {
	Value NewGiveawaysPool
	ExplicitNull bool
}

func (v NullableNewGiveawaysPool) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewGiveawaysPool) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
