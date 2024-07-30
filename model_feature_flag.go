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

// FeatureFlag struct for FeatureFlag
type FeatureFlag struct {
	// The name of the feature flag.
	Name string `json:"name"`
	// The value of the feature flag.
	Value string `json:"value"`
	// The time this entity was last created.
	Created *time.Time `json:"created,omitempty"`
	// The time this entity was last modified.
	Modified *time.Time `json:"modified,omitempty"`
}

// GetName returns the Name field value
func (o *FeatureFlag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *FeatureFlag) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *FeatureFlag) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *FeatureFlag) SetValue(v string) {
	o.Value = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FeatureFlag) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlag) GetCreatedOk() (time.Time, bool) {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FeatureFlag) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *FeatureFlag) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *FeatureFlag) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlag) GetModifiedOk() (time.Time, bool) {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *FeatureFlag) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *FeatureFlag) SetModified(v time.Time) {
	o.Modified = &v
}

type NullableFeatureFlag struct {
	Value FeatureFlag
	ExplicitNull bool
}

func (v NullableFeatureFlag) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFeatureFlag) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
