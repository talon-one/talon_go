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

// AccountDashboardStatistic struct for AccountDashboardStatistic
type AccountDashboardStatistic struct {
	// Aggregated statistic for account revenue.
	Revenue *[]AccountDashboardStatisticRevenue `json:"revenue,omitempty"`
	// Aggregated statistic for account discount.
	Discounts *[]AccountDashboardStatisticDiscount `json:"discounts,omitempty"`
	// Aggregated statistic for account loyalty points.
	LoyaltyPoints *[]AccountDashboardStatisticLoyaltyPoints `json:"loyaltyPoints,omitempty"`
	// Aggregated statistic for account referrals.
	Referrals *[]AccountDashboardStatisticReferrals `json:"referrals,omitempty"`
	Campaigns AccountDashboardStatisticCampaigns `json:"campaigns"`
}

// GetRevenue returns the Revenue field value if set, zero value otherwise.
func (o *AccountDashboardStatistic) GetRevenue() []AccountDashboardStatisticRevenue {
	if o == nil || o.Revenue == nil {
		var ret []AccountDashboardStatisticRevenue
		return ret
	}
	return *o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDashboardStatistic) GetRevenueOk() ([]AccountDashboardStatisticRevenue, bool) {
	if o == nil || o.Revenue == nil {
		var ret []AccountDashboardStatisticRevenue
		return ret, false
	}
	return *o.Revenue, true
}

// HasRevenue returns a boolean if a field has been set.
func (o *AccountDashboardStatistic) HasRevenue() bool {
	if o != nil && o.Revenue != nil {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given []AccountDashboardStatisticRevenue and assigns it to the Revenue field.
func (o *AccountDashboardStatistic) SetRevenue(v []AccountDashboardStatisticRevenue) {
	o.Revenue = &v
}

// GetDiscounts returns the Discounts field value if set, zero value otherwise.
func (o *AccountDashboardStatistic) GetDiscounts() []AccountDashboardStatisticDiscount {
	if o == nil || o.Discounts == nil {
		var ret []AccountDashboardStatisticDiscount
		return ret
	}
	return *o.Discounts
}

// GetDiscountsOk returns a tuple with the Discounts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDashboardStatistic) GetDiscountsOk() ([]AccountDashboardStatisticDiscount, bool) {
	if o == nil || o.Discounts == nil {
		var ret []AccountDashboardStatisticDiscount
		return ret, false
	}
	return *o.Discounts, true
}

// HasDiscounts returns a boolean if a field has been set.
func (o *AccountDashboardStatistic) HasDiscounts() bool {
	if o != nil && o.Discounts != nil {
		return true
	}

	return false
}

// SetDiscounts gets a reference to the given []AccountDashboardStatisticDiscount and assigns it to the Discounts field.
func (o *AccountDashboardStatistic) SetDiscounts(v []AccountDashboardStatisticDiscount) {
	o.Discounts = &v
}

// GetLoyaltyPoints returns the LoyaltyPoints field value if set, zero value otherwise.
func (o *AccountDashboardStatistic) GetLoyaltyPoints() []AccountDashboardStatisticLoyaltyPoints {
	if o == nil || o.LoyaltyPoints == nil {
		var ret []AccountDashboardStatisticLoyaltyPoints
		return ret
	}
	return *o.LoyaltyPoints
}

// GetLoyaltyPointsOk returns a tuple with the LoyaltyPoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDashboardStatistic) GetLoyaltyPointsOk() ([]AccountDashboardStatisticLoyaltyPoints, bool) {
	if o == nil || o.LoyaltyPoints == nil {
		var ret []AccountDashboardStatisticLoyaltyPoints
		return ret, false
	}
	return *o.LoyaltyPoints, true
}

// HasLoyaltyPoints returns a boolean if a field has been set.
func (o *AccountDashboardStatistic) HasLoyaltyPoints() bool {
	if o != nil && o.LoyaltyPoints != nil {
		return true
	}

	return false
}

// SetLoyaltyPoints gets a reference to the given []AccountDashboardStatisticLoyaltyPoints and assigns it to the LoyaltyPoints field.
func (o *AccountDashboardStatistic) SetLoyaltyPoints(v []AccountDashboardStatisticLoyaltyPoints) {
	o.LoyaltyPoints = &v
}

// GetReferrals returns the Referrals field value if set, zero value otherwise.
func (o *AccountDashboardStatistic) GetReferrals() []AccountDashboardStatisticReferrals {
	if o == nil || o.Referrals == nil {
		var ret []AccountDashboardStatisticReferrals
		return ret
	}
	return *o.Referrals
}

// GetReferralsOk returns a tuple with the Referrals field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AccountDashboardStatistic) GetReferralsOk() ([]AccountDashboardStatisticReferrals, bool) {
	if o == nil || o.Referrals == nil {
		var ret []AccountDashboardStatisticReferrals
		return ret, false
	}
	return *o.Referrals, true
}

// HasReferrals returns a boolean if a field has been set.
func (o *AccountDashboardStatistic) HasReferrals() bool {
	if o != nil && o.Referrals != nil {
		return true
	}

	return false
}

// SetReferrals gets a reference to the given []AccountDashboardStatisticReferrals and assigns it to the Referrals field.
func (o *AccountDashboardStatistic) SetReferrals(v []AccountDashboardStatisticReferrals) {
	o.Referrals = &v
}

// GetCampaigns returns the Campaigns field value
func (o *AccountDashboardStatistic) GetCampaigns() AccountDashboardStatisticCampaigns {
	if o == nil {
		var ret AccountDashboardStatisticCampaigns
		return ret
	}

	return o.Campaigns
}

// SetCampaigns sets field value
func (o *AccountDashboardStatistic) SetCampaigns(v AccountDashboardStatisticCampaigns) {
	o.Campaigns = v
}

type NullableAccountDashboardStatistic struct {
	Value AccountDashboardStatistic
	ExplicitNull bool
}

func (v NullableAccountDashboardStatistic) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAccountDashboardStatistic) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
