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

// NewCouponsForMultipleRecipients struct for NewCouponsForMultipleRecipients
type NewCouponsForMultipleRecipients struct {
	// The number of times the coupon code can be redeemed. `0` means unlimited redemptions but any campaign usage limits will still apply.
	UsageLimit int32 `json:"usageLimit"`
	// The total discount value that the code can give. Typically used to represent a gift card value.
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// The number of reservations that can be made with this coupon code.
	ReservationLimit *int32 `json:"reservationLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiration date of the coupon. Coupon never expires if this is omitted.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// Arbitrary properties associated with this item.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The integration IDs for recipients.
	RecipientsIntegrationIds []string `json:"recipientsIntegrationIds"`
	// List of characters used to generate the random parts of a code. By default, the list of characters is equivalent to the `[A-Z, 0-9]` regular expression.
	ValidCharacters *[]string `json:"validCharacters,omitempty"`
	// The pattern used to generate coupon codes. The character `#` is a placeholder and is replaced by a random character from the `validCharacters` set.
	CouponPattern *string `json:"couponPattern,omitempty"`
}

// GetUsageLimit returns the UsageLimit field value
func (o *NewCouponsForMultipleRecipients) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// SetUsageLimit sets field value
func (o *NewCouponsForMultipleRecipients) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetDiscountLimit returns the DiscountLimit field value if set, zero value otherwise.
func (o *NewCouponsForMultipleRecipients) GetDiscountLimit() float32 {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret
	}
	return *o.DiscountLimit
}

// GetDiscountLimitOk returns a tuple with the DiscountLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponsForMultipleRecipients) GetDiscountLimitOk() (float32, bool) {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountLimit, true
}

// HasDiscountLimit returns a boolean if a field has been set.
func (o *NewCouponsForMultipleRecipients) HasDiscountLimit() bool {
	if o != nil && o.DiscountLimit != nil {
		return true
	}

	return false
}

// SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.
func (o *NewCouponsForMultipleRecipients) SetDiscountLimit(v float32) {
	o.DiscountLimit = &v
}

// GetReservationLimit returns the ReservationLimit field value if set, zero value otherwise.
func (o *NewCouponsForMultipleRecipients) GetReservationLimit() int32 {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret
	}
	return *o.ReservationLimit
}

// GetReservationLimitOk returns a tuple with the ReservationLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponsForMultipleRecipients) GetReservationLimitOk() (int32, bool) {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.ReservationLimit, true
}

// HasReservationLimit returns a boolean if a field has been set.
func (o *NewCouponsForMultipleRecipients) HasReservationLimit() bool {
	if o != nil && o.ReservationLimit != nil {
		return true
	}

	return false
}

// SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.
func (o *NewCouponsForMultipleRecipients) SetReservationLimit(v int32) {
	o.ReservationLimit = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *NewCouponsForMultipleRecipients) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponsForMultipleRecipients) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *NewCouponsForMultipleRecipients) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *NewCouponsForMultipleRecipients) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *NewCouponsForMultipleRecipients) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponsForMultipleRecipients) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *NewCouponsForMultipleRecipients) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *NewCouponsForMultipleRecipients) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewCouponsForMultipleRecipients) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponsForMultipleRecipients) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewCouponsForMultipleRecipients) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewCouponsForMultipleRecipients) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetRecipientsIntegrationIds returns the RecipientsIntegrationIds field value
func (o *NewCouponsForMultipleRecipients) GetRecipientsIntegrationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecipientsIntegrationIds
}

// SetRecipientsIntegrationIds sets field value
func (o *NewCouponsForMultipleRecipients) SetRecipientsIntegrationIds(v []string) {
	o.RecipientsIntegrationIds = v
}

// GetValidCharacters returns the ValidCharacters field value if set, zero value otherwise.
func (o *NewCouponsForMultipleRecipients) GetValidCharacters() []string {
	if o == nil || o.ValidCharacters == nil {
		var ret []string
		return ret
	}
	return *o.ValidCharacters
}

// GetValidCharactersOk returns a tuple with the ValidCharacters field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponsForMultipleRecipients) GetValidCharactersOk() ([]string, bool) {
	if o == nil || o.ValidCharacters == nil {
		var ret []string
		return ret, false
	}
	return *o.ValidCharacters, true
}

// HasValidCharacters returns a boolean if a field has been set.
func (o *NewCouponsForMultipleRecipients) HasValidCharacters() bool {
	if o != nil && o.ValidCharacters != nil {
		return true
	}

	return false
}

// SetValidCharacters gets a reference to the given []string and assigns it to the ValidCharacters field.
func (o *NewCouponsForMultipleRecipients) SetValidCharacters(v []string) {
	o.ValidCharacters = &v
}

// GetCouponPattern returns the CouponPattern field value if set, zero value otherwise.
func (o *NewCouponsForMultipleRecipients) GetCouponPattern() string {
	if o == nil || o.CouponPattern == nil {
		var ret string
		return ret
	}
	return *o.CouponPattern
}

// GetCouponPatternOk returns a tuple with the CouponPattern field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCouponsForMultipleRecipients) GetCouponPatternOk() (string, bool) {
	if o == nil || o.CouponPattern == nil {
		var ret string
		return ret, false
	}
	return *o.CouponPattern, true
}

// HasCouponPattern returns a boolean if a field has been set.
func (o *NewCouponsForMultipleRecipients) HasCouponPattern() bool {
	if o != nil && o.CouponPattern != nil {
		return true
	}

	return false
}

// SetCouponPattern gets a reference to the given string and assigns it to the CouponPattern field.
func (o *NewCouponsForMultipleRecipients) SetCouponPattern(v string) {
	o.CouponPattern = &v
}

type NullableNewCouponsForMultipleRecipients struct {
	Value        NewCouponsForMultipleRecipients
	ExplicitNull bool
}

func (v NullableNewCouponsForMultipleRecipients) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewCouponsForMultipleRecipients) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
