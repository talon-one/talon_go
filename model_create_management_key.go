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

// CreateManagementKey struct for CreateManagementKey
type CreateManagementKey struct {
	// Name for management key.
	Name string `json:"name"`
	// The date the management key expires.
	ExpiryDate time.Time `json:"expiryDate"`
	// The list of endpoints that can be accessed with the key
	Endpoints []Endpoint `json:"endpoints"`
	// A list of Application IDs that you can access with the management key. An empty or missing list means the management key can be used for all Applications in the account. 
	AllowedApplicationIds *[]int32 `json:"allowedApplicationIds,omitempty"`
}

// GetName returns the Name field value
func (o *CreateManagementKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CreateManagementKey) SetName(v string) {
	o.Name = v
}

// GetExpiryDate returns the ExpiryDate field value
func (o *CreateManagementKey) GetExpiryDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiryDate
}

// SetExpiryDate sets field value
func (o *CreateManagementKey) SetExpiryDate(v time.Time) {
	o.ExpiryDate = v
}

// GetEndpoints returns the Endpoints field value
func (o *CreateManagementKey) GetEndpoints() []Endpoint {
	if o == nil {
		var ret []Endpoint
		return ret
	}

	return o.Endpoints
}

// SetEndpoints sets field value
func (o *CreateManagementKey) SetEndpoints(v []Endpoint) {
	o.Endpoints = v
}

// GetAllowedApplicationIds returns the AllowedApplicationIds field value if set, zero value otherwise.
func (o *CreateManagementKey) GetAllowedApplicationIds() []int32 {
	if o == nil || o.AllowedApplicationIds == nil {
		var ret []int32
		return ret
	}
	return *o.AllowedApplicationIds
}

// GetAllowedApplicationIdsOk returns a tuple with the AllowedApplicationIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateManagementKey) GetAllowedApplicationIdsOk() ([]int32, bool) {
	if o == nil || o.AllowedApplicationIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.AllowedApplicationIds, true
}

// HasAllowedApplicationIds returns a boolean if a field has been set.
func (o *CreateManagementKey) HasAllowedApplicationIds() bool {
	if o != nil && o.AllowedApplicationIds != nil {
		return true
	}

	return false
}

// SetAllowedApplicationIds gets a reference to the given []int32 and assigns it to the AllowedApplicationIds field.
func (o *CreateManagementKey) SetAllowedApplicationIds(v []int32) {
	o.AllowedApplicationIds = &v
}

type NullableCreateManagementKey struct {
	Value CreateManagementKey
	ExplicitNull bool
}

func (v NullableCreateManagementKey) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateManagementKey) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
