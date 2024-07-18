/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// SetDiscountPerAdditionalCostPerItemEffectProps The properties specific to the \"setDiscountPerAdditionalCostPerItem\" effect. This gets triggered whenever a validated rule contained a \"set discount per additional cost per item\" effect. This is a discount that should be applied on a specific additional cost in a specific item.
type SetDiscountPerAdditionalCostPerItemEffectProps struct {
	// The name / description of this discount
	Name string `json:"name"`
	// The ID of the additional cost.
	AdditionalCostId int32 `json:"additionalCostId"`
	// The total monetary value of the discount.
	Value float32 `json:"value"`
	// The index of the item in the cart item list containing the additional cost to be discounted.
	Position float32 `json:"position"`
	// For cart items with `quantity` > 1, the sub position indicates which item the discount applies to. 
	SubPosition *float32 `json:"subPosition,omitempty"`
	// The name of the additional cost.
	AdditionalCost string `json:"additionalCost"`
	// Only with [partial discounts enabled](https://docs.talon.one/docs/product/campaigns/campaign-evaluation/#partial-discounts). Represents the monetary value of the discount to be applied to additional discount without considering budget limitations. 
	DesiredValue *float32 `json:"desiredValue,omitempty"`
}

// GetName returns the Name field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetName(v string) {
	o.Name = v
}

// GetAdditionalCostId returns the AdditionalCostId field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCostId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AdditionalCostId
}

// SetAdditionalCostId sets field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetAdditionalCostId(v int32) {
	o.AdditionalCostId = v
}

// GetValue returns the Value field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetValue(v float32) {
	o.Value = v
}

// GetPosition returns the Position field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetPosition() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Position
}

// SetPosition sets field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetPosition(v float32) {
	o.Position = v
}

// GetSubPosition returns the SubPosition field value if set, zero value otherwise.
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetSubPosition() float32 {
	if o == nil || o.SubPosition == nil {
		var ret float32
		return ret
	}
	return *o.SubPosition
}

// GetSubPositionOk returns a tuple with the SubPosition field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetSubPositionOk() (float32, bool) {
	if o == nil || o.SubPosition == nil {
		var ret float32
		return ret, false
	}
	return *o.SubPosition, true
}

// HasSubPosition returns a boolean if a field has been set.
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasSubPosition() bool {
	if o != nil && o.SubPosition != nil {
		return true
	}

	return false
}

// SetSubPosition gets a reference to the given float32 and assigns it to the SubPosition field.
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetSubPosition(v float32) {
	o.SubPosition = &v
}

// GetAdditionalCost returns the AdditionalCost field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetAdditionalCost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdditionalCost
}

// SetAdditionalCost sets field value
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetAdditionalCost(v string) {
	o.AdditionalCost = v
}

// GetDesiredValue returns the DesiredValue field value if set, zero value otherwise.
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetDesiredValue() float32 {
	if o == nil || o.DesiredValue == nil {
		var ret float32
		return ret
	}
	return *o.DesiredValue
}

// GetDesiredValueOk returns a tuple with the DesiredValue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) GetDesiredValueOk() (float32, bool) {
	if o == nil || o.DesiredValue == nil {
		var ret float32
		return ret, false
	}
	return *o.DesiredValue, true
}

// HasDesiredValue returns a boolean if a field has been set.
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) HasDesiredValue() bool {
	if o != nil && o.DesiredValue != nil {
		return true
	}

	return false
}

// SetDesiredValue gets a reference to the given float32 and assigns it to the DesiredValue field.
func (o *SetDiscountPerAdditionalCostPerItemEffectProps) SetDesiredValue(v float32) {
	o.DesiredValue = &v
}

type NullableSetDiscountPerAdditionalCostPerItemEffectProps struct {
	Value SetDiscountPerAdditionalCostPerItemEffectProps
	ExplicitNull bool
}

func (v NullableSetDiscountPerAdditionalCostPerItemEffectProps) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSetDiscountPerAdditionalCostPerItemEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
