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
)

// CreateAchievement struct for CreateAchievement
type CreateAchievement struct {
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
}

// GetName returns the Name field value
func (o *CreateAchievement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CreateAchievement) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *CreateAchievement) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *CreateAchievement) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *CreateAchievement) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *CreateAchievement) SetDescription(v string) {
	o.Description = v
}

// GetTarget returns the Target field value
func (o *CreateAchievement) GetTarget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Target
}

// SetTarget sets field value
func (o *CreateAchievement) SetTarget(v float32) {
	o.Target = v
}

// GetPeriod returns the Period field value
func (o *CreateAchievement) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// SetPeriod sets field value
func (o *CreateAchievement) SetPeriod(v string) {
	o.Period = v
}

// GetPeriodEndOverride returns the PeriodEndOverride field value if set, zero value otherwise.
func (o *CreateAchievement) GetPeriodEndOverride() TimePoint {
	if o == nil || o.PeriodEndOverride == nil {
		var ret TimePoint
		return ret
	}
	return *o.PeriodEndOverride
}

// GetPeriodEndOverrideOk returns a tuple with the PeriodEndOverride field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateAchievement) GetPeriodEndOverrideOk() (TimePoint, bool) {
	if o == nil || o.PeriodEndOverride == nil {
		var ret TimePoint
		return ret, false
	}
	return *o.PeriodEndOverride, true
}

// HasPeriodEndOverride returns a boolean if a field has been set.
func (o *CreateAchievement) HasPeriodEndOverride() bool {
	if o != nil && o.PeriodEndOverride != nil {
		return true
	}

	return false
}

// SetPeriodEndOverride gets a reference to the given TimePoint and assigns it to the PeriodEndOverride field.
func (o *CreateAchievement) SetPeriodEndOverride(v TimePoint) {
	o.PeriodEndOverride = &v
}

type NullableCreateAchievement struct {
	Value        CreateAchievement
	ExplicitNull bool
}

func (v NullableCreateAchievement) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateAchievement) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
