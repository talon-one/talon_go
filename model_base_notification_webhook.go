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

// BaseNotificationWebhook struct for BaseNotificationWebhook
type BaseNotificationWebhook struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The time this entity was last modified.
	Modified time.Time `json:"modified"`
	// API URL for the given webhook-based notification.
	Url string `json:"url"`
	// List of API HTTP headers for the given webhook-based notification.
	Headers []string `json:"headers"`
	// Indicates whether the notification is activated.
	Enabled *bool `json:"enabled,omitempty"`
}

// GetId returns the Id field value
func (o *BaseNotificationWebhook) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *BaseNotificationWebhook) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *BaseNotificationWebhook) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *BaseNotificationWebhook) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *BaseNotificationWebhook) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// SetModified sets field value
func (o *BaseNotificationWebhook) SetModified(v time.Time) {
	o.Modified = v
}

// GetUrl returns the Url field value
func (o *BaseNotificationWebhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// SetUrl sets field value
func (o *BaseNotificationWebhook) SetUrl(v string) {
	o.Url = v
}

// GetHeaders returns the Headers field value
func (o *BaseNotificationWebhook) GetHeaders() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Headers
}

// SetHeaders sets field value
func (o *BaseNotificationWebhook) SetHeaders(v []string) {
	o.Headers = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BaseNotificationWebhook) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *BaseNotificationWebhook) GetEnabledOk() (bool, bool) {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret, false
	}
	return *o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BaseNotificationWebhook) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BaseNotificationWebhook) SetEnabled(v bool) {
	o.Enabled = &v
}

type NullableBaseNotificationWebhook struct {
	Value        BaseNotificationWebhook
	ExplicitNull bool
}

func (v NullableBaseNotificationWebhook) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBaseNotificationWebhook) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
