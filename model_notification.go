/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// Notification struct for Notification
type Notification struct {
	// id of the notification
	Id int32 `json:"id"`
	// name of the notification
	Name string `json:"name"`
	// description of the notification
	Description string `json:"description"`
}

// GetId returns the Id field value
func (o *Notification) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Notification) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Notification) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *Notification) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Notification) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *Notification) SetDescription(v string) {
	o.Description = v
}

type NullableNotification struct {
	Value        Notification
	ExplicitNull bool
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
