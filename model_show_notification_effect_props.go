/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerSession](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2) endpoint is `https://mycompany.talon.one/v2/customer_sessions/{Id}`
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// ShowNotificationEffectProps The properties specific to the \"showNotification\" effect. This gets triggered whenever a validated rule contained a \"show notification\" effect.
type ShowNotificationEffectProps struct {
	// The type of notification that should be shown (e.g. error/warning/info).
	NotificationType string `json:"notificationType"`
	// Title of the notification.
	Title string `json:"title"`
	// Body of the notification.
	Body string `json:"body"`
}

// GetNotificationType returns the NotificationType field value
func (o *ShowNotificationEffectProps) GetNotificationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationType
}

// SetNotificationType sets field value
func (o *ShowNotificationEffectProps) SetNotificationType(v string) {
	o.NotificationType = v
}

// GetTitle returns the Title field value
func (o *ShowNotificationEffectProps) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *ShowNotificationEffectProps) SetTitle(v string) {
	o.Title = v
}

// GetBody returns the Body field value
func (o *ShowNotificationEffectProps) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// SetBody sets field value
func (o *ShowNotificationEffectProps) SetBody(v string) {
	o.Body = v
}

type NullableShowNotificationEffectProps struct {
	Value        ShowNotificationEffectProps
	ExplicitNull bool
}

func (v NullableShowNotificationEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableShowNotificationEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
