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
	"time"
)

// InventoryCoupon struct for InventoryCoupon
type InventoryCoupon struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the campaign that owns this entity.
	CampaignId int32 `json:"campaignId"`
	// The actual coupon code.
	Value string `json:"value"`
	// The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.
	UsageLimit int32 `json:"usageLimit"`
	// The amount of discounts that can be given with this coupon code.
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of times this coupon has been successfully used.
	UsageCounter int32 `json:"usageCounter"`
	// The amount of discounts given on rules redeeming this coupon. Only usable if a coupon discount budget was set for this coupon.
	DiscountCounter *float32 `json:"discountCounter,omitempty"`
	// The remaining discount this coupon can give.
	DiscountRemainder *float32 `json:"discountRemainder,omitempty"`
	// Arbitrary properties associated with this item
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The integration ID of the referring customer (if any) for whom this coupon was created as an effect.
	ReferralId *int32 `json:"referralId,omitempty"`
	// The Integration ID of the customer that is allowed to redeem this coupon.
	RecipientIntegrationId *string `json:"recipientIntegrationId,omitempty"`
	// The ID of the Import which created this coupon.
	ImportId *int32 `json:"importId,omitempty"`
	// This value controls what reservations mean to a coupon. If set to true the coupon reservation is used to mark it as a favourite, if set to false the coupon reservation is used as a requirement of usage. This value defaults to true if not specified.
	Reservation *bool `json:"reservation,omitempty"`
	// The id of the batch the coupon belongs to.
	BatchId *string `json:"batchId,omitempty"`
	// The number of times the coupon was redeemed by the profile.
	ProfileRedemptionCount int32 `json:"profileRedemptionCount"`
	// Can be either active, used, expired, or pending. active: reserved coupons that are neither pending nor used nor expired, and have a non-exhausted limit counter. used: coupons that are not pending, and have reached their redemption limit or were redeemed by the profile before expiration. expired: all non-pending, non-active, non-used coupons that were not redeemed by the profile. pending: coupons that have a start date in the future.
	State string `json:"state"`
}

// GetId returns the Id field value
func (o *InventoryCoupon) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *InventoryCoupon) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *InventoryCoupon) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *InventoryCoupon) SetCreated(v time.Time) {
	o.Created = v
}

// GetCampaignId returns the CampaignId field value
func (o *InventoryCoupon) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *InventoryCoupon) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetValue returns the Value field value
func (o *InventoryCoupon) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *InventoryCoupon) SetValue(v string) {
	o.Value = v
}

// GetUsageLimit returns the UsageLimit field value
func (o *InventoryCoupon) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// SetUsageLimit sets field value
func (o *InventoryCoupon) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetDiscountLimit returns the DiscountLimit field value if set, zero value otherwise.
func (o *InventoryCoupon) GetDiscountLimit() float32 {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret
	}
	return *o.DiscountLimit
}

// GetDiscountLimitOk returns a tuple with the DiscountLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetDiscountLimitOk() (float32, bool) {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountLimit, true
}

// HasDiscountLimit returns a boolean if a field has been set.
func (o *InventoryCoupon) HasDiscountLimit() bool {
	if o != nil && o.DiscountLimit != nil {
		return true
	}

	return false
}

// SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.
func (o *InventoryCoupon) SetDiscountLimit(v float32) {
	o.DiscountLimit = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InventoryCoupon) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InventoryCoupon) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *InventoryCoupon) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *InventoryCoupon) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *InventoryCoupon) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *InventoryCoupon) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetUsageCounter returns the UsageCounter field value
func (o *InventoryCoupon) GetUsageCounter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageCounter
}

// SetUsageCounter sets field value
func (o *InventoryCoupon) SetUsageCounter(v int32) {
	o.UsageCounter = v
}

// GetDiscountCounter returns the DiscountCounter field value if set, zero value otherwise.
func (o *InventoryCoupon) GetDiscountCounter() float32 {
	if o == nil || o.DiscountCounter == nil {
		var ret float32
		return ret
	}
	return *o.DiscountCounter
}

// GetDiscountCounterOk returns a tuple with the DiscountCounter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetDiscountCounterOk() (float32, bool) {
	if o == nil || o.DiscountCounter == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountCounter, true
}

// HasDiscountCounter returns a boolean if a field has been set.
func (o *InventoryCoupon) HasDiscountCounter() bool {
	if o != nil && o.DiscountCounter != nil {
		return true
	}

	return false
}

// SetDiscountCounter gets a reference to the given float32 and assigns it to the DiscountCounter field.
func (o *InventoryCoupon) SetDiscountCounter(v float32) {
	o.DiscountCounter = &v
}

// GetDiscountRemainder returns the DiscountRemainder field value if set, zero value otherwise.
func (o *InventoryCoupon) GetDiscountRemainder() float32 {
	if o == nil || o.DiscountRemainder == nil {
		var ret float32
		return ret
	}
	return *o.DiscountRemainder
}

// GetDiscountRemainderOk returns a tuple with the DiscountRemainder field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetDiscountRemainderOk() (float32, bool) {
	if o == nil || o.DiscountRemainder == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountRemainder, true
}

// HasDiscountRemainder returns a boolean if a field has been set.
func (o *InventoryCoupon) HasDiscountRemainder() bool {
	if o != nil && o.DiscountRemainder != nil {
		return true
	}

	return false
}

// SetDiscountRemainder gets a reference to the given float32 and assigns it to the DiscountRemainder field.
func (o *InventoryCoupon) SetDiscountRemainder(v float32) {
	o.DiscountRemainder = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InventoryCoupon) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InventoryCoupon) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *InventoryCoupon) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetReferralId returns the ReferralId field value if set, zero value otherwise.
func (o *InventoryCoupon) GetReferralId() int32 {
	if o == nil || o.ReferralId == nil {
		var ret int32
		return ret
	}
	return *o.ReferralId
}

// GetReferralIdOk returns a tuple with the ReferralId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetReferralIdOk() (int32, bool) {
	if o == nil || o.ReferralId == nil {
		var ret int32
		return ret, false
	}
	return *o.ReferralId, true
}

// HasReferralId returns a boolean if a field has been set.
func (o *InventoryCoupon) HasReferralId() bool {
	if o != nil && o.ReferralId != nil {
		return true
	}

	return false
}

// SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.
func (o *InventoryCoupon) SetReferralId(v int32) {
	o.ReferralId = &v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value if set, zero value otherwise.
func (o *InventoryCoupon) GetRecipientIntegrationId() string {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.RecipientIntegrationId
}

// GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetRecipientIntegrationIdOk() (string, bool) {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.RecipientIntegrationId, true
}

// HasRecipientIntegrationId returns a boolean if a field has been set.
func (o *InventoryCoupon) HasRecipientIntegrationId() bool {
	if o != nil && o.RecipientIntegrationId != nil {
		return true
	}

	return false
}

// SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.
func (o *InventoryCoupon) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = &v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *InventoryCoupon) GetImportId() int32 {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetImportIdOk() (int32, bool) {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret, false
	}
	return *o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *InventoryCoupon) HasImportId() bool {
	if o != nil && o.ImportId != nil {
		return true
	}

	return false
}

// SetImportId gets a reference to the given int32 and assigns it to the ImportId field.
func (o *InventoryCoupon) SetImportId(v int32) {
	o.ImportId = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *InventoryCoupon) GetReservation() bool {
	if o == nil || o.Reservation == nil {
		var ret bool
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetReservationOk() (bool, bool) {
	if o == nil || o.Reservation == nil {
		var ret bool
		return ret, false
	}
	return *o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *InventoryCoupon) HasReservation() bool {
	if o != nil && o.Reservation != nil {
		return true
	}

	return false
}

// SetReservation gets a reference to the given bool and assigns it to the Reservation field.
func (o *InventoryCoupon) SetReservation(v bool) {
	o.Reservation = &v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *InventoryCoupon) GetBatchId() string {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetBatchIdOk() (string, bool) {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret, false
	}
	return *o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *InventoryCoupon) HasBatchId() bool {
	if o != nil && o.BatchId != nil {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *InventoryCoupon) SetBatchId(v string) {
	o.BatchId = &v
}

// GetProfileRedemptionCount returns the ProfileRedemptionCount field value
func (o *InventoryCoupon) GetProfileRedemptionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProfileRedemptionCount
}

// SetProfileRedemptionCount sets field value
func (o *InventoryCoupon) SetProfileRedemptionCount(v int32) {
	o.ProfileRedemptionCount = v
}

// GetState returns the State field value
func (o *InventoryCoupon) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// SetState sets field value
func (o *InventoryCoupon) SetState(v string) {
	o.State = v
}

type NullableInventoryCoupon struct {
	Value        InventoryCoupon
	ExplicitNull bool
}

func (v NullableInventoryCoupon) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInventoryCoupon) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
