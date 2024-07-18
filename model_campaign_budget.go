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

// CampaignBudget struct for CampaignBudget
type CampaignBudget struct {
	// The limitable action to which this limit applies. For example: - `setDiscount` - `setDiscountEffect` - `redeemCoupon` - `createCoupon` 
	Action string `json:"action"`
	// The value to set for the limit.
	Limit float32 `json:"limit"`
	// The number of occurrences of the limited action in the context of the campaign.
	Counter float32 `json:"counter"`
}

// GetAction returns the Action field value
func (o *CampaignBudget) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// SetAction sets field value
func (o *CampaignBudget) SetAction(v string) {
	o.Action = v
}

// GetLimit returns the Limit field value
func (o *CampaignBudget) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// SetLimit sets field value
func (o *CampaignBudget) SetLimit(v float32) {
	o.Limit = v
}

// GetCounter returns the Counter field value
func (o *CampaignBudget) GetCounter() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Counter
}

// SetCounter sets field value
func (o *CampaignBudget) SetCounter(v float32) {
	o.Counter = v
}

type NullableCampaignBudget struct {
	Value CampaignBudget
	ExplicitNull bool
}

func (v NullableCampaignBudget) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignBudget) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
