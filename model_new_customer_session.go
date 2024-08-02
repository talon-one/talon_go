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

// NewCustomerSession struct for NewCustomerSession
type NewCustomerSession struct {
	// A key-value map of the sessions attributes. The potentially valid attributes are configured in your accounts developer settings. 
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Serialized JSON representation.
	CartItems *[]CartItem `json:"cartItems,omitempty"`
	// Any coupon code entered.
	Coupon *string `json:"coupon,omitempty"`
	// Session custom identifiers that you can set limits on or use inside your rules.  For example, you can use IP addresses as identifiers to potentially identify devices and limit discounts abuse in case of customers creating multiple accounts. See the [tutorial](https://docs.talon.one/docs/dev/tutorials/using-identifiers). 
	Identifiers *[]string `json:"identifiers,omitempty"`
	// ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known `profileId`, we recommend you use a guest `profileId`. 
	ProfileId *string `json:"profileId,omitempty"`
	// Any referral code entered.
	Referral *string `json:"referral,omitempty"`
	// Indicates the current state of the session. Sessions can be created as `open` or `closed`. The state transitions are:  1. `open` → `closed` 2. `open` → `cancelled` 3. `closed` → `cancelled` or `partially_returned` 4. `partially_returned` → `cancelled`  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions). 
	State *string `json:"state,omitempty"`
	// The total sum of the cart in one session.
	Total *float32 `json:"total,omitempty"`
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

type NullableNewCustomerSession struct {
	Value NewCustomerSession
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
