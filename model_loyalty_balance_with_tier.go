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
)

// LoyaltyBalanceWithTier struct for LoyaltyBalanceWithTier
type LoyaltyBalanceWithTier struct {
	// Total amount of points awarded to this customer and available to spend.
	ActivePoints *float32 `json:"activePoints,omitempty"`
	// Total amount of points awarded to this customer but not available until their start date.
	PendingPoints *float32 `json:"pendingPoints,omitempty"`
	// Total amount of points already spent by this customer.
	SpentPoints *float32 `json:"spentPoints,omitempty"`
	// Total amount of points awarded but never redeemed. They cannot be used anymore.
	ExpiredPoints *float32 `json:"expiredPoints,omitempty"`
	CurrentTier *Tier `json:"currentTier,omitempty"`
	ProjectedTier *ProjectedTier `json:"projectedTier,omitempty"`
	// The number of points required to move up a tier.
	PointsToNextTier *float32 `json:"pointsToNextTier,omitempty"`
	// The name of the tier consecutive to the current tier.
	NextTierName *string `json:"nextTierName,omitempty"`
}

// GetActivePoints returns the ActivePoints field value if set, zero value otherwise.
func (o *LoyaltyBalanceWithTier) GetActivePoints() float32 {
	if o == nil || o.ActivePoints == nil {
		var ret float32
		return ret
	}
	return *o.ActivePoints
}

// GetActivePointsOk returns a tuple with the ActivePoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalanceWithTier) GetActivePointsOk() (float32, bool) {
	if o == nil || o.ActivePoints == nil {
		var ret float32
		return ret, false
	}
	return *o.ActivePoints, true
}

// HasActivePoints returns a boolean if a field has been set.
func (o *LoyaltyBalanceWithTier) HasActivePoints() bool {
	if o != nil && o.ActivePoints != nil {
		return true
	}

	return false
}

// SetActivePoints gets a reference to the given float32 and assigns it to the ActivePoints field.
func (o *LoyaltyBalanceWithTier) SetActivePoints(v float32) {
	o.ActivePoints = &v
}

// GetPendingPoints returns the PendingPoints field value if set, zero value otherwise.
func (o *LoyaltyBalanceWithTier) GetPendingPoints() float32 {
	if o == nil || o.PendingPoints == nil {
		var ret float32
		return ret
	}
	return *o.PendingPoints
}

// GetPendingPointsOk returns a tuple with the PendingPoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalanceWithTier) GetPendingPointsOk() (float32, bool) {
	if o == nil || o.PendingPoints == nil {
		var ret float32
		return ret, false
	}
	return *o.PendingPoints, true
}

// HasPendingPoints returns a boolean if a field has been set.
func (o *LoyaltyBalanceWithTier) HasPendingPoints() bool {
	if o != nil && o.PendingPoints != nil {
		return true
	}

	return false
}

// SetPendingPoints gets a reference to the given float32 and assigns it to the PendingPoints field.
func (o *LoyaltyBalanceWithTier) SetPendingPoints(v float32) {
	o.PendingPoints = &v
}

// GetSpentPoints returns the SpentPoints field value if set, zero value otherwise.
func (o *LoyaltyBalanceWithTier) GetSpentPoints() float32 {
	if o == nil || o.SpentPoints == nil {
		var ret float32
		return ret
	}
	return *o.SpentPoints
}

// GetSpentPointsOk returns a tuple with the SpentPoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalanceWithTier) GetSpentPointsOk() (float32, bool) {
	if o == nil || o.SpentPoints == nil {
		var ret float32
		return ret, false
	}
	return *o.SpentPoints, true
}

// HasSpentPoints returns a boolean if a field has been set.
func (o *LoyaltyBalanceWithTier) HasSpentPoints() bool {
	if o != nil && o.SpentPoints != nil {
		return true
	}

	return false
}

// SetSpentPoints gets a reference to the given float32 and assigns it to the SpentPoints field.
func (o *LoyaltyBalanceWithTier) SetSpentPoints(v float32) {
	o.SpentPoints = &v
}

// GetExpiredPoints returns the ExpiredPoints field value if set, zero value otherwise.
func (o *LoyaltyBalanceWithTier) GetExpiredPoints() float32 {
	if o == nil || o.ExpiredPoints == nil {
		var ret float32
		return ret
	}
	return *o.ExpiredPoints
}

// GetExpiredPointsOk returns a tuple with the ExpiredPoints field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalanceWithTier) GetExpiredPointsOk() (float32, bool) {
	if o == nil || o.ExpiredPoints == nil {
		var ret float32
		return ret, false
	}
	return *o.ExpiredPoints, true
}

// HasExpiredPoints returns a boolean if a field has been set.
func (o *LoyaltyBalanceWithTier) HasExpiredPoints() bool {
	if o != nil && o.ExpiredPoints != nil {
		return true
	}

	return false
}

// SetExpiredPoints gets a reference to the given float32 and assigns it to the ExpiredPoints field.
func (o *LoyaltyBalanceWithTier) SetExpiredPoints(v float32) {
	o.ExpiredPoints = &v
}

// GetCurrentTier returns the CurrentTier field value if set, zero value otherwise.
func (o *LoyaltyBalanceWithTier) GetCurrentTier() Tier {
	if o == nil || o.CurrentTier == nil {
		var ret Tier
		return ret
	}
	return *o.CurrentTier
}

// GetCurrentTierOk returns a tuple with the CurrentTier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalanceWithTier) GetCurrentTierOk() (Tier, bool) {
	if o == nil || o.CurrentTier == nil {
		var ret Tier
		return ret, false
	}
	return *o.CurrentTier, true
}

// HasCurrentTier returns a boolean if a field has been set.
func (o *LoyaltyBalanceWithTier) HasCurrentTier() bool {
	if o != nil && o.CurrentTier != nil {
		return true
	}

	return false
}

// SetCurrentTier gets a reference to the given Tier and assigns it to the CurrentTier field.
func (o *LoyaltyBalanceWithTier) SetCurrentTier(v Tier) {
	o.CurrentTier = &v
}

// GetProjectedTier returns the ProjectedTier field value if set, zero value otherwise.
func (o *LoyaltyBalanceWithTier) GetProjectedTier() ProjectedTier {
	if o == nil || o.ProjectedTier == nil {
		var ret ProjectedTier
		return ret
	}
	return *o.ProjectedTier
}

// GetProjectedTierOk returns a tuple with the ProjectedTier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalanceWithTier) GetProjectedTierOk() (ProjectedTier, bool) {
	if o == nil || o.ProjectedTier == nil {
		var ret ProjectedTier
		return ret, false
	}
	return *o.ProjectedTier, true
}

// HasProjectedTier returns a boolean if a field has been set.
func (o *LoyaltyBalanceWithTier) HasProjectedTier() bool {
	if o != nil && o.ProjectedTier != nil {
		return true
	}

	return false
}

// SetProjectedTier gets a reference to the given ProjectedTier and assigns it to the ProjectedTier field.
func (o *LoyaltyBalanceWithTier) SetProjectedTier(v ProjectedTier) {
	o.ProjectedTier = &v
}

// GetPointsToNextTier returns the PointsToNextTier field value if set, zero value otherwise.
func (o *LoyaltyBalanceWithTier) GetPointsToNextTier() float32 {
	if o == nil || o.PointsToNextTier == nil {
		var ret float32
		return ret
	}
	return *o.PointsToNextTier
}

// GetPointsToNextTierOk returns a tuple with the PointsToNextTier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalanceWithTier) GetPointsToNextTierOk() (float32, bool) {
	if o == nil || o.PointsToNextTier == nil {
		var ret float32
		return ret, false
	}
	return *o.PointsToNextTier, true
}

// HasPointsToNextTier returns a boolean if a field has been set.
func (o *LoyaltyBalanceWithTier) HasPointsToNextTier() bool {
	if o != nil && o.PointsToNextTier != nil {
		return true
	}

	return false
}

// SetPointsToNextTier gets a reference to the given float32 and assigns it to the PointsToNextTier field.
func (o *LoyaltyBalanceWithTier) SetPointsToNextTier(v float32) {
	o.PointsToNextTier = &v
}

// GetNextTierName returns the NextTierName field value if set, zero value otherwise.
func (o *LoyaltyBalanceWithTier) GetNextTierName() string {
	if o == nil || o.NextTierName == nil {
		var ret string
		return ret
	}
	return *o.NextTierName
}

// GetNextTierNameOk returns a tuple with the NextTierName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyBalanceWithTier) GetNextTierNameOk() (string, bool) {
	if o == nil || o.NextTierName == nil {
		var ret string
		return ret, false
	}
	return *o.NextTierName, true
}

// HasNextTierName returns a boolean if a field has been set.
func (o *LoyaltyBalanceWithTier) HasNextTierName() bool {
	if o != nil && o.NextTierName != nil {
		return true
	}

	return false
}

// SetNextTierName gets a reference to the given string and assigns it to the NextTierName field.
func (o *LoyaltyBalanceWithTier) SetNextTierName(v string) {
	o.NextTierName = &v
}

type NullableLoyaltyBalanceWithTier struct {
	Value LoyaltyBalanceWithTier
	ExplicitNull bool
}

func (v NullableLoyaltyBalanceWithTier) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyBalanceWithTier) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}