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

// IntegrationState Contains all state that might interest application integration plugins. This is the response type returned by all of the Integration API operations. 
type IntegrationState struct {
	Session CustomerSession `json:"session"`
	Profile CustomerProfile `json:"profile"`
	Event Event `json:"event"`
	Loyalty *Loyalty `json:"loyalty,omitempty"`
	Coupon *Coupon `json:"coupon,omitempty"`
}

// GetSession returns the Session field value
func (o *IntegrationState) GetSession() CustomerSession {
	if o == nil {
		var ret CustomerSession
		return ret
	}

	return o.Session
}

// SetSession sets field value
func (o *IntegrationState) SetSession(v CustomerSession) {
	o.Session = v
}

// GetProfile returns the Profile field value
func (o *IntegrationState) GetProfile() CustomerProfile {
	if o == nil {
		var ret CustomerProfile
		return ret
	}

	return o.Profile
}

// SetProfile sets field value
func (o *IntegrationState) SetProfile(v CustomerProfile) {
	o.Profile = v
}

// GetEvent returns the Event field value
func (o *IntegrationState) GetEvent() Event {
	if o == nil {
		var ret Event
		return ret
	}

	return o.Event
}

// SetEvent sets field value
func (o *IntegrationState) SetEvent(v Event) {
	o.Event = v
}

// GetLoyalty returns the Loyalty field value if set, zero value otherwise.
func (o *IntegrationState) GetLoyalty() Loyalty {
	if o == nil || o.Loyalty == nil {
		var ret Loyalty
		return ret
	}
	return *o.Loyalty
}

// GetLoyaltyOk returns a tuple with the Loyalty field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationState) GetLoyaltyOk() (Loyalty, bool) {
	if o == nil || o.Loyalty == nil {
		var ret Loyalty
		return ret, false
	}
	return *o.Loyalty, true
}

// HasLoyalty returns a boolean if a field has been set.
func (o *IntegrationState) HasLoyalty() bool {
	if o != nil && o.Loyalty != nil {
		return true
	}

	return false
}

// SetLoyalty gets a reference to the given Loyalty and assigns it to the Loyalty field.
func (o *IntegrationState) SetLoyalty(v Loyalty) {
	o.Loyalty = &v
}

// GetCoupon returns the Coupon field value if set, zero value otherwise.
func (o *IntegrationState) GetCoupon() Coupon {
	if o == nil || o.Coupon == nil {
		var ret Coupon
		return ret
	}
	return *o.Coupon
}

// GetCouponOk returns a tuple with the Coupon field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationState) GetCouponOk() (Coupon, bool) {
	if o == nil || o.Coupon == nil {
		var ret Coupon
		return ret, false
	}
	return *o.Coupon, true
}

// HasCoupon returns a boolean if a field has been set.
func (o *IntegrationState) HasCoupon() bool {
	if o != nil && o.Coupon != nil {
		return true
	}

	return false
}

// SetCoupon gets a reference to the given Coupon and assigns it to the Coupon field.
func (o *IntegrationState) SetCoupon(v Coupon) {
	o.Coupon = &v
}

type NullableIntegrationState struct {
	Value IntegrationState
	ExplicitNull bool
}

func (v NullableIntegrationState) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableIntegrationState) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
