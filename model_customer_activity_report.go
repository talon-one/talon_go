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
	"time"
)

// CustomerActivityReport A summary report of customer activity for a given time range.
type CustomerActivityReport struct {
	// The integration ID set by your integration layer.
	IntegrationId string `json:"integrationId"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The name for this customer profile.
	Name string `json:"name"`
	// The internal Talon.One ID of the customer.
	CustomerId int32 `json:"customerId"`
	// The last activity of the customer.
	LastActivity *time.Time `json:"lastActivity,omitempty"`
	// Number of coupon redemptions in all customer campaigns.
	CouponRedemptions int32 `json:"couponRedemptions"`
	// Number of coupon use attempts in all customer campaigns.
	CouponUseAttempts int32 `json:"couponUseAttempts"`
	// Number of failed coupon use attempts in all customer campaigns.
	CouponFailedAttempts int32 `json:"couponFailedAttempts"`
	// Number of accrued discounts in all customer campaigns.
	AccruedDiscounts float32 `json:"accruedDiscounts"`
	// Amount of accrued revenue in all customer campaigns.
	AccruedRevenue float32 `json:"accruedRevenue"`
	// Number of orders in all customer campaigns.
	TotalOrders int32 `json:"totalOrders"`
	// Number of orders without coupon used in all customer campaigns.
	TotalOrdersNoCoupon int32 `json:"totalOrdersNoCoupon"`
	// The name of the campaign this customer belongs to.
	CampaignName string `json:"campaignName"`
}

// GetIntegrationId returns the IntegrationId field value
func (o *CustomerActivityReport) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *CustomerActivityReport) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetCreated returns the Created field value
func (o *CustomerActivityReport) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *CustomerActivityReport) SetCreated(v time.Time) {
	o.Created = v
}

// GetName returns the Name field value
func (o *CustomerActivityReport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CustomerActivityReport) SetName(v string) {
	o.Name = v
}

// GetCustomerId returns the CustomerId field value
func (o *CustomerActivityReport) GetCustomerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CustomerId
}

// SetCustomerId sets field value
func (o *CustomerActivityReport) SetCustomerId(v int32) {
	o.CustomerId = v
}

// GetLastActivity returns the LastActivity field value if set, zero value otherwise.
func (o *CustomerActivityReport) GetLastActivity() time.Time {
	if o == nil || o.LastActivity == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActivity
}

// GetLastActivityOk returns a tuple with the LastActivity field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomerActivityReport) GetLastActivityOk() (time.Time, bool) {
	if o == nil || o.LastActivity == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastActivity, true
}

// HasLastActivity returns a boolean if a field has been set.
func (o *CustomerActivityReport) HasLastActivity() bool {
	if o != nil && o.LastActivity != nil {
		return true
	}

	return false
}

// SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.
func (o *CustomerActivityReport) SetLastActivity(v time.Time) {
	o.LastActivity = &v
}

// GetCouponRedemptions returns the CouponRedemptions field value
func (o *CustomerActivityReport) GetCouponRedemptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponRedemptions
}

// SetCouponRedemptions sets field value
func (o *CustomerActivityReport) SetCouponRedemptions(v int32) {
	o.CouponRedemptions = v
}

// GetCouponUseAttempts returns the CouponUseAttempts field value
func (o *CustomerActivityReport) GetCouponUseAttempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponUseAttempts
}

// SetCouponUseAttempts sets field value
func (o *CustomerActivityReport) SetCouponUseAttempts(v int32) {
	o.CouponUseAttempts = v
}

// GetCouponFailedAttempts returns the CouponFailedAttempts field value
func (o *CustomerActivityReport) GetCouponFailedAttempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CouponFailedAttempts
}

// SetCouponFailedAttempts sets field value
func (o *CustomerActivityReport) SetCouponFailedAttempts(v int32) {
	o.CouponFailedAttempts = v
}

// GetAccruedDiscounts returns the AccruedDiscounts field value
func (o *CustomerActivityReport) GetAccruedDiscounts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccruedDiscounts
}

// SetAccruedDiscounts sets field value
func (o *CustomerActivityReport) SetAccruedDiscounts(v float32) {
	o.AccruedDiscounts = v
}

// GetAccruedRevenue returns the AccruedRevenue field value
func (o *CustomerActivityReport) GetAccruedRevenue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccruedRevenue
}

// SetAccruedRevenue sets field value
func (o *CustomerActivityReport) SetAccruedRevenue(v float32) {
	o.AccruedRevenue = v
}

// GetTotalOrders returns the TotalOrders field value
func (o *CustomerActivityReport) GetTotalOrders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalOrders
}

// SetTotalOrders sets field value
func (o *CustomerActivityReport) SetTotalOrders(v int32) {
	o.TotalOrders = v
}

// GetTotalOrdersNoCoupon returns the TotalOrdersNoCoupon field value
func (o *CustomerActivityReport) GetTotalOrdersNoCoupon() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalOrdersNoCoupon
}

// SetTotalOrdersNoCoupon sets field value
func (o *CustomerActivityReport) SetTotalOrdersNoCoupon(v int32) {
	o.TotalOrdersNoCoupon = v
}

// GetCampaignName returns the CampaignName field value
func (o *CustomerActivityReport) GetCampaignName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignName
}

// SetCampaignName sets field value
func (o *CustomerActivityReport) SetCampaignName(v string) {
	o.CampaignName = v
}

type NullableCustomerActivityReport struct {
	Value        CustomerActivityReport
	ExplicitNull bool
}

func (v NullableCustomerActivityReport) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomerActivityReport) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
