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
	"time"
)

// NewReferralsForMultipleAdvocates struct for NewReferralsForMultipleAdvocates
type NewReferralsForMultipleAdvocates struct {
	// Timestamp at which point the referral code becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiry date of the referral code. Referral never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of times a referral code can be used. This can be set to 0 for no limit, but any campaign usage limits will still apply.
	UsageLimit int32 `json:"usageLimit"`
	// The ID of the campaign from which the referral received the referral code.
	CampaignId int32 `json:"campaignId"`
	// An array containing all the respective advocate profiles.
	AdvocateProfileIntegrationIds []string `json:"advocateProfileIntegrationIds"`
	// Arbitrary properties associated with this item.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Set of characters to be used when generating random part of code. Defaults to [A-Z, 0-9] (in terms of RegExp).
	ValidCharacters *[]string `json:"validCharacters,omitempty"`
	// The pattern that will be used to generate referrals. The character `#` acts as a placeholder and will be replaced by a random character from the `validCharacters` set.
	ReferralPattern *string `json:"referralPattern,omitempty"`
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *NewReferralsForMultipleAdvocates) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewReferralsForMultipleAdvocates) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *NewReferralsForMultipleAdvocates) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *NewReferralsForMultipleAdvocates) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *NewReferralsForMultipleAdvocates) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewReferralsForMultipleAdvocates) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *NewReferralsForMultipleAdvocates) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *NewReferralsForMultipleAdvocates) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetUsageLimit returns the UsageLimit field value
func (o *NewReferralsForMultipleAdvocates) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// SetUsageLimit sets field value
func (o *NewReferralsForMultipleAdvocates) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetCampaignId returns the CampaignId field value
func (o *NewReferralsForMultipleAdvocates) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *NewReferralsForMultipleAdvocates) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetAdvocateProfileIntegrationIds returns the AdvocateProfileIntegrationIds field value
func (o *NewReferralsForMultipleAdvocates) GetAdvocateProfileIntegrationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AdvocateProfileIntegrationIds
}

// SetAdvocateProfileIntegrationIds sets field value
func (o *NewReferralsForMultipleAdvocates) SetAdvocateProfileIntegrationIds(v []string) {
	o.AdvocateProfileIntegrationIds = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewReferralsForMultipleAdvocates) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewReferralsForMultipleAdvocates) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewReferralsForMultipleAdvocates) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewReferralsForMultipleAdvocates) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetValidCharacters returns the ValidCharacters field value if set, zero value otherwise.
func (o *NewReferralsForMultipleAdvocates) GetValidCharacters() []string {
	if o == nil || o.ValidCharacters == nil {
		var ret []string
		return ret
	}
	return *o.ValidCharacters
}

// GetValidCharactersOk returns a tuple with the ValidCharacters field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewReferralsForMultipleAdvocates) GetValidCharactersOk() ([]string, bool) {
	if o == nil || o.ValidCharacters == nil {
		var ret []string
		return ret, false
	}
	return *o.ValidCharacters, true
}

// HasValidCharacters returns a boolean if a field has been set.
func (o *NewReferralsForMultipleAdvocates) HasValidCharacters() bool {
	if o != nil && o.ValidCharacters != nil {
		return true
	}

	return false
}

// SetValidCharacters gets a reference to the given []string and assigns it to the ValidCharacters field.
func (o *NewReferralsForMultipleAdvocates) SetValidCharacters(v []string) {
	o.ValidCharacters = &v
}

// GetReferralPattern returns the ReferralPattern field value if set, zero value otherwise.
func (o *NewReferralsForMultipleAdvocates) GetReferralPattern() string {
	if o == nil || o.ReferralPattern == nil {
		var ret string
		return ret
	}
	return *o.ReferralPattern
}

// GetReferralPatternOk returns a tuple with the ReferralPattern field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewReferralsForMultipleAdvocates) GetReferralPatternOk() (string, bool) {
	if o == nil || o.ReferralPattern == nil {
		var ret string
		return ret, false
	}
	return *o.ReferralPattern, true
}

// HasReferralPattern returns a boolean if a field has been set.
func (o *NewReferralsForMultipleAdvocates) HasReferralPattern() bool {
	if o != nil && o.ReferralPattern != nil {
		return true
	}

	return false
}

// SetReferralPattern gets a reference to the given string and assigns it to the ReferralPattern field.
func (o *NewReferralsForMultipleAdvocates) SetReferralPattern(v string) {
	o.ReferralPattern = &v
}

type NullableNewReferralsForMultipleAdvocates struct {
	Value        NewReferralsForMultipleAdvocates
	ExplicitNull bool
}

func (v NullableNewReferralsForMultipleAdvocates) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewReferralsForMultipleAdvocates) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
