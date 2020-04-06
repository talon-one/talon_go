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

// CartItem struct for CartItem
type CartItem struct {
	Name string `json:"name"`
	Sku string `json:"sku"`
	Quantity int32 `json:"quantity"`
	Price float32 `json:"price"`
	Category *string `json:"category,omitempty"`
	// Weight of item in mm
	Weight *float32 `json:"weight,omitempty"`
	// Height of item in mm
	Height *float32 `json:"height,omitempty"`
	// Width of item in mm
	Width *float32 `json:"width,omitempty"`
	// Length of item in mm
	Length *float32 `json:"length,omitempty"`
	// Position of the Cart Item in the Cart (calculated internally)
	Position *float32 `json:"position,omitempty"`
	// Arbitrary properties associated with this item
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	Adjustment *CartItemAdjustment `json:"adjustment,omitempty"`
}

// GetName returns the Name field value
func (o *CartItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CartItem) SetName(v string) {
	o.Name = v
}

// GetSku returns the Sku field value
func (o *CartItem) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// SetSku sets field value
func (o *CartItem) SetSku(v string) {
	o.Sku = v
}

// GetQuantity returns the Quantity field value
func (o *CartItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// SetQuantity sets field value
func (o *CartItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetPrice returns the Price field value
func (o *CartItem) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// SetPrice sets field value
func (o *CartItem) SetPrice(v float32) {
	o.Price = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CartItem) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetCategoryOk() (string, bool) {
	if o == nil || o.Category == nil {
		var ret string
		return ret, false
	}
	return *o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CartItem) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CartItem) SetCategory(v string) {
	o.Category = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *CartItem) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetWeightOk() (float32, bool) {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret, false
	}
	return *o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *CartItem) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *CartItem) SetWeight(v float32) {
	o.Weight = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CartItem) GetHeight() float32 {
	if o == nil || o.Height == nil {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetHeightOk() (float32, bool) {
	if o == nil || o.Height == nil {
		var ret float32
		return ret, false
	}
	return *o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CartItem) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *CartItem) SetHeight(v float32) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CartItem) GetWidth() float32 {
	if o == nil || o.Width == nil {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetWidthOk() (float32, bool) {
	if o == nil || o.Width == nil {
		var ret float32
		return ret, false
	}
	return *o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CartItem) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *CartItem) SetWidth(v float32) {
	o.Width = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *CartItem) GetLength() float32 {
	if o == nil || o.Length == nil {
		var ret float32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetLengthOk() (float32, bool) {
	if o == nil || o.Length == nil {
		var ret float32
		return ret, false
	}
	return *o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *CartItem) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given float32 and assigns it to the Length field.
func (o *CartItem) SetLength(v float32) {
	o.Length = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *CartItem) GetPosition() float32 {
	if o == nil || o.Position == nil {
		var ret float32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetPositionOk() (float32, bool) {
	if o == nil || o.Position == nil {
		var ret float32
		return ret, false
	}
	return *o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *CartItem) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given float32 and assigns it to the Position field.
func (o *CartItem) SetPosition(v float32) {
	o.Position = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CartItem) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CartItem) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *CartItem) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetAdjustment returns the Adjustment field value if set, zero value otherwise.
func (o *CartItem) GetAdjustment() CartItemAdjustment {
	if o == nil || o.Adjustment == nil {
		var ret CartItemAdjustment
		return ret
	}
	return *o.Adjustment
}

// GetAdjustmentOk returns a tuple with the Adjustment field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CartItem) GetAdjustmentOk() (CartItemAdjustment, bool) {
	if o == nil || o.Adjustment == nil {
		var ret CartItemAdjustment
		return ret, false
	}
	return *o.Adjustment, true
}

// HasAdjustment returns a boolean if a field has been set.
func (o *CartItem) HasAdjustment() bool {
	if o != nil && o.Adjustment != nil {
		return true
	}

	return false
}

// SetAdjustment gets a reference to the given CartItemAdjustment and assigns it to the Adjustment field.
func (o *CartItem) SetAdjustment(v CartItemAdjustment) {
	o.Adjustment = &v
}

type NullableCartItem struct {
	Value CartItem
	ExplicitNull bool
}

func (v NullableCartItem) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCartItem) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
