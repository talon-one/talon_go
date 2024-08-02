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

// ApplicationSession struct for ApplicationSession
type ApplicationSession struct {
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Arbitrary properties associated with this item.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Serialized JSON representation.
	CartItems []CartItem `json:"cartItems"`
	// Any coupon code entered.
	Coupon string `json:"coupon"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// **API V1 only.** A map of labeled discount values, in the same currency as the session.  If you are using the V2 endpoints, refer to the `totalDiscounts` property instead. 
	Discounts map[string]float32 `json:"discounts"`
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The integration ID set by your integration layer.
	IntegrationId string `json:"integrationId"`
	// The globally unique Talon.One ID of the customer that created this entity.
	ProfileId *int32 `json:"profileId,omitempty"`
	// Integration ID of the customer for the session.
	Profileintegrationid *string `json:"profileintegrationid,omitempty"`
	// Any referral code entered.
	Referral string `json:"referral"`
	// Indicates the current state of the session. Sessions can be created as `open` or `closed`. The state transitions are:  1. `open` → `closed` 2. `open` → `cancelled` 3. `closed` → `cancelled` or `partially_returned` 4. `partially_returned` → `cancelled`  For more information, see [Customer session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions). 
	State string `json:"state"`
	// The integration ID of the store. You choose this ID when you create a store.
	StoreIntegrationId *string `json:"storeIntegrationId,omitempty"`
	// The total sum of the session before any discounts applied.
	Total float32 `json:"total"`
	// The total sum of the discounts applied to this session.
	TotalDiscounts float32 `json:"totalDiscounts"`
}

// GetApplicationId returns the ApplicationId field value
func (o *ApplicationSession) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *ApplicationSession) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ApplicationSession) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSession) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ApplicationSession) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ApplicationSession) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetCartItems returns the CartItems field value
func (o *ApplicationSession) GetCartItems() []CartItem {
	if o == nil {
		var ret []CartItem
		return ret
	}

	return o.CartItems
}

// SetCartItems sets field value
func (o *ApplicationSession) SetCartItems(v []CartItem) {
	o.CartItems = v
}

// GetCoupon returns the Coupon field value
func (o *ApplicationSession) GetCoupon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coupon
}

// SetCoupon sets field value
func (o *ApplicationSession) SetCoupon(v string) {
	o.Coupon = v
}

// GetCreated returns the Created field value
func (o *ApplicationSession) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *ApplicationSession) SetCreated(v time.Time) {
	o.Created = v
}

// GetDiscounts returns the Discounts field value
func (o *ApplicationSession) GetDiscounts() map[string]float32 {
	if o == nil {
		var ret map[string]float32
		return ret
	}

	return o.Discounts
}

// SetDiscounts sets field value
func (o *ApplicationSession) SetDiscounts(v map[string]float32) {
	o.Discounts = v
}

// GetId returns the Id field value
func (o *ApplicationSession) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *ApplicationSession) SetId(v int32) {
	o.Id = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *ApplicationSession) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *ApplicationSession) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ApplicationSession) GetProfileId() int32 {
	if o == nil || o.ProfileId == nil {
		var ret int32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSession) GetProfileIdOk() (int32, bool) {
	if o == nil || o.ProfileId == nil {
		var ret int32
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ApplicationSession) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.
func (o *ApplicationSession) SetProfileId(v int32) {
	o.ProfileId = &v
}

// GetProfileintegrationid returns the Profileintegrationid field value if set, zero value otherwise.
func (o *ApplicationSession) GetProfileintegrationid() string {
	if o == nil || o.Profileintegrationid == nil {
		var ret string
		return ret
	}
	return *o.Profileintegrationid
}

// GetProfileintegrationidOk returns a tuple with the Profileintegrationid field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSession) GetProfileintegrationidOk() (string, bool) {
	if o == nil || o.Profileintegrationid == nil {
		var ret string
		return ret, false
	}
	return *o.Profileintegrationid, true
}

// HasProfileintegrationid returns a boolean if a field has been set.
func (o *ApplicationSession) HasProfileintegrationid() bool {
	if o != nil && o.Profileintegrationid != nil {
		return true
	}

	return false
}

// SetProfileintegrationid gets a reference to the given string and assigns it to the Profileintegrationid field.
func (o *ApplicationSession) SetProfileintegrationid(v string) {
	o.Profileintegrationid = &v
}

// GetReferral returns the Referral field value
func (o *ApplicationSession) GetReferral() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Referral
}

// SetReferral sets field value
func (o *ApplicationSession) SetReferral(v string) {
	o.Referral = v
}

// GetState returns the State field value
func (o *ApplicationSession) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// SetState sets field value
func (o *ApplicationSession) SetState(v string) {
	o.State = v
}

// GetStoreIntegrationId returns the StoreIntegrationId field value if set, zero value otherwise.
func (o *ApplicationSession) GetStoreIntegrationId() string {
	if o == nil || o.StoreIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.StoreIntegrationId
}

// GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSession) GetStoreIntegrationIdOk() (string, bool) {
	if o == nil || o.StoreIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.StoreIntegrationId, true
}

// HasStoreIntegrationId returns a boolean if a field has been set.
func (o *ApplicationSession) HasStoreIntegrationId() bool {
	if o != nil && o.StoreIntegrationId != nil {
		return true
	}

	return false
}

// SetStoreIntegrationId gets a reference to the given string and assigns it to the StoreIntegrationId field.
func (o *ApplicationSession) SetStoreIntegrationId(v string) {
	o.StoreIntegrationId = &v
}

// GetTotal returns the Total field value
func (o *ApplicationSession) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// SetTotal sets field value
func (o *ApplicationSession) SetTotal(v float32) {
	o.Total = v
}

// GetTotalDiscounts returns the TotalDiscounts field value
func (o *ApplicationSession) GetTotalDiscounts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalDiscounts
}

// SetTotalDiscounts sets field value
func (o *ApplicationSession) SetTotalDiscounts(v float32) {
	o.TotalDiscounts = v
}

type NullableApplicationSession struct {
	Value ApplicationSession
	ExplicitNull bool
}

func (v NullableApplicationSession) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationSession) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
