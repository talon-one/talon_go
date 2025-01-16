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

// ProductSkuUnitAnalytics struct for ProductSkuUnitAnalytics
type ProductSkuUnitAnalytics struct {
	// The start of the aggregation time frame in UTC.
	StartTime time.Time `json:"startTime"`
	// The end of the aggregation time frame in UTC.
	EndTime        time.Time                   `json:"endTime"`
	PurchasedUnits AnalyticsDataPointWithTrend `json:"purchasedUnits"`
	// The SKU linked to the analytics-level product.
	Sku string `json:"sku"`
}

// GetStartTime returns the StartTime field value
func (o *ProductSkuUnitAnalytics) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// SetStartTime sets field value
func (o *ProductSkuUnitAnalytics) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *ProductSkuUnitAnalytics) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// SetEndTime sets field value
func (o *ProductSkuUnitAnalytics) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetPurchasedUnits returns the PurchasedUnits field value
func (o *ProductSkuUnitAnalytics) GetPurchasedUnits() AnalyticsDataPointWithTrend {
	if o == nil {
		var ret AnalyticsDataPointWithTrend
		return ret
	}

	return o.PurchasedUnits
}

// SetPurchasedUnits sets field value
func (o *ProductSkuUnitAnalytics) SetPurchasedUnits(v AnalyticsDataPointWithTrend) {
	o.PurchasedUnits = v
}

// GetSku returns the Sku field value
func (o *ProductSkuUnitAnalytics) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// SetSku sets field value
func (o *ProductSkuUnitAnalytics) SetSku(v string) {
	o.Sku = v
}

type NullableProductSkuUnitAnalytics struct {
	Value        ProductSkuUnitAnalytics
	ExplicitNull bool
}

func (v NullableProductSkuUnitAnalytics) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableProductSkuUnitAnalytics) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
