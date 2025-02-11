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

// Change struct for Change
type Change struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The ID of the user associated with this entity.
	UserId int32 `json:"userId"`
	// ID of application associated with change.
	ApplicationId *int32 `json:"applicationId,omitempty"`
	// API endpoint on which the change was initiated.
	Entity string `json:"entity"`
	// Resource before the change occurred.
	Old *map[string]interface{} `json:"old,omitempty"`
	// Resource after the change occurred.
	New *map[string]interface{} `json:"new,omitempty"`
	// ID of management key used to perform changes.
	ManagementKeyId *int32 `json:"managementKeyId,omitempty"`
}

// GetId returns the Id field value
func (o *Change) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Change) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Change) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Change) SetCreated(v time.Time) {
	o.Created = v
}

// GetUserId returns the UserId field value
func (o *Change) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// SetUserId sets field value
func (o *Change) SetUserId(v int32) {
	o.UserId = v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *Change) GetApplicationId() int32 {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetApplicationIdOk() (int32, bool) {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret, false
	}
	return *o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *Change) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.
func (o *Change) SetApplicationId(v int32) {
	o.ApplicationId = &v
}

// GetEntity returns the Entity field value
func (o *Change) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// SetEntity sets field value
func (o *Change) SetEntity(v string) {
	o.Entity = v
}

// GetOld returns the Old field value if set, zero value otherwise.
func (o *Change) GetOld() map[string]interface{} {
	if o == nil || o.Old == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Old
}

// GetOldOk returns a tuple with the Old field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetOldOk() (map[string]interface{}, bool) {
	if o == nil || o.Old == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Old, true
}

// HasOld returns a boolean if a field has been set.
func (o *Change) HasOld() bool {
	if o != nil && o.Old != nil {
		return true
	}

	return false
}

// SetOld gets a reference to the given map[string]interface{} and assigns it to the Old field.
func (o *Change) SetOld(v map[string]interface{}) {
	o.Old = &v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *Change) GetNew() map[string]interface{} {
	if o == nil || o.New == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetNewOk() (map[string]interface{}, bool) {
	if o == nil || o.New == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *Change) HasNew() bool {
	if o != nil && o.New != nil {
		return true
	}

	return false
}

// SetNew gets a reference to the given map[string]interface{} and assigns it to the New field.
func (o *Change) SetNew(v map[string]interface{}) {
	o.New = &v
}

// GetManagementKeyId returns the ManagementKeyId field value if set, zero value otherwise.
func (o *Change) GetManagementKeyId() int32 {
	if o == nil || o.ManagementKeyId == nil {
		var ret int32
		return ret
	}
	return *o.ManagementKeyId
}

// GetManagementKeyIdOk returns a tuple with the ManagementKeyId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetManagementKeyIdOk() (int32, bool) {
	if o == nil || o.ManagementKeyId == nil {
		var ret int32
		return ret, false
	}
	return *o.ManagementKeyId, true
}

// HasManagementKeyId returns a boolean if a field has been set.
func (o *Change) HasManagementKeyId() bool {
	if o != nil && o.ManagementKeyId != nil {
		return true
	}

	return false
}

// SetManagementKeyId gets a reference to the given int32 and assigns it to the ManagementKeyId field.
func (o *Change) SetManagementKeyId(v int32) {
	o.ManagementKeyId = &v
}

type NullableChange struct {
	Value        Change
	ExplicitNull bool
}

func (v NullableChange) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableChange) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
