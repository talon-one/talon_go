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

// ApplicationCampaignAnalyticsCouponsCount The number of times a coupon was successfully redeemed in influenced sessions.
type ApplicationCampaignAnalyticsCouponsCount struct {
	Value *float32 `json:"value,omitempty"`
	Trend *float32 `json:"trend,omitempty"`
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalyticsCouponsCount) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalyticsCouponsCount) GetValueOk() (float32, bool) {
	if o == nil || o.Value == nil {
		var ret float32
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalyticsCouponsCount) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *ApplicationCampaignAnalyticsCouponsCount) SetValue(v float32) {
	o.Value = &v
}

// GetTrend returns the Trend field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalyticsCouponsCount) GetTrend() float32 {
	if o == nil || o.Trend == nil {
		var ret float32
		return ret
	}
	return *o.Trend
}

// GetTrendOk returns a tuple with the Trend field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalyticsCouponsCount) GetTrendOk() (float32, bool) {
	if o == nil || o.Trend == nil {
		var ret float32
		return ret, false
	}
	return *o.Trend, true
}

// HasTrend returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalyticsCouponsCount) HasTrend() bool {
	if o != nil && o.Trend != nil {
		return true
	}

	return false
}

// SetTrend gets a reference to the given float32 and assigns it to the Trend field.
func (o *ApplicationCampaignAnalyticsCouponsCount) SetTrend(v float32) {
	o.Trend = &v
}

type NullableApplicationCampaignAnalyticsCouponsCount struct {
	Value        ApplicationCampaignAnalyticsCouponsCount
	ExplicitNull bool
}

func (v NullableApplicationCampaignAnalyticsCouponsCount) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationCampaignAnalyticsCouponsCount) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
