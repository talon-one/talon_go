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

// NewAdditionalCost struct for NewAdditionalCost
type NewAdditionalCost struct {
	// A description of this additional cost.
	Description string `json:"description"`
	// The internal name used in API requests.
	Name string `json:"name"`
	// A list of the IDs of the applications that are subscribed to this additional cost.
	SubscribedApplicationsIds *[]int32 `json:"subscribedApplicationsIds,omitempty"`
	// The human-readable name for the additional cost that will be shown in the Campaign Manager. Like `name`, the combination of entity and title must also be unique.
	Title string `json:"title"`
	// The type of additional cost. Possible value: - `session`: Additional cost will be added per session. - `item`: Additional cost will be added per item. - `both`: Additional cost will be added per item and session. 
	Type *string `json:"type,omitempty"`
}

// GetDescription returns the Description field value
func (o *NewAdditionalCost) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *NewAdditionalCost) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *NewAdditionalCost) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewAdditionalCost) SetName(v string) {
	o.Name = v
}

// GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field value if set, zero value otherwise.
func (o *NewAdditionalCost) GetSubscribedApplicationsIds() []int32 {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret
	}
	return *o.SubscribedApplicationsIds
}

// GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewAdditionalCost) GetSubscribedApplicationsIdsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.SubscribedApplicationsIds, true
}

// HasSubscribedApplicationsIds returns a boolean if a field has been set.
func (o *NewAdditionalCost) HasSubscribedApplicationsIds() bool {
	if o != nil && o.SubscribedApplicationsIds != nil {
		return true
	}

	return false
}

// SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.
func (o *NewAdditionalCost) SetSubscribedApplicationsIds(v []int32) {
	o.SubscribedApplicationsIds = &v
}

// GetTitle returns the Title field value
func (o *NewAdditionalCost) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *NewAdditionalCost) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NewAdditionalCost) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewAdditionalCost) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NewAdditionalCost) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NewAdditionalCost) SetType(v string) {
	o.Type = &v
}

type NullableNewAdditionalCost struct {
	Value NewAdditionalCost
	ExplicitNull bool
}

func (v NullableNewAdditionalCost) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewAdditionalCost) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
