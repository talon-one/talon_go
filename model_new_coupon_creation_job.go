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

// NewCouponCreationJob
type NewCouponCreationJob struct {
	// The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.
	UsageLimit int32 `json:"usageLimit"`
	// The amount of discounts that can be given with this coupon code.
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of new coupon codes to generate for the campaign. Must be between 20,001 and 5,000,000.
	NumberOfCoupons int32                  `json:"numberOfCoupons"`
	CouponSettings  *CodeGeneratorSettings `json:"couponSettings,omitempty"`
	// Arbitrary properties associated with coupons
	Attributes map[string]interface{} `json:"attributes"`
}

// GetUsageLimit returns the UsageLimit field value
func (o *NewCouponCreationJob) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// SetUsageLimit sets field value
func (o *NewCouponCreationJob) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetDiscountLimit returns the DiscountLimit field value if set, zero value otherwise.
func (o *NewCouponCreationJob) GetDiscountLimit() float32 {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret
	}
	return *o.DiscountLimit
}

// GetDiscountLimitOk returns a tuple with the DiscountLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponCreationJob) GetDiscountLimitOk() (float32, bool) {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountLimit, true
}

// HasDiscountLimit returns a boolean if a field has been set.
func (o *NewCouponCreationJob) HasDiscountLimit() bool {
	if o != nil && o.DiscountLimit != nil {
		return true
	}

	return false
}

// SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.
func (o *NewCouponCreationJob) SetDiscountLimit(v float32) {
	o.DiscountLimit = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *NewCouponCreationJob) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponCreationJob) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *NewCouponCreationJob) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *NewCouponCreationJob) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *NewCouponCreationJob) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponCreationJob) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *NewCouponCreationJob) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *NewCouponCreationJob) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetNumberOfCoupons returns the NumberOfCoupons field value
func (o *NewCouponCreationJob) GetNumberOfCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfCoupons
}

// SetNumberOfCoupons sets field value
func (o *NewCouponCreationJob) SetNumberOfCoupons(v int32) {
	o.NumberOfCoupons = v
}

// GetCouponSettings returns the CouponSettings field value if set, zero value otherwise.
func (o *NewCouponCreationJob) GetCouponSettings() CodeGeneratorSettings {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.CouponSettings
}

// GetCouponSettingsOk returns a tuple with the CouponSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponCreationJob) GetCouponSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.CouponSettings, true
}

// HasCouponSettings returns a boolean if a field has been set.
func (o *NewCouponCreationJob) HasCouponSettings() bool {
	if o != nil && o.CouponSettings != nil {
		return true
	}

	return false
}

// SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.
func (o *NewCouponCreationJob) SetCouponSettings(v CodeGeneratorSettings) {
	o.CouponSettings = &v
}

// GetAttributes returns the Attributes field value
func (o *NewCouponCreationJob) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// SetAttributes sets field value
func (o *NewCouponCreationJob) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

type NullableNewCouponCreationJob struct {
	Value        NewCouponCreationJob
	ExplicitNull bool
}

func (v NullableNewCouponCreationJob) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewCouponCreationJob) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}