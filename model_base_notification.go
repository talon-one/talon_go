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

// BaseNotification
type BaseNotification struct {
	Policy  map[string]interface{}  `json:"policy"`
	Webhook BaseNotificationWebhook `json:"webhook"`
	// Unique ID for this entity.
	Id int32 `json:"id"`
}

// GetPolicy returns the Policy field value
func (o *BaseNotification) GetPolicy() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Policy
}

// SetPolicy sets field value
func (o *BaseNotification) SetPolicy(v map[string]interface{}) {
	o.Policy = v
}

// GetWebhook returns the Webhook field value
func (o *BaseNotification) GetWebhook() BaseNotificationWebhook {
	if o == nil {
		var ret BaseNotificationWebhook
		return ret
	}

	return o.Webhook
}

// SetWebhook sets field value
func (o *BaseNotification) SetWebhook(v BaseNotificationWebhook) {
	o.Webhook = v
}

// GetId returns the Id field value
func (o *BaseNotification) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *BaseNotification) SetId(v int32) {
	o.Id = v
}

type NullableBaseNotification struct {
	Value        BaseNotification
	ExplicitNull bool
}

func (v NullableBaseNotification) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBaseNotification) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
