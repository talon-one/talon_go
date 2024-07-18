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

// CatalogActionPayload struct for CatalogActionPayload
type CatalogActionPayload struct {
	// The unique SKU of the item to remove.
	Sku string `json:"sku"`
	// Price of the item.
	Price *float32 `json:"price,omitempty"`
	// The attributes of the items to patch.
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	Product *Product `json:"product,omitempty"`
	// Indicates whether to replace the attributes of the item if the same SKU exists.  **Note**: When set to `true`:   - If you do not provide a new `price` value, the existing `price` value is retained.   - If you do not provide a new `product` value, the `product` value is set to `null`. 
	ReplaceIfExists *bool `json:"replaceIfExists,omitempty"`
	// Indicates whether to create an item if the SKU does not exist.
	CreateIfNotExists *bool `json:"createIfNotExists,omitempty"`
	// The list of filters used to select the items to patch, joined by `AND`.  **Note:** Every item in the catalog will be removed if there are no filters. 
	Filters *[]CatalogActionFilter `json:"filters,omitempty"`
}

// GetSku returns the Sku field value
func (o *CatalogActionPayload) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// SetSku sets field value
func (o *CatalogActionPayload) SetSku(v string) {
	o.Sku = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CatalogActionPayload) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatalogActionPayload) GetPriceOk() (float32, bool) {
	if o == nil || o.Price == nil {
		var ret float32
		return ret, false
	}
	return *o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CatalogActionPayload) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *CatalogActionPayload) SetPrice(v float32) {
	o.Price = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CatalogActionPayload) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatalogActionPayload) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CatalogActionPayload) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *CatalogActionPayload) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *CatalogActionPayload) GetProduct() Product {
	if o == nil || o.Product == nil {
		var ret Product
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatalogActionPayload) GetProductOk() (Product, bool) {
	if o == nil || o.Product == nil {
		var ret Product
		return ret, false
	}
	return *o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *CatalogActionPayload) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given Product and assigns it to the Product field.
func (o *CatalogActionPayload) SetProduct(v Product) {
	o.Product = &v
}

// GetReplaceIfExists returns the ReplaceIfExists field value if set, zero value otherwise.
func (o *CatalogActionPayload) GetReplaceIfExists() bool {
	if o == nil || o.ReplaceIfExists == nil {
		var ret bool
		return ret
	}
	return *o.ReplaceIfExists
}

// GetReplaceIfExistsOk returns a tuple with the ReplaceIfExists field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatalogActionPayload) GetReplaceIfExistsOk() (bool, bool) {
	if o == nil || o.ReplaceIfExists == nil {
		var ret bool
		return ret, false
	}
	return *o.ReplaceIfExists, true
}

// HasReplaceIfExists returns a boolean if a field has been set.
func (o *CatalogActionPayload) HasReplaceIfExists() bool {
	if o != nil && o.ReplaceIfExists != nil {
		return true
	}

	return false
}

// SetReplaceIfExists gets a reference to the given bool and assigns it to the ReplaceIfExists field.
func (o *CatalogActionPayload) SetReplaceIfExists(v bool) {
	o.ReplaceIfExists = &v
}

// GetCreateIfNotExists returns the CreateIfNotExists field value if set, zero value otherwise.
func (o *CatalogActionPayload) GetCreateIfNotExists() bool {
	if o == nil || o.CreateIfNotExists == nil {
		var ret bool
		return ret
	}
	return *o.CreateIfNotExists
}

// GetCreateIfNotExistsOk returns a tuple with the CreateIfNotExists field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatalogActionPayload) GetCreateIfNotExistsOk() (bool, bool) {
	if o == nil || o.CreateIfNotExists == nil {
		var ret bool
		return ret, false
	}
	return *o.CreateIfNotExists, true
}

// HasCreateIfNotExists returns a boolean if a field has been set.
func (o *CatalogActionPayload) HasCreateIfNotExists() bool {
	if o != nil && o.CreateIfNotExists != nil {
		return true
	}

	return false
}

// SetCreateIfNotExists gets a reference to the given bool and assigns it to the CreateIfNotExists field.
func (o *CatalogActionPayload) SetCreateIfNotExists(v bool) {
	o.CreateIfNotExists = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CatalogActionPayload) GetFilters() []CatalogActionFilter {
	if o == nil || o.Filters == nil {
		var ret []CatalogActionFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatalogActionPayload) GetFiltersOk() ([]CatalogActionFilter, bool) {
	if o == nil || o.Filters == nil {
		var ret []CatalogActionFilter
		return ret, false
	}
	return *o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CatalogActionPayload) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []CatalogActionFilter and assigns it to the Filters field.
func (o *CatalogActionPayload) SetFilters(v []CatalogActionFilter) {
	o.Filters = &v
}

type NullableCatalogActionPayload struct {
	Value CatalogActionPayload
	ExplicitNull bool
}

func (v NullableCatalogActionPayload) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCatalogActionPayload) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}