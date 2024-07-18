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
	"time"
)

// ApplicationApiHealth Report of health of the API connection of an application.
type ApplicationApiHealth struct {
	// One-word summary of the health of the API connection of an application. Possible values are: - `OK`: The Application has received only successful API requests in the last 5 minutes. - `WARNING`: The Application received at least one failed request in the last 50 minutes. - `ERROR`: More than 50% of received requests failed. - `CRITICAL`: All received requests failed. - `NONE`: During the last 5 minutes, the Application hasn't recorded any integration API requests. 
	Summary string `json:"summary"`
	// time of last request relevant to the API health test.
	LastUsed time.Time `json:"lastUsed"`
}

// GetSummary returns the Summary field value
func (o *ApplicationApiHealth) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// SetSummary sets field value
func (o *ApplicationApiHealth) SetSummary(v string) {
	o.Summary = v
}

// GetLastUsed returns the LastUsed field value
func (o *ApplicationApiHealth) GetLastUsed() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUsed
}

// SetLastUsed sets field value
func (o *ApplicationApiHealth) SetLastUsed(v time.Time) {
	o.LastUsed = v
}

type NullableApplicationApiHealth struct {
	Value ApplicationApiHealth
	ExplicitNull bool
}

func (v NullableApplicationApiHealth) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationApiHealth) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
