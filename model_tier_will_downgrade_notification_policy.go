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

// TierWillDowngradeNotificationPolicy struct for TierWillDowngradeNotificationPolicy
type TierWillDowngradeNotificationPolicy struct {
	// Indicates whether batching is activated.
	BatchingEnabled *bool `json:"batchingEnabled,omitempty"`
	// The name of the notification.
	Name string `json:"name"`
	Triggers []TierWillDowngradeNotificationTrigger `json:"triggers"`
}

// GetBatchingEnabled returns the BatchingEnabled field value if set, zero value otherwise.
func (o *TierWillDowngradeNotificationPolicy) GetBatchingEnabled() bool {
	if o == nil || o.BatchingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BatchingEnabled
}

// GetBatchingEnabledOk returns a tuple with the BatchingEnabled field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TierWillDowngradeNotificationPolicy) GetBatchingEnabledOk() (bool, bool) {
	if o == nil || o.BatchingEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.BatchingEnabled, true
}

// HasBatchingEnabled returns a boolean if a field has been set.
func (o *TierWillDowngradeNotificationPolicy) HasBatchingEnabled() bool {
	if o != nil && o.BatchingEnabled != nil {
		return true
	}

	return false
}

// SetBatchingEnabled gets a reference to the given bool and assigns it to the BatchingEnabled field.
func (o *TierWillDowngradeNotificationPolicy) SetBatchingEnabled(v bool) {
	o.BatchingEnabled = &v
}

// GetName returns the Name field value
func (o *TierWillDowngradeNotificationPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *TierWillDowngradeNotificationPolicy) SetName(v string) {
	o.Name = v
}

// GetTriggers returns the Triggers field value
func (o *TierWillDowngradeNotificationPolicy) GetTriggers() []TierWillDowngradeNotificationTrigger {
	if o == nil {
		var ret []TierWillDowngradeNotificationTrigger
		return ret
	}

	return o.Triggers
}

// SetTriggers sets field value
func (o *TierWillDowngradeNotificationPolicy) SetTriggers(v []TierWillDowngradeNotificationTrigger) {
	o.Triggers = v
}

type NullableTierWillDowngradeNotificationPolicy struct {
	Value TierWillDowngradeNotificationPolicy
	ExplicitNull bool
}

func (v NullableTierWillDowngradeNotificationPolicy) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTierWillDowngradeNotificationPolicy) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
