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

// NewCoupons
type NewCoupons struct {
	// The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.
	UsageLimit int32 `json:"usageLimit"`
	// The amount of discounts that can be given with this coupon code.
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of new coupon codes to generate for the campaign. Must be at least 1.
	NumberOfCoupons int32 `json:"numberOfCoupons"`
	// A unique prefix to prepend to all generated coupons.
	UniquePrefix *string `json:"uniquePrefix,omitempty"`
	// Arbitrary properties associated with this item
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The integration ID for this coupon's beneficiary's profile
	RecipientIntegrationId *string `json:"recipientIntegrationId,omitempty"`
	// Set of characters to be used when generating random part of code. Defaults to [A-Z, 0-9] (in terms of RegExp).
	ValidCharacters *[]string `json:"validCharacters,omitempty"`
	// The pattern that will be used to generate coupon codes. The character `#` acts as a placeholder and will be replaced by a random character from the `validCharacters` set.
	CouponPattern *string `json:"couponPattern,omitempty"`
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
func (o *NewCoupons) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewCoupons) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
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

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *NewCoupons) SetAttributes(v map[string]interface{}) {
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

type NullableNewCoupons struct {
	Value        NewCoupons
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
