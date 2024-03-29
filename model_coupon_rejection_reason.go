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

// CouponRejectionReason Holds a reference to the campaign, the coupon and the reason for which that coupon was rejected. Should only be present when there is a 'rejectCoupon' effect.
type CouponRejectionReason struct {
	CampaignId int32  `json:"campaignId"`
	CouponId   int32  `json:"couponId"`
	Reason     string `json:"reason"`
}

// GetCampaignId returns the CampaignId field value
func (o *CouponRejectionReason) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *CouponRejectionReason) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetCouponId returns the CouponId field value
func (o *CouponRejectionReason) GetCouponId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponId
}

// SetCouponId sets field value
func (o *CouponRejectionReason) SetCouponId(v int32) {
	o.CouponId = v
}

// GetReason returns the Reason field value
func (o *CouponRejectionReason) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// SetReason sets field value
func (o *CouponRejectionReason) SetReason(v string) {
	o.Reason = v
}

type NullableCouponRejectionReason struct {
	Value        CouponRejectionReason
	ExplicitNull bool
}

func (v NullableCouponRejectionReason) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCouponRejectionReason) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
