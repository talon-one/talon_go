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

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// A message describing the error.
	Message string `json:"message"`
	// An array of individual problems encountered during the request.
	Errors *[]ApiError `json:"errors,omitempty"`
}

// GetMessage returns the Message field value
func (o *ErrorResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// SetMessage sets field value
func (o *ErrorResponse) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ErrorResponse) GetErrors() []ApiError {
	if o == nil || o.Errors == nil {
		var ret []ApiError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetErrorsOk() ([]ApiError, bool) {
	if o == nil || o.Errors == nil {
		var ret []ApiError
		return ret, false
	}
	return *o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ErrorResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ApiError and assigns it to the Errors field.
func (o *ErrorResponse) SetErrors(v []ApiError) {
	o.Errors = &v
}

type NullableErrorResponse struct {
	Value ErrorResponse
	ExplicitNull bool
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
