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

// CustomerAnalytics
type CustomerAnalytics struct {
	// Total accepted coupons for this customer
	AcceptedCoupons int32 `json:"acceptedCoupons"`
	// Total created coupons for this customer
	CreatedCoupons int32 `json:"createdCoupons"`
	// Total free items given to this customer
	FreeItems int32 `json:"freeItems"`
	// Total orders made by this customer
	TotalOrders int32 `json:"totalOrders"`
	// Total orders made by this customer that had a discount
	TotalDiscountedOrders int32 `json:"totalDiscountedOrders"`
	// Total Revenue across all closed sessions
	TotalRevenue float32 `json:"totalRevenue"`
	// The sum of discounts that were given across all closed sessions
	TotalDiscounts float32 `json:"totalDiscounts"`
}

// GetAcceptedCoupons returns the AcceptedCoupons field value
func (o *CustomerAnalytics) GetAcceptedCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcceptedCoupons
}

// SetAcceptedCoupons sets field value
func (o *CustomerAnalytics) SetAcceptedCoupons(v int32) {
	o.AcceptedCoupons = v
}

// GetCreatedCoupons returns the CreatedCoupons field value
func (o *CustomerAnalytics) GetCreatedCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedCoupons
}

// SetCreatedCoupons sets field value
func (o *CustomerAnalytics) SetCreatedCoupons(v int32) {
	o.CreatedCoupons = v
}

// GetFreeItems returns the FreeItems field value
func (o *CustomerAnalytics) GetFreeItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FreeItems
}

// SetFreeItems sets field value
func (o *CustomerAnalytics) SetFreeItems(v int32) {
	o.FreeItems = v
}

// GetTotalOrders returns the TotalOrders field value
func (o *CustomerAnalytics) GetTotalOrders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalOrders
}

// SetTotalOrders sets field value
func (o *CustomerAnalytics) SetTotalOrders(v int32) {
	o.TotalOrders = v
}

// GetTotalDiscountedOrders returns the TotalDiscountedOrders field value
func (o *CustomerAnalytics) GetTotalDiscountedOrders() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalDiscountedOrders
}

// SetTotalDiscountedOrders sets field value
func (o *CustomerAnalytics) SetTotalDiscountedOrders(v int32) {
	o.TotalDiscountedOrders = v
}

// GetTotalRevenue returns the TotalRevenue field value
func (o *CustomerAnalytics) GetTotalRevenue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalRevenue
}

// SetTotalRevenue sets field value
func (o *CustomerAnalytics) SetTotalRevenue(v float32) {
	o.TotalRevenue = v
}

// GetTotalDiscounts returns the TotalDiscounts field value
func (o *CustomerAnalytics) GetTotalDiscounts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalDiscounts
}

// SetTotalDiscounts sets field value
func (o *CustomerAnalytics) SetTotalDiscounts(v float32) {
	o.TotalDiscounts = v
}

type NullableCustomerAnalytics struct {
	Value        CustomerAnalytics
	ExplicitNull bool
}

func (v NullableCustomerAnalytics) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomerAnalytics) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
