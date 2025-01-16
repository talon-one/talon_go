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

// Achievement struct for Achievement
type Achievement struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The internal name of the achievement used in API requests.  **Note**: The name should start with a letter. This cannot be changed after the achievement has been created.
	Name string `json:"name"`
	// The display name for the achievement in the Campaign Manager.
	Title string `json:"title"`
	// A description of the achievement.
	Description string `json:"description"`
	// The required number of actions or the transactional milestone to complete the achievement.
	Target float32 `json:"target"`
	// The relative duration after which the achievement ends and resets for a particular customer profile.  **Note**: The `period` does not start when the achievement is created.  The period is a **positive real number** followed by one letter indicating the time unit.  Examples: `30s`, `40m`, `1h`, `5D`, `7W`, `10M`, `15Y`.  Available units:  - `s`: seconds - `m`: minutes - `h`: hours - `D`: days - `W`: weeks - `M`: months - `Y`: years  You can also round certain units down to the beginning of period and up to the end of period.: - `_D` for rounding down days only. Signifies the start of the day. Example: `30D_D` - `_U` for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year. Example: `23W_U`  **Note**: You can either use the round down and round up option or set an absolute period.
	Period            string     `json:"period"`
	PeriodEndOverride *TimePoint `json:"periodEndOverride,omitempty"`
	// The policy that determines if and how the achievement recurs. - `no_recurrence`: The achievement can be completed only once. - `on_expiration`: The achievement resets after it expires and becomes available again.
	RecurrencePolicy *string `json:"recurrencePolicy,omitempty"`
	// The policy that determines how the achievement starts, ends, or resets. - `user_action`: The achievement ends or resets relative to when the customer started the achievement. - `fixed_schedule`: The achievement starts, ends, or resets for all customers following a fixed schedule.
	ActivationPolicy *string `json:"activationPolicy,omitempty"`
	// The achievement's start date when `activationPolicy` is set to `fixed_schedule`.  **Note:** It must be an RFC3339 timestamp string.
	FixedStartDate *time.Time `json:"fixedStartDate,omitempty"`
	// The achievement's end date. If defined, customers cannot participate in the achievement after this date.  **Note:** It must be an RFC3339 timestamp string.
	EndDate *time.Time `json:"endDate,omitempty"`
	// ID of the campaign, to which the achievement belongs to
	CampaignId int32 `json:"campaignId"`
	// ID of the user that created this achievement.
	UserId int32 `json:"userId"`
	// Name of the user that created the achievement.  **Note**: This is not available if the user has been deleted.
	CreatedBy string `json:"createdBy"`
	// Indicates if a customer has made progress in the achievement.
	HasProgress *bool `json:"hasProgress,omitempty"`
}

// GetId returns the Id field value
func (o *Achievement) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Achievement) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Achievement) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Achievement) SetCreated(v time.Time) {
	o.Created = v
}

// GetName returns the Name field value
func (o *Achievement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *Achievement) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *Achievement) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *Achievement) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *Achievement) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *Achievement) SetDescription(v string) {
	o.Description = v
}

// GetTarget returns the Target field value
func (o *Achievement) GetTarget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Target
}

// SetTarget sets field value
func (o *Achievement) SetTarget(v float32) {
	o.Target = v
}

// GetPeriod returns the Period field value
func (o *Achievement) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// SetPeriod sets field value
func (o *Achievement) SetPeriod(v string) {
	o.Period = v
}

// GetPeriodEndOverride returns the PeriodEndOverride field value if set, zero value otherwise.
func (o *Achievement) GetPeriodEndOverride() TimePoint {
	if o == nil || o.PeriodEndOverride == nil {
		var ret TimePoint
		return ret
	}
	return *o.PeriodEndOverride
}

// GetPeriodEndOverrideOk returns a tuple with the PeriodEndOverride field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Achievement) GetPeriodEndOverrideOk() (TimePoint, bool) {
	if o == nil || o.PeriodEndOverride == nil {
		var ret TimePoint
		return ret, false
	}
	return *o.PeriodEndOverride, true
}

// HasPeriodEndOverride returns a boolean if a field has been set.
func (o *Achievement) HasPeriodEndOverride() bool {
	if o != nil && o.PeriodEndOverride != nil {
		return true
	}

	return false
}

// SetPeriodEndOverride gets a reference to the given TimePoint and assigns it to the PeriodEndOverride field.
func (o *Achievement) SetPeriodEndOverride(v TimePoint) {
	o.PeriodEndOverride = &v
}

// GetRecurrencePolicy returns the RecurrencePolicy field value if set, zero value otherwise.
func (o *Achievement) GetRecurrencePolicy() string {
	if o == nil || o.RecurrencePolicy == nil {
		var ret string
		return ret
	}
	return *o.RecurrencePolicy
}

// GetRecurrencePolicyOk returns a tuple with the RecurrencePolicy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Achievement) GetRecurrencePolicyOk() (string, bool) {
	if o == nil || o.RecurrencePolicy == nil {
		var ret string
		return ret, false
	}
	return *o.RecurrencePolicy, true
}

// HasRecurrencePolicy returns a boolean if a field has been set.
func (o *Achievement) HasRecurrencePolicy() bool {
	if o != nil && o.RecurrencePolicy != nil {
		return true
	}

	return false
}

// SetRecurrencePolicy gets a reference to the given string and assigns it to the RecurrencePolicy field.
func (o *Achievement) SetRecurrencePolicy(v string) {
	o.RecurrencePolicy = &v
}

// GetActivationPolicy returns the ActivationPolicy field value if set, zero value otherwise.
func (o *Achievement) GetActivationPolicy() string {
	if o == nil || o.ActivationPolicy == nil {
		var ret string
		return ret
	}
	return *o.ActivationPolicy
}

// GetActivationPolicyOk returns a tuple with the ActivationPolicy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Achievement) GetActivationPolicyOk() (string, bool) {
	if o == nil || o.ActivationPolicy == nil {
		var ret string
		return ret, false
	}
	return *o.ActivationPolicy, true
}

// HasActivationPolicy returns a boolean if a field has been set.
func (o *Achievement) HasActivationPolicy() bool {
	if o != nil && o.ActivationPolicy != nil {
		return true
	}

	return false
}

// SetActivationPolicy gets a reference to the given string and assigns it to the ActivationPolicy field.
func (o *Achievement) SetActivationPolicy(v string) {
	o.ActivationPolicy = &v
}

// GetFixedStartDate returns the FixedStartDate field value if set, zero value otherwise.
func (o *Achievement) GetFixedStartDate() time.Time {
	if o == nil || o.FixedStartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.FixedStartDate
}

// GetFixedStartDateOk returns a tuple with the FixedStartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Achievement) GetFixedStartDateOk() (time.Time, bool) {
	if o == nil || o.FixedStartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.FixedStartDate, true
}

// HasFixedStartDate returns a boolean if a field has been set.
func (o *Achievement) HasFixedStartDate() bool {
	if o != nil && o.FixedStartDate != nil {
		return true
	}

	return false
}

// SetFixedStartDate gets a reference to the given time.Time and assigns it to the FixedStartDate field.
func (o *Achievement) SetFixedStartDate(v time.Time) {
	o.FixedStartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Achievement) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Achievement) GetEndDateOk() (time.Time, bool) {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Achievement) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *Achievement) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetCampaignId returns the CampaignId field value
func (o *Achievement) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *Achievement) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetUserId returns the UserId field value
func (o *Achievement) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// SetUserId sets field value
func (o *Achievement) SetUserId(v int32) {
	o.UserId = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Achievement) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// SetCreatedBy sets field value
func (o *Achievement) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetHasProgress returns the HasProgress field value if set, zero value otherwise.
func (o *Achievement) GetHasProgress() bool {
	if o == nil || o.HasProgress == nil {
		var ret bool
		return ret
	}
	return *o.HasProgress
}

// GetHasProgressOk returns a tuple with the HasProgress field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Achievement) GetHasProgressOk() (bool, bool) {
	if o == nil || o.HasProgress == nil {
		var ret bool
		return ret, false
	}
	return *o.HasProgress, true
}

// HasHasProgress returns a boolean if a field has been set.
func (o *Achievement) HasHasProgress() bool {
	if o != nil && o.HasProgress != nil {
		return true
	}

	return false
}

// SetHasProgress gets a reference to the given bool and assigns it to the HasProgress field.
func (o *Achievement) SetHasProgress(v bool) {
	o.HasProgress = &v
}

type NullableAchievement struct {
	Value        Achievement
	ExplicitNull bool
}

func (v NullableAchievement) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAchievement) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
