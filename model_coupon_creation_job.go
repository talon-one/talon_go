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

// CouponCreationJob struct for CouponCreationJob
type CouponCreationJob struct {
	// Arbitrary properties associated with coupons.
	Attributes map[string]interface{} `json:"attributes"`
	CouponSettings *CodeGeneratorSettings `json:"couponSettings,omitempty"`
	// The number of new coupon codes to generate for the campaign.
	NumberOfCoupons int32 `json:"numberOfCoupons"`
}

// GetAttributes returns the Attributes field value
func (o *CouponCreationJob) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// SetAttributes sets field value
func (o *CouponCreationJob) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetCouponSettings returns the CouponSettings field value if set, zero value otherwise.
func (o *CouponCreationJob) GetCouponSettings() CodeGeneratorSettings {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.CouponSettings
}

// GetCouponSettingsOk returns a tuple with the CouponSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreationJob) GetCouponSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.CouponSettings, true
}

// HasCouponSettings returns a boolean if a field has been set.
func (o *CouponCreationJob) HasCouponSettings() bool {
	if o != nil && o.CouponSettings != nil {
		return true
	}

	return false
}

// SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.
func (o *CouponCreationJob) SetCouponSettings(v CodeGeneratorSettings) {
	o.CouponSettings = &v
}

// GetNumberOfCoupons returns the NumberOfCoupons field value
func (o *CouponCreationJob) GetNumberOfCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfCoupons
}

// SetNumberOfCoupons sets field value
func (o *CouponCreationJob) SetNumberOfCoupons(v int32) {
	o.NumberOfCoupons = v
}

type NullableCouponCreationJob struct {
	Value CouponCreationJob
	ExplicitNull bool
}

func (v NullableCouponCreationJob) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCouponCreationJob) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
