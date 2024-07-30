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

// ReserveCouponEffectProps The properties specific to the \"reserveCoupon\" effect. This gets triggered whenever a validated rule contained a \"reserve coupon\" effect. This reserves the coupon currently on scope to the profile on scope.
type ReserveCouponEffectProps struct {
	// The value of the coupon currently on scope.
	CouponValue string `json:"couponValue"`
	// The ID of this customer profile in the third-party integration.
	ProfileIntegrationId string `json:"profileIntegrationId"`
	// Indicates whether this is a new coupon reservation or not.
	IsNewReservation bool `json:"isNewReservation"`
}

// GetCouponValue returns the CouponValue field value
func (o *ReserveCouponEffectProps) GetCouponValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CouponValue
}

// SetCouponValue sets field value
func (o *ReserveCouponEffectProps) SetCouponValue(v string) {
	o.CouponValue = v
}

// GetProfileIntegrationId returns the ProfileIntegrationId field value
func (o *ReserveCouponEffectProps) GetProfileIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileIntegrationId
}

// SetProfileIntegrationId sets field value
func (o *ReserveCouponEffectProps) SetProfileIntegrationId(v string) {
	o.ProfileIntegrationId = v
}

// GetIsNewReservation returns the IsNewReservation field value
func (o *ReserveCouponEffectProps) GetIsNewReservation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsNewReservation
}

// SetIsNewReservation sets field value
func (o *ReserveCouponEffectProps) SetIsNewReservation(v bool) {
	o.IsNewReservation = v
}

type NullableReserveCouponEffectProps struct {
	Value ReserveCouponEffectProps
	ExplicitNull bool
}

func (v NullableReserveCouponEffectProps) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReserveCouponEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
