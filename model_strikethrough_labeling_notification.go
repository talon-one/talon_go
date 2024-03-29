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

// StrikethroughLabelingNotification The strikethrough labels notification for an application.
type StrikethroughLabelingNotification struct {
	// The ID of the application that catalog items labels belongs to.
	ApplicationId int32 `json:"applicationId"`
	// The batch number of the notification. Notifications might be sent in different batches.
	CurrentBatch int32 `json:"currentBatch"`
	// The total number of batches for the notification.
	TotalBatches int32                      `json:"totalBatches"`
	Trigger      StrikethroughTrigger       `json:"trigger"`
	ChangedItems []StrikethroughChangedItem `json:"changedItems"`
}

// GetApplicationId returns the ApplicationId field value
func (o *StrikethroughLabelingNotification) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *StrikethroughLabelingNotification) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetCurrentBatch returns the CurrentBatch field value
func (o *StrikethroughLabelingNotification) GetCurrentBatch() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentBatch
}

// SetCurrentBatch sets field value
func (o *StrikethroughLabelingNotification) SetCurrentBatch(v int32) {
	o.CurrentBatch = v
}

// GetTotalBatches returns the TotalBatches field value
func (o *StrikethroughLabelingNotification) GetTotalBatches() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalBatches
}

// SetTotalBatches sets field value
func (o *StrikethroughLabelingNotification) SetTotalBatches(v int32) {
	o.TotalBatches = v
}

// GetTrigger returns the Trigger field value
func (o *StrikethroughLabelingNotification) GetTrigger() StrikethroughTrigger {
	if o == nil {
		var ret StrikethroughTrigger
		return ret
	}

	return o.Trigger
}

// SetTrigger sets field value
func (o *StrikethroughLabelingNotification) SetTrigger(v StrikethroughTrigger) {
	o.Trigger = v
}

// GetChangedItems returns the ChangedItems field value
func (o *StrikethroughLabelingNotification) GetChangedItems() []StrikethroughChangedItem {
	if o == nil {
		var ret []StrikethroughChangedItem
		return ret
	}

	return o.ChangedItems
}

// SetChangedItems sets field value
func (o *StrikethroughLabelingNotification) SetChangedItems(v []StrikethroughChangedItem) {
	o.ChangedItems = v
}

type NullableStrikethroughLabelingNotification struct {
	Value        StrikethroughLabelingNotification
	ExplicitNull bool
}

func (v NullableStrikethroughLabelingNotification) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableStrikethroughLabelingNotification) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
