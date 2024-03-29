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

// CouponCreatedEffectProps The properties specific to the \"couponCreated\" effect. This gets triggered whenever a validated rule contained a \"create coupon\" effect, and a coupon was created for a customer. See \"createdCoupons\" on the response for all details of this coupon.
type CouponCreatedEffectProps struct {
	// The coupon code that was created.
	Value string `json:"value"`
	// The integration identifier of the customer for whom this coupon was created.
	ProfileId string `json:"profileId"`
}

// GetValue returns the Value field value
func (o *CouponCreatedEffectProps) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *CouponCreatedEffectProps) SetValue(v string) {
	o.Value = v
}

// GetProfileId returns the ProfileId field value
func (o *CouponCreatedEffectProps) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// SetProfileId sets field value
func (o *CouponCreatedEffectProps) SetProfileId(v string) {
	o.ProfileId = v
}

type NullableCouponCreatedEffectProps struct {
	Value        CouponCreatedEffectProps
	ExplicitNull bool
}

func (v NullableCouponCreatedEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCouponCreatedEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
