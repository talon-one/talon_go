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

// CatalogsStrikethroughNotificationPolicy struct for CatalogsStrikethroughNotificationPolicy
type CatalogsStrikethroughNotificationPolicy struct {
	// Notification name.
	Name string `json:"name"`
	// The number of days in advance that strikethrough pricing updates should be sent.
	AheadOfDaysTrigger *int32 `json:"aheadOfDaysTrigger,omitempty"`
}

// GetName returns the Name field value
func (o *CatalogsStrikethroughNotificationPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CatalogsStrikethroughNotificationPolicy) SetName(v string) {
	o.Name = v
}

// GetAheadOfDaysTrigger returns the AheadOfDaysTrigger field value if set, zero value otherwise.
func (o *CatalogsStrikethroughNotificationPolicy) GetAheadOfDaysTrigger() int32 {
	if o == nil || o.AheadOfDaysTrigger == nil {
		var ret int32
		return ret
	}
	return *o.AheadOfDaysTrigger
}

// GetAheadOfDaysTriggerOk returns a tuple with the AheadOfDaysTrigger field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatalogsStrikethroughNotificationPolicy) GetAheadOfDaysTriggerOk() (int32, bool) {
	if o == nil || o.AheadOfDaysTrigger == nil {
		var ret int32
		return ret, false
	}
	return *o.AheadOfDaysTrigger, true
}

// HasAheadOfDaysTrigger returns a boolean if a field has been set.
func (o *CatalogsStrikethroughNotificationPolicy) HasAheadOfDaysTrigger() bool {
	if o != nil && o.AheadOfDaysTrigger != nil {
		return true
	}

	return false
}

// SetAheadOfDaysTrigger gets a reference to the given int32 and assigns it to the AheadOfDaysTrigger field.
func (o *CatalogsStrikethroughNotificationPolicy) SetAheadOfDaysTrigger(v int32) {
	o.AheadOfDaysTrigger = &v
}

type NullableCatalogsStrikethroughNotificationPolicy struct {
	Value        CatalogsStrikethroughNotificationPolicy
	ExplicitNull bool
}

func (v NullableCatalogsStrikethroughNotificationPolicy) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCatalogsStrikethroughNotificationPolicy) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
