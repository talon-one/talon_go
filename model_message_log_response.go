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

// MessageLogResponse Details of the response.
type MessageLogResponse struct {
	// Timestamp when the response was received.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Raw response data.
	Response *string `json:"response,omitempty"`
	// HTTP status code of the response.
	Status *int32 `json:"status,omitempty"`
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MessageLogResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MessageLogResponse) GetCreatedAtOk() (time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MessageLogResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MessageLogResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *MessageLogResponse) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MessageLogResponse) GetResponseOk() (string, bool) {
	if o == nil || o.Response == nil {
		var ret string
		return ret, false
	}
	return *o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *MessageLogResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *MessageLogResponse) SetResponse(v string) {
	o.Response = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MessageLogResponse) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MessageLogResponse) GetStatusOk() (int32, bool) {
	if o == nil || o.Status == nil {
		var ret int32
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MessageLogResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *MessageLogResponse) SetStatus(v int32) {
	o.Status = &v
}

type NullableMessageLogResponse struct {
	Value MessageLogResponse
	ExplicitNull bool
}

func (v NullableMessageLogResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMessageLogResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
