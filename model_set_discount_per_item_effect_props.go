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

// SetDiscountPerItemEffectProps The properties specific to the \"setDiscountPerItem\" effect. This gets triggered whenever a validated rule contained a \"set per item discount\" effect. This is a discount that should be applied on a specific item.
type SetDiscountPerItemEffectProps struct {
	// The name/description of this discount
	Name string `json:"name"`
	// The total monetary value of the discount
	Value float32 `json:"value"`
	// The index of the item in the cart items list on which this discount should be applied
	Position float32 `json:"position"`
}

// GetName returns the Name field value
func (o *SetDiscountPerItemEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *SetDiscountPerItemEffectProps) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *SetDiscountPerItemEffectProps) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *SetDiscountPerItemEffectProps) SetValue(v float32) {
	o.Value = v
}

// GetPosition returns the Position field value
func (o *SetDiscountPerItemEffectProps) GetPosition() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Position
}

// SetPosition sets field value
func (o *SetDiscountPerItemEffectProps) SetPosition(v float32) {
	o.Position = v
}

type NullableSetDiscountPerItemEffectProps struct {
	Value SetDiscountPerItemEffectProps
	ExplicitNull bool
}

func (v NullableSetDiscountPerItemEffectProps) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSetDiscountPerItemEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
