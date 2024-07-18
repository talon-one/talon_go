/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// ErrorResponseWithStatus struct for ErrorResponseWithStatus
type ErrorResponseWithStatus struct {
	Message *string `json:"message,omitempty"`
	// An array of individual problems encountered during the request.
	Errors *[]ApiError `json:"errors,omitempty"`
	// The error code
	StatusCode *int32 `json:"StatusCode,omitempty"`
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorResponseWithStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseWithStatus) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorResponseWithStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorResponseWithStatus) SetMessage(v string) {
	o.Message = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ErrorResponseWithStatus) GetErrors() []ApiError {
	if o == nil || o.Errors == nil {
		var ret []ApiError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseWithStatus) GetErrorsOk() ([]ApiError, bool) {
	if o == nil || o.Errors == nil {
		var ret []ApiError
		return ret, false
	}
	return *o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ErrorResponseWithStatus) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ApiError and assigns it to the Errors field.
func (o *ErrorResponseWithStatus) SetErrors(v []ApiError) {
	o.Errors = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ErrorResponseWithStatus) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseWithStatus) GetStatusCodeOk() (int32, bool) {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret, false
	}
	return *o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ErrorResponseWithStatus) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ErrorResponseWithStatus) SetStatusCode(v int32) {
	o.StatusCode = &v
}

type NullableErrorResponseWithStatus struct {
	Value ErrorResponseWithStatus
	ExplicitNull bool
}

func (v NullableErrorResponseWithStatus) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableErrorResponseWithStatus) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
