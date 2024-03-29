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

// StrikethroughChangedItem The information of affected items.
type StrikethroughChangedItem struct {
	// The ID of the event that triggered the strikethrough labeling.
	Id int32 `json:"id"`
	// The ID of the catalog that the changed item belongs to.
	CatalogId int32 `json:"catalogId"`
	// The unique SKU of the changed item.
	Sku string `json:"sku"`
	// The version of the changed item.
	Version int32 `json:"version"`
	// The price of the changed item.
	Price float32 `json:"price"`
	// The evaluation time of the changed item.
	EvaluatedAt time.Time              `json:"evaluatedAt"`
	Effects     *[]StrikethroughEffect `json:"effects,omitempty"`
}

// GetId returns the Id field value
func (o *StrikethroughChangedItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *StrikethroughChangedItem) SetId(v int32) {
	o.Id = v
}

// GetCatalogId returns the CatalogId field value
func (o *StrikethroughChangedItem) GetCatalogId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CatalogId
}

// SetCatalogId sets field value
func (o *StrikethroughChangedItem) SetCatalogId(v int32) {
	o.CatalogId = v
}

// GetSku returns the Sku field value
func (o *StrikethroughChangedItem) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// SetSku sets field value
func (o *StrikethroughChangedItem) SetSku(v string) {
	o.Sku = v
}

// GetVersion returns the Version field value
func (o *StrikethroughChangedItem) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// SetVersion sets field value
func (o *StrikethroughChangedItem) SetVersion(v int32) {
	o.Version = v
}

// GetPrice returns the Price field value
func (o *StrikethroughChangedItem) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// SetPrice sets field value
func (o *StrikethroughChangedItem) SetPrice(v float32) {
	o.Price = v
}

// GetEvaluatedAt returns the EvaluatedAt field value
func (o *StrikethroughChangedItem) GetEvaluatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EvaluatedAt
}

// SetEvaluatedAt sets field value
func (o *StrikethroughChangedItem) SetEvaluatedAt(v time.Time) {
	o.EvaluatedAt = v
}

// GetEffects returns the Effects field value if set, zero value otherwise.
func (o *StrikethroughChangedItem) GetEffects() []StrikethroughEffect {
	if o == nil || o.Effects == nil {
		var ret []StrikethroughEffect
		return ret
	}
	return *o.Effects
}

// GetEffectsOk returns a tuple with the Effects field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *StrikethroughChangedItem) GetEffectsOk() ([]StrikethroughEffect, bool) {
	if o == nil || o.Effects == nil {
		var ret []StrikethroughEffect
		return ret, false
	}
	return *o.Effects, true
}

// HasEffects returns a boolean if a field has been set.
func (o *StrikethroughChangedItem) HasEffects() bool {
	if o != nil && o.Effects != nil {
		return true
	}

	return false
}

// SetEffects gets a reference to the given []StrikethroughEffect and assigns it to the Effects field.
func (o *StrikethroughChangedItem) SetEffects(v []StrikethroughEffect) {
	o.Effects = &v
}

type NullableStrikethroughChangedItem struct {
	Value        StrikethroughChangedItem
	ExplicitNull bool
}

func (v NullableStrikethroughChangedItem) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableStrikethroughChangedItem) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
