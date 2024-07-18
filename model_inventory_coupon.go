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
	"time"
)

// InventoryCoupon struct for InventoryCoupon
type InventoryCoupon struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The ID of the campaign that owns this entity.
	CampaignId int32 `json:"campaignId"`
	// The coupon code.
	Value string `json:"value"`
	// The number of times the coupon code can be redeemed. `0` means unlimited redemptions but any campaign usage limits will still apply. 
	UsageLimit int32 `json:"usageLimit"`
	// The total discount value that the code can give. Typically used to represent a gift card value. 
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// The number of reservations that can be made with this coupon code. 
	ReservationLimit *int32 `json:"reservationLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiration date of the coupon. Coupon never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured. 
	Limits *[]LimitConfig `json:"limits,omitempty"`
	// The number of times the coupon has been successfully redeemed.
	UsageCounter int32 `json:"usageCounter"`
	// The amount of discounts given on rules redeeming this coupon. Only usable if a coupon discount budget was set for this coupon.
	DiscountCounter *float32 `json:"discountCounter,omitempty"`
	// The remaining discount this coupon can give.
	DiscountRemainder *float32 `json:"discountRemainder,omitempty"`
	// The number of times this coupon has been reserved.
	ReservationCounter *float32 `json:"reservationCounter,omitempty"`
	// Custom attributes associated with this coupon.
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	// The integration ID of the referring customer (if any) for whom this coupon was created as an effect.
	ReferralId *int32 `json:"referralId,omitempty"`
	// The Integration ID of the customer that is allowed to redeem this coupon.
	RecipientIntegrationId *string `json:"recipientIntegrationId,omitempty"`
	// The ID of the Import which created this coupon.
	ImportId *int32 `json:"importId,omitempty"`
	// Defines the reservation type: - `true`: The coupon can be reserved for multiple customers. - `false`: The coupon can be reserved only for one customer. It is a personal code. 
	Reservation *bool `json:"reservation,omitempty"`
	// The id of the batch the coupon belongs to.
	BatchId *string `json:"batchId,omitempty"`
	// An indication of whether the code can be redeemed only if it has been reserved first.
	IsReservationMandatory *bool `json:"isReservationMandatory,omitempty"`
	// An indication of whether the coupon is implicitly reserved for all customers.
	ImplicitlyReserved *bool `json:"implicitlyReserved,omitempty"`
	// The number of times the coupon was redeemed by the profile.
	ProfileRedemptionCount int32 `json:"profileRedemptionCount"`
	// Can be:  - `active`: The coupon can be used. It is a reserved coupon that is not pending, used, or expired, and it has a non-exhausted limit counter.    **Note:** This coupon state is returned for [scheduled campaigns](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-schedule), but the coupon cannot be used until the campaign is **running**. - `used`: The coupon has been redeemed and cannot be used again. It is not pending and has reached its redemption limit or was redeemed by the profile before expiration. - `expired`: The coupon was never redeemed, and it is now expired. It is non-pending, non-active, and non-used by the profile. - `pending`: The coupon will be usable in the future. - `disabled`: The coupon is part of a non-active campaign. 
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

// GetReservationLimit returns the ReservationLimit field value if set, zero value otherwise.
func (o *InventoryCoupon) GetReservationLimit() int32 {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret
	}
	return *o.ReservationLimit
}

// GetReservationLimitOk returns a tuple with the ReservationLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetReservationLimitOk() (int32, bool) {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.ReservationLimit, true
}

// HasReservationLimit returns a boolean if a field has been set.
func (o *InventoryCoupon) HasReservationLimit() bool {
	if o != nil && o.ReservationLimit != nil {
		return true
	}

	return false
}

// SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.
func (o *InventoryCoupon) SetReservationLimit(v int32) {
	o.ReservationLimit = &v
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

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *InventoryCoupon) GetLimits() []LimitConfig {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetLimitsOk() ([]LimitConfig, bool) {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret, false
	}
	return *o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *InventoryCoupon) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.
func (o *InventoryCoupon) SetLimits(v []LimitConfig) {
	o.Limits = &v
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

// GetReservationCounter returns the ReservationCounter field value if set, zero value otherwise.
func (o *InventoryCoupon) GetReservationCounter() float32 {
	if o == nil || o.ReservationCounter == nil {
		var ret float32
		return ret
	}
	return *o.ReservationCounter
}

// GetReservationCounterOk returns a tuple with the ReservationCounter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetReservationCounterOk() (float32, bool) {
	if o == nil || o.ReservationCounter == nil {
		var ret float32
		return ret, false
	}
	return *o.ReservationCounter, true
}

// HasReservationCounter returns a boolean if a field has been set.
func (o *InventoryCoupon) HasReservationCounter() bool {
	if o != nil && o.ReservationCounter != nil {
		return true
	}

	return false
}

// SetReservationCounter gets a reference to the given float32 and assigns it to the ReservationCounter field.
func (o *InventoryCoupon) SetReservationCounter(v float32) {
	o.ReservationCounter = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InventoryCoupon) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
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

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *InventoryCoupon) SetAttributes(v map[string]map[string]interface{}) {
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

// GetIsReservationMandatory returns the IsReservationMandatory field value if set, zero value otherwise.
func (o *InventoryCoupon) GetIsReservationMandatory() bool {
	if o == nil || o.IsReservationMandatory == nil {
		var ret bool
		return ret
	}
	return *o.IsReservationMandatory
}

// GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetIsReservationMandatoryOk() (bool, bool) {
	if o == nil || o.IsReservationMandatory == nil {
		var ret bool
		return ret, false
	}
	return *o.IsReservationMandatory, true
}

// HasIsReservationMandatory returns a boolean if a field has been set.
func (o *InventoryCoupon) HasIsReservationMandatory() bool {
	if o != nil && o.IsReservationMandatory != nil {
		return true
	}

	return false
}

// SetIsReservationMandatory gets a reference to the given bool and assigns it to the IsReservationMandatory field.
func (o *InventoryCoupon) SetIsReservationMandatory(v bool) {
	o.IsReservationMandatory = &v
}

// GetImplicitlyReserved returns the ImplicitlyReserved field value if set, zero value otherwise.
func (o *InventoryCoupon) GetImplicitlyReserved() bool {
	if o == nil || o.ImplicitlyReserved == nil {
		var ret bool
		return ret
	}
	return *o.ImplicitlyReserved
}

// GetImplicitlyReservedOk returns a tuple with the ImplicitlyReserved field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InventoryCoupon) GetImplicitlyReservedOk() (bool, bool) {
	if o == nil || o.ImplicitlyReserved == nil {
		var ret bool
		return ret, false
	}
	return *o.ImplicitlyReserved, true
}

// HasImplicitlyReserved returns a boolean if a field has been set.
func (o *InventoryCoupon) HasImplicitlyReserved() bool {
	if o != nil && o.ImplicitlyReserved != nil {
		return true
	}

	return false
}

// SetImplicitlyReserved gets a reference to the given bool and assigns it to the ImplicitlyReserved field.
func (o *InventoryCoupon) SetImplicitlyReserved(v bool) {
	o.ImplicitlyReserved = &v
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
	Value InventoryCoupon
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
