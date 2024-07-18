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

// NewCoupons struct for NewCoupons
type NewCoupons struct {
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
	// The number of new coupon codes to generate for the campaign. Must be at least 1.
	NumberOfCoupons int32 `json:"numberOfCoupons"`
	// **DEPRECATED** To create more than 20,000 coupons in one request, use [Create coupons asynchronously](https://docs.talon.one/management-api#operation/createCouponsAsync) endpoint. 
	UniquePrefix *string `json:"uniquePrefix,omitempty"`
	// Arbitrary properties associated with this item.
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	// The integration ID for this coupon's beneficiary's profile.
	RecipientIntegrationId *string `json:"recipientIntegrationId,omitempty"`
	// List of characters used to generate the random parts of a code. By default, the list of characters is equivalent to the `[A-Z, 0-9]` regular expression. 
	ValidCharacters *[]string `json:"validCharacters,omitempty"`
	// The pattern used to generate coupon codes. The character `#` is a placeholder and is replaced by a random character from the `validCharacters` set. 
	CouponPattern *string `json:"couponPattern,omitempty"`
	// An indication of whether the code can be redeemed only if it has been reserved first.
	IsReservationMandatory *bool `json:"isReservationMandatory,omitempty"`
	// An indication of whether the coupon is implicitly reserved for all customers.
	ImplicitlyReserved *bool `json:"implicitlyReserved,omitempty"`
}

// GetUsageLimit returns the UsageLimit field value
func (o *NewCoupons) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// SetUsageLimit sets field value
func (o *NewCoupons) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetDiscountLimit returns the DiscountLimit field value if set, zero value otherwise.
func (o *NewCoupons) GetDiscountLimit() float32 {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret
	}
	return *o.DiscountLimit
}

// GetDiscountLimitOk returns a tuple with the DiscountLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetDiscountLimitOk() (float32, bool) {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountLimit, true
}

// HasDiscountLimit returns a boolean if a field has been set.
func (o *NewCoupons) HasDiscountLimit() bool {
	if o != nil && o.DiscountLimit != nil {
		return true
	}

	return false
}

// SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.
func (o *NewCoupons) SetDiscountLimit(v float32) {
	o.DiscountLimit = &v
}

// GetReservationLimit returns the ReservationLimit field value if set, zero value otherwise.
func (o *NewCoupons) GetReservationLimit() int32 {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret
	}
	return *o.ReservationLimit
}

// GetReservationLimitOk returns a tuple with the ReservationLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetReservationLimitOk() (int32, bool) {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.ReservationLimit, true
}

// HasReservationLimit returns a boolean if a field has been set.
func (o *NewCoupons) HasReservationLimit() bool {
	if o != nil && o.ReservationLimit != nil {
		return true
	}

	return false
}

// SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.
func (o *NewCoupons) SetReservationLimit(v int32) {
	o.ReservationLimit = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *NewCoupons) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *NewCoupons) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *NewCoupons) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *NewCoupons) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *NewCoupons) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *NewCoupons) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *NewCoupons) GetLimits() []LimitConfig {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetLimitsOk() ([]LimitConfig, bool) {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret, false
	}
	return *o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *NewCoupons) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.
func (o *NewCoupons) SetLimits(v []LimitConfig) {
	o.Limits = &v
}

// GetNumberOfCoupons returns the NumberOfCoupons field value
func (o *NewCoupons) GetNumberOfCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfCoupons
}

// SetNumberOfCoupons sets field value
func (o *NewCoupons) SetNumberOfCoupons(v int32) {
	o.NumberOfCoupons = v
}

// GetUniquePrefix returns the UniquePrefix field value if set, zero value otherwise.
func (o *NewCoupons) GetUniquePrefix() string {
	if o == nil || o.UniquePrefix == nil {
		var ret string
		return ret
	}
	return *o.UniquePrefix
}

// GetUniquePrefixOk returns a tuple with the UniquePrefix field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetUniquePrefixOk() (string, bool) {
	if o == nil || o.UniquePrefix == nil {
		var ret string
		return ret, false
	}
	return *o.UniquePrefix, true
}

// HasUniquePrefix returns a boolean if a field has been set.
func (o *NewCoupons) HasUniquePrefix() bool {
	if o != nil && o.UniquePrefix != nil {
		return true
	}

	return false
}

// SetUniquePrefix gets a reference to the given string and assigns it to the UniquePrefix field.
func (o *NewCoupons) SetUniquePrefix(v string) {
	o.UniquePrefix = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NewCoupons) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NewCoupons) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *NewCoupons) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value if set, zero value otherwise.
func (o *NewCoupons) GetRecipientIntegrationId() string {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.RecipientIntegrationId
}

// GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetRecipientIntegrationIdOk() (string, bool) {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.RecipientIntegrationId, true
}

// HasRecipientIntegrationId returns a boolean if a field has been set.
func (o *NewCoupons) HasRecipientIntegrationId() bool {
	if o != nil && o.RecipientIntegrationId != nil {
		return true
	}

	return false
}

// SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.
func (o *NewCoupons) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = &v
}

// GetValidCharacters returns the ValidCharacters field value if set, zero value otherwise.
func (o *NewCoupons) GetValidCharacters() []string {
	if o == nil || o.ValidCharacters == nil {
		var ret []string
		return ret
	}
	return *o.ValidCharacters
}

// GetValidCharactersOk returns a tuple with the ValidCharacters field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetValidCharactersOk() ([]string, bool) {
	if o == nil || o.ValidCharacters == nil {
		var ret []string
		return ret, false
	}
	return *o.ValidCharacters, true
}

// HasValidCharacters returns a boolean if a field has been set.
func (o *NewCoupons) HasValidCharacters() bool {
	if o != nil && o.ValidCharacters != nil {
		return true
	}

	return false
}

// SetValidCharacters gets a reference to the given []string and assigns it to the ValidCharacters field.
func (o *NewCoupons) SetValidCharacters(v []string) {
	o.ValidCharacters = &v
}

// GetCouponPattern returns the CouponPattern field value if set, zero value otherwise.
func (o *NewCoupons) GetCouponPattern() string {
	if o == nil || o.CouponPattern == nil {
		var ret string
		return ret
	}
	return *o.CouponPattern
}

// GetCouponPatternOk returns a tuple with the CouponPattern field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetCouponPatternOk() (string, bool) {
	if o == nil || o.CouponPattern == nil {
		var ret string
		return ret, false
	}
	return *o.CouponPattern, true
}

// HasCouponPattern returns a boolean if a field has been set.
func (o *NewCoupons) HasCouponPattern() bool {
	if o != nil && o.CouponPattern != nil {
		return true
	}

	return false
}

// SetCouponPattern gets a reference to the given string and assigns it to the CouponPattern field.
func (o *NewCoupons) SetCouponPattern(v string) {
	o.CouponPattern = &v
}

// GetIsReservationMandatory returns the IsReservationMandatory field value if set, zero value otherwise.
func (o *NewCoupons) GetIsReservationMandatory() bool {
	if o == nil || o.IsReservationMandatory == nil {
		var ret bool
		return ret
	}
	return *o.IsReservationMandatory
}

// GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetIsReservationMandatoryOk() (bool, bool) {
	if o == nil || o.IsReservationMandatory == nil {
		var ret bool
		return ret, false
	}
	return *o.IsReservationMandatory, true
}

// HasIsReservationMandatory returns a boolean if a field has been set.
func (o *NewCoupons) HasIsReservationMandatory() bool {
	if o != nil && o.IsReservationMandatory != nil {
		return true
	}

	return false
}

// SetIsReservationMandatory gets a reference to the given bool and assigns it to the IsReservationMandatory field.
func (o *NewCoupons) SetIsReservationMandatory(v bool) {
	o.IsReservationMandatory = &v
}

// GetImplicitlyReserved returns the ImplicitlyReserved field value if set, zero value otherwise.
func (o *NewCoupons) GetImplicitlyReserved() bool {
	if o == nil || o.ImplicitlyReserved == nil {
		var ret bool
		return ret
	}
	return *o.ImplicitlyReserved
}

// GetImplicitlyReservedOk returns a tuple with the ImplicitlyReserved field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetImplicitlyReservedOk() (bool, bool) {
	if o == nil || o.ImplicitlyReserved == nil {
		var ret bool
		return ret, false
	}
	return *o.ImplicitlyReserved, true
}

// HasImplicitlyReserved returns a boolean if a field has been set.
func (o *NewCoupons) HasImplicitlyReserved() bool {
	if o != nil && o.ImplicitlyReserved != nil {
		return true
	}

	return false
}

// SetImplicitlyReserved gets a reference to the given bool and assigns it to the ImplicitlyReserved field.
func (o *NewCoupons) SetImplicitlyReserved(v bool) {
	o.ImplicitlyReserved = &v
}

type NullableNewCoupons struct {
	Value NewCoupons
	ExplicitNull bool
}

func (v NullableNewCoupons) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewCoupons) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
