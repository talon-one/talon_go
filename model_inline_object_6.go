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

// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	// The file containing the data that is being imported.
	UpFile *string `json:"upFile,omitempty"`
}

// GetUpFile returns the UpFile field value if set, zero value otherwise.
func (o *InlineObject6) GetUpFile() string {
	if o == nil || o.UpFile == nil {
		var ret string
		return ret
	}
	return *o.UpFile
}

// GetUpFileOk returns a tuple with the UpFile field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetUpFileOk() (string, bool) {
	if o == nil || o.UpFile == nil {
		var ret string
		return ret, false
	}
	return *o.UpFile, true
}

// HasUpFile returns a boolean if a field has been set.
func (o *InlineObject6) HasUpFile() bool {
	if o != nil && o.UpFile != nil {
		return true
	}

	return false
}

// SetUpFile gets a reference to the given string and assigns it to the UpFile field.
func (o *InlineObject6) SetUpFile(v string) {
	o.UpFile = &v
}

type NullableInlineObject6 struct {
	Value InlineObject6
	ExplicitNull bool
}

func (v NullableInlineObject6) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineObject6) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
