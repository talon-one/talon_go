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

// Tier struct for Tier
type Tier struct {
	// The internal ID of the tier.
	Id int32 `json:"id"`
	// The name of the tier.
	Name string `json:"name"`
	// Date and time when the customer moved to this tier. This value uses the loyalty program's time zone setting.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Date when tier level expires in the RFC3339 format (in the Loyalty Program's timezone).
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// Customers's tier downgrade policy. - `one_down`: Once the tier expires and if the user doesn't have enough points to stay in the tier, the user is downgraded one tier down. - `balance_based`: Once the tier expires, the user's tier is evaluated based on the amount of active points the user has at this instant. 
	DowngradePolicy *string `json:"downgradePolicy,omitempty"`
}

// GetId returns the Id field value
func (o *Tier) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Tier) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Tier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *Tier) SetName(v string) {
	o.Name = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Tier) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Tier) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Tier) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Tier) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *Tier) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Tier) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Tier) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *Tier) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetDowngradePolicy returns the DowngradePolicy field value if set, zero value otherwise.
func (o *Tier) GetDowngradePolicy() string {
	if o == nil || o.DowngradePolicy == nil {
		var ret string
		return ret
	}
	return *o.DowngradePolicy
}

// GetDowngradePolicyOk returns a tuple with the DowngradePolicy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Tier) GetDowngradePolicyOk() (string, bool) {
	if o == nil || o.DowngradePolicy == nil {
		var ret string
		return ret, false
	}
	return *o.DowngradePolicy, true
}

// HasDowngradePolicy returns a boolean if a field has been set.
func (o *Tier) HasDowngradePolicy() bool {
	if o != nil && o.DowngradePolicy != nil {
		return true
	}

	return false
}

// SetDowngradePolicy gets a reference to the given string and assigns it to the DowngradePolicy field.
func (o *Tier) SetDowngradePolicy(v string) {
	o.DowngradePolicy = &v
}

type NullableTier struct {
	Value Tier
	ExplicitNull bool
}

func (v NullableTier) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTier) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
