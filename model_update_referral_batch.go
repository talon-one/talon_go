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

// UpdateReferralBatch struct for UpdateReferralBatch
type UpdateReferralBatch struct {
	// Arbitrary properties associated with this item.
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	// The id of the batch the referral belongs to.
	BatchID string `json:"batchID"`
	// Timestamp at which point the referral code becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiration date of the referral code. Referral never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of times a referral code can be used. This can be set to 0 for no limit, but any campaign usage limits will still apply. 
	UsageLimit *int32 `json:"usageLimit,omitempty"`
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateReferralBatch) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReferralBatch) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateReferralBatch) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *UpdateReferralBatch) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

// GetBatchID returns the BatchID field value
func (o *UpdateReferralBatch) GetBatchID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BatchID
}

// SetBatchID sets field value
func (o *UpdateReferralBatch) SetBatchID(v string) {
	o.BatchID = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UpdateReferralBatch) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReferralBatch) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UpdateReferralBatch) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *UpdateReferralBatch) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *UpdateReferralBatch) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReferralBatch) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *UpdateReferralBatch) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *UpdateReferralBatch) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *UpdateReferralBatch) GetUsageLimit() int32 {
	if o == nil || o.UsageLimit == nil {
		var ret int32
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReferralBatch) GetUsageLimitOk() (int32, bool) {
	if o == nil || o.UsageLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *UpdateReferralBatch) HasUsageLimit() bool {
	if o != nil && o.UsageLimit != nil {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.
func (o *UpdateReferralBatch) SetUsageLimit(v int32) {
	o.UsageLimit = &v
}

type NullableUpdateReferralBatch struct {
	Value UpdateReferralBatch
	ExplicitNull bool
}

func (v NullableUpdateReferralBatch) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateReferralBatch) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
