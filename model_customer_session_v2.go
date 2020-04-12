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

// CustomerSessionV2
type CustomerSessionV2 struct {
	// The ID used for this entity in the application system.
	IntegrationId string `json:"integrationId"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID.
	ProfileId string `json:"profileId"`
	// Any coupon codes entered.
	CouponCodes *[]string `json:"couponCodes,omitempty"`
	// Any referral code entered.
	ReferralCode *string `json:"referralCode,omitempty"`
	// Indicates the current state of the session. All sessions must start in the \"open\" state, after which valid transitions are...  1. open -> closed 2. open -> cancelled 3. closed -> cancelled
	State string `json:"state"`
	// All items the customer will be purchasing in this session
	CartItems []CartItem `json:"cartItems"`
	// Any costs associated with the session that can not be explicitly attributed to cart items. Examples include shipping costs and service fees.
	AdditionalCosts *map[string]AdditionalCost `json:"additionalCosts,omitempty"`
	// Identifiers for the customer, this can be used for limits on values such as device ID.
	Identifiers *[]string `json:"identifiers,omitempty"`
	// A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.
	Attributes map[string]interface{} `json:"attributes"`
	// Indicates whether this is the first session for the customer's profile. Will always be true for anonymous sessions.
	FirstSession bool `json:"firstSession"`
	// The total sum of cart-items, as well as additional costs, before any discounts applied
	Total float32 `json:"total"`
	// The total sum of cart-items before any discounts applied
	CartItemTotal float32 `json:"cartItemTotal"`
	// The total sum of additional costs before any discounts applied
	AdditionalCostTotal float32 `json:"additionalCostTotal"`
}

// GetIntegrationId returns the IntegrationId field value
func (o *CustomerSessionV2) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *CustomerSessionV2) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetCreated returns the Created field value
func (o *CustomerSessionV2) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *CustomerSessionV2) SetCreated(v time.Time) {
	o.Created = v
}

// GetApplicationId returns the ApplicationId field value
func (o *CustomerSessionV2) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *CustomerSessionV2) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetProfileId returns the ProfileId field value
func (o *CustomerSessionV2) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// SetProfileId sets field value
func (o *CustomerSessionV2) SetProfileId(v string) {
	o.ProfileId = v
}

// GetCouponCodes returns the CouponCodes field value if set, zero value otherwise.
func (o *CustomerSessionV2) GetCouponCodes() []string {
	if o == nil || o.CouponCodes == nil {
		var ret []string
		return ret
	}
	return *o.CouponCodes
}

// GetCouponCodesOk returns a tuple with the CouponCodes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSessionV2) GetCouponCodesOk() ([]string, bool) {
	if o == nil || o.CouponCodes == nil {
		var ret []string
		return ret, false
	}
	return *o.CouponCodes, true
}

// HasCouponCodes returns a boolean if a field has been set.
func (o *CustomerSessionV2) HasCouponCodes() bool {
	if o != nil && o.CouponCodes != nil {
		return true
	}

	return false
}

// SetCouponCodes gets a reference to the given []string and assigns it to the CouponCodes field.
func (o *CustomerSessionV2) SetCouponCodes(v []string) {
	o.CouponCodes = &v
}

// GetReferralCode returns the ReferralCode field value if set, zero value otherwise.
func (o *CustomerSessionV2) GetReferralCode() string {
	if o == nil || o.ReferralCode == nil {
		var ret string
		return ret
	}
	return *o.ReferralCode
}

// GetReferralCodeOk returns a tuple with the ReferralCode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSessionV2) GetReferralCodeOk() (string, bool) {
	if o == nil || o.ReferralCode == nil {
		var ret string
		return ret, false
	}
	return *o.ReferralCode, true
}

// HasReferralCode returns a boolean if a field has been set.
func (o *CustomerSessionV2) HasReferralCode() bool {
	if o != nil && o.ReferralCode != nil {
		return true
	}

	return false
}

// SetReferralCode gets a reference to the given string and assigns it to the ReferralCode field.
func (o *CustomerSessionV2) SetReferralCode(v string) {
	o.ReferralCode = &v
}

// GetState returns the State field value
func (o *CustomerSessionV2) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// SetState sets field value
func (o *CustomerSessionV2) SetState(v string) {
	o.State = v
}

// GetCartItems returns the CartItems field value
func (o *CustomerSessionV2) GetCartItems() []CartItem {
	if o == nil {
		var ret []CartItem
		return ret
	}

	return o.CartItems
}

// SetCartItems sets field value
func (o *CustomerSessionV2) SetCartItems(v []CartItem) {
	o.CartItems = v
}

// GetAdditionalCosts returns the AdditionalCosts field value if set, zero value otherwise.
func (o *CustomerSessionV2) GetAdditionalCosts() map[string]AdditionalCost {
	if o == nil || o.AdditionalCosts == nil {
		var ret map[string]AdditionalCost
		return ret
	}
	return *o.AdditionalCosts
}

// GetAdditionalCostsOk returns a tuple with the AdditionalCosts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSessionV2) GetAdditionalCostsOk() (map[string]AdditionalCost, bool) {
	if o == nil || o.AdditionalCosts == nil {
		var ret map[string]AdditionalCost
		return ret, false
	}
	return *o.AdditionalCosts, true
}

// HasAdditionalCosts returns a boolean if a field has been set.
func (o *CustomerSessionV2) HasAdditionalCosts() bool {
	if o != nil && o.AdditionalCosts != nil {
		return true
	}

	return false
}

// SetAdditionalCosts gets a reference to the given map[string]AdditionalCost and assigns it to the AdditionalCosts field.
func (o *CustomerSessionV2) SetAdditionalCosts(v map[string]AdditionalCost) {
	o.AdditionalCosts = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *CustomerSessionV2) GetIdentifiers() []string {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret
	}
	return *o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSessionV2) GetIdentifiersOk() ([]string, bool) {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret, false
	}
	return *o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *CustomerSessionV2) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.
func (o *CustomerSessionV2) SetIdentifiers(v []string) {
	o.Identifiers = &v
}

// GetAttributes returns the Attributes field value
func (o *CustomerSessionV2) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// SetAttributes sets field value
func (o *CustomerSessionV2) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetFirstSession returns the FirstSession field value
func (o *CustomerSessionV2) GetFirstSession() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FirstSession
}

// SetFirstSession sets field value
func (o *CustomerSessionV2) SetFirstSession(v bool) {
	o.FirstSession = v
}

// GetTotal returns the Total field value
func (o *CustomerSessionV2) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// SetTotal sets field value
func (o *CustomerSessionV2) SetTotal(v float32) {
	o.Total = v
}

// GetCartItemTotal returns the CartItemTotal field value
func (o *CustomerSessionV2) GetCartItemTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CartItemTotal
}

// SetCartItemTotal sets field value
func (o *CustomerSessionV2) SetCartItemTotal(v float32) {
	o.CartItemTotal = v
}

// GetAdditionalCostTotal returns the AdditionalCostTotal field value
func (o *CustomerSessionV2) GetAdditionalCostTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AdditionalCostTotal
}

// SetAdditionalCostTotal sets field value
func (o *CustomerSessionV2) SetAdditionalCostTotal(v float32) {
	o.AdditionalCostTotal = v
}

type NullableCustomerSessionV2 struct {
	Value        CustomerSessionV2
	ExplicitNull bool
}

func (v NullableCustomerSessionV2) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomerSessionV2) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
