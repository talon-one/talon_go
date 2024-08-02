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

// NewEventType struct for NewEventType
type NewEventType struct {
	// A description of what the event represents. 
	Description *string `json:"description,omitempty"`
	// The integration name for this event type. This will be used in URLs and cannot be changed after an event type has been created.
	Name string `json:"name"`
	// The human-friendly name for this event type.
	Title string `json:"title"`
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewEventType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewEventType) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewEventType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewEventType) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *NewEventType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewEventType) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *NewEventType) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *NewEventType) SetTitle(v string) {
	o.Title = v
}

type NullableNewEventType struct {
	Value NewEventType
	ExplicitNull bool
}

func (v NullableNewEventType) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewEventType) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
