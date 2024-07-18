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

// ApiError struct for ApiError
type ApiError struct {
	// Short description of the problem.
	Title string `json:"title"`
	// Longer description of this specific instance of the problem.
	Details *string `json:"details,omitempty"`
	Source ErrorSource `json:"source"`
}

// GetTitle returns the Title field value
func (o *ApiError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *ApiError) SetTitle(v string) {
	o.Title = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ApiError) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetDetailsOk() (string, bool) {
	if o == nil || o.Details == nil {
		var ret string
		return ret, false
	}
	return *o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ApiError) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ApiError) SetDetails(v string) {
	o.Details = &v
}

// GetSource returns the Source field value
func (o *ApiError) GetSource() ErrorSource {
	if o == nil {
		var ret ErrorSource
		return ret
	}

	return o.Source
}

// SetSource sets field value
func (o *ApiError) SetSource(v ErrorSource) {
	o.Source = v
}

type NullableApiError struct {
	Value ApiError
	ExplicitNull bool
}

func (v NullableApiError) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApiError) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
