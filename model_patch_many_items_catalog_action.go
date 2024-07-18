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

// PatchManyItemsCatalogAction The specific properties of the \"PATCH_MANY\" catalog sync action.
type PatchManyItemsCatalogAction struct {
	// Price of the item.
	Price *float32 `json:"price,omitempty"`
	// The list of filters used to select the items to patch, joined by `AND`.  **Note:** Every item in the catalog will be modified if there are no filters. 
	Filters *[]CatalogActionFilter `json:"filters,omitempty"`
	// The attributes of the items to patch.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PatchManyItemsCatalogAction) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PatchManyItemsCatalogAction) GetPriceOk() (float32, bool) {
	if o == nil || o.Price == nil {
		var ret float32
		return ret, false
	}
	return *o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PatchManyItemsCatalogAction) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *PatchManyItemsCatalogAction) SetPrice(v float32) {
	o.Price = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *PatchManyItemsCatalogAction) GetFilters() []CatalogActionFilter {
	if o == nil || o.Filters == nil {
		var ret []CatalogActionFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PatchManyItemsCatalogAction) GetFiltersOk() ([]CatalogActionFilter, bool) {
	if o == nil || o.Filters == nil {
		var ret []CatalogActionFilter
		return ret, false
	}
	return *o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *PatchManyItemsCatalogAction) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []CatalogActionFilter and assigns it to the Filters field.
func (o *PatchManyItemsCatalogAction) SetFilters(v []CatalogActionFilter) {
	o.Filters = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PatchManyItemsCatalogAction) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PatchManyItemsCatalogAction) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PatchManyItemsCatalogAction) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *PatchManyItemsCatalogAction) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

type NullablePatchManyItemsCatalogAction struct {
	Value PatchManyItemsCatalogAction
	ExplicitNull bool
}

func (v NullablePatchManyItemsCatalogAction) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullablePatchManyItemsCatalogAction) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
