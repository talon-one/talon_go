/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// AccountAnalytics struct for AccountAnalytics
type AccountAnalytics struct {
	// Total number of applications in the account
	Applications int32 `json:"applications"`
	// Total number of live applications in the account
	LiveApplications int32 `json:"liveApplications"`
	// Total number of sandbox applications in the account
	SandboxApplications int32 `json:"sandboxApplications"`
	// Total number of campaigns in the account
	Campaigns int32 `json:"campaigns"`
	// Total number of active campaigns in the account
	ActiveCampaigns int32 `json:"activeCampaigns"`
	// Total number of active campaigns in live applications in the account
	LiveActiveCampaigns int32 `json:"liveActiveCampaigns"`
	// Total number of coupons in the account
	Coupons int32 `json:"coupons"`
	// Total number of active coupons in the account
	ActiveCoupons int32 `json:"activeCoupons"`
	// Total number of expired coupons in the account
	ExpiredCoupons int32 `json:"expiredCoupons"`
	// Total number of referral codes in the account
	ReferralCodes int32 `json:"referralCodes"`
	// Total number of active referral codes in the account
	ActiveReferralCodes int32 `json:"activeReferralCodes"`
	// Total number of expired referral codes in the account
	ExpiredReferralCodes int32 `json:"expiredReferralCodes"`
	// Total number of active rules in the account
	ActiveRules int32 `json:"activeRules"`
	// Total number of users in the account
	Users int32 `json:"users"`
	// Total number of roles in the account
	Roles int32 `json:"roles"`
	// Total number of custom attributes in the account
	CustomAttributes int32 `json:"customAttributes"`
	// Total number of webhooks in the account
	Webhooks int32 `json:"webhooks"`
	// Total number of loyalty programs in the account
	LoyaltyPrograms int32 `json:"loyaltyPrograms"`
}

// GetApplications returns the Applications field value
func (o *AccountAnalytics) GetApplications() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Applications
}

// SetApplications sets field value
func (o *AccountAnalytics) SetApplications(v int32) {
	o.Applications = v
}

// GetLiveApplications returns the LiveApplications field value
func (o *AccountAnalytics) GetLiveApplications() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LiveApplications
}

// SetLiveApplications sets field value
func (o *AccountAnalytics) SetLiveApplications(v int32) {
	o.LiveApplications = v
}

// GetSandboxApplications returns the SandboxApplications field value
func (o *AccountAnalytics) GetSandboxApplications() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SandboxApplications
}

// SetSandboxApplications sets field value
func (o *AccountAnalytics) SetSandboxApplications(v int32) {
	o.SandboxApplications = v
}

// GetCampaigns returns the Campaigns field value
func (o *AccountAnalytics) GetCampaigns() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Campaigns
}

// SetCampaigns sets field value
func (o *AccountAnalytics) SetCampaigns(v int32) {
	o.Campaigns = v
}

// GetActiveCampaigns returns the ActiveCampaigns field value
func (o *AccountAnalytics) GetActiveCampaigns() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActiveCampaigns
}

// SetActiveCampaigns sets field value
func (o *AccountAnalytics) SetActiveCampaigns(v int32) {
	o.ActiveCampaigns = v
}

// GetLiveActiveCampaigns returns the LiveActiveCampaigns field value
func (o *AccountAnalytics) GetLiveActiveCampaigns() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LiveActiveCampaigns
}

// SetLiveActiveCampaigns sets field value
func (o *AccountAnalytics) SetLiveActiveCampaigns(v int32) {
	o.LiveActiveCampaigns = v
}

// GetCoupons returns the Coupons field value
func (o *AccountAnalytics) GetCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Coupons
}

// SetCoupons sets field value
func (o *AccountAnalytics) SetCoupons(v int32) {
	o.Coupons = v
}

// GetActiveCoupons returns the ActiveCoupons field value
func (o *AccountAnalytics) GetActiveCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActiveCoupons
}

// SetActiveCoupons sets field value
func (o *AccountAnalytics) SetActiveCoupons(v int32) {
	o.ActiveCoupons = v
}

// GetExpiredCoupons returns the ExpiredCoupons field value
func (o *AccountAnalytics) GetExpiredCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiredCoupons
}

// SetExpiredCoupons sets field value
func (o *AccountAnalytics) SetExpiredCoupons(v int32) {
	o.ExpiredCoupons = v
}

// GetReferralCodes returns the ReferralCodes field value
func (o *AccountAnalytics) GetReferralCodes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferralCodes
}

// SetReferralCodes sets field value
func (o *AccountAnalytics) SetReferralCodes(v int32) {
	o.ReferralCodes = v
}

// GetActiveReferralCodes returns the ActiveReferralCodes field value
func (o *AccountAnalytics) GetActiveReferralCodes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActiveReferralCodes
}

// SetActiveReferralCodes sets field value
func (o *AccountAnalytics) SetActiveReferralCodes(v int32) {
	o.ActiveReferralCodes = v
}

// GetExpiredReferralCodes returns the ExpiredReferralCodes field value
func (o *AccountAnalytics) GetExpiredReferralCodes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiredReferralCodes
}

// SetExpiredReferralCodes sets field value
func (o *AccountAnalytics) SetExpiredReferralCodes(v int32) {
	o.ExpiredReferralCodes = v
}

// GetActiveRules returns the ActiveRules field value
func (o *AccountAnalytics) GetActiveRules() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActiveRules
}

// SetActiveRules sets field value
func (o *AccountAnalytics) SetActiveRules(v int32) {
	o.ActiveRules = v
}

// GetUsers returns the Users field value
func (o *AccountAnalytics) GetUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Users
}

// SetUsers sets field value
func (o *AccountAnalytics) SetUsers(v int32) {
	o.Users = v
}

// GetRoles returns the Roles field value
func (o *AccountAnalytics) GetRoles() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Roles
}

// SetRoles sets field value
func (o *AccountAnalytics) SetRoles(v int32) {
	o.Roles = v
}

// GetCustomAttributes returns the CustomAttributes field value
func (o *AccountAnalytics) GetCustomAttributes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CustomAttributes
}

// SetCustomAttributes sets field value
func (o *AccountAnalytics) SetCustomAttributes(v int32) {
	o.CustomAttributes = v
}

// GetWebhooks returns the Webhooks field value
func (o *AccountAnalytics) GetWebhooks() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Webhooks
}

// SetWebhooks sets field value
func (o *AccountAnalytics) SetWebhooks(v int32) {
	o.Webhooks = v
}

// GetLoyaltyPrograms returns the LoyaltyPrograms field value
func (o *AccountAnalytics) GetLoyaltyPrograms() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoyaltyPrograms
}

// SetLoyaltyPrograms sets field value
func (o *AccountAnalytics) SetLoyaltyPrograms(v int32) {
	o.LoyaltyPrograms = v
}

type NullableAccountAnalytics struct {
	Value        AccountAnalytics
	ExplicitNull bool
}

func (v NullableAccountAnalytics) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAccountAnalytics) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
