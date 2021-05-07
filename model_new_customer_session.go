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

// NewCustomerSession struct for NewCustomerSession
type NewCustomerSession struct {
	// ID of the customers profile as used within this Talon.One account. May be omitted or set to the empty string if the customer does not yet have a known profile ID.
	ProfileId *string `json:"profileId,omitempty"`
	// Any coupon code entered.
	Coupon *string `json:"coupon,omitempty"`
	// Any referral code entered.
	Referral *string `json:"referral,omitempty"`
	// Indicates the current state of the session. All sessions must start in the \"open\" state, after which valid transitions are...  1. open -> closed 2. open -> cancelled 3. closed -> cancelled
	State *string `json:"state,omitempty"`
	// Serialized JSON representation.
	CartItems *[]CartItem `json:"cartItems,omitempty"`
	// Identifiers for the customer, this can be used for limits on values such as device ID.
	Identifiers *[]string `json:"identifiers,omitempty"`
	// The total sum of the cart in one session.
	Total *float32 `json:"total,omitempty"`
	// A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *NewCustomerSession) GetProfileId() string {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSession) GetProfileIdOk() (string, bool) {
	if o == nil || o.ProfileId == nil {
		var ret string
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *NewCustomerSession) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *NewCustomerSession) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetCoupon returns the Coupon field value if set, zero value otherwise.
func (o *NewCustomerSession) GetCoupon() string {
	if o == nil || o.Coupon == nil {
		var ret string
		return ret
	}
	return *o.Coupon
}

// GetCouponOk returns a tuple with the Coupon field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSession) GetCouponOk() (string, bool) {
	if o == nil || o.Coupon == nil {
		var ret string
		return ret, false
	}
	return *o.Coupon, true
}

// HasCoupon returns a boolean if a field has been set.
func (o *NewCustomerSession) HasCoupon() bool {
	if o != nil && o.Coupon != nil {
		return true
	}

	return false
}

// SetCoupon gets a reference to the given string and assigns it to the Coupon field.
func (o *NewCustomerSession) SetCoupon(v string) {
	o.Coupon = &v
}

// GetReferral returns the Referral field value if set, zero value otherwise.
func (o *NewCustomerSession) GetReferral() string {
	if o == nil || o.Referral == nil {
		var ret string
		return ret
	}
	return *o.Referral
}

// GetReferralOk returns a tuple with the Referral field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSession) GetReferralOk() (string, bool) {
	if o == nil || o.Referral == nil {
		var ret string
		return ret, false
	}
	return *o.Referral, true
}

// HasReferral returns a boolean if a field has been set.
func (o *NewCustomerSession) HasReferral() bool {
	if o != nil && o.Referral != nil {
		return true
	}

	return false
}

// SetReferral gets a reference to the given string and assigns it to the Referral field.
func (o *NewCustomerSession) SetReferral(v string) {
	o.Referral = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NewCustomerSession) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSession) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NewCustomerSession) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NewCustomerSession) SetState(v string) {
	o.State = &v
}

// GetCartItems returns the CartItems field value if set, zero value otherwise.
func (o *NewCustomerSession) GetCartItems() []CartItem {
	if o == nil || o.CartItems == nil {
		var ret []CartItem
		return ret
	}
	return *o.CartItems
}

// GetCartItemsOk returns a tuple with the CartItems field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSession) GetCartItemsOk() ([]CartItem, bool) {
	if o == nil || o.CartItems == nil {
		var ret []CartItem
		return ret, false
	}
	return *o.CartItems, true
}

// HasCartItems returns a boolean if a field has been set.
func (o *NewCustomerSession) HasCartItems() bool {
	if o != nil && o.CartItems != nil {
		return true
	}

	return false
}

// SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.
func (o *NewCustomerSession) SetCartItems(v []CartItem) {
	o.CartItems = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *NewCustomerSession) GetIdentifiers() []string {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret
	}
	return *o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSession) GetIdentifiersOk() ([]string, bool) {
	if o == nil || o.Identifiers == nil {
		var ret []string
		return ret, false
	}
	return *o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *NewCustomerSession) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.
func (o *NewCustomerSession) SetIdentifiers(v []string) {
	o.Identifiers = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *NewCustomerSession) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSession) GetTotalOk() (float32, bool) {
	if o == nil || o.Total == nil {
		var ret float32
		return ret, false
	}
	return *o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *NewCustomerSession) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *NewCustomerSession) SetTotal(v float32) {
	o.Total = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewCustomerSession) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCustomerSession) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewCustomerSession) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewCustomerSession) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

type NullableNewCustomerSession struct {
	Value        NewCustomerSession
	ExplicitNull bool
}

func (v NullableNewCustomerSession) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewCustomerSession) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
