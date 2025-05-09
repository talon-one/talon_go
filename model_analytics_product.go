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

// AnalyticsProduct struct for AnalyticsProduct
type AnalyticsProduct struct {
	// The ID of the product.
	Id int32 `json:"id"`
	// The name of the product.
	Name string `json:"name"`
	// The ID of the catalog. You can find the ID in the Campaign Manager in **Account** > **Tools** > **Cart item catalogs**.
	CatalogId int32                        `json:"catalogId"`
	UnitsSold *AnalyticsDataPointWithTrend `json:"unitsSold,omitempty"`
}

// GetId returns the Id field value
func (o *AnalyticsProduct) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *AnalyticsProduct) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AnalyticsProduct) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *AnalyticsProduct) SetName(v string) {
	o.Name = v
}

// GetCatalogId returns the CatalogId field value
func (o *AnalyticsProduct) GetCatalogId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CatalogId
}

// SetCatalogId sets field value
func (o *AnalyticsProduct) SetCatalogId(v int32) {
	o.CatalogId = v
}

// GetUnitsSold returns the UnitsSold field value if set, zero value otherwise.
func (o *AnalyticsProduct) GetUnitsSold() AnalyticsDataPointWithTrend {
	if o == nil || o.UnitsSold == nil {
		var ret AnalyticsDataPointWithTrend
		return ret
	}
	return *o.UnitsSold
}

// GetUnitsSoldOk returns a tuple with the UnitsSold field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsProduct) GetUnitsSoldOk() (AnalyticsDataPointWithTrend, bool) {
	if o == nil || o.UnitsSold == nil {
		var ret AnalyticsDataPointWithTrend
		return ret, false
	}
	return *o.UnitsSold, true
}

// HasUnitsSold returns a boolean if a field has been set.
func (o *AnalyticsProduct) HasUnitsSold() bool {
	if o != nil && o.UnitsSold != nil {
		return true
	}

	return false
}

// SetUnitsSold gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the UnitsSold field.
func (o *AnalyticsProduct) SetUnitsSold(v AnalyticsDataPointWithTrend) {
	o.UnitsSold = &v
}

type NullableAnalyticsProduct struct {
	Value        AnalyticsProduct
	ExplicitNull bool
}

func (v NullableAnalyticsProduct) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAnalyticsProduct) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
