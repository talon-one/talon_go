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

// AccountDashboardStatisticCampaigns struct for AccountDashboardStatisticCampaigns
type AccountDashboardStatisticCampaigns struct {
	// Number of campaigns that are active and live (across all Applications).
	Live int32 `json:"live"`
	// Campaigns scheduled to expire sometime in the next 7 days.
	EndingSoon int32 `json:"endingSoon"`
	// Campaigns with less than 10% of budget left.
	LowOnBudget int32 `json:"lowOnBudget"`
}

// GetLive returns the Live field value
func (o *AccountDashboardStatisticCampaigns) GetLive() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Live
}

// SetLive sets field value
func (o *AccountDashboardStatisticCampaigns) SetLive(v int32) {
	o.Live = v
}

// GetEndingSoon returns the EndingSoon field value
func (o *AccountDashboardStatisticCampaigns) GetEndingSoon() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndingSoon
}

// SetEndingSoon sets field value
func (o *AccountDashboardStatisticCampaigns) SetEndingSoon(v int32) {
	o.EndingSoon = v
}

// GetLowOnBudget returns the LowOnBudget field value
func (o *AccountDashboardStatisticCampaigns) GetLowOnBudget() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LowOnBudget
}

// SetLowOnBudget sets field value
func (o *AccountDashboardStatisticCampaigns) SetLowOnBudget(v int32) {
	o.LowOnBudget = v
}

type NullableAccountDashboardStatisticCampaigns struct {
	Value AccountDashboardStatisticCampaigns
	ExplicitNull bool
}

func (v NullableAccountDashboardStatisticCampaigns) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAccountDashboardStatisticCampaigns) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
