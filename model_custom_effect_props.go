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

// CustomEffectProps Effect containing custom payload.
type CustomEffectProps struct {
	// The ID of the custom effect that was triggered.
	EffectId int32 `json:"effectId"`
	// The type of the custom effect.
	Name string `json:"name"`
	// The index of the item in the cart item list to which the custom effect is applied.
	CartItemPosition *float32 `json:"cartItemPosition,omitempty"`
	// When cart item flattening is enabled, the sub position indicates to which item unit the custom effect is applied, for cart items with quantity > 1.
	CartItemSubPosition *float32 `json:"cartItemSubPosition,omitempty"`
	// The position of the bundle in a list of item bundles created from the same bundle definition.
	BundleIndex *int32 `json:"bundleIndex,omitempty"`
	// The name of the bundle definition.
	BundleName *string `json:"bundleName,omitempty"`
	// The JSON payload of the custom effect.
	Payload map[string]interface{} `json:"payload"`
}

// GetEffectId returns the EffectId field value
func (o *CustomEffectProps) GetEffectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EffectId
}

// SetEffectId sets field value
func (o *CustomEffectProps) SetEffectId(v int32) {
	o.EffectId = v
}

// GetName returns the Name field value
func (o *CustomEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CustomEffectProps) SetName(v string) {
	o.Name = v
}

// GetCartItemPosition returns the CartItemPosition field value if set, zero value otherwise.
func (o *CustomEffectProps) GetCartItemPosition() float32 {
	if o == nil || o.CartItemPosition == nil {
		var ret float32
		return ret
	}
	return *o.CartItemPosition
}

// GetCartItemPositionOk returns a tuple with the CartItemPosition field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffectProps) GetCartItemPositionOk() (float32, bool) {
	if o == nil || o.CartItemPosition == nil {
		var ret float32
		return ret, false
	}
	return *o.CartItemPosition, true
}

// HasCartItemPosition returns a boolean if a field has been set.
func (o *CustomEffectProps) HasCartItemPosition() bool {
	if o != nil && o.CartItemPosition != nil {
		return true
	}

	return false
}

// SetCartItemPosition gets a reference to the given float32 and assigns it to the CartItemPosition field.
func (o *CustomEffectProps) SetCartItemPosition(v float32) {
	o.CartItemPosition = &v
}

// GetCartItemSubPosition returns the CartItemSubPosition field value if set, zero value otherwise.
func (o *CustomEffectProps) GetCartItemSubPosition() float32 {
	if o == nil || o.CartItemSubPosition == nil {
		var ret float32
		return ret
	}
	return *o.CartItemSubPosition
}

// GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffectProps) GetCartItemSubPositionOk() (float32, bool) {
	if o == nil || o.CartItemSubPosition == nil {
		var ret float32
		return ret, false
	}
	return *o.CartItemSubPosition, true
}

// HasCartItemSubPosition returns a boolean if a field has been set.
func (o *CustomEffectProps) HasCartItemSubPosition() bool {
	if o != nil && o.CartItemSubPosition != nil {
		return true
	}

	return false
}

// SetCartItemSubPosition gets a reference to the given float32 and assigns it to the CartItemSubPosition field.
func (o *CustomEffectProps) SetCartItemSubPosition(v float32) {
	o.CartItemSubPosition = &v
}

// GetBundleIndex returns the BundleIndex field value if set, zero value otherwise.
func (o *CustomEffectProps) GetBundleIndex() int32 {
	if o == nil || o.BundleIndex == nil {
		var ret int32
		return ret
	}
	return *o.BundleIndex
}

// GetBundleIndexOk returns a tuple with the BundleIndex field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffectProps) GetBundleIndexOk() (int32, bool) {
	if o == nil || o.BundleIndex == nil {
		var ret int32
		return ret, false
	}
	return *o.BundleIndex, true
}

// HasBundleIndex returns a boolean if a field has been set.
func (o *CustomEffectProps) HasBundleIndex() bool {
	if o != nil && o.BundleIndex != nil {
		return true
	}

	return false
}

// SetBundleIndex gets a reference to the given int32 and assigns it to the BundleIndex field.
func (o *CustomEffectProps) SetBundleIndex(v int32) {
	o.BundleIndex = &v
}

// GetBundleName returns the BundleName field value if set, zero value otherwise.
func (o *CustomEffectProps) GetBundleName() string {
	if o == nil || o.BundleName == nil {
		var ret string
		return ret
	}
	return *o.BundleName
}

// GetBundleNameOk returns a tuple with the BundleName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffectProps) GetBundleNameOk() (string, bool) {
	if o == nil || o.BundleName == nil {
		var ret string
		return ret, false
	}
	return *o.BundleName, true
}

// HasBundleName returns a boolean if a field has been set.
func (o *CustomEffectProps) HasBundleName() bool {
	if o != nil && o.BundleName != nil {
		return true
	}

	return false
}

// SetBundleName gets a reference to the given string and assigns it to the BundleName field.
func (o *CustomEffectProps) SetBundleName(v string) {
	o.BundleName = &v
}

// GetPayload returns the Payload field value
func (o *CustomEffectProps) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// SetPayload sets field value
func (o *CustomEffectProps) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

type NullableCustomEffectProps struct {
	Value        CustomEffectProps
	ExplicitNull bool
}

func (v NullableCustomEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
