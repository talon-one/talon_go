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

// CouponDeletionFilters struct for CouponDeletionFilters
type CouponDeletionFilters struct {
	// Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally.
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
	// Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally.
	CreatedAfter *time.Time `json:"createdAfter,omitempty"`
	// Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally.
	StartsAfter *time.Time `json:"startsAfter,omitempty"`
	// Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally.
	StartsBefore *time.Time `json:"startsBefore,omitempty"`
	// - `expired`: Matches coupons in which the expiration date is set and in the past. - `validNow`: Matches coupons in which the start date is null or in the past and the expiration date is null or in the future. - `validFuture`: Matches coupons in which the start date is set and in the future.
	Valid *string `json:"valid,omitempty"`
	// - `true`: only coupons where `usageCounter < usageLimit` will be returned. - `false`: only coupons where `usageCounter >= usageLimit` will be returned. - This field cannot be used in conjunction with the `usable` query parameter.
	Usable *bool `json:"usable,omitempty"`
	// - `true`: only coupons where `usageCounter > 0` will be returned. - `false`: only coupons where `usageCounter = 0` will be returned.  **Note:** This field cannot be used in conjunction with the `usable` query parameter.
	Redeemed *bool `json:"redeemed,omitempty"`
	// Filter results by match with a profile id specified in the coupon's `RecipientIntegrationId` field.
	RecipientIntegrationId *string `json:"recipientIntegrationId,omitempty"`
	// Filter results to an exact case-insensitive matching against the coupon code
	ExactMatch *bool `json:"exactMatch,omitempty"`
	// Filter results by the coupon code
	Value *string `json:"value,omitempty"`
	// Filter results by batches of coupons
	BatchId *string `json:"batchId,omitempty"`
	// Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code.
	ReferralId *int32 `json:"referralId,omitempty"`
	// Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally.
	ExpiresAfter *time.Time `json:"expiresAfter,omitempty"`
	// Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally.
	ExpiresBefore *time.Time `json:"expiresBefore,omitempty"`
}

// GetCreatedBefore returns the CreatedBefore field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetCreatedBefore() time.Time {
	if o == nil || o.CreatedBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedBefore
}

// GetCreatedBeforeOk returns a tuple with the CreatedBefore field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetCreatedBeforeOk() (time.Time, bool) {
	if o == nil || o.CreatedBefore == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedBefore, true
}

// HasCreatedBefore returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasCreatedBefore() bool {
	if o != nil && o.CreatedBefore != nil {
		return true
	}

	return false
}

// SetCreatedBefore gets a reference to the given time.Time and assigns it to the CreatedBefore field.
func (o *CouponDeletionFilters) SetCreatedBefore(v time.Time) {
	o.CreatedBefore = &v
}

// GetCreatedAfter returns the CreatedAfter field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetCreatedAfter() time.Time {
	if o == nil || o.CreatedAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAfter
}

// GetCreatedAfterOk returns a tuple with the CreatedAfter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetCreatedAfterOk() (time.Time, bool) {
	if o == nil || o.CreatedAfter == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedAfter, true
}

// HasCreatedAfter returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasCreatedAfter() bool {
	if o != nil && o.CreatedAfter != nil {
		return true
	}

	return false
}

// SetCreatedAfter gets a reference to the given time.Time and assigns it to the CreatedAfter field.
func (o *CouponDeletionFilters) SetCreatedAfter(v time.Time) {
	o.CreatedAfter = &v
}

// GetStartsAfter returns the StartsAfter field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetStartsAfter() time.Time {
	if o == nil || o.StartsAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.StartsAfter
}

// GetStartsAfterOk returns a tuple with the StartsAfter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetStartsAfterOk() (time.Time, bool) {
	if o == nil || o.StartsAfter == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartsAfter, true
}

// HasStartsAfter returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasStartsAfter() bool {
	if o != nil && o.StartsAfter != nil {
		return true
	}

	return false
}

// SetStartsAfter gets a reference to the given time.Time and assigns it to the StartsAfter field.
func (o *CouponDeletionFilters) SetStartsAfter(v time.Time) {
	o.StartsAfter = &v
}

// GetStartsBefore returns the StartsBefore field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetStartsBefore() time.Time {
	if o == nil || o.StartsBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.StartsBefore
}

// GetStartsBeforeOk returns a tuple with the StartsBefore field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetStartsBeforeOk() (time.Time, bool) {
	if o == nil || o.StartsBefore == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartsBefore, true
}

// HasStartsBefore returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasStartsBefore() bool {
	if o != nil && o.StartsBefore != nil {
		return true
	}

	return false
}

// SetStartsBefore gets a reference to the given time.Time and assigns it to the StartsBefore field.
func (o *CouponDeletionFilters) SetStartsBefore(v time.Time) {
	o.StartsBefore = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetValid() string {
	if o == nil || o.Valid == nil {
		var ret string
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetValidOk() (string, bool) {
	if o == nil || o.Valid == nil {
		var ret string
		return ret, false
	}
	return *o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given string and assigns it to the Valid field.
func (o *CouponDeletionFilters) SetValid(v string) {
	o.Valid = &v
}

// GetUsable returns the Usable field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetUsable() bool {
	if o == nil || o.Usable == nil {
		var ret bool
		return ret
	}
	return *o.Usable
}

// GetUsableOk returns a tuple with the Usable field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetUsableOk() (bool, bool) {
	if o == nil || o.Usable == nil {
		var ret bool
		return ret, false
	}
	return *o.Usable, true
}

// HasUsable returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasUsable() bool {
	if o != nil && o.Usable != nil {
		return true
	}

	return false
}

// SetUsable gets a reference to the given bool and assigns it to the Usable field.
func (o *CouponDeletionFilters) SetUsable(v bool) {
	o.Usable = &v
}

// GetRedeemed returns the Redeemed field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetRedeemed() bool {
	if o == nil || o.Redeemed == nil {
		var ret bool
		return ret
	}
	return *o.Redeemed
}

// GetRedeemedOk returns a tuple with the Redeemed field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetRedeemedOk() (bool, bool) {
	if o == nil || o.Redeemed == nil {
		var ret bool
		return ret, false
	}
	return *o.Redeemed, true
}

// HasRedeemed returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasRedeemed() bool {
	if o != nil && o.Redeemed != nil {
		return true
	}

	return false
}

// SetRedeemed gets a reference to the given bool and assigns it to the Redeemed field.
func (o *CouponDeletionFilters) SetRedeemed(v bool) {
	o.Redeemed = &v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetRecipientIntegrationId() string {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.RecipientIntegrationId
}

// GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetRecipientIntegrationIdOk() (string, bool) {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.RecipientIntegrationId, true
}

// HasRecipientIntegrationId returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasRecipientIntegrationId() bool {
	if o != nil && o.RecipientIntegrationId != nil {
		return true
	}

	return false
}

// SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.
func (o *CouponDeletionFilters) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = &v
}

// GetExactMatch returns the ExactMatch field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetExactMatch() bool {
	if o == nil || o.ExactMatch == nil {
		var ret bool
		return ret
	}
	return *o.ExactMatch
}

// GetExactMatchOk returns a tuple with the ExactMatch field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetExactMatchOk() (bool, bool) {
	if o == nil || o.ExactMatch == nil {
		var ret bool
		return ret, false
	}
	return *o.ExactMatch, true
}

// HasExactMatch returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasExactMatch() bool {
	if o != nil && o.ExactMatch != nil {
		return true
	}

	return false
}

// SetExactMatch gets a reference to the given bool and assigns it to the ExactMatch field.
func (o *CouponDeletionFilters) SetExactMatch(v bool) {
	o.ExactMatch = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetValueOk() (string, bool) {
	if o == nil || o.Value == nil {
		var ret string
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CouponDeletionFilters) SetValue(v string) {
	o.Value = &v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetBatchId() string {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetBatchIdOk() (string, bool) {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret, false
	}
	return *o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasBatchId() bool {
	if o != nil && o.BatchId != nil {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *CouponDeletionFilters) SetBatchId(v string) {
	o.BatchId = &v
}

// GetReferralId returns the ReferralId field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetReferralId() int32 {
	if o == nil || o.ReferralId == nil {
		var ret int32
		return ret
	}
	return *o.ReferralId
}

// GetReferralIdOk returns a tuple with the ReferralId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetReferralIdOk() (int32, bool) {
	if o == nil || o.ReferralId == nil {
		var ret int32
		return ret, false
	}
	return *o.ReferralId, true
}

// HasReferralId returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasReferralId() bool {
	if o != nil && o.ReferralId != nil {
		return true
	}

	return false
}

// SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.
func (o *CouponDeletionFilters) SetReferralId(v int32) {
	o.ReferralId = &v
}

// GetExpiresAfter returns the ExpiresAfter field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetExpiresAfter() time.Time {
	if o == nil || o.ExpiresAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAfter
}

// GetExpiresAfterOk returns a tuple with the ExpiresAfter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetExpiresAfterOk() (time.Time, bool) {
	if o == nil || o.ExpiresAfter == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiresAfter, true
}

// HasExpiresAfter returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasExpiresAfter() bool {
	if o != nil && o.ExpiresAfter != nil {
		return true
	}

	return false
}

// SetExpiresAfter gets a reference to the given time.Time and assigns it to the ExpiresAfter field.
func (o *CouponDeletionFilters) SetExpiresAfter(v time.Time) {
	o.ExpiresAfter = &v
}

// GetExpiresBefore returns the ExpiresBefore field value if set, zero value otherwise.
func (o *CouponDeletionFilters) GetExpiresBefore() time.Time {
	if o == nil || o.ExpiresBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresBefore
}

// GetExpiresBeforeOk returns a tuple with the ExpiresBefore field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponDeletionFilters) GetExpiresBeforeOk() (time.Time, bool) {
	if o == nil || o.ExpiresBefore == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiresBefore, true
}

// HasExpiresBefore returns a boolean if a field has been set.
func (o *CouponDeletionFilters) HasExpiresBefore() bool {
	if o != nil && o.ExpiresBefore != nil {
		return true
	}

	return false
}

// SetExpiresBefore gets a reference to the given time.Time and assigns it to the ExpiresBefore field.
func (o *CouponDeletionFilters) SetExpiresBefore(v time.Time) {
	o.ExpiresBefore = &v
}

type NullableCouponDeletionFilters struct {
	Value        CouponDeletionFilters
	ExplicitNull bool
}

func (v NullableCouponDeletionFilters) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCouponDeletionFilters) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
