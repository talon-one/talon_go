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

// AccountDashboardStatisticDiscount struct for AccountDashboardStatisticDiscount
type AccountDashboardStatisticDiscount struct {
	// Average discount percentage.
	Average float32 `json:"average"`
	// Values aggregated for the specified date.
	Datetime time.Time `json:"datetime"`
	// Total discount value redeemed by users.
	Total float32 `json:"total"`
}

// GetAverage returns the Average field value
func (o *AccountDashboardStatisticDiscount) GetAverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Average
}

// SetAverage sets field value
func (o *AccountDashboardStatisticDiscount) SetAverage(v float32) {
	o.Average = v
}

// GetDatetime returns the Datetime field value
func (o *AccountDashboardStatisticDiscount) GetDatetime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Datetime
}

// SetDatetime sets field value
func (o *AccountDashboardStatisticDiscount) SetDatetime(v time.Time) {
	o.Datetime = v
}

// GetTotal returns the Total field value
func (o *AccountDashboardStatisticDiscount) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// SetTotal sets field value
func (o *AccountDashboardStatisticDiscount) SetTotal(v float32) {
	o.Total = v
}

type NullableAccountDashboardStatisticDiscount struct {
	Value AccountDashboardStatisticDiscount
	ExplicitNull bool
}

func (v NullableAccountDashboardStatisticDiscount) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAccountDashboardStatisticDiscount) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
