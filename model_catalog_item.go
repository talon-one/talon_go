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

// CatalogItem
type CatalogItem struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The stock keeping unit of the item.
	Sku string `json:"sku"`
	// Price of the item.
	Price *float32 `json:"price,omitempty"`
	// The ID of the catalog the item belongs to.
	Catalogid int32 `json:"catalogid"`
	// The version of the catalog item.
	Version    int32            `json:"version"`
	Attributes *[]ItemAttribute `json:"attributes,omitempty"`
}

// GetId returns the Id field value
func (o *CatalogItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *CatalogItem) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *CatalogItem) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *CatalogItem) SetCreated(v time.Time) {
	o.Created = v
}

// GetSku returns the Sku field value
func (o *CatalogItem) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// SetSku sets field value
func (o *CatalogItem) SetSku(v string) {
	o.Sku = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CatalogItem) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatalogItem) GetPriceOk() (float32, bool) {
	if o == nil || o.Price == nil {
		var ret float32
		return ret, false
	}
	return *o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CatalogItem) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *CatalogItem) SetPrice(v float32) {
	o.Price = &v
}

// GetCatalogid returns the Catalogid field value
func (o *CatalogItem) GetCatalogid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Catalogid
}

// SetCatalogid sets field value
func (o *CatalogItem) SetCatalogid(v int32) {
	o.Catalogid = v
}

// GetVersion returns the Version field value
func (o *CatalogItem) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// SetVersion sets field value
func (o *CatalogItem) SetVersion(v int32) {
	o.Version = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CatalogItem) GetAttributes() []ItemAttribute {
	if o == nil || o.Attributes == nil {
		var ret []ItemAttribute
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CatalogItem) GetAttributesOk() ([]ItemAttribute, bool) {
	if o == nil || o.Attributes == nil {
		var ret []ItemAttribute
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CatalogItem) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []ItemAttribute and assigns it to the Attributes field.
func (o *CatalogItem) SetAttributes(v []ItemAttribute) {
	o.Attributes = &v
}

type NullableCatalogItem struct {
	Value        CatalogItem
	ExplicitNull bool
}

func (v NullableCatalogItem) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCatalogItem) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}