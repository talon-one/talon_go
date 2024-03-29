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

// SetDiscountEffectProps The properties specific to the \"setDiscount\" effect. This gets triggered whenever a validated rule contained a \"set discount\" effect. This is a discount that should be applied on the scope of defined with it.
type SetDiscountEffectProps struct {
	// The name / description of this discount
	Name string `json:"name"`
	// The total monetary value of the discount.
	Value float32 `json:"value"`
	// The scope which the discount was applied on, can be one of (cartItems,additionalCosts,sessionTotal).
	Scope *string `json:"scope,omitempty"`
	// The original value of the discount.
	DesiredValue *float32 `json:"desiredValue,omitempty"`
}

// GetName returns the Name field value
func (o *SetDiscountEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *SetDiscountEffectProps) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *SetDiscountEffectProps) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *SetDiscountEffectProps) SetValue(v float32) {
	o.Value = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SetDiscountEffectProps) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SetDiscountEffectProps) GetScopeOk() (string, bool) {
	if o == nil || o.Scope == nil {
		var ret string
		return ret, false
	}
	return *o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SetDiscountEffectProps) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *SetDiscountEffectProps) SetScope(v string) {
	o.Scope = &v
}

// GetDesiredValue returns the DesiredValue field value if set, zero value otherwise.
func (o *SetDiscountEffectProps) GetDesiredValue() float32 {
	if o == nil || o.DesiredValue == nil {
		var ret float32
		return ret
	}
	return *o.DesiredValue
}

// GetDesiredValueOk returns a tuple with the DesiredValue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SetDiscountEffectProps) GetDesiredValueOk() (float32, bool) {
	if o == nil || o.DesiredValue == nil {
		var ret float32
		return ret, false
	}
	return *o.DesiredValue, true
}

// HasDesiredValue returns a boolean if a field has been set.
func (o *SetDiscountEffectProps) HasDesiredValue() bool {
	if o != nil && o.DesiredValue != nil {
		return true
	}

	return false
}

// SetDesiredValue gets a reference to the given float32 and assigns it to the DesiredValue field.
func (o *SetDiscountEffectProps) SetDesiredValue(v float32) {
	o.DesiredValue = &v
}

type NullableSetDiscountEffectProps struct {
	Value        SetDiscountEffectProps
	ExplicitNull bool
}

func (v NullableSetDiscountEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSetDiscountEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
