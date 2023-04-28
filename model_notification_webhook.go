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

// NotificationWebhook
type NotificationWebhook struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The time this entity was last modified.
	Modified time.Time `json:"modified"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// API URL for the given webhook-based notification.
	Url string `json:"url"`
	// List of API HTTP headers for the given webhook-based notification.
	Headers []string `json:"headers"`
}

// GetId returns the Id field value
func (o *NotificationWebhook) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *NotificationWebhook) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *NotificationWebhook) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *NotificationWebhook) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *NotificationWebhook) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// SetModified sets field value
func (o *NotificationWebhook) SetModified(v time.Time) {
	o.Modified = v
}

// GetApplicationId returns the ApplicationId field value
func (o *NotificationWebhook) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *NotificationWebhook) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetUrl returns the Url field value
func (o *NotificationWebhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// SetUrl sets field value
func (o *NotificationWebhook) SetUrl(v string) {
	o.Url = v
}

// GetHeaders returns the Headers field value
func (o *NotificationWebhook) GetHeaders() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Headers
}

// SetHeaders sets field value
func (o *NotificationWebhook) SetHeaders(v []string) {
	o.Headers = v
}

type NullableNotificationWebhook struct {
	Value        NotificationWebhook
	ExplicitNull bool
}

func (v NullableNotificationWebhook) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNotificationWebhook) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
