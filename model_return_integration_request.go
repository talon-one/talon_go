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

// ReturnIntegrationRequest The body of a return integration API request. Next to the cart items details, this contains an optional listing of extra properties that should be returned in the response.
type ReturnIntegrationRequest struct {
	Return NewReturn `json:"return"`
	// Extends the response with the chosen data entities. Use this property to get as much data as you need in one _Update customer session_ request instead of sending extra requests to other endpoints.
	ResponseContent *[]string `json:"responseContent,omitempty"`
}

// GetReturn returns the Return field value
func (o *ReturnIntegrationRequest) GetReturn() NewReturn {
	if o == nil {
		var ret NewReturn
		return ret
	}

	return o.Return
}

// SetReturn sets field value
func (o *ReturnIntegrationRequest) SetReturn(v NewReturn) {
	o.Return = v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *ReturnIntegrationRequest) GetResponseContent() []string {
	if o == nil || o.ResponseContent == nil {
		var ret []string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReturnIntegrationRequest) GetResponseContentOk() ([]string, bool) {
	if o == nil || o.ResponseContent == nil {
		var ret []string
		return ret, false
	}
	return *o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *ReturnIntegrationRequest) HasResponseContent() bool {
	if o != nil && o.ResponseContent != nil {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given []string and assigns it to the ResponseContent field.
func (o *ReturnIntegrationRequest) SetResponseContent(v []string) {
	o.ResponseContent = &v
}

type NullableReturnIntegrationRequest struct {
	Value        ReturnIntegrationRequest
	ExplicitNull bool
}

func (v NullableReturnIntegrationRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReturnIntegrationRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
