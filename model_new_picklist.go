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

// NewPicklist struct for NewPicklist
type NewPicklist struct {
	// The type of allowed values in the picklist. If type time is chosen, it must be an RFC3339 timestamp string.
	Type string `json:"type"`
	// The list of allowed values provided by this picklist.
	Values []string `json:"values"`
}

// GetType returns the Type field value
func (o *NewPicklist) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *NewPicklist) SetType(v string) {
	o.Type = v
}

// GetValues returns the Values field value
func (o *NewPicklist) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// SetValues sets field value
func (o *NewPicklist) SetValues(v []string) {
	o.Values = v
}

type NullableNewPicklist struct {
	Value        NewPicklist
	ExplicitNull bool
}

func (v NullableNewPicklist) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewPicklist) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
