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

// BulkApplicationNotification struct for BulkApplicationNotification
type BulkApplicationNotification struct {
	TotalResultSize int32                     `json:"totalResultSize"`
	Data            []ApplicationNotification `json:"data"`
}

// GetTotalResultSize returns the TotalResultSize field value
func (o *BulkApplicationNotification) GetTotalResultSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResultSize
}

// SetTotalResultSize sets field value
func (o *BulkApplicationNotification) SetTotalResultSize(v int32) {
	o.TotalResultSize = v
}

// GetData returns the Data field value
func (o *BulkApplicationNotification) GetData() []ApplicationNotification {
	if o == nil {
		var ret []ApplicationNotification
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *BulkApplicationNotification) SetData(v []ApplicationNotification) {
	o.Data = v
}

type NullableBulkApplicationNotification struct {
	Value        BulkApplicationNotification
	ExplicitNull bool
}

func (v NullableBulkApplicationNotification) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBulkApplicationNotification) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
